# Audio
(*Audio*)

## Overview

Learn how to turn audio into text or text into audio.

### Available Operations

* [CreateSpeech](#createspeech) - Generates audio from the input text.
* [CreateTranscription](#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](#createtranslation) - Translates audio into English.

## CreateSpeech

Generates audio from the input text.

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
    res, err := s.Audio.CreateSpeech(ctx, shared.CreateSpeechRequest{
        Input: "<value>",
        Model: shared.CreateCreateSpeechRequestModelCreateSpeechRequest2(
        shared.CreateSpeechRequest2Tts1,
        ),
        Voice: shared.VoiceFable,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Stream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateSpeechRequest](../../pkg/models/shared/createspeechrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateSpeechResponse](../../pkg/models/operations/createspeechresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateTranscription

Transcribes audio into the input language.

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
    res, err := s.Audio.CreateTranscription(ctx, shared.CreateTranscriptionRequest{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("0xe08fcc1Fd5"),
            FileName: "buckinghamshire.gif",
        },
        Model: shared.CreateCreateTranscriptionRequestModelCreateTranscriptionRequest2(
        shared.CreateTranscriptionRequest2Whisper1,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.CreateTranscriptionRequest](../../pkg/models/shared/createtranscriptionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateTranscriptionResponse](../../pkg/models/operations/createtranscriptionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateTranslation

Translates audio into English.

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
    res, err := s.Audio.CreateTranslation(ctx, shared.CreateTranslationRequest{
        File: shared.CreateTranslationRequestFile{
            Content: []byte("0xa45ca6c4DE"),
            FileName: "reggae_toys_silver.gif",
        },
        Model: shared.CreateCreateTranslationRequestModelCreateTranslationRequest2(
        shared.CreateTranslationRequest2Whisper1,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.CreateTranslationRequest](../../pkg/models/shared/createtranslationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateTranslationResponse](../../pkg/models/operations/createtranslationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
