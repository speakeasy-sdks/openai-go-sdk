package operations

import (
	"net/http"
)

type RetrieveFileRequest struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type RetrieveFileResponse struct {
	ContentType string
	OpenAIFile  interface{}
	StatusCode  int
	RawResponse *http.Response
}
