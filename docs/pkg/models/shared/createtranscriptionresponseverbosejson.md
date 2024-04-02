# CreateTranscriptionResponseVerboseJSON

Represents a verbose json transcription response returned by model, based on the provided input.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Duration`                                                                          | *string*                                                                            | :heavy_check_mark:                                                                  | The duration of the input audio.                                                    |
| `Language`                                                                          | *string*                                                                            | :heavy_check_mark:                                                                  | The language of the input audio.                                                    |
| `Segments`                                                                          | [][shared.TranscriptionSegment](../../../pkg/models/shared/transcriptionsegment.md) | :heavy_minus_sign:                                                                  | Segments of the transcribed text and their corresponding details.                   |
| `Text`                                                                              | *string*                                                                            | :heavy_check_mark:                                                                  | The transcribed text.                                                               |
| `Words`                                                                             | [][shared.TranscriptionWord](../../../pkg/models/shared/transcriptionword.md)       | :heavy_minus_sign:                                                                  | Extracted words and their corresponding timestamps.                                 |