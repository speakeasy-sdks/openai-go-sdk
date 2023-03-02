package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateTranslationRequest struct {
	Request shared.CreateTranslationRequest `request:"mediaType=multipart/form-data"`
}

type CreateTranslationResponse struct {
	ContentType               string
	CreateTranslationResponse *shared.CreateTranslationResponse
	StatusCode                int
}
