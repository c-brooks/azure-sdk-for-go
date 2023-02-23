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

// ScriptExecutionsClient contains the methods for the ScriptExecutions group.
// Don't use this type directly, use NewScriptExecutionsClient() instead.
type ScriptExecutionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewScriptExecutionsClient creates a new instance of ScriptExecutionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewScriptExecutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScriptExecutionsClient, error) {
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
	client := &ScriptExecutionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a script execution in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - The name of the private cloud.
// scriptExecutionName - Name of the user-invoked script execution resource
// scriptExecution - A script running in the private cloud
// options - ScriptExecutionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ScriptExecutionsClient.BeginCreateOrUpdate
// method.
func (client *ScriptExecutionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution, options *ScriptExecutionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ScriptExecutionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, scriptExecutionName, scriptExecution, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ScriptExecutionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ScriptExecutionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a script execution in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *ScriptExecutionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution, options *ScriptExecutionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, scriptExecution, options)
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
func (client *ScriptExecutionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, scriptExecution ScriptExecution, options *ScriptExecutionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}"
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
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, scriptExecution)
}

// BeginDelete - Cancel a ScriptExecution in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// scriptExecutionName - Name of the user-invoked script execution resource
// options - ScriptExecutionsClientBeginDeleteOptions contains the optional parameters for the ScriptExecutionsClient.BeginDelete
// method.
func (client *ScriptExecutionsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientBeginDeleteOptions) (*runtime.Poller[ScriptExecutionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ScriptExecutionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ScriptExecutionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Cancel a ScriptExecution in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *ScriptExecutionsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
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
func (client *ScriptExecutionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}"
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
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
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

// Get - Get an script execution by name in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// scriptExecutionName - Name of the user-invoked script execution resource
// options - ScriptExecutionsClientGetOptions contains the optional parameters for the ScriptExecutionsClient.Get method.
func (client *ScriptExecutionsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetOptions) (ScriptExecutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
	if err != nil {
		return ScriptExecutionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptExecutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptExecutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ScriptExecutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}"
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
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
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
func (client *ScriptExecutionsClient) getHandleResponse(resp *http.Response) (ScriptExecutionsClientGetResponse, error) {
	result := ScriptExecutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptExecution); err != nil {
		return ScriptExecutionsClientGetResponse{}, err
	}
	return result, nil
}

// GetExecutionLogs - Return the logs for a script execution resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// scriptExecutionName - Name of the user-invoked script execution resource
// options - ScriptExecutionsClientGetExecutionLogsOptions contains the optional parameters for the ScriptExecutionsClient.GetExecutionLogs
// method.
func (client *ScriptExecutionsClient) GetExecutionLogs(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetExecutionLogsOptions) (ScriptExecutionsClientGetExecutionLogsResponse, error) {
	req, err := client.getExecutionLogsCreateRequest(ctx, resourceGroupName, privateCloudName, scriptExecutionName, options)
	if err != nil {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getExecutionLogsHandleResponse(resp)
}

// getExecutionLogsCreateRequest creates the GetExecutionLogs request.
func (client *ScriptExecutionsClient) getExecutionLogsCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptExecutionName string, options *ScriptExecutionsClientGetExecutionLogsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions/{scriptExecutionName}/getExecutionLogs"
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
	if scriptExecutionName == "" {
		return nil, errors.New("parameter scriptExecutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptExecutionName}", url.PathEscape(scriptExecutionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ScriptOutputStreamType != nil {
		return req, runtime.MarshalAsJSON(req, options.ScriptOutputStreamType)
	}
	return req, nil
}

// getExecutionLogsHandleResponse handles the GetExecutionLogs response.
func (client *ScriptExecutionsClient) getExecutionLogsHandleResponse(resp *http.Response) (ScriptExecutionsClientGetExecutionLogsResponse, error) {
	result := ScriptExecutionsClientGetExecutionLogsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptExecution); err != nil {
		return ScriptExecutionsClientGetExecutionLogsResponse{}, err
	}
	return result, nil
}

// NewListPager - List script executions in a private cloud
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - ScriptExecutionsClientListOptions contains the optional parameters for the ScriptExecutionsClient.List method.
func (client *ScriptExecutionsClient) NewListPager(resourceGroupName string, privateCloudName string, options *ScriptExecutionsClientListOptions) *runtime.Pager[ScriptExecutionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScriptExecutionsClientListResponse]{
		More: func(page ScriptExecutionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScriptExecutionsClientListResponse) (ScriptExecutionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ScriptExecutionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ScriptExecutionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ScriptExecutionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ScriptExecutionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *ScriptExecutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptExecutions"
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
func (client *ScriptExecutionsClient) listHandleResponse(resp *http.Response) (ScriptExecutionsClientListResponse, error) {
	result := ScriptExecutionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptExecutionsList); err != nil {
		return ScriptExecutionsClientListResponse{}, err
	}
	return result, nil
}
