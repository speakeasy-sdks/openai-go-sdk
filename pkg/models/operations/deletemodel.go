// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"net/http"
)

type DeleteModelRequest struct {
	// The model to delete
	Model string `pathParam:"style=simple,explode=false,name=model"`
}

func (o *DeleteModelRequest) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

type DeleteModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DeleteModelResponse *shared.DeleteModelResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteModelResponse) GetDeleteModelResponse() *shared.DeleteModelResponse {
	if o == nil {
		return nil
	}
	return o.DeleteModelResponse
}

func (o *DeleteModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
