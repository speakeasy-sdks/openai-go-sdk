// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type RetrieveModelRequest struct {
	// The ID of the model to use for this request
	Model string `pathParam:"style=simple,explode=false,name=model"`
}

func (o *RetrieveModelRequest) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

type RetrieveModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	Model *shared.Model
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RetrieveModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetrieveModelResponse) GetModel() *shared.Model {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *RetrieveModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetrieveModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
