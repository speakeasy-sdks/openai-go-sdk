// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// OpenAIFile - OK
type OpenAIFile struct {
	Bytes         int64   `json:"bytes"`
	CreatedAt     int64   `json:"created_at"`
	Filename      string  `json:"filename"`
	ID            string  `json:"id"`
	Object        string  `json:"object"`
	Purpose       string  `json:"purpose"`
	Status        *string `json:"status,omitempty"`
	StatusDetails *string `json:"status_details,omitempty"`
}

func (o *OpenAIFile) GetBytes() int64 {
	if o == nil {
		return 0
	}
	return o.Bytes
}

func (o *OpenAIFile) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *OpenAIFile) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

func (o *OpenAIFile) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OpenAIFile) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}

func (o *OpenAIFile) GetPurpose() string {
	if o == nil {
		return ""
	}
	return o.Purpose
}

func (o *OpenAIFile) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *OpenAIFile) GetStatusDetails() *string {
	if o == nil {
		return nil
	}
	return o.StatusDetails
}
