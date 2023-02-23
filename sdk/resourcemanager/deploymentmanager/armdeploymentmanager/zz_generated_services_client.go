//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

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

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServicesClient, error) {
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
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Synchronously creates a new service or updates an existing service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// serviceName - The name of the service resource.
// serviceInfo - The service object
// options - ServicesClientCreateOrUpdateOptions contains the optional parameters for the ServicesClient.CreateOrUpdate method.
func (client *ServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceInfo ServiceResource, options *ServicesClientCreateOrUpdateOptions) (ServicesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceTopologyName, serviceName, serviceInfo, options)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ServicesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceInfo ServiceResource, options *ServicesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, serviceInfo)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServicesClient) createOrUpdateHandleResponse(resp *http.Response) (ServicesClientCreateOrUpdateResponse, error) {
	result := ServicesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// serviceName - The name of the service resource.
// options - ServicesClientDeleteOptions contains the optional parameters for the ServicesClient.Delete method.
func (client *ServicesClient) Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, options *ServicesClientDeleteOptions) (ServicesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceTopologyName, serviceName, options)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ServicesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ServicesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, options *ServicesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// serviceName - The name of the service resource.
// options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceTopologyName, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the services in the service topology.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serviceTopologyName - The name of the service topology .
// options - ServicesClientListOptions contains the optional parameters for the ServicesClient.List method.
func (client *ServicesClient) List(ctx context.Context, resourceGroupName string, serviceTopologyName string, options *ServicesClientListOptions) (ServicesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, serviceTopologyName, options)
	if err != nil {
		return ServicesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceTopologyName string, options *ServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceTopologyName == "" {
		return nil, errors.New("parameter serviceTopologyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceTopologyName}", url.PathEscape(serviceTopologyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesClientListResponse, error) {
	result := ServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceArray); err != nil {
		return ServicesClientListResponse{}, err
	}
	return result, nil
}
