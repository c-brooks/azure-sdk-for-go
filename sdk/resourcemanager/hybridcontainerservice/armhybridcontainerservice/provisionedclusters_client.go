//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcontainerservice

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

// ProvisionedClustersClient contains the methods for the ProvisionedClusters group.
// Don't use this type directly, use NewProvisionedClustersClient() instead.
type ProvisionedClustersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProvisionedClustersClient creates a new instance of ProvisionedClustersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProvisionedClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProvisionedClustersClient, error) {
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
	client := &ProvisionedClustersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates the Hybrid AKS provisioned cluster
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// provisionedClustersName - Parameter for the name of the provisioned cluster
// options - ProvisionedClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ProvisionedClustersClient.BeginCreateOrUpdate
// method.
func (client *ProvisionedClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, provisionedClustersName string, provisionedClusters ProvisionedClusters, options *ProvisionedClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ProvisionedClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, provisionedClustersName, provisionedClusters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ProvisionedClustersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProvisionedClustersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates the Hybrid AKS provisioned cluster
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *ProvisionedClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, provisionedClustersName string, provisionedClusters ProvisionedClusters, options *ProvisionedClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, provisionedClustersName, provisionedClusters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProvisionedClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, provisionedClustersName string, provisionedClusters ProvisionedClusters, options *ProvisionedClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/provisionedClusters/{provisionedClustersName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisionedClustersName == "" {
		return nil, errors.New("parameter provisionedClustersName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisionedClustersName}", url.PathEscape(provisionedClustersName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, provisionedClusters)
}

// Delete - Deletes the Hybrid AKS provisioned cluster
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// provisionedClustersName - Parameter for the name of the provisioned cluster
// options - ProvisionedClustersClientDeleteOptions contains the optional parameters for the ProvisionedClustersClient.Delete
// method.
func (client *ProvisionedClustersClient) Delete(ctx context.Context, resourceGroupName string, provisionedClustersName string, options *ProvisionedClustersClientDeleteOptions) (ProvisionedClustersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, provisionedClustersName, options)
	if err != nil {
		return ProvisionedClustersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvisionedClustersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProvisionedClustersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProvisionedClustersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProvisionedClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, provisionedClustersName string, options *ProvisionedClustersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/provisionedClusters/{provisionedClustersName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisionedClustersName == "" {
		return nil, errors.New("parameter provisionedClustersName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisionedClustersName}", url.PathEscape(provisionedClustersName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the Hybrid AKS provisioned cluster
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// provisionedClustersName - Parameter for the name of the provisioned cluster
// options - ProvisionedClustersClientGetOptions contains the optional parameters for the ProvisionedClustersClient.Get method.
func (client *ProvisionedClustersClient) Get(ctx context.Context, resourceGroupName string, provisionedClustersName string, options *ProvisionedClustersClientGetOptions) (ProvisionedClustersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, provisionedClustersName, options)
	if err != nil {
		return ProvisionedClustersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProvisionedClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProvisionedClustersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProvisionedClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, provisionedClustersName string, options *ProvisionedClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/provisionedClusters/{provisionedClustersName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisionedClustersName == "" {
		return nil, errors.New("parameter provisionedClustersName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisionedClustersName}", url.PathEscape(provisionedClustersName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProvisionedClustersClient) getHandleResponse(resp *http.Response) (ProvisionedClustersClientGetResponse, error) {
	result := ProvisionedClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClustersResponse); err != nil {
		return ProvisionedClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets the Hybrid AKS provisioned cluster in a resource group
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ProvisionedClustersClientListByResourceGroupOptions contains the optional parameters for the ProvisionedClustersClient.ListByResourceGroup
// method.
func (client *ProvisionedClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *ProvisionedClustersClientListByResourceGroupOptions) *runtime.Pager[ProvisionedClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProvisionedClustersClientListByResourceGroupResponse]{
		More: func(page ProvisionedClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProvisionedClustersClientListByResourceGroupResponse) (ProvisionedClustersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProvisionedClustersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProvisionedClustersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProvisionedClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProvisionedClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProvisionedClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/provisionedClusters"
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
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ProvisionedClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (ProvisionedClustersClientListByResourceGroupResponse, error) {
	result := ProvisionedClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClustersResponseListResult); err != nil {
		return ProvisionedClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets the Hybrid AKS provisioned cluster in a subscription
// Generated from API version 2022-05-01-preview
// options - ProvisionedClustersClientListBySubscriptionOptions contains the optional parameters for the ProvisionedClustersClient.ListBySubscription
// method.
func (client *ProvisionedClustersClient) NewListBySubscriptionPager(options *ProvisionedClustersClientListBySubscriptionOptions) *runtime.Pager[ProvisionedClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProvisionedClustersClientListBySubscriptionResponse]{
		More: func(page ProvisionedClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProvisionedClustersClientListBySubscriptionResponse) (ProvisionedClustersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProvisionedClustersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProvisionedClustersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProvisionedClustersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProvisionedClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProvisionedClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridContainerService/provisionedClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProvisionedClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (ProvisionedClustersClientListBySubscriptionResponse, error) {
	result := ProvisionedClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProvisionedClustersResponseListResult); err != nil {
		return ProvisionedClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the Hybrid AKS provisioned cluster
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// provisionedClustersName - Parameter for the name of the provisioned cluster
// options - ProvisionedClustersClientBeginUpdateOptions contains the optional parameters for the ProvisionedClustersClient.BeginUpdate
// method.
func (client *ProvisionedClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, provisionedClustersName string, provisionedClusters ProvisionedClustersPatch, options *ProvisionedClustersClientBeginUpdateOptions) (*runtime.Poller[ProvisionedClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, provisionedClustersName, provisionedClusters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ProvisionedClustersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ProvisionedClustersClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the Hybrid AKS provisioned cluster
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01-preview
func (client *ProvisionedClustersClient) update(ctx context.Context, resourceGroupName string, provisionedClustersName string, provisionedClusters ProvisionedClustersPatch, options *ProvisionedClustersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, provisionedClustersName, provisionedClusters, options)
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
func (client *ProvisionedClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, provisionedClustersName string, provisionedClusters ProvisionedClustersPatch, options *ProvisionedClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridContainerService/provisionedClusters/{provisionedClustersName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisionedClustersName == "" {
		return nil, errors.New("parameter provisionedClustersName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisionedClustersName}", url.PathEscape(provisionedClustersName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, provisionedClusters)
}
