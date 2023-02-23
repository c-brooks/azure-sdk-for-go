//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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

// GalleriesClient contains the methods for the Galleries group.
// Don't use this type directly, use NewGalleriesClient() instead.
type GalleriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGalleriesClient creates a new instance of GalleriesClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGalleriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleriesClient, error) {
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
	client := &GalleriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods
// allowed in the middle. The maximum length is 80 characters.
// gallery - Parameters supplied to the create or update Shared Image Gallery operation.
// options - GalleriesClientBeginCreateOrUpdateOptions contains the optional parameters for the GalleriesClient.BeginCreateOrUpdate
// method.
func (client *GalleriesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery, options *GalleriesClientBeginCreateOrUpdateOptions) (*runtime.Poller[GalleriesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, gallery, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GalleriesClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GalleriesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
func (client *GalleriesClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery, options *GalleriesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, gallery, options)
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
func (client *GalleriesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, gallery Gallery, options *GalleriesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, gallery)
}

// BeginDelete - Delete a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery to be deleted.
// options - GalleriesClientBeginDeleteOptions contains the optional parameters for the GalleriesClient.BeginDelete method.
func (client *GalleriesClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesClientBeginDeleteOptions) (*runtime.Poller[GalleriesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GalleriesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GalleriesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
func (client *GalleriesClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, options)
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
func (client *GalleriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves information about a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery.
// options - GalleriesClientGetOptions contains the optional parameters for the GalleriesClient.Get method.
func (client *GalleriesClient) Get(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesClientGetOptions) (GalleriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, options)
	if err != nil {
		return GalleriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GalleriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GalleriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, options *GalleriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-03")
	if options != nil && options.Select != nil {
		reqQP.Set("$select", string(*options.Select))
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleriesClient) getHandleResponse(resp *http.Response) (GalleriesClientGetResponse, error) {
	result := GalleriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Gallery); err != nil {
		return GalleriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List galleries under a subscription.
// Generated from API version 2022-03-03
// options - GalleriesClientListOptions contains the optional parameters for the GalleriesClient.List method.
func (client *GalleriesClient) NewListPager(options *GalleriesClientListOptions) *runtime.Pager[GalleriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GalleriesClientListResponse]{
		More: func(page GalleriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleriesClientListResponse) (GalleriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GalleriesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GalleriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GalleriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GalleriesClient) listCreateRequest(ctx context.Context, options *GalleriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GalleriesClient) listHandleResponse(resp *http.Response) (GalleriesClientListResponse, error) {
	result := GalleriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryList); err != nil {
		return GalleriesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List galleries under a resource group.
// Generated from API version 2022-03-03
// resourceGroupName - The name of the resource group.
// options - GalleriesClientListByResourceGroupOptions contains the optional parameters for the GalleriesClient.ListByResourceGroup
// method.
func (client *GalleriesClient) NewListByResourceGroupPager(resourceGroupName string, options *GalleriesClientListByResourceGroupOptions) *runtime.Pager[GalleriesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[GalleriesClientListByResourceGroupResponse]{
		More: func(page GalleriesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleriesClientListByResourceGroupResponse) (GalleriesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GalleriesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GalleriesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GalleriesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GalleriesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GalleriesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries"
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
	reqQP.Set("api-version", "2022-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GalleriesClient) listByResourceGroupHandleResponse(resp *http.Response) (GalleriesClientListByResourceGroupResponse, error) {
	result := GalleriesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryList); err != nil {
		return GalleriesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
// resourceGroupName - The name of the resource group.
// galleryName - The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods
// allowed in the middle. The maximum length is 80 characters.
// gallery - Parameters supplied to the update Shared Image Gallery operation.
// options - GalleriesClientBeginUpdateOptions contains the optional parameters for the GalleriesClient.BeginUpdate method.
func (client *GalleriesClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate, options *GalleriesClientBeginUpdateOptions) (*runtime.Poller[GalleriesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, gallery, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GalleriesClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GalleriesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update a Shared Image Gallery.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-03
func (client *GalleriesClient) update(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate, options *GalleriesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, gallery, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, gallery GalleryUpdate, options *GalleriesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, gallery)
}
