package shared

type CreateSearchResponseData struct {
	Document *int64   `json:"document,omitempty"`
	Object   *string  `json:"object,omitempty"`
	Score    *float64 `json:"score,omitempty"`
}

type CreateSearchResponse struct {
	Data   []CreateSearchResponseData `json:"data,omitempty"`
	Model  *string                    `json:"model,omitempty"`
	Object *string                    `json:"object,omitempty"`
}
