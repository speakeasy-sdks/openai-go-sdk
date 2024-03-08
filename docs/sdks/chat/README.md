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
    res, err := s.Chat.CreateChatCompletion(ctx, shared.CreateChatCompletionRequest{
        Messages: []shared.ChatCompletionRequestMessage{
            shared.CreateChatCompletionRequestMessageChatCompletionRequestToolMessage(
                shared.ChatCompletionRequestToolMessage{
                    Content: "<value>",
                    Role: shared.ChatCompletionRequestToolMessageRoleTool,
                    ToolCallID: "<value>",
                },
            ),
        },
        Model: shared.CreateCreateChatCompletionRequestModelTwo(
        shared.TwoGpt35Turbo,
        ),
        N: openaigosdk.Int64(1),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.CreateChatCompletionRequest](../../pkg/models/shared/createchatcompletionrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateChatCompletionResponse](../../pkg/models/operations/createchatcompletionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
