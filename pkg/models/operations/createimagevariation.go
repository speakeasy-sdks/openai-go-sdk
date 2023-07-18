// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
	"net/http"
)

type CreateImageVariationResponse struct {
	ContentType string
	// OK
	ImagesResponse *shared.ImagesResponse
	StatusCode     int
	RawResponse    *http.Response
}

func (o *CreateImageVariationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateImageVariationResponse) GetImagesResponse() *shared.ImagesResponse {
	if o == nil {
		return nil
	}
	return o.ImagesResponse
}

func (o *CreateImageVariationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateImageVariationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
