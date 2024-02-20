// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RunStepDetailsToolCallsCodeOutputLogsObjectType - Always `logs`.
type RunStepDetailsToolCallsCodeOutputLogsObjectType string

const (
	RunStepDetailsToolCallsCodeOutputLogsObjectTypeLogs RunStepDetailsToolCallsCodeOutputLogsObjectType = "logs"
)

func (e RunStepDetailsToolCallsCodeOutputLogsObjectType) ToPointer() *RunStepDetailsToolCallsCodeOutputLogsObjectType {
	return &e
}

func (e *RunStepDetailsToolCallsCodeOutputLogsObjectType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "logs":
		*e = RunStepDetailsToolCallsCodeOutputLogsObjectType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RunStepDetailsToolCallsCodeOutputLogsObjectType: %v", v)
	}
}

// RunStepDetailsToolCallsCodeOutputLogsObject - Text output from the Code Interpreter tool call as part of a run step.
type RunStepDetailsToolCallsCodeOutputLogsObject struct {
	// The text output from the Code Interpreter tool call.
	Logs string `json:"logs"`
	// Always `logs`.
	Type RunStepDetailsToolCallsCodeOutputLogsObjectType `json:"type"`
}

func (o *RunStepDetailsToolCallsCodeOutputLogsObject) GetLogs() string {
	if o == nil {
		return ""
	}
	return o.Logs
}

func (o *RunStepDetailsToolCallsCodeOutputLogsObject) GetType() RunStepDetailsToolCallsCodeOutputLogsObjectType {
	if o == nil {
		return RunStepDetailsToolCallsCodeOutputLogsObjectType("")
	}
	return o.Type
}
