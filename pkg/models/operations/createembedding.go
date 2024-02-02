// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"net/http"
)

type CreateEmbeddingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateEmbeddingResponse *shared.CreateEmbeddingResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateEmbeddingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateEmbeddingResponse) GetCreateEmbeddingResponse() *shared.CreateEmbeddingResponse {
	if o == nil {
		return nil
	}
	return o.CreateEmbeddingResponse
}

func (o *CreateEmbeddingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateEmbeddingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
