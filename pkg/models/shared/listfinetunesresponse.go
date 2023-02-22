package shared

type ListFineTunesResponse struct {
	Data   []interface{} `json:"data"`
	Object string        `json:"object"`
}
