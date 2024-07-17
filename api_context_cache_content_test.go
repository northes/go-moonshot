package moonshot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContextCacheContentWithId(t *testing.T) {
	const cacheId = "my_cache_id"
	expected := &ContextCacheContent{
		CacheId:  cacheId,
		ResetTTL: ResetTTLNever,
	}
	actual := NewContextCacheContentWithId(cacheId)
	assert.Equal(t, expected, actual)

	const resetTTL = 3600
	expected = &ContextCacheContent{
		CacheId:  cacheId,
		ResetTTL: resetTTL,
	}
	actual = NewContextCacheContentWithId(cacheId).
		WithResetTTL(resetTTL)
	assert.Equal(t, expected, actual)

	const dryRun = true
	expected = &ContextCacheContent{
		CacheId:  cacheId,
		ResetTTL: resetTTL,
		DryRun:   dryRun,
	}
	actual = NewContextCacheContentWithId(cacheId).
		WithResetTTL(resetTTL).
		WithDryRun(dryRun)
	assert.Equal(t, expected, actual)

	const tag = "my_tag"
	expected = &ContextCacheContent{
		Tag:      tag,
		ResetTTL: ResetTTLNever,
	}
	actual = NewContextCacheContentWithTag(tag)
	assert.Equal(t, expected, actual)
}
