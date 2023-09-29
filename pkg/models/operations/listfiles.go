// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
	"net/http"
)

type ListFilesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListFilesResponse *shared.ListFilesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListFilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFilesResponse) GetListFilesResponse() *shared.ListFilesResponse {
	if o == nil {
		return nil
	}
	return o.ListFilesResponse
}

func (o *ListFilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
