// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

type CreateTranslationRequestFile struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *CreateTranslationRequestFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CreateTranslationRequestFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// CreateTranslationRequest2 - ID of the model to use. Only `whisper-1` is currently available.
type CreateTranslationRequest2 string

const (
	CreateTranslationRequest2Whisper1 CreateTranslationRequest2 = "whisper-1"
)

func (e CreateTranslationRequest2) ToPointer() *CreateTranslationRequest2 {
	return &e
}

func (e *CreateTranslationRequest2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "whisper-1":
		*e = CreateTranslationRequest2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTranslationRequest2: %v", v)
	}
}

type CreateTranslationRequestModelType string

const (
	CreateTranslationRequestModelTypeStr                       CreateTranslationRequestModelType = "str"
	CreateTranslationRequestModelTypeCreateTranslationRequest2 CreateTranslationRequestModelType = "CreateTranslationRequest_2"
)

type CreateTranslationRequestModel struct {
	Str                       *string
	CreateTranslationRequest2 *CreateTranslationRequest2

	Type CreateTranslationRequestModelType
}

func CreateCreateTranslationRequestModelStr(str string) CreateTranslationRequestModel {
	typ := CreateTranslationRequestModelTypeStr

	return CreateTranslationRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateTranslationRequestModelCreateTranslationRequest2(createTranslationRequest2 CreateTranslationRequest2) CreateTranslationRequestModel {
	typ := CreateTranslationRequestModelTypeCreateTranslationRequest2

	return CreateTranslationRequestModel{
		CreateTranslationRequest2: &createTranslationRequest2,
		Type:                      typ,
	}
}

func (u *CreateTranslationRequestModel) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateTranslationRequestModelTypeStr
		return nil
	}

	createTranslationRequest2 := CreateTranslationRequest2("")
	if err := utils.UnmarshalJSON(data, &createTranslationRequest2, "", true, true); err == nil {
		u.CreateTranslationRequest2 = &createTranslationRequest2
		u.Type = CreateTranslationRequestModelTypeCreateTranslationRequest2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateTranslationRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateTranslationRequest2 != nil {
		return utils.MarshalJSON(u.CreateTranslationRequest2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateTranslationRequest struct {
	// The audio file object (not file name) translate, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm.
	//
	File CreateTranslationRequestFile `multipartForm:"file"`
	// ID of the model to use. Only `whisper-1` is currently available.
	//
	Model CreateTranslationRequestModel `multipartForm:"name=model"`
	// An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text/prompting) should be in English.
	//
	Prompt *string `multipartForm:"name=prompt"`
	// The format of the transcript output, in one of these options: `json`, `text`, `srt`, `verbose_json`, or `vtt`.
	//
	ResponseFormat *string `default:"json" multipartForm:"name=response_format"`
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.
	//
	Temperature *float64 `default:"0" multipartForm:"name=temperature"`
}

func (c CreateTranslationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTranslationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTranslationRequest) GetFile() CreateTranslationRequestFile {
	if o == nil {
		return CreateTranslationRequestFile{}
	}
	return o.File
}

func (o *CreateTranslationRequest) GetModel() CreateTranslationRequestModel {
	if o == nil {
		return CreateTranslationRequestModel{}
	}
	return o.Model
}

func (o *CreateTranslationRequest) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *CreateTranslationRequest) GetResponseFormat() *string {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateTranslationRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}
