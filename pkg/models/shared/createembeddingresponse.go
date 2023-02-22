package shared

type CreateEmbeddingResponseData struct {
	Embedding []float64 `json:"embedding"`
	Index     int64     `json:"index"`
	Object    string    `json:"object"`
}

type CreateEmbeddingResponseUsage struct {
	PromptTokens int64 `json:"prompt_tokens"`
	TotalTokens  int64 `json:"total_tokens"`
}

type CreateEmbeddingResponse struct {
	Data   []CreateEmbeddingResponseData `json:"data"`
	Model  string                        `json:"model"`
	Object string                        `json:"object"`
	Usage  CreateEmbeddingResponseUsage  `json:"usage"`
}
