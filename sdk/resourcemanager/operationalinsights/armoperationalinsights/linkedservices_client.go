//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights

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

// LinkedServicesClient contains the methods for the LinkedServices group.
// Don't use this type directly, use NewLinkedServicesClient() instead.
type LinkedServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLinkedServicesClient creates a new instance of LinkedServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLinkedServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkedServicesClient, error) {
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
	client := &LinkedServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// linkedServiceName - Name of the linkedServices resource
// parameters - The parameters required to create or update a linked service.
// options - LinkedServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the LinkedServicesClient.BeginCreateOrUpdate
// method.
func (client *LinkedServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters LinkedService, options *LinkedServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[LinkedServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, linkedServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LinkedServicesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a linked service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
func (client *LinkedServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters LinkedService, options *LinkedServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, linkedServiceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LinkedServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters LinkedService, options *LinkedServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices/{linkedServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a linked service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// linkedServiceName - Name of the linked service.
// options - LinkedServicesClientBeginDeleteOptions contains the optional parameters for the LinkedServicesClient.BeginDelete
// method.
func (client *LinkedServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, options *LinkedServicesClientBeginDeleteOptions) (*runtime.Poller[LinkedServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, linkedServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LinkedServicesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LinkedServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a linked service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
func (client *LinkedServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, options *LinkedServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, linkedServiceName, options)
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
func (client *LinkedServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, options *LinkedServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices/{linkedServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a linked service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// linkedServiceName - Name of the linked service.
// options - LinkedServicesClientGetOptions contains the optional parameters for the LinkedServicesClient.Get method.
func (client *LinkedServicesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, options *LinkedServicesClientGetOptions) (LinkedServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, linkedServiceName, options)
	if err != nil {
		return LinkedServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkedServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkedServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkedServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, options *LinkedServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices/{linkedServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if linkedServiceName == "" {
		return nil, errors.New("parameter linkedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkedServiceName}", url.PathEscape(linkedServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedServicesClient) getHandleResponse(resp *http.Response) (LinkedServicesClientGetResponse, error) {
	result := LinkedServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedService); err != nil {
		return LinkedServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Gets the linked services instances in a workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-08-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - LinkedServicesClientListByWorkspaceOptions contains the optional parameters for the LinkedServicesClient.ListByWorkspace
// method.
func (client *LinkedServicesClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *LinkedServicesClientListByWorkspaceOptions) *runtime.Pager[LinkedServicesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkedServicesClientListByWorkspaceResponse]{
		More: func(page LinkedServicesClientListByWorkspaceResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LinkedServicesClientListByWorkspaceResponse) (LinkedServicesClientListByWorkspaceResponse, error) {
			req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return LinkedServicesClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LinkedServicesClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkedServicesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LinkedServicesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LinkedServicesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/linkedServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *LinkedServicesClient) listByWorkspaceHandleResponse(resp *http.Response) (LinkedServicesClientListByWorkspaceResponse, error) {
	result := LinkedServicesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedServiceListResult); err != nil {
		return LinkedServicesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
