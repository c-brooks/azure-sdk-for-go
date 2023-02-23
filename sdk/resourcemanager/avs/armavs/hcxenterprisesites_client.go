//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armavs

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

// HcxEnterpriseSitesClient contains the methods for the HcxEnterpriseSites group.
// Don't use this type directly, use NewHcxEnterpriseSitesClient() instead.
type HcxEnterpriseSitesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHcxEnterpriseSitesClient creates a new instance of HcxEnterpriseSitesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHcxEnterpriseSitesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HcxEnterpriseSitesClient, error) {
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
	client := &HcxEnterpriseSitesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an HCX Enterprise Site in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - The name of the private cloud.
// hcxEnterpriseSiteName - Name of the HCX Enterprise Site in the private cloud
// hcxEnterpriseSite - The HCX Enterprise Site
// options - HcxEnterpriseSitesClientCreateOrUpdateOptions contains the optional parameters for the HcxEnterpriseSitesClient.CreateOrUpdate
// method.
func (client *HcxEnterpriseSitesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, hcxEnterpriseSite HcxEnterpriseSite, options *HcxEnterpriseSitesClientCreateOrUpdateOptions) (HcxEnterpriseSitesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, hcxEnterpriseSiteName, hcxEnterpriseSite, options)
	if err != nil {
		return HcxEnterpriseSitesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HcxEnterpriseSitesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HcxEnterpriseSitesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HcxEnterpriseSitesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, hcxEnterpriseSite HcxEnterpriseSite, options *HcxEnterpriseSitesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites/{hcxEnterpriseSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if hcxEnterpriseSiteName == "" {
		return nil, errors.New("parameter hcxEnterpriseSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcxEnterpriseSiteName}", url.PathEscape(hcxEnterpriseSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, hcxEnterpriseSite)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HcxEnterpriseSitesClient) createOrUpdateHandleResponse(resp *http.Response) (HcxEnterpriseSitesClientCreateOrUpdateResponse, error) {
	result := HcxEnterpriseSitesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcxEnterpriseSite); err != nil {
		return HcxEnterpriseSitesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an HCX Enterprise Site in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// hcxEnterpriseSiteName - Name of the HCX Enterprise Site in the private cloud
// options - HcxEnterpriseSitesClientDeleteOptions contains the optional parameters for the HcxEnterpriseSitesClient.Delete
// method.
func (client *HcxEnterpriseSitesClient) Delete(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesClientDeleteOptions) (HcxEnterpriseSitesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, hcxEnterpriseSiteName, options)
	if err != nil {
		return HcxEnterpriseSitesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HcxEnterpriseSitesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HcxEnterpriseSitesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HcxEnterpriseSitesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HcxEnterpriseSitesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites/{hcxEnterpriseSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if hcxEnterpriseSiteName == "" {
		return nil, errors.New("parameter hcxEnterpriseSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcxEnterpriseSiteName}", url.PathEscape(hcxEnterpriseSiteName))
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

// Get - Get an HCX Enterprise Site by name in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// hcxEnterpriseSiteName - Name of the HCX Enterprise Site in the private cloud
// options - HcxEnterpriseSitesClientGetOptions contains the optional parameters for the HcxEnterpriseSitesClient.Get method.
func (client *HcxEnterpriseSitesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesClientGetOptions) (HcxEnterpriseSitesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, hcxEnterpriseSiteName, options)
	if err != nil {
		return HcxEnterpriseSitesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HcxEnterpriseSitesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HcxEnterpriseSitesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HcxEnterpriseSitesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, hcxEnterpriseSiteName string, options *HcxEnterpriseSitesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites/{hcxEnterpriseSiteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if hcxEnterpriseSiteName == "" {
		return nil, errors.New("parameter hcxEnterpriseSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcxEnterpriseSiteName}", url.PathEscape(hcxEnterpriseSiteName))
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
func (client *HcxEnterpriseSitesClient) getHandleResponse(resp *http.Response) (HcxEnterpriseSitesClientGetResponse, error) {
	result := HcxEnterpriseSitesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcxEnterpriseSite); err != nil {
		return HcxEnterpriseSitesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List HCX Enterprise Sites in a private cloud
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - HcxEnterpriseSitesClientListOptions contains the optional parameters for the HcxEnterpriseSitesClient.List method.
func (client *HcxEnterpriseSitesClient) NewListPager(resourceGroupName string, privateCloudName string, options *HcxEnterpriseSitesClientListOptions) *runtime.Pager[HcxEnterpriseSitesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HcxEnterpriseSitesClientListResponse]{
		More: func(page HcxEnterpriseSitesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HcxEnterpriseSitesClientListResponse) (HcxEnterpriseSitesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HcxEnterpriseSitesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HcxEnterpriseSitesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HcxEnterpriseSitesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HcxEnterpriseSitesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *HcxEnterpriseSitesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/hcxEnterpriseSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
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

// listHandleResponse handles the List response.
func (client *HcxEnterpriseSitesClient) listHandleResponse(resp *http.Response) (HcxEnterpriseSitesClientListResponse, error) {
	result := HcxEnterpriseSitesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcxEnterpriseSiteList); err != nil {
		return HcxEnterpriseSitesClientListResponse{}, err
	}
	return result, nil
}
