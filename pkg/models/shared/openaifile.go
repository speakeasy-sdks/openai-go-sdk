// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OpenAIFileStatusDetails struct {
}

// OpenAIFile - OK
type OpenAIFile struct {
	Bytes         int64                    `json:"bytes"`
	CreatedAt     int64                    `json:"created_at"`
	Filename      string                   `json:"filename"`
	ID            string                   `json:"id"`
	Object        string                   `json:"object"`
	Purpose       string                   `json:"purpose"`
	Status        *string                  `json:"status,omitempty"`
	StatusDetails *OpenAIFileStatusDetails `json:"status_details,omitempty"`
}
