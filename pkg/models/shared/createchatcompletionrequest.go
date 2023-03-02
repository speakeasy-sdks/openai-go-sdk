package shared

type CreateChatCompletionRequest struct {
	FrequencyPenalty *float64                       `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]interface{}         `json:"logit_bias,omitempty"`
	MaxTokens        *int64                         `json:"max_tokens,omitempty"`
	Messages         []ChatCompletionRequestMessage `json:"messages"`
	Model            string                         `json:"model"`
	N                *int64                         `json:"n,omitempty"`
	PresencePenalty  *float64                       `json:"presence_penalty,omitempty"`
	Stop             interface{}                    `json:"stop,omitempty"`
	Stream           *bool                          `json:"stream,omitempty"`
	Temperature      *float64                       `json:"temperature,omitempty"`
	TopP             *float64                       `json:"top_p,omitempty"`
	User             interface{}                    `json:"user,omitempty"`
}
