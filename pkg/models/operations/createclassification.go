package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateClassificationRequest struct {
	Request shared.CreateClassificationRequest `request:"mediaType=application/json"`
}

type CreateClassificationResponse struct {
	ContentType                  string
	CreateClassificationResponse *shared.CreateClassificationResponse
	StatusCode                   int
}
