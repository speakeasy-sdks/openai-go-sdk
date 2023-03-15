package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateEditResponse struct {
	ContentType        string
	CreateEditResponse *shared.CreateEditResponse
	StatusCode         int
	RawResponse        *http.Response
}
