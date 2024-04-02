# Assistants
(*Assistants*)

## Overview

Build Assistants that can call models and use tools.

### Available Operations

* [CancelRun](#cancelrun) - Cancels a run that is `in_progress`.
* [CreateAssistant](#createassistant) - Create an assistant with a model and instructions.
* [CreateAssistantFile](#createassistantfile) - Create an assistant file by attaching a [File](/docs/api-reference/files) to an [assistant](/docs/api-reference/assistants).
* [CreateMessage](#createmessage) - Create a message.
* [CreateRun](#createrun) - Create a run.
* [CreateThread](#createthread) - Create a thread.
* [CreateThreadAndRun](#createthreadandrun) - Create a thread and run it in one request.
* [DeleteAssistant](#deleteassistant) - Delete an assistant.
* [DeleteAssistantFile](#deleteassistantfile) - Delete an assistant file.
* [DeleteThread](#deletethread) - Delete a thread.
* [GetAssistant](#getassistant) - Retrieves an assistant.
* [GetAssistantFile](#getassistantfile) - Retrieves an AssistantFile.
* [GetMessage](#getmessage) - Retrieve a message.
* [GetMessageFile](#getmessagefile) - Retrieves a message file.
* [GetRun](#getrun) - Retrieves a run.
* [GetRunStep](#getrunstep) - Retrieves a run step.
* [GetThread](#getthread) - Retrieves a thread.
* [ListAssistantFiles](#listassistantfiles) - Returns a list of assistant files.
* [ListAssistants](#listassistants) - Returns a list of assistants.
* [ListMessageFiles](#listmessagefiles) - Returns a list of message files.
* [ListMessages](#listmessages) - Returns a list of messages for a given thread.
* [ListRunSteps](#listrunsteps) - Returns a list of run steps belonging to a run.
* [ListRuns](#listruns) - Returns a list of runs belonging to a thread.
* [ModifyAssistant](#modifyassistant) - Modifies an assistant.
* [ModifyMessage](#modifymessage) - Modifies a message.
* [ModifyRun](#modifyrun) - Modifies a run.
* [ModifyThread](#modifythread) - Modifies a thread.
* [SubmitToolOuputsToRun](#submittoolouputstorun) - When a run has the `status: "requires_action"` and `required_action.type` is `submit_tool_outputs`, this endpoint can be used to submit the outputs from the tool calls once they're all completed. All outputs must be submitted in a single request.


## CancelRun

Cancels a run that is `in_progress`.

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `runID`                                               | *string*                                              | :heavy_check_mark:                                    | The ID of the run to cancel.                          |
| `threadID`                                            | *string*                                              | :heavy_check_mark:                                    | The ID of the thread to which this run belongs.       |


### Response

**[*operations.CancelRunResponse](../../pkg/models/operations/cancelrunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateAssistant

Create an assistant with a model and instructions.

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
    res, err := s.Assistants.CreateAssistant(ctx, shared.CreateAssistantRequest{
        Model: "XTS",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AssistantObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateAssistantRequest](../../pkg/models/shared/createassistantrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateAssistantResponse](../../pkg/models/operations/createassistantresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateAssistantFile

Create an assistant file by attaching a [File](/docs/api-reference/files) to an [assistant](/docs/api-reference/assistants).

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


    createAssistantFileRequest := shared.CreateAssistantFileRequest{
        FileID: "<value>",
    }

    var assistantID string = "file-abc123"

    ctx := context.Background()
    res, err := s.Assistants.CreateAssistantFile(ctx, createAssistantFileRequest, assistantID)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssistantFileObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `createAssistantFileRequest`                                                               | [shared.CreateAssistantFileRequest](../../pkg/models/shared/createassistantfilerequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `assistantID`                                                                              | *string*                                                                                   | :heavy_check_mark:                                                                         | The ID of the assistant for which to create a File.<br/>                                   | file-abc123                                                                                |


### Response

**[*operations.CreateAssistantFileResponse](../../pkg/models/operations/createassistantfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateMessage

Create a message.

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


    createMessageRequest := shared.CreateMessageRequest{
        Content: "<value>",
        Role: shared.CreateMessageRequestRoleUser,
    }

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.CreateMessage(ctx, createMessageRequest, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `createMessageRequest`                                                         | [shared.CreateMessageRequest](../../pkg/models/shared/createmessagerequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `threadID`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | The ID of the [thread](/docs/api-reference/threads) to create a message for.   |


### Response

**[*operations.CreateMessageResponse](../../pkg/models/operations/createmessageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateRun

Create a run.

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


    createRunRequest := shared.CreateRunRequest{
        AssistantID: "<value>",
        Temperature: openaigosdk.Float64(1),
    }

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.CreateRun(ctx, createRunRequest, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RunObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `createRunRequest`                                                     | [shared.CreateRunRequest](../../pkg/models/shared/createrunrequest.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `threadID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | The ID of the thread to run.                                           |


### Response

**[*operations.CreateRunResponse](../../pkg/models/operations/createrunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateThread

Create a thread.

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
    res, err := s.Assistants.CreateThread(ctx, &shared.CreateThreadRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ThreadObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateThreadRequest](../../pkg/models/shared/createthreadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateThreadResponse](../../pkg/models/operations/createthreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateThreadAndRun

Create a thread and run it in one request.

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
    res, err := s.Assistants.CreateThreadAndRun(ctx, shared.CreateThreadAndRunRequest{
        AssistantID: "<value>",
        Temperature: openaigosdk.Float64(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RunObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.CreateThreadAndRunRequest](../../pkg/models/shared/createthreadandrunrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateThreadAndRunResponse](../../pkg/models/operations/createthreadandrunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAssistant

Delete an assistant.

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


    var assistantID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.DeleteAssistant(ctx, assistantID)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAssistantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `assistantID`                                         | *string*                                              | :heavy_check_mark:                                    | The ID of the assistant to delete.                    |


### Response

**[*operations.DeleteAssistantResponse](../../pkg/models/operations/deleteassistantresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAssistantFile

Delete an assistant file.

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


    var assistantID string = "<value>"

    var fileID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.DeleteAssistantFile(ctx, assistantID, fileID)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAssistantFileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `assistantID`                                         | *string*                                              | :heavy_check_mark:                                    | The ID of the assistant that the file belongs to.     |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | The ID of the file to delete.                         |


### Response

**[*operations.DeleteAssistantFileResponse](../../pkg/models/operations/deleteassistantfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteThread

Delete a thread.

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


    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.DeleteThread(ctx, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteThreadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `threadID`                                            | *string*                                              | :heavy_check_mark:                                    | The ID of the thread to delete.                       |


### Response

**[*operations.DeleteThreadResponse](../../pkg/models/operations/deletethreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAssistant

Retrieves an assistant.

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


    var assistantID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.GetAssistant(ctx, assistantID)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssistantObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `assistantID`                                         | *string*                                              | :heavy_check_mark:                                    | The ID of the assistant to retrieve.                  |


### Response

**[*operations.GetAssistantResponse](../../pkg/models/operations/getassistantresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAssistantFile

Retrieves an AssistantFile.

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


    var assistantID string = "<value>"

    var fileID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.GetAssistantFile(ctx, assistantID, fileID)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssistantFileObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `assistantID`                                         | *string*                                              | :heavy_check_mark:                                    | The ID of the assistant who the file belongs to.      |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | The ID of the file we're getting.                     |


### Response

**[*operations.GetAssistantFileResponse](../../pkg/models/operations/getassistantfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMessage

Retrieve a message.

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


    var messageID string = "<value>"

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.GetMessage(ctx, messageID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `messageID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | The ID of the message to retrieve.                                                 |
| `threadID`                                                                         | *string*                                                                           | :heavy_check_mark:                                                                 | The ID of the [thread](/docs/api-reference/threads) to which this message belongs. |


### Response

**[*operations.GetMessageResponse](../../pkg/models/operations/getmessageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMessageFile

Retrieves a message file.

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


    var fileID string = "file-abc123"

    var messageID string = "msg_abc123"

    var threadID string = "thread_abc123"

    ctx := context.Background()
    res, err := s.Assistants.GetMessageFile(ctx, fileID, messageID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageFileObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `fileID`                                                   | *string*                                                   | :heavy_check_mark:                                         | The ID of the file being retrieved.                        | file-abc123                                                |
| `messageID`                                                | *string*                                                   | :heavy_check_mark:                                         | The ID of the message the file belongs to.                 | msg_abc123                                                 |
| `threadID`                                                 | *string*                                                   | :heavy_check_mark:                                         | The ID of the thread to which the message and File belong. | thread_abc123                                              |


### Response

**[*operations.GetMessageFileResponse](../../pkg/models/operations/getmessagefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRun

Retrieves a run.

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


    var runID string = "<value>"

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.GetRun(ctx, runID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RunObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `runID`                                                           | *string*                                                          | :heavy_check_mark:                                                | The ID of the run to retrieve.                                    |
| `threadID`                                                        | *string*                                                          | :heavy_check_mark:                                                | The ID of the [thread](/docs/api-reference/threads) that was run. |


### Response

**[*operations.GetRunResponse](../../pkg/models/operations/getrunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRunStep

Retrieves a run step.

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


    var runID string = "<value>"

    var stepID string = "<value>"

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.GetRunStep(ctx, runID, stepID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RunStepObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `runID`                                                     | *string*                                                    | :heavy_check_mark:                                          | The ID of the run to which the run step belongs.            |
| `stepID`                                                    | *string*                                                    | :heavy_check_mark:                                          | The ID of the run step to retrieve.                         |
| `threadID`                                                  | *string*                                                    | :heavy_check_mark:                                          | The ID of the thread to which the run and run step belongs. |


### Response

**[*operations.GetRunStepResponse](../../pkg/models/operations/getrunstepresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetThread

Retrieves a thread.

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


    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.GetThread(ctx, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ThreadObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `threadID`                                            | *string*                                              | :heavy_check_mark:                                    | The ID of the thread to retrieve.                     |


### Response

**[*operations.GetThreadResponse](../../pkg/models/operations/getthreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListAssistantFiles

Returns a list of assistant files.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"context"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/operations"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListAssistantFiles(ctx, operations.ListAssistantFilesRequest{
        AssistantID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAssistantFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListAssistantFilesRequest](../../pkg/models/operations/listassistantfilesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.ListAssistantFilesResponse](../../pkg/models/operations/listassistantfilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListAssistants

Returns a list of assistants.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var after *string = openaigosdk.String("<value>")

    var before *string = openaigosdk.String("<value>")

    var limit *int64 = openaigosdk.Int64(20)

    var order *operations.QueryParamOrder = operations.QueryParamOrderDesc.ToPointer()

    ctx := context.Background()
    res, err := s.Assistants.ListAssistants(ctx, after, before, limit, order)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAssistantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                    |
| `after`                                                                                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | A cursor for use in pagination. `after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include after=obj_foo in order to fetch the next page of the list.<br/>   |
| `before`                                                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | A cursor for use in pagination. `before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with obj_foo, your subsequent call can include before=obj_foo in order to fetch the previous page of the list.<br/> |
| `limit`                                                                                                                                                                                                                                                                                | **int64*                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 20.<br/>                                                                                                                                                                        |
| `order`                                                                                                                                                                                                                                                                                | [*operations.QueryParamOrder](../../pkg/models/operations/queryparamorder.md)                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.<br/>                                                                                                                                                               |


### Response

**[*operations.ListAssistantsResponse](../../pkg/models/operations/listassistantsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListMessageFiles

Returns a list of message files.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"context"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/operations"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListMessageFiles(ctx, operations.ListMessageFilesRequest{
        MessageID: "<value>",
        ThreadID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMessageFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListMessageFilesRequest](../../pkg/models/operations/listmessagefilesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListMessageFilesResponse](../../pkg/models/operations/listmessagefilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListMessages

Returns a list of messages for a given thread.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"context"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/operations"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListMessages(ctx, operations.ListMessagesRequest{
        ThreadID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListMessagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListMessagesRequest](../../pkg/models/operations/listmessagesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListMessagesResponse](../../pkg/models/operations/listmessagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListRunSteps

Returns a list of run steps belonging to a run.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"context"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/operations"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListRunSteps(ctx, operations.ListRunStepsRequest{
        RunID: "<value>",
        ThreadID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListRunStepsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListRunStepsRequest](../../pkg/models/operations/listrunstepsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListRunStepsResponse](../../pkg/models/operations/listrunstepsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListRuns

Returns a list of runs belonging to a thread.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"context"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/operations"
	"log"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListRuns(ctx, operations.ListRunsRequest{
        ThreadID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListRunsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListRunsRequest](../../pkg/models/operations/listrunsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ListRunsResponse](../../pkg/models/operations/listrunsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifyAssistant

Modifies an assistant.

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


    modifyAssistantRequest := shared.ModifyAssistantRequest{}

    var assistantID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.ModifyAssistant(ctx, modifyAssistantRequest, assistantID)
    if err != nil {
        log.Fatal(err)
    }
    if res.AssistantObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `modifyAssistantRequest`                                                           | [shared.ModifyAssistantRequest](../../pkg/models/shared/modifyassistantrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `assistantID`                                                                      | *string*                                                                           | :heavy_check_mark:                                                                 | The ID of the assistant to modify.                                                 |


### Response

**[*operations.ModifyAssistantResponse](../../pkg/models/operations/modifyassistantresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifyMessage

Modifies a message.

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


    modifyMessageRequest := shared.ModifyMessageRequest{}

    var messageID string = "<value>"

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.ModifyMessage(ctx, modifyMessageRequest, messageID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `modifyMessageRequest`                                                         | [shared.ModifyMessageRequest](../../pkg/models/shared/modifymessagerequest.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `messageID`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The ID of the message to modify.                                               |
| `threadID`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | The ID of the thread to which this message belongs.                            |


### Response

**[*operations.ModifyMessageResponse](../../pkg/models/operations/modifymessageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifyRun

Modifies a run.

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


    modifyRunRequest := shared.ModifyRunRequest{}

    var runID string = "<value>"

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.ModifyRun(ctx, modifyRunRequest, runID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RunObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `modifyRunRequest`                                                     | [shared.ModifyRunRequest](../../pkg/models/shared/modifyrunrequest.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `runID`                                                                | *string*                                                               | :heavy_check_mark:                                                     | The ID of the run to modify.                                           |
| `threadID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | The ID of the [thread](/docs/api-reference/threads) that was run.      |


### Response

**[*operations.ModifyRunResponse](../../pkg/models/operations/modifyrunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ModifyThread

Modifies a thread.

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


    modifyThreadRequest := shared.ModifyThreadRequest{}

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.ModifyThread(ctx, modifyThreadRequest, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.ThreadObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `modifyThreadRequest`                                                        | [shared.ModifyThreadRequest](../../pkg/models/shared/modifythreadrequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `threadID`                                                                   | *string*                                                                     | :heavy_check_mark:                                                           | The ID of the thread to modify. Only the `metadata` can be modified.         |


### Response

**[*operations.ModifyThreadResponse](../../pkg/models/operations/modifythreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SubmitToolOuputsToRun

When a run has the `status: "requires_action"` and `required_action.type` is `submit_tool_outputs`, this endpoint can be used to submit the outputs from the tool calls once they're all completed. All outputs must be submitted in a single request.


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


    submitToolOutputsRunRequest := shared.SubmitToolOutputsRunRequest{
        ToolOutputs: []shared.ToolOutputs{
            shared.ToolOutputs{},
        },
    }

    var runID string = "<value>"

    var threadID string = "<value>"

    ctx := context.Background()
    res, err := s.Assistants.SubmitToolOuputsToRun(ctx, submitToolOutputsRunRequest, runID, threadID)
    if err != nil {
        log.Fatal(err)
    }
    if res.RunObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `submitToolOutputsRunRequest`                                                                | [shared.SubmitToolOutputsRunRequest](../../pkg/models/shared/submittooloutputsrunrequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `runID`                                                                                      | *string*                                                                                     | :heavy_check_mark:                                                                           | The ID of the run that requires the tool output submission.                                  |
| `threadID`                                                                                   | *string*                                                                                     | :heavy_check_mark:                                                                           | The ID of the [thread](/docs/api-reference/threads) to which this run belongs.               |


### Response

**[*operations.SubmitToolOuputsToRunResponse](../../pkg/models/operations/submittoolouputstorunresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
