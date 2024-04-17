package moonshot_test

import (
	"context"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/northes/go-moonshot/test"
)

func TestTokenizer(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.Tokenizers().EstimateTokenCount(context.Background(), &moonshot.TokenizersEstimateTokenCountRequest{
		Model: moonshot.ModelMoonshotV18K,
		Messages: []*moonshot.ChatCompletionsMessage{
			{
				Role:    moonshot.RoleSystem,
				Content: "你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。",
			},
			{
				Role:    moonshot.RoleUser,
				Content: "你好，我叫李雷，1+1等于多少？",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", test.MarshalJsonToStringX(resp))
}
