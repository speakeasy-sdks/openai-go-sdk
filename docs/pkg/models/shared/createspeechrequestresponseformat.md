# CreateSpeechRequestResponseFormat

The format to return audio in. 
Supported formats are `mp3`, `opus`, `aac`, `flac`, `pcm`, and `wav`. 

The `pcm` audio format, similar to `wav` but without a header, utilizes a 24kHz sample rate, mono channel, and 16-bit depth in signed little-endian format.


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `CreateSpeechRequestResponseFormatMp3`  | mp3                                     |
| `CreateSpeechRequestResponseFormatOpus` | opus                                    |
| `CreateSpeechRequestResponseFormatAac`  | aac                                     |
| `CreateSpeechRequestResponseFormatFlac` | flac                                    |
| `CreateSpeechRequestResponseFormatPcm`  | pcm                                     |
| `CreateSpeechRequestResponseFormatWav`  | wav                                     |