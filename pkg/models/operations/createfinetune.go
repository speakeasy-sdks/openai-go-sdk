package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateFineTuneRequest struct {
	Request shared.CreateFineTuneRequest `request:"mediaType=application/json"`
}

type CreateFineTuneResponse struct {
	ContentType string
	FineTune    interface{}
	StatusCode  int
	RawResponse *http.Response
}
