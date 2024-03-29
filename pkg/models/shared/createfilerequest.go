// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type File struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *File) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// Purpose - The intended purpose of the uploaded file.
//
// Use "fine-tune" for [Fine-tuning](/docs/api-reference/fine-tuning) and "assistants" for [Assistants](/docs/api-reference/assistants) and [Messages](/docs/api-reference/messages). This allows us to validate the format of the uploaded file is correct for fine-tuning.
type Purpose string

const (
	PurposeFineTune   Purpose = "fine-tune"
	PurposeAssistants Purpose = "assistants"
)

func (e Purpose) ToPointer() *Purpose {
	return &e
}

func (e *Purpose) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fine-tune":
		fallthrough
	case "assistants":
		*e = Purpose(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Purpose: %v", v)
	}
}

type CreateFileRequest struct {
	// The File object (not file name) to be uploaded.
	//
	File File `multipartForm:"file"`
	// The intended purpose of the uploaded file.
	//
	// Use "fine-tune" for [Fine-tuning](/docs/api-reference/fine-tuning) and "assistants" for [Assistants](/docs/api-reference/assistants) and [Messages](/docs/api-reference/messages). This allows us to validate the format of the uploaded file is correct for fine-tuning.
	//
	Purpose Purpose `multipartForm:"name=purpose"`
}

func (o *CreateFileRequest) GetFile() File {
	if o == nil {
		return File{}
	}
	return o.File
}

func (o *CreateFileRequest) GetPurpose() Purpose {
	if o == nil {
		return Purpose("")
	}
	return o.Purpose
}
