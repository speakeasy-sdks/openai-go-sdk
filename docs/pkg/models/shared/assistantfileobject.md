# AssistantFileObject

A list of [Files](/docs/api-reference/files) attached to an `assistant`.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AssistantID`                                                            | *string*                                                                 | :heavy_check_mark:                                                       | The assistant ID that the file is attached to.                           |
| `CreatedAt`                                                              | *int64*                                                                  | :heavy_check_mark:                                                       | The Unix timestamp (in seconds) for when the assistant file was created. |
| `ID`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | The identifier, which can be referenced in API endpoints.                |
| `Object`                                                                 | [shared.Object](../../../pkg/models/shared/object.md)                    | :heavy_check_mark:                                                       | The object type, which is always `assistant.file`.                       |