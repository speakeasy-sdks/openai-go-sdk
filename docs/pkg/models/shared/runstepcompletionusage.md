# RunStepCompletionUsage

Usage statistics related to the run step. This value will be `null` while the run step's status is `in_progress`.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CompletionTokens`                                                | *int64*                                                           | :heavy_check_mark:                                                | Number of completion tokens used over the course of the run step. |
| `PromptTokens`                                                    | *int64*                                                           | :heavy_check_mark:                                                | Number of prompt tokens used over the course of the run step.     |
| `TotalTokens`                                                     | *int64*                                                           | :heavy_check_mark:                                                | Total number of tokens used (prompt + completion).                |