package moonshot_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/northes/go-moonshot"
	"github.com/northes/go-moonshot/internal/httpx"
	"github.com/northes/go-moonshot/test"
)

func TestChat(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	builder := moonshot.NewChatCompletionsBuilder()
	builder.AddPrompt("你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。").
		AddUserContent("你好，我叫李雷，1+1等于多少？").
		SetTemperature(0.3)

	resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(test.MarshalJsonToStringX(resp))
	/*
		{"id":"chatcmpl-dafad118ba6a4d1bb3e10be1734c6213","object":"chat.completion","created":15893254,"model":"moonshot-v1-8k","choices":[{"index":0,"message":{"Role":"assistant","Content":"你好，李雷！1+1等于2。如果你有更复杂的数学问题或者其他问题，也可以随时问我。"},"finish_reason":"stop"}],"usage":{"prompt_tokens":83,"completion_tokens":25,"total_tokens":108}}
	*/
}

func TestChatStream(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}

	builder := moonshot.NewChatCompletionsBuilder()
	builder.SetModel(moonshot.ModelMoonshotV18K).
		AddUserContent("你好，我叫李雷，1+1等于多少？").
		SetStream(true)

	resp, err := cli.Chat().CompletionsStream(context.Background(), builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}

	for receive := range resp.Receive() {
		msg, err := receive.GetMessage()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// Finish usage
				if len(receive.Choices) != 0 {
					choice := receive.Choices[0]
					if choice.FinishReason == moonshot.FinishReasonStop && choice.Usage != nil {
						t.Logf("Finish Usage: PromptTokens: %d, CompletionTokens: %d, TotalTokens: %d", choice.Usage.PromptTokens, choice.Usage.CompletionTokens, choice.Usage.TotalTokens)
					}
				}
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

func TestChatWithContext(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	builder := moonshot.NewChatCompletionsBuilder()
	builder.AddPrompt("你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。").
		AddUserContent("你好，我叫李雷，1+1等于多少？").
		SetTemperature(0.3)

	resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(test.MarshalJsonToStringX(resp))

	for _, choice := range resp.Choices {
		builder.AddMessage(choice.Message)
	}

	builder.AddUserContent("在这个基础上再加3等于多少")

	resp, err = cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(test.MarshalJsonToStringX(resp))
}

func TestPartialMode(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	const leadingText = "{"

	builder := moonshot.NewChatCompletionsBuilder()
	builder.AddSystemContent("请从产品描述中提取名称、尺寸、价格和颜色，并在一个 JSON 对象中输出。")
	builder.AddUserContent("大米 SmartHome Mini 是一款小巧的智能家居助手，有黑色和银色两种颜色，售价为 998 元，尺寸为 256 x 128 x 128mm。可让您通过语音或应用程序控制灯光、恒温器和其他联网设备，无论您将它放在家中的任何位置。")
	builder.AddAssistantContent(leadingText, true)
	builder.SetTemperature(0.3)

	resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}
	message, err := resp.GetMessage()
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, strings.HasPrefix(message.Content, leadingText), false, "message content should not start with '{'")

	jsonStr := fmt.Sprintf("%s%s", leadingText, message.Content)

	t.Log(jsonStr)
}

func TestUseTools(t *testing.T) {
	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	const functionName = "locate_the_ip_address"
	const ip = "15.202.141.157"

	builder := moonshot.NewChatCompletionsBuilder()
	builder.AddUserContent(fmt.Sprintf("IP地址`%s`的定位是哪里？", ip))
	builder.SetTool(&moonshot.ChatCompletionsTool{
		Type: moonshot.ChatCompletionsToolTypeFunction,
		Function: &moonshot.ChatCompletionsToolFunction{
			Name:        functionName,
			Description: "定位IP地址，查询传入的IP地址的位置信息",
			Parameters: &moonshot.ChatCompletionsToolFunctionParameters{
				Type: moonshot.ChatCompletionsParametersTypeObject,
				Properties: map[string]*moonshot.ChatCompletionsToolFunctionProperties{
					"ip_address": {
						Type:        "string",
						Description: "IP地址",
					},
				},
			},
		},
	})

	resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}

	// check tool calls
	if len(resp.Choices) != 0 {
		if resp.Choices[0].FinishReason == moonshot.FinishReasonToolCalls {
			for _, tool := range resp.Choices[0].Message.ToolCalls {
				t.Logf("should tool calls: %v", test.MarshalJsonToStringX(tool))
				if strings.HasPrefix(tool.ID, functionName) {
					// tool calls
					ipInfo, err := IPLocate(ip)
					if err != nil {
						t.Fatal(err)
					}

					builder.AddMessageFromChoices(resp.Choices)

					t.Logf("tool calls result: %s", ipInfo)

					builder.AddToolContent(ipInfo, functionName, tool.ID)
				}
			}
		}
	}

	resp, err = cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}

	finalMessage, err := resp.GetMessage()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test.MarshalJsonToStringX(finalMessage.Content))
	// IP地址`15.202.141.157`的定位结果显示，该IP地址位于美国华盛顿州，属于Amazon.com。具体的城市信息没有提供，但经纬度坐标为(-95.71289, 37.090245)。这些信息来源于高德地图，最后更新时间为2024年6月6日。
}

func IPLocate(ip string) (string, error) {
	response, err := httpx.NewClient(fmt.Sprintf("https://apihut.co/ip/%s", ip)).Get(context.Background())
	if err != nil {
		return "", err
	}
	defer func() {
		_ = response.Raw().Body.Close()
	}()

	body, err := io.ReadAll(response.Raw().Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func TestBuiltinFunctionWebSearch(t *testing.T) {
	if test.IsGithubActions() {
		return
	}

	cli, err := NewTestClient()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	builder := moonshot.NewChatCompletionsBuilder()
	builder.SetModel(moonshot.ModelMoonshotV1128K)
	builder.AddUserContent("请搜索 Moonshot AI Context Caching 技术，并告诉我它是什么。")
	builder.SetTool(&moonshot.ChatCompletionsTool{
		Type: moonshot.ChatCompletionsToolTypeBuiltinFunction,
		Function: &moonshot.ChatCompletionsToolFunction{
			Name: moonshot.BuiltinFunctionWebSearch,
		},
	})

	resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Choices) != 0 {
		choice := resp.Choices[0]
		if choice.FinishReason == moonshot.FinishReasonToolCalls {
			for _, tool := range choice.Message.ToolCalls {
				t.Logf("tool calls: %v", test.MarshalJsonToStringX(tool))
				if tool.Function.Name == moonshot.BuiltinFunctionWebSearch {
					// web search
					arguments := new(moonshot.ChatCompletionsToolBuiltinFunctionWebSearchArguments)
					if err = json.Unmarshal([]byte(tool.Function.Arguments), arguments); err != nil {
						t.Errorf("unmarshal tool arguments error: %v", err)
						continue
					}

					t.Logf("tool calls result: search_id: %s, total_tokens: %d", arguments.SearchResult.SearchId, arguments.Usage.TotalTokens)

					builder.AddMessageFromChoices(resp.Choices)
					builder.AddToolContent(tool.Function.Arguments, tool.Function.Name, tool.ID)
				}
			}
		}
	}

	t.Logf("builder: %v", test.MarshalJsonToStringX(builder.ToRequest()))

	resp, err = cli.Chat().Completions(ctx, builder.ToRequest())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(test.MarshalJsonToStringX(resp))
}
