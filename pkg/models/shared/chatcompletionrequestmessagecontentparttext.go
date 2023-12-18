// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequestMessageContentPartTextType - The type of the content part.
type ChatCompletionRequestMessageContentPartTextType string

const (
	ChatCompletionRequestMessageContentPartTextTypeText ChatCompletionRequestMessageContentPartTextType = "text"
)

func (e ChatCompletionRequestMessageContentPartTextType) ToPointer() *ChatCompletionRequestMessageContentPartTextType {
	return &e
}

func (e *ChatCompletionRequestMessageContentPartTextType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = ChatCompletionRequestMessageContentPartTextType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestMessageContentPartTextType: %v", v)
	}
}

type ChatCompletionRequestMessageContentPartText struct {
	// The text content.
	Text string `json:"text"`
	// The type of the content part.
	Type ChatCompletionRequestMessageContentPartTextType `json:"type"`
}

func (o *ChatCompletionRequestMessageContentPartText) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

func (o *ChatCompletionRequestMessageContentPartText) GetType() ChatCompletionRequestMessageContentPartTextType {
	if o == nil {
		return ChatCompletionRequestMessageContentPartTextType("")
	}
	return o.Type
}