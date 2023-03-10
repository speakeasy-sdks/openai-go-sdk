package operations

import (
	"net/http"
)

type RetrieveModelPathParams struct {
	Model string `pathParam:"style=simple,explode=false,name=model"`
}

type RetrieveModelRequest struct {
	PathParams RetrieveModelPathParams
}

type RetrieveModelResponse struct {
	ContentType string
	Model       interface{}
	StatusCode  int
	RawResponse *http.Response
}
