//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub

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

// SharedPrivateLinkResourcesClient contains the methods for the WebPubSubSharedPrivateLinkResources group.
// Don't use this type directly, use NewSharedPrivateLinkResourcesClient() instead.
type SharedPrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient with the specified values.
// subscriptionID - Gets subscription Id which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSharedPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharedPrivateLinkResourcesClient, error) {
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
	client := &SharedPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
// sharedPrivateLinkResourceName - The name of the shared private link resource
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// parameters - The shared private link resource
// options - SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginCreateOrUpdate
// method.
func (client *SharedPrivateLinkResourcesClient) BeginCreateOrUpdate(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, parameters SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SharedPrivateLinkResourcesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SharedPrivateLinkResourcesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SharedPrivateLinkResourcesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
func (client *SharedPrivateLinkResourcesClient) createOrUpdate(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, parameters SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, parameters, options)
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
func (client *SharedPrivateLinkResourcesClient) createOrUpdateCreateRequest(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, parameters SharedPrivateLinkResource, options *SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete the specified shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
// sharedPrivateLinkResourceName - The name of the shared private link resource
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - SharedPrivateLinkResourcesClientBeginDeleteOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginDelete
// method.
func (client *SharedPrivateLinkResourcesClient) BeginDelete(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*runtime.Poller[SharedPrivateLinkResourcesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SharedPrivateLinkResourcesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SharedPrivateLinkResourcesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the specified shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
func (client *SharedPrivateLinkResourcesClient) deleteOperation(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, options)
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
func (client *SharedPrivateLinkResourcesClient) deleteCreateRequest(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified shared private link resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01-preview
// sharedPrivateLinkResourceName - The name of the shared private link resource
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - SharedPrivateLinkResourcesClientGetOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.Get
// method.
func (client *SharedPrivateLinkResourcesClient) Get(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientGetOptions) (SharedPrivateLinkResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName, options)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharedPrivateLinkResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharedPrivateLinkResourcesClient) getCreateRequest(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}"
	if sharedPrivateLinkResourceName == "" {
		return nil, errors.New("parameter sharedPrivateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sharedPrivateLinkResourceName}", url.PathEscape(sharedPrivateLinkResourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedPrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientGetResponse, error) {
	result := SharedPrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResource); err != nil {
		return SharedPrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List shared private link resources
// Generated from API version 2022-08-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - SharedPrivateLinkResourcesClientListOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.List
// method.
func (client *SharedPrivateLinkResourcesClient) NewListPager(resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientListOptions) *runtime.Pager[SharedPrivateLinkResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharedPrivateLinkResourcesClientListResponse]{
		More: func(page SharedPrivateLinkResourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharedPrivateLinkResourcesClientListResponse) (SharedPrivateLinkResourcesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SharedPrivateLinkResourcesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SharedPrivateLinkResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SharedPrivateLinkResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SharedPrivateLinkResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *SharedPrivateLinkResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/sharedPrivateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SharedPrivateLinkResourcesClient) listHandleResponse(resp *http.Response) (SharedPrivateLinkResourcesClientListResponse, error) {
	result := SharedPrivateLinkResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedPrivateLinkResourceList); err != nil {
		return SharedPrivateLinkResourcesClientListResponse{}, err
	}
	return result, nil
}
