package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

type CreateEditRequest struct {
	Request shared.CreateEditRequest `request:"mediaType=application/json"`
}

type CreateEditResponse struct {
	ContentType        string
	CreateEditResponse *shared.CreateEditResponse
	StatusCode         int
}
