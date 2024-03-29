// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"net/http"
)

type CreateAssistantFileRequest struct {
	CreateAssistantFileRequest shared.CreateAssistantFileRequest `request:"mediaType=application/json"`
	// The ID of the assistant for which to create a File.
	//
	AssistantID string `pathParam:"style=simple,explode=false,name=assistant_id"`
}

func (o *CreateAssistantFileRequest) GetCreateAssistantFileRequest() shared.CreateAssistantFileRequest {
	if o == nil {
		return shared.CreateAssistantFileRequest{}
	}
	return o.CreateAssistantFileRequest
}

func (o *CreateAssistantFileRequest) GetAssistantID() string {
	if o == nil {
		return ""
	}
	return o.AssistantID
}

type CreateAssistantFileResponse struct {
	// OK
	AssistantFileObject *shared.AssistantFileObject
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateAssistantFileResponse) GetAssistantFileObject() *shared.AssistantFileObject {
	if o == nil {
		return nil
	}
	return o.AssistantFileObject
}

func (o *CreateAssistantFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAssistantFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAssistantFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
