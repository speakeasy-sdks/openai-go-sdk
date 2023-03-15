package operations

import (
	"net/http"
)

type CreateFineTuneResponse struct {
	ContentType string
	FineTune    interface{}
	StatusCode  int
	RawResponse *http.Response
}
