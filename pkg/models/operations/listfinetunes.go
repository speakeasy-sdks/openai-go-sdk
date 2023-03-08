package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type ListFineTunesResponse struct {
	ContentType           string
	ListFineTunesResponse *shared.ListFineTunesResponse
	StatusCode            int
	RawResponse           *http.Response
}
