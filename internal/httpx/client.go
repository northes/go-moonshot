package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/northes/go-moonshot/internal/httpx/tools"
)

type Client struct {
	method      string
	head        http.Header
	body        any
	url         *url.URL
	contentType string
	timeout     time.Duration
	response    http.Response
	logger      Logger
	debug       bool
	error       error
}

type Response struct {
	response *http.Response
	logger   Logger
}

func NewClient(rawURL string, opts ...Option) *Client {
	cli := new(Client)
	u, err := url.Parse(rawURL)
	if err != nil {
		cli.error = err
		return cli
	}
	cli.url = u
	cli.head = make(http.Header)
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(cli)
		}
	}
	if cli.logger == nil {
		cli.logger = slog.Default()
	}
	return cli
}

func (c *Client) SetBody(body any) *Client {
	c.body = body
	return c
}

func (c *Client) SetContentType(contentType string) *Client {
	c.contentType = contentType
	return c
}

func (c *Client) AddHeader(k, v string) *Client {
	c.head.Add(k, v)
	return c
}

func (c *Client) AddHeaders(kv map[string]string) *Client {
	for k, v := range kv {
		c.head.Add(k, v)
	}
	return c
}

func (c *Client) AddPath(paths ...string) *Client {
	c.url = c.url.JoinPath(paths...)
	//c.url.JoinPath(paths...)
	return c
}

func (c *Client) AddParam(key string, value string) *Client {
	c.url.Query().Set(key, value)
	return c
}

func (c *Client) AddParams(kv map[string]string) *Client {
	for k, v := range kv {
		c.url.Query().Add(k, v)
	}
	return c
}

func (c *Client) Get(ctx ...context.Context) (*Response, error) {
	c.method = http.MethodGet
	return c.do(ctx...)
}

func (c *Client) Post(ctx ...context.Context) (*Response, error) {
	c.method = http.MethodPost
	return c.do(ctx...)
}

func (c *Client) Delete(ctx ...context.Context) (*Response, error) {
	c.method = http.MethodDelete
	return c.do(ctx...)
}

func (c *Client) Patch(ctx ...context.Context) (*Response, error) {
	c.method = http.MethodPatch
	return c.do(ctx...)
}

func (c *Client) Put(ctx ...context.Context) (*Response, error) {
	c.method = http.MethodPut
	return c.do(ctx...)
}

func (c *Client) do(ctxs ...context.Context) (*Response, error) {
	var (
		req  *http.Request
		err  error
		body *bytes.Reader
	)

	ctx := context.Background()
	if len(ctxs) != 0 {
		ctx = ctxs[0]
	}

	if c.error != nil {
		return nil, c.error
	}

	if c.body != nil {
		switch c.body.(type) {
		case string:
			b := c.body.(string)
			body = bytes.NewReader([]byte(b))
		case []byte:
			b := c.body.([]byte)
			body = bytes.NewReader(b)
		case *bytes.Buffer:
			b := c.body.(*bytes.Buffer)
			body = bytes.NewReader(b.Bytes())
		default:
			b, err := json.Marshal(c.body)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("body unmarshal(%v): %v", c.body, err))
			}
			body = bytes.NewReader(b)
			c.contentType = tools.ApplicationJson.String()
		}
	}

	if body == nil {
		req, err = http.NewRequest(c.method, c.url.String(), nil)
	} else {
		req, err = http.NewRequest(c.method, c.url.String(), body)
	}
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new request: %v", err))
	}

	req.Header = c.head
	if c.contentType != "" {
		req.Header.Set(tools.ContentTypeHeaderKey, c.contentType)
	}

	if c.debug {
		c.logger.Info("request",
			slog.String("url", c.url.String()),
			slog.String("body", fmt.Sprintf("%+v", c.body)),
			slog.String("header", fmt.Sprintf("%+v", req.Header)),
		)
	}

	client := http.Client{
		Timeout: c.timeout,
	}

	if c.debug {
		c.logger.Info("client",
			slog.Duration("timeout", client.Timeout),
		)
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("client Do(): %v", err))
	}

	response := &Response{
		response: resp,
		logger:   c.logger,
	}

	if c.debug {
		slog.Info("response",
			slog.String("resp", response.String()),
		)
	}

	return response, nil
}

func (r *Response) String() string {
	b := r.response.Body
	if b == nil {
		return ""
	}
	defer func() {
		_ = r.response.Body.Close()
	}()
	body, err := io.ReadAll(b)
	if err != nil {
		r.logger.Error(err.Error())
		return ""
	}
	r.response.Body = io.NopCloser(bytes.NewBuffer(body))
	return fmt.Sprintf("Status: %s\n  Body: %s", r.response.Status, string(body))
}

func (r *Response) Unmarshal(body any) error {
	if body == nil || r.Raw() == nil {
		return errors.New(fmt.Sprintf("response is nil or input body id nil"))
	}
	b := r.response.Body
	defer func() {
		_ = r.response.Body.Close()
	}()
	bb, err := io.ReadAll(b)
	if err != nil {
		return err
	}
	return json.Unmarshal(bb, body)
}

func (r *Response) Raw() *http.Response {
	return r.response
}

func (r *Response) StatusOK() bool {
	return r.response.StatusCode == http.StatusOK
}
