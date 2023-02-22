package shared

type CreateEmbeddingRequest struct {
	Input interface{}  `json:"input"`
	Model interface{}  `json:"model"`
	User  *interface{} `json:"user,omitempty"`
}
