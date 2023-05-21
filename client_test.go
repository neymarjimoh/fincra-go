package fincra

import (
	"net/url"
	"testing"

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
		"not speciying second parameter should return live url": {
			NewClient("anotherlive"),
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
