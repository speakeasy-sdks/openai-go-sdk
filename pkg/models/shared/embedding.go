// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Embedding - Represents an embedding vector returned by embedding endpoint.
type Embedding struct {
	// The embedding vector, which is a list of floats. The length of vector depends on the model as listed in the [embedding guide](/docs/guides/embeddings).
	//
	Embedding []float64 `json:"embedding"`
	// The index of the embedding in the list of embeddings.
	Index int64 `json:"index"`
	// The object type, which is always "embedding".
	Object string `json:"object"`
}

func (o *Embedding) GetEmbedding() []float64 {
	if o == nil {
		return []float64{}
	}
	return o.Embedding
}

func (o *Embedding) GetIndex() int64 {
	if o == nil {
		return 0
	}
	return o.Index
}

func (o *Embedding) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}