//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage

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

// BlobInventoryPoliciesClient contains the methods for the BlobInventoryPolicies group.
// Don't use this type directly, use NewBlobInventoryPoliciesClient() instead.
type BlobInventoryPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBlobInventoryPoliciesClient creates a new instance of BlobInventoryPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBlobInventoryPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BlobInventoryPoliciesClient, error) {
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
	client := &BlobInventoryPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Sets the blob inventory policy to the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// blobInventoryPolicyName - The name of the storage account blob inventory policy. It should always be 'default'
// properties - The blob inventory policy set to a storage account.
// options - BlobInventoryPoliciesClientCreateOrUpdateOptions contains the optional parameters for the BlobInventoryPoliciesClient.CreateOrUpdate
// method.
func (client *BlobInventoryPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, properties BlobInventoryPolicy, options *BlobInventoryPoliciesClientCreateOrUpdateOptions) (BlobInventoryPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, blobInventoryPolicyName, properties, options)
	if err != nil {
		return BlobInventoryPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobInventoryPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobInventoryPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BlobInventoryPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, properties BlobInventoryPolicy, options *BlobInventoryPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies/{blobInventoryPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if blobInventoryPolicyName == "" {
		return nil, errors.New("parameter blobInventoryPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobInventoryPolicyName}", url.PathEscape(string(blobInventoryPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BlobInventoryPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (BlobInventoryPoliciesClientCreateOrUpdateResponse, error) {
	result := BlobInventoryPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobInventoryPolicy); err != nil {
		return BlobInventoryPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the blob inventory policy associated with the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// blobInventoryPolicyName - The name of the storage account blob inventory policy. It should always be 'default'
// options - BlobInventoryPoliciesClientDeleteOptions contains the optional parameters for the BlobInventoryPoliciesClient.Delete
// method.
func (client *BlobInventoryPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesClientDeleteOptions) (BlobInventoryPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, blobInventoryPolicyName, options)
	if err != nil {
		return BlobInventoryPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobInventoryPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return BlobInventoryPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return BlobInventoryPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BlobInventoryPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies/{blobInventoryPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if blobInventoryPolicyName == "" {
		return nil, errors.New("parameter blobInventoryPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobInventoryPolicyName}", url.PathEscape(string(blobInventoryPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the blob inventory policy associated with the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// blobInventoryPolicyName - The name of the storage account blob inventory policy. It should always be 'default'
// options - BlobInventoryPoliciesClientGetOptions contains the optional parameters for the BlobInventoryPoliciesClient.Get
// method.
func (client *BlobInventoryPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesClientGetOptions) (BlobInventoryPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, blobInventoryPolicyName, options)
	if err != nil {
		return BlobInventoryPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobInventoryPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobInventoryPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BlobInventoryPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, blobInventoryPolicyName BlobInventoryPolicyName, options *BlobInventoryPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies/{blobInventoryPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if blobInventoryPolicyName == "" {
		return nil, errors.New("parameter blobInventoryPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobInventoryPolicyName}", url.PathEscape(string(blobInventoryPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BlobInventoryPoliciesClient) getHandleResponse(resp *http.Response) (BlobInventoryPoliciesClientGetResponse, error) {
	result := BlobInventoryPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobInventoryPolicy); err != nil {
		return BlobInventoryPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the blob inventory policy associated with the specified storage account.
// Generated from API version 2022-09-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - BlobInventoryPoliciesClientListOptions contains the optional parameters for the BlobInventoryPoliciesClient.List
// method.
func (client *BlobInventoryPoliciesClient) NewListPager(resourceGroupName string, accountName string, options *BlobInventoryPoliciesClientListOptions) *runtime.Pager[BlobInventoryPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BlobInventoryPoliciesClientListResponse]{
		More: func(page BlobInventoryPoliciesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BlobInventoryPoliciesClientListResponse) (BlobInventoryPoliciesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return BlobInventoryPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BlobInventoryPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BlobInventoryPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BlobInventoryPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BlobInventoryPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/inventoryPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BlobInventoryPoliciesClient) listHandleResponse(resp *http.Response) (BlobInventoryPoliciesClientListResponse, error) {
	result := BlobInventoryPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListBlobInventoryPolicy); err != nil {
		return BlobInventoryPoliciesClientListResponse{}, err
	}
	return result, nil
}
