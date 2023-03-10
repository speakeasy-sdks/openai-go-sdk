package shared

type CreateEditResponseChoicesLogprobs struct {
	TextOffset    []int64                  `json:"text_offset,omitempty"`
	TokenLogprobs []float64                `json:"token_logprobs,omitempty"`
	Tokens        []string                 `json:"tokens,omitempty"`
	TopLogprobs   []map[string]interface{} `json:"top_logprobs,omitempty"`
}

type CreateEditResponseChoices struct {
	FinishReason *string                            `json:"finish_reason,omitempty"`
	Index        *int64                             `json:"index,omitempty"`
	Logprobs     *CreateEditResponseChoicesLogprobs `json:"logprobs,omitempty"`
	Text         *string                            `json:"text,omitempty"`
}

type CreateEditResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type CreateEditResponse struct {
	Choices []CreateEditResponseChoices `json:"choices"`
	Created int64                       `json:"created"`
	Object  string                      `json:"object"`
	Usage   CreateEditResponseUsage     `json:"usage"`
}
