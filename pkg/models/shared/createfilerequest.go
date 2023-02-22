package shared

type CreateFileRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type CreateFileRequest struct {
	File    CreateFileRequestFile `multipartForm:"file"`
	Purpose string                `multipartForm:"name=purpose"`
}
