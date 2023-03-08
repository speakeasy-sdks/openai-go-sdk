package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type ListEnginesResponse struct {
	ContentType         string
	ListEnginesResponse *shared.ListEnginesResponse
	StatusCode          int
	RawResponse         *http.Response
}
