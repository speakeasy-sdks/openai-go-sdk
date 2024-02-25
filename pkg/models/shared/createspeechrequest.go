// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
)

type CreateSpeechRequest2 string

const (
	CreateSpeechRequest2Tts1   CreateSpeechRequest2 = "tts-1"
	CreateSpeechRequest2Tts1Hd CreateSpeechRequest2 = "tts-1-hd"
)

func (e CreateSpeechRequest2) ToPointer() *CreateSpeechRequest2 {
	return &e
}

func (e *CreateSpeechRequest2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tts-1":
		fallthrough
	case "tts-1-hd":
		*e = CreateSpeechRequest2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSpeechRequest2: %v", v)
	}
}

type CreateSpeechRequestModelType string

const (
	CreateSpeechRequestModelTypeStr                  CreateSpeechRequestModelType = "str"
	CreateSpeechRequestModelTypeCreateSpeechRequest2 CreateSpeechRequestModelType = "CreateSpeechRequest_2"
)

// CreateSpeechRequestModel - One of the available [TTS models](/docs/models/tts): `tts-1` or `tts-1-hd`
type CreateSpeechRequestModel struct {
	Str                  *string
	CreateSpeechRequest2 *CreateSpeechRequest2

	Type CreateSpeechRequestModelType
}

func CreateCreateSpeechRequestModelStr(str string) CreateSpeechRequestModel {
	typ := CreateSpeechRequestModelTypeStr

	return CreateSpeechRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateSpeechRequestModelCreateSpeechRequest2(createSpeechRequest2 CreateSpeechRequest2) CreateSpeechRequestModel {
	typ := CreateSpeechRequestModelTypeCreateSpeechRequest2

	return CreateSpeechRequestModel{
		CreateSpeechRequest2: &createSpeechRequest2,
		Type:                 typ,
	}
}

func (u *CreateSpeechRequestModel) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateSpeechRequestModelTypeStr
		return nil
	}

	createSpeechRequest2 := CreateSpeechRequest2("")
	if err := utils.UnmarshalJSON(data, &createSpeechRequest2, "", true, true); err == nil {
		u.CreateSpeechRequest2 = &createSpeechRequest2
		u.Type = CreateSpeechRequestModelTypeCreateSpeechRequest2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateSpeechRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateSpeechRequest2 != nil {
		return utils.MarshalJSON(u.CreateSpeechRequest2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// CreateSpeechRequestResponseFormat - The format to return audio in.
// Supported formats are `mp3`, `opus`, `aac`, `flac`, `pcm`, and `wav`.
//
// The `pcm` audio format, similar to `wav` but without a header, utilizes a 24kHz sample rate, mono channel, and 16-bit depth in signed little-endian format.
type CreateSpeechRequestResponseFormat string

const (
	CreateSpeechRequestResponseFormatMp3  CreateSpeechRequestResponseFormat = "mp3"
	CreateSpeechRequestResponseFormatOpus CreateSpeechRequestResponseFormat = "opus"
	CreateSpeechRequestResponseFormatAac  CreateSpeechRequestResponseFormat = "aac"
	CreateSpeechRequestResponseFormatFlac CreateSpeechRequestResponseFormat = "flac"
	CreateSpeechRequestResponseFormatPcm  CreateSpeechRequestResponseFormat = "pcm"
	CreateSpeechRequestResponseFormatWav  CreateSpeechRequestResponseFormat = "wav"
)

func (e CreateSpeechRequestResponseFormat) ToPointer() *CreateSpeechRequestResponseFormat {
	return &e
}

func (e *CreateSpeechRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mp3":
		fallthrough
	case "opus":
		fallthrough
	case "aac":
		fallthrough
	case "flac":
		fallthrough
	case "pcm":
		fallthrough
	case "wav":
		*e = CreateSpeechRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateSpeechRequestResponseFormat: %v", v)
	}
}

// Voice - The voice to use when generating the audio. Supported voices are `alloy`, `echo`, `fable`, `onyx`, `nova`, and `shimmer`. Previews of the voices are available in the [Text to speech guide](/docs/guides/text-to-speech/voice-options).
type Voice string

const (
	VoiceAlloy   Voice = "alloy"
	VoiceEcho    Voice = "echo"
	VoiceFable   Voice = "fable"
	VoiceOnyx    Voice = "onyx"
	VoiceNova    Voice = "nova"
	VoiceShimmer Voice = "shimmer"
)

func (e Voice) ToPointer() *Voice {
	return &e
}

func (e *Voice) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "alloy":
		fallthrough
	case "echo":
		fallthrough
	case "fable":
		fallthrough
	case "onyx":
		fallthrough
	case "nova":
		fallthrough
	case "shimmer":
		*e = Voice(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Voice: %v", v)
	}
}

type CreateSpeechRequest struct {
	// The text to generate audio for. The maximum length is 4096 characters.
	Input string `json:"input"`
	// One of the available [TTS models](/docs/models/tts): `tts-1` or `tts-1-hd`
	//
	Model CreateSpeechRequestModel `json:"model"`
	// The format to return audio in.
	// Supported formats are `mp3`, `opus`, `aac`, `flac`, `pcm`, and `wav`.
	//
	// The `pcm` audio format, similar to `wav` but without a header, utilizes a 24kHz sample rate, mono channel, and 16-bit depth in signed little-endian format.
	ResponseFormat *CreateSpeechRequestResponseFormat `default:"mp3" json:"response_format"`
	// The speed of the generated audio. Select a value from `0.25` to `4.0`. `1.0` is the default.
	Speed *float64 `default:"1" json:"speed"`
	// The voice to use when generating the audio. Supported voices are `alloy`, `echo`, `fable`, `onyx`, `nova`, and `shimmer`. Previews of the voices are available in the [Text to speech guide](/docs/guides/text-to-speech/voice-options).
	Voice Voice `json:"voice"`
}

func (c CreateSpeechRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateSpeechRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateSpeechRequest) GetInput() string {
	if o == nil {
		return ""
	}
	return o.Input
}

func (o *CreateSpeechRequest) GetModel() CreateSpeechRequestModel {
	if o == nil {
		return CreateSpeechRequestModel{}
	}
	return o.Model
}

func (o *CreateSpeechRequest) GetResponseFormat() *CreateSpeechRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateSpeechRequest) GetSpeed() *float64 {
	if o == nil {
		return nil
	}
	return o.Speed
}

func (o *CreateSpeechRequest) GetVoice() Voice {
	if o == nil {
		return Voice("")
	}
	return o.Voice
}
