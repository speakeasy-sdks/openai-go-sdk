package shared

type CreateModerationRequest struct {
	Input interface{} `json:"input"`
	Model *string     `json:"model,omitempty"`
}
