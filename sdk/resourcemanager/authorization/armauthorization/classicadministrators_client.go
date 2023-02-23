//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// ClassicAdministratorsClient contains the methods for the ClassicAdministrators group.
// Don't use this type directly, use NewClassicAdministratorsClient() instead.
type ClassicAdministratorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClassicAdministratorsClient creates a new instance of ClassicAdministratorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClassicAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClassicAdministratorsClient, error) {
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
	client := &ClassicAdministratorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Gets service administrator, account administrator, and co-administrators for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-07-01
// options - ClassicAdministratorsClientListOptions contains the optional parameters for the ClassicAdministratorsClient.List
// method.
func (client *ClassicAdministratorsClient) NewListPager(options *ClassicAdministratorsClientListOptions) *runtime.Pager[ClassicAdministratorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClassicAdministratorsClientListResponse]{
		More: func(page ClassicAdministratorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClassicAdministratorsClientListResponse) (ClassicAdministratorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClassicAdministratorsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ClassicAdministratorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClassicAdministratorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ClassicAdministratorsClient) listCreateRequest(ctx context.Context, options *ClassicAdministratorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/classicAdministrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ClassicAdministratorsClient) listHandleResponse(resp *http.Response) (ClassicAdministratorsClientListResponse, error) {
	result := ClassicAdministratorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClassicAdministratorListResult); err != nil {
		return ClassicAdministratorsClientListResponse{}, err
	}
	return result, nil
}
