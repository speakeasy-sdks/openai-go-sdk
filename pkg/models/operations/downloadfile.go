package operations

type DownloadFilePathParams struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DownloadFileRequest struct {
	PathParams DownloadFilePathParams
}

type DownloadFileResponse struct {
	ContentType                          string
	StatusCode                           int
	DownloadFile200ApplicationJSONString *string
}
