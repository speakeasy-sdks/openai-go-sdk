// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateMessageRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type CreateMessageRequestMetadata struct {
}

// CreateMessageRequestRole - The role of the entity that is creating the message. Currently only `user` is supported.
type CreateMessageRequestRole string

const (
	CreateMessageRequestRoleUser CreateMessageRequestRole = "user"
)

func (e CreateMessageRequestRole) ToPointer() *CreateMessageRequestRole {
	return &e
}

func (e *CreateMessageRequestRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = CreateMessageRequestRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateMessageRequestRole: %v", v)
	}
}

type CreateMessageRequest struct {
	// The content of the message.
	Content string `json:"content"`
	// A list of [File](/docs/api-reference/files) IDs that the message should use. There can be a maximum of 10 files attached to a message. Useful for tools like `retrieval` and `code_interpreter` that can access and use files.
	FileIds []string `json:"file_ids,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *CreateMessageRequestMetadata `json:"metadata,omitempty"`
	// The role of the entity that is creating the message. Currently only `user` is supported.
	Role CreateMessageRequestRole `json:"role"`
}

func (o *CreateMessageRequest) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *CreateMessageRequest) GetFileIds() []string {
	if o == nil {
		return nil
	}
	return o.FileIds
}

func (o *CreateMessageRequest) GetMetadata() *CreateMessageRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CreateMessageRequest) GetRole() CreateMessageRequestRole {
	if o == nil {
		return CreateMessageRequestRole("")
	}
	return o.Role
}