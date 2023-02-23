//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// VendorSKUPreviewClient contains the methods for the VendorSKUPreview group.
// Don't use this type directly, use NewVendorSKUPreviewClient() instead.
type VendorSKUPreviewClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVendorSKUPreviewClient creates a new instance of VendorSKUPreviewClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVendorSKUPreviewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VendorSKUPreviewClient, error) {
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
	client := &VendorSKUPreviewClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates preview information of a vendor sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the vendor sku.
// previewSubscription - Preview subscription ID.
// parameters - Parameters supplied to the create or update vendor preview subscription operation.
// options - VendorSKUPreviewClientBeginCreateOrUpdateOptions contains the optional parameters for the VendorSKUPreviewClient.BeginCreateOrUpdate
// method.
func (client *VendorSKUPreviewClient) BeginCreateOrUpdate(ctx context.Context, vendorName string, skuName string, previewSubscription string, parameters PreviewSubscription, options *VendorSKUPreviewClientBeginCreateOrUpdateOptions) (*runtime.Poller[VendorSKUPreviewClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, vendorName, skuName, previewSubscription, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VendorSKUPreviewClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VendorSKUPreviewClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates preview information of a vendor sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *VendorSKUPreviewClient) createOrUpdate(ctx context.Context, vendorName string, skuName string, previewSubscription string, parameters PreviewSubscription, options *VendorSKUPreviewClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, vendorName, skuName, previewSubscription, parameters, options)
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
func (client *VendorSKUPreviewClient) createOrUpdateCreateRequest(ctx context.Context, vendorName string, skuName string, previewSubscription string, parameters PreviewSubscription, options *VendorSKUPreviewClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions/{previewSubscription}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if previewSubscription == "" {
		return nil, errors.New("parameter previewSubscription cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{previewSubscription}", url.PathEscape(previewSubscription))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the preview information of a vendor sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the vendor sku.
// previewSubscription - Preview subscription ID.
// options - VendorSKUPreviewClientBeginDeleteOptions contains the optional parameters for the VendorSKUPreviewClient.BeginDelete
// method.
func (client *VendorSKUPreviewClient) BeginDelete(ctx context.Context, vendorName string, skuName string, previewSubscription string, options *VendorSKUPreviewClientBeginDeleteOptions) (*runtime.Poller[VendorSKUPreviewClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, vendorName, skuName, previewSubscription, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VendorSKUPreviewClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VendorSKUPreviewClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the preview information of a vendor sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *VendorSKUPreviewClient) deleteOperation(ctx context.Context, vendorName string, skuName string, previewSubscription string, options *VendorSKUPreviewClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, vendorName, skuName, previewSubscription, options)
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
func (client *VendorSKUPreviewClient) deleteCreateRequest(ctx context.Context, vendorName string, skuName string, previewSubscription string, options *VendorSKUPreviewClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions/{previewSubscription}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if previewSubscription == "" {
		return nil, errors.New("parameter previewSubscription cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{previewSubscription}", url.PathEscape(previewSubscription))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the preview information of a vendor sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the vendor sku.
// previewSubscription - Preview subscription ID.
// options - VendorSKUPreviewClientGetOptions contains the optional parameters for the VendorSKUPreviewClient.Get method.
func (client *VendorSKUPreviewClient) Get(ctx context.Context, vendorName string, skuName string, previewSubscription string, options *VendorSKUPreviewClientGetOptions) (VendorSKUPreviewClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vendorName, skuName, previewSubscription, options)
	if err != nil {
		return VendorSKUPreviewClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VendorSKUPreviewClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VendorSKUPreviewClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VendorSKUPreviewClient) getCreateRequest(ctx context.Context, vendorName string, skuName string, previewSubscription string, options *VendorSKUPreviewClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions/{previewSubscription}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if previewSubscription == "" {
		return nil, errors.New("parameter previewSubscription cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{previewSubscription}", url.PathEscape(previewSubscription))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VendorSKUPreviewClient) getHandleResponse(resp *http.Response) (VendorSKUPreviewClientGetResponse, error) {
	result := VendorSKUPreviewClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PreviewSubscription); err != nil {
		return VendorSKUPreviewClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the preview information of a vendor sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the sku.
// options - VendorSKUPreviewClientListOptions contains the optional parameters for the VendorSKUPreviewClient.List method.
func (client *VendorSKUPreviewClient) NewListPager(vendorName string, skuName string, options *VendorSKUPreviewClientListOptions) *runtime.Pager[VendorSKUPreviewClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VendorSKUPreviewClientListResponse]{
		More: func(page VendorSKUPreviewClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VendorSKUPreviewClientListResponse) (VendorSKUPreviewClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vendorName, skuName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VendorSKUPreviewClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VendorSKUPreviewClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VendorSKUPreviewClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VendorSKUPreviewClient) listCreateRequest(ctx context.Context, vendorName string, skuName string, options *VendorSKUPreviewClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/previewSubscriptions"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VendorSKUPreviewClient) listHandleResponse(resp *http.Response) (VendorSKUPreviewClientListResponse, error) {
	result := VendorSKUPreviewClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PreviewSubscriptionsList); err != nil {
		return VendorSKUPreviewClientListResponse{}, err
	}
	return result, nil
}
