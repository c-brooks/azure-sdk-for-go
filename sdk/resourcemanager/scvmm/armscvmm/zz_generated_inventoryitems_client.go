//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

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

// InventoryItemsClient contains the methods for the InventoryItems group.
// Don't use this type directly, use NewInventoryItemsClient() instead.
type InventoryItemsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInventoryItemsClient creates a new instance of InventoryItemsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInventoryItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InventoryItemsClient, error) {
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
	client := &InventoryItemsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create Or Update InventoryItem.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// vmmServerName - Name of the VMMServer.
// inventoryItemName - Name of the inventoryItem.
// options - InventoryItemsClientCreateOptions contains the optional parameters for the InventoryItemsClient.Create method.
func (client *InventoryItemsClient) Create(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *InventoryItemsClientCreateOptions) (InventoryItemsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, vmmServerName, inventoryItemName, options)
	if err != nil {
		return InventoryItemsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InventoryItemsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InventoryItemsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *InventoryItemsClient) createCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *InventoryItemsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}/inventoryItems/{inventoryItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	if inventoryItemName == "" {
		return nil, errors.New("parameter inventoryItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inventoryItemName}", url.PathEscape(inventoryItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *InventoryItemsClient) createHandleResponse(resp *http.Response) (InventoryItemsClientCreateResponse, error) {
	result := InventoryItemsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InventoryItem); err != nil {
		return InventoryItemsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an inventoryItem.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// vmmServerName - Name of the VMMServer.
// inventoryItemName - Name of the inventoryItem.
// options - InventoryItemsClientDeleteOptions contains the optional parameters for the InventoryItemsClient.Delete method.
func (client *InventoryItemsClient) Delete(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *InventoryItemsClientDeleteOptions) (InventoryItemsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmmServerName, inventoryItemName, options)
	if err != nil {
		return InventoryItemsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InventoryItemsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return InventoryItemsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return InventoryItemsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InventoryItemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *InventoryItemsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}/inventoryItems/{inventoryItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	if inventoryItemName == "" {
		return nil, errors.New("parameter inventoryItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inventoryItemName}", url.PathEscape(inventoryItemName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Shows an inventory item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// vmmServerName - Name of the VMMServer.
// inventoryItemName - Name of the inventoryItem.
// options - InventoryItemsClientGetOptions contains the optional parameters for the InventoryItemsClient.Get method.
func (client *InventoryItemsClient) Get(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *InventoryItemsClientGetOptions) (InventoryItemsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmmServerName, inventoryItemName, options)
	if err != nil {
		return InventoryItemsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InventoryItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InventoryItemsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InventoryItemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, inventoryItemName string, options *InventoryItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}/inventoryItems/{inventoryItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	if inventoryItemName == "" {
		return nil, errors.New("parameter inventoryItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inventoryItemName}", url.PathEscape(inventoryItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InventoryItemsClient) getHandleResponse(resp *http.Response) (InventoryItemsClientGetResponse, error) {
	result := InventoryItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InventoryItem); err != nil {
		return InventoryItemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVMMServerPager - Returns the list of inventoryItems in the given VMMServer.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// vmmServerName - Name of the VMMServer.
// options - InventoryItemsClientListByVMMServerOptions contains the optional parameters for the InventoryItemsClient.ListByVMMServer
// method.
func (client *InventoryItemsClient) NewListByVMMServerPager(resourceGroupName string, vmmServerName string, options *InventoryItemsClientListByVMMServerOptions) *runtime.Pager[InventoryItemsClientListByVMMServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[InventoryItemsClientListByVMMServerResponse]{
		More: func(page InventoryItemsClientListByVMMServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InventoryItemsClientListByVMMServerResponse) (InventoryItemsClientListByVMMServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVMMServerCreateRequest(ctx, resourceGroupName, vmmServerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InventoryItemsClientListByVMMServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InventoryItemsClientListByVMMServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InventoryItemsClientListByVMMServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVMMServerHandleResponse(resp)
		},
	})
}

// listByVMMServerCreateRequest creates the ListByVMMServer request.
func (client *InventoryItemsClient) listByVMMServerCreateRequest(ctx context.Context, resourceGroupName string, vmmServerName string, options *InventoryItemsClientListByVMMServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/vmmServers/{vmmServerName}/inventoryItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmmServerName == "" {
		return nil, errors.New("parameter vmmServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmmServerName}", url.PathEscape(vmmServerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVMMServerHandleResponse handles the ListByVMMServer response.
func (client *InventoryItemsClient) listByVMMServerHandleResponse(resp *http.Response) (InventoryItemsClientListByVMMServerResponse, error) {
	result := InventoryItemsClientListByVMMServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InventoryItemsList); err != nil {
		return InventoryItemsClientListByVMMServerResponse{}, err
	}
	return result, nil
}
