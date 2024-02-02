# Completions
(*Completions*)

## Overview

Given a prompt, the model will return one or more predicted completions, and can also return the probabilities of alternative tokens at each position.

### Available Operations

* [CreateCompletion](#createcompletion) - Creates a completion for the provided prompt and parameters.

## CreateCompletion

Creates a completion for the provided prompt and parameters.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Completions.CreateCompletion(ctx, shared.CreateCompletionRequest{
        LogitBias: map[string]int64{
            "key": 160667,
        },
        MaxTokens: openaigosdk.Int64(16),
        Model: shared.CreateCreateCompletionRequestModelStr(
        "string",
        ),
        N: openaigosdk.Int64(1),
        Prompt: shared.CreatePromptStr(
        "This is a test.",
        ),
        Stop: shared.CreateCreateCompletionRequestStopStr(
        "
        ",
        ),
        Suffix: openaigosdk.String("test."),
        Temperature: openaigosdk.Float64(1),
        TopP: openaigosdk.Float64(1),
        User: openaigosdk.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.CreateCompletionRequest](../../pkg/models/shared/createcompletionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateCompletionResponse](../../pkg/models/operations/createcompletionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
