package openai

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

type Openai struct {
	OpenAI *openAI

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*Openai)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Openai) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Openai) {
		sdk._defaultClient = client
	}
}

func New(opts ...SDKOption) *Openai {
	sdk := &Openai{
		_language:   "go",
		_sdkVersion: "1.3.1",
		_genVersion: "1.5.4",
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
