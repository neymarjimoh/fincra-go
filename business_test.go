package fincra

import (
	"fmt"
	"testing"
)

func TestGetBusinessId(t *testing.T) {
	t.Run("fetch business info", func(t *testing.T) {
		client := defaultTestClient()

		resp, err := client.GetBusinessId()
		if err != nil {
			t.Errorf("error fetching business info: %v", err)
		}

		fmt.Println(resp)

		want := map[string]interface{}{
			"success": true,
			"message": "Parent business fetched successfully",
		}

		got := make(map[string]interface{}, len(want))
		for k, v := range resp {
			if k == "message" || k == "success" {
				got[k] = v
			}
		}

		testEqual(t, got, want)
	})
}
