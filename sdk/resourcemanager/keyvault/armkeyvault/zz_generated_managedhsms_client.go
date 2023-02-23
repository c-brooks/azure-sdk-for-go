//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

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
	"strconv"
	"strings"
)

// ManagedHsmsClient contains the methods for the ManagedHsms group.
// Don't use this type directly, use NewManagedHsmsClient() instead.
type ManagedHsmsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedHsmsClient creates a new instance of ManagedHsmsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedHsmsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedHsmsClient, error) {
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
	client := &ManagedHsmsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a managed HSM Pool in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - Name of the resource group that contains the managed HSM pool.
// name - Name of the managed HSM Pool
// parameters - Parameters to create or update the managed HSM Pool
// options - ManagedHsmsClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedHsmsClient.BeginCreateOrUpdate
// method.
func (client *ManagedHsmsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedHsmsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, name, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedHsmsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedHsmsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a managed HSM Pool in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
func (client *ManagedHsmsClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedHsmsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified managed HSM Pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - Name of the resource group that contains the managed HSM pool.
// name - The name of the managed HSM Pool to delete
// options - ManagedHsmsClientBeginDeleteOptions contains the optional parameters for the ManagedHsmsClient.BeginDelete method.
func (client *ManagedHsmsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsClientBeginDeleteOptions) (*runtime.Poller[ManagedHsmsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedHsmsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedHsmsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified managed HSM Pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
func (client *ManagedHsmsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedHsmsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified managed HSM Pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - Name of the resource group that contains the managed HSM pool.
// name - The name of the managed HSM Pool.
// options - ManagedHsmsClientGetOptions contains the optional parameters for the ManagedHsmsClient.Get method.
func (client *ManagedHsmsClient) Get(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsClientGetOptions) (ManagedHsmsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return ManagedHsmsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedHsmsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ManagedHsmsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedHsmsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ManagedHsmsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedHsmsClient) getHandleResponse(resp *http.Response) (ManagedHsmsClientGetResponse, error) {
	result := ManagedHsmsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsm); err != nil {
		return ManagedHsmsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeleted - Gets the specified deleted managed HSM.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// name - The name of the deleted managed HSM.
// location - The location of the deleted managed HSM.
// options - ManagedHsmsClientGetDeletedOptions contains the optional parameters for the ManagedHsmsClient.GetDeleted method.
func (client *ManagedHsmsClient) GetDeleted(ctx context.Context, name string, location string, options *ManagedHsmsClientGetDeletedOptions) (ManagedHsmsClientGetDeletedResponse, error) {
	req, err := client.getDeletedCreateRequest(ctx, name, location, options)
	if err != nil {
		return ManagedHsmsClientGetDeletedResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedHsmsClientGetDeletedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedHsmsClientGetDeletedResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDeletedHandleResponse(resp)
}

// getDeletedCreateRequest creates the GetDeleted request.
func (client *ManagedHsmsClient) getDeletedCreateRequest(ctx context.Context, name string, location string, options *ManagedHsmsClientGetDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedManagedHSMs/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeletedHandleResponse handles the GetDeleted response.
func (client *ManagedHsmsClient) getDeletedHandleResponse(resp *http.Response) (ManagedHsmsClientGetDeletedResponse, error) {
	result := ManagedHsmsClientGetDeletedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedManagedHsm); err != nil {
		return ManagedHsmsClientGetDeletedResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The List operation gets information about the managed HSM Pools associated with the subscription
// and within the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - Name of the resource group that contains the managed HSM pool.
// options - ManagedHsmsClientListByResourceGroupOptions contains the optional parameters for the ManagedHsmsClient.ListByResourceGroup
// method.
func (client *ManagedHsmsClient) NewListByResourceGroupPager(resourceGroupName string, options *ManagedHsmsClientListByResourceGroupOptions) *runtime.Pager[ManagedHsmsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedHsmsClientListByResourceGroupResponse]{
		More: func(page ManagedHsmsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedHsmsClientListByResourceGroupResponse) (ManagedHsmsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedHsmsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedHsmsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedHsmsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedHsmsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedHsmsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedHsmsClient) listByResourceGroupHandleResponse(resp *http.Response) (ManagedHsmsClientListByResourceGroupResponse, error) {
	result := ManagedHsmsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmListResult); err != nil {
		return ManagedHsmsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The List operation gets information about the managed HSM Pools associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// options - ManagedHsmsClientListBySubscriptionOptions contains the optional parameters for the ManagedHsmsClient.ListBySubscription
// method.
func (client *ManagedHsmsClient) NewListBySubscriptionPager(options *ManagedHsmsClientListBySubscriptionOptions) *runtime.Pager[ManagedHsmsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedHsmsClientListBySubscriptionResponse]{
		More: func(page ManagedHsmsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedHsmsClientListBySubscriptionResponse) (ManagedHsmsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedHsmsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedHsmsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedHsmsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagedHsmsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagedHsmsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/managedHSMs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagedHsmsClient) listBySubscriptionHandleResponse(resp *http.Response) (ManagedHsmsClientListBySubscriptionResponse, error) {
	result := ManagedHsmsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedHsmListResult); err != nil {
		return ManagedHsmsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListDeletedPager - The List operation gets information about the deleted managed HSMs associated with the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// options - ManagedHsmsClientListDeletedOptions contains the optional parameters for the ManagedHsmsClient.ListDeleted method.
func (client *ManagedHsmsClient) NewListDeletedPager(options *ManagedHsmsClientListDeletedOptions) *runtime.Pager[ManagedHsmsClientListDeletedResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedHsmsClientListDeletedResponse]{
		More: func(page ManagedHsmsClientListDeletedResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedHsmsClientListDeletedResponse) (ManagedHsmsClientListDeletedResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listDeletedCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedHsmsClientListDeletedResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedHsmsClientListDeletedResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedHsmsClientListDeletedResponse{}, runtime.NewResponseError(resp)
			}
			return client.listDeletedHandleResponse(resp)
		},
	})
}

// listDeletedCreateRequest creates the ListDeleted request.
func (client *ManagedHsmsClient) listDeletedCreateRequest(ctx context.Context, options *ManagedHsmsClientListDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedManagedHSMs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDeletedHandleResponse handles the ListDeleted response.
func (client *ManagedHsmsClient) listDeletedHandleResponse(resp *http.Response) (ManagedHsmsClientListDeletedResponse, error) {
	result := ManagedHsmsClientListDeletedResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedManagedHsmListResult); err != nil {
		return ManagedHsmsClientListDeletedResponse{}, err
	}
	return result, nil
}

// BeginPurgeDeleted - Permanently deletes the specified managed HSM.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// name - The name of the soft-deleted managed HSM.
// location - The location of the soft-deleted managed HSM.
// options - ManagedHsmsClientBeginPurgeDeletedOptions contains the optional parameters for the ManagedHsmsClient.BeginPurgeDeleted
// method.
func (client *ManagedHsmsClient) BeginPurgeDeleted(ctx context.Context, name string, location string, options *ManagedHsmsClientBeginPurgeDeletedOptions) (*runtime.Poller[ManagedHsmsClientPurgeDeletedResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purgeDeleted(ctx, name, location, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedHsmsClientPurgeDeletedResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedHsmsClientPurgeDeletedResponse](options.ResumeToken, client.pl, nil)
	}
}

// PurgeDeleted - Permanently deletes the specified managed HSM.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
func (client *ManagedHsmsClient) purgeDeleted(ctx context.Context, name string, location string, options *ManagedHsmsClientBeginPurgeDeletedOptions) (*http.Response, error) {
	req, err := client.purgeDeletedCreateRequest(ctx, name, location, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// purgeDeletedCreateRequest creates the PurgeDeleted request.
func (client *ManagedHsmsClient) purgeDeletedCreateRequest(ctx context.Context, name string, location string, options *ManagedHsmsClientBeginPurgeDeletedOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedManagedHSMs/{name}/purge"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update a managed HSM Pool in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
// resourceGroupName - Name of the resource group that contains the managed HSM pool.
// name - Name of the managed HSM Pool
// parameters - Parameters to patch the managed HSM Pool
// options - ManagedHsmsClientBeginUpdateOptions contains the optional parameters for the ManagedHsmsClient.BeginUpdate method.
func (client *ManagedHsmsClient) BeginUpdate(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsClientBeginUpdateOptions) (*runtime.Poller[ManagedHsmsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, name, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedHsmsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedHsmsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a managed HSM Pool in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01-preview
func (client *ManagedHsmsClient) update(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ManagedHsmsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, parameters ManagedHsm, options *ManagedHsmsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/managedHSMs/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
