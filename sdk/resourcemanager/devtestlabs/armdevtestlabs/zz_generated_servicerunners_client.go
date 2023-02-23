//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// ServiceRunnersClient contains the methods for the ServiceRunners group.
// Don't use this type directly, use NewServiceRunnersClient() instead.
type ServiceRunnersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceRunnersClient creates a new instance of ServiceRunnersClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceRunnersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceRunnersClient, error) {
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
	client := &ServiceRunnersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or replace an existing service runner.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the service runner.
// serviceRunner - A container for a managed identity to execute DevTest lab services.
// options - ServiceRunnersClientCreateOrUpdateOptions contains the optional parameters for the ServiceRunnersClient.CreateOrUpdate
// method.
func (client *ServiceRunnersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, serviceRunner ServiceRunner, options *ServiceRunnersClientCreateOrUpdateOptions) (ServiceRunnersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, serviceRunner, options)
	if err != nil {
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ServiceRunnersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceRunnersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, serviceRunner ServiceRunner, options *ServiceRunnersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, serviceRunner)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceRunnersClient) createOrUpdateHandleResponse(resp *http.Response) (ServiceRunnersClientCreateOrUpdateResponse, error) {
	result := ServiceRunnersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRunner); err != nil {
		return ServiceRunnersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete service runner.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the service runner.
// options - ServiceRunnersClientDeleteOptions contains the optional parameters for the ServiceRunnersClient.Delete method.
func (client *ServiceRunnersClient) Delete(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientDeleteOptions) (ServiceRunnersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return ServiceRunnersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceRunnersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ServiceRunnersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ServiceRunnersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceRunnersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get service runner.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-15
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the service runner.
// options - ServiceRunnersClientGetOptions contains the optional parameters for the ServiceRunnersClient.Get method.
func (client *ServiceRunnersClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientGetOptions) (ServiceRunnersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return ServiceRunnersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceRunnersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceRunnersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceRunnersClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *ServiceRunnersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceRunnersClient) getHandleResponse(resp *http.Response) (ServiceRunnersClientGetResponse, error) {
	result := ServiceRunnersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRunner); err != nil {
		return ServiceRunnersClientGetResponse{}, err
	}
	return result, nil
}
