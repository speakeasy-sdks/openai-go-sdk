<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/openai-go-sdk"
    "github.com/speakeasy-sdks/openai-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/openai-go-sdk/pkg/models/operations"
)

func main() {
    s := openai.New()
    
    req := operations.CancelFineTuneRequest{
        PathParams: operations.CancelFineTunePathParams{
            FineTuneID: "unde",
        },
    }
    
    res, err := s.OpenAI.CancelFineTune(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.FineTune != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->