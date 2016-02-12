package opentsdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Options struct {
	// Host value for the opentsdb server
	// Default: 127.0.0.1
	Host string

	// Port for the opentsdb server
	// Default: 4242
	Port int
}

type Client struct {
	url        *url.URL
	httpClient *http.Client
}

func NewClient(opt Options) (*Client, error) {
	if opt.Host == "" {
		opt.Host = "127.0.0.1"
	}
	if opt.Port == 0 {
		opt.Port = 4242
	}

	u, err := url.Parse(fmt.Sprintf("http://%s:%d", opt.Host, opt.Port))
	if err != nil {
		return nil, err
	}

	return &Client{
		url:        u,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) Aggregators() error {
	return nil
}

func (c *Client) Annotation() error {
	return nil
}

func (c *Client) Config() error {
	return nil
}

func (c *Client) Dropcaches() error {
	return nil
}

func (c *Client) Put(bp *BatchPoints, params string) ([]byte, error) {
	data, err := bp.ToJson()
	if err != nil {
		return nil, err
	}

	u := c.url
	u.Path = "api/put"
	u.RawQuery = params

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Query(q *QueryParams) ([]byte, error) {
	data, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	u := c.url
	u.Path = "api/query"

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Search() error {
	return nil
}

func (c *Client) Serializers() error {
	return nil
}

func (c *Client) Stats() error {
	return nil
}

func (c *Client) Suggest() error {
	return nil
}

func (c *Client) Tree() error {
	return nil
}

func (c *Client) Uid() error {
	return nil
}

func (c *Client) Version() error {
	return nil
}
