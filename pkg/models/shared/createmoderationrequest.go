// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/utils"
)

type CreateModerationRequestInputType string

const (
	CreateModerationRequestInputTypeStr        CreateModerationRequestInputType = "str"
	CreateModerationRequestInputTypeArrayOfstr CreateModerationRequestInputType = "arrayOfstr"
)

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

// CreateModerationRequestModel2 - Two content moderations models are available: `text-moderation-stable` and `text-moderation-latest`.
//
// The default is `text-moderation-latest` which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use `text-moderation-stable`, we will provide advanced notice before updating the model. Accuracy of `text-moderation-stable` may be slightly lower than for `text-moderation-latest`.
type CreateModerationRequestModel2 string

const (
	CreateModerationRequestModel2TextModerationLatest CreateModerationRequestModel2 = "text-moderation-latest"
	CreateModerationRequestModel2TextModerationStable CreateModerationRequestModel2 = "text-moderation-stable"
)

func (e CreateModerationRequestModel2) ToPointer() *CreateModerationRequestModel2 {
	return &e
}

func (e *CreateModerationRequestModel2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-moderation-latest":
		fallthrough
	case "text-moderation-stable":
		*e = CreateModerationRequestModel2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateModerationRequestModel2: %v", v)
	}
}

type CreateModerationRequestModelType string

const (
	CreateModerationRequestModelTypeStr                           CreateModerationRequestModelType = "str"
	CreateModerationRequestModelTypeCreateModerationRequestModel2 CreateModerationRequestModelType = "CreateModerationRequest_model_2"
)

type CreateModerationRequestModel struct {
	Str                           *string
	CreateModerationRequestModel2 *CreateModerationRequestModel2

	Type CreateModerationRequestModelType
}

func CreateCreateModerationRequestModelStr(str string) CreateModerationRequestModel {
	typ := CreateModerationRequestModelTypeStr

	return CreateModerationRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateModerationRequestModelCreateModerationRequestModel2(createModerationRequestModel2 CreateModerationRequestModel2) CreateModerationRequestModel {
	typ := CreateModerationRequestModelTypeCreateModerationRequestModel2

	return CreateModerationRequestModel{
		CreateModerationRequestModel2: &createModerationRequestModel2,
		Type:                          typ,
	}
}

func (u *CreateModerationRequestModel) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateModerationRequestModelTypeStr
		return nil
	}

	createModerationRequestModel2 := CreateModerationRequestModel2("")
	if err := utils.UnmarshalJSON(data, &createModerationRequestModel2, "", true, true); err == nil {
		u.CreateModerationRequestModel2 = &createModerationRequestModel2
		u.Type = CreateModerationRequestModelTypeCreateModerationRequestModel2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateModerationRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateModerationRequestModel2 != nil {
		return utils.MarshalJSON(u.CreateModerationRequestModel2, "", true)
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
