//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice

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

// WorkflowsClient contains the methods for the Workflows group.
// Don't use this type directly, use NewWorkflowsClient() instead.
type WorkflowsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkflowsClient creates a new instance of WorkflowsClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkflowsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowsClient, error) {
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
	client := &WorkflowsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// RegenerateAccessKey - Regenerates the callback URL access key for request triggers.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// keyType - The access key type.
// options - WorkflowsClientRegenerateAccessKeyOptions contains the optional parameters for the WorkflowsClient.RegenerateAccessKey
// method.
func (client *WorkflowsClient) RegenerateAccessKey(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType RegenerateActionParameter, options *WorkflowsClientRegenerateAccessKeyOptions) (WorkflowsClientRegenerateAccessKeyResponse, error) {
	req, err := client.regenerateAccessKeyCreateRequest(ctx, resourceGroupName, name, workflowName, keyType, options)
	if err != nil {
		return WorkflowsClientRegenerateAccessKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowsClientRegenerateAccessKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowsClientRegenerateAccessKeyResponse{}, runtime.NewResponseError(resp)
	}
	return WorkflowsClientRegenerateAccessKeyResponse{}, nil
}

// regenerateAccessKeyCreateRequest creates the RegenerateAccessKey request.
func (client *WorkflowsClient) regenerateAccessKeyCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType RegenerateActionParameter, options *WorkflowsClientRegenerateAccessKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/regenerateAccessKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, keyType)
}

// Validate - Validates the workflow definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Site name.
// workflowName - The workflow name.
// validate - The workflow.
// options - WorkflowsClientValidateOptions contains the optional parameters for the WorkflowsClient.Validate method.
func (client *WorkflowsClient) Validate(ctx context.Context, resourceGroupName string, name string, workflowName string, validate Workflow, options *WorkflowsClientValidateOptions) (WorkflowsClientValidateResponse, error) {
	req, err := client.validateCreateRequest(ctx, resourceGroupName, name, workflowName, validate, options)
	if err != nil {
		return WorkflowsClientValidateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowsClientValidateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowsClientValidateResponse{}, runtime.NewResponseError(resp)
	}
	return WorkflowsClientValidateResponse{}, nil
}

// validateCreateRequest creates the Validate request.
func (client *WorkflowsClient) validateCreateRequest(ctx context.Context, resourceGroupName string, name string, workflowName string, validate Workflow, options *WorkflowsClientValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/validate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, validate)
}
