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

// BuildpackBindingClient contains the methods for the BuildpackBinding group.
// Don't use this type directly, use NewBuildpackBindingClient() instead.
type BuildpackBindingClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBuildpackBindingClient creates a new instance of BuildpackBindingClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBuildpackBindingClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BuildpackBindingClient, error) {
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
	client := &BuildpackBindingClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a buildpack binding.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// buildpackBindingName - The name of the Buildpack Binding Name
// buildpackBinding - The target buildpack binding for the create or update operation
// options - BuildpackBindingClientBeginCreateOrUpdateOptions contains the optional parameters for the BuildpackBindingClient.BeginCreateOrUpdate
// method.
func (client *BuildpackBindingClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, buildpackBinding BuildpackBindingResource, options *BuildpackBindingClientBeginCreateOrUpdateOptions) (*runtime.Poller[BuildpackBindingClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, buildServiceName, builderName, buildpackBindingName, buildpackBinding, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BuildpackBindingClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BuildpackBindingClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a buildpack binding.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *BuildpackBindingClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, buildpackBinding BuildpackBindingResource, options *BuildpackBindingClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, buildpackBindingName, buildpackBinding, options)
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
func (client *BuildpackBindingClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, buildpackBinding BuildpackBindingResource, options *BuildpackBindingClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}/buildpackBindings/{buildpackBindingName}"
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
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	if buildpackBindingName == "" {
		return nil, errors.New("parameter buildpackBindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildpackBindingName}", url.PathEscape(buildpackBindingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, buildpackBinding)
}

// BeginDelete - Operation to delete a Buildpack Binding
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// buildpackBindingName - The name of the Buildpack Binding Name
// options - BuildpackBindingClientBeginDeleteOptions contains the optional parameters for the BuildpackBindingClient.BeginDelete
// method.
func (client *BuildpackBindingClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *BuildpackBindingClientBeginDeleteOptions) (*runtime.Poller[BuildpackBindingClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, buildServiceName, builderName, buildpackBindingName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BuildpackBindingClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[BuildpackBindingClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Operation to delete a Buildpack Binding
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
func (client *BuildpackBindingClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *BuildpackBindingClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, buildpackBindingName, options)
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
func (client *BuildpackBindingClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *BuildpackBindingClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}/buildpackBindings/{buildpackBindingName}"
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
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	if buildpackBindingName == "" {
		return nil, errors.New("parameter buildpackBindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildpackBindingName}", url.PathEscape(buildpackBindingName))
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

// Get - Get a buildpack binding by name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// buildpackBindingName - The name of the Buildpack Binding Name
// options - BuildpackBindingClientGetOptions contains the optional parameters for the BuildpackBindingClient.Get method.
func (client *BuildpackBindingClient) Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *BuildpackBindingClientGetOptions) (BuildpackBindingClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, buildpackBindingName, options)
	if err != nil {
		return BuildpackBindingClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BuildpackBindingClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BuildpackBindingClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BuildpackBindingClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, options *BuildpackBindingClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}/buildpackBindings/{buildpackBindingName}"
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
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
	if buildpackBindingName == "" {
		return nil, errors.New("parameter buildpackBindingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildpackBindingName}", url.PathEscape(buildpackBindingName))
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
func (client *BuildpackBindingClient) getHandleResponse(resp *http.Response) (BuildpackBindingClientGetResponse, error) {
	result := BuildpackBindingClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuildpackBindingResource); err != nil {
		return BuildpackBindingClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all buildpack bindings in a builder.
// Generated from API version 2022-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// buildServiceName - The name of the build service resource.
// builderName - The name of the builder resource.
// options - BuildpackBindingClientListOptions contains the optional parameters for the BuildpackBindingClient.List method.
func (client *BuildpackBindingClient) NewListPager(resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildpackBindingClientListOptions) *runtime.Pager[BuildpackBindingClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BuildpackBindingClientListResponse]{
		More: func(page BuildpackBindingClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BuildpackBindingClientListResponse) (BuildpackBindingClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, builderName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BuildpackBindingClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BuildpackBindingClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BuildpackBindingClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BuildpackBindingClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *BuildpackBindingClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}/buildpackBindings"
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
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if builderName == "" {
		return nil, errors.New("parameter builderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{builderName}", url.PathEscape(builderName))
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
func (client *BuildpackBindingClient) listHandleResponse(resp *http.Response) (BuildpackBindingClientListResponse, error) {
	result := BuildpackBindingClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuildpackBindingResourceCollection); err != nil {
		return BuildpackBindingClientListResponse{}, err
	}
	return result, nil
}
