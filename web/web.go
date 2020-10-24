package web

import (
	"fmt"
	"net/http"
	"time"
)

func MakeAPIRequest(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "juicybot")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client request error: %v", err)
	}
	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("%v not found", url)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code not 200: %v", resp)
	}

	return resp, nil
}
