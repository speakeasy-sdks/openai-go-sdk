# Files
(*Files*)

## Overview

Files are used to upload documents that can be used with features like fine-tuning.

### Available Operations

* [CreateFile](#createfile) - Upload a file that can be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please [contact us](https://help.openai.com/) if you need to increase the storage limit.

* [DeleteFile](#deletefile) - Delete a file.
* [DownloadFile](#downloadfile) - Returns the contents of the specified file.
* [ListFiles](#listfiles) - Returns a list of files that belong to the user's organization.
* [RetrieveFile](#retrievefile) - Returns information about a specific file.

## CreateFile

Upload a file that can be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please [contact us](https://help.openai.com/) if you need to increase the storage limit.


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
    res, err := s.Files.CreateFile(ctx, shared.CreateFileRequest{
        File: shared.CreateFileRequestFile{
            Content: []byte("`'$Z`(L/RH"),
            File: "string",
        },
        Purpose: "string",
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


## DeleteFile

Delete a file.

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


    var fileID string = "string"

    ctx := context.Background()
    res, err := s.Files.DeleteFile(ctx, fileID)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteFileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | The ID of the file to use for this request.           |


### Response

**[*operations.DeleteFileResponse](../../models/operations/deletefileresponse.md), error**


## DownloadFile

Returns the contents of the specified file.

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


    var fileID string = "string"

    ctx := context.Background()
    res, err := s.Files.DownloadFile(ctx, fileID)
    if err != nil {
        log.Fatal(err)
    }

    if res.DownloadFile200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | The ID of the file to use for this request.           |


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
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v2"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Files.ListFiles(ctx)
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


## RetrieveFile

Returns information about a specific file.

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


    var fileID string = "string"

    ctx := context.Background()
    res, err := s.Files.RetrieveFile(ctx, fileID)
    if err != nil {
        log.Fatal(err)
    }

    if res.OpenAIFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | The ID of the file to use for this request.           |


### Response

**[*operations.RetrieveFileResponse](../../models/operations/retrievefileresponse.md), error**

