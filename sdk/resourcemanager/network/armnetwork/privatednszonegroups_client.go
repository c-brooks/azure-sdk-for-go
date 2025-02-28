//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateDNSZoneGroupsClient contains the methods for the PrivateDNSZoneGroups group.
// Don't use this type directly, use NewPrivateDNSZoneGroupsClient() instead.
type PrivateDNSZoneGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateDNSZoneGroupsClient creates a new instance of PrivateDNSZoneGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateDNSZoneGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateDNSZoneGroupsClient, error) {
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
	client := &PrivateDNSZoneGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// privateEndpointName - The name of the private endpoint.
// privateDNSZoneGroupName - The name of the private dns zone group.
// parameters - Parameters supplied to the create or update private dns zone group operation.
// options - PrivateDNSZoneGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateDNSZoneGroupsClient.BeginCreateOrUpdate
// method.
func (client *PrivateDNSZoneGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateDNSZoneGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateDNSZoneGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateDNSZoneGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a private dns zone group in the specified private endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *PrivateDNSZoneGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateDNSZoneGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters PrivateDNSZoneGroup, options *PrivateDNSZoneGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if privateDNSZoneGroupName == "" {
		return nil, errors.New("parameter privateDNSZoneGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDNSZoneGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified private dns zone group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// privateEndpointName - The name of the private endpoint.
// privateDNSZoneGroupName - The name of the private dns zone group.
// options - PrivateDNSZoneGroupsClientBeginDeleteOptions contains the optional parameters for the PrivateDNSZoneGroupsClient.BeginDelete
// method.
func (client *PrivateDNSZoneGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsClientBeginDeleteOptions) (*runtime.Poller[PrivateDNSZoneGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[PrivateDNSZoneGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[PrivateDNSZoneGroupsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified private dns zone group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
func (client *PrivateDNSZoneGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, options)
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
func (client *PrivateDNSZoneGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if privateDNSZoneGroupName == "" {
		return nil, errors.New("parameter privateDNSZoneGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDNSZoneGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the private dns zone group resource by specified private dns zone group name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01
// resourceGroupName - The name of the resource group.
// privateEndpointName - The name of the private endpoint.
// privateDNSZoneGroupName - The name of the private dns zone group.
// options - PrivateDNSZoneGroupsClientGetOptions contains the optional parameters for the PrivateDNSZoneGroupsClient.Get
// method.
func (client *PrivateDNSZoneGroupsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsClientGetOptions) (PrivateDNSZoneGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, options)
	if err != nil {
		return PrivateDNSZoneGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateDNSZoneGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateDNSZoneGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateDNSZoneGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, options *PrivateDNSZoneGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
	if privateDNSZoneGroupName == "" {
		return nil, errors.New("parameter privateDNSZoneGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateDnsZoneGroupName}", url.PathEscape(privateDNSZoneGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateDNSZoneGroupsClient) getHandleResponse(resp *http.Response) (PrivateDNSZoneGroupsClientGetResponse, error) {
	result := PrivateDNSZoneGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateDNSZoneGroup); err != nil {
		return PrivateDNSZoneGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all private dns zone groups in a private endpoint.
// Generated from API version 2022-07-01
// privateEndpointName - The name of the private endpoint.
// resourceGroupName - The name of the resource group.
// options - PrivateDNSZoneGroupsClientListOptions contains the optional parameters for the PrivateDNSZoneGroupsClient.List
// method.
func (client *PrivateDNSZoneGroupsClient) NewListPager(privateEndpointName string, resourceGroupName string, options *PrivateDNSZoneGroupsClientListOptions) *runtime.Pager[PrivateDNSZoneGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateDNSZoneGroupsClientListResponse]{
		More: func(page PrivateDNSZoneGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateDNSZoneGroupsClientListResponse) (PrivateDNSZoneGroupsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, privateEndpointName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateDNSZoneGroupsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PrivateDNSZoneGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateDNSZoneGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateDNSZoneGroupsClient) listCreateRequest(ctx context.Context, privateEndpointName string, resourceGroupName string, options *PrivateDNSZoneGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups"
	if privateEndpointName == "" {
		return nil, errors.New("parameter privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(privateEndpointName))
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
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateDNSZoneGroupsClient) listHandleResponse(resp *http.Response) (PrivateDNSZoneGroupsClientListResponse, error) {
	result := PrivateDNSZoneGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateDNSZoneGroupListResult); err != nil {
		return PrivateDNSZoneGroupsClientListResponse{}, err
	}
	return result, nil
}
