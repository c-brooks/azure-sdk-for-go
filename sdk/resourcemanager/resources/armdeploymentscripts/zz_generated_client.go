//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentscripts

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
	"strconv"
	"strings"
)

// Client contains the methods for the DeploymentScripts group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - Subscription Id which forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
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
	client := &Client{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a deployment script.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scriptName - Name of the deployment script.
// deploymentScript - Deployment script supplied to the operation.
// options - ClientBeginCreateOptions contains the optional parameters for the Client.BeginCreate method.
func (client *Client) BeginCreate(ctx context.Context, resourceGroupName string, scriptName string, deploymentScript DeploymentScriptClassification, options *ClientBeginCreateOptions) (*runtime.Poller[ClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, scriptName, deploymentScript, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a deployment script.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
func (client *Client) create(ctx context.Context, resourceGroupName string, scriptName string, deploymentScript DeploymentScriptClassification, options *ClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, scriptName, deploymentScript, options)
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

// createCreateRequest creates the Create request.
func (client *Client) createCreateRequest(ctx context.Context, resourceGroupName string, scriptName string, deploymentScript DeploymentScriptClassification, options *ClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, deploymentScript)
}

// Delete - Deletes a deployment script. When operation completes, status code 200 returned without content.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scriptName - Name of the deployment script.
// options - ClientDeleteOptions contains the optional parameters for the Client.Delete method.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, scriptName string, options *ClientDeleteOptions) (ClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scriptName, options)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Client) deleteCreateRequest(ctx context.Context, resourceGroupName string, scriptName string, options *ClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a deployment script with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scriptName - Name of the deployment script.
// options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, resourceGroupName string, scriptName string, options *ClientGetOptions) (ClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scriptName, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, resourceGroupName string, scriptName string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// GetLogs - Gets deployment script logs for a given deployment script name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scriptName - Name of the deployment script.
// options - ClientGetLogsOptions contains the optional parameters for the Client.GetLogs method.
func (client *Client) GetLogs(ctx context.Context, resourceGroupName string, scriptName string, options *ClientGetLogsOptions) (ClientGetLogsResponse, error) {
	req, err := client.getLogsCreateRequest(ctx, resourceGroupName, scriptName, options)
	if err != nil {
		return ClientGetLogsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetLogsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetLogsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLogsHandleResponse(resp)
}

// getLogsCreateRequest creates the GetLogs request.
func (client *Client) getLogsCreateRequest(ctx context.Context, resourceGroupName string, scriptName string, options *ClientGetLogsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts/{scriptName}/logs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogsHandleResponse handles the GetLogs response.
func (client *Client) getLogsHandleResponse(resp *http.Response) (ClientGetLogsResponse, error) {
	result := ClientGetLogsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptLogsList); err != nil {
		return ClientGetLogsResponse{}, err
	}
	return result, nil
}

// GetLogsDefault - Gets deployment script logs for a given deployment script name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scriptName - Name of the deployment script.
// options - ClientGetLogsDefaultOptions contains the optional parameters for the Client.GetLogsDefault method.
func (client *Client) GetLogsDefault(ctx context.Context, resourceGroupName string, scriptName string, options *ClientGetLogsDefaultOptions) (ClientGetLogsDefaultResponse, error) {
	req, err := client.getLogsDefaultCreateRequest(ctx, resourceGroupName, scriptName, options)
	if err != nil {
		return ClientGetLogsDefaultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetLogsDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetLogsDefaultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getLogsDefaultHandleResponse(resp)
}

// getLogsDefaultCreateRequest creates the GetLogsDefault request.
func (client *Client) getLogsDefaultCreateRequest(ctx context.Context, resourceGroupName string, scriptName string, options *ClientGetLogsDefaultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts/{scriptName}/logs/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	if options != nil && options.Tail != nil {
		reqQP.Set("tail", strconv.FormatInt(int64(*options.Tail), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getLogsDefaultHandleResponse handles the GetLogsDefault response.
func (client *Client) getLogsDefaultHandleResponse(resp *http.Response) (ClientGetLogsDefaultResponse, error) {
	result := ClientGetLogsDefaultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptLog); err != nil {
		return ClientGetLogsDefaultResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists deployments scripts.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ClientListByResourceGroupOptions contains the optional parameters for the Client.ListByResourceGroup method.
func (client *Client) NewListByResourceGroupPager(resourceGroupName string, options *ClientListByResourceGroupOptions) *runtime.Pager[ClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListByResourceGroupResponse]{
		More: func(page ClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListByResourceGroupResponse) (ClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *Client) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *Client) listByResourceGroupHandleResponse(resp *http.Response) (ClientListByResourceGroupResponse, error) {
	result := ClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentScriptListResult); err != nil {
		return ClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all deployment scripts for a given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// options - ClientListBySubscriptionOptions contains the optional parameters for the Client.ListBySubscription method.
func (client *Client) NewListBySubscriptionPager(options *ClientListBySubscriptionOptions) *runtime.Pager[ClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListBySubscriptionResponse]{
		More: func(page ClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListBySubscriptionResponse) (ClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *Client) listBySubscriptionCreateRequest(ctx context.Context, options *ClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deploymentScripts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *Client) listBySubscriptionHandleResponse(resp *http.Response) (ClientListBySubscriptionResponse, error) {
	result := ClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeploymentScriptListResult); err != nil {
		return ClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates deployment script tags with specified values.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// scriptName - Name of the deployment script.
// options - ClientUpdateOptions contains the optional parameters for the Client.Update method.
func (client *Client) Update(ctx context.Context, resourceGroupName string, scriptName string, options *ClientUpdateOptions) (ClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, scriptName, options)
	if err != nil {
		return ClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *Client) updateCreateRequest(ctx context.Context, resourceGroupName string, scriptName string, options *ClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deploymentScripts/{scriptName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scriptName == "" {
		return nil, errors.New("parameter scriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptName}", url.PathEscape(scriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.DeploymentScript != nil {
		return req, runtime.MarshalAsJSON(req, *options.DeploymentScript)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *Client) updateHandleResponse(resp *http.Response) (ClientUpdateResponse, error) {
	result := ClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ClientUpdateResponse{}, err
	}
	return result, nil
}
