# OpenAI

## Overview

The OpenAI REST API

### Available Operations

* [CancelFineTune](#cancelfinetune) - Immediately cancel a fine-tune job.

* [~~CreateAnswer~~](#createanswer) - Answers the specified question using the provided documents and examples.

The endpoint first [searches](/docs/api-reference/searches) over provided documents or files to find relevant context. The relevant context is combined with the provided examples and question to create the prompt for [completion](/docs/api-reference/completions).
 :warning: **Deprecated**
* [CreateChatCompletion](#createchatcompletion) - Creates a model response for the given chat conversation.
* [~~CreateClassification~~](#createclassification) - Classifies the specified `query` using provided examples.

The endpoint first [searches](/docs/api-reference/searches) over the labeled examples
to select the ones most relevant for the particular query. Then, the relevant examples
are combined with the query to construct a prompt to produce the final label via the
[completions](/docs/api-reference/completions) endpoint.

Labeled examples can be provided via an uploaded `file`, or explicitly listed in the
request using the `examples` parameter for quick tests and small scale use cases.
 :warning: **Deprecated**
* [CreateCompletion](#createcompletion) - Creates a completion for the provided prompt and parameters.
* [CreateEdit](#createedit) - Creates a new edit for the provided input, instruction, and parameters.
* [CreateEmbedding](#createembedding) - Creates an embedding vector representing the input text.
* [CreateFile](#createfile) - Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit.

* [CreateFineTune](#createfinetune) - Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about Fine-tuning](/docs/guides/fine-tuning)

* [CreateImage](#createimage) - Creates an image given a prompt.
* [CreateImageEdit](#createimageedit) - Creates an edited or extended image given an original image and a prompt.
* [CreateImageVariation](#createimagevariation) - Creates a variation of a given image.
* [CreateModeration](#createmoderation) - Classifies if text violates OpenAI's Content Policy
* [~~CreateSearch~~](#createsearch) - The search endpoint computes similarity scores between provided query and documents. Documents can be passed directly to the API if there are no more than 200 of them.

To go beyond the 200 document limit, documents can be processed offline and then used for efficient retrieval at query time. When `file` is set, the search endpoint searches over all the documents in the given file and returns up to the `max_rerank` number of documents. These documents will be returned along with their search scores.

The similarity score is a positive score that usually ranges from 0 to 300 (but can sometimes go higher), where a score above 200 usually means the document is semantically similar to the query.
 :warning: **Deprecated**
* [CreateTranscription](#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](#createtranslation) - Translates audio into into English.
* [DeleteFile](#deletefile) - Delete a file.
* [DeleteModel](#deletemodel) - Delete a fine-tuned model. You must have the Owner role in your organization.
* [DownloadFile](#downloadfile) - Returns the contents of the specified file
* [~~ListEngines~~](#listengines) - Lists the currently available (non-finetuned) models, and provides basic information about each one such as the owner and availability. :warning: **Deprecated**
* [ListFiles](#listfiles) - Returns a list of files that belong to the user's organization.
* [ListFineTuneEvents](#listfinetuneevents) - Get fine-grained status updates for a fine-tune job.

* [ListFineTunes](#listfinetunes) - List your organization's fine-tuning jobs

* [ListModels](#listmodels) - Lists the currently available models, and provides basic information about each one such as the owner and availability.
* [~~RetrieveEngine~~](#retrieveengine) - Retrieves a model instance, providing basic information about it such as the owner and availability. :warning: **Deprecated**
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
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


## ~~CreateAnswer~~

Answers the specified question using the provided documents and examples.

The endpoint first [searches](/docs/api-reference/searches) over provided documents or files to find relevant context. The relevant context is combined with the provided examples and question to create the prompt for [completion](/docs/api-reference/completions).


> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateAnswer(ctx, shared.CreateAnswerRequest{
        Documents: []string{
            "provident",
            "distinctio",
            "quibusdam",
        },
        Examples: [][]string{
            []string{
                "corrupti",
                "illum",
                "vel",
                "error",
            },
            []string{
                "suscipit",
                "iure",
                "magnam",
            },
            []string{
                "ipsa",
                "delectus",
                "tempora",
                "suscipit",
            },
        },
        ExamplesContext: "Ottawa, Canada's capital, is located in the east of southern Ontario, near the city of Montréal and the U.S. border.",
        Expand: []interface{}{
            "minus",
            "placeat",
        },
        File: gpt.String("voluptatum"),
        LogitBias: gpt.String("iusto"),
        Logprobs: gpt.Int64(568045),
        MaxRerank: gpt.Int64(392785),
        MaxTokens: gpt.Int64(925597),
        Model: "temporibus",
        N: gpt.Int64(71036),
        Question: "What is the capital of Japan?",
        ReturnMetadata: gpt.String("quis"),
        ReturnPrompt: gpt.Bool(false),
        SearchModel: gpt.String("veritatis"),
        Stop: &shared.CreateAnswerRequestStop{},
        Temperature: gpt.Float64(6481.72),
        User: gpt.String("perferendis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAnswerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateAnswerRequest](../../models/shared/createanswerrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateAnswerResponse](../../models/operations/createanswerresponse.md), error**


## CreateChatCompletion

Creates a model response for the given chat conversation.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateChatCompletion(ctx, shared.CreateChatCompletionRequest{
        FrequencyPenalty: gpt.Float64(3682.41),
        FunctionCall: &shared.CreateChatCompletionRequestFunctionCall{},
        Functions: []shared.ChatCompletionFunctions{
            shared.ChatCompletionFunctions{
                Description: gpt.String("sapiente"),
                Name: "Fred Strosin",
                Parameters: map[string]interface{}{
                    "quod": "quod",
                    "esse": "totam",
                },
            },
            shared.ChatCompletionFunctions{
                Description: gpt.String("porro"),
                Name: "Samuel Reichel",
                Parameters: map[string]interface{}{
                    "deleniti": "hic",
                },
            },
            shared.ChatCompletionFunctions{
                Description: gpt.String("optio"),
                Name: "Jack Johns",
                Parameters: map[string]interface{}{
                    "impedit": "cum",
                },
            },
            shared.ChatCompletionFunctions{
                Description: gpt.String("esse"),
                Name: "Mrs. Miriam Collier",
                Parameters: map[string]interface{}{
                    "iste": "dolor",
                },
            },
        },
        LogitBias: &shared.CreateChatCompletionRequestLogitBias{},
        MaxTokens: gpt.Int64(616934),
        Messages: []shared.ChatCompletionRequestMessage{
            shared.ChatCompletionRequestMessage{
                Content: gpt.String("hic"),
                FunctionCall: &shared.ChatCompletionRequestMessageFunctionCall{
                    Arguments: gpt.String("saepe"),
                    Name: gpt.String("Harvey Hessel"),
                },
                Name: gpt.String("Dr. Rickey Boyle"),
                Role: shared.ChatCompletionRequestMessageRoleAssistant,
            },
            shared.ChatCompletionRequestMessage{
                Content: gpt.String("laborum"),
                FunctionCall: &shared.ChatCompletionRequestMessageFunctionCall{
                    Arguments: gpt.String("dolores"),
                    Name: gpt.String("Stacy Champlin"),
                },
                Name: gpt.String("Corey Hane III"),
                Role: shared.ChatCompletionRequestMessageRoleAssistant,
            },
        },
        Model: "doloribus",
        N: gpt.Int64(1),
        PresencePenalty: gpt.Float64(9589.5),
        Stop: &shared.CreateChatCompletionRequestStop{},
        Stream: gpt.Bool(false),
        Temperature: gpt.Float64(1),
        TopP: gpt.Float64(1),
        User: gpt.String("architecto"),
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


## ~~CreateClassification~~

Classifies the specified `query` using provided examples.

The endpoint first [searches](/docs/api-reference/searches) over the labeled examples
to select the ones most relevant for the particular query. Then, the relevant examples
are combined with the query to construct a prompt to produce the final label via the
[completions](/docs/api-reference/completions) endpoint.

Labeled examples can be provided via an uploaded `file`, or explicitly listed in the
request using the `examples` parameter for quick tests and small scale use cases.


> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateClassification(ctx, shared.CreateClassificationRequest{
        Examples: [][]string{
            []string{
                "culpa",
            },
            []string{
                "repellat",
            },
            []string{
                "occaecati",
                "numquam",
                "commodi",
            },
        },
        Expand: gpt.String("quam"),
        File: gpt.String("molestiae"),
        Labels: []string{
            "error",
        },
        LogitBias: gpt.String("quia"),
        Logprobs: gpt.String("quis"),
        MaxExamples: gpt.Int64(110375),
        Model: "laborum",
        Query: "The plot is not very attractive.",
        ReturnMetadata: gpt.String("animi"),
        ReturnPrompt: gpt.String("enim"),
        SearchModel: gpt.String("odit"),
        Temperature: gpt.Float64(0),
        User: gpt.String("quo"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateClassificationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CreateClassificationRequest](../../models/shared/createclassificationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateClassificationResponse](../../models/operations/createclassificationresponse.md), error**


## CreateCompletion

Creates a completion for the provided prompt and parameters.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateCompletion(ctx, shared.CreateCompletionRequest{
        BestOf: gpt.Int64(196582),
        Echo: gpt.Bool(false),
        FrequencyPenalty: gpt.Float64(9495.72),
        LogitBias: &shared.CreateCompletionRequestLogitBias{},
        Logprobs: gpt.Int64(368725),
        MaxTokens: gpt.Int64(16),
        Model: "id",
        N: gpt.Int64(1),
        PresencePenalty: gpt.Float64(8209.94),
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


## CreateEdit

Creates a new edit for the provided input, instruction, and parameters.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateEdit(ctx, shared.CreateEditRequest{
        Input: gpt.String("What day of the wek is it?"),
        Instruction: "Fix the spelling mistakes.",
        Model: "aut",
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateEmbedding(ctx, shared.CreateEmbeddingRequest{
        Input: shared.CreateEmbeddingRequestInput{},
        Model: "quasi",
        User: gpt.String("error"),
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateFile(ctx, shared.CreateFileRequest{
        File: shared.CreateFileRequestFile{
            Content: []byte("temporibus"),
            File: "laborum",
        },
        Purpose: "quasi",
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateFineTune(ctx, shared.CreateFineTuneRequest{
        BatchSize: gpt.Int64(971945),
        ClassificationBetas: []float64{
            8781.94,
            4686.51,
            5096.24,
            9767.62,
        },
        ClassificationNClasses: gpt.Int64(55714),
        ClassificationPositiveClass: gpt.String("omnis"),
        ComputeClassificationMetrics: gpt.Bool(false),
        LearningRateMultiplier: gpt.Float64(4511.59),
        Model: gpt.String("cum"),
        NEpochs: gpt.Int64(19987),
        PromptLossWeight: gpt.Float64(391.87),
        Suffix: gpt.String("reprehenderit"),
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateImage(ctx, shared.CreateImageRequest{
        N: gpt.Int64(1),
        Prompt: "A cute baby sea otter",
        ResponseFormat: shared.CreateImageRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: gpt.String("ut"),
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateImageEdit(ctx, shared.CreateImageEditRequest{
        Image: shared.CreateImageEditRequestImage{
            Content: []byte("maiores"),
            Image: "dicta",
        },
        Mask: &shared.CreateImageEditRequestMask{
            Content: []byte("corporis"),
            Mask: "dolore",
        },
        N: gpt.String("iusto"),
        Prompt: "A cute baby sea otter wearing a beret",
        ResponseFormat: gpt.String("dicta"),
        Size: gpt.String("harum"),
        User: gpt.String("enim"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateImageEditRequest](../../models/shared/createimageeditrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateImageVariation(ctx, shared.CreateImageVariationRequest{
        Image: shared.CreateImageVariationRequestImage{
            Content: []byte("accusamus"),
            Image: "commodi",
        },
        N: gpt.String("repudiandae"),
        ResponseFormat: gpt.String("quae"),
        Size: gpt.String("ipsum"),
        User: gpt.String("quidem"),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CreateImageVariationRequest](../../models/shared/createimagevariationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateModeration(ctx, shared.CreateModerationRequest{
        Input: shared.CreateModerationRequestInput{},
        Model: gpt.String("text-moderation-stable"),
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


## ~~CreateSearch~~

The search endpoint computes similarity scores between provided query and documents. Documents can be passed directly to the API if there are no more than 200 of them.

To go beyond the 200 document limit, documents can be processed offline and then used for efficient retrieval at query time. When `file` is set, the search endpoint searches over all the documents in the given file and returns up to the `max_rerank` number of documents. These documents will be returned along with their search scores.

The similarity score is a positive score that usually ranges from 0 to 300 (but can sometimes go higher), where a score above 200 usually means the document is semantically similar to the query.


> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateSearch(ctx, operations.CreateSearchRequest{
        CreateSearchRequest: shared.CreateSearchRequest{
            Documents: []string{
                "excepturi",
                "pariatur",
                "modi",
            },
            File: gpt.String("praesentium"),
            MaxRerank: gpt.Int64(523248),
            Query: "the president",
            ReturnMetadata: gpt.Bool(false),
            User: gpt.String("voluptates"),
        },
        EngineID: "davinci",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateSearchRequest](../../models/operations/createsearchrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateSearchResponse](../../models/operations/createsearchresponse.md), error**


## CreateTranscription

Transcribes audio into the input language.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateTranscription(ctx, shared.CreateTranscriptionRequest{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("quasi"),
            File: "repudiandae",
        },
        Language: gpt.String("sint"),
        Model: "veritatis",
        Prompt: gpt.String("itaque"),
        ResponseFormat: gpt.String("incidunt"),
        Temperature: gpt.Float64(3185.69),
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

Translates audio into into English.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.CreateTranslation(ctx, shared.CreateTranslationRequest{
        File: shared.CreateTranslationRequestFile{
            Content: []byte("consequatur"),
            File: "est",
        },
        Model: "quibusdam",
        Prompt: gpt.String("explicabo"),
        ResponseFormat: gpt.String("deserunt"),
        Temperature: gpt.Float64(7163.27),
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.DeleteFile(ctx, operations.DeleteFileRequest{
        FileID: "quibusdam",
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.DownloadFile(ctx, operations.DownloadFileRequest{
        FileID: "labore",
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


## ~~ListEngines~~

Lists the currently available (non-finetuned) models, and provides basic information about each one such as the owner and availability.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

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
    res, err := s.OpenAI.ListEngines(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListEnginesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ListEnginesResponse](../../models/operations/listenginesresponse.md), error**


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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
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


## ~~RetrieveEngine~~

Retrieves a model instance, providing basic information about it such as the owner and availability.

> :warning: **DEPRECATED**: this method will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.RetrieveEngine(ctx, operations.RetrieveEngineRequest{
        EngineID: "davinci",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Engine != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RetrieveEngineRequest](../../models/operations/retrieveenginerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.RetrieveEngineResponse](../../models/operations/retrieveengineresponse.md), error**


## RetrieveFile

Returns information about a specific file.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/openai-go-sdk"
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
)

func main() {
    s := gpt.New()

    ctx := context.Background()
    res, err := s.OpenAI.RetrieveFile(ctx, operations.RetrieveFileRequest{
        FileID: "modi",
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
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
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
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

