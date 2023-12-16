// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

// Metadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type Metadata struct {
}

// AssistantObjectObject - The object type, which is always `assistant`.
type AssistantObjectObject string

const (
	AssistantObjectObjectAssistant AssistantObjectObject = "assistant"
)

func (e AssistantObjectObject) ToPointer() *AssistantObjectObject {
	return &e
}

func (e *AssistantObjectObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "assistant":
		*e = AssistantObjectObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AssistantObjectObject: %v", v)
	}
}

type ToolsType string

const (
	ToolsTypeAssistantToolsCode      ToolsType = "AssistantToolsCode"
	ToolsTypeAssistantToolsRetrieval ToolsType = "AssistantToolsRetrieval"
	ToolsTypeAssistantToolsFunction  ToolsType = "AssistantToolsFunction"
)

type Tools struct {
	AssistantToolsCode      *AssistantToolsCode
	AssistantToolsRetrieval *AssistantToolsRetrieval
	AssistantToolsFunction  *AssistantToolsFunction

	Type ToolsType
}

func CreateToolsAssistantToolsCode(assistantToolsCode AssistantToolsCode) Tools {
	typ := ToolsTypeAssistantToolsCode

	return Tools{
		AssistantToolsCode: &assistantToolsCode,
		Type:               typ,
	}
}

func CreateToolsAssistantToolsRetrieval(assistantToolsRetrieval AssistantToolsRetrieval) Tools {
	typ := ToolsTypeAssistantToolsRetrieval

	return Tools{
		AssistantToolsRetrieval: &assistantToolsRetrieval,
		Type:                    typ,
	}
}

func CreateToolsAssistantToolsFunction(assistantToolsFunction AssistantToolsFunction) Tools {
	typ := ToolsTypeAssistantToolsFunction

	return Tools{
		AssistantToolsFunction: &assistantToolsFunction,
		Type:                   typ,
	}
}

func (u *Tools) UnmarshalJSON(data []byte) error {

	assistantToolsCode := AssistantToolsCode{}
	if err := utils.UnmarshalJSON(data, &assistantToolsCode, "", true, true); err == nil {
		u.AssistantToolsCode = &assistantToolsCode
		u.Type = ToolsTypeAssistantToolsCode
		return nil
	}

	assistantToolsRetrieval := AssistantToolsRetrieval{}
	if err := utils.UnmarshalJSON(data, &assistantToolsRetrieval, "", true, true); err == nil {
		u.AssistantToolsRetrieval = &assistantToolsRetrieval
		u.Type = ToolsTypeAssistantToolsRetrieval
		return nil
	}

	assistantToolsFunction := AssistantToolsFunction{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFunction, "", true, true); err == nil {
		u.AssistantToolsFunction = &assistantToolsFunction
		u.Type = ToolsTypeAssistantToolsFunction
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Tools) MarshalJSON() ([]byte, error) {
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

// AssistantObject - Represents an `assistant` that can call the model and use tools.
type AssistantObject struct {
	// The Unix timestamp (in seconds) for when the assistant was created.
	CreatedAt int64 `json:"created_at"`
	// The description of the assistant. The maximum length is 512 characters.
	//
	Description *string `json:"description"`
	// A list of [file](/docs/api-reference/files) IDs attached to this assistant. There can be a maximum of 20 files attached to the assistant. Files are ordered by their creation date in ascending order.
	//
	FileIds []string `json:"file_ids"`
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// The system instructions that the assistant uses. The maximum length is 32768 characters.
	//
	Instructions *string `json:"instructions"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *Metadata `json:"metadata"`
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
	//
	Model string `json:"model"`
	// The name of the assistant. The maximum length is 256 characters.
	//
	Name *string `json:"name"`
	// The object type, which is always `assistant`.
	Object AssistantObjectObject `json:"object"`
	// A list of tool enabled on the assistant. There can be a maximum of 128 tools per assistant. Tools can be of types `code_interpreter`, `retrieval`, or `function`.
	//
	Tools []Tools `json:"tools"`
}

func (o *AssistantObject) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *AssistantObject) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AssistantObject) GetFileIds() []string {
	if o == nil {
		return []string{}
	}
	return o.FileIds
}

func (o *AssistantObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AssistantObject) GetInstructions() *string {
	if o == nil {
		return nil
	}
	return o.Instructions
}

func (o *AssistantObject) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AssistantObject) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *AssistantObject) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AssistantObject) GetObject() AssistantObjectObject {
	if o == nil {
		return AssistantObjectObject("")
	}
	return o.Object
}

func (o *AssistantObject) GetTools() []Tools {
	if o == nil {
		return []Tools{}
	}
	return o.Tools
}
