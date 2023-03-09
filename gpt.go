package gpt

import (
	"github.com/speakeasy-sdks/openai-go-sdk/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://api.openai.com/v1",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// SDK Documentation: APIs for sampling from and fine-tuning language models
type Gpt struct {
	OpenAI *openAI

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*Gpt)

func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Gpt) {
		sdk._serverURL = serverURL
	}
}

func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Gpt) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Gpt) {
		sdk._defaultClient = client
	}
}

func New(opts ...SDKOption) *Gpt {
	sdk := &Gpt{
		_language:   "go",
		_sdkVersion: "1.6.1",
		_genVersion: "1.8.7",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		sdk._securityClient = sdk._defaultClient

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.OpenAI = newOpenAI(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
