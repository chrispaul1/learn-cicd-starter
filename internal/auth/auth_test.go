package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPI(t *testing.T) {

	cases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"valid api key", "ApiKey XX-56-Y7", "XX-56-Y7"},
		{"Empty Header", "", ""},
		{"Invalid Header Format", "InvalidInput", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://example.com", nil)
			if err != nil {
				panic(err)
			}
			req.Header.Set("Authorization", c.input)

			res, _ := GetAPIKey(req.Header)

			if res != c.expectedOutput {
				t.Errorf(`--------------------------------------------------
				Expecting Output: %s
				Actual Output: %s
				Fail, Incorrect output`, c.expectedOutput, res)
				fmt.Printf("\n")
			}
		})
	}

}
