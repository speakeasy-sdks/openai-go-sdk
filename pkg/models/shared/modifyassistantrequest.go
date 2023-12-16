// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

// ModifyAssistantRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type ModifyAssistantRequestMetadata struct {
}

type ModifyAssistantRequestToolsType string

const (
	ModifyAssistantRequestToolsTypeAssistantToolsCode      ModifyAssistantRequestToolsType = "AssistantToolsCode"
	ModifyAssistantRequestToolsTypeAssistantToolsRetrieval ModifyAssistantRequestToolsType = "AssistantToolsRetrieval"
	ModifyAssistantRequestToolsTypeAssistantToolsFunction  ModifyAssistantRequestToolsType = "AssistantToolsFunction"
)

type ModifyAssistantRequestTools struct {
	AssistantToolsCode      *AssistantToolsCode
	AssistantToolsRetrieval *AssistantToolsRetrieval
	AssistantToolsFunction  *AssistantToolsFunction

	Type ModifyAssistantRequestToolsType
}

func CreateModifyAssistantRequestToolsAssistantToolsCode(assistantToolsCode AssistantToolsCode) ModifyAssistantRequestTools {
	typ := ModifyAssistantRequestToolsTypeAssistantToolsCode

	return ModifyAssistantRequestTools{
		AssistantToolsCode: &assistantToolsCode,
		Type:               typ,
	}
}

func CreateModifyAssistantRequestToolsAssistantToolsRetrieval(assistantToolsRetrieval AssistantToolsRetrieval) ModifyAssistantRequestTools {
	typ := ModifyAssistantRequestToolsTypeAssistantToolsRetrieval

	return ModifyAssistantRequestTools{
		AssistantToolsRetrieval: &assistantToolsRetrieval,
		Type:                    typ,
	}
}

func CreateModifyAssistantRequestToolsAssistantToolsFunction(assistantToolsFunction AssistantToolsFunction) ModifyAssistantRequestTools {
	typ := ModifyAssistantRequestToolsTypeAssistantToolsFunction

	return ModifyAssistantRequestTools{
		AssistantToolsFunction: &assistantToolsFunction,
		Type:                   typ,
	}
}

func (u *ModifyAssistantRequestTools) UnmarshalJSON(data []byte) error {

	assistantToolsCode := AssistantToolsCode{}
	if err := utils.UnmarshalJSON(data, &assistantToolsCode, "", true, true); err == nil {
		u.AssistantToolsCode = &assistantToolsCode
		u.Type = ModifyAssistantRequestToolsTypeAssistantToolsCode
		return nil
	}

	assistantToolsRetrieval := AssistantToolsRetrieval{}
	if err := utils.UnmarshalJSON(data, &assistantToolsRetrieval, "", true, true); err == nil {
		u.AssistantToolsRetrieval = &assistantToolsRetrieval
		u.Type = ModifyAssistantRequestToolsTypeAssistantToolsRetrieval
		return nil
	}

	assistantToolsFunction := AssistantToolsFunction{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFunction, "", true, true); err == nil {
		u.AssistantToolsFunction = &assistantToolsFunction
		u.Type = ModifyAssistantRequestToolsTypeAssistantToolsFunction
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ModifyAssistantRequestTools) MarshalJSON() ([]byte, error) {
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

type ModifyAssistantRequest struct {
	// The description of the assistant. The maximum length is 512 characters.
	//
	Description *string `json:"description,omitempty"`
	// A list of [File](/docs/api-reference/files) IDs attached to this assistant. There can be a maximum of 20 files attached to the assistant. Files are ordered by their creation date in ascending order. If a file was previosuly attached to the list but does not show up in the list, it will be deleted from the assistant.
	//
	FileIds []string `json:"file_ids,omitempty"`
	// The system instructions that the assistant uses. The maximum length is 32768 characters.
	//
	Instructions *string `json:"instructions,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *ModifyAssistantRequestMetadata `json:"metadata,omitempty"`
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
	//
	Model *string `json:"model,omitempty"`
	// The name of the assistant. The maximum length is 256 characters.
	//
	Name *string `json:"name,omitempty"`
	// A list of tool enabled on the assistant. There can be a maximum of 128 tools per assistant. Tools can be of types `code_interpreter`, `retrieval`, or `function`.
	//
	Tools []ModifyAssistantRequestTools `json:"tools,omitempty"`
}

func (o *ModifyAssistantRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ModifyAssistantRequest) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

func (o *ModifyAssistantRequest) GetInstructions() *string {
	if o == nil {
		return nil
	}
	return o.Instructions
}

func (o *ModifyAssistantRequest) GetMetadata() *ModifyAssistantRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ModifyAssistantRequest) GetModel() *string {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *ModifyAssistantRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ModifyAssistantRequest) GetTools() []ModifyAssistantRequestTools {
	if o == nil {
		return nil
	}
	return o.Tools
}
