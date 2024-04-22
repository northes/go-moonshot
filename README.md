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
go get github.com/northes/go-moonshot
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
        moonshot.WithAPIKey("xxxx"),
    ),
)
```

### Call API

```go
// List Models
resp, err := cli.Models().List(context.Background())
if err != nil {
    return err
}
```

```go
// Chat completions(stream)
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