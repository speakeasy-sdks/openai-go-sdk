# FineTuning
(*FineTuning*)

## Overview

Manage fine-tuning jobs to tailor a model to your specific training data.

### Available Operations

* [CancelFineTuningJob](#cancelfinetuningjob) - Immediately cancel a fine-tune job.

* [CreateFineTuningJob](#createfinetuningjob) - Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/fine-tuning)

* [ListFineTuningEvents](#listfinetuningevents) - Get status updates for a fine-tuning job.

* [ListPaginatedFineTuningJobs](#listpaginatedfinetuningjobs) - List your organization's fine-tuning jobs

* [RetrieveFineTuningJob](#retrievefinetuningjob) - Get info about a fine-tuning job.

[Learn more about fine-tuning](/docs/guides/fine-tuning)


## CancelFineTuningJob

Immediately cancel a fine-tune job.


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


    var fineTuningJobID string = "ft-AF1WoRqd3aJAHsqc9NY7iL8F"

    ctx := context.Background()
    res, err := s.FineTuning.CancelFineTuningJob(ctx, fineTuningJobID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTuningJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `fineTuningJobID`                                     | *string*                                              | :heavy_check_mark:                                    | The ID of the fine-tuning job to cancel.<br/>         | ft-AF1WoRqd3aJAHsqc9NY7iL8F                           |


### Response

**[*operations.CancelFineTuningJobResponse](../../models/operations/cancelfinetuningjobresponse.md), error**


## CreateFineTuningJob

Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/fine-tuning)


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
    res, err := s.FineTuning.CreateFineTuningJob(ctx, shared.CreateFineTuningJobRequest{
        Hyperparameters: &shared.CreateFineTuningJobRequestHyperparameters{
            NEpochs: shared.CreateCreateFineTuningJobRequestHyperparametersNEpochsCreateFineTuningJobRequestHyperparametersNEpochs1(
            shared.CreateFineTuningJobRequestHyperparametersNEpochs1Auto,
            ),
        },
        Model: shared.CreateCreateFineTuningJobRequestModelCreateFineTuningJobRequestModel2(
        shared.CreateFineTuningJobRequestModel2Gpt35Turbo,
        ),
        TrainingFile: "file-abc123",
        ValidationFile: openaigosdk.String("file-abc123"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTuningJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.CreateFineTuningJobRequest](../../models/shared/createfinetuningjobrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateFineTuningJobResponse](../../models/operations/createfinetuningjobresponse.md), error**


## ListFineTuningEvents

Get status updates for a fine-tuning job.


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


    var fineTuningJobID string = "ft-AF1WoRqd3aJAHsqc9NY7iL8F"

    var after *string = "string"

    var limit *int64 = 896841

    ctx := context.Background()
    res, err := s.FineTuning.ListFineTuningEvents(ctx, fineTuningJobID, after, limit)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListFineTuningJobEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |                                                                     |
| `fineTuningJobID`                                                   | *string*                                                            | :heavy_check_mark:                                                  | The ID of the fine-tuning job to get events for.<br/>               | ft-AF1WoRqd3aJAHsqc9NY7iL8F                                         |
| `after`                                                             | **string*                                                           | :heavy_minus_sign:                                                  | Identifier for the last event from the previous pagination request. |                                                                     |
| `limit`                                                             | **int64*                                                            | :heavy_minus_sign:                                                  | Number of events to retrieve.                                       |                                                                     |


### Response

**[*operations.ListFineTuningEventsResponse](../../models/operations/listfinetuningeventsresponse.md), error**


## ListPaginatedFineTuningJobs

List your organization's fine-tuning jobs


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


    var after *string = "string"

    var limit *int64 = 385496

    ctx := context.Background()
    res, err := s.FineTuning.ListPaginatedFineTuningJobs(ctx, after, limit)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPaginatedFineTuningJobsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `after`                                                           | **string*                                                         | :heavy_minus_sign:                                                | Identifier for the last job from the previous pagination request. |
| `limit`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | Number of fine-tuning jobs to retrieve.                           |


### Response

**[*operations.ListPaginatedFineTuningJobsResponse](../../models/operations/listpaginatedfinetuningjobsresponse.md), error**


## RetrieveFineTuningJob

Get info about a fine-tuning job.

[Learn more about fine-tuning](/docs/guides/fine-tuning)


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


    var fineTuningJobID string = "ft-AF1WoRqd3aJAHsqc9NY7iL8F"

    ctx := context.Background()
    res, err := s.FineTuning.RetrieveFineTuningJob(ctx, fineTuningJobID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTuningJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `fineTuningJobID`                                     | *string*                                              | :heavy_check_mark:                                    | The ID of the fine-tuning job.<br/>                   | ft-AF1WoRqd3aJAHsqc9NY7iL8F                           |


### Response

**[*operations.RetrieveFineTuningJobResponse](../../models/operations/retrievefinetuningjobresponse.md), error**

