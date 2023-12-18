// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"net/http"
)

type GetAssistantFileRequest struct {
	// The ID of the assistant who the file belongs to.
	AssistantID string `pathParam:"style=simple,explode=false,name=assistant_id"`
	// The ID of the file we're getting.
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

func (o *GetAssistantFileRequest) GetAssistantID() string {
	if o == nil {
		return ""
	}
	return o.AssistantID
}

func (o *GetAssistantFileRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

type GetAssistantFileResponse struct {
	// OK
	AssistantFileObject *shared.AssistantFileObject
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAssistantFileResponse) GetAssistantFileObject() *shared.AssistantFileObject {
	if o == nil {
		return nil
	}
	return o.AssistantFileObject
}

func (o *GetAssistantFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAssistantFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAssistantFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}