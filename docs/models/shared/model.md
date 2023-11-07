# Model

Describes an OpenAI model offering that can be used with the API.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Created`                                                           | *int64*                                                             | :heavy_check_mark:                                                  | The Unix timestamp (in seconds) when the model was created.         |
| `ID`                                                                | *string*                                                            | :heavy_check_mark:                                                  | The model identifier, which can be referenced in the API endpoints. |
| `Object`                                                            | [shared.ModelObject](../../models/shared/modelobject.md)            | :heavy_check_mark:                                                  | The object type, which is always "model".                           |
| `OwnedBy`                                                           | *string*                                                            | :heavy_check_mark:                                                  | The organization that owns the model.                               |