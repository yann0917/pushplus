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

func TestTopicList(t *testing.T) {
	result, err := HttpClient.TopicList(1, 20, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
