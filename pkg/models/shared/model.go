// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Model - Describes an OpenAI model offering that can be used with the API.
type Model struct {
	// The Unix timestamp (in seconds) when the model was created.
	Created int64 `json:"created"`
	// The model identifier, which can be referenced in the API endpoints.
	ID string `json:"id"`
	// The object type, which is always "model".
	Object string `json:"object"`
	// The organization that owns the model.
	OwnedBy string `json:"owned_by"`
}

func (o *Model) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *Model) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Model) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}

func (o *Model) GetOwnedBy() string {
	if o == nil {
		return ""
	}
	return o.OwnedBy
}
