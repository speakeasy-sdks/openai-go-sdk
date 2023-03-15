package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTranscriptionResponse struct {
	ContentType                 string
	CreateTranscriptionResponse *shared.CreateTranscriptionResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
