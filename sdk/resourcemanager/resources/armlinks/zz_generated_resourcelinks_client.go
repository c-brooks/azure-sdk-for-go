//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlinks

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

// ResourceLinksClient contains the methods for the ResourceLinks group.
// Don't use this type directly, use NewResourceLinksClient() instead.
type ResourceLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceLinksClient creates a new instance of ResourceLinksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceLinksClient, error) {
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
	client := &ResourceLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a resource link between the specified resources.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-09-01
// linkID - The fully qualified ID of the resource link. Use the format,
// /subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/{provider-namespace}/{resource-type}/{resource-name}/Microsoft.Resources/links/{link-name}.
// For example,
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup/Microsoft.Web/sites/mySite/Microsoft.Resources/links/myLink
// parameters - Parameters for creating or updating a resource link.
// options - ResourceLinksClientCreateOrUpdateOptions contains the optional parameters for the ResourceLinksClient.CreateOrUpdate
// method.
func (client *ResourceLinksClient) CreateOrUpdate(ctx context.Context, linkID string, parameters ResourceLink, options *ResourceLinksClientCreateOrUpdateOptions) (ResourceLinksClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, linkID, parameters, options)
	if err != nil {
		return ResourceLinksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceLinksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ResourceLinksClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ResourceLinksClient) createOrUpdateCreateRequest(ctx context.Context, linkID string, parameters ResourceLink, options *ResourceLinksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{linkId}"
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", linkID)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ResourceLinksClient) createOrUpdateHandleResponse(resp *http.Response) (ResourceLinksClientCreateOrUpdateResponse, error) {
	result := ResourceLinksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceLink); err != nil {
		return ResourceLinksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a resource link with the specified ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-09-01
// linkID - The fully qualified ID of the resource link. Use the format,
// /subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/{provider-namespace}/{resource-type}/{resource-name}/Microsoft.Resources/links/{link-name}.
// For example,
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup/Microsoft.Web/sites/mySite/Microsoft.Resources/links/myLink
// options - ResourceLinksClientDeleteOptions contains the optional parameters for the ResourceLinksClient.Delete method.
func (client *ResourceLinksClient) Delete(ctx context.Context, linkID string, options *ResourceLinksClientDeleteOptions) (ResourceLinksClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, linkID, options)
	if err != nil {
		return ResourceLinksClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceLinksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ResourceLinksClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ResourceLinksClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ResourceLinksClient) deleteCreateRequest(ctx context.Context, linkID string, options *ResourceLinksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{linkId}"
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", linkID)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a resource link with the specified ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-09-01
// linkID - The fully qualified Id of the resource link. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup/Microsoft.Web/sites/mySite/Microsoft.Resources/links/myLink
// options - ResourceLinksClientGetOptions contains the optional parameters for the ResourceLinksClient.Get method.
func (client *ResourceLinksClient) Get(ctx context.Context, linkID string, options *ResourceLinksClientGetOptions) (ResourceLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, linkID, options)
	if err != nil {
		return ResourceLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourceLinksClient) getCreateRequest(ctx context.Context, linkID string, options *ResourceLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/{linkId}"
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", linkID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceLinksClient) getHandleResponse(resp *http.Response) (ResourceLinksClientGetResponse, error) {
	result := ResourceLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceLink); err != nil {
		return ResourceLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListAtSourceScopePager - Gets a list of resource links at and below the specified source scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-09-01
// scope - The fully qualified ID of the scope for getting the resource links. For example, to list resource links at and
// under a resource group, set the scope to
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup.
// options - ResourceLinksClientListAtSourceScopeOptions contains the optional parameters for the ResourceLinksClient.ListAtSourceScope
// method.
func (client *ResourceLinksClient) NewListAtSourceScopePager(scope string, options *ResourceLinksClientListAtSourceScopeOptions) *runtime.Pager[ResourceLinksClientListAtSourceScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceLinksClientListAtSourceScopeResponse]{
		More: func(page ResourceLinksClientListAtSourceScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceLinksClientListAtSourceScopeResponse) (ResourceLinksClientListAtSourceScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtSourceScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceLinksClientListAtSourceScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceLinksClientListAtSourceScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceLinksClientListAtSourceScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtSourceScopeHandleResponse(resp)
		},
	})
}

// listAtSourceScopeCreateRequest creates the ListAtSourceScope request.
func (client *ResourceLinksClient) listAtSourceScopeCreateRequest(ctx context.Context, scope string, options *ResourceLinksClientListAtSourceScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/links"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", "atScope()")
	}
	reqQP.Set("api-version", "2016-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtSourceScopeHandleResponse handles the ListAtSourceScope response.
func (client *ResourceLinksClient) listAtSourceScopeHandleResponse(resp *http.Response) (ResourceLinksClientListAtSourceScopeResponse, error) {
	result := ResourceLinksClientListAtSourceScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceLinkResult); err != nil {
		return ResourceLinksClientListAtSourceScopeResponse{}, err
	}
	return result, nil
}

// NewListAtSubscriptionPager - Gets all the linked resources for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2016-09-01
// options - ResourceLinksClientListAtSubscriptionOptions contains the optional parameters for the ResourceLinksClient.ListAtSubscription
// method.
func (client *ResourceLinksClient) NewListAtSubscriptionPager(options *ResourceLinksClientListAtSubscriptionOptions) *runtime.Pager[ResourceLinksClientListAtSubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceLinksClientListAtSubscriptionResponse]{
		More: func(page ResourceLinksClientListAtSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceLinksClientListAtSubscriptionResponse) (ResourceLinksClientListAtSubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAtSubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceLinksClientListAtSubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceLinksClientListAtSubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceLinksClientListAtSubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAtSubscriptionHandleResponse(resp)
		},
	})
}

// listAtSubscriptionCreateRequest creates the ListAtSubscription request.
func (client *ResourceLinksClient) listAtSubscriptionCreateRequest(ctx context.Context, options *ResourceLinksClientListAtSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/links"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2016-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAtSubscriptionHandleResponse handles the ListAtSubscription response.
func (client *ResourceLinksClient) listAtSubscriptionHandleResponse(resp *http.Response) (ResourceLinksClientListAtSubscriptionResponse, error) {
	result := ResourceLinksClientListAtSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceLinkResult); err != nil {
		return ResourceLinksClientListAtSubscriptionResponse{}, err
	}
	return result, nil
}
