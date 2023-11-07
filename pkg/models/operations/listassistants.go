// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
	"net/http"
)

// QueryParamOrder - Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.
type QueryParamOrder string

const (
	QueryParamOrderAsc  QueryParamOrder = "asc"
	QueryParamOrderDesc QueryParamOrder = "desc"
)

func (e QueryParamOrder) ToPointer() *QueryParamOrder {
	return &e
}

func (e *QueryParamOrder) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = QueryParamOrder(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamOrder: %v", v)
	}
}

type ListAssistantsRequest struct {
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
	Order *QueryParamOrder `default:"desc" queryParam:"style=form,explode=true,name=order"`
}

func (l ListAssistantsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAssistantsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAssistantsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListAssistantsRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListAssistantsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAssistantsRequest) GetOrder() *QueryParamOrder {
	if o == nil {
		return nil
	}
	return o.Order
}

type ListAssistantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ListAssistantsResponse *shared.ListAssistantsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAssistantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAssistantsResponse) GetListAssistantsResponse() *shared.ListAssistantsResponse {
	if o == nil {
		return nil
	}
	return o.ListAssistantsResponse
}

func (o *ListAssistantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAssistantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
