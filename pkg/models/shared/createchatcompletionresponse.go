package shared

type CreateChatCompletionResponseChoices struct {
	FinishReason *string                        `json:"finish_reason,omitempty"`
	Index        *int64                         `json:"index,omitempty"`
	Message      *ChatCompletionResponseMessage `json:"message,omitempty"`
}

type CreateChatCompletionResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type CreateChatCompletionResponse struct {
	Choices []CreateChatCompletionResponseChoices `json:"choices"`
	Created int64                                 `json:"created"`
	ID      string                                `json:"id"`
	Model   string                                `json:"model"`
	Object  string                                `json:"object"`
	Usage   *CreateChatCompletionResponseUsage    `json:"usage,omitempty"`
}
