// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type GetThreadRequest struct {
	// The ID of the thread to retrieve.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
}

func (o *GetThreadRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

type GetThreadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ThreadObject *shared.ThreadObject
}

func (o *GetThreadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetThreadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetThreadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetThreadResponse) GetThreadObject() *shared.ThreadObject {
	if o == nil {
		return nil
	}
	return o.ThreadObject
}