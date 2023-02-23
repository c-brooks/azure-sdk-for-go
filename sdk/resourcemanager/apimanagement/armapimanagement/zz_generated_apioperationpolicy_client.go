//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// APIOperationPolicyClient contains the methods for the APIOperationPolicy group.
// Don't use this type directly, use NewAPIOperationPolicyClient() instead.
type APIOperationPolicyClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIOperationPolicyClient creates a new instance of APIOperationPolicyClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIOperationPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIOperationPolicyClient, error) {
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
	client := &APIOperationPolicyClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates policy configuration for the API Operation level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// parameters - The policy contents to apply.
// options - APIOperationPolicyClientCreateOrUpdateOptions contains the optional parameters for the APIOperationPolicyClient.CreateOrUpdate
// method.
func (client *APIOperationPolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, parameters PolicyContract, options *APIOperationPolicyClientCreateOrUpdateOptions) (APIOperationPolicyClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, operationID, policyID, parameters, options)
	if err != nil {
		return APIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return APIOperationPolicyClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIOperationPolicyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, parameters PolicyContract, options *APIOperationPolicyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIOperationPolicyClient) createOrUpdateHandleResponse(resp *http.Response) (APIOperationPolicyClientCreateOrUpdateResponse, error) {
	result := APIOperationPolicyClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return APIOperationPolicyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the policy configuration at the Api Operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - APIOperationPolicyClientDeleteOptions contains the optional parameters for the APIOperationPolicyClient.Delete
// method.
func (client *APIOperationPolicyClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, ifMatch string, options *APIOperationPolicyClientDeleteOptions) (APIOperationPolicyClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, operationID, policyID, ifMatch, options)
	if err != nil {
		return APIOperationPolicyClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIOperationPolicyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return APIOperationPolicyClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return APIOperationPolicyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIOperationPolicyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, ifMatch string, options *APIOperationPolicyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the policy configuration at the API Operation level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// options - APIOperationPolicyClientGetOptions contains the optional parameters for the APIOperationPolicyClient.Get method.
func (client *APIOperationPolicyClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, options *APIOperationPolicyClientGetOptions) (APIOperationPolicyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, operationID, policyID, options)
	if err != nil {
		return APIOperationPolicyClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIOperationPolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIOperationPolicyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIOperationPolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, options *APIOperationPolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("format", string(*options.Format))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIOperationPolicyClient) getHandleResponse(resp *http.Response) (APIOperationPolicyClientGetResponse, error) {
	result := APIOperationPolicyClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return APIOperationPolicyClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the API operation policy specified by its identifier.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// options - APIOperationPolicyClientGetEntityTagOptions contains the optional parameters for the APIOperationPolicyClient.GetEntityTag
// method.
func (client *APIOperationPolicyClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, options *APIOperationPolicyClientGetEntityTagOptions) (APIOperationPolicyClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, operationID, policyID, options)
	if err != nil {
		return APIOperationPolicyClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIOperationPolicyClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIOperationPolicyClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIOperationPolicyClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, policyID PolicyIDName, options *APIOperationPolicyClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIOperationPolicyClient) getEntityTagHandleResponse(resp *http.Response) (APIOperationPolicyClientGetEntityTagResponse, error) {
	result := APIOperationPolicyClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// ListByOperation - Get the list of policy configuration at the API Operation level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
// ;rev=n as a suffix where n is the revision number.
// operationID - Operation identifier within an API. Must be unique in the current API Management service instance.
// options - APIOperationPolicyClientListByOperationOptions contains the optional parameters for the APIOperationPolicyClient.ListByOperation
// method.
func (client *APIOperationPolicyClient) ListByOperation(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, options *APIOperationPolicyClientListByOperationOptions) (APIOperationPolicyClientListByOperationResponse, error) {
	req, err := client.listByOperationCreateRequest(ctx, resourceGroupName, serviceName, apiID, operationID, options)
	if err != nil {
		return APIOperationPolicyClientListByOperationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIOperationPolicyClientListByOperationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIOperationPolicyClientListByOperationResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByOperationHandleResponse(resp)
}

// listByOperationCreateRequest creates the ListByOperation request.
func (client *APIOperationPolicyClient) listByOperationCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, operationID string, options *APIOperationPolicyClientListByOperationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByOperationHandleResponse handles the ListByOperation response.
func (client *APIOperationPolicyClient) listByOperationHandleResponse(resp *http.Response) (APIOperationPolicyClientListByOperationResponse, error) {
	result := APIOperationPolicyClientListByOperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyCollection); err != nil {
		return APIOperationPolicyClientListByOperationResponse{}, err
	}
	return result, nil
}
