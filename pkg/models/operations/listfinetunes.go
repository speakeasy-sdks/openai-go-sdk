// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
	"net/http"
)

type ListFineTunesResponse struct {
	ContentType string
	// OK
	ListFineTunesResponse *shared.ListFineTunesResponse
	StatusCode            int
	RawResponse           *http.Response
}
