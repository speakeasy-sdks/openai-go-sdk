package operations

type RetrieveEnginePathParams struct {
	EngineID string `pathParam:"style=simple,explode=false,name=engine_id"`
}

type RetrieveEngineRequest struct {
	PathParams RetrieveEnginePathParams
}

type RetrieveEngineResponse struct {
	ContentType string
	Engine      interface{}
	StatusCode  int
}
