package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateModerationRequest struct {
	Request shared.CreateModerationRequest `request:"mediaType=application/json"`
}

type CreateModerationResponse struct {
	ContentType              string
	CreateModerationResponse *shared.CreateModerationResponse
	StatusCode               int
}
