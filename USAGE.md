<!-- Start SDK Example Usage -->


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
    res, err := s.Audio.CreateTranscription(ctx, shared.CreateTranscriptionRequest{
        File: shared.CreateTranscriptionRequestFile{
            Content: []byte("\#BbTW'zX9"),
            File: "string",
        },
        Model: shared.CreateCreateTranscriptionRequestModelCreateTranscriptionRequestModel2(
        shared.CreateTranscriptionRequestModel2Whisper1,
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTranscriptionResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->