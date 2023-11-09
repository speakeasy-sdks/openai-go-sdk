# Moderations
(*Moderations*)

## Overview

Given a input text, outputs if the model classifies it as violating OpenAI's content policy.

### Available Operations

* [CreateModeration](#createmoderation) - Classifies if text violates OpenAI's Content Policy

## CreateModeration

Classifies if text violates OpenAI's Content Policy

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
    res, err := s.Moderations.CreateModeration(ctx, shared.CreateModerationRequest{
        Input: shared.CreateCreateModerationRequestInputArrayOfstr(
                []string{
                    "I",
                    " ",
                    "w",
                    "a",
                    "n",
                    "t",
                    " ",
                    "t",
                    "o",
                    " ",
                    "k",
                    "i",
                    "l",
                    "l",
                    " ",
                    "t",
                    "h",
                    "e",
                    "m",
                    ".",
                },
        ),
        Model: shared.CreateCreateModerationRequestModelCreateModerationRequest2(
        shared.CreateModerationRequest2TextModerationStable,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateModerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.CreateModerationRequest](../../pkg/models/shared/createmoderationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateModerationResponse](../../pkg/models/operations/createmoderationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
