package shared

type CreateEditRequest struct {
	Input       *string     `json:"input,omitempty"`
	Instruction string      `json:"instruction"`
	Model       interface{} `json:"model"`
	N           *int64      `json:"n,omitempty"`
	Temperature *float64    `json:"temperature,omitempty"`
	TopP        *float64    `json:"top_p,omitempty"`
}
