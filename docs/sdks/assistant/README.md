# Assistant
(*Assistant*)

### Available Operations

* [ModifyAssistant](#modifyassistant) - Modifies an assistant.

## ModifyAssistant

Modifies an assistant.

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
        openaigosdk.WithSecurity(""),
    )


    modifyAssistantRequest := shared.ModifyAssistantRequest{
        FileIds: []string{
            "string",
        },
        Metadata: &shared.ModifyAssistantRequestMetadata{},
        Tools: []shared.ModifyAssistantRequestTools{
            shared.CreateModifyAssistantRequestToolsCodeInterpreterTool(
                shared.CodeInterpreterTool{
                    Type: shared.TypeCodeInterpreter,
                },
            ),
        },
    }

    var assistantID string = "string"

    ctx := context.Background()
    res, err := s.Assistant.ModifyAssistant(ctx, modifyAssistantRequest, assistantID)
    if err != nil {
        log.Fatal(err)
    }

    if res.AssistantObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `modifyAssistantRequest`                                                              | [shared.ModifyAssistantRequest](../../../pkg/models/shared/modifyassistantrequest.md) | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `assistantID`                                                                         | *string*                                                                              | :heavy_check_mark:                                                                    | The ID of the assistant to modify.                                                    |


### Response

**[*operations.ModifyAssistantResponse](../../pkg/models/operations/modifyassistantresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
