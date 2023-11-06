# Models
(*Models*)

## Overview

List and describe the various models available in the API.

### Available Operations

* [DeleteModel](#deletemodel) - Delete a fine-tuned model. You must have the Owner role in your organization to delete a model.
* [ListModels](#listmodels) - Lists the currently available models, and provides basic information about each one such as the owner and availability.
* [RetrieveModel](#retrievemodel) - Retrieves a model instance, providing basic information about the model such as the owner and permissioning.

## DeleteModel

Delete a fine-tuned model. You must have the Owner role in your organization to delete a model.

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


    var model string = "ft:gpt-3.5-turbo:acemeco:suffix:abc123"

    ctx := context.Background()
    res, err := s.Models.DeleteModel(ctx, model)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteModelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `model`                                               | *string*                                              | :heavy_check_mark:                                    | The model to delete                                   | ft:gpt-3.5-turbo:acemeco:suffix:abc123                |


### Response

**[*operations.DeleteModelResponse](../../models/operations/deletemodelresponse.md), error**


## ListModels

Lists the currently available models, and provides basic information about each one such as the owner and availability.

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
    res, err := s.Models.ListModels(ctx)
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


## RetrieveModel

Retrieves a model instance, providing basic information about the model such as the owner and permissioning.

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


    var model string = "gpt-3.5-turbo"

    ctx := context.Background()
    res, err := s.Models.RetrieveModel(ctx, model)
    if err != nil {
        log.Fatal(err)
    }

    if res.Model != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `model`                                               | *string*                                              | :heavy_check_mark:                                    | The ID of the model to use for this request           | gpt-3.5-turbo                                         |


### Response

**[*operations.RetrieveModelResponse](../../models/operations/retrievemodelresponse.md), error**

