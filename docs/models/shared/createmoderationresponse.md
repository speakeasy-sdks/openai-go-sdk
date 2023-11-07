# CreateModerationResponse

Represents policy compliance report by OpenAI's content moderation model against a given input.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ID`                                               | *string*                                           | :heavy_check_mark:                                 | The unique identifier for the moderation request.  |
| `Model`                                            | *string*                                           | :heavy_check_mark:                                 | The model used to generate the moderation results. |
| `Results`                                          | [][shared.Results](../../models/shared/results.md) | :heavy_check_mark:                                 | A list of moderation objects.                      |