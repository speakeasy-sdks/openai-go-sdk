# Chat
(*.Chat*)

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
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Chat.CreateChatCompletion(ctx, shared.CreateChatCompletionRequest{
        FunctionCall: shared.CreateCreateChatCompletionRequestFunctionCallSchemas(
                shared.Schemas{
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
            shared.CreateChatCompletionRequestMessageUserMessage(
                shared.UserMessage{
                    Content: shared.CreateContentArrayOfChatCompletionRequestMessageContentPart(
                            []shared.ChatCompletionRequestMessageContentPart{
                                shared.CreateChatCompletionRequestMessageContentPartTextContentPart(
                                    shared.TextContentPart{
                                        Text: "string",
                                        Type: shared.SchemasChatCompletionRequestMessageContentPartTextTypeText,
                                    },
                                ),
                            },
                    ),
                    Role: shared.SchemasChatCompletionRequestUserMessageRoleUser,
                },
            ),
        },
        Model: shared.CreateCreateChatCompletionRequestModelTwo(
        shared.TwoGpt35Turbo,
        ),
        N: openaigosdk.Int64(1),
        ResponseFormat: &shared.ResponseFormat{
            Type: shared.CreateChatCompletionRequestTypeJSONObject.ToPointer(),
        },
        Stop: shared.CreateStopStr(
        "string",
        ),
        Temperature: openaigosdk.Float64(1),
        ToolChoice: shared.CreateChatCompletionToolChoiceOptionChatCompletionToolChoiceOption1(
        shared.ChatCompletionToolChoiceOption1None,
        ),
        Tools: []shared.ChatCompletionTool{
            shared.ChatCompletionTool{
                Function: shared.ChatCompletionToolFunction{
                    Name: "string",
                    Parameters: map[string]interface{}{
                        "key": "string",
                    },
                },
                Type: shared.ChatCompletionToolTypeFunction,
            },
        },
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

