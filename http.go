package moonshot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IRequest interface {
	Path() string
}

type Httpx struct {
	Method     string
	URL        string
	reqBuilder func(in IRequest) (*http.Request, error)
}

func (c *Client) RegisterAPI(req IRequest, method string, reqBuilders ...func(in IRequest) (*http.Request, error)) error {
	if _, ok := c.requestToURLMap[req.Path()]; ok {
		return fmt.Errorf("httpx: request already registered for path %s", req.Path())
	}

	reqBuilder := func(in IRequest) (*http.Request, error) {
		v, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		r, err := http.NewRequest(method, c.cfg.Host+in.Path(), bytes.NewReader(v))
		if err != nil {
			return nil, err
		}
		return r, nil
	}

	if len(reqBuilders) != 0 {
		if reqBuilders[0] != nil {
			reqBuilder = reqBuilders[0]
		}
	}

	c.requestToURLMap[req.Path()] = &Httpx{
		Method:     method,
		URL:        c.cfg.Host + req.Path(),
		reqBuilder: reqBuilder,
	}
	return nil
}

func (c *Client) Do(ctx context.Context, req IRequest, resp ...any) (*http.Response, error) {
	ctrl, ok := c.requestToURLMap[req.Path()]
	if !ok {
		return nil, fmt.Errorf("httpx: request already registered for path %s", req.Path())
	}
	r, err := ctrl.reqBuilder(req)
	if err != nil {
		return nil, fmt.Errorf("httpx: builder request %w", err)
	}
	client := http.DefaultClient
	if ctx != nil {
		if hc, ok := ctx.Value(ContextHTTPClient).(*http.Client); ok {
			client = hc
		}
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.cfg.APIKey))
	rsp, err := client.Do(r.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("httpx: http request %w", err)
	}

	if err = StatusCodeToError(rsp.StatusCode); err != nil {
		return rsp, fmt.Errorf("httpx: http response [%d] %w", rsp.StatusCode, err)
	}

	if len(resp) != 0 {
		rr := resp[0]
		body, err := io.ReadAll(rsp.Body)
		if err != nil {
			return nil, fmt.Errorf("httpx: read body %w", err)
		}
		defer rsp.Body.Close()
		err = json.Unmarshal(body, rr)
		if err != nil {
			return rsp, fmt.Errorf("httpx: unmarshal body %w", err)
		}
	}

	return rsp, nil
}
