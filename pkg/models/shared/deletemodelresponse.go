package shared

type DeleteModelResponse struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}
