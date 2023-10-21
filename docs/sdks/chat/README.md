# Chat
(*Chat*)

## Overview

Given a list of messages comprising a conversation, the model will return a response.

### Available Operations

* [CreateChatCompletion](#createchatcompletion) - Creates a model response for the given chat conversation.

## CreateChatCompletion

Creates a model response for the given chat conversation.

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
    res, err := s.Chat.CreateChatCompletion(ctx, shared.CreateChatCompletionRequest{
        FunctionCall: shared.CreateCreateChatCompletionRequestFunctionCallChatCompletionFunctionCallOption(
                shared.ChatCompletionFunctionCallOption{
                    Name: "string",
                },
        ),
        Functions: []shared.ChatCompletionFunctions{
            shared.ChatCompletionFunctions{
                Name: "string",
                Parameters: map[string]interface{}{
                    "key": "string",
                },
            },
        },
        LogitBias: map[string]int64{
            "key": 544683,
        },
        Messages: []shared.ChatCompletionRequestMessage{
            shared.ChatCompletionRequestMessage{
                Content: "string",
                FunctionCall: &shared.ChatCompletionRequestMessageFunctionCall{
                    Arguments: "string",
                    Name: "string",
                },
                Role: shared.ChatCompletionRequestMessageRoleUser,
            },
        },
        Model: shared.CreateCreateChatCompletionRequestModelCreateChatCompletionRequestModel2(
        shared.CreateChatCompletionRequestModel2Gpt35Turbo,
        ),
        N: openaigosdk.Int64(1),
        Stop: shared.CreateCreateChatCompletionRequestStopArrayOfstr(
                []string{
                    "string",
                },
        ),
        Temperature: openaigosdk.Float64(1),
        TopP: openaigosdk.Float64(1),
        User: openaigosdk.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateChatCompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CreateChatCompletionRequest](../../models/shared/createchatcompletionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateChatCompletionResponse](../../models/operations/createchatcompletionresponse.md), error**

