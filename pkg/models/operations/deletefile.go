// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type DeleteFileRequest struct {
	// The ID of the file to use for this request.
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

func (o *DeleteFileRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

type DeleteFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DeleteFileResponse *shared.DeleteFileResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteFileResponse) GetDeleteFileResponse() *shared.DeleteFileResponse {
	if o == nil {
		return nil
	}
	return o.DeleteFileResponse
}

func (o *DeleteFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
