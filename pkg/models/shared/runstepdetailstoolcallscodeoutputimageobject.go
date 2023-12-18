// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type RunStepDetailsToolCallsCodeOutputImageObjectImage struct {
	// The [file](/docs/api-reference/files) ID of the image.
	FileID string `json:"file_id"`
}

func (o *RunStepDetailsToolCallsCodeOutputImageObjectImage) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

// RunStepDetailsToolCallsCodeOutputImageObjectType - Always `image`.
type RunStepDetailsToolCallsCodeOutputImageObjectType string

const (
	RunStepDetailsToolCallsCodeOutputImageObjectTypeImage RunStepDetailsToolCallsCodeOutputImageObjectType = "image"
)

func (e RunStepDetailsToolCallsCodeOutputImageObjectType) ToPointer() *RunStepDetailsToolCallsCodeOutputImageObjectType {
	return &e
}

func (e *RunStepDetailsToolCallsCodeOutputImageObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "image":
		*e = RunStepDetailsToolCallsCodeOutputImageObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunStepDetailsToolCallsCodeOutputImageObjectType: %v", v)
	}
}

type RunStepDetailsToolCallsCodeOutputImageObject struct {
	Image RunStepDetailsToolCallsCodeOutputImageObjectImage `json:"image"`
	// Always `image`.
	Type RunStepDetailsToolCallsCodeOutputImageObjectType `json:"type"`
}

func (o *RunStepDetailsToolCallsCodeOutputImageObject) GetImage() RunStepDetailsToolCallsCodeOutputImageObjectImage {
	if o == nil {
		return RunStepDetailsToolCallsCodeOutputImageObjectImage{}
	}
	return o.Image
}

func (o *RunStepDetailsToolCallsCodeOutputImageObject) GetType() RunStepDetailsToolCallsCodeOutputImageObjectType {
	if o == nil {
		return RunStepDetailsToolCallsCodeOutputImageObjectType("")
	}
	return o.Type
}