package moonshot_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/northes/go-moonshot"
)

func TestNewChatCompletionsBuilder(t *testing.T) {
	tt := require.New(t)

	builder := moonshot.NewChatCompletionsBuilder()
	tt.NotNil(builder)

	const (
		promptContent    = "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。"
		userContent      = "你好，我叫李雷，1+1等于多少？"
		assistantContent = "我是小助手！"
		functionName1    = "function1"
		functionName2    = "function2"
	)

	var wantedReq = &moonshot.ChatCompletionsRequest{
		Messages: []*moonshot.ChatCompletionsMessage{
			{
				Role:    moonshot.RoleContextCache,
				Content: "tag=tag1;reset_ttl=3600;dry_run=1",
			},
			{
				Role:    moonshot.RoleSystem,
				Content: promptContent,
			},
			{
				Role:    moonshot.RoleUser,
				Content: userContent,
			},
			{
				Role:    moonshot.RoleAssistant,
				Content: assistantContent,
			},
			{
				Role:    moonshot.RoleUser,
				Content: userContent,
			},
		},
		Model:            moonshot.ModelMoonshotV132K,
		MaxTokens:        1024,
		Temperature:      0.3,
		TopP:             1.0,
		N:                1,
		PresencePenalty:  1.2,
		FrequencyPenalty: 1.5,
		ResponseFormat: &moonshot.ChatCompletionsRequestResponseFormat{
			Type: moonshot.ChatCompletionsResponseFormatJSONObject,
		},
		Stop:   []string{"结束"},
		Stream: true,
		Tools: []*moonshot.ChatCompletionsTool{{
			Type: moonshot.ChatCompletionsToolTypeFunction,
			Function: &moonshot.ChatCompletionsToolFunction{
				Name:        functionName1,
				Description: "",
				Parameters:  nil,
			},
		}, {
			Type: moonshot.ChatCompletionsToolTypeFunction,
			Function: &moonshot.ChatCompletionsToolFunction{
				Name: functionName2,
			},
		}, {
			Type: moonshot.ChatCompletionsToolTypeFunction,
			Function: &moonshot.ChatCompletionsToolFunction{
				Name: functionName2,
			},
		}},
	}

	builder.SetContextCacheContent(
		moonshot.NewContextCacheContentWithTag("tag1").
			WithResetTTL(3600).
			WithDryRun(true)).
		AddPrompt(promptContent).
		AddUserContent(userContent).
		AddAssistantContent(assistantContent).
		AddMessage(&moonshot.ChatCompletionsMessage{
			Role:    moonshot.RoleUser,
			Content: userContent,
		}).
		SetModel(moonshot.ModelMoonshotV132K).
		SetMaxTokens(1024).
		SetTemperature(0.3).
		SetTopP(1.0).
		SetN(1).
		SetPresencePenalty(1.2).
		SetFrequencyPenalty(1.5).
		SetResponseFormat(moonshot.ChatCompletionsResponseFormatJSONObject).
		SetStop([]string{"结束"}).
		SetStream(true).
		SetTool(&moonshot.ChatCompletionsTool{
			Type: moonshot.ChatCompletionsToolTypeFunction,
			Function: &moonshot.ChatCompletionsToolFunction{
				Name:        functionName1,
				Description: "",
				Parameters:  nil,
			},
		}).SetTools([]*moonshot.ChatCompletionsTool{
		{
			Type: moonshot.ChatCompletionsToolTypeFunction,
			Function: &moonshot.ChatCompletionsToolFunction{
				Name: functionName2,
			},
		},
		{
			Type: moonshot.ChatCompletionsToolTypeFunction,
			Function: &moonshot.ChatCompletionsToolFunction{
				Name: functionName2,
			},
		},
	})

	builder.SetTools(nil)
	tt.Equal(wantedReq, builder.ToRequest())

	tt.Equal(wantedReq, builder.ToRequest())
	tt.NotEqual(wantedReq, builder.SetModel(moonshot.ModelMoonshotV1128K).ToRequest())

	builder2 := moonshot.NewChatCompletionsBuilder(*wantedReq)
	tt.Equal(wantedReq, builder2.ToRequest())

	builder2.SetPresencePenalty(2)
	tt.NotEqual(wantedReq, builder2.ToRequest())

	builder2.SetResponseFormat(moonshot.ChatCompletionsResponseFormatText)
	tt.NotEqual(wantedReq, builder2.ToRequest())
}
