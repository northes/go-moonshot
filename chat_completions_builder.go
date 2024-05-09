package moonshot

type IChatCompletionsBuilder interface {
	AppendUser(content string) IChatCompletionsBuilder
	AppendPrompt(prompt string) IChatCompletionsBuilder
	AppendSystem(content string) IChatCompletionsBuilder
	AppendAssistant(content string) IChatCompletionsBuilder
	AppendMessage(message *ChatCompletionsMessage) IChatCompletionsBuilder

	WithModel(model ChatCompletionsModelID) IChatCompletionsBuilder
	WithTemperature(temperature float64) IChatCompletionsBuilder
	WithStream() IChatCompletionsBuilder

	ToRequest() *ChatCompletionsRequest
}

type chatCompletionsBuilder struct {
	req *ChatCompletionsRequest
}

func NewChatCompletionsBuilder(req ...*ChatCompletionsRequest) IChatCompletionsBuilder {
	builder := new(chatCompletionsBuilder)

	if len(req) > 0 && req[0] != nil {
		builder.req = req[0]
	}

	builder.preCheck()

	return builder
}

func (c *chatCompletionsBuilder) preCheck() {
	if c.req.Messages == nil {
		c.req.Messages = make([]*ChatCompletionsMessage, 0)
	}
	if c.req.Model == "" {
		c.req.Model = ModelMoonshotV18K
	}
}

func (c *chatCompletionsBuilder) AppendUser(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleUser,
		Content: content,
	})
	return c
}

func (c *chatCompletionsBuilder) AppendSystem(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleSystem,
		Content: content,
	})
	return c
}

func (c *chatCompletionsBuilder) AppendAssistant(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleAssistant,
		Content: content,
	})
	return c
}

func (c *chatCompletionsBuilder) AppendPrompt(prompt string) IChatCompletionsBuilder {
	return c.AppendSystem(prompt)
}

func (c *chatCompletionsBuilder) AppendMessage(message *ChatCompletionsMessage) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, message)
	return c
}

func (c *chatCompletionsBuilder) WithModel(model ChatCompletionsModelID) IChatCompletionsBuilder {
	c.req.Model = model
	return c
}

func (c *chatCompletionsBuilder) WithTemperature(temperature float64) IChatCompletionsBuilder {
	c.req.Temperature = temperature
	return c
}

func (c *chatCompletionsBuilder) WithStream() IChatCompletionsBuilder {
	c.req.Stream = true
	return c
}

func (c *chatCompletionsBuilder) ToRequest() *ChatCompletionsRequest {
	return c.req
}
