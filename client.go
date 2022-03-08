package pushplus

import (
	"errors"
	"time"

	"github.com/imroc/req/v3"
)

var (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
	BaseURL   = "https://www.pushplus.plus"
)

const (
	SendMsgURL      = "/send"
	GetAccessKeyURL = "/api/common/openApi/getAccessKey"

	MsgListURL   = "/api/open/message/list"
	MsgResultURL = "/api/open/message/sendMessageResult?shortCode="

	UserTokenURL     = "/api/open/user/token"
	UserInfoURL      = "/api/open/user/myInfo"
	UserLimitTimeURL = "/api/open/user/userLimitTime"

	TopicAddURL        = "/api/open/topic/add"
	TopicListURL       = "/api/open/topic/list"
	ExitTopicURL       = "/api/open/topic/exitTopic?topicId="
	TopicDetailURL     = "/api/open/topic/detail?topicId="
	JoinTopicDetailURL = "/api/open/topic/joinTopicDetail?topicId="
	TopicQRCode        = "/api/open/topic/qrCode?topicId=" // 二维码类型；0-临时二维码，1-永久二维码

	SubscriberListURL  = "/api/open/topicUser/subscriberList"
	DeleteTopicUserURL = "/api/open/topicUser/deleteTopicUser?topicRelationId="

	WebhookListURL   = "/api/open/webhook/list"
	WebhookDetailURL = "/api/open/webhook/detail?webhookId="
	WebhookAddURL    = "/api/open/webhook/add"
	WebhookEditURL   = "/api/open/webhook/edit"

	GetUserSettingsURL      = "/api/open/setting/getUserSettings"
	ChangeDefaultChannelURL = "/api/open/setting/changeDefaultChannel"
	ChangeReceiveLimitURL   = "/api/open/setting/changeRecevieLimit?recevieLimit="
)

type Authorizer interface {
	Authorize(*Client) error
}

type DefaultAuthorizer struct{}

type Client struct {
	Req        *req.Client
	AccessKey  string
	expireAt   time.Time
	Authorizer Authorizer
	Token      string
	SecretKey  string
	Template   string
	Channel    string
	Webhook    string
	Callback   string
	DebugLog   bool
}

type Option func(*Client)

func NewClient(token string, options ...Option) *Client {
	c := Client{
		Req: req.C().
			SetBaseURL(BaseURL).
			SetUserAgent(UserAgent),
		Token:      token,
		Template:   "html",
		Channel:    "wechat",
		Authorizer: DefaultAuthorizer{},
	}

	for _, option := range options {
		option(&c)
	}
	return &c
}

func SecretKey(secretKey string) Option {
	return func(client *Client) {
		client.SecretKey = secretKey
	}
}

func Template(tpl string) Option {
	return func(client *Client) {
		client.Template = tpl
	}
}

func DebugLog(isDebug bool) Option {
	return func(client *Client) {
		if isDebug {
			client.Req = client.Req.EnableDumpAllAsync()
		}
		client.DebugLog = isDebug
	}
}

func Channel(ch string) Option {
	return func(client *Client) {
		client.Channel = ch
	}
}

func Webhook(hook string) Option {
	return func(client *Client) {
		client.Webhook = hook
	}
}

func Callback(url string) Option {
	return func(client *Client) {
		client.Callback = url
	}
}

func (da DefaultAuthorizer) Authorize(client *Client) (err error) {
	param := AuthReq{
		Token:     client.Token,
		SecretKey: client.SecretKey,
	}
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(GetAccessKeyURL)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		err = errors.New("HTTP Failed")
		return
	}

	var authResp AuthResp
	err = resp.UnmarshalJson(&authResp)
	if err != nil {
		return
	}

	if authResp.Data.AccessKey == "" {
		err = errors.New("授权失败:" + authResp.Msg)
		return
	}

	client.AccessKey = authResp.Data.AccessKey
	expireIn := authResp.Data.ExpiresIn
	// 记录 accessKey 过期时间，过期时间前60秒就触发刷新
	client.expireAt = time.Now().Add(time.Second * time.Duration(expireIn-60))
	return
}

func (client *Client) Auth() error {

	if client.AccessKey != "" && time.Now().Before(client.expireAt) {
		return nil
	}

	if err := client.Authorizer.Authorize(client); err != nil {
		return err
	}
	client.Req = client.Req.SetCommonHeader("access-key", client.AccessKey)
	return nil
}
