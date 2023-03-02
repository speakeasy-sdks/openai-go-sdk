package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateImageVariationRequest struct {
	Request shared.CreateImageVariationRequest `request:"mediaType=multipart/form-data"`
}

type CreateImageVariationResponse struct {
	ContentType    string
	ImagesResponse interface{}
	StatusCode     int
}
