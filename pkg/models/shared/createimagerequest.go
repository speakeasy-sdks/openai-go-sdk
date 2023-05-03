// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateImageRequestResponseFormatEnum - The format in which the generated images are returned. Must be one of `url` or `b64_json`.
type CreateImageRequestResponseFormatEnum string

const (
	CreateImageRequestResponseFormatEnumURL     CreateImageRequestResponseFormatEnum = "url"
	CreateImageRequestResponseFormatEnumB64JSON CreateImageRequestResponseFormatEnum = "b64_json"
)

func (e CreateImageRequestResponseFormatEnum) ToPointer() *CreateImageRequestResponseFormatEnum {
	return &e
}

func (e *CreateImageRequestResponseFormatEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageRequestResponseFormatEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequestResponseFormatEnum: %v", v)
	}
}

// CreateImageRequestSizeEnum - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
type CreateImageRequestSizeEnum string

const (
	CreateImageRequestSizeEnumTwoHundredAndFiftySixx256     CreateImageRequestSizeEnum = "256x256"
	CreateImageRequestSizeEnumFiveHundredAndTwelvex512      CreateImageRequestSizeEnum = "512x512"
	CreateImageRequestSizeEnumOneThousandAndTwentyFourx1024 CreateImageRequestSizeEnum = "1024x1024"
)

func (e CreateImageRequestSizeEnum) ToPointer() *CreateImageRequestSizeEnum {
	return &e
}

func (e *CreateImageRequestSizeEnum) UnmarshalJSON(data []byte) error {
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
		*e = CreateImageRequestSizeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequestSizeEnum: %v", v)
	}
}

type CreateImageRequest struct {
	// The number of images to generate. Must be between 1 and 10.
	N *int64 `json:"n,omitempty"`
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`.
	ResponseFormat *CreateImageRequestResponseFormatEnum `json:"response_format,omitempty"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size *CreateImageRequestSizeEnum `json:"size,omitempty"`
	User interface{}                 `json:"user,omitempty"`
}
