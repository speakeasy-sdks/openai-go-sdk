package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateClassificationResponse struct {
	ContentType                  string
	CreateClassificationResponse *shared.CreateClassificationResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
