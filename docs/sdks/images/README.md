# Images
(*Images*)

## Overview

Given a prompt and/or an input image, the model will generate a new image.

### Available Operations

* [CreateImage](#createimage) - Creates an image given a prompt.
* [CreateImageEdit](#createimageedit) - Creates an edited or extended image given an original image and a prompt.
* [CreateImageVariation](#createimagevariation) - Creates a variation of a given image.

## CreateImage

Creates an image given a prompt.

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
    res, err := s.Images.CreateImage(ctx, shared.CreateImageRequest{
        N: openaigosdk.Int64(1),
        Prompt: "A cute baby sea otter",
        ResponseFormat: shared.CreateImageRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: openaigosdk.String("user-1234"),
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
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v2"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Images.CreateImageEdit(ctx, shared.CreateImageEditRequest{
        Image: shared.CreateImageEditRequestImage{
            Content: []byte("0]/(|3W_T9"),
            Image: "https://loremflickr.com/640/480",
        },
        Mask: &shared.CreateImageEditRequestMask{
            Content: []byte("`^YjrpxopK"),
            Mask: "plum",
        },
        N: openaigosdk.Int64(1),
        Prompt: "A cute baby sea otter wearing a beret",
        ResponseFormat: shared.CreateImageEditRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageEditRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: openaigosdk.String("user-1234"),
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
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v2"
	"github.com/speakeasy-sdks/openai-go-sdk/v2/pkg/models/shared"
)

func main() {
    s := openaigosdk.New(
        openaigosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Images.CreateImageVariation(ctx, shared.CreateImageVariationRequest{
        Image: shared.CreateImageVariationRequestImage{
            Content: []byte("`YY7PCrWuK"),
            Image: "https://loremflickr.com/640/480",
        },
        N: openaigosdk.Int64(1),
        ResponseFormat: shared.CreateImageVariationRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageVariationRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        User: openaigosdk.String("user-1234"),
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

