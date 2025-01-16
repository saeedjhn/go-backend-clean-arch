package setuptest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var _errResponseNotInitial = errors.New("response is not initialized")

type TinyRequest struct {
	buf     bytes.Buffer
	headers map[string]string
	req     *http.Request
	resp    *http.Response
}

func NewTinyRequest() *TinyRequest {
	return &TinyRequest{}
}

func (r *TinyRequest) SetHeader(headers map[string]string) *TinyRequest {
	r.headers = headers

	return r
}

func (r *TinyRequest) SetBody(body interface{}) error {
	if body != nil {
		err := json.NewEncoder(&r.buf).Encode(body)
		if err != nil {
			return fmt.Errorf("could not encode body: %w", err)
		}
	}

	return nil
}

func (r *TinyRequest) Status() string {
	return r.resp.Status
}

func (r *TinyRequest) StatusCode() int {
	return r.resp.StatusCode
}

func (r *TinyRequest) Body() ([]byte, error) {
	if r.resp == nil {
		return nil, _errResponseNotInitial
	}

	var closeErr error
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			closeErr = fmt.Errorf("error closing response body: %w", err)
		}
	}(r.resp.Body)

	data, readErr := io.ReadAll(r.resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("error reading response body: %w", readErr)
	}

	if closeErr != nil {
		return nil, closeErr
	}

	return data, nil
}

func (r *TinyRequest) UnmarshallBody(v interface{}) error {
	if r.resp == nil {
		return _errResponseNotInitial
	}

	var closeErr error
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			closeErr = fmt.Errorf("error closing response body: %w", err)
		}
	}(r.resp.Body)

	if decodeErr := json.NewDecoder(r.resp.Body).Decode(v); decodeErr != nil {
		return decodeErr
	}

	return closeErr
}

func (r *TinyRequest) Fetch(method, uri string) (*TinyRequest, error) {
	var err error

	r.req, err = http.NewRequest(method, uri, &r.buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request with method %s and URI %s: %w", method, uri, err)
	}

	r.setHeaders()

	client := &http.Client{}
	resp, err := client.Do(r.req) //nolint:bodyclose // nothing
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	r.resp = resp

	return r, nil
}

func (r *TinyRequest) setHeaders() {
	if len(r.headers) != 0 {
		for k, v := range r.headers {
			r.req.Header.Set(k, v)
		}
	}
}
