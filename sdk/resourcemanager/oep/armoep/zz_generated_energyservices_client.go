//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoep

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

// EnergyServicesClient contains the methods for the EnergyServices group.
// Don't use this type directly, use NewEnergyServicesClient() instead.
type EnergyServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEnergyServicesClient creates a new instance of EnergyServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEnergyServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnergyServicesClient, error) {
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
	client := &EnergyServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginAddPartition - Method that gets called if new partition is to be added in a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientBeginAddPartitionOptions contains the optional parameters for the EnergyServicesClient.BeginAddPartition
// method.
func (client *EnergyServicesClient) BeginAddPartition(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginAddPartitionOptions) (*runtime.Poller[EnergyServicesClientAddPartitionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.addPartition(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EnergyServicesClientAddPartitionResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EnergyServicesClientAddPartitionResponse](options.ResumeToken, client.pl, nil)
	}
}

// AddPartition - Method that gets called if new partition is to be added in a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
func (client *EnergyServicesClient) addPartition(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginAddPartitionOptions) (*http.Response, error) {
	req, err := client.addPartitionCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// addPartitionCreateRequest creates the AddPartition request.
func (client *EnergyServicesClient) addPartitionCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginAddPartitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}/addPartition"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginCreate - Method that gets called if subscribed for ResourceCreationBegin trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientBeginCreateOptions contains the optional parameters for the EnergyServicesClient.BeginCreate
// method.
func (client *EnergyServicesClient) BeginCreate(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginCreateOptions) (*runtime.Poller[EnergyServicesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EnergyServicesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EnergyServicesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Method that gets called if subscribed for ResourceCreationBegin trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
func (client *EnergyServicesClient) create(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *EnergyServicesClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Deletes oep resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientBeginDeleteOptions contains the optional parameters for the EnergyServicesClient.BeginDelete
// method.
func (client *EnergyServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginDeleteOptions) (*runtime.Poller[EnergyServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[EnergyServicesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[EnergyServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes oep resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
func (client *EnergyServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
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
func (client *EnergyServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns oep resource for a given name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientGetOptions contains the optional parameters for the EnergyServicesClient.Get method.
func (client *EnergyServicesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientGetOptions) (EnergyServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return EnergyServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnergyServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnergyServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EnergyServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnergyServicesClient) getHandleResponse(resp *http.Response) (EnergyServicesClientGetResponse, error) {
	result := EnergyServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnergyService); err != nil {
		return EnergyServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns list of oep resources..
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - EnergyServicesClientListByResourceGroupOptions contains the optional parameters for the EnergyServicesClient.ListByResourceGroup
// method.
func (client *EnergyServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *EnergyServicesClientListByResourceGroupOptions) *runtime.Pager[EnergyServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnergyServicesClientListByResourceGroupResponse]{
		More: func(page EnergyServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnergyServicesClientListByResourceGroupResponse) (EnergyServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnergyServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EnergyServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnergyServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *EnergyServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *EnergyServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *EnergyServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (EnergyServicesClientListByResourceGroupResponse, error) {
	result := EnergyServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnergyServiceList); err != nil {
		return EnergyServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists a collection of oep resources under the given Azure Subscription ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// options - EnergyServicesClientListBySubscriptionOptions contains the optional parameters for the EnergyServicesClient.ListBySubscription
// method.
func (client *EnergyServicesClient) NewListBySubscriptionPager(options *EnergyServicesClientListBySubscriptionOptions) *runtime.Pager[EnergyServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnergyServicesClientListBySubscriptionResponse]{
		More: func(page EnergyServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnergyServicesClientListBySubscriptionResponse) (EnergyServicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnergyServicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EnergyServicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnergyServicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *EnergyServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *EnergyServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OpenEnergyPlatform/energyServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *EnergyServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (EnergyServicesClientListBySubscriptionResponse, error) {
	result := EnergyServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnergyServiceList); err != nil {
		return EnergyServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListPartitions - Method that gets called when list of partitions is requested.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientListPartitionsOptions contains the optional parameters for the EnergyServicesClient.ListPartitions
// method.
func (client *EnergyServicesClient) ListPartitions(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientListPartitionsOptions) (EnergyServicesClientListPartitionsResponse, error) {
	req, err := client.listPartitionsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return EnergyServicesClientListPartitionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnergyServicesClientListPartitionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnergyServicesClientListPartitionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listPartitionsHandleResponse(resp)
}

// listPartitionsCreateRequest creates the ListPartitions request.
func (client *EnergyServicesClient) listPartitionsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientListPartitionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}/listPartitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listPartitionsHandleResponse handles the ListPartitions response.
func (client *EnergyServicesClient) listPartitionsHandleResponse(resp *http.Response) (EnergyServicesClientListPartitionsResponse, error) {
	result := EnergyServicesClientListPartitionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataPartitionsListResult); err != nil {
		return EnergyServicesClientListPartitionsResponse{}, err
	}
	return result, nil
}

// BeginRemovePartition - Method that gets called if new partition is to be removed from a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientBeginRemovePartitionOptions contains the optional parameters for the EnergyServicesClient.BeginRemovePartition
// method.
func (client *EnergyServicesClient) BeginRemovePartition(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginRemovePartitionOptions) (*runtime.Poller[EnergyServicesClientRemovePartitionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.removePartition(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EnergyServicesClientRemovePartitionResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EnergyServicesClientRemovePartitionResponse](options.ResumeToken, client.pl, nil)
	}
}

// RemovePartition - Method that gets called if new partition is to be removed from a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
func (client *EnergyServicesClient) removePartition(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginRemovePartitionOptions) (*http.Response, error) {
	req, err := client.removePartitionCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// removePartitionCreateRequest creates the RemovePartition request.
func (client *EnergyServicesClient) removePartitionCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientBeginRemovePartitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}/removePartition"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// Update -
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-04-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The resource name.
// options - EnergyServicesClientUpdateOptions contains the optional parameters for the EnergyServicesClient.Update method.
func (client *EnergyServicesClient) Update(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientUpdateOptions) (EnergyServicesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return EnergyServicesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EnergyServicesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnergyServicesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *EnergyServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *EnergyServicesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OpenEnergyPlatform/energyServices/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *EnergyServicesClient) updateHandleResponse(resp *http.Response) (EnergyServicesClientUpdateResponse, error) {
	result := EnergyServicesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnergyService); err != nil {
		return EnergyServicesClientUpdateResponse{}, err
	}
	return result, nil
}
