# CreateModerationResponseResultsCategories

A list of the categories, and whether they are flagged or not.


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Hate`                                                 | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'hate'.             |
| `HateThreatening`                                      | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'hate/threatening'. |
| `SelfHarm`                                             | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'self-harm'.        |
| `Sexual`                                               | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'sexual'.           |
| `SexualMinors`                                         | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'sexual/minors'.    |
| `Violence`                                             | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'violence'.         |
| `ViolenceGraphic`                                      | *bool*                                                 | :heavy_check_mark:                                     | Whether the content was flagged as 'violence/graphic'. |