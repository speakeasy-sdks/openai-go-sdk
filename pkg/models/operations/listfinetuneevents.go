package operations

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
	"net/http"
)

type ListFineTuneEventsPathParams struct {
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
}

type ListFineTuneEventsQueryParams struct {
	Stream *bool `queryParam:"style=form,explode=true,name=stream"`
}

type ListFineTuneEventsRequest struct {
	PathParams  ListFineTuneEventsPathParams
	QueryParams ListFineTuneEventsQueryParams
}

type ListFineTuneEventsResponse struct {
	ContentType                string
	ListFineTuneEventsResponse *shared.ListFineTuneEventsResponse
	StatusCode                 int
	RawResponse                *http.Response
}
