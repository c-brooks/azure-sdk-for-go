//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"errors"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/c-brooks/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IngestionSettingsClient contains the methods for the IngestionSettings group.
// Don't use this type directly, use NewIngestionSettingsClient() instead.
type IngestionSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIngestionSettingsClient creates a new instance of IngestionSettingsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIngestionSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IngestionSettingsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &IngestionSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create setting for ingesting security data and logs to correlate with resources associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-15-preview
// ingestionSettingName - Name of the ingestion setting
// ingestionSetting - Ingestion setting object
// options - IngestionSettingsClientCreateOptions contains the optional parameters for the IngestionSettingsClient.Create
// method.
func (client *IngestionSettingsClient) Create(ctx context.Context, ingestionSettingName string, ingestionSetting IngestionSetting, options *IngestionSettingsClientCreateOptions) (IngestionSettingsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, ingestionSettingName, ingestionSetting, options)
	if err != nil {
		return IngestionSettingsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IngestionSettingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IngestionSettingsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *IngestionSettingsClient) createCreateRequest(ctx context.Context, ingestionSettingName string, ingestionSetting IngestionSetting, options *IngestionSettingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings/{ingestionSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ingestionSettingName == "" {
		return nil, errors.New("parameter ingestionSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ingestionSettingName}", url.PathEscape(ingestionSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, ingestionSetting)
}

// createHandleResponse handles the Create response.
func (client *IngestionSettingsClient) createHandleResponse(resp *http.Response) (IngestionSettingsClientCreateResponse, error) {
	result := IngestionSettingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IngestionSetting); err != nil {
		return IngestionSettingsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the ingestion settings for this subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-15-preview
// ingestionSettingName - Name of the ingestion setting
// options - IngestionSettingsClientDeleteOptions contains the optional parameters for the IngestionSettingsClient.Delete
// method.
func (client *IngestionSettingsClient) Delete(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientDeleteOptions) (IngestionSettingsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, ingestionSettingName, options)
	if err != nil {
		return IngestionSettingsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IngestionSettingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IngestionSettingsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IngestionSettingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IngestionSettingsClient) deleteCreateRequest(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings/{ingestionSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ingestionSettingName == "" {
		return nil, errors.New("parameter ingestionSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ingestionSettingName}", url.PathEscape(ingestionSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Settings for ingesting security data and logs to correlate with resources associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-15-preview
// ingestionSettingName - Name of the ingestion setting
// options - IngestionSettingsClientGetOptions contains the optional parameters for the IngestionSettingsClient.Get method.
func (client *IngestionSettingsClient) Get(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientGetOptions) (IngestionSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, ingestionSettingName, options)
	if err != nil {
		return IngestionSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IngestionSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IngestionSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IngestionSettingsClient) getCreateRequest(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings/{ingestionSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ingestionSettingName == "" {
		return nil, errors.New("parameter ingestionSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ingestionSettingName}", url.PathEscape(ingestionSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IngestionSettingsClient) getHandleResponse(resp *http.Response) (IngestionSettingsClientGetResponse, error) {
	result := IngestionSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IngestionSetting); err != nil {
		return IngestionSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Settings for ingesting security data and logs to correlate with resources associated with the subscription.
// Generated from API version 2021-01-15-preview
// options - IngestionSettingsClientListOptions contains the optional parameters for the IngestionSettingsClient.List method.
func (client *IngestionSettingsClient) NewListPager(options *IngestionSettingsClientListOptions) *runtime.Pager[IngestionSettingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IngestionSettingsClientListResponse]{
		More: func(page IngestionSettingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IngestionSettingsClientListResponse) (IngestionSettingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IngestionSettingsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IngestionSettingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IngestionSettingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IngestionSettingsClient) listCreateRequest(ctx context.Context, options *IngestionSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IngestionSettingsClient) listHandleResponse(resp *http.Response) (IngestionSettingsClientListResponse, error) {
	result := IngestionSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IngestionSettingList); err != nil {
		return IngestionSettingsClientListResponse{}, err
	}
	return result, nil
}

// ListConnectionStrings - Connection strings for ingesting security scan logs and data.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-15-preview
// ingestionSettingName - Name of the ingestion setting
// options - IngestionSettingsClientListConnectionStringsOptions contains the optional parameters for the IngestionSettingsClient.ListConnectionStrings
// method.
func (client *IngestionSettingsClient) ListConnectionStrings(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientListConnectionStringsOptions) (IngestionSettingsClientListConnectionStringsResponse, error) {
	req, err := client.listConnectionStringsCreateRequest(ctx, ingestionSettingName, options)
	if err != nil {
		return IngestionSettingsClientListConnectionStringsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IngestionSettingsClientListConnectionStringsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IngestionSettingsClientListConnectionStringsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listConnectionStringsHandleResponse(resp)
}

// listConnectionStringsCreateRequest creates the ListConnectionStrings request.
func (client *IngestionSettingsClient) listConnectionStringsCreateRequest(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientListConnectionStringsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings/{ingestionSettingName}/listConnectionStrings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ingestionSettingName == "" {
		return nil, errors.New("parameter ingestionSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ingestionSettingName}", url.PathEscape(ingestionSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConnectionStringsHandleResponse handles the ListConnectionStrings response.
func (client *IngestionSettingsClient) listConnectionStringsHandleResponse(resp *http.Response) (IngestionSettingsClientListConnectionStringsResponse, error) {
	result := IngestionSettingsClientListConnectionStringsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionStrings); err != nil {
		return IngestionSettingsClientListConnectionStringsResponse{}, err
	}
	return result, nil
}

// ListTokens - Returns the token that is used for correlating ingested telemetry with the resources in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-01-15-preview
// ingestionSettingName - Name of the ingestion setting
// options - IngestionSettingsClientListTokensOptions contains the optional parameters for the IngestionSettingsClient.ListTokens
// method.
func (client *IngestionSettingsClient) ListTokens(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientListTokensOptions) (IngestionSettingsClientListTokensResponse, error) {
	req, err := client.listTokensCreateRequest(ctx, ingestionSettingName, options)
	if err != nil {
		return IngestionSettingsClientListTokensResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IngestionSettingsClientListTokensResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IngestionSettingsClientListTokensResponse{}, runtime.NewResponseError(resp)
	}
	return client.listTokensHandleResponse(resp)
}

// listTokensCreateRequest creates the ListTokens request.
func (client *IngestionSettingsClient) listTokensCreateRequest(ctx context.Context, ingestionSettingName string, options *IngestionSettingsClientListTokensOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/ingestionSettings/{ingestionSettingName}/listTokens"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ingestionSettingName == "" {
		return nil, errors.New("parameter ingestionSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ingestionSettingName}", url.PathEscape(ingestionSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listTokensHandleResponse handles the ListTokens response.
func (client *IngestionSettingsClient) listTokensHandleResponse(resp *http.Response) (IngestionSettingsClientListTokensResponse, error) {
	result := IngestionSettingsClientListTokensResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IngestionSettingToken); err != nil {
		return IngestionSettingsClientListTokensResponse{}, err
	}
	return result, nil
}
