# MessageFileObject

A list of files attached to a `message`.


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `CreatedAt`                                                                                                      | *int64*                                                                                                          | :heavy_check_mark:                                                                                               | The Unix timestamp (in seconds) for when the message file was created.                                           |
| `ID`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The identifier, which can be referenced in API endpoints.                                                        |
| `MessageID`                                                                                                      | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The ID of the [message](/docs/api-reference/messages) that the [File](/docs/api-reference/files) is attached to. |
| `Object`                                                                                                         | [shared.MessageFileObjectObject](../../../pkg/models/shared/messagefileobjectobject.md)                          | :heavy_check_mark:                                                                                               | The object type, which is always `thread.message.file`.                                                          |