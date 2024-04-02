package moonshot_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/northes/go-moonshot"
	"github.com/northes/gox"
)

func TestChat(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	resp, err := cli.Chat().Completions(ctx, &moonshot.ChatCompletionsRequest{
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
		Temperature: 0.3,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(gox.JsonMarshalToStringX(resp))
	/*
		{"id":"chatcmpl-dafad118ba6a4d1bb3e10be1734c6213","object":"chat.completion","created":15893254,"model":"moonshot-v1-8k","choices":[{"index":0,"message":{"Role":"assistant","Content":"你好，李雷！1+1等于2。如果你有更复杂的数学问题或者其他问题，也可以随时问我。"},"finish_reason":"stop"}],"usage":{"prompt_tokens":83,"completion_tokens":25,"total_tokens":108}}
	*/
}

//func TestChatStream(t *testing.T) {
//	cli, err := NewTestClient()
//	if err != nil {
//		t.Fatal(err)
//	}
//	ctx := context.Background()
//	ch := make(chan *moonshot.ChatCompletionsResponse)
//	done := make(chan struct{})
//
//	go func() {
//		for {
//			select {
//			case resp := <-ch:
//				//t.Log("got resp")
//				//t.Logf("%+v", moonshot.MarshalToStringX(resp))
//				delta := resp.Choices[0].Delta
//				//t.Log(delta.Role)
//				//if delta.Role == enum.RoleUser {
//				//t.Log(delta.Content)
//				//}
//				if resp.CanGetContent() {
//					t.Log(delta.Content)
//				}
//			case <-done:
//				close(ch)
//				close(done)
//				t.Log("done!")
//				return
//			}
//		}
//	}()
//
//	err = cli.Chat().CompletionsStream(ctx, &moonshot.ChatCompletionsRequest{
//		Model: moonshot.ModelMoonshotV18K,
//		Messages: []*moonshot.ChatCompletionsMessage{
//			{
//				Role:    moonshot.RoleUser,
//				Content: "你好，我叫李雷，1+1等于多少？",
//			},
//		},
//		Temperature: 0.3,
//		Stream:      true,
//	}, ch, done)
//	if err != nil {
//		t.Fatal(err)
//	}
//}

func TestChatStream(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.Chat().CompletionsStream(context.Background(), &moonshot.ChatCompletionsRequest{
		Model: moonshot.ModelMoonshotV18K,
		Messages: []*moonshot.ChatCompletionsMessage{
			{
				Role:    moonshot.RoleUser,
				Content: "你好，我叫李雷，1+1等于多少？",
			},
		},
		Temperature: 0.3,
		Stream:      true,
	})
	if err != nil {
		t.Fatal(err)
	}

	for receive := range resp.Next() {
		msg, err := receive.GetMessage()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			t.Error(err)
			continue
		}

		switch msg.Role {
		case moonshot.RoleSystem:
		case moonshot.RoleUser:
		case moonshot.RoleAssistant:
		default:
			t.Logf("Role: %s,Content: %s", msg.Role, msg.Content)
		}

	}
}
