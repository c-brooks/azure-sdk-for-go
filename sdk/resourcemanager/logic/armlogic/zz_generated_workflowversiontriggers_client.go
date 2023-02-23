//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// WorkflowVersionTriggersClient contains the methods for the WorkflowVersionTriggers group.
// Don't use this type directly, use NewWorkflowVersionTriggersClient() instead.
type WorkflowVersionTriggersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowVersionTriggersClient creates a new instance of WorkflowVersionTriggersClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowVersionTriggersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowVersionTriggersClient, error) {
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
	client := &WorkflowVersionTriggersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListCallbackURL - Get the callback url for a trigger of a workflow version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// workflowName - The workflow name.
// versionID - The workflow versionId.
// triggerName - The workflow trigger name.
// options - WorkflowVersionTriggersClientListCallbackURLOptions contains the optional parameters for the WorkflowVersionTriggersClient.ListCallbackURL
// method.
func (client *WorkflowVersionTriggersClient) ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, versionID string, triggerName string, options *WorkflowVersionTriggersClientListCallbackURLOptions) (WorkflowVersionTriggersClientListCallbackURLResponse, error) {
	req, err := client.listCallbackURLCreateRequest(ctx, resourceGroupName, workflowName, versionID, triggerName, options)
	if err != nil {
		return WorkflowVersionTriggersClientListCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowVersionTriggersClientListCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowVersionTriggersClientListCallbackURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.listCallbackURLHandleResponse(resp)
}

// listCallbackURLCreateRequest creates the ListCallbackURL request.
func (client *WorkflowVersionTriggersClient) listCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, versionID string, triggerName string, options *WorkflowVersionTriggersClientListCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}/triggers/{triggerName}/listCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// listCallbackURLHandleResponse handles the ListCallbackURL response.
func (client *WorkflowVersionTriggersClient) listCallbackURLHandleResponse(resp *http.Response) (WorkflowVersionTriggersClientListCallbackURLResponse, error) {
	result := WorkflowVersionTriggersClientListCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return WorkflowVersionTriggersClientListCallbackURLResponse{}, err
	}
	return result, nil
}
