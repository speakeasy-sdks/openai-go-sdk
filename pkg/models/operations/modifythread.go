// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"net/http"
)

type ModifyThreadRequest struct {
	ModifyThreadRequest shared.ModifyThreadRequest `request:"mediaType=application/json"`
	// The ID of the thread to modify. Only the `metadata` can be modified.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
}

func (o *ModifyThreadRequest) GetModifyThreadRequest() shared.ModifyThreadRequest {
	if o == nil {
		return shared.ModifyThreadRequest{}
	}
	return o.ModifyThreadRequest
}

func (o *ModifyThreadRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

type ModifyThreadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ThreadObject *shared.ThreadObject
}

func (o *ModifyThreadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ModifyThreadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ModifyThreadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ModifyThreadResponse) GetThreadObject() *shared.ThreadObject {
	if o == nil {
		return nil
	}
	return o.ThreadObject
}
