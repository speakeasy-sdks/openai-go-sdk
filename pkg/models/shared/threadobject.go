// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ThreadObjectMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type ThreadObjectMetadata struct {
}

// ThreadObjectObject - The object type, which is always `thread`.
type ThreadObjectObject string

const (
	ThreadObjectObjectThread ThreadObjectObject = "thread"
)

func (e ThreadObjectObject) ToPointer() *ThreadObjectObject {
	return &e
}

func (e *ThreadObjectObject) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "thread":
		*e = ThreadObjectObject(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ThreadObjectObject: %v", v)
	}
}

// ThreadObject - Represents a thread that contains [messages](/docs/api-reference/messages).
type ThreadObject struct {
	// The Unix timestamp (in seconds) for when the thread was created.
	CreatedAt int64 `json:"created_at"`
	// The identifier, which can be referenced in API endpoints.
	ID string `json:"id"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *ThreadObjectMetadata `json:"metadata"`
	// The object type, which is always `thread`.
	Object ThreadObjectObject `json:"object"`
}

func (o *ThreadObject) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *ThreadObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ThreadObject) GetMetadata() *ThreadObjectMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ThreadObject) GetObject() ThreadObjectObject {
	if o == nil {
		return ThreadObjectObject("")
	}
	return o.Object
}