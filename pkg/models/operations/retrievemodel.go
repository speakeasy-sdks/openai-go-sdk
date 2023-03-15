package operations

import (
	"net/http"
)

type RetrieveModelRequest struct {
	Model string `pathParam:"style=simple,explode=false,name=model"`
}

type RetrieveModelResponse struct {
	ContentType string
	Model       interface{}
	StatusCode  int
	RawResponse *http.Response
}
