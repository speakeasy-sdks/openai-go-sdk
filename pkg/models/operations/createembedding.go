package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateEmbeddingResponse struct {
	ContentType             string
	CreateEmbeddingResponse *shared.CreateEmbeddingResponse
	StatusCode              int
	RawResponse             *http.Response
}
