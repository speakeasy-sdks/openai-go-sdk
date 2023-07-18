// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DeleteFileResponse - OK
type DeleteFileResponse struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}

func (o *DeleteFileResponse) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *DeleteFileResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteFileResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}
