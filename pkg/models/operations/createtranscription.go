package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateTranscriptionRequest struct {
	Request shared.CreateTranscriptionRequest `request:"mediaType=multipart/form-data"`
}

type CreateTranscriptionResponse struct {
	ContentType                 string
	CreateTranscriptionResponse *shared.CreateTranscriptionResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
