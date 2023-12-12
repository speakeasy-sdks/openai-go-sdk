# FineTunes
(*FineTunes*)

## Overview

Manage legacy fine-tuning jobs to tailor a model to your specific training data.

### Available Operations

* [~~CancelFineTune~~](#cancelfinetune) - Immediately cancel a fine-tune job.
 :warning: **Deprecated**
* [~~CreateFineTune~~](#createfinetune) - Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/legacy-fine-tuning)
 :warning: **Deprecated**
* [~~ListFineTuneEvents~~](#listfinetuneevents) - Get fine-grained status updates for a fine-tune job.
 :warning: **Deprecated**
* [~~ListFineTunes~~](#listfinetunes) - List your organization's fine-tuning jobs
 :warning: **Deprecated**
* [~~RetrieveFineTune~~](#retrievefinetune) - Gets info about the fine-tune job.

[Learn more about fine-tuning](/docs/guides/legacy-fine-tuning)
 :warning: **Deprecated**

## ~~CancelFineTune~~

Immediately cancel a fine-tune job.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var fineTuneID string = "ft-AF1WoRqd3aJAHsqc9NY7iL8F"

    ctx := context.Background()
    res, err := s.FineTunes.CancelFineTune(ctx, fineTuneID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTune != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `fineTuneID`                                          | *string*                                              | :heavy_check_mark:                                    | The ID of the fine-tune job to cancel<br/>            | ft-AF1WoRqd3aJAHsqc9NY7iL8F                           |


### Response

**[*operations.CancelFineTuneResponse](../../pkg/models/operations/cancelfinetuneresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~CreateFineTune~~

Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/legacy-fine-tuning)


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.FineTunes.CreateFineTune(ctx, shared.CreateFineTuneRequest{
        ClassificationBetas: []float64{
            0.6,
            1,
            1.5,
            2,
        },
        Hyperparameters: &shared.Hyperparameters{
            NEpochs: shared.CreateNEpochsInteger(
            399302,
            ),
        },
        Model: shared.CreateCreateFineTuneRequestModelCreateFineTuneRequest2(
        shared.CreateFineTuneRequest2Curie,
        ),
        TrainingFile: "file-abc123",
        ValidationFile: openaigosdk.String("file-abc123"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTune != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateFineTuneRequest](../../pkg/models/shared/createfinetunerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateFineTuneResponse](../../pkg/models/operations/createfinetuneresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~ListFineTuneEvents~~

Get fine-grained status updates for a fine-tune job.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var fineTuneID string = "ft-AF1WoRqd3aJAHsqc9NY7iL8F"

    var stream *bool = false

    ctx := context.Background()
    res, err := s.FineTunes.ListFineTuneEvents(ctx, fineTuneID, stream)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListFineTuneEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                               | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                             | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                      | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `fineTuneID`                                                                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                      | The ID of the fine-tune job to get events for.<br/>                                                                                                                                                                                                                                                                                                                                                                                                     | ft-AF1WoRqd3aJAHsqc9NY7iL8F                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `stream`                                                                                                                                                                                                                                                                                                                                                                                                                                                | **bool*                                                                                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                      | Whether to stream events for the fine-tune job. If set to true,<br/>events will be sent as data-only<br/>[server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format)<br/>as they become available. The stream will terminate with a<br/>`data: [DONE]` message when the job is finished (succeeded, cancelled,<br/>or failed).<br/><br/>If set to false, only events generated so far will be returned.<br/> |                                                                                                                                                                                                                                                                                                                                                                                                                                                         |


### Response

**[*operations.ListFineTuneEventsResponse](../../pkg/models/operations/listfinetuneeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~ListFineTunes~~

List your organization's fine-tuning jobs


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.FineTunes.ListFineTunes(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListFineTunesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListFineTunesResponse](../../pkg/models/operations/listfinetunesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ~~RetrieveFineTune~~

Gets info about the fine-tune job.

[Learn more about fine-tuning](/docs/guides/legacy-fine-tuning)


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var fineTuneID string = "ft-AF1WoRqd3aJAHsqc9NY7iL8F"

    ctx := context.Background()
    res, err := s.FineTunes.RetrieveFineTune(ctx, fineTuneID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTune != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `fineTuneID`                                          | *string*                                              | :heavy_check_mark:                                    | The ID of the fine-tune job<br/>                      | ft-AF1WoRqd3aJAHsqc9NY7iL8F                           |


### Response

**[*operations.RetrieveFineTuneResponse](../../pkg/models/operations/retrievefinetuneresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
