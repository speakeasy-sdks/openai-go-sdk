package shared

type CreateImageRequestResponseFormatEnum string

const (
	CreateImageRequestResponseFormatEnumURL     CreateImageRequestResponseFormatEnum = "url"
	CreateImageRequestResponseFormatEnumB64JSON CreateImageRequestResponseFormatEnum = "b64_json"
)

type CreateImageRequestSizeEnum string

const (
	CreateImageRequestSizeEnumTwoHundredAndFiftySixx256     CreateImageRequestSizeEnum = "256x256"
	CreateImageRequestSizeEnumFiveHundredAndTwelvex512      CreateImageRequestSizeEnum = "512x512"
	CreateImageRequestSizeEnumOneThousandAndTwentyFourx1024 CreateImageRequestSizeEnum = "1024x1024"
)

type CreateImageRequest struct {
	N              *int64                                `json:"n,omitempty"`
	Prompt         string                                `json:"prompt"`
	ResponseFormat *CreateImageRequestResponseFormatEnum `json:"response_format,omitempty"`
	Size           *CreateImageRequestSizeEnum           `json:"size,omitempty"`
	User           *interface{}                          `json:"user,omitempty"`
}
