//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

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

// OrganizationClient contains the methods for the Organization group.
// Don't use this type directly, use NewOrganizationClient() instead.
type OrganizationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOrganizationClient creates a new instance of OrganizationClient with the specified values.
// subscriptionID - Microsoft Azure subscription id
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOrganizationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrganizationClient, error) {
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
	client := &OrganizationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientBeginCreateOptions contains the optional parameters for the OrganizationClient.BeginCreate
// method.
func (client *OrganizationClient) BeginCreate(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginCreateOptions) (*runtime.Poller[OrganizationClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, organizationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[OrganizationClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[OrganizationClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
func (client *OrganizationClient) create(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, organizationName, options)
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
func (client *OrganizationClient) createCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Delete Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientBeginDeleteOptions contains the optional parameters for the OrganizationClient.BeginDelete
// method.
func (client *OrganizationClient) BeginDelete(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginDeleteOptions) (*runtime.Poller[OrganizationClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, organizationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[OrganizationClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[OrganizationClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
func (client *OrganizationClient) deleteOperation(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, options)
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
func (client *OrganizationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a specific Organization resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientGetOptions contains the optional parameters for the OrganizationClient.Get method.
func (client *OrganizationClient) Get(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientGetOptions) (OrganizationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrganizationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrganizationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OrganizationClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrganizationClient) getHandleResponse(resp *http.Response) (OrganizationClientGetResponse, error) {
	result := OrganizationClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResource); err != nil {
		return OrganizationClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Organizations under the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - Resource group name
// options - OrganizationClientListByResourceGroupOptions contains the optional parameters for the OrganizationClient.ListByResourceGroup
// method.
func (client *OrganizationClient) NewListByResourceGroupPager(resourceGroupName string, options *OrganizationClientListByResourceGroupOptions) *runtime.Pager[OrganizationClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrganizationClientListByResourceGroupResponse]{
		More: func(page OrganizationClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationClientListByResourceGroupResponse) (OrganizationClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrganizationClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrganizationClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrganizationClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OrganizationClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OrganizationClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations"
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
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OrganizationClient) listByResourceGroupHandleResponse(resp *http.Response) (OrganizationClientListByResourceGroupResponse, error) {
	result := OrganizationClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResourceListResult); err != nil {
		return OrganizationClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all organizations under the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// options - OrganizationClientListBySubscriptionOptions contains the optional parameters for the OrganizationClient.ListBySubscription
// method.
func (client *OrganizationClient) NewListBySubscriptionPager(options *OrganizationClientListBySubscriptionOptions) *runtime.Pager[OrganizationClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrganizationClientListBySubscriptionResponse]{
		More: func(page OrganizationClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationClientListBySubscriptionResponse) (OrganizationClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrganizationClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrganizationClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrganizationClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OrganizationClient) listBySubscriptionCreateRequest(ctx context.Context, options *OrganizationClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/organizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OrganizationClient) listBySubscriptionHandleResponse(resp *http.Response) (OrganizationClientListBySubscriptionResponse, error) {
	result := OrganizationClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResourceListResult); err != nil {
		return OrganizationClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update Organization resource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01
// resourceGroupName - Resource group name
// organizationName - Organization resource name
// options - OrganizationClientUpdateOptions contains the optional parameters for the OrganizationClient.Update method.
func (client *OrganizationClient) Update(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientUpdateOptions) (OrganizationClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrganizationClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrganizationClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *OrganizationClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Confluent/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *OrganizationClient) updateHandleResponse(resp *http.Response) (OrganizationClientUpdateResponse, error) {
	result := OrganizationClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrganizationResource); err != nil {
		return OrganizationClientUpdateResponse{}, err
	}
	return result, nil
}
