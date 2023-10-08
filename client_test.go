package fincra

import (
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestClient(t *testing.T) {
	testCases := map[string]struct {
		input *Client
	}{
		"speciying sandbox true should return sandbox url": {
			NewClient("test", WithSandbox(true)),
		},
		"speciying sandbox false should return live url": {
			NewClient("live", WithSandbox(false)),
		},
		"not specifying second parameter should return live url": {
			NewClient("anotherlive"),
		},
		"speciying timeout should set the timeout": {
			NewClient("timeoutset", WithTimeout(5*time.Second)),
		},
	}
	var i = 0
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_url, _ := url.Parse(liveUrl)
			if i == 0 {
				_url, _ = url.Parse(sandboxUrl)
			}
			if tc.input.BaseUrl.String() != _url.String() {
				t.Errorf("expected %v, got %v", _url, tc.input.BaseUrl)
			}
			if i == 3 {
				if tc.input.HttpClient.Timeout != 5*time.Second {
					t.Errorf("expected %v, got %v", 5*time.Second, tc.input.HttpClient.Timeout)
				}
			}
			i++
		})
	}
}

func TestClientSendRequest(t *testing.T) {
	t.Run("send request", func(t *testing.T) {
		_, err := defaultTestClient().sendRequest("GET", "/", nil)
		if err != nil {
			t.Errorf("cannot send request %v", err)
		}
	})
}

func defaultTestClient() *Client {
	return NewClient("rrMFfLEf43q7L2lNDgdM8hDzZnsDxZos", WithSandbox(true))
}

func testEqual(t *testing.T, want, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("values not equal (-want / +got):\n%s", diff)
	}
}
