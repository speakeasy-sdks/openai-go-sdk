package shared

type CreateAnswerRequest struct {
	Documents       []string      `json:"documents,omitempty"`
	Examples        [][]string    `json:"examples"`
	ExamplesContext string        `json:"examples_context"`
	Expand          []interface{} `json:"expand,omitempty"`
	File            *string       `json:"file,omitempty"`
	LogitBias       interface{}   `json:"logit_bias,omitempty"`
	Logprobs        *int64        `json:"logprobs,omitempty"`
	MaxRerank       *int64        `json:"max_rerank,omitempty"`
	MaxTokens       *int64        `json:"max_tokens,omitempty"`
	Model           string        `json:"model"`
	N               *int64        `json:"n,omitempty"`
	Question        string        `json:"question"`
	ReturnMetadata  interface{}   `json:"return_metadata,omitempty"`
	ReturnPrompt    *bool         `json:"return_prompt,omitempty"`
	SearchModel     *string       `json:"search_model,omitempty"`
	Stop            interface{}   `json:"stop,omitempty"`
	Temperature     *float64      `json:"temperature,omitempty"`
	User            interface{}   `json:"user,omitempty"`
}
