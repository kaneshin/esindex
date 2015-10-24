package elasticsearch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	config *Config
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) Do(req *Request) (Result, error) {
	url, err := url.Parse(c.config.url)
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(url.Path, req.path)

	httpReq, err := http.NewRequest(req.method, url.String(), req.body)
	if err != nil {
		return nil, err
	}

	req.Request = httpReq
	resp, err := c.config.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := Result{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
