# OpenAI

## Overview

The OpenAI REST API

### Available Operations

* [CancelFineTune](#cancelfinetune) - Immediately cancel a fine-tune job.

* [~~CreateAnswer~~](#createanswer) - Answers the specified question using the provided documents and examples.

The endpoint first [searches](/docs/api-reference/searches) over provided documents or files to find relevant context. The relevant context is combined with the provided examples and question to create the prompt for [completion](/docs/api-reference/completions).
 :warning: **Deprecated**
* [CreateChatCompletion](#createchatcompletion) - Creates a completion for the chat message
* [~~CreateClassification~~](#createclassification) - Classifies the specified `query` using provided examples.

The endpoint first [searches](/docs/api-reference/searches) over the labeled examples
to select the ones most relevant for the particular query. Then, the relevant examples
are combined with the query to construct a prompt to produce the final label via the
[completions](/docs/api-reference/completions) endpoint.

Labeled examples can be provided via an uploaded `file`, or explicitly listed in the
request using the `examples` parameter for quick tests and small scale use cases.
 :warning: **Deprecated**
* [CreateCompletion](#createcompletion) - Creates a completion for the provided prompt and parameters
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

## CreateChatCompletion

Creates a completion for the chat message

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
        LogitBias: map[string]interface{}{
            "sapiente": "quo",
            "odit": "at",
            "at": "maiores",
            "molestiae": "quod",
        },
        MaxTokens: gpt.Int64(800911),
        Messages: []shared.ChatCompletionRequestMessage{
            shared.ChatCompletionRequestMessage{
                Content: "totam",
                Name: gpt.String("Omar Carroll"),
                Role: shared.ChatCompletionRequestMessageRoleUser,
            },
            shared.ChatCompletionRequestMessage{
                Content: "fugit",
                Name: gpt.String("Irvin Rosenbaum III"),
                Role: shared.ChatCompletionRequestMessageRoleUser,
            },
        },
        Model: "modi",
        N: gpt.Int64(1),
        PresencePenalty: gpt.Float64(1863.32),
        Stop: &shared.CreateChatCompletionRequestStop{},
        Stream: gpt.Bool(false),
        Temperature: gpt.Float64(1),
        TopP: gpt.Float64(1),
        User: gpt.String("impedit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateChatCompletionResponse != nil {
        // handle response
    }
}
```

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
                "ipsum",
                "excepturi",
            },
            []string{
                "perferendis",
            },
            []string{
                "natus",
                "sed",
            },
        },
        Expand: gpt.String("iste"),
        File: gpt.String("dolor"),
        Labels: []string{
            "laboriosam",
            "hic",
            "saepe",
        },
        LogitBias: gpt.String("fuga"),
        Logprobs: gpt.String("in"),
        MaxExamples: gpt.Int64(359508),
        Model: "iste",
        Query: "The plot is not very attractive.",
        ReturnMetadata: gpt.String("iure"),
        ReturnPrompt: gpt.String("saepe"),
        SearchModel: gpt.String("quidem"),
        Temperature: gpt.Float64(0),
        User: gpt.String("architecto"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateClassificationResponse != nil {
        // handle response
    }
}
```

## CreateCompletion

Creates a completion for the provided prompt and parameters

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
        BestOf: gpt.Int64(60225),
        Echo: gpt.Bool(false),
        FrequencyPenalty: gpt.Float64(9698.1),
        LogitBias: map[string]interface{}{
            "mollitia": "laborum",
            "dolores": "dolorem",
            "corporis": "explicabo",
        },
        Logprobs: gpt.Int64(750686),
        MaxTokens: gpt.Int64(16),
        Model: "enim",
        N: gpt.Int64(1),
        PresencePenalty: gpt.Float64(6078.31),
        Prompt: &shared.CreateCompletionRequestPrompt{},
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
        Model: "nemo",
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
        Model: "minima",
        User: gpt.String("excepturi"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateEmbeddingResponse != nil {
        // handle response
    }
}
```

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
            Content: []byte("accusantium"),
            File: "iure",
        },
        Purpose: "culpa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAIFile != nil {
        // handle response
    }
}
```

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
        BatchSize: gpt.Int64(988374),
        ClassificationBetas: []float64{
            1020.44,
            6527.9,
            2088.76,
            6350.59,
        },
        ClassificationNClasses: gpt.Int64(161309),
        ClassificationPositiveClass: gpt.String("repellat"),
        ComputeClassificationMetrics: gpt.Bool(false),
        LearningRateMultiplier: gpt.Float64(6531.08),
        Model: gpt.String("occaecati"),
        NEpochs: gpt.Int64(253291),
        PromptLossWeight: gpt.Float64(4143.69),
        Suffix: gpt.String("quam"),
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
        User: gpt.String("molestiae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagesResponse != nil {
        // handle response
    }
}
```

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
            Content: []byte("velit"),
            Image: "error",
        },
        Mask: &shared.CreateImageEditRequestMask{
            Content: []byte("quia"),
            Mask: "quis",
        },
        N: gpt.String("vitae"),
        Prompt: "A cute baby sea otter wearing a beret",
        ResponseFormat: gpt.String("laborum"),
        Size: gpt.String("animi"),
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
            Content: []byte("odit"),
            Image: "quo",
        },
        N: gpt.String("sequi"),
        ResponseFormat: gpt.String("tenetur"),
        Size: gpt.String("ipsam"),
        User: gpt.String("id"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImagesResponse != nil {
        // handle response
    }
}
```

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
                "aut",
                "quasi",
                "error",
                "temporibus",
            },
            File: gpt.String("laborum"),
            MaxRerank: gpt.Int64(96098),
            Query: "the president",
            ReturnMetadata: gpt.Bool(false),
            User: gpt.String("reiciendis"),
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
            Content: []byte("voluptatibus"),
            File: "vero",
        },
        Language: gpt.String("nihil"),
        Model: "praesentium",
        Prompt: gpt.String("voluptatibus"),
        ResponseFormat: gpt.String("ipsa"),
        Temperature: gpt.Float64(6048.46),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTranscriptionResponse != nil {
        // handle response
    }
}
```

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
            Content: []byte("voluptate"),
            File: "cum",
        },
        Model: "perferendis",
        Prompt: gpt.String("doloremque"),
        ResponseFormat: gpt.String("reprehenderit"),
        Temperature: gpt.Float64(2828.07),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTranslationResponse != nil {
        // handle response
    }
}
```

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
        FileID: "maiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteFileResponse != nil {
        // handle response
    }
}
```

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
        FileID: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DownloadFile200ApplicationJSONString != nil {
        // handle response
    }
}
```

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
        FileID: "corporis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAIFile != nil {
        // handle response
    }
}
```

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
