package apiclient

import "testing"

func TestMakeNewsApiHTTPClient(t *testing.T) {
	type test struct {
		apiUrl      string
		expectedUrl string
	}

	tests := []test{
		{"https://someapi.com/", "https://someapi.com/"},
		{"https://someapi.com", "https://someapi.com/"},
	}

	for _, test := range tests {
		client := MakeNewsApiHTTPClient(ApiAuthDetails{
			ApiKey: "apikey",
			ApiUrl: test.apiUrl,
		})
		if client.apiAuthDetails.ApiUrl != test.expectedUrl {
			t.Fatalf("Expected %v, got %v", test.expectedUrl, client.apiAuthDetails.ApiUrl)
		}
	}
}
