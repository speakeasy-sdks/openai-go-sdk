// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type FineTuningJobEventLevel string

const (
	FineTuningJobEventLevelInfo  FineTuningJobEventLevel = "info"
	FineTuningJobEventLevelWarn  FineTuningJobEventLevel = "warn"
	FineTuningJobEventLevelError FineTuningJobEventLevel = "error"
)

func (e FineTuningJobEventLevel) ToPointer() *FineTuningJobEventLevel {
	return &e
}

func (e *FineTuningJobEventLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		*e = FineTuningJobEventLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FineTuningJobEventLevel: %v", v)
	}
}

type FineTuningJobEvent struct {
	CreatedAt int64                   `json:"created_at"`
	ID        string                  `json:"id"`
	Level     FineTuningJobEventLevel `json:"level"`
	Message   string                  `json:"message"`
	Object    string                  `json:"object"`
}

func (o *FineTuningJobEvent) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *FineTuningJobEvent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *FineTuningJobEvent) GetLevel() FineTuningJobEventLevel {
	if o == nil {
		return FineTuningJobEventLevel("")
	}
	return o.Level
}

func (o *FineTuningJobEvent) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *FineTuningJobEvent) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}
