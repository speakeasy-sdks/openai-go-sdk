package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type DeleteModelRequest struct {
	Model string `pathParam:"style=simple,explode=false,name=model"`
}

type DeleteModelResponse struct {
	ContentType         string
	DeleteModelResponse *shared.DeleteModelResponse
	StatusCode          int
	RawResponse         *http.Response
}
