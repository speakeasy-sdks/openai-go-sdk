# CreateEditResponse

OK


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Choices`                                                                       | [][CreateEditResponseChoices](../../models/shared/createeditresponsechoices.md) | :heavy_check_mark:                                                              | A list of edit choices. Can be more than one if `n` is greater than 1.          |
| `Created`                                                                       | *int64*                                                                         | :heavy_check_mark:                                                              | A unix timestamp of when the edit was created.                                  |
| `Object`                                                                        | *string*                                                                        | :heavy_check_mark:                                                              | The object type, which is always `edit`.                                        |
| `Usage`                                                                         | [CompletionUsage](../../models/shared/completionusage.md)                       | :heavy_check_mark:                                                              | Usage statistics for the completion request.                                    |