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
    res, err := s.Completions.CreateCompletion(ctx, shared.CreateCompletionRequest{
        LogitBias: map[string]int64{
            "red": 242695,
        },
        MaxTokens: openaigosdk.Int64(16),
        Model: shared.CreateCreateCompletionRequestModelStr(
        "optimistic",
        ),
        N: openaigosdk.Int64(1),
        Prompt: shared.CreateCreateCompletionRequestPromptArrayOfinteger(
                []int64{
                    [,
                    1,
                    2,
                    1,
                    2,
                    ,,
                     ,
                    3,
                    1,
                    8,
                    ,,
                     ,
                    2,
                    5,
                    7,
                    ,,
                     ,
                    1,
                    3,
                    3,
                    2,
                    ,,
                     ,
                    1,
                    3,
                    ],
                },
        ),
        Stop: shared.CreateCreateCompletionRequestStopArrayOfstr(
                []string{
                    "[",
                    "\"",
                    "\",
                    "n",
                    "\"",
                    "]",
                },
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateCompletionRequest](../../models/shared/createcompletionrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateCompletionResponse](../../models/operations/createcompletionresponse.md), error**

