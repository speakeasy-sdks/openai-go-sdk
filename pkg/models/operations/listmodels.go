// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type ListModelsResponse struct {
	ContentType string
	// OK
	ListModelsResponse *shared.ListModelsResponse
	StatusCode         int
	RawResponse        *http.Response
}
