package moonshot_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/northes/go-moonshot"
	"github.com/northes/go-moonshot/test"
)

// Due to some problems, this test is temporarily shielded in the GitHub Actions environment.
// waiting for the official Mock Server.
// https://github.com/MoonshotAI/moonpalace

func TestContextCache(t *testing.T) {
	if test.IsGithubActions() {
		return
	}
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	// Create
	var createResponse *moonshot.ContextCacheCreateResponse
	createResponse, err = cli.ContextCache().Create(ctx, &moonshot.ContextCacheCreateRequest{
		Model: moonshot.ModelFamilyMoonshotV1,
		Messages: []moonshot.ChatCompletionsMessage{
			{
				Role:    moonshot.RoleSystem,
				Content: "你是一个翻译机器人，我会告诉你一些英文，请帮我翻译成中文。",
			},
		},
		ExpiredAt: -1,
		TTL:       60,
	})
	if err != nil {
		t.Fatal(err)
	}
	id := createResponse.Id
	assert.Equal(t, time.Now().Unix()+60, createResponse.ExpiredAt)

	// Get
	var getResponse *moonshot.ContextCacheGetResponse
	getResponse, err = cli.ContextCache().Get(ctx, &moonshot.ContextCacheGetRequest{
		Id: id,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, id, getResponse.Id)

	// Update
	var updateResponse *moonshot.ContextCacheUpdateResponse
	updateResponse, err = cli.ContextCache().Update(ctx, &moonshot.ContextCacheUpdateRequest{
		Id:        id,
		ExpiredAt: -1,
		TTL:       120,
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = updateResponse
	// assert.Equal(t, time.Now().Unix()+120, updateResponse.ExpiredAt)

	// List
	var listResponse *moonshot.ContextCacheListResponse
	listResponse, err = cli.ContextCache().List(ctx, &moonshot.ContextCacheListRequest{
		// Limit:  10,
		// Order:  moonshot.ContextCacheOrderAsc,
		// After:  "",
		// Before: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.GreaterOrEqual(t, len(listResponse.Data), 1)

	// Delete
	var deleteResponse *moonshot.ContextCacheDeleteResponse
	deleteResponse, err = cli.ContextCache().Delete(ctx, &moonshot.ContextCacheDeleteRequest{
		Id: id,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, id, deleteResponse.Id)
}

func TestContextCache_Create(t *testing.T) {
	if test.IsGithubActions() {
		return
	}
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	response, err := cli.ContextCache().Create(ctx, &moonshot.ContextCacheCreateRequest{
		Model: moonshot.ModelFamilyMoonshotV1,
		Messages: []moonshot.ChatCompletionsMessage{
			{
				Role:    moonshot.RoleSystem,
				Content: "你是一个翻译机器人，我会告诉你一些英文，请帮我翻译成中文。",
			},
		},
		Tools:       nil,
		Description: "这是一个测试的描述",
		Metadata:    nil,
		ExpiredAt:   -1,
		TTL:         60,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test.MarshalJsonToStringX(response))
}

func TestContextCache_Delete(t *testing.T) {
	if test.IsGithubActions() {
		return
	}
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	response, err := cli.ContextCache().Delete(ctx, &moonshot.ContextCacheDeleteRequest{
		Id: "cache-xxxx",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test.MarshalJsonToStringX(response))
}

func TestContextCache_List(t *testing.T) {
	if test.IsGithubActions() {
		return
	}
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	response, err := cli.ContextCache().List(ctx, &moonshot.ContextCacheListRequest{})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test.MarshalJsonToStringX(response))
}

func TestContextCache_CreateTag(t *testing.T) {
	if test.IsGithubActions() {
		return
	}
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	// Create
	var createResponse *moonshot.ContextCacheCreateTagResponse
	createResponse, err = cli.ContextCache().CreateTag(ctx, &moonshot.ContextCacheCreateTagRequest{
		Tag:     "MyCacheTag",
		CacheId: "cache-",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "MyCacheTag", createResponse.Tag)
}
