package shared

type ListFilesResponse struct {
	Data   []interface{} `json:"data"`
	Object string        `json:"object"`
}
