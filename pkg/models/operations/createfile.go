// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateFileResponse struct {
	ContentType string
	// OK
	OpenAIFile  *shared.OpenAIFile
	StatusCode  int
	RawResponse *http.Response
}
