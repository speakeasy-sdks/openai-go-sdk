package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateFileRequest struct {
	Request shared.CreateFileRequest `request:"mediaType=multipart/form-data"`
}

type CreateFileResponse struct {
	ContentType string
	OpenAIFile  *interface{}
	StatusCode  int
}
