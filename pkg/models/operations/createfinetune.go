// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type CreateFineTuneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	FineTune *shared.FineTune
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateFineTuneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFineTuneResponse) GetFineTune() *shared.FineTune {
	if o == nil {
		return nil
	}
	return o.FineTune
}

func (o *CreateFineTuneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFineTuneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
