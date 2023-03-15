package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateCompletionResponse struct {
	ContentType              string
	CreateCompletionResponse *shared.CreateCompletionResponse
	StatusCode               int
	RawResponse              *http.Response
}
