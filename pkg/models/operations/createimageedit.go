// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"net/http"
)

type CreateImageEditResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ImagesResponse *shared.ImagesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateImageEditResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateImageEditResponse) GetImagesResponse() *shared.ImagesResponse {
	if o == nil {
		return nil
	}
	return o.ImagesResponse
}

func (o *CreateImageEditResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateImageEditResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
