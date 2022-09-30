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

func (client *Client) UserSendCount() (result UserLimitTimeResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(UserSendCount)
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

// MpList  获取微信公众号渠道列表
// page 当前所在分页数, pageSize 每页大小最大值为50
func (client *Client) MpList(page, pageSize int) (list MpListResp, err error) {
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
		Post(WebhookMpList)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// CpList  获取企业微信渠道列表
// page 当前所在分页数, pageSize 每页大小最大值为50
func (client *Client) CpList(page, pageSize int) (list CpListResp, err error) {
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
		Post(WebhookCpList)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// MailList  获取邮箱渠道列表
// page 当前所在分页数, pageSize 每页大小最大值为50
func (client *Client) MailList(page, pageSize int) (list MailListResp, err error) {
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
		Post(WebhookMailList)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// MailDetail  获取邮箱渠道详情
// mailID 邮箱编号
func (client *Client) MailDetail(mailID int) (detail MailDetailResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(WebhookMailDetail + Int2String(mailID))
	if err != nil {
		return
	}
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

// ChangeIsSend 开启/关闭发送消息功能
// isSend 发送消息功能；0-禁用，1-启用
func (client *Client) ChangeIsSend(isSend int) (detail IntResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(ChangeIsSend + Int2String(isSend))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// ChangeOpenMessageType 修改打开消息方式
// msgType 消息打开类型；0:H5，1:小程序
func (client *Client) ChangeOpenMessageType(msgType int) (detail IntResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(changeOpenMessageType + Int2String(msgType))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// FriendGetQRCode 获取个人二维码
func (client *Client) FriendGetQRCode() (detail FriendQRCode, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(FriendGetQrCode)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// FriendList  获取好友列表
// page 当前所在分页数, pageSize 每页大小最大值为50
func (client *Client) FriendList(page, pageSize int) (list FriendListResp, err error) {
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
		Post(FriendList)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&list)
	if err != nil {
		return
	}
	return
}

// FriendDel 删除好友
// id 好友id
func (client *Client) FriendDel(id int) (detail IntResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}

	resp, err := client.Req.R().
		Get(DeleteFriend + Int2String(id))
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}

// FriendEditRemark  修改好友备注
// id 好友编号, remark 好友备注
func (client *Client) FriendEditRemark(id int, remark string) (detail IntResp, err error) {
	err = client.Auth()
	if err != nil {
		return
	}
	param := FriendData{
		Id:     id,
		Remark: remark,
	}
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(FriendEditRemark)
	if err != nil {
		return
	}

	err = resp.UnmarshalJson(&detail)
	if err != nil {
		return
	}
	return
}
