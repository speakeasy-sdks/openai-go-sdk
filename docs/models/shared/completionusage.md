# CompletionUsage

Usage statistics for the completion request.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CompletionTokens`                                                | *int64*                                                           | :heavy_check_mark:                                                | Number of tokens in the generated completion.                     |
| `PromptTokens`                                                    | *int64*                                                           | :heavy_check_mark:                                                | Number of tokens in the prompt.                                   |
| `TotalTokens`                                                     | *int64*                                                           | :heavy_check_mark:                                                | Total number of tokens used in the request (prompt + completion). |