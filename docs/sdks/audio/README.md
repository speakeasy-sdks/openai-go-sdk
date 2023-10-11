# Audio
(*Audio*)

## Overview

Learn how to turn audio into text.

### Available Operations

* [CreateTranscription](#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](#createtranslation) - Translates audio into English.

## CreateTranscription

Transcribes audio into the input language.

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
    res, err := s.Audio.CreateTranscription(ctx, shared.CreateTranscriptionRequest{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("\#BbTW'zX9"),
            File: "Buckinghamshire",
        },
        Model: shared.CreateCreateTranscriptionRequestModelCreateTranscriptionRequestModel2(
        shared.CreateTranscriptionRequestModel2Whisper1,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTranscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.CreateTranscriptionRequest](../../models/shared/createtranscriptionrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateTranscriptionResponse](../../models/operations/createtranscriptionresponse.md), error**


## CreateTranslation

Translates audio into English.

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
    res, err := s.Audio.CreateTranslation(ctx, shared.CreateTranslationRequest{
        File: shared.CreateTranslationRequestFile{
            Content: []byte("M57UL;W3rx"),
            File: "Reggae Toys silver",
        },
        Model: shared.CreateCreateTranslationRequestModelCreateTranslationRequestModel2(
        shared.CreateTranslationRequestModel2Whisper1,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTranslationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateTranslationRequest](../../models/shared/createtranslationrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateTranslationResponse](../../models/operations/createtranslationresponse.md), error**

