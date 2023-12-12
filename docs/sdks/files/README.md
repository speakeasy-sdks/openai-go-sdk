# Files
(*Files*)

## Overview

Files are used to upload documents that can be used with features like Assistants and Fine-tuning.

### Available Operations

* [CreateFile](#createfile) - Upload a file that can be used across various endpoints. The size of all the files uploaded by one organization can be up to 100 GB.

The size of individual files can be a maximum of 512 MB. See the [Assistants Tools guide](/docs/assistants/tools) to learn more about the types of files supported. The Fine-tuning API only supports `.jsonl` files.

Please [contact us](https://help.openai.com/) if you need to increase these storage limits.

* [DeleteFile](#deletefile) - Delete a file.
* [DownloadFile](#downloadfile) - Returns the contents of the specified file.
* [ListFiles](#listfiles) - Returns a list of files that belong to the user's organization.
* [RetrieveFile](#retrievefile) - Returns information about a specific file.

## CreateFile

Upload a file that can be used across various endpoints. The size of all the files uploaded by one organization can be up to 100 GB.

The size of individual files can be a maximum of 512 MB. See the [Assistants Tools guide](/docs/assistants/tools) to learn more about the types of files supported. The Fine-tuning API only supports `.jsonl` files.

Please [contact us](https://help.openai.com/) if you need to increase these storage limits.


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
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Files.CreateFile(ctx, shared.CreateFileRequest{
        File: shared.File{
            Content: []byte("0xf10df1a3b9"),
            FileName: "rap_national.mp4v",
        },
        Purpose: shared.PurposeFineTune,
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreateFileRequest](../../pkg/models/shared/createfilerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateFileResponse](../../pkg/models/operations/createfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteFile

Delete a file.

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
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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

**[*operations.DeleteFileResponse](../../pkg/models/operations/deletefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DownloadFile

Returns the contents of the specified file.

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
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var fileID string = "string"

    ctx := context.Background()
    res, err := s.Files.DownloadFile(ctx, fileID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
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

**[*operations.DownloadFileResponse](../../pkg/models/operations/downloadfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListFiles

Returns a list of files that belong to the user's organization.

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
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var purpose *string = "string"

    ctx := context.Background()
    res, err := s.Files.ListFiles(ctx, purpose)
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
| `purpose`                                             | **string*                                             | :heavy_minus_sign:                                    | Only return files with the given purpose.             |


### Response

**[*operations.ListFilesResponse](../../pkg/models/operations/listfilesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RetrieveFile

Returns information about a specific file.

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
        openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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

**[*operations.RetrieveFileResponse](../../pkg/models/operations/retrievefileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
