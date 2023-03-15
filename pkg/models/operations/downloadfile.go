package operations

import (
	"net/http"
)

type DownloadFileRequest struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DownloadFileResponse struct {
	ContentType                          string
	StatusCode                           int
	RawResponse                          *http.Response
	DownloadFile200ApplicationJSONString *string
}
