//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomanage

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

// BestPracticesClient contains the methods for the BestPractices group.
// Don't use this type directly, use NewBestPracticesClient() instead.
type BestPracticesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewBestPracticesClient creates a new instance of BestPracticesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBestPracticesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*BestPracticesClient, error) {
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
	client := &BestPracticesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get information about a Automanage best practice
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-04
// bestPracticeName - The Automanage best practice name.
// options - BestPracticesClientGetOptions contains the optional parameters for the BestPracticesClient.Get method.
func (client *BestPracticesClient) Get(ctx context.Context, bestPracticeName string, options *BestPracticesClientGetOptions) (BestPracticesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, bestPracticeName, options)
	if err != nil {
		return BestPracticesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BestPracticesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BestPracticesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BestPracticesClient) getCreateRequest(ctx context.Context, bestPracticeName string, options *BestPracticesClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Automanage/bestPractices/{bestPracticeName}"
	if bestPracticeName == "" {
		return nil, errors.New("parameter bestPracticeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bestPracticeName}", url.PathEscape(bestPracticeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BestPracticesClient) getHandleResponse(resp *http.Response) (BestPracticesClientGetResponse, error) {
	result := BestPracticesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BestPractice); err != nil {
		return BestPracticesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByTenantPager - Retrieve a list of Automanage best practices
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-04
// options - BestPracticesClientListByTenantOptions contains the optional parameters for the BestPracticesClient.ListByTenant
// method.
func (client *BestPracticesClient) NewListByTenantPager(options *BestPracticesClientListByTenantOptions) *runtime.Pager[BestPracticesClientListByTenantResponse] {
	return runtime.NewPager(runtime.PagingHandler[BestPracticesClientListByTenantResponse]{
		More: func(page BestPracticesClientListByTenantResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *BestPracticesClientListByTenantResponse) (BestPracticesClientListByTenantResponse, error) {
			req, err := client.listByTenantCreateRequest(ctx, options)
			if err != nil {
				return BestPracticesClientListByTenantResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return BestPracticesClientListByTenantResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BestPracticesClientListByTenantResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByTenantHandleResponse(resp)
		},
	})
}

// listByTenantCreateRequest creates the ListByTenant request.
func (client *BestPracticesClient) listByTenantCreateRequest(ctx context.Context, options *BestPracticesClientListByTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Automanage/bestPractices"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTenantHandleResponse handles the ListByTenant response.
func (client *BestPracticesClient) listByTenantHandleResponse(resp *http.Response) (BestPracticesClientListByTenantResponse, error) {
	result := BestPracticesClientListByTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BestPracticeList); err != nil {
		return BestPracticesClientListByTenantResponse{}, err
	}
	return result, nil
}
