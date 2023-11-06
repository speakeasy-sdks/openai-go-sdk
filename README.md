# OpenAPI Go SDK

![OpenAI_Logo_Black](https://user-images.githubusercontent.com/6267663/220744241-48f469af-40b6-4d7f-ab48-8426b30189f0.svg#gh-light-mode-only)
![OpenAI_Logo_White](https://user-images.githubusercontent.com/6267663/220744513-66c99d0e-ed91-4577-982f-e7128d35ce95.svg#gh-dark-mode-only)

<div align="center">
   <p><strong>This is an unofficial SDK for the OpenAI API.</strong>  The OpenAI API can be applied to virtually any task that involves understanding or generating natural language or code. We offer a spectrum of models with different levels of power suitable for different tasks, as well as the ability to fine-tune your own custom models. These models can be used for everything from content generation to semantic search and classification.</p>
   <a href="https://github.com/speakeasy-sdks/openai-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/openai-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://platform.openai.com/docs/introduction"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=2ca47c&style=for-the-badge" /></a>
</div> 

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/openai-go-sdk
```
<!-- End SDK Installation -->

## Authentication

The OpenAI API uses API keys for authentication. Visit your API Keys page to retrieve the API key you'll use in your requests.

**Remember that your API key is a secret!** Do not share it with others or expose it in any client-side code (browsers, apps). Production requests must be routed through your own backend server where your API key can be securely loaded from an environment variable or key management service.

All API requests should include your API key in an Authorization HTTP header as follows:

```bash
Authorization: Bearer YOUR_API_KEY
```

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
            File: "string",
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Audio](docs/sdks/audio/README.md)

* [CreateTranscription](docs/sdks/audio/README.md#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](docs/sdks/audio/README.md#createtranslation) - Translates audio into English.

### [Chat](docs/sdks/chat/README.md)

* [CreateChatCompletion](docs/sdks/chat/README.md#createchatcompletion) - Creates a model response for the given chat conversation.

### [Completions](docs/sdks/completions/README.md)

* [CreateCompletion](docs/sdks/completions/README.md#createcompletion) - Creates a completion for the provided prompt and parameters.

### [Edits](docs/sdks/edits/README.md)

* [~~CreateEdit~~](docs/sdks/edits/README.md#createedit) - Creates a new edit for the provided input, instruction, and parameters. :warning: **Deprecated**

### [Embeddings](docs/sdks/embeddings/README.md)

* [CreateEmbedding](docs/sdks/embeddings/README.md#createembedding) - Creates an embedding vector representing the input text.

### [Files](docs/sdks/files/README.md)

* [CreateFile](docs/sdks/files/README.md#createfile) - Upload a file that can be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please [contact us](https://help.openai.com/) if you need to increase the storage limit.

* [DeleteFile](docs/sdks/files/README.md#deletefile) - Delete a file.
* [DownloadFile](docs/sdks/files/README.md#downloadfile) - Returns the contents of the specified file.
* [ListFiles](docs/sdks/files/README.md#listfiles) - Returns a list of files that belong to the user's organization.
* [RetrieveFile](docs/sdks/files/README.md#retrievefile) - Returns information about a specific file.

### [FineTunes](docs/sdks/finetunes/README.md)

* [~~CancelFineTune~~](docs/sdks/finetunes/README.md#cancelfinetune) - Immediately cancel a fine-tune job.
 :warning: **Deprecated**
* [~~CreateFineTune~~](docs/sdks/finetunes/README.md#createfinetune) - Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/legacy-fine-tuning)
 :warning: **Deprecated**
* [~~ListFineTuneEvents~~](docs/sdks/finetunes/README.md#listfinetuneevents) - Get fine-grained status updates for a fine-tune job.
 :warning: **Deprecated**
* [~~ListFineTunes~~](docs/sdks/finetunes/README.md#listfinetunes) - List your organization's fine-tuning jobs
 :warning: **Deprecated**
* [~~RetrieveFineTune~~](docs/sdks/finetunes/README.md#retrievefinetune) - Gets info about the fine-tune job.

[Learn more about fine-tuning](/docs/guides/legacy-fine-tuning)
 :warning: **Deprecated**

### [FineTuning](docs/sdks/finetuning/README.md)

* [CancelFineTuningJob](docs/sdks/finetuning/README.md#cancelfinetuningjob) - Immediately cancel a fine-tune job.

* [CreateFineTuningJob](docs/sdks/finetuning/README.md#createfinetuningjob) - Creates a job that fine-tunes a specified model from a given dataset.

Response includes details of the enqueued job including job status and the name of the fine-tuned models once complete.

[Learn more about fine-tuning](/docs/guides/fine-tuning)

* [ListFineTuningEvents](docs/sdks/finetuning/README.md#listfinetuningevents) - Get status updates for a fine-tuning job.

* [ListPaginatedFineTuningJobs](docs/sdks/finetuning/README.md#listpaginatedfinetuningjobs) - List your organization's fine-tuning jobs

* [RetrieveFineTuningJob](docs/sdks/finetuning/README.md#retrievefinetuningjob) - Get info about a fine-tuning job.

[Learn more about fine-tuning](/docs/guides/fine-tuning)


### [Images](docs/sdks/images/README.md)

* [CreateImage](docs/sdks/images/README.md#createimage) - Creates an image given a prompt.
* [CreateImageEdit](docs/sdks/images/README.md#createimageedit) - Creates an edited or extended image given an original image and a prompt.
* [CreateImageVariation](docs/sdks/images/README.md#createimagevariation) - Creates a variation of a given image.

### [Models](docs/sdks/models/README.md)

* [DeleteModel](docs/sdks/models/README.md#deletemodel) - Delete a fine-tuned model. You must have the Owner role in your organization to delete a model.
* [ListModels](docs/sdks/models/README.md#listmodels) - Lists the currently available models, and provides basic information about each one such as the owner and availability.
* [RetrieveModel](docs/sdks/models/README.md#retrievemodel) - Retrieves a model instance, providing basic information about the model such as the owner and permissioning.

### [Moderations](docs/sdks/moderations/README.md)

* [CreateModeration](docs/sdks/moderations/README.md#createmoderation) - Classifies if text violates OpenAI's Content Policy
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->



<!-- End Dev Containers -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.openai.com/v1` | None |

For example:


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
        openaigosdk.WithServerIndex(0),
    )

    ctx := context.Background()
    res, err := s.Audio.CreateTranscription(ctx, shared.CreateTranscriptionRequest{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("\#BbTW'zX9"),
            File: "string",
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


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


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
        openaigosdk.WithServerURL("https://api.openai.com/v1"),
    )

    ctx := context.Background()
    res, err := s.Audio.CreateTranscription(ctx, shared.CreateTranscriptionRequest{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("\#BbTW'zX9"),
            File: "string",
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
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
