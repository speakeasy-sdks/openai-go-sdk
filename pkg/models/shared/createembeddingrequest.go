// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type CreateEmbeddingRequestInputType string

const (
	CreateEmbeddingRequestInputTypeStr                   CreateEmbeddingRequestInputType = "str"
	CreateEmbeddingRequestInputTypeArrayOfstr            CreateEmbeddingRequestInputType = "arrayOfstr"
	CreateEmbeddingRequestInputTypeArrayOfinteger        CreateEmbeddingRequestInputType = "arrayOfinteger"
	CreateEmbeddingRequestInputTypeArrayOfarrayOfinteger CreateEmbeddingRequestInputType = "arrayOfarrayOfinteger"
)

type CreateEmbeddingRequestInput struct {
	Str                   *string
	ArrayOfstr            []string
	ArrayOfinteger        []int64
	ArrayOfarrayOfinteger [][]int64

	Type CreateEmbeddingRequestInputType
}

func CreateCreateEmbeddingRequestInputStr(str string) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeStr

	return CreateEmbeddingRequestInput{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateEmbeddingRequestInputArrayOfstr(arrayOfstr []string) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeArrayOfstr

	return CreateEmbeddingRequestInput{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func CreateCreateEmbeddingRequestInputArrayOfinteger(arrayOfinteger []int64) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeArrayOfinteger

	return CreateEmbeddingRequestInput{
		ArrayOfinteger: arrayOfinteger,
		Type:           typ,
	}
}

func CreateCreateEmbeddingRequestInputArrayOfarrayOfinteger(arrayOfarrayOfinteger [][]int64) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeArrayOfarrayOfinteger

	return CreateEmbeddingRequestInput{
		ArrayOfarrayOfinteger: arrayOfarrayOfinteger,
		Type:                  typ,
	}
}

func (u *CreateEmbeddingRequestInput) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateEmbeddingRequestInputTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateEmbeddingRequestInputTypeArrayOfstr
		return nil
	}

	arrayOfinteger := []int64{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfinteger); err == nil {
		u.ArrayOfinteger = arrayOfinteger
		u.Type = CreateEmbeddingRequestInputTypeArrayOfinteger
		return nil
	}

	arrayOfarrayOfinteger := [][]int64{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfarrayOfinteger); err == nil {
		u.ArrayOfarrayOfinteger = arrayOfarrayOfinteger
		u.Type = CreateEmbeddingRequestInputTypeArrayOfarrayOfinteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateEmbeddingRequestInput) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	if u.ArrayOfinteger != nil {
		return json.Marshal(u.ArrayOfinteger)
	}

	if u.ArrayOfarrayOfinteger != nil {
		return json.Marshal(u.ArrayOfarrayOfinteger)
	}

	return nil, nil
}

type CreateEmbeddingRequest struct {
	// Input text to embed, encoded as a string or array of tokens. To embed multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed the max input tokens for the model (8191 tokens for `text-embedding-ada-002`). [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.
	//
	Input CreateEmbeddingRequestInput `json:"input"`
	Model interface{}                 `json:"model"`
	User  interface{}                 `json:"user,omitempty"`
}
