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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IdentityProviderClient contains the methods for the IdentityProvider group.
// Don't use this type directly, use NewIdentityProviderClient() instead.
type IdentityProviderClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIdentityProviderClient creates a new instance of IdentityProviderClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIdentityProviderClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IdentityProviderClient, error) {
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
	client := &IdentityProviderClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates the IdentityProvider configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// identityProviderName - Identity Provider Type identifier.
// parameters - Create parameters.
// options - IdentityProviderClientCreateOrUpdateOptions contains the optional parameters for the IdentityProviderClient.CreateOrUpdate
// method.
func (client *IdentityProviderClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, parameters IdentityProviderCreateContract, options *IdentityProviderClientCreateOrUpdateOptions) (IdentityProviderClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, identityProviderName, parameters, options)
	if err != nil {
		return IdentityProviderClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IdentityProviderClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IdentityProviderClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IdentityProviderClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, parameters IdentityProviderCreateContract, options *IdentityProviderClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if identityProviderName == "" {
		return nil, errors.New("parameter identityProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityProviderName}", url.PathEscape(string(identityProviderName)))
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
func (client *IdentityProviderClient) createOrUpdateHandleResponse(resp *http.Response) (IdentityProviderClientCreateOrUpdateResponse, error) {
	result := IdentityProviderClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IdentityProviderContract); err != nil {
		return IdentityProviderClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified identity provider configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// identityProviderName - Identity Provider Type identifier.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - IdentityProviderClientDeleteOptions contains the optional parameters for the IdentityProviderClient.Delete method.
func (client *IdentityProviderClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, ifMatch string, options *IdentityProviderClientDeleteOptions) (IdentityProviderClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, identityProviderName, ifMatch, options)
	if err != nil {
		return IdentityProviderClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IdentityProviderClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IdentityProviderClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IdentityProviderClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IdentityProviderClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, ifMatch string, options *IdentityProviderClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if identityProviderName == "" {
		return nil, errors.New("parameter identityProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityProviderName}", url.PathEscape(string(identityProviderName)))
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

// Get - Gets the configuration details of the identity Provider configured in specified service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// identityProviderName - Identity Provider Type identifier.
// options - IdentityProviderClientGetOptions contains the optional parameters for the IdentityProviderClient.Get method.
func (client *IdentityProviderClient) Get(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, options *IdentityProviderClientGetOptions) (IdentityProviderClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, identityProviderName, options)
	if err != nil {
		return IdentityProviderClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IdentityProviderClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IdentityProviderClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IdentityProviderClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, options *IdentityProviderClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if identityProviderName == "" {
		return nil, errors.New("parameter identityProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityProviderName}", url.PathEscape(string(identityProviderName)))
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

// getHandleResponse handles the Get response.
func (client *IdentityProviderClient) getHandleResponse(resp *http.Response) (IdentityProviderClientGetResponse, error) {
	result := IdentityProviderClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IdentityProviderContract); err != nil {
		return IdentityProviderClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the identityProvider specified by its identifier.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// identityProviderName - Identity Provider Type identifier.
// options - IdentityProviderClientGetEntityTagOptions contains the optional parameters for the IdentityProviderClient.GetEntityTag
// method.
func (client *IdentityProviderClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, options *IdentityProviderClientGetEntityTagOptions) (IdentityProviderClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, identityProviderName, options)
	if err != nil {
		return IdentityProviderClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IdentityProviderClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IdentityProviderClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *IdentityProviderClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, options *IdentityProviderClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if identityProviderName == "" {
		return nil, errors.New("parameter identityProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityProviderName}", url.PathEscape(string(identityProviderName)))
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
func (client *IdentityProviderClient) getEntityTagHandleResponse(resp *http.Response) (IdentityProviderClientGetEntityTagResponse, error) {
	result := IdentityProviderClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists a collection of Identity Provider configured in the specified service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - IdentityProviderClientListByServiceOptions contains the optional parameters for the IdentityProviderClient.ListByService
// method.
func (client *IdentityProviderClient) NewListByServicePager(resourceGroupName string, serviceName string, options *IdentityProviderClientListByServiceOptions) *runtime.Pager[IdentityProviderClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[IdentityProviderClientListByServiceResponse]{
		More: func(page IdentityProviderClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IdentityProviderClientListByServiceResponse) (IdentityProviderClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IdentityProviderClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IdentityProviderClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IdentityProviderClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *IdentityProviderClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *IdentityProviderClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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

// listByServiceHandleResponse handles the ListByService response.
func (client *IdentityProviderClient) listByServiceHandleResponse(resp *http.Response) (IdentityProviderClientListByServiceResponse, error) {
	result := IdentityProviderClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IdentityProviderList); err != nil {
		return IdentityProviderClientListByServiceResponse{}, err
	}
	return result, nil
}

// ListSecrets - Gets the client secret details of the Identity Provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// identityProviderName - Identity Provider Type identifier.
// options - IdentityProviderClientListSecretsOptions contains the optional parameters for the IdentityProviderClient.ListSecrets
// method.
func (client *IdentityProviderClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, options *IdentityProviderClientListSecretsOptions) (IdentityProviderClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, identityProviderName, options)
	if err != nil {
		return IdentityProviderClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IdentityProviderClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IdentityProviderClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *IdentityProviderClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, options *IdentityProviderClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}/listSecrets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if identityProviderName == "" {
		return nil, errors.New("parameter identityProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityProviderName}", url.PathEscape(string(identityProviderName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *IdentityProviderClient) listSecretsHandleResponse(resp *http.Response) (IdentityProviderClientListSecretsResponse, error) {
	result := IdentityProviderClientListSecretsResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClientSecretContract); err != nil {
		return IdentityProviderClientListSecretsResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing IdentityProvider configuration.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// identityProviderName - Identity Provider Type identifier.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update parameters.
// options - IdentityProviderClientUpdateOptions contains the optional parameters for the IdentityProviderClient.Update method.
func (client *IdentityProviderClient) Update(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, ifMatch string, parameters IdentityProviderUpdateParameters, options *IdentityProviderClientUpdateOptions) (IdentityProviderClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, identityProviderName, ifMatch, parameters, options)
	if err != nil {
		return IdentityProviderClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IdentityProviderClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IdentityProviderClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *IdentityProviderClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName IdentityProviderType, ifMatch string, parameters IdentityProviderUpdateParameters, options *IdentityProviderClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if identityProviderName == "" {
		return nil, errors.New("parameter identityProviderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{identityProviderName}", url.PathEscape(string(identityProviderName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *IdentityProviderClient) updateHandleResponse(resp *http.Response) (IdentityProviderClientUpdateResponse, error) {
	result := IdentityProviderClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.IdentityProviderContract); err != nil {
		return IdentityProviderClientUpdateResponse{}, err
	}
	return result, nil
}
