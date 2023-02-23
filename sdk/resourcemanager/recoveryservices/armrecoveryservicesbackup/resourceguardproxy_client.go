//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

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

// ResourceGuardProxyClient contains the methods for the ResourceGuardProxy group.
// Don't use this type directly, use NewResourceGuardProxyClient() instead.
type ResourceGuardProxyClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceGuardProxyClient creates a new instance of ResourceGuardProxyClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceGuardProxyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceGuardProxyClient, error) {
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
	client := &ResourceGuardProxyClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Delete ResourceGuardProxy under vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ResourceGuardProxyClientDeleteOptions contains the optional parameters for the ResourceGuardProxyClient.Delete
//     method.
func (client *ResourceGuardProxyClient) Delete(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *ResourceGuardProxyClientDeleteOptions) (ResourceGuardProxyClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, options)
	if err != nil {
		return ResourceGuardProxyClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGuardProxyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ResourceGuardProxyClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ResourceGuardProxyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ResourceGuardProxyClient) deleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *ResourceGuardProxyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns ResourceGuardProxy under vault and with the name referenced in request
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ResourceGuardProxyClientGetOptions contains the optional parameters for the ResourceGuardProxyClient.Get method.
func (client *ResourceGuardProxyClient) Get(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *ResourceGuardProxyClientGetOptions) (ResourceGuardProxyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, options)
	if err != nil {
		return ResourceGuardProxyClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGuardProxyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceGuardProxyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourceGuardProxyClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *ResourceGuardProxyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceGuardProxyClient) getHandleResponse(resp *http.Response) (ResourceGuardProxyClientGetResponse, error) {
	result := ResourceGuardProxyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return ResourceGuardProxyClientGetResponse{}, err
	}
	return result, nil
}

// Put - Add or Update ResourceGuardProxy under vault Secures vault critical operations
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - Request body for operation
//   - options - ResourceGuardProxyClientPutOptions contains the optional parameters for the ResourceGuardProxyClient.Put method.
func (client *ResourceGuardProxyClient) Put(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *ResourceGuardProxyClientPutOptions) (ResourceGuardProxyClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return ResourceGuardProxyClientPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGuardProxyClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceGuardProxyClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *ResourceGuardProxyClient) putCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *ResourceGuardProxyClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// putHandleResponse handles the Put response.
func (client *ResourceGuardProxyClient) putHandleResponse(resp *http.Response) (ResourceGuardProxyClientPutResponse, error) {
	result := ResourceGuardProxyClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return ResourceGuardProxyClientPutResponse{}, err
	}
	return result, nil
}

// UnlockDelete - Secures delete ResourceGuardProxy operations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - Request body for operation
//   - options - ResourceGuardProxyClientUnlockDeleteOptions contains the optional parameters for the ResourceGuardProxyClient.UnlockDelete
//     method.
func (client *ResourceGuardProxyClient) UnlockDelete(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *ResourceGuardProxyClientUnlockDeleteOptions) (ResourceGuardProxyClientUnlockDeleteResponse, error) {
	req, err := client.unlockDeleteCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.unlockDeleteHandleResponse(resp)
}

// unlockDeleteCreateRequest creates the UnlockDelete request.
func (client *ResourceGuardProxyClient) unlockDeleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *ResourceGuardProxyClientUnlockDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}/unlockDelete"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// unlockDeleteHandleResponse handles the UnlockDelete response.
func (client *ResourceGuardProxyClient) unlockDeleteHandleResponse(resp *http.Response) (ResourceGuardProxyClientUnlockDeleteResponse, error) {
	result := ResourceGuardProxyClientUnlockDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnlockDeleteResponse); err != nil {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	return result, nil
}
