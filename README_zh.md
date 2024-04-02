## Go Moonshot

[![Go Report Card](https://goreportcard.com/badge/github.com/northes/go-moonshot)](https://goreportcard.com/report/github.com/northes/go-moonshot)
[![Go Reference](https://pkg.go.dev/badge/github.com/northes/go-moonshot.svg)](https://pkg.go.dev/github.com/northes/go-moonshot)
[![codecov](https://codecov.io/gh/northes/go-moonshot/graph/badge.svg?token=81O85CA9KL)](https://codecov.io/gh/northes/go-moonshot)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot?ref=badge_shield&issueType=license)
[![License](https://img.shields.io/github/license/northes/go-moonshot)](https://github.com/northes/go-moonshot)

**简体中文** | [English](README.md)

为 [月之暗面/MoonshotAI](https://moonshot.cn) 创建的AI [Kimi](https://kimi.moonshot.cn) 提供的 Go SDK。

## 特性

- 人体工学的API，链式操作
- 完全的API支持
- 预定义枚举

## 已支持的 API

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

## 使用

### 初始化

1. 申请一个 MoonshotAI API Key: [https://platform.moonshot.cn](https://platform.moonshot.cn)
2. 使用配置文件或环境变量设置密钥

> :warning: 注意：您的API密钥是敏感信息。
不要和任何人分享。

#### 仅使用Key

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

#### 使用配置

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

### 调用 API

```go
// 列出可用模型
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

## License

This is open-sourced library licensed under the [MIT license](LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot.svg?type=large&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot?ref=badge_large&issueType=license)