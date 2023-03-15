package operations

import (
	"net/http"
)

type RetrieveFineTuneRequest struct {
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
}

type RetrieveFineTuneResponse struct {
	ContentType string
	FineTune    interface{}
	StatusCode  int
	RawResponse *http.Response
}
