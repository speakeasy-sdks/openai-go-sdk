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
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Embeddings.CreateEmbedding(ctx, shared.CreateEmbeddingRequest{
        EncodingFormat: shared.EncodingFormatFloat.ToPointer(),
        Input: shared.CreateInputStr(
        "The quick brown fox jumped over the lazy dog",
        ),
        Model: shared.CreateCreateEmbeddingRequestModelCreateEmbeddingRequest2(
        shared.CreateEmbeddingRequest2TextEmbedding3Small,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateEmbeddingRequest](../../pkg/models/shared/createembeddingrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateEmbeddingResponse](../../pkg/models/operations/createembeddingresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
