package moonshot

type IChatCompletionsBuilder interface {
	AddUserContent(content string) IChatCompletionsBuilder
	AddSystemContent(content string) IChatCompletionsBuilder
	AddAssistantContent(content string) IChatCompletionsBuilder
	AddPrompt(prompt string) IChatCompletionsBuilder
	AddMessage(message *ChatCompletionsMessage) IChatCompletionsBuilder

	SetModel(model ChatCompletionsModelID) IChatCompletionsBuilder
	SetTemperature(temperature float64) IChatCompletionsBuilder
	SetStream(enable bool) IChatCompletionsBuilder
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

// NewChatCompletionsBuilder creates a new chat completions builder, or with the given request
func NewChatCompletionsBuilder(req ...ChatCompletionsRequest) IChatCompletionsBuilder {
	builder := &chatCompletionsBuilder{
		req: &ChatCompletionsRequest{},
	}

	if len(req) > 0 {
		builder.req = &req[0]
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

// AddUserContent add a message with the role of user
func (c *chatCompletionsBuilder) AddUserContent(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleUser,
		Content: content,
	})
	return c
}

// AddSystemContent add a message with the role of system
func (c *chatCompletionsBuilder) AddSystemContent(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleSystem,
		Content: content,
	})
	return c
}

// AddAssistantContent adds a message with the role of assistant
func (c *chatCompletionsBuilder) AddAssistantContent(content string) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, &ChatCompletionsMessage{
		Role:    RoleAssistant,
		Content: content,
	})
	return c
}

// AddPrompt is an alias of AddSystemContent
func (c *chatCompletionsBuilder) AddPrompt(prompt string) IChatCompletionsBuilder {
	return c.AddSystemContent(prompt)
}

// AddMessage add ChatCompletionsMessage to the request
func (c *chatCompletionsBuilder) AddMessage(message *ChatCompletionsMessage) IChatCompletionsBuilder {
	c.req.Messages = append(c.req.Messages, message)
	return c
}

// SetModel sets the model of the request
func (c *chatCompletionsBuilder) SetModel(model ChatCompletionsModelID) IChatCompletionsBuilder {
	c.req.Model = model
	return c
}

// SetTemperature sets the temperature of the request
func (c *chatCompletionsBuilder) SetTemperature(temperature float64) IChatCompletionsBuilder {
	c.req.Temperature = temperature
	return c
}

// SetMaxTokens sets the max tokens of the request
func (c *chatCompletionsBuilder) SetMaxTokens(num int) IChatCompletionsBuilder {
	c.req.MaxTokens = num
	return c
}

// SetTopP sets the top p of the request
func (c *chatCompletionsBuilder) SetTopP(num float64) IChatCompletionsBuilder {
	c.req.TopP = num
	return c
}

// SetN sets the n of the request
func (c *chatCompletionsBuilder) SetN(num int) IChatCompletionsBuilder {
	c.req.N = num
	return c
}

// SetPresencePenalty sets the presence penalty of the request
func (c *chatCompletionsBuilder) SetPresencePenalty(num float64) IChatCompletionsBuilder {
	c.req.PresencePenalty = num
	return c
}

// SetFrequencyPenalty sets the frequency penalty of the request
func (c *chatCompletionsBuilder) SetFrequencyPenalty(num float64) IChatCompletionsBuilder {
	c.req.FrequencyPenalty = num
	return c
}

// SetStop sets the stop of the request
func (c *chatCompletionsBuilder) SetStop(stop []string) IChatCompletionsBuilder {
	c.req.Stop = stop
	return c
}

// SetStream sets the stream of the request
func (c *chatCompletionsBuilder) SetStream(enable bool) IChatCompletionsBuilder {
	c.req.Stream = enable
	return c
}

// ToRequest returns the ChatCompletionsRequest
func (c *chatCompletionsBuilder) ToRequest() *ChatCompletionsRequest {
	return c.req
}
