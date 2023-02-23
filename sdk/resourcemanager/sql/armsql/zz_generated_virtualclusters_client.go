//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// VirtualClustersClient contains the methods for the VirtualClusters group.
// Don't use this type directly, use NewVirtualClustersClient() instead.
type VirtualClustersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualClustersClient creates a new instance of VirtualClustersClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualClustersClient, error) {
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
	client := &VirtualClustersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Deletes a virtual cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// virtualClusterName - The name of the virtual cluster.
// options - VirtualClustersClientBeginDeleteOptions contains the optional parameters for the VirtualClustersClient.BeginDelete
// method.
func (client *VirtualClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientBeginDeleteOptions) (*runtime.Poller[VirtualClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualClusterName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualClustersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualClustersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a virtual cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
func (client *VirtualClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualClusterName, options)
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
func (client *VirtualClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/virtualClusters/{virtualClusterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualClusterName == "" {
		return nil, errors.New("parameter virtualClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualClusterName}", url.PathEscape(virtualClusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a virtual cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// virtualClusterName - The name of the virtual cluster.
// options - VirtualClustersClientGetOptions contains the optional parameters for the VirtualClustersClient.Get method.
func (client *VirtualClustersClient) Get(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientGetOptions) (VirtualClustersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualClusterName, options)
	if err != nil {
		return VirtualClustersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualClustersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/virtualClusters/{virtualClusterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualClusterName == "" {
		return nil, errors.New("parameter virtualClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualClusterName}", url.PathEscape(virtualClusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualClustersClient) getHandleResponse(resp *http.Response) (VirtualClustersClientGetResponse, error) {
	result := VirtualClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualCluster); err != nil {
		return VirtualClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all virtualClusters in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// options - VirtualClustersClientListOptions contains the optional parameters for the VirtualClustersClient.List method.
func (client *VirtualClustersClient) NewListPager(options *VirtualClustersClientListOptions) *runtime.Pager[VirtualClustersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualClustersClientListResponse]{
		More: func(page VirtualClustersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualClustersClientListResponse) (VirtualClustersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualClustersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualClustersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualClustersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualClustersClient) listCreateRequest(ctx context.Context, options *VirtualClustersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/virtualClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualClustersClient) listHandleResponse(resp *http.Response) (VirtualClustersClientListResponse, error) {
	result := VirtualClustersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualClusterListResult); err != nil {
		return VirtualClustersClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of virtual clusters in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - VirtualClustersClientListByResourceGroupOptions contains the optional parameters for the VirtualClustersClient.ListByResourceGroup
// method.
func (client *VirtualClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualClustersClientListByResourceGroupOptions) *runtime.Pager[VirtualClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualClustersClientListByResourceGroupResponse]{
		More: func(page VirtualClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualClustersClientListByResourceGroupResponse) (VirtualClustersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualClustersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualClustersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/virtualClusters"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualClustersClientListByResourceGroupResponse, error) {
	result := VirtualClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualClusterListResult); err != nil {
		return VirtualClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a virtual cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// virtualClusterName - The name of the virtual cluster.
// parameters - The requested virtual cluster resource state.
// options - VirtualClustersClientBeginUpdateOptions contains the optional parameters for the VirtualClustersClient.BeginUpdate
// method.
func (client *VirtualClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, virtualClusterName string, parameters VirtualClusterUpdate, options *VirtualClustersClientBeginUpdateOptions) (*runtime.Poller[VirtualClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, virtualClusterName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualClustersClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualClustersClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a virtual cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
func (client *VirtualClustersClient) update(ctx context.Context, resourceGroupName string, virtualClusterName string, parameters VirtualClusterUpdate, options *VirtualClustersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualClusterName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualClusterName string, parameters VirtualClusterUpdate, options *VirtualClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/virtualClusters/{virtualClusterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualClusterName == "" {
		return nil, errors.New("parameter virtualClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualClusterName}", url.PathEscape(virtualClusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// UpdateDNSServers - Synchronizes the DNS server settings used by the managed instances inside the given virtual cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// virtualClusterName - The name of the virtual cluster.
// options - VirtualClustersClientUpdateDNSServersOptions contains the optional parameters for the VirtualClustersClient.UpdateDNSServers
// method.
func (client *VirtualClustersClient) UpdateDNSServers(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientUpdateDNSServersOptions) (VirtualClustersClientUpdateDNSServersResponse, error) {
	req, err := client.updateDNSServersCreateRequest(ctx, resourceGroupName, virtualClusterName, options)
	if err != nil {
		return VirtualClustersClientUpdateDNSServersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualClustersClientUpdateDNSServersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualClustersClientUpdateDNSServersResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateDNSServersHandleResponse(resp)
}

// updateDNSServersCreateRequest creates the UpdateDNSServers request.
func (client *VirtualClustersClient) updateDNSServersCreateRequest(ctx context.Context, resourceGroupName string, virtualClusterName string, options *VirtualClustersClientUpdateDNSServersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/virtualClusters/{virtualClusterName}/updateManagedInstanceDnsServers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualClusterName == "" {
		return nil, errors.New("parameter virtualClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualClusterName}", url.PathEscape(virtualClusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// updateDNSServersHandleResponse handles the UpdateDNSServers response.
func (client *VirtualClustersClient) updateDNSServersHandleResponse(resp *http.Response) (VirtualClustersClientUpdateDNSServersResponse, error) {
	result := VirtualClustersClientUpdateDNSServersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateManagedInstanceDNSServersOperation); err != nil {
		return VirtualClustersClientUpdateDNSServersResponse{}, err
	}
	return result, nil
}
