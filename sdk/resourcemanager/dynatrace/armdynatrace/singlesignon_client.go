//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdynatrace

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

// SingleSignOnClient contains the methods for the SingleSignOn group.
// Don't use this type directly, use NewSingleSignOnClient() instead.
type SingleSignOnClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSingleSignOnClient creates a new instance of SingleSignOnClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSingleSignOnClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SingleSignOnClient, error) {
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
	client := &SingleSignOnClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a DynatraceSingleSignOnResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// configurationName - Single Sign On Configuration Name
// resource - Resource create parameters.
// options - SingleSignOnClientBeginCreateOrUpdateOptions contains the optional parameters for the SingleSignOnClient.BeginCreateOrUpdate
// method.
func (client *SingleSignOnClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, resource SingleSignOnResource, options *SingleSignOnClientBeginCreateOrUpdateOptions) (*runtime.Poller[SingleSignOnClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, monitorName, configurationName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SingleSignOnClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SingleSignOnClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create a DynatraceSingleSignOnResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01
func (client *SingleSignOnClient) createOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, resource SingleSignOnResource, options *SingleSignOnClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, configurationName, resource, options)
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
func (client *SingleSignOnClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, resource SingleSignOnResource, options *SingleSignOnClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Dynatrace.Observability/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// Get - Get a DynatraceSingleSignOnResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// configurationName - Single Sign On Configuration Name
// options - SingleSignOnClientGetOptions contains the optional parameters for the SingleSignOnClient.Get method.
func (client *SingleSignOnClient) Get(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientGetOptions) (SingleSignOnClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return SingleSignOnClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SingleSignOnClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SingleSignOnClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SingleSignOnClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *SingleSignOnClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Dynatrace.Observability/monitors/{monitorName}/singleSignOnConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SingleSignOnClient) getHandleResponse(resp *http.Response) (SingleSignOnClientGetResponse, error) {
	result := SingleSignOnClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SingleSignOnResource); err != nil {
		return SingleSignOnClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all DynatraceSingleSignOnResource by monitorName
// Generated from API version 2021-09-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - SingleSignOnClientListOptions contains the optional parameters for the SingleSignOnClient.List method.
func (client *SingleSignOnClient) NewListPager(resourceGroupName string, monitorName string, options *SingleSignOnClientListOptions) *runtime.Pager[SingleSignOnClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SingleSignOnClientListResponse]{
		More: func(page SingleSignOnClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SingleSignOnClientListResponse) (SingleSignOnClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SingleSignOnClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SingleSignOnClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SingleSignOnClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SingleSignOnClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *SingleSignOnClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Dynatrace.Observability/monitors/{monitorName}/singleSignOnConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SingleSignOnClient) listHandleResponse(resp *http.Response) (SingleSignOnClientListResponse, error) {
	result := SingleSignOnClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SingleSignOnResourceListResult); err != nil {
		return SingleSignOnClientListResponse{}, err
	}
	return result, nil
}
