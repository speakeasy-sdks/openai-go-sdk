// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type GetRunRequest struct {
	// The ID of the run to retrieve.
	RunID string `pathParam:"style=simple,explode=false,name=run_id"`
	// The ID of the [thread](/docs/api-reference/threads) that was run.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
}

func (o *GetRunRequest) GetRunID() string {
	if o == nil {
		return ""
	}
	return o.RunID
}

func (o *GetRunRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

type GetRunResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	RunObject *shared.RunObject
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRunResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRunResponse) GetRunObject() *shared.RunObject {
	if o == nil {
		return nil
	}
	return o.RunObject
}

func (o *GetRunResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRunResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
