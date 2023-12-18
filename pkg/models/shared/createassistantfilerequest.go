// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateAssistantFileRequest struct {
	// A [File](/docs/api-reference/files) ID (with `purpose="assistants"`) that the assistant should use. Useful for tools like `retrieval` and `code_interpreter` that can access files.
	FileID string `json:"file_id"`
}

func (o *CreateAssistantFileRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}