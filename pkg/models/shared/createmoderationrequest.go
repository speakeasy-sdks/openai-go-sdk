// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
)

type CreateModerationRequestInputType string

const (
	CreateModerationRequestInputTypeStr        CreateModerationRequestInputType = "str"
	CreateModerationRequestInputTypeArrayOfstr CreateModerationRequestInputType = "arrayOfstr"
)

// CreateModerationRequestInput - The input text to classify
type CreateModerationRequestInput struct {
	Str        *string
	ArrayOfstr []string

	Type CreateModerationRequestInputType
}

func CreateCreateModerationRequestInputStr(str string) CreateModerationRequestInput {
	typ := CreateModerationRequestInputTypeStr

	return CreateModerationRequestInput{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateModerationRequestInputArrayOfstr(arrayOfstr []string) CreateModerationRequestInput {
	typ := CreateModerationRequestInputTypeArrayOfstr

	return CreateModerationRequestInput{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *CreateModerationRequestInput) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateModerationRequestInputTypeStr
		return nil
	}

	arrayOfstr := []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfstr, "", true, true); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateModerationRequestInputTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateModerationRequestInput) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfstr != nil {
		return utils.MarshalJSON(u.ArrayOfstr, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateModerationRequest2 string

const (
	CreateModerationRequest2TextModerationLatest CreateModerationRequest2 = "text-moderation-latest"
	CreateModerationRequest2TextModerationStable CreateModerationRequest2 = "text-moderation-stable"
)

func (e CreateModerationRequest2) ToPointer() *CreateModerationRequest2 {
	return &e
}

func (e *CreateModerationRequest2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-moderation-latest":
		fallthrough
	case "text-moderation-stable":
		*e = CreateModerationRequest2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateModerationRequest2: %v", v)
	}
}

type CreateModerationRequestModelType string

const (
	CreateModerationRequestModelTypeStr                      CreateModerationRequestModelType = "str"
	CreateModerationRequestModelTypeCreateModerationRequest2 CreateModerationRequestModelType = "CreateModerationRequest_2"
)

// CreateModerationRequestModel - Two content moderations models are available: `text-moderation-stable` and `text-moderation-latest`.
//
// The default is `text-moderation-latest` which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use `text-moderation-stable`, we will provide advanced notice before updating the model. Accuracy of `text-moderation-stable` may be slightly lower than for `text-moderation-latest`.
type CreateModerationRequestModel struct {
	Str                      *string
	CreateModerationRequest2 *CreateModerationRequest2

	Type CreateModerationRequestModelType
}

func CreateCreateModerationRequestModelStr(str string) CreateModerationRequestModel {
	typ := CreateModerationRequestModelTypeStr

	return CreateModerationRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateModerationRequestModelCreateModerationRequest2(createModerationRequest2 CreateModerationRequest2) CreateModerationRequestModel {
	typ := CreateModerationRequestModelTypeCreateModerationRequest2

	return CreateModerationRequestModel{
		CreateModerationRequest2: &createModerationRequest2,
		Type:                     typ,
	}
}

func (u *CreateModerationRequestModel) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateModerationRequestModelTypeStr
		return nil
	}

	createModerationRequest2 := CreateModerationRequest2("")
	if err := utils.UnmarshalJSON(data, &createModerationRequest2, "", true, true); err == nil {
		u.CreateModerationRequest2 = &createModerationRequest2
		u.Type = CreateModerationRequestModelTypeCreateModerationRequest2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateModerationRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateModerationRequest2 != nil {
		return utils.MarshalJSON(u.CreateModerationRequest2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateModerationRequest struct {
	// The input text to classify
	Input CreateModerationRequestInput `json:"input"`
	// Two content moderations models are available: `text-moderation-stable` and `text-moderation-latest`.
	//
	// The default is `text-moderation-latest` which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use `text-moderation-stable`, we will provide advanced notice before updating the model. Accuracy of `text-moderation-stable` may be slightly lower than for `text-moderation-latest`.
	//
	Model *CreateModerationRequestModel `json:"model,omitempty"`
}

func (o *CreateModerationRequest) GetInput() CreateModerationRequestInput {
	if o == nil {
		return CreateModerationRequestInput{}
	}
	return o.Input
}

func (o *CreateModerationRequest) GetModel() *CreateModerationRequestModel {
	if o == nil {
		return nil
	}
	return o.Model
}
