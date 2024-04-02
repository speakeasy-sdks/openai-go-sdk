// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
	"net/http"
)

type CreateTranscriptionResponseBodyType string

const (
	CreateTranscriptionResponseBodyTypeCreateTranscriptionResponseJSON        CreateTranscriptionResponseBodyType = "CreateTranscriptionResponseJson"
	CreateTranscriptionResponseBodyTypeCreateTranscriptionResponseVerboseJSON CreateTranscriptionResponseBodyType = "CreateTranscriptionResponseVerboseJson"
)

// CreateTranscriptionResponseBody - OK
type CreateTranscriptionResponseBody struct {
	CreateTranscriptionResponseJSON        *shared.CreateTranscriptionResponseJSON
	CreateTranscriptionResponseVerboseJSON *shared.CreateTranscriptionResponseVerboseJSON

	Type CreateTranscriptionResponseBodyType
}

func CreateCreateTranscriptionResponseBodyCreateTranscriptionResponseJSON(createTranscriptionResponseJSON shared.CreateTranscriptionResponseJSON) CreateTranscriptionResponseBody {
	typ := CreateTranscriptionResponseBodyTypeCreateTranscriptionResponseJSON

	return CreateTranscriptionResponseBody{
		CreateTranscriptionResponseJSON: &createTranscriptionResponseJSON,
		Type:                            typ,
	}
}

func CreateCreateTranscriptionResponseBodyCreateTranscriptionResponseVerboseJSON(createTranscriptionResponseVerboseJSON shared.CreateTranscriptionResponseVerboseJSON) CreateTranscriptionResponseBody {
	typ := CreateTranscriptionResponseBodyTypeCreateTranscriptionResponseVerboseJSON

	return CreateTranscriptionResponseBody{
		CreateTranscriptionResponseVerboseJSON: &createTranscriptionResponseVerboseJSON,
		Type:                                   typ,
	}
}

func (u *CreateTranscriptionResponseBody) UnmarshalJSON(data []byte) error {

	createTranscriptionResponseJSON := shared.CreateTranscriptionResponseJSON{}
	if err := utils.UnmarshalJSON(data, &createTranscriptionResponseJSON, "", true, true); err == nil {
		u.CreateTranscriptionResponseJSON = &createTranscriptionResponseJSON
		u.Type = CreateTranscriptionResponseBodyTypeCreateTranscriptionResponseJSON
		return nil
	}

	createTranscriptionResponseVerboseJSON := shared.CreateTranscriptionResponseVerboseJSON{}
	if err := utils.UnmarshalJSON(data, &createTranscriptionResponseVerboseJSON, "", true, true); err == nil {
		u.CreateTranscriptionResponseVerboseJSON = &createTranscriptionResponseVerboseJSON
		u.Type = CreateTranscriptionResponseBodyTypeCreateTranscriptionResponseVerboseJSON
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateTranscriptionResponseBody) MarshalJSON() ([]byte, error) {
	if u.CreateTranscriptionResponseJSON != nil {
		return utils.MarshalJSON(u.CreateTranscriptionResponseJSON, "", true)
	}

	if u.CreateTranscriptionResponseVerboseJSON != nil {
		return utils.MarshalJSON(u.CreateTranscriptionResponseVerboseJSON, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateTranscriptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	OneOf *CreateTranscriptionResponseBody
}

func (o *CreateTranscriptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTranscriptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTranscriptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTranscriptionResponse) GetOneOf() *CreateTranscriptionResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
