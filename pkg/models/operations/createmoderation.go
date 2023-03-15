package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateModerationResponse struct {
	ContentType              string
	CreateModerationResponse *shared.CreateModerationResponse
	StatusCode               int
	RawResponse              *http.Response
}
