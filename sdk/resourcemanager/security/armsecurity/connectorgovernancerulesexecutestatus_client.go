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

// ConnectorGovernanceRulesExecuteStatusClient contains the methods for the SecurityConnectorGovernanceRulesExecuteStatus group.
// Don't use this type directly, use NewConnectorGovernanceRulesExecuteStatusClient() instead.
type ConnectorGovernanceRulesExecuteStatusClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectorGovernanceRulesExecuteStatusClient creates a new instance of ConnectorGovernanceRulesExecuteStatusClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConnectorGovernanceRulesExecuteStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectorGovernanceRulesExecuteStatusClient, error) {
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
	client := &ConnectorGovernanceRulesExecuteStatusClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginGet - Get a specific governanceRule execution status for the requested scope by ruleId and operationId
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// securityConnectorName - The security connector name.
// ruleID - The security GovernanceRule key - unique key for the standard GovernanceRule
// operationID - The security GovernanceRule execution key - unique key for the execution of GovernanceRule
// options - ConnectorGovernanceRulesExecuteStatusClientBeginGetOptions contains the optional parameters for the ConnectorGovernanceRulesExecuteStatusClient.BeginGet
// method.
func (client *ConnectorGovernanceRulesExecuteStatusClient) BeginGet(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, operationID string, options *ConnectorGovernanceRulesExecuteStatusClientBeginGetOptions) (*runtime.Poller[ConnectorGovernanceRulesExecuteStatusClientGetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.get(ctx, resourceGroupName, securityConnectorName, ruleID, operationID, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ConnectorGovernanceRulesExecuteStatusClientGetResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConnectorGovernanceRulesExecuteStatusClientGetResponse](options.ResumeToken, client.pl, nil)
	}
}

// Get - Get a specific governanceRule execution status for the requested scope by ruleId and operationId
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *ConnectorGovernanceRulesExecuteStatusClient) get(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, operationID string, options *ConnectorGovernanceRulesExecuteStatusClientBeginGetOptions) (*http.Response, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, ruleID, operationID, options)
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

// getCreateRequest creates the Get request.
func (client *ConnectorGovernanceRulesExecuteStatusClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ruleID string, operationID string, options *ConnectorGovernanceRulesExecuteStatusClientBeginGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/providers/Microsoft.Security/governanceRules/{ruleId}/operationResults/{operationId}"
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
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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
