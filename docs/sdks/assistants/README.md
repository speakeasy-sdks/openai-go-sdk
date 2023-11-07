# Assistants
(*.Assistants*)

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
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var runID string = "string"

    var threadID string = "string"

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

**[*operations.CancelRunResponse](../../models/operations/cancelrunresponse.md), error**


## CreateAssistant

Create an assistant with a model and instructions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.CreateAssistant(ctx, shared.CreateAssistantRequest{
        FileIds: []string{
            "string",
        },
        Metadata: &shared.CreateAssistantRequestMetadata{},
        Model: "XTS",
        Tools: []shared.CreateAssistantRequestTools{
            shared.CreateCreateAssistantRequestToolsRetrievalTool(
                shared.RetrievalTool{
                    Type: shared.SchemasAssistantToolsRetrievalTypeRetrieval,
                },
            ),
        },
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateAssistantRequest](../../models/shared/createassistantrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateAssistantResponse](../../models/operations/createassistantresponse.md), error**


## CreateAssistantFile

Create an assistant file by attaching a [File](/docs/api-reference/files) to an [assistant](/docs/api-reference/assistants).

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    createAssistantFileRequest := shared.CreateAssistantFileRequest{
        FileID: "string",
    }

    var assistantID string = "file-AF1WoRqd3aJAHsqc9NY7iL8F"

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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `createAssistantFileRequest`                                                           | [shared.CreateAssistantFileRequest](../../models/shared/createassistantfilerequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `assistantID`                                                                          | *string*                                                                               | :heavy_check_mark:                                                                     | The ID of the assistant for which to create a File.<br/>                               | file-AF1WoRqd3aJAHsqc9NY7iL8F                                                          |


### Response

**[*operations.CreateAssistantFileResponse](../../models/operations/createassistantfileresponse.md), error**


## CreateMessage

Create a message.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    createMessageRequest := shared.CreateMessageRequest{
        Content: "string",
        FileIds: []string{
            "string",
        },
        Metadata: &shared.CreateMessageRequestMetadata{},
        Role: shared.CreateMessageRequestRoleUser,
    }

    var threadID string = "string"

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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `createMessageRequest`                                                       | [shared.CreateMessageRequest](../../models/shared/createmessagerequest.md)   | :heavy_check_mark:                                                           | N/A                                                                          |
| `threadID`                                                                   | *string*                                                                     | :heavy_check_mark:                                                           | The ID of the [thread](/docs/api-reference/threads) to create a message for. |


### Response

**[*operations.CreateMessageResponse](../../models/operations/createmessageresponse.md), error**


## CreateRun

Create a run.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    createRunRequest := shared.CreateRunRequest{
        AssistantID: "string",
        Metadata: &shared.CreateRunRequestMetadata{},
        Tools: []shared.CreateRunRequestTools{
            shared.CreateCreateRunRequestToolsCodeInterpreterTool(
                shared.CodeInterpreterTool{
                    Type: shared.TypeCodeInterpreter,
                },
            ),
        },
    }

    var threadID string = "string"

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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `createRunRequest`                                                 | [shared.CreateRunRequest](../../models/shared/createrunrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `threadID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | The ID of the thread to run.                                       |


### Response

**[*operations.CreateRunResponse](../../models/operations/createrunresponse.md), error**


## CreateThread

Create a thread.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.CreateThread(ctx, shared.CreateThreadRequest{
        Messages: []shared.CreateMessageRequest{
            shared.CreateMessageRequest{
                Content: "string",
                FileIds: []string{
                    "string",
                },
                Metadata: &shared.CreateMessageRequestMetadata{},
                Role: shared.CreateMessageRequestRoleUser,
            },
        },
        Metadata: &shared.CreateThreadRequestMetadata{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ThreadObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateThreadRequest](../../models/shared/createthreadrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateThreadResponse](../../models/operations/createthreadresponse.md), error**


## CreateThreadAndRun

Create a thread and run it in one request.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.CreateThreadAndRun(ctx, shared.CreateThreadAndRunRequest{
        AssistantID: "string",
        Metadata: &shared.CreateThreadAndRunRequestMetadata{},
        Thread: &shared.CreateThreadRequest{
            Messages: []shared.CreateMessageRequest{
                shared.CreateMessageRequest{
                    Content: "string",
                    FileIds: []string{
                        "string",
                    },
                    Metadata: &shared.CreateMessageRequestMetadata{},
                    Role: shared.CreateMessageRequestRoleUser,
                },
            },
            Metadata: &shared.CreateThreadRequestMetadata{},
        },
        Tools: []shared.CreateThreadAndRunRequestTools{
            shared.CreateCreateThreadAndRunRequestToolsAssistantToolsFunction(
                shared.AssistantToolsFunction{
                    Function: shared.AssistantToolsFunctionFunction{
                        Description: "Exclusive holistic moratorium",
                        Name: "string",
                        Parameters: map[string]interface{}{
                            "key": "string",
                        },
                    },
                    Type: shared.AssistantToolsFunctionTypeFunction,
                },
            ),
        },
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.CreateThreadAndRunRequest](../../models/shared/createthreadandrunrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateThreadAndRunResponse](../../models/operations/createthreadandrunresponse.md), error**


## DeleteAssistant

Delete an assistant.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var assistantID string = "string"

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

**[*operations.DeleteAssistantResponse](../../models/operations/deleteassistantresponse.md), error**


## DeleteAssistantFile

Delete an assistant file.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var assistantID string = "string"

    var fileID string = "string"

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

**[*operations.DeleteAssistantFileResponse](../../models/operations/deleteassistantfileresponse.md), error**


## DeleteThread

Delete a thread.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var threadID string = "string"

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

**[*operations.DeleteThreadResponse](../../models/operations/deletethreadresponse.md), error**


## GetAssistant

Retrieves an assistant.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var assistantID string = "string"

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

**[*operations.GetAssistantResponse](../../models/operations/getassistantresponse.md), error**


## GetAssistantFile

Retrieves an AssistantFile.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var assistantID string = "string"

    var fileID string = "string"

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

**[*operations.GetAssistantFileResponse](../../models/operations/getassistantfileresponse.md), error**


## GetMessage

Retrieve a message.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var messageID string = "string"

    var threadID string = "string"

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

**[*operations.GetMessageResponse](../../models/operations/getmessageresponse.md), error**


## GetMessageFile

Retrieves a message file.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var fileID string = "file-AF1WoRqd3aJAHsqc9NY7iL8F"

    var messageID string = "msg_AF1WoRqd3aJAHsqc9NY7iL8F"

    var threadID string = "thread_AF1WoRqd3aJAHsqc9NY7iL8F"

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
| `fileID`                                                   | *string*                                                   | :heavy_check_mark:                                         | The ID of the file being retrieved.                        | file-AF1WoRqd3aJAHsqc9NY7iL8F                              |
| `messageID`                                                | *string*                                                   | :heavy_check_mark:                                         | The ID of the message the file belongs to.                 | msg_AF1WoRqd3aJAHsqc9NY7iL8F                               |
| `threadID`                                                 | *string*                                                   | :heavy_check_mark:                                         | The ID of the thread to which the message and File belong. | thread_AF1WoRqd3aJAHsqc9NY7iL8F                            |


### Response

**[*operations.GetMessageFileResponse](../../models/operations/getmessagefileresponse.md), error**


## GetRun

Retrieves a run.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var runID string = "string"

    var threadID string = "string"

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

**[*operations.GetRunResponse](../../models/operations/getrunresponse.md), error**


## GetRunStep

Retrieves a run step.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var runID string = "string"

    var stepID string = "string"

    var threadID string = "string"

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

**[*operations.GetRunStepResponse](../../models/operations/getrunstepresponse.md), error**


## GetThread

Retrieves a thread.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var threadID string = "string"

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

**[*operations.GetThreadResponse](../../models/operations/getthreadresponse.md), error**


## ListAssistantFiles

Returns a list of assistant files.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListAssistantFiles(ctx, operations.ListAssistantFilesRequest{
        AssistantID: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAssistantFilesRequest](../../models/operations/listassistantfilesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListAssistantFilesResponse](../../models/operations/listassistantfilesresponse.md), error**


## ListAssistants

Returns a list of assistants.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    var after *string = "string"

    var before *string = "string"

    var limit *int64 = 948776

    var order *operations.QueryParamOrder = operations.QueryParamOrderAsc

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
| `order`                                                                                                                                                                                                                                                                                | [*operations.QueryParamOrder](../../models/operations/queryparamorder.md)                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                     | Sort order by the `created_at` timestamp of the objects. `asc` for ascending order and `desc` for descending order.<br/>                                                                                                                                                               |


### Response

**[*operations.ListAssistantsResponse](../../models/operations/listassistantsresponse.md), error**


## ListMessageFiles

Returns a list of message files.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListMessageFiles(ctx, operations.ListMessageFilesRequest{
        MessageID: "string",
        ThreadID: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListMessageFilesRequest](../../models/operations/listmessagefilesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListMessageFilesResponse](../../models/operations/listmessagefilesresponse.md), error**


## ListMessages

Returns a list of messages for a given thread.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListMessages(ctx, operations.ListMessagesRequest{
        ThreadID: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListMessagesRequest](../../models/operations/listmessagesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListMessagesResponse](../../models/operations/listmessagesresponse.md), error**


## ListRunSteps

Returns a list of run steps belonging to a run.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListRunSteps(ctx, operations.ListRunStepsRequest{
        RunID: "string",
        ThreadID: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListRunStepsRequest](../../models/operations/listrunstepsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListRunStepsResponse](../../models/operations/listrunstepsresponse.md), error**


## ListRuns

Returns a list of runs belonging to a thread.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Assistants.ListRuns(ctx, operations.ListRunsRequest{
        ThreadID: "string",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListRunsRequest](../../models/operations/listrunsrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.ListRunsResponse](../../models/operations/listrunsresponse.md), error**


## ModifyMessage

Modifies a message.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    modifyMessageRequest := shared.ModifyMessageRequest{
        Metadata: &shared.ModifyMessageRequestMetadata{},
    }

    var messageID string = "string"

    var threadID string = "string"

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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `modifyMessageRequest`                                                     | [shared.ModifyMessageRequest](../../models/shared/modifymessagerequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `messageID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | The ID of the message to modify.                                           |
| `threadID`                                                                 | *string*                                                                   | :heavy_check_mark:                                                         | The ID of the thread to which this message belongs.                        |


### Response

**[*operations.ModifyMessageResponse](../../models/operations/modifymessageresponse.md), error**


## ModifyRun

Modifies a run.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    modifyRunRequest := shared.ModifyRunRequest{
        Metadata: &shared.ModifyRunRequestMetadata{},
    }

    var runID string = "string"

    var threadID string = "string"

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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `modifyRunRequest`                                                 | [shared.ModifyRunRequest](../../models/shared/modifyrunrequest.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `runID`                                                            | *string*                                                           | :heavy_check_mark:                                                 | The ID of the run to modify.                                       |
| `threadID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | The ID of the [thread](/docs/api-reference/threads) that was run.  |


### Response

**[*operations.ModifyRunResponse](../../models/operations/modifyrunresponse.md), error**


## ModifyThread

Modifies a thread.

### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    modifyThreadRequest := shared.ModifyThreadRequest{
        Metadata: &shared.ModifyThreadRequestMetadata{},
    }

    var threadID string = "string"

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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `modifyThreadRequest`                                                    | [shared.ModifyThreadRequest](../../models/shared/modifythreadrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `threadID`                                                               | *string*                                                                 | :heavy_check_mark:                                                       | The ID of the thread to modify. Only the `metadata` can be modified.     |


### Response

**[*operations.ModifyThreadResponse](../../models/operations/modifythreadresponse.md), error**


## SubmitToolOuputsToRun

When a run has the `status: "requires_action"` and `required_action.type` is `submit_tool_outputs`, this endpoint can be used to submit the outputs from the tool calls once they're all completed. All outputs must be submitted in a single request.


### Example Usage

```go
package main

import(
	"context"
	"log"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v3"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )


    submitToolOutputsRunRequest := shared.SubmitToolOutputsRunRequest{
        ToolOutputs: []shared.ToolOutputs{
            shared.ToolOutputs{},
        },
    }

    var runID string = "string"

    var threadID string = "string"

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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `submitToolOutputsRunRequest`                                                            | [shared.SubmitToolOutputsRunRequest](../../models/shared/submittooloutputsrunrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `runID`                                                                                  | *string*                                                                                 | :heavy_check_mark:                                                                       | The ID of the run that requires the tool output submission.                              |
| `threadID`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | The ID of the [thread](/docs/api-reference/threads) to which this run belongs.           |


### Response

**[*operations.SubmitToolOuputsToRunResponse](../../models/operations/submittoolouputstorunresponse.md), error**

