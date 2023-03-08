package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateClassificationRequest struct {
	Request shared.CreateClassificationRequest `request:"mediaType=application/json"`
}

type CreateClassificationResponse struct {
	ContentType                  string
	CreateClassificationResponse *shared.CreateClassificationResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
