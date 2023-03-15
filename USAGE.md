<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/openai-go-sdk"
    "github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
)

func main() {
    s := gpt.New()

    req := operations.CancelFineTuneRequest{
        FineTuneID: "unde",
    }

    ctx := context.Background()
    res, err := s.OpenAI.CancelFineTune(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTune != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->