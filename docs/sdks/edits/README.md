# Edits
(*Edits*)

## Overview

Given a prompt and an instruction, the model will return an edited version of the prompt.

### Available Operations

* [~~CreateEdit~~](#createedit) - Creates a new edit for the provided input, instruction, and parameters. :warning: **Deprecated**

## ~~CreateEdit~~

Creates a new edit for the provided input, instruction, and parameters.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.Edits.CreateEdit(ctx, shared.CreateEditRequest{
        Input: openaigosdk.String("What day of the wek is it?"),
        Instruction: "Fix the spelling mistakes.",
        Model: shared.CreateCreateEditRequestModelCreateEditRequest2(
        shared.CreateEditRequest2TextDavinciEdit001,
        ),
        N: openaigosdk.Int64(1),
        Temperature: openaigosdk.Float64(1),
        TopP: openaigosdk.Float64(1),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateEditResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateEditRequest](../../pkg/models/shared/createeditrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateEditResponse](../../pkg/models/operations/createeditresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
