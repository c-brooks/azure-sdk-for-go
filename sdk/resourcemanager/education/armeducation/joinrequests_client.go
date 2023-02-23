//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeducation

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

// JoinRequestsClient contains the methods for the JoinRequests group.
// Don't use this type directly, use NewJoinRequestsClient() instead.
type JoinRequestsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewJoinRequestsClient creates a new instance of JoinRequestsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJoinRequestsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*JoinRequestsClient, error) {
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
	client := &JoinRequestsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Approve - Approve student joining the redeemable lab
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// joinRequestName - Join name
// options - JoinRequestsClientApproveOptions contains the optional parameters for the JoinRequestsClient.Approve method.
func (client *JoinRequestsClient) Approve(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *JoinRequestsClientApproveOptions) (JoinRequestsClientApproveResponse, error) {
	req, err := client.approveCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, joinRequestName, options)
	if err != nil {
		return JoinRequestsClientApproveResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JoinRequestsClientApproveResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JoinRequestsClientApproveResponse{}, runtime.NewResponseError(resp)
	}
	return JoinRequestsClientApproveResponse{}, nil
}

// approveCreateRequest creates the Approve request.
func (client *JoinRequestsClient) approveCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *JoinRequestsClientApproveOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests/{joinRequestName}/approve"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	if joinRequestName == "" {
		return nil, errors.New("parameter joinRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{joinRequestName}", url.PathEscape(joinRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Deny - Deny student joining the redeemable lab
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// joinRequestName - Join name
// options - JoinRequestsClientDenyOptions contains the optional parameters for the JoinRequestsClient.Deny method.
func (client *JoinRequestsClient) Deny(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *JoinRequestsClientDenyOptions) (JoinRequestsClientDenyResponse, error) {
	req, err := client.denyCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, joinRequestName, options)
	if err != nil {
		return JoinRequestsClientDenyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JoinRequestsClientDenyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JoinRequestsClientDenyResponse{}, runtime.NewResponseError(resp)
	}
	return JoinRequestsClientDenyResponse{}, nil
}

// denyCreateRequest creates the Deny request.
func (client *JoinRequestsClient) denyCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *JoinRequestsClientDenyOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests/{joinRequestName}/deny"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	if joinRequestName == "" {
		return nil, errors.New("parameter joinRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{joinRequestName}", url.PathEscape(joinRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - get student join requests
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// joinRequestName - Join name
// options - JoinRequestsClientGetOptions contains the optional parameters for the JoinRequestsClient.Get method.
func (client *JoinRequestsClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *JoinRequestsClientGetOptions) (JoinRequestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, joinRequestName, options)
	if err != nil {
		return JoinRequestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JoinRequestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JoinRequestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JoinRequestsClient) getCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, joinRequestName string, options *JoinRequestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests/{joinRequestName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	if joinRequestName == "" {
		return nil, errors.New("parameter joinRequestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{joinRequestName}", url.PathEscape(joinRequestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JoinRequestsClient) getHandleResponse(resp *http.Response) (JoinRequestsClientGetResponse, error) {
	result := JoinRequestsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JoinRequestDetails); err != nil {
		return JoinRequestsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - get student join requests
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// billingAccountName - Billing account name.
// billingProfileName - Billing profile name.
// invoiceSectionName - Invoice section name.
// options - JoinRequestsClientListOptions contains the optional parameters for the JoinRequestsClient.List method.
func (client *JoinRequestsClient) NewListPager(billingAccountName string, billingProfileName string, invoiceSectionName string, options *JoinRequestsClientListOptions) *runtime.Pager[JoinRequestsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[JoinRequestsClientListResponse]{
		More: func(page JoinRequestsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JoinRequestsClientListResponse) (JoinRequestsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JoinRequestsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JoinRequestsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JoinRequestsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *JoinRequestsClient) listCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *JoinRequestsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Education/labs/default/joinRequests"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.IncludeDenied != nil {
		reqQP.Set("includeDenied", strconv.FormatBool(*options.IncludeDenied))
	}
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *JoinRequestsClient) listHandleResponse(resp *http.Response) (JoinRequestsClientListResponse, error) {
	result := JoinRequestsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JoinRequestList); err != nil {
		return JoinRequestsClientListResponse{}, err
	}
	return result, nil
}
