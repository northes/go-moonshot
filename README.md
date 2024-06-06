## Go Moonshot

[![Go Report Card](https://goreportcard.com/badge/github.com/northes/go-moonshot)](https://goreportcard.com/report/github.com/northes/go-moonshot)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![tag](https://img.shields.io/github/tag/northes/go-moonshot.svg)](https://github.com/northes/go-moonshot/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/northes/go-moonshot.svg)](https://pkg.go.dev/github.com/northes/go-moonshot)
[![codecov](https://codecov.io/gh/northes/go-moonshot/graph/badge.svg?token=81O85CA9KL)](https://codecov.io/gh/northes/go-moonshot)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot?ref=badge_shield&issueType=license)
[![License](https://img.shields.io/github/license/northes/go-moonshot)](https://github.com/northes/go-moonshot)

[简体中文](README_zh.md) | **English**

A Go SDK for [Kimi](https://kimi.moonshot.cn) which created
by [MoonshotAI](https://moonshot.cn).

## 🚀 Installation

```bash
go get github.com/northes/go-moonshot@v0.4.3
```

You can find the docs at [go docs](https://pkg.go.dev/github.com/northes/go-moonshot).

## 🤘 Feature

- Easy to use and simple API, chain operation.
- Full API support.
- Predefined enumeration.

##  📄 Supported API

| API                     | Done |
|-------------------------|------|
| Chat Completion         | ✅    |
| Chat Completion(stream) | ✅    |
| List Models             | ✅    |
| List Files              | ✅    |
| Upload File             | ✅    |
| Delete File             | ✅    |
| Get File Info           | ✅    |
| Get File Contents       | ✅    |
| Estimate Token Count    | ✅    |
| User Balance            | ✅    |

## 🥪 Usage

### Initialize client

1. Get a MoonshotAI API Key: [https://platform.moonshot.cn](https://platform.moonshot.cn).
2. Set up key using a configuration file or environment variable.

> :warning: Note: Your API key is sensitive information. Do not share it with anyone.

#### With Only Key

```go
key, ok := os.LookupEnv("MOONSHOT_KEY")
if !ok {
    return errors.New("missing environment variable: moonshot_key")
}

cli, err := moonshot.NewClient(key)
if err != nil {
    return err
}
```

#### With Config

```go
key, ok := os.LookupEnv("MOONSHOT_KEY")
if !ok {
    return errors.New("missing environment variable: moonshot_key")
}

cli, err := moonshot.NewClientWithConfig(
    moonshot.NewConfig(
        moonshot.WithAPIKey(key),
    ),
)
```

### API

#### List Models

```go
resp, err := cli.Models().List(context.Background())
if err != nil {
    return err
}
```

#### Chat Completions

```go
// Use builder to build a request more conveniently
builder := moonshot.NewChatCompletionsBuilder()
builder.AppendPrompt("你是 Kimi，由 Moonshot AI 提供的人工智能助手，你更擅长中文和英文的对话。你会为用户提供安全，有帮助，准确的回答。同时，你会拒绝一切涉及恐怖主义，种族歧视，黄色暴力等问题的回答。Moonshot AI 为专有名词，不可翻译成其他语言。").
	AppendUser("你好，我叫李雷，1+1等于多少？").
	WithTemperature(0.3)

resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
if err != nil {
    return err
}
// {"id":"cmpl-eb8e8474fbae4e42bea9f6bbf38d56ed","object":"chat.completion","created":2647921,"model":"moonshot-v1-8k","choices":[{"index":0,"message":{"role":"assistant","content":"你好，李雷！1+1等于2。这是一个基本的数学加法运算。如果你有任何其他问题或需要帮助，请随时告诉我。"},"finish_reason":"stop"}],"usage":{"prompt_tokens":87,"completion_tokens":31,"total_tokens":118}}

// do something...

// append context
for _, choice := range resp.Choices {
    builder.AppendMessage(choice.Message)
}

builder.AppendUser("在这个基础上再加3等于多少")

resp, err := cli.Chat().Completions(ctx, builder.ToRequest())
if err != nil {
    return err
}
// {"id":"cmpl-a7b938eaddc04fbf85fe578a980040ac","object":"chat.completion","created":5455796,"model":"moonshot-v1-8k","choices":[{"index":0,"message":{"role":"assistant","content":"在这个基础上，即1+1=2的结果上再加3，等于5。所以，2+3=5。"},"finish_reason":"stop"}],"usage":{"prompt_tokens":131,"completion_tokens":26,"total_tokens":157}}
```

#### Chat completions with stream

```go
// use struct
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
    return err
}

for receive := range resp.Receive() {
    msg, err := receive.GetMessage()
    if err != nil {
        if errors.Is(err, io.EOF) {
            break
        }
        break
    }
    switch msg.Role {
        case moonshot.RoleSystem,moonshot.RoleUser,moonshot.RoleAssistant:
        // do something...
        default:
        // do something...
    }
}
```

## 🤝  Missing a Feature?

Feel free to open a new issue, or contact me.

## 📘 License

This is open-sourced library licensed under the [MIT license](LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot.svg?type=large&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot?ref=badge_large&issueType=license)