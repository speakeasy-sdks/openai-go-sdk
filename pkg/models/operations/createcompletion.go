package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCompletionRequest struct {
	Request shared.CreateCompletionRequest `request:"mediaType=application/json"`
}

type CreateCompletionResponse struct {
	ContentType              string
	CreateCompletionResponse *shared.CreateCompletionResponse
	StatusCode               int
	RawResponse              *http.Response
}
