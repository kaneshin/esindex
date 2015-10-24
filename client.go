package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

func indexNameFromAliasName(aliasName string) string {
	return fmt.Sprintf("%s_v%s", aliasName, time.Now().Format(versionFormat))
}

func aliasNameFromIndexName(indexName string) string {
	return indexName[:len(indexName)-len("_v"+versionFormat)]
}

type ElasticsearchClient struct {
	url        string
	HTTPClient *http.Client
}

type Request struct {
	*http.Request
	method string
	path   string
	body   io.Reader
}

func NewRequest(method, path string, body io.Reader) *Request {
	return &Request{
		method: method,
		path:   path,
		body:   body,
	}
}

var DefaultClient = &ElasticsearchClient{*urlStr, &http.Client{}}

type Result map[string]interface{}

func (c *ElasticsearchClient) Do(req *Request) (Result, error) {
	url, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	url.Path = path.Join(url.Path, req.path)
	httpReq, err := http.NewRequest(req.method, url.String(), req.body)
	if err != nil {
		return nil, err
	}
	req.Request = httpReq
	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := Result{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, err
}
