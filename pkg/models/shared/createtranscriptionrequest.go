package shared

type CreateTranscriptionRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type CreateTranscriptionRequest struct {
	File           CreateTranscriptionRequestFile `multipartForm:"file"`
	Language       *string                        `multipartForm:"name=language"`
	Model          string                         `multipartForm:"name=model"`
	Prompt         *string                        `multipartForm:"name=prompt"`
	ResponseFormat *string                        `multipartForm:"name=response_format"`
	Temperature    *float64                       `multipartForm:"name=temperature"`
}
