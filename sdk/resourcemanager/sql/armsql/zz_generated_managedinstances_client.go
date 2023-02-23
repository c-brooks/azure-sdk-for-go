//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedInstancesClient contains the methods for the ManagedInstances group.
// Don't use this type directly, use NewManagedInstancesClient() instead.
type ManagedInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedInstancesClient creates a new instance of ManagedInstancesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstancesClient, error) {
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
	client := &ManagedInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// parameters - The requested managed instance resource state.
// options - ManagedInstancesClientBeginCreateOrUpdateOptions contains the optional parameters for the ManagedInstancesClient.BeginCreateOrUpdate
// method.
func (client *ManagedInstancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstancesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstancesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance, options *ManagedInstancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstancesClientBeginDeleteOptions contains the optional parameters for the ManagedInstancesClient.BeginDelete
// method.
func (client *ManagedInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginDeleteOptions) (*runtime.Poller[ManagedInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstancesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
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
func (client *ManagedInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginFailover - Failovers a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance to failover.
// options - ManagedInstancesClientBeginFailoverOptions contains the optional parameters for the ManagedInstancesClient.BeginFailover
// method.
func (client *ManagedInstancesClient) BeginFailover(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginFailoverOptions) (*runtime.Poller[ManagedInstancesClientFailoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.failover(ctx, resourceGroupName, managedInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstancesClientFailoverResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstancesClientFailoverResponse](options.ResumeToken, client.pl, nil)
	}
}

// Failover - Failovers a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) failover(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
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

// failoverCreateRequest creates the Failover request.
func (client *ManagedInstancesClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/failover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ReplicaType != nil {
		reqQP.Set("replicaType", string(*options.ReplicaType))
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstancesClientGetOptions contains the optional parameters for the ManagedInstancesClient.Get method.
func (client *ManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientGetOptions) (ManagedInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
	if err != nil {
		return ManagedInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstancesClient) getHandleResponse(resp *http.Response) (ManagedInstancesClientGetResponse, error) {
	result := ManagedInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstance); err != nil {
		return ManagedInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all managed instances in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// options - ManagedInstancesClientListOptions contains the optional parameters for the ManagedInstancesClient.List method.
func (client *ManagedInstancesClient) NewListPager(options *ManagedInstancesClientListOptions) *runtime.Pager[ManagedInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListResponse]{
		More: func(page ManagedInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListResponse) (ManagedInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ManagedInstancesClient) listCreateRequest(ctx context.Context, options *ManagedInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedInstancesClient) listHandleResponse(resp *http.Response) (ManagedInstancesClientListResponse, error) {
	result := ManagedInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesClientListResponse{}, err
	}
	return result, nil
}

// NewListByInstancePoolPager - Gets a list of all managed instances in an instance pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// instancePoolName - The instance pool name.
// options - ManagedInstancesClientListByInstancePoolOptions contains the optional parameters for the ManagedInstancesClient.ListByInstancePool
// method.
func (client *ManagedInstancesClient) NewListByInstancePoolPager(resourceGroupName string, instancePoolName string, options *ManagedInstancesClientListByInstancePoolOptions) *runtime.Pager[ManagedInstancesClientListByInstancePoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListByInstancePoolResponse]{
		More: func(page ManagedInstancesClientListByInstancePoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListByInstancePoolResponse) (ManagedInstancesClientListByInstancePoolResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstancePoolCreateRequest(ctx, resourceGroupName, instancePoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstancesClientListByInstancePoolResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstancesClientListByInstancePoolResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstancesClientListByInstancePoolResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstancePoolHandleResponse(resp)
		},
	})
}

// listByInstancePoolCreateRequest creates the ListByInstancePool request.
func (client *ManagedInstancesClient) listByInstancePoolCreateRequest(ctx context.Context, resourceGroupName string, instancePoolName string, options *ManagedInstancesClientListByInstancePoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/managedInstances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instancePoolName == "" {
		return nil, errors.New("parameter instancePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instancePoolName}", url.PathEscape(instancePoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstancePoolHandleResponse handles the ListByInstancePool response.
func (client *ManagedInstancesClient) listByInstancePoolHandleResponse(resp *http.Response) (ManagedInstancesClientListByInstancePoolResponse, error) {
	result := ManagedInstancesClientListByInstancePoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesClientListByInstancePoolResponse{}, err
	}
	return result, nil
}

// NewListByManagedInstancePager - Get top resource consuming queries of a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - ManagedInstancesClientListByManagedInstanceOptions contains the optional parameters for the ManagedInstancesClient.ListByManagedInstance
// method.
func (client *ManagedInstancesClient) NewListByManagedInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstancesClientListByManagedInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListByManagedInstanceResponse]{
		More: func(page ManagedInstancesClientListByManagedInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListByManagedInstanceResponse) (ManagedInstancesClientListByManagedInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstancesClientListByManagedInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstancesClientListByManagedInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstancesClientListByManagedInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagedInstanceHandleResponse(resp)
		},
	})
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstancesClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancesClientListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/topqueries"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.NumberOfQueries != nil {
		reqQP.Set("numberOfQueries", strconv.FormatInt(int64(*options.NumberOfQueries), 10))
	}
	if options != nil && options.Databases != nil {
		reqQP.Set("databases", *options.Databases)
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", string(*options.Interval))
	}
	if options != nil && options.AggregationFunction != nil {
		reqQP.Set("aggregationFunction", string(*options.AggregationFunction))
	}
	if options != nil && options.ObservationMetric != nil {
		reqQP.Set("observationMetric", string(*options.ObservationMetric))
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstancesClient) listByManagedInstanceHandleResponse(resp *http.Response) (ManagedInstancesClientListByManagedInstanceResponse, error) {
	result := ManagedInstancesClientListByManagedInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopQueriesListResult); err != nil {
		return ManagedInstancesClientListByManagedInstanceResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of managed instances in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - ManagedInstancesClientListByResourceGroupOptions contains the optional parameters for the ManagedInstancesClient.ListByResourceGroup
// method.
func (client *ManagedInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *ManagedInstancesClientListByResourceGroupOptions) *runtime.Pager[ManagedInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancesClientListByResourceGroupResponse]{
		More: func(page ManagedInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancesClientListByResourceGroupResponse) (ManagedInstancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedInstancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (ManagedInstancesClientListByResourceGroupResponse, error) {
	result := ManagedInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceListResult); err != nil {
		return ManagedInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// parameters - The requested managed instance resource state.
// options - ManagedInstancesClientBeginUpdateOptions contains the optional parameters for the ManagedInstancesClient.BeginUpdate
// method.
func (client *ManagedInstancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesClientBeginUpdateOptions) (*runtime.Poller[ManagedInstancesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedInstanceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagedInstancesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstancesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a managed instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-05-01-preview
func (client *ManagedInstancesClient) update(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedInstanceName, parameters, options)
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
func (client *ManagedInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate, options *ManagedInstancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
