# ChatCompletionNamedToolChoice

Specifies a tool the model should use. Use to force the model to call a specific function.


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `Function`                                                                                                          | [shared.ChatCompletionNamedToolChoiceFunction](../../../pkg/models/shared/chatcompletionnamedtoolchoicefunction.md) | :heavy_check_mark:                                                                                                  | N/A                                                                                                                 |
| `Type`                                                                                                              | [shared.ChatCompletionNamedToolChoiceType](../../../pkg/models/shared/chatcompletionnamedtoolchoicetype.md)         | :heavy_check_mark:                                                                                                  | The type of the tool. Currently, only `function` is supported.                                                      |