//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

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

// IotConnectorsClient contains the methods for the IotConnectors group.
// Don't use this type directly, use NewIotConnectorsClient() instead.
type IotConnectorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIotConnectorsClient creates a new instance of IotConnectorsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIotConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IotConnectorsClient, error) {
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
	client := &IotConnectorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an IoT Connector resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// iotConnectorName - The name of IoT Connector resource.
// iotConnector - The parameters for creating or updating an IoT Connectors resource.
// options - IotConnectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the IotConnectorsClient.BeginCreateOrUpdate
// method.
func (client *IotConnectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IotConnectorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, iotConnectorName, iotConnector, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[IotConnectorsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[IotConnectorsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an IoT Connector resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
func (client *IotConnectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, iotConnector, options)
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
func (client *IotConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, iotConnector)
}

// BeginDelete - Deletes an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// iotConnectorName - The name of IoT Connector resource.
// workspaceName - The name of workspace resource.
// options - IotConnectorsClientBeginDeleteOptions contains the optional parameters for the IotConnectorsClient.BeginDelete
// method.
func (client *IotConnectorsClient) BeginDelete(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsClientBeginDeleteOptions) (*runtime.Poller[IotConnectorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, iotConnectorName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[IotConnectorsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[IotConnectorsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
func (client *IotConnectorsClient) deleteOperation(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, iotConnectorName, workspaceName, options)
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
func (client *IotConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// iotConnectorName - The name of IoT Connector resource.
// options - IotConnectorsClientGetOptions contains the optional parameters for the IotConnectorsClient.Get method.
func (client *IotConnectorsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *IotConnectorsClientGetOptions) (IotConnectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, options)
	if err != nil {
		return IotConnectorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IotConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IotConnectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IotConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *IotConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotConnectorsClient) getHandleResponse(resp *http.Response) (IotConnectorsClientGetResponse, error) {
	result := IotConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotConnector); err != nil {
		return IotConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists all IoT Connectors for the given workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// options - IotConnectorsClientListByWorkspaceOptions contains the optional parameters for the IotConnectorsClient.ListByWorkspace
// method.
func (client *IotConnectorsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *IotConnectorsClientListByWorkspaceOptions) *runtime.Pager[IotConnectorsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[IotConnectorsClientListByWorkspaceResponse]{
		More: func(page IotConnectorsClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IotConnectorsClientListByWorkspaceResponse) (IotConnectorsClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IotConnectorsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IotConnectorsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IotConnectorsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *IotConnectorsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IotConnectorsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *IotConnectorsClient) listByWorkspaceHandleResponse(resp *http.Response) (IotConnectorsClientListByWorkspaceResponse, error) {
	result := IotConnectorsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotConnectorCollection); err != nil {
		return IotConnectorsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the resource group that contains the service instance.
// iotConnectorName - The name of IoT Connector resource.
// workspaceName - The name of workspace resource.
// iotConnectorPatchResource - The parameters for updating an IoT Connector.
// options - IotConnectorsClientBeginUpdateOptions contains the optional parameters for the IotConnectorsClient.BeginUpdate
// method.
func (client *IotConnectorsClient) BeginUpdate(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsClientBeginUpdateOptions) (*runtime.Poller[IotConnectorsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, iotConnectorName, workspaceName, iotConnectorPatchResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[IotConnectorsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[IotConnectorsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Patch an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
func (client *IotConnectorsClient) update(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, iotConnectorName, workspaceName, iotConnectorPatchResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *IotConnectorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, iotConnectorPatchResource)
}
