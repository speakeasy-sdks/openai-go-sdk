package operations

import (
	"net/http"
)

type CreateImageEditResponse struct {
	ContentType    string
	ImagesResponse interface{}
	StatusCode     int
	RawResponse    *http.Response
}
