package operations

type CancelFineTunePathParams struct {
	FineTuneID string `pathParam:"style=simple,explode=false,name=fine_tune_id"`
}

type CancelFineTuneRequest struct {
	PathParams CancelFineTunePathParams
}

type CancelFineTuneResponse struct {
	ContentType string
	FineTune    *interface{}
	StatusCode  int
}
