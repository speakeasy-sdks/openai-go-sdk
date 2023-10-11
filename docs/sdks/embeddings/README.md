# Embeddings
(*Embeddings*)

## Overview

Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.

### Available Operations

* [CreateEmbedding](#createembedding) - Creates an embedding vector representing the input text.

## CreateEmbedding

Creates an embedding vector representing the input text.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v2"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Embeddings.CreateEmbedding(ctx, shared.CreateEmbeddingRequest{
        Input: shared.CreateCreateEmbeddingRequestInputStr(
        "The quick brown fox jumped over the lazy dog",
        ),
        Model: shared.CreateCreateEmbeddingRequestModelCreateEmbeddingRequestModel2(
        shared.CreateEmbeddingRequestModel2TextEmbeddingAda002,
        ),
        User: openaigosdk.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateEmbeddingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateEmbeddingRequest](../../models/shared/createembeddingrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateEmbeddingResponse](../../models/operations/createembeddingresponse.md), error**

