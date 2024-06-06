package moonshot

type ChatCompletionsTool struct {
	Type     ChatCompletionsToolType      `json:"type"`
	Function *ChatCompletionsToolFunction `json:"function"`
}

type ChatCompletionsToolFunction struct {
	Name        string                                 `json:"name"`
	Description string                                 `json:"description"`
	Parameters  *ChatCompletionsToolFunctionParameters `json:"parameters"`
}

type ChatCompletionsToolFunctionParameters struct {
	Type       ChatCompletionsParametersType                     `json:"type"`
	Properties map[string]*ChatCompletionsToolFunctionProperties `json:"properties"`
	Required   []string                                          `json:"required,omitempty"`
}

type ChatCompletionsToolFunctionProperties struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
}

type ChatCompletionsResponseToolCalls struct {
	Index    int64                                     `json:"index"`
	ID       string                                    `json:"id"`
	Type     string                                    `json:"type"`
	Function *ChatCompletionsResponseToolCallsFunction `json:"function"`
}

type ChatCompletionsResponseToolCallsFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}
