# RunStepDetailsToolCallsRetrievalObject


## Fields

| Field                                                                                                                         | Type                                                                                                                          | Required                                                                                                                      | Description                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                                          | *string*                                                                                                                      | :heavy_check_mark:                                                                                                            | The ID of the tool call object.                                                                                               |
| `Retrieval`                                                                                                                   | [shared.Retrieval](../../../pkg/models/shared/retrieval.md)                                                                   | :heavy_check_mark:                                                                                                            | For now, this is always going to be an empty object.                                                                          |
| `Type`                                                                                                                        | [shared.RunStepDetailsToolCallsRetrievalObjectType](../../../pkg/models/shared/runstepdetailstoolcallsretrievalobjecttype.md) | :heavy_check_mark:                                                                                                            | The type of tool call. This is always going to be `retrieval` for this type of tool call.                                     |