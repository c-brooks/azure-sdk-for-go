//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// DataMaskingRulesClient contains the methods for the DataMaskingRules group.
// Don't use this type directly, use NewDataMaskingRulesClient() instead.
type DataMaskingRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataMaskingRulesClient creates a new instance of DataMaskingRulesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataMaskingRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataMaskingRulesClient, error) {
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
	client := &DataMaskingRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Sql pool data masking rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// dataMaskingRuleName - The name of the data masking rule.
// parameters - The required parameters for creating or updating a data masking rule.
// options - DataMaskingRulesClientCreateOrUpdateOptions contains the optional parameters for the DataMaskingRulesClient.CreateOrUpdate
// method.
func (client *DataMaskingRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataMaskingRuleName string, parameters DataMaskingRule, options *DataMaskingRulesClientCreateOrUpdateOptions) (DataMaskingRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, dataMaskingRuleName, parameters, options)
	if err != nil {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataMaskingRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataMaskingRuleName string, parameters DataMaskingRule, options *DataMaskingRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	if dataMaskingRuleName == "" {
		return nil, errors.New("parameter dataMaskingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingRuleName}", url.PathEscape(dataMaskingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataMaskingRulesClient) createOrUpdateHandleResponse(resp *http.Response) (DataMaskingRulesClientCreateOrUpdateResponse, error) {
	result := DataMaskingRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingRule); err != nil {
		return DataMaskingRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets the specific Sql pool data masking rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// dataMaskingRuleName - The name of the data masking rule.
// options - DataMaskingRulesClientGetOptions contains the optional parameters for the DataMaskingRulesClient.Get method.
func (client *DataMaskingRulesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataMaskingRuleName string, options *DataMaskingRulesClientGetOptions) (DataMaskingRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, dataMaskingRuleName, options)
	if err != nil {
		return DataMaskingRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataMaskingRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataMaskingRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataMaskingRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, dataMaskingRuleName string, options *DataMaskingRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	if dataMaskingRuleName == "" {
		return nil, errors.New("parameter dataMaskingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingRuleName}", url.PathEscape(dataMaskingRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataMaskingRulesClient) getHandleResponse(resp *http.Response) (DataMaskingRulesClientGetResponse, error) {
	result := DataMaskingRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingRule); err != nil {
		return DataMaskingRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySQLPoolPager - Gets a list of Sql pool data masking rules.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// options - DataMaskingRulesClientListBySQLPoolOptions contains the optional parameters for the DataMaskingRulesClient.ListBySQLPool
// method.
func (client *DataMaskingRulesClient) NewListBySQLPoolPager(resourceGroupName string, workspaceName string, sqlPoolName string, options *DataMaskingRulesClientListBySQLPoolOptions) *runtime.Pager[DataMaskingRulesClientListBySQLPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataMaskingRulesClientListBySQLPoolResponse]{
		More: func(page DataMaskingRulesClientListBySQLPoolResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DataMaskingRulesClientListBySQLPoolResponse) (DataMaskingRulesClientListBySQLPoolResponse, error) {
			req, err := client.listBySQLPoolCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
			if err != nil {
				return DataMaskingRulesClientListBySQLPoolResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataMaskingRulesClientListBySQLPoolResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataMaskingRulesClientListBySQLPoolResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySQLPoolHandleResponse(resp)
		},
	})
}

// listBySQLPoolCreateRequest creates the ListBySQLPool request.
func (client *DataMaskingRulesClient) listBySQLPoolCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *DataMaskingRulesClientListBySQLPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	urlPath = strings.ReplaceAll(urlPath, "{dataMaskingPolicyName}", url.PathEscape("Default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySQLPoolHandleResponse handles the ListBySQLPool response.
func (client *DataMaskingRulesClient) listBySQLPoolHandleResponse(resp *http.Response) (DataMaskingRulesClientListBySQLPoolResponse, error) {
	result := DataMaskingRulesClientListBySQLPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataMaskingRuleListResult); err != nil {
		return DataMaskingRulesClientListBySQLPoolResponse{}, err
	}
	return result, nil
}
