package moonshot

import (
	"fmt"
	"strings"
)

const (
	ResetTTLNever     = -1 // ResetTTLNever is the value for never reset ttl
	ResetTTLImmediate = 0  // ResetTTLImmediate is the value for immediate reset ttl
)

// ContextCacheContent is the content for the context cache
type ContextCacheContent struct {
	CacheId  string `json:"cache_id"`
	Tag      string `json:"tag"`
	ResetTTL int64  `json:"reset_ttl"`
	DryRun   bool   `json:"dry_run"`
}

func NewContextCacheContentWithId(cacheId string) *ContextCacheContent {
	return &ContextCacheContent{CacheId: cacheId, ResetTTL: ResetTTLNever}
}

func NewContextCacheContentWithTag(tag string) *ContextCacheContent {
	return &ContextCacheContent{Tag: tag, ResetTTL: ResetTTLNever}
}

// WithResetTTL set the reset ttl for the context cache
func (c *ContextCacheContent) WithResetTTL(resetTTL int64) *ContextCacheContent {
	c.ResetTTL = resetTTL
	return c
}

// WithDryRun set the dry run for the context cache
func (c *ContextCacheContent) WithDryRun(dryRun bool) *ContextCacheContent {
	c.DryRun = dryRun
	return c
}

func (c *ContextCacheContent) Content() string {
	var slice []string

	if c.CacheId != "" {
		slice = append(slice, fmt.Sprintf("cache_id=%s", c.CacheId))
	} else if c.Tag != "" {
		slice = append(slice, fmt.Sprintf("tag=%s", c.Tag))
	}

	if c.ResetTTL >= 0 {
		slice = append(slice, fmt.Sprintf("reset_ttl=%d", c.ResetTTL))
	}

	if c.DryRun {
		slice = append(slice, "dry_run=1")
	}

	return strings.Join(slice, ";")
}
