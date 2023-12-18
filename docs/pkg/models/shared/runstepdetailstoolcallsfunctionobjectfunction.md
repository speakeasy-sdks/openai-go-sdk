# RunStepDetailsToolCallsFunctionObjectFunction

The definition of the function that was called.


## Fields

| Field                                                                                                                                     | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `Arguments`                                                                                                                               | *string*                                                                                                                                  | :heavy_check_mark:                                                                                                                        | The arguments passed to the function.                                                                                                     |
| `Name`                                                                                                                                    | *string*                                                                                                                                  | :heavy_check_mark:                                                                                                                        | The name of the function.                                                                                                                 |
| `Output`                                                                                                                                  | *string*                                                                                                                                  | :heavy_check_mark:                                                                                                                        | The output of the function. This will be `null` if the outputs have not been [submitted](/docs/api-reference/runs/submitToolOutputs) yet. |