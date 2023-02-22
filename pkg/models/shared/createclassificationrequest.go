package shared

type CreateClassificationRequest struct {
	Examples       [][]string   `json:"examples,omitempty"`
	Expand         *interface{} `json:"expand,omitempty"`
	File           *string      `json:"file,omitempty"`
	Labels         []string     `json:"labels,omitempty"`
	LogitBias      *interface{} `json:"logit_bias,omitempty"`
	Logprobs       *interface{} `json:"logprobs,omitempty"`
	MaxExamples    *int64       `json:"max_examples,omitempty"`
	Model          interface{}  `json:"model"`
	Query          string       `json:"query"`
	ReturnMetadata *interface{} `json:"return_metadata,omitempty"`
	ReturnPrompt   *interface{} `json:"return_prompt,omitempty"`
	SearchModel    *interface{} `json:"search_model,omitempty"`
	Temperature    *float64     `json:"temperature,omitempty"`
	User           *interface{} `json:"user,omitempty"`
}
