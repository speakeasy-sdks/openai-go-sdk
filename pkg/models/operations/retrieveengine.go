package operations

import (
	"net/http"
)

type RetrieveEngineRequest struct {
	EngineID string `pathParam:"style=simple,explode=false,name=engine_id"`
}

type RetrieveEngineResponse struct {
	ContentType string
	Engine      interface{}
	StatusCode  int
	RawResponse *http.Response
}
