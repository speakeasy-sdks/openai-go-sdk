// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

// CreateRunRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type CreateRunRequestMetadata struct {
}

type CreateRunRequestToolsType string

const (
	CreateRunRequestToolsTypeAssistantToolsCode      CreateRunRequestToolsType = "AssistantToolsCode"
	CreateRunRequestToolsTypeAssistantToolsRetrieval CreateRunRequestToolsType = "AssistantToolsRetrieval"
	CreateRunRequestToolsTypeAssistantToolsFunction  CreateRunRequestToolsType = "AssistantToolsFunction"
)

type CreateRunRequestTools struct {
	AssistantToolsCode      *AssistantToolsCode
	AssistantToolsRetrieval *AssistantToolsRetrieval
	AssistantToolsFunction  *AssistantToolsFunction

	Type CreateRunRequestToolsType
}

func CreateCreateRunRequestToolsAssistantToolsCode(assistantToolsCode AssistantToolsCode) CreateRunRequestTools {
	typ := CreateRunRequestToolsTypeAssistantToolsCode

	return CreateRunRequestTools{
		AssistantToolsCode: &assistantToolsCode,
		Type:               typ,
	}
}

func CreateCreateRunRequestToolsAssistantToolsRetrieval(assistantToolsRetrieval AssistantToolsRetrieval) CreateRunRequestTools {
	typ := CreateRunRequestToolsTypeAssistantToolsRetrieval

	return CreateRunRequestTools{
		AssistantToolsRetrieval: &assistantToolsRetrieval,
		Type:                    typ,
	}
}

func CreateCreateRunRequestToolsAssistantToolsFunction(assistantToolsFunction AssistantToolsFunction) CreateRunRequestTools {
	typ := CreateRunRequestToolsTypeAssistantToolsFunction

	return CreateRunRequestTools{
		AssistantToolsFunction: &assistantToolsFunction,
		Type:                   typ,
	}
}

func (u *CreateRunRequestTools) UnmarshalJSON(data []byte) error {

	assistantToolsCode := AssistantToolsCode{}
	if err := utils.UnmarshalJSON(data, &assistantToolsCode, "", true, true); err == nil {
		u.AssistantToolsCode = &assistantToolsCode
		u.Type = CreateRunRequestToolsTypeAssistantToolsCode
		return nil
	}

	assistantToolsRetrieval := AssistantToolsRetrieval{}
	if err := utils.UnmarshalJSON(data, &assistantToolsRetrieval, "", true, true); err == nil {
		u.AssistantToolsRetrieval = &assistantToolsRetrieval
		u.Type = CreateRunRequestToolsTypeAssistantToolsRetrieval
		return nil
	}

	assistantToolsFunction := AssistantToolsFunction{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFunction, "", true, true); err == nil {
		u.AssistantToolsFunction = &assistantToolsFunction
		u.Type = CreateRunRequestToolsTypeAssistantToolsFunction
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateRunRequestTools) MarshalJSON() ([]byte, error) {
	if u.AssistantToolsCode != nil {
		return utils.MarshalJSON(u.AssistantToolsCode, "", true)
	}

	if u.AssistantToolsRetrieval != nil {
		return utils.MarshalJSON(u.AssistantToolsRetrieval, "", true)
	}

	if u.AssistantToolsFunction != nil {
		return utils.MarshalJSON(u.AssistantToolsFunction, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateRunRequest struct {
	// The ID of the [assistant](/docs/api-reference/assistants) to use to execute this run.
	AssistantID string `json:"assistant_id"`
	// Override the default system message of the assistant. This is useful for modifying the behavior on a per-run basis.
	Instructions *string `json:"instructions,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *CreateRunRequestMetadata `json:"metadata,omitempty"`
	// The ID of the [Model](/docs/api-reference/models) to be used to execute this run. If a value is provided here, it will override the model associated with the assistant. If not, the model associated with the assistant will be used.
	Model *string `json:"model,omitempty"`
	// Override the tools the assistant can use for this run. This is useful for modifying the behavior on a per-run basis.
	Tools []CreateRunRequestTools `json:"tools,omitempty"`
}

func (o *CreateRunRequest) GetAssistantID() string {
	if o == nil {
		return ""
	}
	return o.AssistantID
}

func (o *CreateRunRequest) GetInstructions() *string {
	if o == nil {
		return nil
	}
	return o.Instructions
}

func (o *CreateRunRequest) GetMetadata() *CreateRunRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CreateRunRequest) GetModel() *string {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *CreateRunRequest) GetTools() []CreateRunRequestTools {
	if o == nil {
		return nil
	}
	return o.Tools
}