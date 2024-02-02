// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package openaigosdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v4/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://api.openai.com/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Gpt - OpenAI API: The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
type Gpt struct {
	// Build Assistants that can call models and use tools.
	Assistants *Assistants
	// Learn how to turn audio into text or text into audio.
	Audio *Audio
	// Given a list of messages comprising a conversation, the model will return a response.
	Chat *Chat
	// Given a prompt, the model will return one or more predicted completions, and can also return the probabilities of alternative tokens at each position.
	Completions *Completions
	// Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.
	Embeddings *Embeddings
	// Files are used to upload documents that can be used with features like Assistants and Fine-tuning.
	Files *Files
	// Manage fine-tuning jobs to tailor a model to your specific training data.
	FineTuning *FineTuning
	// Given a prompt and/or an input image, the model will generate a new image.
	Images *Images
	// List and describe the various models available in the API.
	Models *Models
	// Given a input text, outputs if the model classifies it as violating OpenAI's content policy.
	Moderations *Moderations

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Gpt)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Gpt) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Gpt) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Gpt) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Gpt) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKeyAuth string) SDKOption {
	return func(sdk *Gpt) {
		security := shared.Security{APIKeyAuth: apiKeyAuth}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Gpt) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Gpt) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Gpt {
	sdk := &Gpt{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.0.0",
			SDKVersion:        "4.0.0",
			GenVersion:        "2.248.1",
			UserAgent:         "speakeasy-sdk/go 4.0.0 2.248.1 2.0.0 github.com/speakeasy-sdks/openai-go-sdk",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Assistants = newAssistants(sdk.sdkConfiguration)

	sdk.Audio = newAudio(sdk.sdkConfiguration)

	sdk.Chat = newChat(sdk.sdkConfiguration)

	sdk.Completions = newCompletions(sdk.sdkConfiguration)

	sdk.Embeddings = newEmbeddings(sdk.sdkConfiguration)

	sdk.Files = newFiles(sdk.sdkConfiguration)

	sdk.FineTuning = newFineTuning(sdk.sdkConfiguration)

	sdk.Images = newImages(sdk.sdkConfiguration)

	sdk.Models = newModels(sdk.sdkConfiguration)

	sdk.Moderations = newModerations(sdk.sdkConfiguration)

	return sdk
}
