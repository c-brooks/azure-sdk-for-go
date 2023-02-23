//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// ConnectorGovernanceRulesClient contains the methods for the SecurityConnectorGovernanceRules group.
// Don't use this type directly, use NewConnectorGovernanceRulesClient() instead.
type ConnectorGovernanceRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectorGovernanceRulesClient creates a new instance of ConnectorGovernanceRulesClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConnectorGovernanceRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectorGovernanceRulesClient, error) {
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
	client := &ConnectorGovernanceRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or update a security GovernanceRule on the given security connector.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// securityConnectorName - The security connector name.
// ruleID - The security GovernanceRule key - unique key for the standard GovernanceRule
// governanceRule - GovernanceRule over a subscription scope
// options - ConnectorGovernanceRulesClientCreateOrUpdateOptions contains the optional parameters for the ConnectorGovernanceRulesClient.CreateOrUpdate
// method.
func (client *ConnectorGovernanceRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, governanceRule GovernanceRule, options *ConnectorGovernanceRulesClientCreateOrUpdateOptions) (ConnectorGovernanceRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, securityConnectorName, ruleID, governanceRule, options)
	if err != nil {
		return ConnectorGovernanceRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectorGovernanceRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConnectorGovernanceRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectorGovernanceRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, governanceRule GovernanceRule, options *ConnectorGovernanceRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, governanceRule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectorGovernanceRulesClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectorGovernanceRulesClientCreateOrUpdateResponse, error) {
	result := ConnectorGovernanceRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceRule); err != nil {
		return ConnectorGovernanceRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a GovernanceRule over a given scope
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// securityConnectorName - The security connector name.
// ruleID - The security GovernanceRule key - unique key for the standard GovernanceRule
// options - ConnectorGovernanceRulesClientDeleteOptions contains the optional parameters for the ConnectorGovernanceRulesClient.Delete
// method.
func (client *ConnectorGovernanceRulesClient) Delete(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, options *ConnectorGovernanceRulesClientDeleteOptions) (ConnectorGovernanceRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, securityConnectorName, ruleID, options)
	if err != nil {
		return ConnectorGovernanceRulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectorGovernanceRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConnectorGovernanceRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConnectorGovernanceRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectorGovernanceRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, options *ConnectorGovernanceRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific governanceRule for the requested scope by ruleId
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// securityConnectorName - The security connector name.
// ruleID - The security GovernanceRule key - unique key for the standard GovernanceRule
// options - ConnectorGovernanceRulesClientGetOptions contains the optional parameters for the ConnectorGovernanceRulesClient.Get
// method.
func (client *ConnectorGovernanceRulesClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, options *ConnectorGovernanceRulesClientGetOptions) (ConnectorGovernanceRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, ruleID, options)
	if err != nil {
		return ConnectorGovernanceRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectorGovernanceRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectorGovernanceRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectorGovernanceRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, options *ConnectorGovernanceRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectorGovernanceRulesClient) getHandleResponse(resp *http.Response) (ConnectorGovernanceRulesClientGetResponse, error) {
	result := ConnectorGovernanceRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceRule); err != nil {
		return ConnectorGovernanceRulesClientGetResponse{}, err
	}
	return result, nil
}
