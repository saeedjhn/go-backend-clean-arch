package main

import (
	"context"
	"fmt"
	"maps"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type HTTPAdaptor struct {
	addr   string
	Header map[string]string
	Params map[string]string
	Paths  map[string]string
	client *resty.Client
}

func NewHTTPAdaptor(addr string) *HTTPAdaptor {
	return &HTTPAdaptor{
		addr: addr,
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Params: make(map[string]string),
		Paths:  make(map[string]string),
		client: resty.New(),
	}
}

func (c *HTTPAdaptor) WithHeader(m map[string]string) *HTTPAdaptor {
	maps.Copy(c.Header, m)

	return c
}

func (c *HTTPAdaptor) WithParam(key, value string) *HTTPAdaptor {
	c.Params[key] = value

	return c
}

func (c *HTTPAdaptor) WithParams(m map[string]string) *HTTPAdaptor {
	c.Params = m

	return c
}

func (c *HTTPAdaptor) WithPath(key, value string) *HTTPAdaptor {
	c.Paths[key] = value

	return c
}

func (c *HTTPAdaptor) WithPathParams(m map[string]string) *HTTPAdaptor {
	c.Paths = m

	return c
}

func (c *HTTPAdaptor) FetchByID(ctx context.Context, req Request) (Response, error) {
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

	r.SetPathParam("postId", strconv.FormatUint(req.ID, 10))

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
