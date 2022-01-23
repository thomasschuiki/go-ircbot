package web

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func MakeAPIRequest(url string, headers map[string]string, queryParams map[string]string, v interface{}) error {
	client := resty.New()
	client.SetTimeout(time.Second * 10)
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

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
		return fmt.Errorf("client request error: %w", err)
	}
	if resp.StatusCode() == 404 {
		return fmt.Errorf("%s not found", url)
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("response status code not 200: %#v", resp)
	}

	return err
}

func GetWebpage(url string, queryParams map[string]string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-controle", "no-cache")
	req.Header.Set("dnt", "1")
	req.Header.Set("upgrade-insecure-requests", "1")

	q := req.URL.Query() // Get a copy of the query values.
	for k, v := range queryParams {
		q.Add(k, v) // Add a new value to the set.
	}
	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return resp.Body, nil
	}
	return nil, err
}
