package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateFineTuneRequest struct {
	Request shared.CreateFineTuneRequest `request:"mediaType=application/json"`
}

type CreateFineTuneResponse struct {
	ContentType string
	FineTune    interface{}
	StatusCode  int
}
