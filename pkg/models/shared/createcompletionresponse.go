package shared

type CreateCompletionResponseChoicesLogprobs struct {
	TextOffset    []int64                  `json:"text_offset,omitempty"`
	TokenLogprobs []float64                `json:"token_logprobs,omitempty"`
	Tokens        []string                 `json:"tokens,omitempty"`
	TopLogprobs   []map[string]interface{} `json:"top_logprobs,omitempty"`
}

type CreateCompletionResponseChoices struct {
	FinishReason *string                                  `json:"finish_reason,omitempty"`
	Index        *int64                                   `json:"index,omitempty"`
	Logprobs     *CreateCompletionResponseChoicesLogprobs `json:"logprobs,omitempty"`
	Text         *string                                  `json:"text,omitempty"`
}

type CreateCompletionResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type CreateCompletionResponse struct {
	Choices []CreateCompletionResponseChoices `json:"choices"`
	Created int64                             `json:"created"`
	ID      string                            `json:"id"`
	Model   string                            `json:"model"`
	Object  string                            `json:"object"`
	Usage   *CreateCompletionResponseUsage    `json:"usage,omitempty"`
}
