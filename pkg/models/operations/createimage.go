package operations

import (
	"net/http"
)

type CreateImageResponse struct {
	ContentType    string
	ImagesResponse interface{}
	StatusCode     int
	RawResponse    *http.Response
}
