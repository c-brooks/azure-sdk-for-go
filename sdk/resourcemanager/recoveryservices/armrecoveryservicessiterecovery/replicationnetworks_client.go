//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery

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

// ReplicationNetworksClient contains the methods for the ReplicationNetworks group.
// Don't use this type directly, use NewReplicationNetworksClient() instead.
type ReplicationNetworksClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationNetworksClient creates a new instance of ReplicationNetworksClient with the specified values.
// resourceName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReplicationNetworksClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationNetworksClient, error) {
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
	client := &ReplicationNetworksClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Get - Gets the details of a network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// fabricName - Server Id.
// networkName - Primary network name.
// options - ReplicationNetworksClientGetOptions contains the optional parameters for the ReplicationNetworksClient.Get method.
func (client *ReplicationNetworksClient) Get(ctx context.Context, fabricName string, networkName string, options *ReplicationNetworksClientGetOptions) (ReplicationNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, networkName, options)
	if err != nil {
		return ReplicationNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationNetworksClient) getCreateRequest(ctx context.Context, fabricName string, networkName string, options *ReplicationNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationNetworksClient) getHandleResponse(resp *http.Response) (ReplicationNetworksClientGetResponse, error) {
	result := ReplicationNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Network); err != nil {
		return ReplicationNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the networks available in a vault.
// Generated from API version 2022-10-01
// options - ReplicationNetworksClientListOptions contains the optional parameters for the ReplicationNetworksClient.List
// method.
func (client *ReplicationNetworksClient) NewListPager(options *ReplicationNetworksClientListOptions) *runtime.Pager[ReplicationNetworksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationNetworksClientListResponse]{
		More: func(page ReplicationNetworksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationNetworksClientListResponse) (ReplicationNetworksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationNetworksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationNetworksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationNetworksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationNetworksClient) listCreateRequest(ctx context.Context, options *ReplicationNetworksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworks"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationNetworksClient) listHandleResponse(resp *http.Response) (ReplicationNetworksClientListResponse, error) {
	result := ReplicationNetworksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkCollection); err != nil {
		return ReplicationNetworksClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationFabricsPager - Lists the networks available for a fabric.
// Generated from API version 2022-10-01
// fabricName - Fabric name.
// options - ReplicationNetworksClientListByReplicationFabricsOptions contains the optional parameters for the ReplicationNetworksClient.ListByReplicationFabrics
// method.
func (client *ReplicationNetworksClient) NewListByReplicationFabricsPager(fabricName string, options *ReplicationNetworksClientListByReplicationFabricsOptions) *runtime.Pager[ReplicationNetworksClientListByReplicationFabricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationNetworksClientListByReplicationFabricsResponse]{
		More: func(page ReplicationNetworksClientListByReplicationFabricsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationNetworksClientListByReplicationFabricsResponse) (ReplicationNetworksClientListByReplicationFabricsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationFabricsCreateRequest(ctx, fabricName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationNetworksClientListByReplicationFabricsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationNetworksClientListByReplicationFabricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationNetworksClientListByReplicationFabricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationFabricsHandleResponse(resp)
		},
	})
}

// listByReplicationFabricsCreateRequest creates the ListByReplicationFabrics request.
func (client *ReplicationNetworksClient) listByReplicationFabricsCreateRequest(ctx context.Context, fabricName string, options *ReplicationNetworksClientListByReplicationFabricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationFabricsHandleResponse handles the ListByReplicationFabrics response.
func (client *ReplicationNetworksClient) listByReplicationFabricsHandleResponse(resp *http.Response) (ReplicationNetworksClientListByReplicationFabricsResponse, error) {
	result := ReplicationNetworksClientListByReplicationFabricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkCollection); err != nil {
		return ReplicationNetworksClientListByReplicationFabricsResponse{}, err
	}
	return result, nil
}
