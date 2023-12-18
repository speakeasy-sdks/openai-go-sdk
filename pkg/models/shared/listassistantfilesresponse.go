// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListAssistantFilesResponse struct {
	Data    []AssistantFileObject `json:"data"`
	FirstID string                `json:"first_id"`
	HasMore bool                  `json:"has_more"`
	LastID  string                `json:"last_id"`
	Object  string                `json:"object"`
}

func (o *ListAssistantFilesResponse) GetData() []AssistantFileObject {
	if o == nil {
		return []AssistantFileObject{}
	}
	return o.Data
}

func (o *ListAssistantFilesResponse) GetFirstID() string {
	if o == nil {
		return ""
	}
	return o.FirstID
}

func (o *ListAssistantFilesResponse) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *ListAssistantFilesResponse) GetLastID() string {
	if o == nil {
		return ""
	}
	return o.LastID
}

func (o *ListAssistantFilesResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}