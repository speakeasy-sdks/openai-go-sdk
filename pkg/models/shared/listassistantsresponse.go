// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListAssistantsResponse struct {
	Data    []AssistantObject `json:"data"`
	FirstID string            `json:"first_id"`
	HasMore bool              `json:"has_more"`
	LastID  string            `json:"last_id"`
	Object  string            `json:"object"`
}

func (o *ListAssistantsResponse) GetData() []AssistantObject {
	if o == nil {
		return []AssistantObject{}
	}
	return o.Data
}

func (o *ListAssistantsResponse) GetFirstID() string {
	if o == nil {
		return ""
	}
	return o.FirstID
}

func (o *ListAssistantsResponse) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *ListAssistantsResponse) GetLastID() string {
	if o == nil {
		return ""
	}
	return o.LastID
}

func (o *ListAssistantsResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}