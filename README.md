# Go SDK

![OpenAI_Logo_Black](https://user-images.githubusercontent.com/6267663/220744241-48f469af-40b6-4d7f-ab48-8426b30189f0.svg#gh-light-mode-only)
![OpenAI_Logo_White](https://user-images.githubusercontent.com/6267663/220744513-66c99d0e-ed91-4577-982f-e7128d35ce95.svg#gh-dark-mode-only)

<div align="center">
   <p><strong>This is an unofficial SDK for the OpenAI API.</strong>  The OpenAI API can be applied to virtually any task that involves understanding or generating natural language or code. We offer a spectrum of models with different levels of power suitable for different tasks, as well as the ability to fine-tune your own custom models. These models can be used for everything from content generation to semantic search and classification.</p>
  <a href="https://platform.openai.com/docs/introduction"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=2ca47c&style=for-the-badge" /></a>
</div> 

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/openai-go-sdk
```
<!-- End SDK Installation [installation] -->

## Authentication

The OpenAI API uses API keys for authentication. Visit your API Keys page to retrieve the API key you'll use in your requests.

**Remember that your API key is a secret!** Do not share it with others or expose it in any client-side code (browsers, apps). Production requests must be routed through your own backend server where your API key can be securely loaded from an environment variable or key management service.

All API requests should include your API key in an Authorization HTTP header as follows:

```bash
Authorization: Bearer YOUR_API_KEY
```

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"log"
)

func main() {
	s := openaigosdk.New(
		openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var runID string = "<value>"

	var threadID string = "<value>"

	ctx := context.Background()
	res, err := s.Assistants.CancelRun(ctx, runID, threadID)
	if err != nil {
		log.Fatal(err)
	}
	if res.RunObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Assistants](docs/sdks/assistants/README.md)

* [CancelRun](docs/sdks/assistants/README.md#cancelrun) - Cancels a run that is `in_progress`.
* [CreateAssistant](docs/sdks/assistants/README.md#createassistant) - Create an assistant with a model and instructions.
* [CreateAssistantFile](docs/sdks/assistants/README.md#createassistantfile) - Create an assistant file by attaching a [File](/docs/api-reference/files) to an [assistant](/docs/api-reference/assistants).
* [CreateMessage](docs/sdks/assistants/README.md#createmessage) - Create a message.
* [CreateRun](docs/sdks/assistants/README.md#createrun) - Create a run.
* [CreateThread](docs/sdks/assistants/README.md#createthread) - Create a thread.
* [CreateThreadAndRun](docs/sdks/assistants/README.md#createthreadandrun) - Create a thread and run it in one request.
* [DeleteAssistant](docs/sdks/assistants/README.md#deleteassistant) - Delete an assistant.
* [DeleteAssistantFile](docs/sdks/assistants/README.md#deleteassistantfile) - Delete an assistant file.
* [DeleteThread](docs/sdks/assistants/README.md#deletethread) - Delete a thread.
* [GetAssistant](docs/sdks/assistants/README.md#getassistant) - Retrieves an assistant.
* [GetAssistantFile](docs/sdks/assistants/README.md#getassistantfile) - Retrieves an AssistantFile.
* [GetMessage](docs/sdks/assistants/README.md#getmessage) - Retrieve a message.
* [GetMessageFile](docs/sdks/assistants/README.md#getmessagefile) - Retrieves a message file.
* [GetRun](docs/sdks/assistants/README.md#getrun) - Retrieves a run.
* [GetRunStep](docs/sdks/assistants/README.md#getrunstep) - Retrieves a run step.
* [GetThread](docs/sdks/assistants/README.md#getthread) - Retrieves a thread.
* [ListAssistantFiles](docs/sdks/assistants/README.md#listassistantfiles) - Returns a list of assistant files.
* [ListAssistants](docs/sdks/assistants/README.md#listassistants) - Returns a list of assistants.
* [ListMessageFiles](docs/sdks/assistants/README.md#listmessagefiles) - Returns a list of message files.
* [ListMessages](docs/sdks/assistants/README.md#listmessages) - Returns a list of messages for a given thread.
* [ListRunSteps](docs/sdks/assistants/README.md#listrunsteps) - Returns a list of run steps belonging to a run.
* [ListRuns](docs/sdks/assistants/README.md#listruns) - Returns a list of runs belonging to a thread.
* [ModifyAssistant](docs/sdks/assistants/README.md#modifyassistant) - Modifies an assistant.
* [ModifyMessage](docs/sdks/assistants/README.md#modifymessage) - Modifies a message.
* [ModifyRun](docs/sdks/assistants/README.md#modifyrun) - Modifies a run.
* [ModifyThread](docs/sdks/assistants/README.md#modifythread) - Modifies a thread.
* [SubmitToolOuputsToRun](docs/sdks/assistants/README.md#submittoolouputstorun) - When a run has the `status: "requires_action"` and `required_action.type` is `submit_tool_outputs`, this endpoint can be used to submit the outputs from the tool calls once they're all completed. All outputs must be submitted in a single request.


### [Audio](docs/sdks/audio/README.md)

* [CreateSpeech](docs/sdks/audio/README.md#createspeech) - Generates audio from the input text.
* [CreateTranscription](docs/sdks/audio/README.md#createtranscription) - Transcribes audio into the input language.
* [CreateTranslation](docs/sdks/audio/README.md#createtranslation) - Translates audio into English.

### [Chat](docs/sdks/chat/README.md)

* [CreateChatCompletion](docs/sdks/chat/README.md#createchatcompletion) - Creates a model response for the given chat conversation.

### [Completions](docs/sdks/completions/README.md)

* [CreateCompletion](docs/sdks/completions/README.md#createcompletion) - Creates a completion for the provided prompt and parameters.

### [Embeddings](docs/sdks/embeddings/README.md)

* [CreateEmbedding](docs/sdks/embeddings/README.md#createembedding) - Creates an embedding vector representing the input text.

### [Files](docs/sdks/files/README.md)

* [CreateFile](docs/sdks/files/README.md#createfile) - Upload a file that can be used across various endpoints. The size of all the files uploaded by one organization can be up to 100 GB.

The size of individual files can be a maximum of 512 MB or 2 million tokens for Assistants. See the [Assistants Tools guide](/docs/assistants/tools) to learn more about the types of files supported. The Fine-tuning API only supports `.jsonl` files.

Please [contact us](https://help.openai.com/) if you need to increase these storage limits.

* [DeleteFile](docs/sdks/files/README.md#deletefile) - Delete a file.
* [DownloadFile](docs/sdks/files/README.md#downloadfile) - Returns the contents of the specified file.
* [ListFiles](docs/sdks/files/README.md#listfiles) - Returns a list of files that belong to the user's organization.
* [RetrieveFile](docs/sdks/files/README.md#retrievefile) - Returns information about a specific file.

### [FineTuning](docs/sdks/finetuning/README.md)

* [CancelFineTuningJob](docs/sdks/finetuning/README.md#cancelfinetuningjob) - Immediately cancel a fine-tune job.

* [CreateFineTuningJob](docs/sdks/finetuning/README.md#createfinetuningjob) - Creates a fine-tuning job which begins the process of creating a new model from a given dataset.

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

* [CreateModeration](docs/sdks/moderations/README.md#createmoderation) - Classifies if text is potentially harmful.
<!-- End Available Resources and Operations [operations] -->





<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"log"
)

func main() {
	s := openaigosdk.New(
		openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var runID string = "<value>"

	var threadID string = "<value>"

	ctx := context.Background()
	res, err := s.Assistants.CancelRun(ctx, runID, threadID)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.openai.com/v1` | None |

#### Example

```go
package main

import (
	"context"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"log"
)

func main() {
	s := openaigosdk.New(
		openaigosdk.WithServerIndex(0),
		openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var runID string = "<value>"

	var threadID string = "<value>"

	ctx := context.Background()
	res, err := s.Assistants.CancelRun(ctx, runID, threadID)
	if err != nil {
		log.Fatal(err)
	}
	if res.RunObject != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"log"
)

func main() {
	s := openaigosdk.New(
		openaigosdk.WithServerURL("https://api.openai.com/v1"),
		openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var runID string = "<value>"

	var threadID string = "<value>"

	ctx := context.Background()
	res, err := s.Assistants.CancelRun(ctx, runID, threadID)
	if err != nil {
		log.Fatal(err)
	}
	if res.RunObject != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

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
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `APIKeyAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"log"
)

func main() {
	s := openaigosdk.New(
		openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var runID string = "<value>"

	var threadID string = "<value>"

	ctx := context.Background()
	res, err := s.Assistants.CancelRun(ctx, runID, threadID)
	if err != nil {
		log.Fatal(err)
	}
	if res.RunObject != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
