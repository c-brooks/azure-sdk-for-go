//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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

// KeyValuesClient contains the methods for the KeyValues group.
// Don't use this type directly, use NewKeyValuesClient() instead.
type KeyValuesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewKeyValuesClient creates a new instance of KeyValuesClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewKeyValuesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KeyValuesClient, error) {
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
	client := &KeyValuesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a key-value.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// keyValueName - Identifier of key and label combination. Key and label are joined by $ character. Label is optional.
// options - KeyValuesClientCreateOrUpdateOptions contains the optional parameters for the KeyValuesClient.CreateOrUpdate
// method.
func (client *KeyValuesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientCreateOrUpdateOptions) (KeyValuesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, configStoreName, keyValueName, options)
	if err != nil {
		return KeyValuesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KeyValuesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KeyValuesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *KeyValuesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues/{keyValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if keyValueName == "" {
		return nil, errors.New("parameter keyValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyValueName}", url.PathEscape(keyValueName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.KeyValueParameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.KeyValueParameters)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *KeyValuesClient) createOrUpdateHandleResponse(resp *http.Response) (KeyValuesClientCreateOrUpdateResponse, error) {
	result := KeyValuesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyValue); err != nil {
		return KeyValuesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a key-value.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// keyValueName - Identifier of key and label combination. Key and label are joined by $ character. Label is optional.
// options - KeyValuesClientBeginDeleteOptions contains the optional parameters for the KeyValuesClient.BeginDelete method.
func (client *KeyValuesClient) BeginDelete(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientBeginDeleteOptions) (*runtime.Poller[KeyValuesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, configStoreName, keyValueName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[KeyValuesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[KeyValuesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a key-value.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *KeyValuesClient) deleteOperation(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configStoreName, keyValueName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *KeyValuesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues/{keyValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if keyValueName == "" {
		return nil, errors.New("parameter keyValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyValueName}", url.PathEscape(keyValueName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified key-value.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// keyValueName - Identifier of key and label combination. Key and label are joined by $ character. Label is optional.
// options - KeyValuesClientGetOptions contains the optional parameters for the KeyValuesClient.Get method.
func (client *KeyValuesClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientGetOptions) (KeyValuesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, configStoreName, keyValueName, options)
	if err != nil {
		return KeyValuesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KeyValuesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KeyValuesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *KeyValuesClient) getCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, keyValueName string, options *KeyValuesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues/{keyValueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if keyValueName == "" {
		return nil, errors.New("parameter keyValueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyValueName}", url.PathEscape(keyValueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KeyValuesClient) getHandleResponse(resp *http.Response) (KeyValuesClientGetResponse, error) {
	result := KeyValuesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyValue); err != nil {
		return KeyValuesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByConfigurationStorePager - Lists the key-values for a given configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group to which the container registry belongs.
// configStoreName - The name of the configuration store.
// options - KeyValuesClientListByConfigurationStoreOptions contains the optional parameters for the KeyValuesClient.ListByConfigurationStore
// method.
func (client *KeyValuesClient) NewListByConfigurationStorePager(resourceGroupName string, configStoreName string, options *KeyValuesClientListByConfigurationStoreOptions) *runtime.Pager[KeyValuesClientListByConfigurationStoreResponse] {
	return runtime.NewPager(runtime.PagingHandler[KeyValuesClientListByConfigurationStoreResponse]{
		More: func(page KeyValuesClientListByConfigurationStoreResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *KeyValuesClientListByConfigurationStoreResponse) (KeyValuesClientListByConfigurationStoreResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByConfigurationStoreCreateRequest(ctx, resourceGroupName, configStoreName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return KeyValuesClientListByConfigurationStoreResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return KeyValuesClientListByConfigurationStoreResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return KeyValuesClientListByConfigurationStoreResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConfigurationStoreHandleResponse(resp)
		},
	})
}

// listByConfigurationStoreCreateRequest creates the ListByConfigurationStore request.
func (client *KeyValuesClient) listByConfigurationStoreCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *KeyValuesClientListByConfigurationStoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/keyValues"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByConfigurationStoreHandleResponse handles the ListByConfigurationStore response.
func (client *KeyValuesClient) listByConfigurationStoreHandleResponse(resp *http.Response) (KeyValuesClientListByConfigurationStoreResponse, error) {
	result := KeyValuesClientListByConfigurationStoreResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyValueListResult); err != nil {
		return KeyValuesClientListByConfigurationStoreResponse{}, err
	}
	return result, nil
}
