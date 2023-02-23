//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredisenterprise

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
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
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByClusterPager - Gets the private link resources that need to be created for a RedisEnterprise cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the RedisEnterprise cluster.
// options - PrivateLinkResourcesClientListByClusterOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByCluster
// method.
func (client *PrivateLinkResourcesClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *PrivateLinkResourcesClientListByClusterOptions) *runtime.Pager[PrivateLinkResourcesClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListByClusterResponse]{
		More: func(page PrivateLinkResourcesClientListByClusterResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListByClusterResponse) (PrivateLinkResourcesClientListByClusterResponse, error) {
			req, err := client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			if err != nil {
				return PrivateLinkResourcesClientListByClusterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateLinkResourcesClientListByClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkResourcesClientListByClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterHandleResponse(resp)
		},
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *PrivateLinkResourcesClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *PrivateLinkResourcesClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/privateLinkResources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *PrivateLinkResourcesClient) listByClusterHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByClusterResponse, error) {
	result := PrivateLinkResourcesClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByClusterResponse{}, err
	}
	return result, nil
}
