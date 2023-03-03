package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateImageEditRequest struct {
	Request shared.CreateImageEditRequest `request:"mediaType=multipart/form-data"`
}

type CreateImageEditResponse struct {
	ContentType    string
	ImagesResponse interface{}
	StatusCode     int
	RawResponse    *http.Response
}
