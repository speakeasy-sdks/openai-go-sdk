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

// Function - The function definition.
type Function struct {
	// A description of what the function does, used by the model to choose when and how to call the function.
	Description string `json:"description"`
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.
	Name string `json:"name"`
	// The parameters the functions accepts, described as a JSON Schema object. See the [guide](/docs/guides/gpt/function-calling) for examples, and the [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for documentation about the format.
	//
	// To describe a function that accepts no parameters, provide the value `{"type": "object", "properties": {}}`.
	Parameters map[string]interface{} `json:"parameters"`
}

func (o *Function) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Function) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Function) GetParameters() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Parameters
}

// SchemasType - The type of tool being defined: `function`
type SchemasType string

const (
	SchemasTypeFunction SchemasType = "function"
)

func (e SchemasType) ToPointer() *SchemasType {
	return &e
}

func (e *SchemasType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "function":
		*e = SchemasType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasType: %v", v)
	}
}

type FunctionTool struct {
	// The function definition.
	Function Function `json:"function"`
	// The type of tool being defined: `function`
	Type SchemasType `json:"type"`
}

func (o *FunctionTool) GetFunction() Function {
	if o == nil {
		return Function{}
	}
	return o.Function
}

func (o *FunctionTool) GetType() SchemasType {
	if o == nil {
		return SchemasType("")
	}
	return o.Type
}

// SchemasAssistantToolsRetrievalType - The type of tool being defined: `retrieval`
type SchemasAssistantToolsRetrievalType string

const (
	SchemasAssistantToolsRetrievalTypeRetrieval SchemasAssistantToolsRetrievalType = "retrieval"
)

func (e SchemasAssistantToolsRetrievalType) ToPointer() *SchemasAssistantToolsRetrievalType {
	return &e
}

func (e *SchemasAssistantToolsRetrievalType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "retrieval":
		*e = SchemasAssistantToolsRetrievalType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasAssistantToolsRetrievalType: %v", v)
	}
}

type RetrievalTool struct {
	// The type of tool being defined: `retrieval`
	Type SchemasAssistantToolsRetrievalType `json:"type"`
}

func (o *RetrievalTool) GetType() SchemasAssistantToolsRetrievalType {
	if o == nil {
		return SchemasAssistantToolsRetrievalType("")
	}
	return o.Type
}

// Type - The type of tool being defined: `code_interpreter`
type Type string

const (
	TypeCodeInterpreter Type = "code_interpreter"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "code_interpreter":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type CodeInterpreterTool struct {
	// The type of tool being defined: `code_interpreter`
	Type Type `json:"type"`
}

func (o *CodeInterpreterTool) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

type ToolsType string

const (
	ToolsTypeCodeInterpreterTool ToolsType = "Code interpreter tool"
	ToolsTypeRetrievalTool       ToolsType = "Retrieval tool"
	ToolsTypeFunctionTool        ToolsType = "Function tool"
)

type Tools struct {
	CodeInterpreterTool *CodeInterpreterTool
	RetrievalTool       *RetrievalTool
	FunctionTool        *FunctionTool

	Type ToolsType
}

func CreateToolsCodeInterpreterTool(codeInterpreterTool CodeInterpreterTool) Tools {
	typ := ToolsTypeCodeInterpreterTool

	return Tools{
		CodeInterpreterTool: &codeInterpreterTool,
		Type:                typ,
	}
}

func CreateToolsRetrievalTool(retrievalTool RetrievalTool) Tools {
	typ := ToolsTypeRetrievalTool

	return Tools{
		RetrievalTool: &retrievalTool,
		Type:          typ,
	}
}

func CreateToolsFunctionTool(functionTool FunctionTool) Tools {
	typ := ToolsTypeFunctionTool

	return Tools{
		FunctionTool: &functionTool,
		Type:         typ,
	}
}

func (u *Tools) UnmarshalJSON(data []byte) error {

	codeInterpreterTool := CodeInterpreterTool{}
	if err := utils.UnmarshalJSON(data, &codeInterpreterTool, "", true, true); err == nil {
		u.CodeInterpreterTool = &codeInterpreterTool
		u.Type = ToolsTypeCodeInterpreterTool
		return nil
	}

	retrievalTool := RetrievalTool{}
	if err := utils.UnmarshalJSON(data, &retrievalTool, "", true, true); err == nil {
		u.RetrievalTool = &retrievalTool
		u.Type = ToolsTypeRetrievalTool
		return nil
	}

	functionTool := FunctionTool{}
	if err := utils.UnmarshalJSON(data, &functionTool, "", true, true); err == nil {
		u.FunctionTool = &functionTool
		u.Type = ToolsTypeFunctionTool
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Tools) MarshalJSON() ([]byte, error) {
	if u.CodeInterpreterTool != nil {
		return utils.MarshalJSON(u.CodeInterpreterTool, "", true)
	}

	if u.RetrievalTool != nil {
		return utils.MarshalJSON(u.RetrievalTool, "", true)
	}

	if u.FunctionTool != nil {
		return utils.MarshalJSON(u.FunctionTool, "", true)
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
