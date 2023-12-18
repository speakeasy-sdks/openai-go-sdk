// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

// Code - One of `server_error` or `rate_limit_exceeded`.
type Code string

const (
	CodeServerError       Code = "server_error"
	CodeRateLimitExceeded Code = "rate_limit_exceeded"
)

func (e Code) ToPointer() *Code {
	return &e
}

func (e *Code) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "server_error":
		fallthrough
	case "rate_limit_exceeded":
		*e = Code(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Code: %v", v)
	}
}

// LastError - The last error associated with this run. Will be `null` if there are no errors.
type LastError struct {
	// One of `server_error` or `rate_limit_exceeded`.
	Code Code `json:"code"`
	// A human-readable description of the error.
	Message string `json:"message"`
}

func (o *LastError) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}

func (o *LastError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

// RunObjectMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type RunObjectMetadata struct {
}

// RunObjectObject - The object type, which is always `thread.run`.
type RunObjectObject string

const (
	RunObjectObjectThreadRun RunObjectObject = "thread.run"
)

func (e RunObjectObject) ToPointer() *RunObjectObject {
	return &e
}

func (e *RunObjectObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "thread.run":
		*e = RunObjectObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunObjectObject: %v", v)
	}
}

// SubmitToolOutputs - Details on the tool outputs needed for this run to continue.
type SubmitToolOutputs struct {
	// A list of the relevant tool calls.
	ToolCalls []RunToolCallObject `json:"tool_calls"`
}

func (o *SubmitToolOutputs) GetToolCalls() []RunToolCallObject {
	if o == nil {
		return []RunToolCallObject{}
	}
	return o.ToolCalls
}

// RunObjectType - For now, this is always `submit_tool_outputs`.
type RunObjectType string

const (
	RunObjectTypeSubmitToolOutputs RunObjectType = "submit_tool_outputs"
)

func (e RunObjectType) ToPointer() *RunObjectType {
	return &e
}

func (e *RunObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "submit_tool_outputs":
		*e = RunObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunObjectType: %v", v)
	}
}

// RequiredAction - Details on the action required to continue the run. Will be `null` if no action is required.
type RequiredAction struct {
	// Details on the tool outputs needed for this run to continue.
	SubmitToolOutputs SubmitToolOutputs `json:"submit_tool_outputs"`
	// For now, this is always `submit_tool_outputs`.
	Type RunObjectType `json:"type"`
}

func (o *RequiredAction) GetSubmitToolOutputs() SubmitToolOutputs {
	if o == nil {
		return SubmitToolOutputs{}
	}
	return o.SubmitToolOutputs
}

func (o *RequiredAction) GetType() RunObjectType {
	if o == nil {
		return RunObjectType("")
	}
	return o.Type
}

// RunObjectStatus - The status of the run, which can be either `queued`, `in_progress`, `requires_action`, `cancelling`, `cancelled`, `failed`, `completed`, or `expired`.
type RunObjectStatus string

const (
	RunObjectStatusQueued         RunObjectStatus = "queued"
	RunObjectStatusInProgress     RunObjectStatus = "in_progress"
	RunObjectStatusRequiresAction RunObjectStatus = "requires_action"
	RunObjectStatusCancelling     RunObjectStatus = "cancelling"
	RunObjectStatusCancelled      RunObjectStatus = "cancelled"
	RunObjectStatusFailed         RunObjectStatus = "failed"
	RunObjectStatusCompleted      RunObjectStatus = "completed"
	RunObjectStatusExpired        RunObjectStatus = "expired"
)

func (e RunObjectStatus) ToPointer() *RunObjectStatus {
	return &e
}

func (e *RunObjectStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "queued":
		fallthrough
	case "in_progress":
		fallthrough
	case "requires_action":
		fallthrough
	case "cancelling":
		fallthrough
	case "cancelled":
		fallthrough
	case "failed":
		fallthrough
	case "completed":
		fallthrough
	case "expired":
		*e = RunObjectStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunObjectStatus: %v", v)
	}
}

type RunObjectToolsType string

const (
	RunObjectToolsTypeAssistantToolsCode      RunObjectToolsType = "AssistantToolsCode"
	RunObjectToolsTypeAssistantToolsRetrieval RunObjectToolsType = "AssistantToolsRetrieval"
	RunObjectToolsTypeAssistantToolsFunction  RunObjectToolsType = "AssistantToolsFunction"
)

type RunObjectTools struct {
	AssistantToolsCode      *AssistantToolsCode
	AssistantToolsRetrieval *AssistantToolsRetrieval
	AssistantToolsFunction  *AssistantToolsFunction

	Type RunObjectToolsType
}

func CreateRunObjectToolsAssistantToolsCode(assistantToolsCode AssistantToolsCode) RunObjectTools {
	typ := RunObjectToolsTypeAssistantToolsCode

	return RunObjectTools{
		AssistantToolsCode: &assistantToolsCode,
		Type:               typ,
	}
}

func CreateRunObjectToolsAssistantToolsRetrieval(assistantToolsRetrieval AssistantToolsRetrieval) RunObjectTools {
	typ := RunObjectToolsTypeAssistantToolsRetrieval

	return RunObjectTools{
		AssistantToolsRetrieval: &assistantToolsRetrieval,
		Type:                    typ,
	}
}

func CreateRunObjectToolsAssistantToolsFunction(assistantToolsFunction AssistantToolsFunction) RunObjectTools {
	typ := RunObjectToolsTypeAssistantToolsFunction

	return RunObjectTools{
		AssistantToolsFunction: &assistantToolsFunction,
		Type:                   typ,
	}
}

func (u *RunObjectTools) UnmarshalJSON(data []byte) error {

	assistantToolsCode := AssistantToolsCode{}
	if err := utils.UnmarshalJSON(data, &assistantToolsCode, "", true, true); err == nil {
		u.AssistantToolsCode = &assistantToolsCode
		u.Type = RunObjectToolsTypeAssistantToolsCode
		return nil
	}

	assistantToolsRetrieval := AssistantToolsRetrieval{}
	if err := utils.UnmarshalJSON(data, &assistantToolsRetrieval, "", true, true); err == nil {
		u.AssistantToolsRetrieval = &assistantToolsRetrieval
		u.Type = RunObjectToolsTypeAssistantToolsRetrieval
		return nil
	}

	assistantToolsFunction := AssistantToolsFunction{}
	if err := utils.UnmarshalJSON(data, &assistantToolsFunction, "", true, true); err == nil {
		u.AssistantToolsFunction = &assistantToolsFunction
		u.Type = RunObjectToolsTypeAssistantToolsFunction
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RunObjectTools) MarshalJSON() ([]byte, error) {
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

// RunObject - Represents an execution run on a [thread](/docs/api-reference/threads).
type RunObject struct {
	// The ID of the [assistant](/docs/api-reference/assistants) used for execution of this run.
	AssistantID string `json:"assistant_id"`
	// The Unix timestamp (in seconds) for when the run was cancelled.
	CancelledAt *int64 `json:"cancelled_at"`
	// The Unix timestamp (in seconds) for when the run was completed.
	CompletedAt *int64 `json:"completed_at"`
	// The Unix timestamp (in seconds) for when the run was created.
	CreatedAt int64 `json:"created_at"`
	// The Unix timestamp (in seconds) for when the run will expire.
	ExpiresAt int64 `json:"expires_at"`
	// The Unix timestamp (in seconds) for when the run failed.
	FailedAt *int64 `json:"failed_at"`
	// The list of [File](/docs/api-reference/files) IDs the [assistant](/docs/api-reference/assistants) used for this run.
	FileIds []string `json:"file_ids"`
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// The instructions that the [assistant](/docs/api-reference/assistants) used for this run.
	Instructions string `json:"instructions"`
	// The last error associated with this run. Will be `null` if there are no errors.
	LastError *LastError `json:"last_error"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *RunObjectMetadata `json:"metadata"`
	// The model that the [assistant](/docs/api-reference/assistants) used for this run.
	Model string `json:"model"`
	// The object type, which is always `thread.run`.
	Object RunObjectObject `json:"object"`
	// Details on the action required to continue the run. Will be `null` if no action is required.
	RequiredAction *RequiredAction `json:"required_action"`
	// The Unix timestamp (in seconds) for when the run was started.
	StartedAt *int64 `json:"started_at"`
	// The status of the run, which can be either `queued`, `in_progress`, `requires_action`, `cancelling`, `cancelled`, `failed`, `completed`, or `expired`.
	Status RunObjectStatus `json:"status"`
	// The ID of the [thread](/docs/api-reference/threads) that was executed on as a part of this run.
	ThreadID string `json:"thread_id"`
	// The list of tools that the [assistant](/docs/api-reference/assistants) used for this run.
	Tools []RunObjectTools `json:"tools"`
}

func (o *RunObject) GetAssistantID() string {
	if o == nil {
		return ""
	}
	return o.AssistantID
}

func (o *RunObject) GetCancelledAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CancelledAt
}

func (o *RunObject) GetCompletedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *RunObject) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *RunObject) GetExpiresAt() int64 {
	if o == nil {
		return 0
	}
	return o.ExpiresAt
}

func (o *RunObject) GetFailedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.FailedAt
}

func (o *RunObject) GetFileIds() []string {
	if o == nil {
		return []string{}
	}
	return o.FileIds
}

func (o *RunObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RunObject) GetInstructions() string {
	if o == nil {
		return ""
	}
	return o.Instructions
}

func (o *RunObject) GetLastError() *LastError {
	if o == nil {
		return nil
	}
	return o.LastError
}

func (o *RunObject) GetMetadata() *RunObjectMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *RunObject) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *RunObject) GetObject() RunObjectObject {
	if o == nil {
		return RunObjectObject("")
	}
	return o.Object
}

func (o *RunObject) GetRequiredAction() *RequiredAction {
	if o == nil {
		return nil
	}
	return o.RequiredAction
}

func (o *RunObject) GetStartedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *RunObject) GetStatus() RunObjectStatus {
	if o == nil {
		return RunObjectStatus("")
	}
	return o.Status
}

func (o *RunObject) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

func (o *RunObject) GetTools() []RunObjectTools {
	if o == nil {
		return []RunObjectTools{}
	}
	return o.Tools
}