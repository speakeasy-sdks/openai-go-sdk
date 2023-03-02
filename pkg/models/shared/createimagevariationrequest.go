package shared

type CreateImageVariationRequestImage struct {
	Content []byte `multipartForm:"content"`
	Image   string `multipartForm:"name=image"`
}

type CreateImageVariationRequest struct {
	Image          CreateImageVariationRequestImage `multipartForm:"file"`
	N              interface{}                      `multipartForm:"name=n"`
	ResponseFormat interface{}                      `multipartForm:"name=response_format"`
	Size           interface{}                      `multipartForm:"name=size"`
	User           interface{}                      `multipartForm:"name=user"`
}
