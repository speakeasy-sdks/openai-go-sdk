// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

// Detail - Specifies the detail level of the image. Learn more in the [Vision guide](/docs/guides/vision/low-or-high-fidelity-image-understanding).
type Detail string

const (
	DetailAuto Detail = "auto"
	DetailLow  Detail = "low"
	DetailHigh Detail = "high"
)

func (e Detail) ToPointer() *Detail {
	return &e
}

func (e *Detail) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "low":
		fallthrough
	case "high":
		*e = Detail(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Detail: %v", v)
	}
}

type ImageURL struct {
	// Specifies the detail level of the image. Learn more in the [Vision guide](/docs/guides/vision/low-or-high-fidelity-image-understanding).
	Detail *Detail `default:"auto" json:"detail"`
	// Either a URL of the image or the base64 encoded image data.
	URL string `json:"url"`
}

func (i ImageURL) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *ImageURL) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ImageURL) GetDetail() *Detail {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *ImageURL) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

// SchemasChatCompletionRequestMessageContentPartImageType - The type of the content part.
type SchemasChatCompletionRequestMessageContentPartImageType string

const (
	SchemasChatCompletionRequestMessageContentPartImageTypeImageURL SchemasChatCompletionRequestMessageContentPartImageType = "image_url"
)

func (e SchemasChatCompletionRequestMessageContentPartImageType) ToPointer() *SchemasChatCompletionRequestMessageContentPartImageType {
	return &e
}

func (e *SchemasChatCompletionRequestMessageContentPartImageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "image_url":
		*e = SchemasChatCompletionRequestMessageContentPartImageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasChatCompletionRequestMessageContentPartImageType: %v", v)
	}
}

type ImageContentPart struct {
	ImageURL ImageURL `json:"image_url"`
	// The type of the content part.
	Type SchemasChatCompletionRequestMessageContentPartImageType `json:"type"`
}

func (o *ImageContentPart) GetImageURL() ImageURL {
	if o == nil {
		return ImageURL{}
	}
	return o.ImageURL
}

func (o *ImageContentPart) GetType() SchemasChatCompletionRequestMessageContentPartImageType {
	if o == nil {
		return SchemasChatCompletionRequestMessageContentPartImageType("")
	}
	return o.Type
}

// SchemasChatCompletionRequestMessageContentPartTextType - The type of the content part.
type SchemasChatCompletionRequestMessageContentPartTextType string

const (
	SchemasChatCompletionRequestMessageContentPartTextTypeText SchemasChatCompletionRequestMessageContentPartTextType = "text"
)

func (e SchemasChatCompletionRequestMessageContentPartTextType) ToPointer() *SchemasChatCompletionRequestMessageContentPartTextType {
	return &e
}

func (e *SchemasChatCompletionRequestMessageContentPartTextType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = SchemasChatCompletionRequestMessageContentPartTextType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasChatCompletionRequestMessageContentPartTextType: %v", v)
	}
}

type TextContentPart struct {
	// The text content.
	Text string `json:"text"`
	// The type of the content part.
	Type SchemasChatCompletionRequestMessageContentPartTextType `json:"type"`
}

func (o *TextContentPart) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

func (o *TextContentPart) GetType() SchemasChatCompletionRequestMessageContentPartTextType {
	if o == nil {
		return SchemasChatCompletionRequestMessageContentPartTextType("")
	}
	return o.Type
}

type ChatCompletionRequestMessageContentPartType string

const (
	ChatCompletionRequestMessageContentPartTypeTextContentPart  ChatCompletionRequestMessageContentPartType = "Text content part"
	ChatCompletionRequestMessageContentPartTypeImageContentPart ChatCompletionRequestMessageContentPartType = "Image content part"
)

type ChatCompletionRequestMessageContentPart struct {
	TextContentPart  *TextContentPart
	ImageContentPart *ImageContentPart

	Type ChatCompletionRequestMessageContentPartType
}

func CreateChatCompletionRequestMessageContentPartTextContentPart(textContentPart TextContentPart) ChatCompletionRequestMessageContentPart {
	typ := ChatCompletionRequestMessageContentPartTypeTextContentPart

	return ChatCompletionRequestMessageContentPart{
		TextContentPart: &textContentPart,
		Type:            typ,
	}
}

func CreateChatCompletionRequestMessageContentPartImageContentPart(imageContentPart ImageContentPart) ChatCompletionRequestMessageContentPart {
	typ := ChatCompletionRequestMessageContentPartTypeImageContentPart

	return ChatCompletionRequestMessageContentPart{
		ImageContentPart: &imageContentPart,
		Type:             typ,
	}
}

func (u *ChatCompletionRequestMessageContentPart) UnmarshalJSON(data []byte) error {

	textContentPart := TextContentPart{}
	if err := utils.UnmarshalJSON(data, &textContentPart, "", true, true); err == nil {
		u.TextContentPart = &textContentPart
		u.Type = ChatCompletionRequestMessageContentPartTypeTextContentPart
		return nil
	}

	imageContentPart := ImageContentPart{}
	if err := utils.UnmarshalJSON(data, &imageContentPart, "", true, true); err == nil {
		u.ImageContentPart = &imageContentPart
		u.Type = ChatCompletionRequestMessageContentPartTypeImageContentPart
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ChatCompletionRequestMessageContentPart) MarshalJSON() ([]byte, error) {
	if u.TextContentPart != nil {
		return utils.MarshalJSON(u.TextContentPart, "", true)
	}

	if u.ImageContentPart != nil {
		return utils.MarshalJSON(u.ImageContentPart, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
