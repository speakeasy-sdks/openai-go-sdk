package operations

type RetrieveFilePathParams struct {
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type RetrieveFileRequest struct {
	PathParams RetrieveFilePathParams
}

type RetrieveFileResponse struct {
	ContentType string
	OpenAIFile  *interface{}
	StatusCode  int
}
