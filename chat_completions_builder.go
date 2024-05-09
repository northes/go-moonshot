package moonshot

type IChatCompletionsBuilder interface {
	AddUserContent(content string) IChatCompletionsBuilder
	AddSystemContent(content string) IChatCompletionsBuilder
	AddAssistantContent(content string) IChatCompletionsBuilder
	AddPrompt(prompt string) IChatCompletionsBuilder
	AddMessage(message *ChatCompletionsMessage) IChatCompletionsBuilder

	SetModel(model ChatCompletionsModelID) IChatCompletionsBuilder
	SetTemperature(temperature float64) IChatCompletionsBuilder
	SetStream() IChatCompletionsBuilder
	SetMaxTokens(num int) IChatCompletionsBuilder
	SetTopP(num float64) IChatCompletionsBuilder
	SetN(num int) IChatCompletionsBuilder
	SetPresencePenalty(num float64) IChatCompletionsBuilder
	SetFrequencyPenalty(num float64) IChatCompletionsBuilder
	SetStop(stop []string) IChatCompletionsBuilder

	ToRequest() *ChatCompletionsRequest
}

type chatCompletionsBuilder struct {
	req *ChatCompletionsRequest
}

func NewChatCompletionsBuilder(req ...*ChatCompletionsRequest) IChatCompletionsBuilder {
	builder := &chatCompletionsBuilder{
		req: &ChatCompletionsRequest{},
	}

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

func (c *chatCompletionsBuilder) AddUserContent(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleUser,
		Content: content,
	})
	return c
}

func (c *chatCompletionsBuilder) AddSystemContent(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleSystem,
		Content: content,
	})
	return c
}

func (c *chatCompletionsBuilder) AddAssistantContent(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleAssistant,
		Content: content,
	})
	return c
}

func (c *chatCompletionsBuilder) AddPrompt(prompt string) IChatCompletionsBuilder {
	return c.AddSystemContent(prompt)
}

func (c *chatCompletionsBuilder) AddMessage(message *ChatCompletionsMessage) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, message)
	return c
}

func (c *chatCompletionsBuilder) SetModel(model ChatCompletionsModelID) IChatCompletionsBuilder {
	c.req.Model = model
	return c
}

func (c *chatCompletionsBuilder) SetTemperature(temperature float64) IChatCompletionsBuilder {
	c.req.Temperature = temperature
	return c
}

func (c *chatCompletionsBuilder) SetMaxTokens(num int) IChatCompletionsBuilder {
	c.req.MaxTokens = num
	return c
}

func (c *chatCompletionsBuilder) SetTopP(num float64) IChatCompletionsBuilder {
	c.req.TopP = num
	return c
}

func (c *chatCompletionsBuilder) SetN(num int) IChatCompletionsBuilder {
	c.req.N = num
	return c
}

func (c *chatCompletionsBuilder) SetPresencePenalty(num float64) IChatCompletionsBuilder {
	c.req.PresencePenalty = num
	return c
}

func (c *chatCompletionsBuilder) SetFrequencyPenalty(num float64) IChatCompletionsBuilder {
	c.req.FrequencyPenalty = num
	return c
}

func (c *chatCompletionsBuilder) SetStop(stop []string) IChatCompletionsBuilder {
	c.req.Stop = stop
	return c
}

func (c *chatCompletionsBuilder) SetStream() IChatCompletionsBuilder {
	c.req.Stream = true
	return c
}

func (c *chatCompletionsBuilder) ToRequest() *ChatCompletionsRequest {
	return c.req
}
