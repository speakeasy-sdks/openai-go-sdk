package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type ListFilesResponse struct {
	ContentType       string
	ListFilesResponse *shared.ListFilesResponse
	StatusCode        int
	RawResponse       *http.Response
}
