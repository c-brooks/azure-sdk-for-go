//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

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

// IntegrationRuntimeNodesClient contains the methods for the IntegrationRuntimeNodes group.
// Don't use this type directly, use NewIntegrationRuntimeNodesClient() instead.
type IntegrationRuntimeNodesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationRuntimeNodesClient creates a new instance of IntegrationRuntimeNodesClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationRuntimeNodesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeNodesClient, error) {
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
	client := &IntegrationRuntimeNodesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Deletes a self-hosted integration runtime node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// integrationRuntimeName - The integration runtime name.
// nodeName - The integration runtime node name.
// options - IntegrationRuntimeNodesClientDeleteOptions contains the optional parameters for the IntegrationRuntimeNodesClient.Delete
// method.
func (client *IntegrationRuntimeNodesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientDeleteOptions) (IntegrationRuntimeNodesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName, options)
	if err != nil {
		return IntegrationRuntimeNodesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationRuntimeNodesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IntegrationRuntimeNodesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationRuntimeNodesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a self-hosted integration runtime node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// integrationRuntimeName - The integration runtime name.
// nodeName - The integration runtime node name.
// options - IntegrationRuntimeNodesClientGetOptions contains the optional parameters for the IntegrationRuntimeNodesClient.Get
// method.
func (client *IntegrationRuntimeNodesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientGetOptions) (IntegrationRuntimeNodesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName, options)
	if err != nil {
		return IntegrationRuntimeNodesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeNodesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeNodesClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimeNodesClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeNodesClientGetResponse, error) {
	result := IntegrationRuntimeNodesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SelfHostedIntegrationRuntimeNode); err != nil {
		return IntegrationRuntimeNodesClientGetResponse{}, err
	}
	return result, nil
}

// GetIPAddress - Get the IP address of self-hosted integration runtime node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// integrationRuntimeName - The integration runtime name.
// nodeName - The integration runtime node name.
// options - IntegrationRuntimeNodesClientGetIPAddressOptions contains the optional parameters for the IntegrationRuntimeNodesClient.GetIPAddress
// method.
func (client *IntegrationRuntimeNodesClient) GetIPAddress(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientGetIPAddressOptions) (IntegrationRuntimeNodesClientGetIPAddressResponse, error) {
	req, err := client.getIPAddressCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName, options)
	if err != nil {
		return IntegrationRuntimeNodesClientGetIPAddressResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientGetIPAddressResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeNodesClientGetIPAddressResponse{}, runtime.NewResponseError(resp)
	}
	return client.getIPAddressHandleResponse(resp)
}

// getIPAddressCreateRequest creates the GetIPAddress request.
func (client *IntegrationRuntimeNodesClient) getIPAddressCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientGetIPAddressOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}/ipAddress"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getIPAddressHandleResponse handles the GetIPAddress response.
func (client *IntegrationRuntimeNodesClient) getIPAddressHandleResponse(resp *http.Response) (IntegrationRuntimeNodesClientGetIPAddressResponse, error) {
	result := IntegrationRuntimeNodesClientGetIPAddressResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeNodeIPAddress); err != nil {
		return IntegrationRuntimeNodesClientGetIPAddressResponse{}, err
	}
	return result, nil
}

// Update - Updates a self-hosted integration runtime node.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// integrationRuntimeName - The integration runtime name.
// nodeName - The integration runtime node name.
// updateIntegrationRuntimeNodeRequest - The parameters for updating an integration runtime node.
// options - IntegrationRuntimeNodesClientUpdateOptions contains the optional parameters for the IntegrationRuntimeNodesClient.Update
// method.
func (client *IntegrationRuntimeNodesClient) Update(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest UpdateIntegrationRuntimeNodeRequest, options *IntegrationRuntimeNodesClientUpdateOptions) (IntegrationRuntimeNodesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName, updateIntegrationRuntimeNodeRequest, options)
	if err != nil {
		return IntegrationRuntimeNodesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeNodesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *IntegrationRuntimeNodesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest UpdateIntegrationRuntimeNodeRequest, options *IntegrationRuntimeNodesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateIntegrationRuntimeNodeRequest)
}

// updateHandleResponse handles the Update response.
func (client *IntegrationRuntimeNodesClient) updateHandleResponse(resp *http.Response) (IntegrationRuntimeNodesClientUpdateResponse, error) {
	result := IntegrationRuntimeNodesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SelfHostedIntegrationRuntimeNode); err != nil {
		return IntegrationRuntimeNodesClientUpdateResponse{}, err
	}
	return result, nil
}
