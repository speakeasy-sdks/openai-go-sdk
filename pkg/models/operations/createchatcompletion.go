// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type CreateChatCompletionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateChatCompletionResponse *shared.CreateChatCompletionResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateChatCompletionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateChatCompletionResponse) GetCreateChatCompletionResponse() *shared.CreateChatCompletionResponse {
	if o == nil {
		return nil
	}
	return o.CreateChatCompletionResponse
}

func (o *CreateChatCompletionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateChatCompletionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
