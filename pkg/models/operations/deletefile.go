package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteFileRequest struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DeleteFileResponse struct {
	ContentType        string
	DeleteFileResponse *shared.DeleteFileResponse
	StatusCode         int
	RawResponse        *http.Response
}
