<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	openaigosdk "github.com/speakeasy-sdks/openai-go-sdk/v4"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"log"
)

func main() {
	s := openaigosdk.New(
		openaigosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var runID string = "<value>"

	var threadID string = "<value>"

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
<!-- End SDK Example Usage [usage] -->