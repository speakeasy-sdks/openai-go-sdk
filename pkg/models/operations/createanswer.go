package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateAnswerRequest struct {
	Request shared.CreateAnswerRequest `request:"mediaType=application/json"`
}

type CreateAnswerResponse struct {
	ContentType          string
	CreateAnswerResponse *shared.CreateAnswerResponse
	StatusCode           int
}
