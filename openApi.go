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
	err = client.EnableDebugLog().Auth()
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
