package shared

type CreateSearchRequest struct {
	Documents      []string     `json:"documents,omitempty"`
	File           *string      `json:"file,omitempty"`
	MaxRerank      *int64       `json:"max_rerank,omitempty"`
	Query          string       `json:"query"`
	ReturnMetadata *bool        `json:"return_metadata,omitempty"`
	User           *interface{} `json:"user,omitempty"`
}
