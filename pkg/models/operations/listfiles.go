package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type ListFilesResponse struct {
	ContentType       string
	ListFilesResponse *shared.ListFilesResponse
	StatusCode        int
}
