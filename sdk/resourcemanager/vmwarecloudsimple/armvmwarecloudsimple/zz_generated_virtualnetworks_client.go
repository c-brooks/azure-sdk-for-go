//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworksClient, error) {
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
	client := &VirtualNetworksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Return virtual network by its name
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// virtualNetworkName - virtual network id (vsphereId)
// options - VirtualNetworksClientGetOptions contains the optional parameters for the VirtualNetworksClient.Get method.
func (client *VirtualNetworksClient) Get(ctx context.Context, regionID string, pcName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (VirtualNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, regionID, pcName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworksClient) getCreateRequest(ctx context.Context, regionID string, pcName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworksClient) getHandleResponse(resp *http.Response) (VirtualNetworksClientGetResponse, error) {
	result := VirtualNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Return list of virtual networks in location for private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-04-01
// regionID - The region Id (westus, eastus)
// pcName - The private cloud name
// resourcePoolName - Resource pool used to derive vSphere cluster which contains virtual networks
// options - VirtualNetworksClientListOptions contains the optional parameters for the VirtualNetworksClient.List method.
func (client *VirtualNetworksClient) NewListPager(regionID string, pcName string, resourcePoolName string, options *VirtualNetworksClientListOptions) *runtime.Pager[VirtualNetworksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListResponse]{
		More: func(page VirtualNetworksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListResponse) (VirtualNetworksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, regionID, pcName, resourcePoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualNetworksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualNetworksClient) listCreateRequest(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *VirtualNetworksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	reqQP.Set("resourcePoolName", resourcePoolName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworksClient) listHandleResponse(resp *http.Response) (VirtualNetworksClientListResponse, error) {
	result := VirtualNetworksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResponse); err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	return result, nil
}
