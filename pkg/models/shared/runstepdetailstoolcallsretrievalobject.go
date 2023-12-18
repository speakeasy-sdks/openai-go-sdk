// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Retrieval - For now, this is always going to be an empty object.
type Retrieval struct {
}

// RunStepDetailsToolCallsRetrievalObjectType - The type of tool call. This is always going to be `retrieval` for this type of tool call.
type RunStepDetailsToolCallsRetrievalObjectType string

const (
	RunStepDetailsToolCallsRetrievalObjectTypeRetrieval RunStepDetailsToolCallsRetrievalObjectType = "retrieval"
)

func (e RunStepDetailsToolCallsRetrievalObjectType) ToPointer() *RunStepDetailsToolCallsRetrievalObjectType {
	return &e
}

func (e *RunStepDetailsToolCallsRetrievalObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "retrieval":
		*e = RunStepDetailsToolCallsRetrievalObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunStepDetailsToolCallsRetrievalObjectType: %v", v)
	}
}

type RunStepDetailsToolCallsRetrievalObject struct {
	// The ID of the tool call object.
	ID string `json:"id"`
	// For now, this is always going to be an empty object.
	Retrieval Retrieval `json:"retrieval"`
	// The type of tool call. This is always going to be `retrieval` for this type of tool call.
	Type RunStepDetailsToolCallsRetrievalObjectType `json:"type"`
}

func (o *RunStepDetailsToolCallsRetrievalObject) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RunStepDetailsToolCallsRetrievalObject) GetRetrieval() Retrieval {
	if o == nil {
		return Retrieval{}
	}
	return o.Retrieval
}

func (o *RunStepDetailsToolCallsRetrievalObject) GetType() RunStepDetailsToolCallsRetrievalObjectType {
	if o == nil {
		return RunStepDetailsToolCallsRetrievalObjectType("")
	}
	return o.Type
}