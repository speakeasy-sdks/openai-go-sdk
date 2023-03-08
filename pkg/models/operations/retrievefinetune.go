package operations

import (
	"net/http"
)

type RetrieveFineTunePathParams struct {
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
}

type RetrieveFineTuneRequest struct {
	PathParams RetrieveFineTunePathParams
}

type RetrieveFineTuneResponse struct {
	ContentType string
	FineTune    interface{}
	StatusCode  int
	RawResponse *http.Response
}
