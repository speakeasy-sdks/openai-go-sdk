# Results


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Categories`                                                            | [shared.Categories](../../../pkg/models/shared/categories.md)           | :heavy_check_mark:                                                      | A list of the categories, and whether they are flagged or not.          |
| `CategoryScores`                                                        | [shared.CategoryScores](../../../pkg/models/shared/categoryscores.md)   | :heavy_check_mark:                                                      | A list of the categories along with their scores as predicted by model. |
| `Flagged`                                                               | *bool*                                                                  | :heavy_check_mark:                                                      | Whether any of the below categories are flagged.                        |