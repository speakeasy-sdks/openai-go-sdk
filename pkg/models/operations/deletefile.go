package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type DeleteFilePathParams struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DeleteFileRequest struct {
	PathParams DeleteFilePathParams
}

type DeleteFileResponse struct {
	ContentType        string
	DeleteFileResponse *shared.DeleteFileResponse
	StatusCode         int
}
