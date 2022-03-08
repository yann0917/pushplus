package pushplus

// SendMsg 发送消息
func (client *Client) SendMsg(title, content string) (result Resp, err error) {
	param := SendMsgReq{
		Token:       client.Token,
		Title:       title,
		Content:     content,
		Template:    client.Template,
		CallbackUrl: client.Callback,
		Channel:     client.Channel,
		Webhook:     client.Webhook,
	}
	resp, err := client.Req.R().
		SetBodyJsonMarshal(param).
		Post(SendMsgURL)
	if err != nil {
		return
	}
	err = resp.UnmarshalJson(&result)
	return
}
