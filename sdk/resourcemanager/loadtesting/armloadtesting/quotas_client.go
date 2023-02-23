//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armloadtesting

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

// QuotasClient contains the methods for the Quotas group.
// Don't use this type directly, use NewQuotasClient() instead.
type QuotasClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewQuotasClient creates a new instance of QuotasClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQuotasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QuotasClient, error) {
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
	client := &QuotasClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckAvailability - Check Quota Availability on quota bucket per region per subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01
//   - location - The name of Azure region.
//   - quotaBucketName - Quota Bucket name.
//   - quotaBucketRequest - Quota Bucket Request data
//   - options - QuotasClientCheckAvailabilityOptions contains the optional parameters for the QuotasClient.CheckAvailability
//     method.
func (client *QuotasClient) CheckAvailability(ctx context.Context, location string, quotaBucketName string, quotaBucketRequest QuotaBucketRequest, options *QuotasClientCheckAvailabilityOptions) (QuotasClientCheckAvailabilityResponse, error) {
	req, err := client.checkAvailabilityCreateRequest(ctx, location, quotaBucketName, quotaBucketRequest, options)
	if err != nil {
		return QuotasClientCheckAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QuotasClientCheckAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotasClientCheckAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkAvailabilityHandleResponse(resp)
}

// checkAvailabilityCreateRequest creates the CheckAvailability request.
func (client *QuotasClient) checkAvailabilityCreateRequest(ctx context.Context, location string, quotaBucketName string, quotaBucketRequest QuotaBucketRequest, options *QuotasClientCheckAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LoadTestService/locations/{location}/quotas/{quotaBucketName}/checkAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if quotaBucketName == "" {
		return nil, errors.New("parameter quotaBucketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaBucketName}", url.PathEscape(quotaBucketName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, quotaBucketRequest)
}

// checkAvailabilityHandleResponse handles the CheckAvailability response.
func (client *QuotasClient) checkAvailabilityHandleResponse(resp *http.Response) (QuotasClientCheckAvailabilityResponse, error) {
	result := QuotasClientCheckAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckQuotaAvailabilityResponse); err != nil {
		return QuotasClientCheckAvailabilityResponse{}, err
	}
	return result, nil
}

// Get - Get the available quota for a quota bucket per region per subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01
//   - location - The name of Azure region.
//   - quotaBucketName - Quota Bucket name.
//   - options - QuotasClientGetOptions contains the optional parameters for the QuotasClient.Get method.
func (client *QuotasClient) Get(ctx context.Context, location string, quotaBucketName string, options *QuotasClientGetOptions) (QuotasClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, quotaBucketName, options)
	if err != nil {
		return QuotasClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QuotasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QuotasClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QuotasClient) getCreateRequest(ctx context.Context, location string, quotaBucketName string, options *QuotasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LoadTestService/locations/{location}/quotas/{quotaBucketName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if quotaBucketName == "" {
		return nil, errors.New("parameter quotaBucketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaBucketName}", url.PathEscape(quotaBucketName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QuotasClient) getHandleResponse(resp *http.Response) (QuotasClientGetResponse, error) {
	result := QuotasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaResource); err != nil {
		return QuotasClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the available quota per region per subscription.
//
// Generated from API version 2022-12-01
//   - location - The name of Azure region.
//   - options - QuotasClientListOptions contains the optional parameters for the QuotasClient.NewListPager method.
func (client *QuotasClient) NewListPager(location string, options *QuotasClientListOptions) *runtime.Pager[QuotasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[QuotasClientListResponse]{
		More: func(page QuotasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QuotasClientListResponse) (QuotasClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return QuotasClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return QuotasClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return QuotasClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *QuotasClient) listCreateRequest(ctx context.Context, location string, options *QuotasClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.LoadTestService/locations/{location}/quotas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QuotasClient) listHandleResponse(resp *http.Response) (QuotasClientListResponse, error) {
	result := QuotasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaResourceList); err != nil {
		return QuotasClientListResponse{}, err
	}
	return result, nil
}
