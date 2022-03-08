package pushplus

// MessageList 消息列表
// page 当前所在分页数, pageSize 每页大小最大值为50
func (client *Client) MessageList(page, pageSize int) (list MsgListResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}
	param := PageReq{
		Current:  page,
		PageSize: pageSize,
	}
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(MsgListURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// SendMessageResult 查询消息发送结果
// shortCode 消息短链码
func (client *Client) SendMessageResult(shortCode string) (result SendMsgResultResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(MsgResultURL + shortCode)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

func (client *Client) GetUserToken() (result Resp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(UserTokenURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

func (client *Client) GetUserInfo() (result UserInfo, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(UserInfoURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

func (client *Client) GetUserLimitTime() (result UserLimitTimeResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(UserLimitTimeURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

// TopicList  群组列表
// page 当前所在分页数, pageSize 每页大小最大值为50, topicType 群组类型: 0-我创建的，1-我加入的
func (client *Client) TopicList(page, pageSize, topicType int) (list TopicListResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}
	param := TopicListReq{}
	param.Current = page
	param.PageSize = pageSize
	param.Params.TopicType = topicType
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(TopicListURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// TopicDetail  获取我创建的群组详情
// topicID 群组 ID
func (client *Client) TopicDetail(topicID int) (detail TopicDetailResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(TopicDetailURL + Int2String(topicID))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// TopicJoinDetail  获取我加入的群详情
// topicID 群组 ID
func (client *Client) TopicJoinDetail(topicID int) (detail TopicJoinDetailResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(JoinTopicDetailURL + Int2String(topicID))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// TopicAdd  新增群组
func (client *Client) TopicAdd(req TopicAddReq) (result TopicAddResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		SetBodyJsonMarshal(req).
		Post(TopicAddURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

// GetTopicQrCode  获取群组二维码
// topicID 群组 ID
// forever 二维码类型；0-临时二维码，1-永久二维码
func (client *Client) GetTopicQrCode(topicID, forever int) (detail QrCodeResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(TopicQRCode + Int2String(topicID) + "&forever=" + Int2String(forever))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// ExistTopic  退出群组
// topicID 群组 ID
func (client *Client) ExistTopic(topicID int) (detail Resp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(ExitTopicURL + Int2String(topicID))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// TopicUserList  群组内用户列表
// page 当前所在分页数, pageSize 每页大小最大值为50, topicType 群组类型: 0-我创建的，1-我加入的
func (client *Client) TopicUserList(page, pageSize, topicID int) (list TopicUserListResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}
	param := TopicUserListReq{}
	param.Current = page
	param.PageSize = pageSize
	param.Params.TopicId = topicID
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(SubscriberListURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// DeleteTopicUser  删除群组内用户
// id 用户编号
func (client *Client) DeleteTopicUser(id int) (result Resp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Post(DeleteTopicUserURL + Int2String(id))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

// WebhookList  Webhook列表
// page 当前所在分页数, pageSize 每页大小最大值为50
func (client *Client) WebhookList(page, pageSize int) (list WebhookListResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}
	param := PageReq{
		Current:  page,
		PageSize: pageSize,
	}
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(WebhookListURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// WebhookAdd  新增Webhook
func (client *Client) WebhookAdd(req WebhookReq) (result TopicAddResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		SetBodyJsonMarshal(req).
		Post(WebhookAddURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

// WebhookEdit  编辑Webhook
func (client *Client) WebhookEdit(req WebhookData) (result Resp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		SetBodyJsonMarshal(req).
		Post(WebhookEditURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&result)
	if err != nil {
		return
	}
	return
}

// WebhookDetail  获取我加入的群详情
// topicID 群组 ID
func (client *Client) WebhookDetail(id int) (detail WebhookResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(WebhookDetailURL + Int2String(id))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

func (client *Client) GetUserSettings() (detail UserSetting, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(GetUserSettingsURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// ChangeDefaultChannel 修改默认发送渠道
// channel 默认渠道；wechat-微信公众号,mail-邮件,cp-企业微信应用,webhook-第三方webhook
// webhook 渠道参数；webhook和cp渠道需要填写具体的webhook编号或自定义编码
func (client *Client) ChangeDefaultChannel(channel, webhook string) (detail Resp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	param := struct {
		DefaultChannel string `json:"defaultChannel"`
		DefaultWebhook string `json:"defaultWebhook"`
	}{
		DefaultChannel: channel,
		DefaultWebhook: webhook,
	}

	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(ChangeDefaultChannelURL)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// ChangeReceiveLimit 修改接收消息限制
// limit 接收消息限制；0-接收全部，1-不接收消息
func (client *Client) ChangeReceiveLimit(limit int) (detail Resp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(ChangeReceiveLimitURL + Int2String(limit))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}
