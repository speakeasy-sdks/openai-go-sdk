# Moderations
(*Moderations*)

## Overview

Given a input text, outputs if the model classifies it as potentially harmful.

### Available Operations

* [CreateModeration](#createmoderation) - Classifies if text is potentially harmful.

## CreateModeration

Classifies if text is potentially harmful.

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
    res, err := s.Moderations.CreateModeration(ctx, shared.CreateModerationRequest{
        Input: shared.CreateCreateModerationRequestInputArrayOfstr(
                []string{
                    "I want to kill them.",
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
| sdkerrors.SDKError | 4xx-5xx            | */*                |
