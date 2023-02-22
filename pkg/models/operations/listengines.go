package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type ListEnginesResponse struct {
	ContentType         string
	ListEnginesResponse *shared.ListEnginesResponse
	StatusCode          int
}
