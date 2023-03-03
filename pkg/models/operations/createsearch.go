package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSearchPathParams struct {
	EngineID string `pathParam:"style=simple,explode=false,name=engine_id"`
}

type CreateSearchRequest struct {
	PathParams CreateSearchPathParams
	Request    shared.CreateSearchRequest `request:"mediaType=application/json"`
}

type CreateSearchResponse struct {
	ContentType          string
	CreateSearchResponse *shared.CreateSearchResponse
	StatusCode           int
	RawResponse          *http.Response
}
