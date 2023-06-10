# FineTune

OK


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `CreatedAt`                                                       | *int64*                                                           | :heavy_check_mark:                                                | N/A                                                               |
| `Events`                                                          | [][FineTuneEvent](../../models/shared/finetuneevent.md)           | :heavy_minus_sign:                                                | N/A                                                               |
| `FineTunedModel`                                                  | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `Hyperparams`                                                     | [FineTuneHyperparams](../../models/shared/finetunehyperparams.md) | :heavy_check_mark:                                                | N/A                                                               |
| `ID`                                                              | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `Model`                                                           | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `Object`                                                          | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `OrganizationID`                                                  | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `ResultFiles`                                                     | [][OpenAIFile](../../models/shared/openaifile.md)                 | :heavy_check_mark:                                                | N/A                                                               |
| `Status`                                                          | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `TrainingFiles`                                                   | [][OpenAIFile](../../models/shared/openaifile.md)                 | :heavy_check_mark:                                                | N/A                                                               |
| `UpdatedAt`                                                       | *int64*                                                           | :heavy_check_mark:                                                | N/A                                                               |
| `ValidationFiles`                                                 | [][OpenAIFile](../../models/shared/openaifile.md)                 | :heavy_check_mark:                                                | N/A                                                               |