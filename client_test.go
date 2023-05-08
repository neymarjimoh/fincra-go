package fincra

import (
	"net/url"
	"testing"
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
