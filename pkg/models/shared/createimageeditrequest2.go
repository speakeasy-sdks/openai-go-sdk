// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/utils"
)

type CreateImageEditRequestImage struct {
	Content []byte `multipartForm:"content"`
	Image   string `multipartForm:"name=image"`
}

func (o *CreateImageEditRequestImage) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CreateImageEditRequestImage) GetImage() string {
	if o == nil {
		return ""
	}
	return o.Image
}

type CreateImageEditRequestMask struct {
	Content []byte `multipartForm:"content"`
	Mask    string `multipartForm:"name=mask"`
}

func (o *CreateImageEditRequestMask) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CreateImageEditRequestMask) GetMask() string {
	if o == nil {
		return ""
	}
	return o.Mask
}

// CreateImageEditRequestResponseFormat - The format in which the generated images are returned. Must be one of `url` or `b64_json`.
type CreateImageEditRequestResponseFormat string

const (
	CreateImageEditRequestResponseFormatURL     CreateImageEditRequestResponseFormat = "url"
	CreateImageEditRequestResponseFormatB64JSON CreateImageEditRequestResponseFormat = "b64_json"
)

func (e CreateImageEditRequestResponseFormat) ToPointer() *CreateImageEditRequestResponseFormat {
	return &e
}

func (e *CreateImageEditRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageEditRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageEditRequestResponseFormat: %v", v)
	}
}

// CreateImageEditRequestSize - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
type CreateImageEditRequestSize string

const (
	CreateImageEditRequestSizeTwoHundredAndFiftySixx256     CreateImageEditRequestSize = "256x256"
	CreateImageEditRequestSizeFiveHundredAndTwelvex512      CreateImageEditRequestSize = "512x512"
	CreateImageEditRequestSizeOneThousandAndTwentyFourx1024 CreateImageEditRequestSize = "1024x1024"
)

func (e CreateImageEditRequestSize) ToPointer() *CreateImageEditRequestSize {
	return &e
}

func (e *CreateImageEditRequestSize) UnmarshalJSON(data []byte) error {
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
		*e = CreateImageEditRequestSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageEditRequestSize: %v", v)
	}
}

type CreateImageEditRequest2 struct {
	// The image to edit. Must be a valid PNG file, less than 4MB, and square. If mask is not provided, image must have transparency, which will be used as the mask.
	Image CreateImageEditRequestImage `multipartForm:"file"`
	// An additional image whose fully transparent areas (e.g. where alpha is zero) indicate where `image` should be edited. Must be a valid PNG file, less than 4MB, and have the same dimensions as `image`.
	Mask *CreateImageEditRequestMask `multipartForm:"file"`
	// The number of images to generate. Must be between 1 and 10.
	N *int64 `default:"1" multipartForm:"name=n"`
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `multipartForm:"name=prompt"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`.
	ResponseFormat *CreateImageEditRequestResponseFormat `default:"url" multipartForm:"name=response_format"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size *CreateImageEditRequestSize `default:"1024x1024" multipartForm:"name=size"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `multipartForm:"name=user"`
}

func (c CreateImageEditRequest2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateImageEditRequest2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateImageEditRequest2) GetImage() CreateImageEditRequestImage {
	if o == nil {
		return CreateImageEditRequestImage{}
	}
	return o.Image
}

func (o *CreateImageEditRequest2) GetMask() *CreateImageEditRequestMask {
	if o == nil {
		return nil
	}
	return o.Mask
}

func (o *CreateImageEditRequest2) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateImageEditRequest2) GetPrompt() string {
	if o == nil {
		return ""
	}
	return o.Prompt
}

func (o *CreateImageEditRequest2) GetResponseFormat() *CreateImageEditRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateImageEditRequest2) GetSize() *CreateImageEditRequestSize {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateImageEditRequest2) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
