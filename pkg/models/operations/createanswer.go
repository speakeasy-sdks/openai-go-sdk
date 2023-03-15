package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateAnswerResponse struct {
	ContentType          string
	CreateAnswerResponse *shared.CreateAnswerResponse
	StatusCode           int
	RawResponse          *http.Response
}
