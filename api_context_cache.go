package moonshot

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type IContextCache interface {
	Create(ctx context.Context, req *ContextCacheCreateRequest) (*ContextCacheCreateResponse, error)
	List(ctx context.Context, req *ContextCacheListRequest) (*ContextCacheListResponse, error)
	Delete(ctx context.Context, req *ContextCacheDeleteRequest) (*ContextCacheDeleteResponse, error)
	Update(ctx context.Context, req *ContextCacheUpdateRequest) (*ContextCacheUpdateResponse, error)
	Get(ctx context.Context, req *ContextCacheGetRequest) (*ContextCacheGetResponse, error)

	CreateTag(ctx context.Context, req *ContextCacheCreateTagRequest) (*ContextCacheCreateTagResponse, error)
	ListTag(ctx context.Context, req *ContextCacheListTagRequest) (*ContextCacheListTagResponse, error)
	DeleteTag(ctx context.Context, req *ContextCacheDeleteTagRequest) (*ContextCacheDeleteTagResponse, error)
	GetTag(ctx context.Context, req *ContextCacheGetTagRequest) (*ContextCacheGetTagResponse, error)
	GetTagContent(ctx context.Context, req *ContextCacheGetTagContentRequest) (*ContextCacheGetTagContentResponse, error)
}

type contextCache struct {
	client *Client
}

// ContextCache returns a new context cache controller
func (c *Client) ContextCache() IContextCache {
	return &contextCache{
		client: c,
	}
}

// ContextCache is the cache of the context
type ContextCache struct {
	Id          string                   `json:"id"`          // 缓存的唯一标识
	Status      ContextCacheStatus       `json:"status"`      // 缓存的状态
	Object      string                   `json:"object"`      // 缓存的类型
	CreatedAt   int64                    `json:"created_at"`  // 缓存的创建时间
	ExpiredAt   int64                    `json:"expired_at"`  // 缓存的过期时间
	Tokens      int                      `json:"tokens"`      // 缓存的 Token 数量
	Model       ChatCompletionsModelID   `json:"model"`       // 缓存的模型组名称
	Messages    []ChatCompletionsMessage `json:"messages"`    // 缓存的消息内容
	Tools       []ChatCompletionsTool    `json:"tools"`       // 缓存使用的工具
	Name        string                   `json:"name"`        // 缓存的名称
	Description string                   `json:"description"` // 缓存的描述信息
	Metadata    map[string]string        `json:"metadata"`    // 缓存的元信息
}

// ContextCacheCreateRequest is the request for creating a context cache
type ContextCacheCreateRequest struct {
	Model       ChatCompletionsModelID   `json:"model"`                 // 模型组（model family）名称
	Messages    []ChatCompletionsMessage `json:"messages"`              // 消息内容
	Tools       []ChatCompletionsTool    `json:"tools,omitempty"`       // 使用的工具
	Name        string                   `json:"name,omitempty"`        // 缓存名称
	Description string                   `json:"description,omitempty"` // 缓存描述信息
	Metadata    map[string]string        `json:"metadata,omitempty"`    // 缓存的元信息
	ExpiredAt   int64                    `json:"expired_at"`            // 缓存的过期时间
	TTL         int64                    `json:"ttl,omitempty"`         // 缓存的有效期
}

// ContextCacheCreateResponse is the response for creating a context cache
type ContextCacheCreateResponse ContextCache

// ContextCacheListRequest is the request for listing context caches
type ContextCacheListRequest struct {
	Limit    int               `json:"limit,omitempty"`    // 当前请求单页返回的缓存数量
	Order    ContextCacheOrder `json:"order,omitempty"`    // 当前请求时查询缓存的排序规则
	After    string            `json:"after,omitempty"`    // 当前请求时，应该从哪一个缓存开始进行查找
	Before   string            `json:"before,omitempty"`   // 当前请求时，应该查询到哪一个缓存为止
	Metadata map[string]string `json:"metadata,omitempty"` // 用于筛选缓存的 metadata 信息
}

// ContextCacheListResponse is the response for listing context caches
type ContextCacheListResponse struct {
	Object string         `json:"object"` // 返回的数据类型
	Data   []ContextCache `json:"data"`   // 返回的缓存列表
}

// ContextCacheDeleteRequest is the request for deleting a context cache
type ContextCacheDeleteRequest struct {
	Id string `json:"id"` // 缓存的唯一标识
}

// ContextCacheDeleteResponse is the response for deleting a context cache
type ContextCacheDeleteResponse struct {
	Deleted bool   `json:"deleted"` // 缓存是否被删除
	Id      string `json:"id"`      // 被删除的缓存的唯一标识
	Object  string `json:"object"`  // 返回的数据类型
}

// ContextCacheUpdateRequest is the request for updating a context cache
type ContextCacheUpdateRequest struct {
	Id        string            `json:"_"`                  // 缓存的唯一标识
	Metadata  map[string]string `json:"metadata,omitempty"` // 缓存的元信息
	ExpiredAt int64             `json:"expired_at"`         // 缓存的过期时间
	TTL       int64             `json:"ttl,omitempty"`      // 缓存的有效期
}

// ContextCacheUpdateResponse is the response for updating a context cache
type ContextCacheUpdateResponse ContextCache

// ContextCacheGetRequest is the request for getting a context cache
type ContextCacheGetRequest struct {
	Id string `json:"_"` // 缓存的唯一标识
}

// ContextCacheGetResponse is the response for getting a context cache
type ContextCacheGetResponse ContextCache

// ContextCacheTag is the tag of the context cache
type ContextCacheTag struct {
	Tag       string `json:"tag"`        // 缓存的标签
	CacheId   string `json:"cache_id"`   // 缓存的唯一标识
	Object    string `json:"object"`     // 缓存的类型
	OwnedBy   string `json:"owned_by"`   // 缓存的拥有者
	CreatedAt int    `json:"created_at"` // 缓存的创建时间
}

// ContextCacheCreateTagRequest is the request for creating a context cache tag
type ContextCacheCreateTagRequest struct {
	Tag     string `json:"tag"`      // 缓存的标签
	CacheId string `json:"cache_id"` // 缓存的唯一标识
}

// ContextCacheCreateTagResponse is the response for creating a context cache tag
type ContextCacheCreateTagResponse ContextCacheTag

// ContextCacheListTagRequest is the request for listing context cache tags
type ContextCacheListTagRequest struct {
	Limit  int               `json:"limit,omitempty"`  // 当前请求单页返回的缓存数量
	Order  ContextCacheOrder `json:"order,omitempty"`  // 当前请求时查询缓存的排序规则
	After  string            `json:"after,omitempty"`  // 当前请求时，应该从哪一个缓存开始进行查找
	Before string            `json:"before,omitempty"` // 当前请求时，应该查询到哪一个缓存为止
}

// ContextCacheListTagResponse is the response for listing context cache tags
type ContextCacheListTagResponse struct {
	Object string            `json:"object"` // 返回的数据类型
	Data   []ContextCacheTag `json:"data"`   // 返回的缓存标签列表
}

// ContextCacheDeleteTagRequest is the request for deleting a context cache tag
type ContextCacheDeleteTagRequest struct {
	Tag string `json:"_"` // 缓存的标签
}

// ContextCacheDeleteTagResponse is the response for deleting a context cache tag
type ContextCacheDeleteTagResponse struct {
	Deleted bool   `json:"deleted"` // 缓存是否被删除
	Object  string `json:"object"`  // 返回的数据类型
	Tag     string `json:"tag"`     // 被删除的缓存的标签
}

// ContextCacheGetTagRequest is the request for getting a context cache tag
type ContextCacheGetTagRequest struct {
	Tag string `json:"_"` // 缓存的标签
}

// ContextCacheGetTagResponse is the response for getting a context cache tag
type ContextCacheGetTagResponse ContextCacheTag

// ContextCacheGetTagContentRequest is the request for getting a context cache tag content
type ContextCacheGetTagContentRequest struct {
	Tag string `json:"_"` // 缓存的标签
}

// ContextCacheGetTagContentResponse is the response for getting a context cache tag content
type ContextCacheGetTagContentResponse ContextCache

func (c *contextCache) Create(ctx context.Context, req *ContextCacheCreateRequest) (*ContextCacheCreateResponse, error) {
	const path = "/v1/caching"
	contextCacheCreateResp := new(ContextCacheCreateResponse)
	resp, err := c.client.HTTPClient().SetPath(path).SetBody(req).Post(ctx)
	if err != nil {
		return contextCacheCreateResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheCreateResp)
	if err != nil {
		return nil, err
	}
	return contextCacheCreateResp, nil
}

func (c *contextCache) List(ctx context.Context, req *ContextCacheListRequest) (*ContextCacheListResponse, error) {
	path := "/v1/caching"
	params := url.Values{}
	if req.Limit > 0 {
		params.Add("limit", strconv.Itoa(req.Limit))
	}
	if req.Order != "" {
		params.Add("order", req.Order.String())
	}
	if req.After != "" {
		params.Add("after", req.After)
	}
	if req.Before != "" {
		params.Add("before", req.Before)
	}
	if len(req.Metadata) > 0 {
		for k, v := range req.Metadata {
			params.Add(fmt.Sprintf("metadata[%s]", k), v)
		}
	}
	if len(params) > 0 {
		path = fmt.Sprintf("%s?%s", path, params.Encode())
	}
	contextCacheListResp := new(ContextCacheListResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Get(ctx)
	if err != nil {
		return contextCacheListResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheListResp)
	if err != nil {
		return nil, err
	}
	return contextCacheListResp, nil
}

func (c *contextCache) Delete(ctx context.Context, req *ContextCacheDeleteRequest) (*ContextCacheDeleteResponse, error) {
	path := fmt.Sprintf("/v1/caching/%s", req.Id)
	contextCacheDeleteResp := new(ContextCacheDeleteResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Delete(ctx)
	if err != nil {
		return contextCacheDeleteResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheDeleteResp)
	if err != nil {
		return nil, err
	}
	return contextCacheDeleteResp, nil
}

func (c *contextCache) Update(ctx context.Context, req *ContextCacheUpdateRequest) (*ContextCacheUpdateResponse, error) {
	path := fmt.Sprintf("/v1/caching/%s", req.Id)
	contextCacheUpdateResp := new(ContextCacheUpdateResponse)
	resp, err := c.client.HTTPClient().SetPath(path).SetBody(req).Put(ctx)
	if err != nil {
		return contextCacheUpdateResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheUpdateResp)
	if err != nil {
		return nil, err
	}
	return contextCacheUpdateResp, nil
}

func (c *contextCache) Get(ctx context.Context, req *ContextCacheGetRequest) (*ContextCacheGetResponse, error) {
	path := fmt.Sprintf("/v1/caching/%s", req.Id)
	contextCacheGetResp := new(ContextCacheGetResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Get(ctx)
	if err != nil {
		return contextCacheGetResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheGetResp)
	if err != nil {
		return nil, err
	}
	return contextCacheGetResp, nil
}

func (c *contextCache) CreateTag(ctx context.Context, req *ContextCacheCreateTagRequest) (*ContextCacheCreateTagResponse, error) {
	const path = "/v1/caching/refs/tags"
	contextCacheCreateTagResp := new(ContextCacheCreateTagResponse)
	resp, err := c.client.HTTPClient().SetPath(path).SetBody(req).Post(ctx)
	if err != nil {
		return contextCacheCreateTagResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheCreateTagResp)
	if err != nil {
		return nil, err
	}
	return contextCacheCreateTagResp, nil
}

func (c *contextCache) ListTag(ctx context.Context, req *ContextCacheListTagRequest) (*ContextCacheListTagResponse, error) {
	path := "/v1/caching/refs/tags"
	params := url.Values{}
	if req.Limit > 0 {
		params.Add("limit", strconv.Itoa(req.Limit))
	}
	if req.Order != "" {
		params.Add("order", req.Order.String())
	}
	if req.After != "" {
		params.Add("after", req.After)
	}
	if req.Before != "" {
		params.Add("before", req.Before)
	}
	if len(params) > 0 {
		path = fmt.Sprintf("%s?%s", path, params.Encode())
	}
	contextCacheListTagResp := new(ContextCacheListTagResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Get(ctx)
	if err != nil {
		return contextCacheListTagResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheListTagResp)
	if err != nil {
		return nil, err
	}
	return contextCacheListTagResp, nil
}

func (c *contextCache) DeleteTag(ctx context.Context, req *ContextCacheDeleteTagRequest) (*ContextCacheDeleteTagResponse, error) {
	path := fmt.Sprintf("/v1/caching/refs/tags/%s", req.Tag)
	contextCacheDeleteTagResp := new(ContextCacheDeleteTagResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Delete(ctx)
	if err != nil {
		return contextCacheDeleteTagResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheDeleteTagResp)
	if err != nil {
		return nil, err
	}
	return contextCacheDeleteTagResp, nil
}

func (c *contextCache) GetTag(ctx context.Context, req *ContextCacheGetTagRequest) (*ContextCacheGetTagResponse, error) {
	path := fmt.Sprintf("/v1/caching/refs/tags/%s", req.Tag)
	contextCacheGetTagResp := new(ContextCacheGetTagResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Get(ctx)
	if err != nil {
		return contextCacheGetTagResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheGetTagResp)
	if err != nil {
		return nil, err
	}
	return contextCacheGetTagResp, nil
}

func (c *contextCache) GetTagContent(ctx context.Context, req *ContextCacheGetTagContentRequest) (*ContextCacheGetTagContentResponse, error) {
	path := fmt.Sprintf("/v1/caching/refs/tags/%s/content", req.Tag)
	contextCacheGetTagContentResp := new(ContextCacheGetTagContentResponse)
	resp, err := c.client.HTTPClient().SetPath(path).Get(ctx)
	if err != nil {
		return contextCacheGetTagContentResp, err
	}
	if !resp.StatusOK() {
		return nil, ResponseToError(resp)
	}
	err = resp.Unmarshal(contextCacheGetTagContentResp)
	if err != nil {
		return nil, err
	}
	return contextCacheGetTagContentResp, nil
}
