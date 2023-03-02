package shared

type CreateCompletionRequest struct {
	BestOf           *int64                 `json:"best_of,omitempty"`
	Echo             *bool                  `json:"echo,omitempty"`
	FrequencyPenalty *float64               `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`
	Logprobs         *int64                 `json:"logprobs,omitempty"`
	MaxTokens        *int64                 `json:"max_tokens,omitempty"`
	Model            string                 `json:"model"`
	N                *int64                 `json:"n,omitempty"`
	PresencePenalty  *float64               `json:"presence_penalty,omitempty"`
	Prompt           interface{}            `json:"prompt,omitempty"`
	Stop             interface{}            `json:"stop,omitempty"`
	Stream           *bool                  `json:"stream,omitempty"`
	Suffix           *string                `json:"suffix,omitempty"`
	Temperature      *float64               `json:"temperature,omitempty"`
	TopP             *float64               `json:"top_p,omitempty"`
	User             *string                `json:"user,omitempty"`
}
