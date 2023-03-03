package operations

import (
	"net/http"
)

type DownloadFilePathParams struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DownloadFileRequest struct {
	PathParams DownloadFilePathParams
}

type DownloadFileResponse struct {
	ContentType                          string
	StatusCode                           int
	RawResponse                          *http.Response
	DownloadFile200ApplicationJSONString *string
}
