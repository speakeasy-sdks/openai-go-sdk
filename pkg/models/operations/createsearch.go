package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSearchRequest struct {
	CreateSearchRequest shared.CreateSearchRequest `request:"mediaType=application/json"`
	EngineID            string                     `pathParam:"style=simple,explode=false,name=engine_id"`
}

type CreateSearchResponse struct {
	ContentType          string
	CreateSearchResponse *shared.CreateSearchResponse
	StatusCode           int
	RawResponse          *http.Response
}
