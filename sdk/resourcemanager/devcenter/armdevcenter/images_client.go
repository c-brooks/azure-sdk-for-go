//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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

// ImagesClient contains the methods for the Images group.
// Don't use this type directly, use NewImagesClient() instead.
type ImagesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewImagesClient creates a new instance of ImagesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImagesClient, error) {
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
	client := &ImagesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// galleryName - The name of the gallery.
// imageName - The name of the image.
// options - ImagesClientGetOptions contains the optional parameters for the ImagesClient.Get method.
func (client *ImagesClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, imageName string, options *ImagesClientGetOptions) (ImagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, galleryName, imageName, options)
	if err != nil {
		return ImagesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, imageName string, options *ImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/galleries/{galleryName}/images/{imageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImagesClient) getHandleResponse(resp *http.Response) (ImagesClientGetResponse, error) {
	result := ImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Image); err != nil {
		return ImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDevCenterPager - Lists images for a devcenter.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// options - ImagesClientListByDevCenterOptions contains the optional parameters for the ImagesClient.ListByDevCenter method.
func (client *ImagesClient) NewListByDevCenterPager(resourceGroupName string, devCenterName string, options *ImagesClientListByDevCenterOptions) *runtime.Pager[ImagesClientListByDevCenterResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImagesClientListByDevCenterResponse]{
		More: func(page ImagesClientListByDevCenterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImagesClientListByDevCenterResponse) (ImagesClientListByDevCenterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDevCenterCreateRequest(ctx, resourceGroupName, devCenterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImagesClientListByDevCenterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ImagesClientListByDevCenterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImagesClientListByDevCenterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDevCenterHandleResponse(resp)
		},
	})
}

// listByDevCenterCreateRequest creates the ListByDevCenter request.
func (client *ImagesClient) listByDevCenterCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *ImagesClientListByDevCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDevCenterHandleResponse handles the ListByDevCenter response.
func (client *ImagesClient) listByDevCenterHandleResponse(resp *http.Response) (ImagesClientListByDevCenterResponse, error) {
	result := ImagesClientListByDevCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageListResult); err != nil {
		return ImagesClientListByDevCenterResponse{}, err
	}
	return result, nil
}

// NewListByGalleryPager - Lists images for a gallery.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// devCenterName - The name of the devcenter.
// galleryName - The name of the gallery.
// options - ImagesClientListByGalleryOptions contains the optional parameters for the ImagesClient.ListByGallery method.
func (client *ImagesClient) NewListByGalleryPager(resourceGroupName string, devCenterName string, galleryName string, options *ImagesClientListByGalleryOptions) *runtime.Pager[ImagesClientListByGalleryResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImagesClientListByGalleryResponse]{
		More: func(page ImagesClientListByGalleryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImagesClientListByGalleryResponse) (ImagesClientListByGalleryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGalleryCreateRequest(ctx, resourceGroupName, devCenterName, galleryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImagesClientListByGalleryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ImagesClientListByGalleryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImagesClientListByGalleryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGalleryHandleResponse(resp)
		},
	})
}

// listByGalleryCreateRequest creates the ListByGallery request.
func (client *ImagesClient) listByGalleryCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, options *ImagesClientListByGalleryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/galleries/{galleryName}/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByGalleryHandleResponse handles the ListByGallery response.
func (client *ImagesClient) listByGalleryHandleResponse(resp *http.Response) (ImagesClientListByGalleryResponse, error) {
	result := ImagesClientListByGalleryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageListResult); err != nil {
		return ImagesClientListByGalleryResponse{}, err
	}
	return result, nil
}
