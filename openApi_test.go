package pushplus

import (
	"testing"
)

func TestMessageList(t *testing.T) {
	result, err := HttpClient.MessageList(1, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendMessageResult(t *testing.T) {
	result, err := HttpClient.SendMessageResult("5e9a4f36d04e49c2b323f5b87c8b5a2d")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetUserToken(t *testing.T) {
	result, err := HttpClient.GetUserToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetUserInfo(t *testing.T) {
	result, err := HttpClient.GetUserInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetUserLimitTime(t *testing.T) {
	result, err := HttpClient.GetUserLimitTime()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestUserSendCount(t *testing.T) {
	result, err := HttpClient.UserSendCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestTopicList(t *testing.T) {
	result, err := HttpClient.TopicList(1, 20, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_TopicAdd(t *testing.T) {
	req := TopicAddReq{
		TopicCode:      "pushplus",
		TopicName:      "test",
		Contact:        "me",
		Introduction:   "Brief",
		ReceiptMessage: "欢迎",
	}
	result, err := HttpClient.TopicAdd(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_TopicDetail(t *testing.T) {
	result, err := HttpClient.TopicDetail(19821)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_TopicJoinDetail(t *testing.T) {
	result, err := HttpClient.TopicJoinDetail(19821)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_GetTopicQrCode(t *testing.T) {
	result, err := HttpClient.GetTopicQrCode(19821, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_ExistTopic(t *testing.T) {
	result, err := HttpClient.ExistTopic(19821)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_TopicUserList(t *testing.T) {
	result, err := HttpClient.TopicUserList(1, 20, 19821)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_DeleteTopicUser(t *testing.T) {
	result, err := HttpClient.DeleteTopicUser(107396)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_WebhookList(t *testing.T) {
	result, err := HttpClient.WebhookList(1, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_WebhookAdd(t *testing.T) {
	req := WebhookReq{
		WebhookCode: "",
		WebhookName: "",
		WebhookType: 0,
		WebhookUrl:  "",
	}
	result, err := HttpClient.WebhookAdd(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_WebhookEdit(t *testing.T) {
	req := WebhookData{
		Id:          0,
		WebhookCode: "",
		WebhookName: "",
		WebhookType: 0,
		WebhookUrl:  "",
		CreateTime:  "",
	}
	result, err := HttpClient.WebhookEdit(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_WebhookDetail(t *testing.T) {
	result, err := HttpClient.WebhookDetail(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_MpList(t *testing.T) {
	result, err := HttpClient.MpList(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_CList(t *testing.T) {
	result, err := HttpClient.CpList(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_MailList(t *testing.T) {
	result, err := HttpClient.MailList(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_MailDetail(t *testing.T) {
	result, err := HttpClient.MailDetail(100)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_ChangeIsSend(t *testing.T) {
	result, err := HttpClient.ChangeIsSend(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_ChangeOpenMessageType(t *testing.T) {
	result, err := HttpClient.ChangeOpenMessageType(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_GetUserSettings(t *testing.T) {
	result, err := HttpClient.GetUserSettings()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_ChangeDefaultChannel(t *testing.T) {
	result, err := HttpClient.ChangeDefaultChannel("wechat", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_ChangeReceiveLimit(t *testing.T) {
	result, err := HttpClient.ChangeReceiveLimit(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)

}

func TestClient_FriendGetQRCode(t *testing.T) {
	result, err := HttpClient.FriendGetQRCode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_FriendList(t *testing.T) {
	result, err := HttpClient.FriendList(1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_FriendDel(t *testing.T) {
	result, err := HttpClient.FriendDel(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
