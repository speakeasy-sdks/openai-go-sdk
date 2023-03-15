package operations

import (
	"net/http"
)

type CreateImageVariationResponse struct {
	ContentType    string
	ImagesResponse interface{}
	StatusCode     int
	RawResponse    *http.Response
}
