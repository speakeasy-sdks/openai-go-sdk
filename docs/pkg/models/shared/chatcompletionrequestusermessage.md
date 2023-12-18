# ChatCompletionRequestUserMessage


## Fields

| Field                                                                                                                        | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `Content`                                                                                                                    | [shared.Content](../../../pkg/models/shared/content.md)                                                                      | :heavy_check_mark:                                                                                                           | The contents of the user message.<br/>                                                                                       |
| `Name`                                                                                                                       | **string*                                                                                                                    | :heavy_minus_sign:                                                                                                           | An optional name for the participant. Provides the model information to differentiate between participants of the same role. |
| `Role`                                                                                                                       | [shared.ChatCompletionRequestUserMessageRole](../../../pkg/models/shared/chatcompletionrequestusermessagerole.md)            | :heavy_check_mark:                                                                                                           | The role of the messages author, in this case `user`.                                                                        |