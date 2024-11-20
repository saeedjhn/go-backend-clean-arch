package main

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"maps"
)

type HTTPClient struct {
	addr   string
	Header map[string]string
	Params map[string]string
	client *resty.Client
}

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{
		addr: addr,
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		client: resty.New(),
	}
}

func (c *HTTPClient) WithHeader(m map[string]string) *HTTPClient {
	maps.Copy(c.Header, m)

	return c
}

func (c *HTTPClient) WithParams(m map[string]string) *HTTPClient {
	c.Params = m

	return c
}

func (c *HTTPClient) WithParam(key, value string) *HTTPClient {
	c.Params[key] = value

	return c
}

func (c *HTTPClient) Get(ctx context.Context, req Request) (Response, error) {
	var rs Response

	r := c.client.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&rs)

	if len(c.Header) != 0 {
		for k, v := range c.Header {
			r.SetHeader(k, v)
		}
	}

	if len(c.Params) != 0 {
		r.SetQueryParams(c.Params)
	}

	resp, err := r.Get(c.addr)

	if err != nil {
		return rs, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return rs, fmt.Errorf(
			"error: status code %d, body: %s",
			resp.StatusCode(),
			resp.String(),
		)
	}

	return rs, nil
}
