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
    res, err := s.Images.CreateImage(ctx, shared.CreateImageRequest{
        Model: shared.CreateCreateImageRequestModelCreateImageRequest2(
        shared.CreateImageRequest2DallE3,
        ),
        N: openaigosdk.Int64(1),
        Prompt: "A cute baby sea otter",
        Quality: shared.QualityStandard.ToPointer(),
        ResponseFormat: shared.CreateImageRequestResponseFormatURL.ToPointer(),
        Size: shared.CreateImageRequestSizeOneThousandAndTwentyFourx1024.ToPointer(),
        Style: shared.StyleVivid.ToPointer(),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.CreateImageRequest](../../pkg/models/shared/createimagerequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CreateImageResponse](../../pkg/models/operations/createimageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateImageEdit

Creates an edited or extended image given an original image and a prompt.

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
    res, err := s.Images.CreateImageEdit(ctx, shared.CreateImageEditRequest{
        Image: shared.CreateImageEditRequestImage{
            Content: []byte("0x3e31F4cec5"),
            FileName: "facilitator_gosh_hatchback.mpe",
        },
        Mask: &shared.Mask{
            Content: []byte("0xFC5456e4eC"),
            FileName: "electric_cambridgeshire.jpeg",
        },
        Model: shared.CreateCreateImageEditRequestModelCreateImageEditRequest2(
        shared.CreateImageEditRequest2DallE2,
        ),
        N: openaigosdk.Int64(1),
        Prompt: "A cute baby sea otter wearing a beret",
        ResponseFormat: shared.CreateImageEditRequestResponseFormatURL.ToPointer(),
        Size: shared.SizeOneThousandAndTwentyFourx1024.ToPointer(),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateImageEditRequest](../../pkg/models/shared/createimageeditrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateImageEditResponse](../../pkg/models/operations/createimageeditresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateImageVariation

Creates a variation of a given image.

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
    res, err := s.Images.CreateImageVariation(ctx, shared.CreateImageVariationRequest{
        Image: shared.CreateImageVariationRequestImage{
            Content: []byte("0xfdd5b8DcDa"),
            FileName: "fantastic.gif",
        },
        Model: shared.CreateCreateImageVariationRequestModelCreateImageVariationRequest2(
        shared.CreateImageVariationRequest2DallE2,
        ),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.CreateImageVariationRequest](../../pkg/models/shared/createimagevariationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateImageVariationResponse](../../pkg/models/operations/createimagevariationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
