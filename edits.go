// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package openaigosdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/openai-go-sdk/v3/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// Edits - Given a prompt and an instruction, the model will return an edited version of the prompt.
type Edits struct {
	sdkConfiguration sdkConfiguration
}

func newEdits(sdkConfig sdkConfiguration) *Edits {
	return &Edits{
		sdkConfiguration: sdkConfig,
	}
}

// CreateEdit - Creates a new edit for the provided input, instruction, and parameters.
//
// Deprecated method: This will be removed in a future release, please migrate away from it as soon as possible.
func (s *Edits) CreateEdit(ctx context.Context, request shared.CreateEditRequest) (*operations.CreateEditResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/edits"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateEditResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.CreateEditResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CreateEditResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
