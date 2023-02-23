//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

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

// SQLManagedInstancesClient contains the methods for the SQLManagedInstances group.
// Don't use this type directly, use NewSQLManagedInstancesClient() instead.
type SQLManagedInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLManagedInstancesClient creates a new instance of SQLManagedInstancesClient with the specified values.
// subscriptionID - The ID of the Azure subscription
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLManagedInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLManagedInstancesClient, error) {
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
	client := &SQLManagedInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates or replaces a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// resourceGroupName - The name of the Azure resource group
// sqlManagedInstanceName - Name of SQL Managed Instance
// sqlManagedInstance - The SQL Managed Instance to be created or updated.
// options - SQLManagedInstancesClientBeginCreateOptions contains the optional parameters for the SQLManagedInstancesClient.BeginCreate
// method.
func (client *SQLManagedInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, sqlManagedInstance SQLManagedInstance, options *SQLManagedInstancesClientBeginCreateOptions) (*runtime.Poller[SQLManagedInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sqlManagedInstanceName, sqlManagedInstance, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SQLManagedInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SQLManagedInstancesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates or replaces a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
func (client *SQLManagedInstancesClient) create(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, sqlManagedInstance SQLManagedInstance, options *SQLManagedInstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, sqlManagedInstance, options)
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
func (client *SQLManagedInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, sqlManagedInstance SQLManagedInstance, options *SQLManagedInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sqlManagedInstance)
}

// BeginDelete - Deletes a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// resourceGroupName - The name of the Azure resource group
// sqlManagedInstanceName - Name of SQL Managed Instance
// options - SQLManagedInstancesClientBeginDeleteOptions contains the optional parameters for the SQLManagedInstancesClient.BeginDelete
// method.
func (client *SQLManagedInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientBeginDeleteOptions) (*runtime.Poller[SQLManagedInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlManagedInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SQLManagedInstancesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SQLManagedInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
func (client *SQLManagedInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, options)
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
func (client *SQLManagedInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// resourceGroupName - The name of the Azure resource group
// sqlManagedInstanceName - Name of SQL Managed Instance
// options - SQLManagedInstancesClientGetOptions contains the optional parameters for the SQLManagedInstancesClient.Get method.
func (client *SQLManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientGetOptions) (SQLManagedInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, options)
	if err != nil {
		return SQLManagedInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLManagedInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLManagedInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLManagedInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLManagedInstancesClient) getHandleResponse(resp *http.Response) (SQLManagedInstancesClientGetResponse, error) {
	result := SQLManagedInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstance); err != nil {
		return SQLManagedInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List sqlManagedInstance resources in the subscription
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// options - SQLManagedInstancesClientListOptions contains the optional parameters for the SQLManagedInstancesClient.List
// method.
func (client *SQLManagedInstancesClient) NewListPager(options *SQLManagedInstancesClientListOptions) *runtime.Pager[SQLManagedInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLManagedInstancesClientListResponse]{
		More: func(page SQLManagedInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLManagedInstancesClientListResponse) (SQLManagedInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLManagedInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLManagedInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLManagedInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLManagedInstancesClient) listCreateRequest(ctx context.Context, options *SQLManagedInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/sqlManagedInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLManagedInstancesClient) listHandleResponse(resp *http.Response) (SQLManagedInstancesClientListResponse, error) {
	result := SQLManagedInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstanceListResult); err != nil {
		return SQLManagedInstancesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all sqlManagedInstances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// resourceGroupName - The name of the Azure resource group
// options - SQLManagedInstancesClientListByResourceGroupOptions contains the optional parameters for the SQLManagedInstancesClient.ListByResourceGroup
// method.
func (client *SQLManagedInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *SQLManagedInstancesClientListByResourceGroupOptions) *runtime.Pager[SQLManagedInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLManagedInstancesClientListByResourceGroupResponse]{
		More: func(page SQLManagedInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLManagedInstancesClientListByResourceGroupResponse) (SQLManagedInstancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLManagedInstancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLManagedInstancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLManagedInstancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLManagedInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLManagedInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances"
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
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLManagedInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLManagedInstancesClientListByResourceGroupResponse, error) {
	result := SQLManagedInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstanceListResult); err != nil {
		return SQLManagedInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// resourceGroupName - The name of the Azure resource group
// sqlManagedInstanceName - Name of SQL Managed Instance
// parameters - The SQL Managed Instance.
// options - SQLManagedInstancesClientUpdateOptions contains the optional parameters for the SQLManagedInstancesClient.Update
// method.
func (client *SQLManagedInstancesClient) Update(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, parameters SQLManagedInstanceUpdate, options *SQLManagedInstancesClientUpdateOptions) (SQLManagedInstancesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, parameters, options)
	if err != nil {
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLManagedInstancesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SQLManagedInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, parameters SQLManagedInstanceUpdate, options *SQLManagedInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *SQLManagedInstancesClient) updateHandleResponse(resp *http.Response) (SQLManagedInstancesClientUpdateResponse, error) {
	result := SQLManagedInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstance); err != nil {
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
