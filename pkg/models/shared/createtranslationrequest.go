package shared

type CreateTranslationRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type CreateTranslationRequest struct {
	File           CreateTranslationRequestFile `multipartForm:"file"`
	Model          string                       `multipartForm:"name=model"`
	Prompt         *string                      `multipartForm:"name=prompt"`
	ResponseFormat *string                      `multipartForm:"name=response_format"`
	Temperature    *float64                     `multipartForm:"name=temperature"`
}
