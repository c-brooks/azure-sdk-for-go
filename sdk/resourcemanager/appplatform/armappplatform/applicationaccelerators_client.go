//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

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

// ApplicationAcceleratorsClient contains the methods for the ApplicationAccelerators group.
// Don't use this type directly, use NewApplicationAcceleratorsClient() instead.
type ApplicationAcceleratorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationAcceleratorsClient creates a new instance of ApplicationAcceleratorsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationAcceleratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationAcceleratorsClient, error) {
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
	client := &ApplicationAcceleratorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the application accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// applicationAcceleratorName - The name of the application accelerator.
// applicationAcceleratorResource - The application accelerator for the create or update operation
// options - ApplicationAcceleratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationAcceleratorsClient.BeginCreateOrUpdate
// method.
func (client *ApplicationAcceleratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, applicationAcceleratorResource ApplicationAcceleratorResource, options *ApplicationAcceleratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationAcceleratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, applicationAcceleratorName, applicationAcceleratorResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ApplicationAcceleratorsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationAcceleratorsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the application accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *ApplicationAcceleratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, applicationAcceleratorResource ApplicationAcceleratorResource, options *ApplicationAcceleratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, applicationAcceleratorResource, options)
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
func (client *ApplicationAcceleratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, applicationAcceleratorResource ApplicationAcceleratorResource, options *ApplicationAcceleratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, applicationAcceleratorResource)
}

// BeginDelete - Delete the application accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// applicationAcceleratorName - The name of the application accelerator.
// options - ApplicationAcceleratorsClientBeginDeleteOptions contains the optional parameters for the ApplicationAcceleratorsClient.BeginDelete
// method.
func (client *ApplicationAcceleratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *ApplicationAcceleratorsClientBeginDeleteOptions) (*runtime.Poller[ApplicationAcceleratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, applicationAcceleratorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ApplicationAcceleratorsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationAcceleratorsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the application accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *ApplicationAcceleratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *ApplicationAcceleratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, options)
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
func (client *ApplicationAcceleratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *ApplicationAcceleratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the application accelerator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// applicationAcceleratorName - The name of the application accelerator.
// options - ApplicationAcceleratorsClientGetOptions contains the optional parameters for the ApplicationAcceleratorsClient.Get
// method.
func (client *ApplicationAcceleratorsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *ApplicationAcceleratorsClientGetOptions) (ApplicationAcceleratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, applicationAcceleratorName, options)
	if err != nil {
		return ApplicationAcceleratorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationAcceleratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationAcceleratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationAcceleratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, options *ApplicationAcceleratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if applicationAcceleratorName == "" {
		return nil, errors.New("parameter applicationAcceleratorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationAcceleratorName}", url.PathEscape(applicationAcceleratorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationAcceleratorsClient) getHandleResponse(resp *http.Response) (ApplicationAcceleratorsClientGetResponse, error) {
	result := ApplicationAcceleratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationAcceleratorResource); err != nil {
		return ApplicationAcceleratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handle requests to list all application accelerator.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ApplicationAcceleratorsClientListOptions contains the optional parameters for the ApplicationAcceleratorsClient.List
// method.
func (client *ApplicationAcceleratorsClient) NewListPager(resourceGroupName string, serviceName string, options *ApplicationAcceleratorsClientListOptions) *runtime.Pager[ApplicationAcceleratorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationAcceleratorsClientListResponse]{
		More: func(page ApplicationAcceleratorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationAcceleratorsClientListResponse) (ApplicationAcceleratorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationAcceleratorsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationAcceleratorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationAcceleratorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationAcceleratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ApplicationAcceleratorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationAcceleratorsClient) listHandleResponse(resp *http.Response) (ApplicationAcceleratorsClientListResponse, error) {
	result := ApplicationAcceleratorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationAcceleratorResourceCollection); err != nil {
		return ApplicationAcceleratorsClientListResponse{}, err
	}
	return result, nil
}
