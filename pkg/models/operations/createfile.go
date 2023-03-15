package operations

import (
	"net/http"
)

type CreateFileResponse struct {
	ContentType string
	OpenAIFile  interface{}
	StatusCode  int
	RawResponse *http.Response
}
