// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
	"net/http"
)

// ListMessagesQueryParamOrder - Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
type ListMessagesQueryParamOrder string

const (
	ListMessagesQueryParamOrderAsc  ListMessagesQueryParamOrder = "asc"
	ListMessagesQueryParamOrderDesc ListMessagesQueryParamOrder = "desc"
)

func (e ListMessagesQueryParamOrder) ToPointer() *ListMessagesQueryParamOrder {
	return &e
}

func (e *ListMessagesQueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = ListMessagesQueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListMessagesQueryParamOrder: %v", v)
	}
}

type ListMessagesRequest struct {
	// A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.
	//
	After *string `queryParam:"style=form,explode=true,name=after"`
	// A cursor for use in pagination. `before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include before=obj_foo in order to fetch the previous page of the list.
	//
	Before *string `queryParam:"style=form,explode=true,name=before"`
	// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.
	//
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
	//
	Order *ListMessagesQueryParamOrder `default:"desc" queryParam:"style=form,explode=true,name=order"`
	// The ID of the [thread](/docs/api-reference/threads) the messages belong to.
	ThreadID string `pathParam:"style=simple,explode=false,name=thread_id"`
}

func (l ListMessagesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListMessagesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListMessagesRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListMessagesRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListMessagesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMessagesRequest) GetOrder() *ListMessagesQueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ListMessagesRequest) GetThreadID() string {
	if o == nil {
		return ""
	}
	return o.ThreadID
}

type ListMessagesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListMessagesResponse *shared.ListMessagesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListMessagesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListMessagesResponse) GetListMessagesResponse() *shared.ListMessagesResponse {
	if o == nil {
		return nil
	}
	return o.ListMessagesResponse
}

func (o *ListMessagesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListMessagesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
