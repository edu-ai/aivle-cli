package operations

import (
	"net/http"
)

func apiGetRequest(apiRoot string, token string, api string) (*http.Request, error) {
	req, err := http.NewRequest("GET", apiRoot+api, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Token "+token)
	return req, nil
}
