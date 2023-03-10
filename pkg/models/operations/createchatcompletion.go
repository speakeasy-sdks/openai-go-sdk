package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateChatCompletionRequest struct {
	Request shared.CreateChatCompletionRequest `request:"mediaType=application/json"`
}

type CreateChatCompletionResponse struct {
	ContentType                  string
	CreateChatCompletionResponse *shared.CreateChatCompletionResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
