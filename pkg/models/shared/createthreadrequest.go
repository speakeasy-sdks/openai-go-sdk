// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateThreadRequestMetadata - Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
type CreateThreadRequestMetadata struct {
}

type CreateThreadRequest struct {
	// A list of [messages](/docs/api-reference/messages) to start the thread with.
	Messages []CreateMessageRequest `json:"messages,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
	//
	Metadata *CreateThreadRequestMetadata `json:"metadata,omitempty"`
}

func (o *CreateThreadRequest) GetMessages() []CreateMessageRequest {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *CreateThreadRequest) GetMetadata() *CreateThreadRequestMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}
