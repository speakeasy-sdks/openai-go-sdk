package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateImageRequest struct {
	Request shared.CreateImageRequest `request:"mediaType=application/json"`
}

type CreateImageResponse struct {
	ContentType    string
	ImagesResponse interface{}
	StatusCode     int
	RawResponse    *http.Response
}
