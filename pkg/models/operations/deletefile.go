// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteFileRequest struct {
	// The ID of the file to use for this request
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DeleteFileResponse struct {
	ContentType string
	// OK
	DeleteFileResponse *shared.DeleteFileResponse
	StatusCode         int
	RawResponse        *http.Response
}
