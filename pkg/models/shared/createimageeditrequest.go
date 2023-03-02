package shared

type CreateImageEditRequestImage struct {
	Content []byte `multipartForm:"content"`
	Image   string `multipartForm:"name=image"`
}

type CreateImageEditRequestMask struct {
	Content []byte `multipartForm:"content"`
	Mask    string `multipartForm:"name=mask"`
}

type CreateImageEditRequest struct {
	Image          CreateImageEditRequestImage `multipartForm:"file"`
	Mask           *CreateImageEditRequestMask `multipartForm:"file"`
	N              interface{}                 `multipartForm:"name=n"`
	Prompt         string                      `multipartForm:"name=prompt"`
	ResponseFormat interface{}                 `multipartForm:"name=response_format"`
	Size           interface{}                 `multipartForm:"name=size"`
	User           interface{}                 `multipartForm:"name=user"`
}
