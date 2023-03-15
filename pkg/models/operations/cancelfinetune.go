package operations

import (
	"net/http"
)

type CancelFineTuneRequest struct {
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
}

type CancelFineTuneResponse struct {
	ContentType string
	FineTune    interface{}
	StatusCode  int
	RawResponse *http.Response
}
