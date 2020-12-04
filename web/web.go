package web

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func MakeAPIRequest(url string, headers map[string]string, queryParams map[string]string, v interface{}) error {
	client := resty.New()
	client.SetTimeout(time.Second * 10)
	if headers == nil {
		headers = make(map[string]string)
	}
	if queryParams == nil {
		queryParams = make(map[string]string)
	}
	if headers["Accept"] == "" {
	headers["Accept"] = "application/json"
	}

	if headers["User-Agent"] == "" {
	headers["User-Agent"] = "juicybot"
	}

	resp, err := client.R().
		SetHeaders(headers).
		SetQueryParams(queryParams).
		SetResult(&v).
		Get(url)

	if err != nil {
		return fmt.Errorf("client request error: %v", err)
	}
	if resp.StatusCode() == 404 {
		return fmt.Errorf("%v not found", url)
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("response status code not 200: %v", resp)
	}

	return err
}
