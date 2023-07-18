// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Model - OK
type Model struct {
	Created int64  `json:"created"`
	ID      string `json:"id"`
	Object  string `json:"object"`
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
