package pushplus

type SendMsgReq struct {
	Token       string `json:"token"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content"`
	Template    string `json:"template,omitempty"` // 消息模板 html,json,cloudMonitor 默认 html
	Topic       string `json:"topic,omitempty"`    // 群组编码，不填仅发送给自己；channel为webhook时无效
	Channel     string `json:"channel,omitempty"`  // 发送渠道, wechat,webhook,cp 默认 wechat
	Webhook     string `json:"webhook,omitempty"`
	CallbackUrl string `json:"callback_url,omitempty"` // 发送结果回调地址
	Timestamp   string `json:"timestamp,omitempty"`    // 毫秒时间戳。格式如：1632993318000。服务器时间戳大于此时间戳，则消息不会发送
	To          string `json:"to,omitempty"`           // 好友令牌，微信公众号渠道填写好友令牌，企业微信渠道填写企业微信用户id
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type IntResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

type SendMsgCallbackResp struct {
	ShortCode  string `json:"shortCode"`  // 消息流水号
	SendStatus int    `json:"sendStatus"` // 发送状态；0未发送，1发送中，2发送成功，3发送失败
	Message    string `json:"message"`    // 推送错误内容（如有）
}

type AuthReq struct {
	Token     string `json:"token,omitempty"`     // 用户令牌，同发送消息接口令牌
	SecretKey string `json:"secretKey,omitempty"` // 用户密钥
}

type AuthResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessKey string `json:"accessKey"` // 要获取的 accessKey
		ExpiresIn int64  `json:"expiresIn"` // accessKey 的有效期(秒为单位，7200秒之内的值)；
	} `json:"data"`
}

type PageReq struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}
type Pager struct {
	PageNum  int `json:"pageNum"`  // 当前页码
	PageSize int `json:"pageSize"` // 分页大小
	Total    int `json:"total"`    // 总行数
	Pages    int `json:"pages"`    // 总页数
}

type MsgListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []struct {
			TopicName   string `json:"topicName"`   // 群组名称，一对多消息才有值
			MessageType int    `json:"messageType"` // 消息类型;1-一对一消息,2-一对多消息
			Title       string `json:"title"`       // 消息标题
			ShortCode   string `json:"shortCode"`   // 消息短链码;可用于查询消息发送结果
			Channel     string `json:"channel"`     // 消息发送渠道； wechat-微信公众号,mail-邮件,cp-企业微信应用,webhook-第三方webhook
			UpdateTime  string `json:"updateTime"`  // 更新日期
		} `json:"list"`
	} `json:"data"`
}

type SendMsgResultResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Status       int    `json:"status"`       // 消息投递状态；0-未投递，1-发送中，2-已发送，3-发送失败
		ErrorMessage string `json:"errorMessage"` // 发送失败原因
		UpdateTime   string `json:"updateTime"`
	} `json:"data"`
}

type UserInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OpenId      string  `json:"openId"`      // 用户微信的openId
		UnionId     string  `json:"unionId"`     // 用户微信的unionId
		NickName    string  `json:"nickName"`    // 昵称
		HeadImgUrl  string  `json:"headImgUrl"`  // 头像
		UserSex     int     `json:"userSex"`     // 性别: 0-未设置，1-男，2-女
		Token       string  `json:"token"`       // 用户令牌
		PhoneNumber string  `json:"phoneNumber"` // 手机号
		Email       string  `json:"email"`       // 邮箱
		EmailStatus int     `json:"emailStatus"` // 邮箱验证状态；0-未验证，1-待验证，2-已验证
		Birthday    string  `json:"birthday"`    // 生日
		Points      float64 `json:"points"`      // 用户积分
	} `json:"data"`
}

type UserLimitTimeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		SendLimit     int    `json:"sendLimit"`     // 发送限制状态;1-无限制，2-短期限制，3-永久限制
		UserLimitTime string `json:"userLimitTime"` // 解封时间
	} `json:"data"`
}

type UserSendCountResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		WechatSendCount  int `json:"wechatSendCount"`  // 微信公众号渠道请求次数
		CpSendCount      int `json:"cpSendCount"`      // 企业微信应用渠道请求次数
		WebhookSendCount int `json:"webhookSendCount"` // webhook渠道请求次数
		MailSendCount    int `json:"mailSendCount"`    // 邮件渠道请求次数
	} `json:"data"`
}
type TopicListReq struct {
	PageReq
	Params struct {
		TopicType int `json:"topicType"` // 群组类型;0-我创建的，1-我加入的
	} `json:"params"`
}

type TopicListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []struct {
			TopicId    int    `json:"topicId"`    // 群组编号
			TopicCode  string `json:"topicCode"`  // 群组编码
			TopicName  string `json:"topicName"`  // 群组名称
			CreateTime string `json:"createTime"` // 创建时间
		} `json:"list"`
	} `json:"data"`
}

type TopicDetailResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TopicId        int    `json:"topicId"`
		TopicName      string `json:"topicName"`
		TopicCode      string `json:"topicCode"`
		QrCodeImgUrl   string `json:"qrCodeImgUrl"`
		Contact        string `json:"contact"`
		Introduction   string `json:"introduction"`
		ReceiptMessage string `json:"receiptMessage"`
		CreateTime     string `json:"createTime"`
		TopicUserCount int    `json:"topicUserCount"`
	} `json:"data"`
}

type TopicJoinDetailResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TopicName    string `json:"topicName"`
		TopicCode    string `json:"topicCode"`
		TopicId      int    `json:"topicId"`
		Contact      string `json:"contact"`
		Introduction string `json:"introduction"`
		CreateTime   string `json:"createTime"`
	} `json:"data"`
}

type TopicAddReq struct {
	TopicCode      string `json:"topicCode"`      // 群组编码
	TopicName      string `json:"topicName"`      // 群组名称
	Contact        string `json:"contact"`        // 联系方式
	Introduction   string `json:"introduction"`   // 群组简介
	ReceiptMessage string `json:"receiptMessage"` // 加入后回复内容
}

type TopicAddResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"` // 新建群组的群组编号
}

type QrCodeResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		QrCodeImgUrl string `json:"qrCodeImgUrl"` // 群组二维码图片路径
		Forever      int    `json:"forever"`      // 二维码类型；0-临时二维码，1-永久二维码
	} `json:"data"`
}

type TopicUserListReq struct {
	PageReq
	Params struct {
		TopicId int `json:"topicId"`
	} `json:"params"`
}

type TopicUserListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []struct {
			Id          int    `json:"id"`          // 用户编号；可用于删除用户
			NickName    string `json:"nickName"`    // 昵称
			OpenId      string `json:"openId"`      // 用户微信openId
			HeadImgUrl  string `json:"headImgUrl"`  // 头像url地址
			UserSex     int    `json:"userSex"`     // 性别；0-未设置，1-男，2-女
			HavePhone   int    `json:"havePhone"`   // 是否绑定手机；0-未绑定，1-已绑定
			IsFollow    int    `json:"isFollow"`    // 是否关注微信公众号；0-未关注，1-已关注
			EmailStatus int    `json:"emailStatus"` // 邮箱验证状态；0-未验证，1-待验证，2-已验证
		} `json:"list"`
	} `json:"data"`
}

type WebhookListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []WebhookData `json:"list"`
	} `json:"data"`
}

type WebhookResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data WebhookData `json:"data"`
}

type WebhookData struct {
	Id          int    `json:"id"`
	WebhookCode string `json:"webhookCode"` // 编码
	WebhookName string `json:"webhookName"` // 名称
	WebhookType int    `json:"webhookType"` // 类型；1-企业微信，2-钉钉，3-飞书，4-server酱
	WebhookUrl  string `json:"webhookUrl"`  // 调用的url地址
	CreateTime  string `json:"createTime"`
}

type WebhookReq struct {
	WebhookCode string `json:"webhookCode"` // 编码
	WebhookName string `json:"webhookName"` // 名称
	WebhookType int    `json:"webhookType"` // 类型；1-企业微信，2-钉钉，3-飞书，4-server酱
	WebhookUrl  string `json:"webhookUrl"`  // 调用的url地址
}

type MpListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []MpData `json:"list"`
	} `json:"data"`
}

type MpData struct {
	Id                 int    `json:"id"`                 // 微信公众号编号
	NickName           string `json:"nickName"`           // 微信公众号名称
	HeadImg            string `json:"headImg"`            // 微信公众号头像
	PrincipalName      string `json:"principalName"`      // 公众号的主体名称
	AuthorizationAppid string `json:"authorizationAppid"` // 授权方appid
	FuncInfo           string `json:"funcInfo"`           // 权限集列表
	ServiceType        int    `json:"serviceType"`        // 授权方公众号类型，0 代表订阅号，1 代表由历史老账号升级后的订阅号，2 代表服务号
	VerifyType         int    `json:"verifyType"`         // 授权方认证类型，-1 代表未认证，0 代表微信认证
	Alias              string `json:"alias"`              // 公众号所设置的微信号
	UpdateTime         string `json:"updateTime"`
}

type CpListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []CpData `json:"list"`
	} `json:"data"`
}

type CpData struct {
	Id     int    `json:"id"`     // 企业微信应用编号
	CpName string `json:"cpName"` // 企业微信应用名称
	CpCode string `json:"cpCode"` // 企业微信应用编码
}

type MailListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []struct {
			Id       int    `json:"id"`
			MailName string `json:"mailName"`
			MailCode string `json:"mailCode"`
		} `json:"list"`
	} `json:"data"`
}

type MailDetailResp struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data MailData `json:"data"`
}
type MailData struct {
	Id         int    `json:"id"`       // 邮箱编号
	MailName   string `json:"mailName"` // 邮箱名称
	MailCode   string `json:"mailCode"` // 邮箱编码
	Account    string `json:"account,omitempty"`
	Password   string `json:"password,omitempty"`
	SmtpServer string `json:"smtpServer,omitempty"`
	SmtpSsl    int    `json:"smtpSsl,omitempty"`
	SmtpPort   int    `json:"smtpPort,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
}

type UserSetting struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DefaultChannel    string `json:"defaultChannel"`    // 默认渠道编码
		DefaultChannelTxt string `json:"defaultChannelTxt"` // 默认渠道名称
		DefaultWebhook    string `json:"defaultWebhook"`    // webhook 参数
		SendLimit         int    `json:"sendLimit"`         // 发送限制；0-无限制，1-禁止所有渠道发送，2-限制微信渠道，3-限制邮件渠道
		RecevieLimit      int    `json:"recevieLimit"`      // 接收限制；0-接收全部，1-不接收消息

	} `json:"data"`
}

type FriendQRCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		QrCodeImgUrl string `json:"qrCodeImgUrl"`
	} `json:"data"`
}

type FriendListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Pager
		List []FriendData `json:"list"`
	} `json:"data"`
}

type FriendData struct {
	Id          int    `json:"id"`          // 好友编号
	FriendId    int    `json:"friendId"`    // 好友id
	Token       string `json:"token"`       // 好友令牌；发送好友消息使用
	HeadImgUrl  string `json:"headImgUrl"`  // 好友头像
	NickName    string `json:"nickName"`    // 好友昵称
	EmailStatus int    `json:"emailStatus"` // 邮箱验证状态；0-未验证，1-待验证，2-已验证
	HavePhone   int    `json:"havePhone"`   // 是否绑定手机；0-未绑定，1-已绑定
	IsFollow    int    `json:"isFollow"`    // 是否关注微信公众号；0-未关注，1-已关注
	Remark      string `json:"remark"`      // 备注
	CreateTime  string `json:"createTime"`  // 创建日期
}

type EditFriendReq struct {
	Id     int    `json:"id"`     // 好友编号
	Remark string `json:"remark"` // 备注
}
