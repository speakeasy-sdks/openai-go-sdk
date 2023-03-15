package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateChatCompletionResponse struct {
	ContentType                  string
	CreateChatCompletionResponse *shared.CreateChatCompletionResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
