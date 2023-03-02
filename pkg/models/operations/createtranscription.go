package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateTranscriptionRequest struct {
	Request shared.CreateTranscriptionRequest `request:"mediaType=multipart/form-data"`
}

type CreateTranscriptionResponse struct {
	ContentType                 string
	CreateTranscriptionResponse *shared.CreateTranscriptionResponse
	StatusCode                  int
}
