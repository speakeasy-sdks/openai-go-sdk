# OpenAI

## Overview

The OpenAI REST API

### Available Operations

* [CancelFineTune](#cancelfinetune) - Immediately cancel a fine-tune job.

* [CreateChatCompletion](#createchatcompletion) - Creates a model response for the given chat conversation.
* [CreateCompletion](#createcompletion) - Creates a completion for the provided prompt and parameters.
* [~~CreateEdit~~](#createedit) - Creates a new edit for the provided input, instruction, and parameters. :warning: **Deprecated**
* [CreateEmbedding](#createembedding) - Creates an embedding vector representing the input text.
* [CreateFile](#createfile) - Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit.

* [CreateFineTune](#createfinetune) - Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about Fine-tuning](/docs/guides/fine-tuning)

* [CreateImage](#createimage) - Creates an image given a prompt.
* [CreateImageEdit](#createimageedit) - Creates an edited or extended image given an original image and a prompt.
* [CreateImageVariation](#createimagevariation) - Creates a variation of a given image.
* [CreateModeration](#createmoderation) - Classifies if text violates OpenAI's Content Policy
* [CreateTranscription](#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](#createtranslation) - Translates audio into English.
* [DeleteFile](#deletefile) - Delete a file.
* [DeleteModel](#deletemodel) - Delete a fine-tuned model. You must have the Owner role in your organization.
* [DownloadFile](#downloadfile) - Returns the contents of the specified file
* [ListFiles](#listfiles) - Returns a list of files that belong to the user's organization.
* [ListFineTuneEvents](#listfinetuneevents) - Get fine-grained status updates for a fine-tune job.

* [ListFineTunes](#listfinetunes) - List your organization's fine-tuning jobs

* [ListModels](#listmodels) - Lists the currently available models, and provides basic information about each one such as the owner and availability.
* [RetrieveFile](#retrievefile) - Returns information about a specific file.
* [RetrieveFineTune](#retrievefinetune) - Gets info about the fine-tune job.

[Learn more about Fine-tuning](/docs/guides/fine-tuning)

* [RetrieveModel](#retrievemodel) - Retrieves a model instance, providing basic information about the model such as the owner and permissioning.

## CancelFineTune

Immediately cancel a fine-tune job.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CancelFineTune(ctx, operations.CancelFineTuneRequest{
        FineTuneID: "ft-AF1WoRqd3aJAHsqc9NY7iL8F",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CancelFineTuneRequest](../../models/operations/cancelfinetunerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CancelFineTuneResponse](../../models/operations/cancelfinetuneresponse.md), error**


## CreateChatCompletion

Creates a model response for the given chat conversation.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateChatCompletion(ctx, shared.CreateChatCompletionRequest{
        FrequencyPenalty: gpt.Float64(5488.14),
        FunctionCall: &shared.CreateChatCompletionRequestFunctionCall{},
        Functions: []shared.ChatCompletionFunctions{
            shared.ChatCompletionFunctions{
                Description: gpt.String("distinctio"),
                Name: "Stuart Stiedemann",
                Parameters: map[string]interface{}{
                    "error": "deserunt",
                    "suscipit": "iure",
                },
            },
            shared.ChatCompletionFunctions{
                Description: gpt.String("magnam"),
                Name: "Larry Windler",
                Parameters: map[string]interface{}{
                    "minus": "placeat",
                    "voluptatum": "iusto",
                },
            },
            shared.ChatCompletionFunctions{
                Description: gpt.String("excepturi"),
                Name: "Mrs. Sophie Smith MD",
                Parameters: map[string]interface{}{
                    "ipsam": "repellendus",
                },
            },
        },
        LogitBias: map[string]int64{
            "quo": 140350,
            "at": 870088,
            "maiores": 473608,
            "quod": 800911,
        },
        MaxTokens: gpt.Int64(461479),
        Messages: []shared.ChatCompletionRequestMessage{
            shared.ChatCompletionRequestMessage{
                Content: "porro",
                FunctionCall: &shared.ChatCompletionRequestMessageFunctionCall{
                    Arguments: "dolorum",
                    Name: "Antoinette Nikolaus",
                },
                Name: gpt.String("Irvin Rosenbaum III"),
                Role: shared.ChatCompletionRequestMessageRoleUser,
            },
            shared.ChatCompletionRequestMessage{
                Content: "modi",
                FunctionCall: &shared.ChatCompletionRequestMessageFunctionCall{
                    Arguments: "qui",
                    Name: "Jonathon Klocko",
                },
                Name: gpt.String("Dorothy Hane"),
                Role: shared.ChatCompletionRequestMessageRoleAssistant,
            },
            shared.ChatCompletionRequestMessage{
                Content: "dolor",
                FunctionCall: &shared.ChatCompletionRequestMessageFunctionCall{
                    Arguments: "natus",
                    Name: "May Turcotte",
                },
                Name: gpt.String("Sheryl Kertzmann"),
                Role: shared.ChatCompletionRequestMessageRoleSystem,
            },
        },
        Model: "gpt-3.5-turbo",
        N: gpt.Int64(1),
        PresencePenalty: gpt.Float64(9698.1),
        Stop: &shared.CreateChatCompletionRequestStop{},
        Stream: gpt.Bool(false),
        Temperature: gpt.Float64(1),
        TopP: gpt.Float64(1),
        User: gpt.String("user-1234"),
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


## CreateCompletion

Creates a completion for the provided prompt and parameters.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateCompletion(ctx, shared.CreateCompletionRequest{
        BestOf: gpt.Int64(666767),
        Echo: gpt.Bool(false),
        FrequencyPenalty: gpt.Float64(6531.4),
        LogitBias: map[string]int64{
            "dolores": 210382,
            "corporis": 128926,
            "nobis": 315428,
        },
        Logprobs: gpt.Int64(607831),
        MaxTokens: gpt.Int64(16),
        Model: "minima",
        N: gpt.Int64(1),
        PresencePenalty: gpt.Float64(5701.97),
        Prompt: shared.CreateCompletionRequestPrompt{},
        Stop: &shared.CreateCompletionRequestStop{},
        Stream: gpt.Bool(false),
        Suffix: gpt.String("test."),
        Temperature: gpt.Float64(1),
        TopP: gpt.Float64(1),
        User: gpt.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCompletionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateCompletionRequest](../../models/shared/createcompletionrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateCompletionResponse](../../models/operations/createcompletionresponse.md), error**


## ~~CreateEdit~~

Creates a new edit for the provided input, instruction, and parameters.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateEdit(ctx, shared.CreateEditRequest{
        Input: gpt.String("What day of the wek is it?"),
        Instruction: "Fix the spelling mistakes.",
        Model: "text-davinci-edit-001",
        N: gpt.Int64(1),
        Temperature: gpt.Float64(1),
        TopP: gpt.Float64(1),
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.CreateEditRequest](../../models/shared/createeditrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.CreateEditResponse](../../models/operations/createeditresponse.md), error**


## CreateEmbedding

Creates an embedding vector representing the input text.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateEmbedding(ctx, shared.CreateEmbeddingRequest{
        Input: shared.CreateEmbeddingRequestInput{},
        Model: "text-embedding-ada-002",
        User: gpt.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateEmbeddingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateEmbeddingRequest](../../models/shared/createembeddingrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateEmbeddingResponse](../../models/operations/createembeddingresponse.md), error**


## CreateFile

Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateFile(ctx, shared.CreateFileRequest{
        File: shared.CreateFileRequestFile{
            Content: []byte("culpa"),
            File: "doloribus",
        },
        Purpose: "sapiente",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAIFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.CreateFileRequest](../../models/shared/createfilerequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.CreateFileResponse](../../models/operations/createfileresponse.md), error**


## CreateFineTune

Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about Fine-tuning](/docs/guides/fine-tuning)


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateFineTune(ctx, shared.CreateFineTuneRequest{
        BatchSize: gpt.Int64(102044),
        ClassificationBetas: []float64{
            2088.76,
            6350.59,
            1613.09,
        },
        ClassificationNClasses: gpt.Int64(995300),
        ClassificationPositiveClass: gpt.String("mollitia"),
        ComputeClassificationMetrics: gpt.Bool(false),
        LearningRateMultiplier: gpt.Float64(5818.5),
        Model: gpt.String("curie"),
        NEpochs: gpt.Int64(414369),
        PromptLossWeight: gpt.Float64(4663.11),
        Suffix: gpt.String("molestiae"),
        TrainingFile: "file-ajSREls59WBbvgSzJSVWxMCB",
        ValidationFile: gpt.String("file-XjSREls59WBbvgSzJSVWxMCa"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateFineTuneRequest](../../models/shared/createfinetunerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateFineTuneResponse](../../models/operations/createfinetuneresponse.md), error**


## CreateImage

Creates an image given a prompt.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateImage(ctx, shared.CreateImageRequest{
        N: gpt.Int64(1),
        Prompt: "A cute baby sea otter",
        ResponseFormat: shared.CreateImageRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: gpt.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.CreateImageRequest](../../models/shared/createimagerequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.CreateImageResponse](../../models/operations/createimageresponse.md), error**


## CreateImageEdit

Creates an edited or extended image given an original image and a prompt.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateImageEdit(ctx, shared.CreateImageEditRequest2{
        Image: shared.CreateImageEditRequestImage{
            Content: []byte("velit"),
            Image: "error",
        },
        Mask: &shared.CreateImageEditRequestMask{
            Content: []byte("quia"),
            Mask: "quis",
        },
        N: gpt.Int64(1),
        Prompt: "A cute baby sea otter wearing a beret",
        ResponseFormat: shared.CreateImageEditRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageEditRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: gpt.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateImageEditRequest2](../../models/shared/createimageeditrequest2.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateImageEditResponse](../../models/operations/createimageeditresponse.md), error**


## CreateImageVariation

Creates a variation of a given image.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateImageVariation(ctx, shared.CreateImageVariationRequest2{
        Image: shared.CreateImageVariationRequestImage{
            Content: []byte("vitae"),
            Image: "laborum",
        },
        N: gpt.Int64(1),
        ResponseFormat: shared.CreateImageVariationRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageVariationRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: gpt.String("user-1234"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.CreateImageVariationRequest2](../../models/shared/createimagevariationrequest2.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateImageVariationResponse](../../models/operations/createimagevariationresponse.md), error**


## CreateModeration

Classifies if text violates OpenAI's Content Policy

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateModeration(ctx, shared.CreateModerationRequest{
        Input: shared.CreateModerationRequestInput{},
        Model: shared.CreateModerationRequestModel2TextModerationStable.ToPointer(),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CreateModerationRequest](../../models/shared/createmoderationrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateModerationResponse](../../models/operations/createmoderationresponse.md), error**


## CreateTranscription

Transcribes audio into the input language.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateTranscription(ctx, shared.CreateTranscriptionRequest1{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("enim"),
            File: "odit",
        },
        Language: gpt.String("quo"),
        Model: "whisper-1",
        Prompt: gpt.String("tenetur"),
        ResponseFormat: shared.CreateTranscriptionRequestResponseFormatText.ToPointer(),
        Temperature: gpt.Float64(6625.27),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CreateTranscriptionRequest1](../../models/shared/createtranscriptionrequest1.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateTranslation(ctx, shared.CreateTranslationRequest{
        File: shared.CreateTranslationRequestFile{
            Content: []byte("possimus"),
            File: "aut",
        },
        Model: "whisper-1",
        Prompt: gpt.String("error"),
        ResponseFormat: gpt.String("temporibus"),
        Temperature: gpt.Float64(6736.6),
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


## DeleteFile

Delete a file.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.DeleteFile(ctx, operations.DeleteFileRequest{
        FileID: "quasi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteFileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteFileRequest](../../models/operations/deletefilerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.DeleteFileResponse](../../models/operations/deletefileresponse.md), error**


## DeleteModel

Delete a fine-tuned model. You must have the Owner role in your organization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.DeleteModel(ctx, operations.DeleteModelRequest{
        Model: "curie:ft-acmeco-2021-03-03-21-44-20",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteModelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteModelRequest](../../models/operations/deletemodelrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DeleteModelResponse](../../models/operations/deletemodelresponse.md), error**


## DownloadFile

Returns the contents of the specified file

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.DownloadFile(ctx, operations.DownloadFileRequest{
        FileID: "reiciendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DownloadFile200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DownloadFileRequest](../../models/operations/downloadfilerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DownloadFileResponse](../../models/operations/downloadfileresponse.md), error**


## ListFiles

Returns a list of files that belong to the user's organization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.ListFiles(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListFilesResponse](../../models/operations/listfilesresponse.md), error**


## ListFineTuneEvents

Get fine-grained status updates for a fine-tune job.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.ListFineTuneEvents(ctx, operations.ListFineTuneEventsRequest{
        FineTuneID: "ft-AF1WoRqd3aJAHsqc9NY7iL8F",
        Stream: gpt.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListFineTuneEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListFineTuneEventsRequest](../../models/operations/listfinetuneeventsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListFineTuneEventsResponse](../../models/operations/listfinetuneeventsresponse.md), error**


## ListFineTunes

List your organization's fine-tuning jobs


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.ListFineTunes(ctx)
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

**[*operations.ListFineTunesResponse](../../models/operations/listfinetunesresponse.md), error**


## ListModels

Lists the currently available models, and provides basic information about each one such as the owner and availability.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.ListModels(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListModelsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListModelsResponse](../../models/operations/listmodelsresponse.md), error**


## RetrieveFile

Returns information about a specific file.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.RetrieveFile(ctx, operations.RetrieveFileRequest{
        FileID: "voluptatibus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAIFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.RetrieveFileRequest](../../models/operations/retrievefilerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.RetrieveFileResponse](../../models/operations/retrievefileresponse.md), error**


## RetrieveFineTune

Gets info about the fine-tune job.

[Learn more about Fine-tuning](/docs/guides/fine-tuning)


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.RetrieveFineTune(ctx, operations.RetrieveFineTuneRequest{
        FineTuneID: "ft-AF1WoRqd3aJAHsqc9NY7iL8F",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RetrieveFineTuneRequest](../../models/operations/retrievefinetunerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RetrieveFineTuneResponse](../../models/operations/retrievefinetuneresponse.md), error**


## RetrieveModel

Retrieves a model instance, providing basic information about the model such as the owner and permissioning.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.RetrieveModel(ctx, operations.RetrieveModelRequest{
        Model: "text-davinci-001",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Model != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.RetrieveModelRequest](../../models/operations/retrievemodelrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.RetrieveModelResponse](../../models/operations/retrievemodelresponse.md), error**

