# GetMessageFileRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `FileID`                                                   | *string*                                                   | :heavy_check_mark:                                         | The ID of the file being retrieved.                        | file-abc123                                                |
| `MessageID`                                                | *string*                                                   | :heavy_check_mark:                                         | The ID of the message the file belongs to.                 | msg_abc123                                                 |
| `ThreadID`                                                 | *string*                                                   | :heavy_check_mark:                                         | The ID of the thread to which the message and File belong. | thread_abc123                                              |