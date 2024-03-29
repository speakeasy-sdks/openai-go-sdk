// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"net/http"
)

type CreateCompletionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateCompletionResponse *shared.CreateCompletionResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateCompletionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCompletionResponse) GetCreateCompletionResponse() *shared.CreateCompletionResponse {
	if o == nil {
		return nil
	}
	return o.CreateCompletionResponse
}

func (o *CreateCompletionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCompletionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
