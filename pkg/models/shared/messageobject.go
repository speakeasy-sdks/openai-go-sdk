// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
)

type MessageObjectContentType string

const (
	MessageObjectContentTypeMessageContentImageFileObject MessageObjectContentType = "MessageContentImageFileObject"
	MessageObjectContentTypeMessageContentTextObject      MessageObjectContentType = "MessageContentTextObject"
)

type MessageObjectContent struct {
	MessageContentImageFileObject *MessageContentImageFileObject
	MessageContentTextObject      *MessageContentTextObject

	Type MessageObjectContentType
}

func CreateMessageObjectContentMessageContentImageFileObject(messageContentImageFileObject MessageContentImageFileObject) MessageObjectContent {
	typ := MessageObjectContentTypeMessageContentImageFileObject

	return MessageObjectContent{
		MessageContentImageFileObject: &messageContentImageFileObject,
		Type:                          typ,
	}
}

func CreateMessageObjectContentMessageContentTextObject(messageContentTextObject MessageContentTextObject) MessageObjectContent {
	typ := MessageObjectContentTypeMessageContentTextObject

	return MessageObjectContent{
		MessageContentTextObject: &messageContentTextObject,
		Type:                     typ,
	}
}

func (u *MessageObjectContent) UnmarshalJSON(data []byte) error {

	messageContentImageFileObject := MessageContentImageFileObject{}
	if err := utils.UnmarshalJSON(data, &messageContentImageFileObject, "", true, true); err == nil {
		u.MessageContentImageFileObject = &messageContentImageFileObject
		u.Type = MessageObjectContentTypeMessageContentImageFileObject
		return nil
	}

	messageContentTextObject := MessageContentTextObject{}
	if err := utils.UnmarshalJSON(data, &messageContentTextObject, "", true, true); err == nil {
		u.MessageContentTextObject = &messageContentTextObject
		u.Type = MessageObjectContentTypeMessageContentTextObject
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u MessageObjectContent) MarshalJSON() ([]byte, error) {
	if u.MessageContentImageFileObject != nil {
		return utils.MarshalJSON(u.MessageContentImageFileObject, "", true)
	}

	if u.MessageContentTextObject != nil {
		return utils.MarshalJSON(u.MessageContentTextObject, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Reason - The reason the message is incomplete.
type Reason string

const (
	ReasonContentFilter Reason = "content_filter"
	ReasonMaxTokens     Reason = "max_tokens"
	ReasonRunCancelled  Reason = "run_cancelled"
	ReasonRunExpired    Reason = "run_expired"
	ReasonRunFailed     Reason = "run_failed"
)

func (e Reason) ToPointer() *Reason {
	return &e
}

func (e *Reason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "content_filter":
		fallthrough
	case "max_tokens":
		fallthrough
	case "run_cancelled":
		fallthrough
	case "run_expired":
		fallthrough
	case "run_failed":
		*e = Reason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Reason: %v", v)
	}
}

// IncompleteDetails - On an incomplete message, details about why the message is incomplete.
type IncompleteDetails struct {
	// The reason the message is incomplete.
	Reason Reason `json:"reason"`
}

func (o *IncompleteDetails) GetReason() Reason {
	if o == nil {
		return Reason("")
	}
	return o.Reason
}

// MessageObjectMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type MessageObjectMetadata struct {
}

// MessageObjectObject - The object type, which is always `thread.message`.
type MessageObjectObject string

const (
	MessageObjectObjectThreadMessage MessageObjectObject = "thread.message"
)

func (e MessageObjectObject) ToPointer() *MessageObjectObject {
	return &e
}

func (e *MessageObjectObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "thread.message":
		*e = MessageObjectObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageObjectObject: %v", v)
	}
}

// MessageObjectRole - The entity that produced the message. One of `user` or `assistant`.
type MessageObjectRole string

const (
	MessageObjectRoleUser      MessageObjectRole = "user"
	MessageObjectRoleAssistant MessageObjectRole = "assistant"
)

func (e MessageObjectRole) ToPointer() *MessageObjectRole {
	return &e
}

func (e *MessageObjectRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		fallthrough
	case "assistant":
		*e = MessageObjectRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageObjectRole: %v", v)
	}
}

// MessageObjectStatus - The status of the message, which can be either `in_progress`, `incomplete`, or `completed`.
type MessageObjectStatus string

const (
	MessageObjectStatusInProgress MessageObjectStatus = "in_progress"
	MessageObjectStatusIncomplete MessageObjectStatus = "incomplete"
	MessageObjectStatusCompleted  MessageObjectStatus = "completed"
)

func (e MessageObjectStatus) ToPointer() *MessageObjectStatus {
	return &e
}

func (e *MessageObjectStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "in_progress":
		fallthrough
	case "incomplete":
		fallthrough
	case "completed":
		*e = MessageObjectStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageObjectStatus: %v", v)
	}
}

// MessageObject - Represents a message within a [thread](/docs/api-reference/threads).
type MessageObject struct {
	// If applicable, the ID of the [assistant](/docs/api-reference/assistants) that authored this message.
	AssistantID *string `json:"assistant_id"`
	// The Unix timestamp (in seconds) for when the message was completed.
	CompletedAt *int64 `json:"completed_at"`
	// The content of the message in array of text and/or images.
	Content []MessageObjectContent `json:"content"`
	// The Unix timestamp (in seconds) for when the message was created.
	CreatedAt int64 `json:"created_at"`
	// A list of [file](/docs/api-reference/files) IDs that the assistant should use. Useful for tools like retrieval and code_interpreter that can access files. A maximum of 10 files can be attached to a message.
	FileIds []string `json:"file_ids"`
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// The Unix timestamp (in seconds) for when the message was marked as incomplete.
	IncompleteAt *int64 `json:"incomplete_at"`
	// On an incomplete message, details about why the message is incomplete.
	IncompleteDetails *IncompleteDetails `json:"incomplete_details"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *MessageObjectMetadata `json:"metadata"`
	// The object type, which is always `thread.message`.
	Object MessageObjectObject `json:"object"`
	// The entity that produced the message. One of `user` or `assistant`.
	Role MessageObjectRole `json:"role"`
	// The ID of the [run](/docs/api-reference/runs) associated with the creation of this message. Value is `null` when messages are created manually using the create message or create thread endpoints.
	RunID *string `json:"run_id"`
	// The status of the message, which can be either `in_progress`, `incomplete`, or `completed`.
	Status MessageObjectStatus `json:"status"`
	// The [thread](/docs/api-reference/threads) ID that this message belongs to.
	ThreadID string `json:"thread_id"`
}

func (o *MessageObject) GetAssistantID() *string {
	if o == nil {
		return nil
	}
	return o.AssistantID
}

func (o *MessageObject) GetCompletedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *MessageObject) GetContent() []MessageObjectContent {
	if o == nil {
		return []MessageObjectContent{}
	}
	return o.Content
}

func (o *MessageObject) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *MessageObject) GetFileIds() []string {
	if o == nil {
		return []string{}
	}
	return o.FileIds
}

func (o *MessageObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MessageObject) GetIncompleteAt() *int64 {
	if o == nil {
		return nil
	}
	return o.IncompleteAt
}

func (o *MessageObject) GetIncompleteDetails() *IncompleteDetails {
	if o == nil {
		return nil
	}
	return o.IncompleteDetails
}

func (o *MessageObject) GetMetadata() *MessageObjectMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *MessageObject) GetObject() MessageObjectObject {
	if o == nil {
		return MessageObjectObject("")
	}
	return o.Object
}

func (o *MessageObject) GetRole() MessageObjectRole {
	if o == nil {
		return MessageObjectRole("")
	}
	return o.Role
}

func (o *MessageObject) GetRunID() *string {
	if o == nil {
		return nil
	}
	return o.RunID
}

func (o *MessageObject) GetStatus() MessageObjectStatus {
	if o == nil {
		return MessageObjectStatus("")
	}
	return o.Status
}

func (o *MessageObject) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}
