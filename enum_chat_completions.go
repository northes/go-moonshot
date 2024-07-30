package moonshot

type ChatCompletionsMessageRole string

const (
	RoleSystem       ChatCompletionsMessageRole = "system"
	RoleUser         ChatCompletionsMessageRole = "user"
	RoleAssistant    ChatCompletionsMessageRole = "assistant"
	RoleTool         ChatCompletionsMessageRole = "tool"
	RoleContextCache ChatCompletionsMessageRole = "cache"
)

func (c ChatCompletionsMessageRole) String() string {
	return string(c)
}

type ChatCompletionsModelID string

const (
	ModelMoonshotV18K   ChatCompletionsModelID = "moonshot-v1-8k"
	ModelMoonshotV132K  ChatCompletionsModelID = "moonshot-v1-32k"
	ModelMoonshotV1128K ChatCompletionsModelID = "moonshot-v1-128k"
)

func (c ChatCompletionsModelID) String() string {
	return string(c)
}

type ChatCompletionsModelFamily string

const (
	ModelFamilyMoonshotV1 ChatCompletionsModelFamily = "moonshot-v1"
)

func (c ChatCompletionsModelFamily) String() string {
	return string(c)
}

type ChatCompletionsFinishReason string

const (
	FinishReasonStop      ChatCompletionsFinishReason = "stop"
	FinishReasonLength    ChatCompletionsFinishReason = "length"
	FinishReasonToolCalls ChatCompletionsFinishReason = "tool_calls"
)

func (c ChatCompletionsFinishReason) String() string {
	return string(c)
}

type ChatCompletionsToolType string

const (
	ChatCompletionsToolTypeFunction ChatCompletionsToolType = "function"
)

func (c ChatCompletionsToolType) String() string {
	return string(c)
}

type ChatCompletionsParametersType string

const (
	ChatCompletionsParametersTypeObject ChatCompletionsParametersType = "object"
)

func (c ChatCompletionsParametersType) String() string {
	return string(c)
}
