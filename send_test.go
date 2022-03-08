package pushplus

import (
	"os"
	"testing"
)

var HttpClient *Client

func TestMain(m *testing.M) {
	token := "0787bab863f84f02ab3c5a1510336e66"
	key := "5b2e1c522248443d85b3281de67eeeed"
	HttpClient = NewClient(token, SecretKey(key), DebugLog(true))
	code := m.Run()
	os.Exit(code)
}

func TestSendMsg(t *testing.T) {
	result, err := HttpClient.SendMsg("11111", "2222")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestAuthorize(t *testing.T) {
	var da DefaultAuthorizer
	da.Authorize(HttpClient)
}

func TestAuth(t *testing.T) {
	HttpClient.Auth()
}
