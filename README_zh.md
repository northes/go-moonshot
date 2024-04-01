## Go Moonshot

[![Go Report Card](https://goreportcard.com/badge/github.com/northes/go-moonshot)](https://goreportcard.com/report/github.com/northes/go-moonshot)
[![Go Reference](https://pkg.go.dev/badge/github.com/northes/go-moonshot.svg)](https://pkg.go.dev/github.com/northes/go-moonshot)
[![codecov](https://codecov.io/gh/northes/go-moonshot/graph/badge.svg?token=81O85CA9KL)](https://codecov.io/gh/northes/go-moonshot)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot.svg?type=shield&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot?ref=badge_shield&issueType=license)
[![License](https://img.shields.io/github/license/northes/go-moonshot)](https://github.com/northes/go-moonshot)

**简体中文** | [English](README.md)

为 [MoonshotAI](https://moonshot.cn) 创建的AI [Kimi](https://kimi.moonshot.cn) 提供的 Go SDK。

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

1. Get a MoonshotAI API Key: [https://platform.moonshot.cn](https://platform.moonshot.cn).
2. Set up key using a configuration file or environment variable.

> :warning: Note: Your API key is sensitive information. Do not share it with anyone.

```go
key, ok := os.LookupEnv("moonshot_key")
if !ok {
return nil, errors.New("missing environment variable: moonshot_key")
}
return moonshot.NewClient(moonshot.NewConfig(
moonshot.SetAPIKey(key),
))
```

### 调用 API

```go
// List Models
resp, err := cli.Models().List(context.Background())
if err != nil {
return err
}
```

## License

This is open-sourced library licensed under the [MIT license](LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot.svg?type=large&issueType=license)](https://app.fossa.com/projects/git%2Bgithub.com%2Fnorthes%2Fgo-moonshot?ref=badge_large&issueType=license)