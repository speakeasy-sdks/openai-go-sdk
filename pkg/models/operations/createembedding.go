package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateEmbeddingRequest struct {
	Request shared.CreateEmbeddingRequest `request:"mediaType=application/json"`
}

type CreateEmbeddingResponse struct {
	ContentType             string
	CreateEmbeddingResponse *shared.CreateEmbeddingResponse
	StatusCode              int
}
