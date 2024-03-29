// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
)

type CreateFineTuningJobRequest1 string

const (
	CreateFineTuningJobRequest1Auto CreateFineTuningJobRequest1 = "auto"
)

func (e CreateFineTuningJobRequest1) ToPointer() *CreateFineTuningJobRequest1 {
	return &e
}

func (e *CreateFineTuningJobRequest1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = CreateFineTuningJobRequest1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateFineTuningJobRequest1: %v", v)
	}
}

type BatchSizeType string

const (
	BatchSizeTypeCreateFineTuningJobRequest1 BatchSizeType = "CreateFineTuningJobRequest_1"
	BatchSizeTypeInteger                     BatchSizeType = "integer"
)

// BatchSize - Number of examples in each batch. A larger batch size means that model parameters
// are updated less frequently, but with lower variance.
type BatchSize struct {
	CreateFineTuningJobRequest1 *CreateFineTuningJobRequest1
	Integer                     *int64

	Type BatchSizeType
}

func CreateBatchSizeCreateFineTuningJobRequest1(createFineTuningJobRequest1 CreateFineTuningJobRequest1) BatchSize {
	typ := BatchSizeTypeCreateFineTuningJobRequest1

	return BatchSize{
		CreateFineTuningJobRequest1: &createFineTuningJobRequest1,
		Type:                        typ,
	}
}

func CreateBatchSizeInteger(integer int64) BatchSize {
	typ := BatchSizeTypeInteger

	return BatchSize{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *BatchSize) UnmarshalJSON(data []byte) error {

	createFineTuningJobRequest1 := CreateFineTuningJobRequest1("")
	if err := utils.UnmarshalJSON(data, &createFineTuningJobRequest1, "", true, true); err == nil {
		u.CreateFineTuningJobRequest1 = &createFineTuningJobRequest1
		u.Type = BatchSizeTypeCreateFineTuningJobRequest1
		return nil
	}

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = BatchSizeTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u BatchSize) MarshalJSON() ([]byte, error) {
	if u.CreateFineTuningJobRequest1 != nil {
		return utils.MarshalJSON(u.CreateFineTuningJobRequest1, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateFineTuningJobRequestSchemas1 string

const (
	CreateFineTuningJobRequestSchemas1Auto CreateFineTuningJobRequestSchemas1 = "auto"
)

func (e CreateFineTuningJobRequestSchemas1) ToPointer() *CreateFineTuningJobRequestSchemas1 {
	return &e
}

func (e *CreateFineTuningJobRequestSchemas1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = CreateFineTuningJobRequestSchemas1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateFineTuningJobRequestSchemas1: %v", v)
	}
}

type LearningRateMultiplierType string

const (
	LearningRateMultiplierTypeCreateFineTuningJobRequestSchemas1 LearningRateMultiplierType = "CreateFineTuningJobRequest_Schemas_1"
	LearningRateMultiplierTypeNumber                             LearningRateMultiplierType = "number"
)

// LearningRateMultiplier - Scaling factor for the learning rate. A smaller learning rate may be useful to avoid
// overfitting.
type LearningRateMultiplier struct {
	CreateFineTuningJobRequestSchemas1 *CreateFineTuningJobRequestSchemas1
	Number                             *float64

	Type LearningRateMultiplierType
}

func CreateLearningRateMultiplierCreateFineTuningJobRequestSchemas1(createFineTuningJobRequestSchemas1 CreateFineTuningJobRequestSchemas1) LearningRateMultiplier {
	typ := LearningRateMultiplierTypeCreateFineTuningJobRequestSchemas1

	return LearningRateMultiplier{
		CreateFineTuningJobRequestSchemas1: &createFineTuningJobRequestSchemas1,
		Type:                               typ,
	}
}

func CreateLearningRateMultiplierNumber(number float64) LearningRateMultiplier {
	typ := LearningRateMultiplierTypeNumber

	return LearningRateMultiplier{
		Number: &number,
		Type:   typ,
	}
}

func (u *LearningRateMultiplier) UnmarshalJSON(data []byte) error {

	createFineTuningJobRequestSchemas1 := CreateFineTuningJobRequestSchemas1("")
	if err := utils.UnmarshalJSON(data, &createFineTuningJobRequestSchemas1, "", true, true); err == nil {
		u.CreateFineTuningJobRequestSchemas1 = &createFineTuningJobRequestSchemas1
		u.Type = LearningRateMultiplierTypeCreateFineTuningJobRequestSchemas1
		return nil
	}

	number := float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = LearningRateMultiplierTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u LearningRateMultiplier) MarshalJSON() ([]byte, error) {
	if u.CreateFineTuningJobRequestSchemas1 != nil {
		return utils.MarshalJSON(u.CreateFineTuningJobRequestSchemas1, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateFineTuningJobRequestSchemasHyperparameters1 string

const (
	CreateFineTuningJobRequestSchemasHyperparameters1Auto CreateFineTuningJobRequestSchemasHyperparameters1 = "auto"
)

func (e CreateFineTuningJobRequestSchemasHyperparameters1) ToPointer() *CreateFineTuningJobRequestSchemasHyperparameters1 {
	return &e
}

func (e *CreateFineTuningJobRequestSchemasHyperparameters1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		*e = CreateFineTuningJobRequestSchemasHyperparameters1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateFineTuningJobRequestSchemasHyperparameters1: %v", v)
	}
}

type NEpochsType string

const (
	NEpochsTypeCreateFineTuningJobRequestSchemasHyperparameters1 NEpochsType = "CreateFineTuningJobRequest_Schemas_hyperparameters_1"
	NEpochsTypeInteger                                           NEpochsType = "integer"
)

// NEpochs - The number of epochs to train the model for. An epoch refers to one full cycle
// through the training dataset.
type NEpochs struct {
	CreateFineTuningJobRequestSchemasHyperparameters1 *CreateFineTuningJobRequestSchemasHyperparameters1
	Integer                                           *int64

	Type NEpochsType
}

func CreateNEpochsCreateFineTuningJobRequestSchemasHyperparameters1(createFineTuningJobRequestSchemasHyperparameters1 CreateFineTuningJobRequestSchemasHyperparameters1) NEpochs {
	typ := NEpochsTypeCreateFineTuningJobRequestSchemasHyperparameters1

	return NEpochs{
		CreateFineTuningJobRequestSchemasHyperparameters1: &createFineTuningJobRequestSchemasHyperparameters1,
		Type: typ,
	}
}

func CreateNEpochsInteger(integer int64) NEpochs {
	typ := NEpochsTypeInteger

	return NEpochs{
		Integer: &integer,
		Type:    typ,
	}
}

func (u *NEpochs) UnmarshalJSON(data []byte) error {

	createFineTuningJobRequestSchemasHyperparameters1 := CreateFineTuningJobRequestSchemasHyperparameters1("")
	if err := utils.UnmarshalJSON(data, &createFineTuningJobRequestSchemasHyperparameters1, "", true, true); err == nil {
		u.CreateFineTuningJobRequestSchemasHyperparameters1 = &createFineTuningJobRequestSchemasHyperparameters1
		u.Type = NEpochsTypeCreateFineTuningJobRequestSchemasHyperparameters1
		return nil
	}

	integer := int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = NEpochsTypeInteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u NEpochs) MarshalJSON() ([]byte, error) {
	if u.CreateFineTuningJobRequestSchemasHyperparameters1 != nil {
		return utils.MarshalJSON(u.CreateFineTuningJobRequestSchemasHyperparameters1, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Hyperparameters - The hyperparameters used for the fine-tuning job.
type Hyperparameters struct {
	// Number of examples in each batch. A larger batch size means that model parameters
	// are updated less frequently, but with lower variance.
	//
	BatchSize *BatchSize `json:"batch_size,omitempty"`
	// Scaling factor for the learning rate. A smaller learning rate may be useful to avoid
	// overfitting.
	//
	LearningRateMultiplier *LearningRateMultiplier `json:"learning_rate_multiplier,omitempty"`
	// The number of epochs to train the model for. An epoch refers to one full cycle
	// through the training dataset.
	//
	NEpochs *NEpochs `json:"n_epochs,omitempty"`
}

func (o *Hyperparameters) GetBatchSize() *BatchSize {
	if o == nil {
		return nil
	}
	return o.BatchSize
}

func (o *Hyperparameters) GetLearningRateMultiplier() *LearningRateMultiplier {
	if o == nil {
		return nil
	}
	return o.LearningRateMultiplier
}

func (o *Hyperparameters) GetNEpochs() *NEpochs {
	if o == nil {
		return nil
	}
	return o.NEpochs
}

type CreateFineTuningJobRequest2 string

const (
	CreateFineTuningJobRequest2Babbage002 CreateFineTuningJobRequest2 = "babbage-002"
	CreateFineTuningJobRequest2Davinci002 CreateFineTuningJobRequest2 = "davinci-002"
	CreateFineTuningJobRequest2Gpt35Turbo CreateFineTuningJobRequest2 = "gpt-3.5-turbo"
)

func (e CreateFineTuningJobRequest2) ToPointer() *CreateFineTuningJobRequest2 {
	return &e
}

func (e *CreateFineTuningJobRequest2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "babbage-002":
		fallthrough
	case "davinci-002":
		fallthrough
	case "gpt-3.5-turbo":
		*e = CreateFineTuningJobRequest2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateFineTuningJobRequest2: %v", v)
	}
}

type CreateFineTuningJobRequestModelType string

const (
	CreateFineTuningJobRequestModelTypeStr                         CreateFineTuningJobRequestModelType = "str"
	CreateFineTuningJobRequestModelTypeCreateFineTuningJobRequest2 CreateFineTuningJobRequestModelType = "CreateFineTuningJobRequest_2"
)

// CreateFineTuningJobRequestModel - The name of the model to fine-tune. You can select one of the
// [supported models](/docs/guides/fine-tuning/what-models-can-be-fine-tuned).
type CreateFineTuningJobRequestModel struct {
	Str                         *string
	CreateFineTuningJobRequest2 *CreateFineTuningJobRequest2

	Type CreateFineTuningJobRequestModelType
}

func CreateCreateFineTuningJobRequestModelStr(str string) CreateFineTuningJobRequestModel {
	typ := CreateFineTuningJobRequestModelTypeStr

	return CreateFineTuningJobRequestModel{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateFineTuningJobRequestModelCreateFineTuningJobRequest2(createFineTuningJobRequest2 CreateFineTuningJobRequest2) CreateFineTuningJobRequestModel {
	typ := CreateFineTuningJobRequestModelTypeCreateFineTuningJobRequest2

	return CreateFineTuningJobRequestModel{
		CreateFineTuningJobRequest2: &createFineTuningJobRequest2,
		Type:                        typ,
	}
}

func (u *CreateFineTuningJobRequestModel) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateFineTuningJobRequestModelTypeStr
		return nil
	}

	createFineTuningJobRequest2 := CreateFineTuningJobRequest2("")
	if err := utils.UnmarshalJSON(data, &createFineTuningJobRequest2, "", true, true); err == nil {
		u.CreateFineTuningJobRequest2 = &createFineTuningJobRequest2
		u.Type = CreateFineTuningJobRequestModelTypeCreateFineTuningJobRequest2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateFineTuningJobRequestModel) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.CreateFineTuningJobRequest2 != nil {
		return utils.MarshalJSON(u.CreateFineTuningJobRequest2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateFineTuningJobRequest struct {
	// The hyperparameters used for the fine-tuning job.
	Hyperparameters *Hyperparameters `json:"hyperparameters,omitempty"`
	// The name of the model to fine-tune. You can select one of the
	// [supported models](/docs/guides/fine-tuning/what-models-can-be-fine-tuned).
	//
	Model CreateFineTuningJobRequestModel `json:"model"`
	// A string of up to 18 characters that will be added to your fine-tuned model name.
	//
	// For example, a `suffix` of "custom-model-name" would produce a model name like `ft:gpt-3.5-turbo:openai:custom-model-name:7p4lURel`.
	//
	Suffix *string `default:"null" json:"suffix"`
	// The ID of an uploaded file that contains training data.
	//
	// See [upload file](/docs/api-reference/files/upload) for how to upload a file.
	//
	// Your dataset must be formatted as a JSONL file. Additionally, you must upload your file with the purpose `fine-tune`.
	//
	// See the [fine-tuning guide](/docs/guides/fine-tuning) for more details.
	//
	TrainingFile string `json:"training_file"`
	// The ID of an uploaded file that contains validation data.
	//
	// If you provide this file, the data is used to generate validation
	// metrics periodically during fine-tuning. These metrics can be viewed in
	// the fine-tuning results file.
	// The same data should not be present in both train and validation files.
	//
	// Your dataset must be formatted as a JSONL file. You must upload your file with the purpose `fine-tune`.
	//
	// See the [fine-tuning guide](/docs/guides/fine-tuning) for more details.
	//
	ValidationFile *string `json:"validation_file,omitempty"`
}

func (c CreateFineTuningJobRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateFineTuningJobRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateFineTuningJobRequest) GetHyperparameters() *Hyperparameters {
	if o == nil {
		return nil
	}
	return o.Hyperparameters
}

func (o *CreateFineTuningJobRequest) GetModel() CreateFineTuningJobRequestModel {
	if o == nil {
		return CreateFineTuningJobRequestModel{}
	}
	return o.Model
}

func (o *CreateFineTuningJobRequest) GetSuffix() *string {
	if o == nil {
		return nil
	}
	return o.Suffix
}

func (o *CreateFineTuningJobRequest) GetTrainingFile() string {
	if o == nil {
		return ""
	}
	return o.TrainingFile
}

func (o *CreateFineTuningJobRequest) GetValidationFile() *string {
	if o == nil {
		return nil
	}
	return o.ValidationFile
}
