//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

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

// RequestStatusClient contains the methods for the QuotaRequestStatus group.
// Don't use this type directly, use NewRequestStatusClient() instead.
type RequestStatusClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRequestStatusClient creates a new instance of RequestStatusClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRequestStatusClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RequestStatusClient, error) {
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
	client := &RequestStatusClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get the quota request details and status by quota request ID for the resources of the resource provider at a specific
// location. The quota request ID id is returned in the response of the PUT
// operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-15-preview
// id - Quota request ID.
// scope - The target Azure resource URI. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/qms-test/providers/Microsoft.Batch/batchAccounts/testAccount/.
// This is the target Azure
// resource URI for the List GET operation. If a {resourceName} is added after /quotas, then it's the target Azure resource
// URI in the GET operation for the specific resource.
// options - RequestStatusClientGetOptions contains the optional parameters for the RequestStatusClient.Get method.
func (client *RequestStatusClient) Get(ctx context.Context, id string, scope string, options *RequestStatusClientGetOptions) (RequestStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, id, scope, options)
	if err != nil {
		return RequestStatusClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RequestStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RequestStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RequestStatusClient) getCreateRequest(ctx context.Context, id string, scope string, options *RequestStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Quota/quotaRequests/{id}"
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RequestStatusClient) getHandleResponse(resp *http.Response) (RequestStatusClientGetResponse, error) {
	result := RequestStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestDetails); err != nil {
		return RequestStatusClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - For the specified scope, get the current quota requests for a one year period ending at the time is made.
// Use the oData filter to select quota requests.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-03-15-preview
// scope - The target Azure resource URI. For example, /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/qms-test/providers/Microsoft.Batch/batchAccounts/testAccount/.
// This is the target Azure
// resource URI for the List GET operation. If a {resourceName} is added after /quotas, then it's the target Azure resource
// URI in the GET operation for the specific resource.
// options - RequestStatusClientListOptions contains the optional parameters for the RequestStatusClient.List method.
func (client *RequestStatusClient) NewListPager(scope string, options *RequestStatusClientListOptions) *runtime.Pager[RequestStatusClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RequestStatusClientListResponse]{
		More: func(page RequestStatusClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RequestStatusClientListResponse) (RequestStatusClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RequestStatusClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RequestStatusClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RequestStatusClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RequestStatusClient) listCreateRequest(ctx context.Context, scope string, options *RequestStatusClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Quota/quotaRequests"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RequestStatusClient) listHandleResponse(resp *http.Response) (RequestStatusClientListResponse, error) {
	result := RequestStatusClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequestDetailsList); err != nil {
		return RequestStatusClientListResponse{}, err
	}
	return result, nil
}
