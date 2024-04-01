package moonshot

type ChatCompletionsMessageRole string

const (
	RoleSystem    ChatCompletionsMessageRole = "system"
	RoleUser      ChatCompletionsMessageRole = "user"
	RoleAssistant ChatCompletionsMessageRole = "assistant"
)

type ChatCompletionsModelID string

const (
	ModelMoonshotV18K   ChatCompletionsModelID = "moonshot-v1-8k"
	ModelMoonshotV132K  ChatCompletionsModelID = "moonshot-v1-32k"
	ModelMoonshotV1128K ChatCompletionsModelID = "moonshot-v1-128k"
)

type ChatCompletionsFinishReason string

const (
	FinishReasonStop   ChatCompletionsFinishReason = "stop"
	FinishReasonLength ChatCompletionsFinishReason = "length"
)
