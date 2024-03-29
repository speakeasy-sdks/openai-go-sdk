// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
)

type CreateImageRequest2 string

const (
	CreateImageRequest2DallE2 CreateImageRequest2 = "dall-e-2"
	CreateImageRequest2DallE3 CreateImageRequest2 = "dall-e-3"
)

func (e CreateImageRequest2) ToPointer() *CreateImageRequest2 {
	return &e
}

func (e *CreateImageRequest2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dall-e-2":
		fallthrough
	case "dall-e-3":
		*e = CreateImageRequest2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequest2: %v", v)
	}
}

type CreateImageRequestModelType string

const (
	CreateImageRequestModelTypeStr                 CreateImageRequestModelType = "str"
	CreateImageRequestModelTypeCreateImageRequest2 CreateImageRequestModelType = "CreateImageRequest_2"
)

// CreateImageRequestModel - The model to use for image generation.
type CreateImageRequestModel struct {
	Str                 *string
	CreateImageRequest2 *CreateImageRequest2

	Type CreateImageRequestModelType
}

func CreateCreateImageRequestModelStr(str string) CreateImageRequestModel {
	typ := CreateImageRequestModelTypeStr

	return CreateImageRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateImageRequestModelCreateImageRequest2(createImageRequest2 CreateImageRequest2) CreateImageRequestModel {
	typ := CreateImageRequestModelTypeCreateImageRequest2

	return CreateImageRequestModel{
		CreateImageRequest2: &createImageRequest2,
		Type:                typ,
	}
}

func (u *CreateImageRequestModel) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateImageRequestModelTypeStr
		return nil
	}

	createImageRequest2 := CreateImageRequest2("")
	if err := utils.UnmarshalJSON(data, &createImageRequest2, "", true, true); err == nil {
		u.CreateImageRequest2 = &createImageRequest2
		u.Type = CreateImageRequestModelTypeCreateImageRequest2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateImageRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateImageRequest2 != nil {
		return utils.MarshalJSON(u.CreateImageRequest2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Quality - The quality of the image that will be generated. `hd` creates images with finer details and greater consistency across the image. This param is only supported for `dall-e-3`.
type Quality string

const (
	QualityStandard Quality = "standard"
	QualityHd       Quality = "hd"
)

func (e Quality) ToPointer() *Quality {
	return &e
}

func (e *Quality) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "hd":
		*e = Quality(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Quality: %v", v)
	}
}

// CreateImageRequestResponseFormat - The format in which the generated images are returned. Must be one of `url` or `b64_json`.
type CreateImageRequestResponseFormat string

const (
	CreateImageRequestResponseFormatURL     CreateImageRequestResponseFormat = "url"
	CreateImageRequestResponseFormatB64JSON CreateImageRequestResponseFormat = "b64_json"
)

func (e CreateImageRequestResponseFormat) ToPointer() *CreateImageRequestResponseFormat {
	return &e
}

func (e *CreateImageRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequestResponseFormat: %v", v)
	}
}

// CreateImageRequestSize - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024` for `dall-e-2`. Must be one of `1024x1024`, `1792x1024`, or `1024x1792` for `dall-e-3` models.
type CreateImageRequestSize string

const (
	CreateImageRequestSizeTwoHundredAndFiftySixx256                CreateImageRequestSize = "256x256"
	CreateImageRequestSizeFiveHundredAndTwelvex512                 CreateImageRequestSize = "512x512"
	CreateImageRequestSizeOneThousandAndTwentyFourx1024            CreateImageRequestSize = "1024x1024"
	CreateImageRequestSizeOneThousandSevenHundredAndNinetyTwox1024 CreateImageRequestSize = "1792x1024"
	CreateImageRequestSizeOneThousandAndTwentyFourx1792            CreateImageRequestSize = "1024x1792"
)

func (e CreateImageRequestSize) ToPointer() *CreateImageRequestSize {
	return &e
}

func (e *CreateImageRequestSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "256x256":
		fallthrough
	case "512x512":
		fallthrough
	case "1024x1024":
		fallthrough
	case "1792x1024":
		fallthrough
	case "1024x1792":
		*e = CreateImageRequestSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequestSize: %v", v)
	}
}

// Style - The style of the generated images. Must be one of `vivid` or `natural`. Vivid causes the model to lean towards generating hyper-real and dramatic images. Natural causes the model to produce more natural, less hyper-real looking images. This param is only supported for `dall-e-3`.
type Style string

const (
	StyleVivid   Style = "vivid"
	StyleNatural Style = "natural"
)

func (e Style) ToPointer() *Style {
	return &e
}

func (e *Style) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vivid":
		fallthrough
	case "natural":
		*e = Style(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Style: %v", v)
	}
}

type CreateImageRequest struct {
	// The model to use for image generation.
	Model *CreateImageRequestModel `json:"model,omitempty"`
	// The number of images to generate. Must be between 1 and 10. For `dall-e-3`, only `n=1` is supported.
	N *int64 `default:"1" json:"n"`
	// A text description of the desired image(s). The maximum length is 1000 characters for `dall-e-2` and 4000 characters for `dall-e-3`.
	Prompt string `json:"prompt"`
	// The quality of the image that will be generated. `hd` creates images with finer details and greater consistency across the image. This param is only supported for `dall-e-3`.
	Quality *Quality `default:"standard" json:"quality"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`.
	ResponseFormat *CreateImageRequestResponseFormat `default:"url" json:"response_format"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024` for `dall-e-2`. Must be one of `1024x1024`, `1792x1024`, or `1024x1792` for `dall-e-3` models.
	Size *CreateImageRequestSize `default:"1024x1024" json:"size"`
	// The style of the generated images. Must be one of `vivid` or `natural`. Vivid causes the model to lean towards generating hyper-real and dramatic images. Natural causes the model to produce more natural, less hyper-real looking images. This param is only supported for `dall-e-3`.
	Style *Style `default:"vivid" json:"style"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `json:"user,omitempty"`
}

func (c CreateImageRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateImageRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateImageRequest) GetModel() *CreateImageRequestModel {
	if o == nil {
		return nil
	}
	return o.Model
}

func (o *CreateImageRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateImageRequest) GetPrompt() string {
	if o == nil {
		return ""
	}
	return o.Prompt
}

func (o *CreateImageRequest) GetQuality() *Quality {
	if o == nil {
		return nil
	}
	return o.Quality
}

func (o *CreateImageRequest) GetResponseFormat() *CreateImageRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateImageRequest) GetSize() *CreateImageRequestSize {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateImageRequest) GetStyle() *Style {
	if o == nil {
		return nil
	}
	return o.Style
}

func (o *CreateImageRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
