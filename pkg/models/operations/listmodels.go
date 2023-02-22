package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type ListModelsResponse struct {
	ContentType        string
	ListModelsResponse *shared.ListModelsResponse
	StatusCode         int
}
