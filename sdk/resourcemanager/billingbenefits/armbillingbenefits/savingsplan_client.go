//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbillingbenefits

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SavingsPlanClient contains the methods for the SavingsPlan group.
// Don't use this type directly, use NewSavingsPlanClient() instead.
type SavingsPlanClient struct {
	host   string
	expand *string
	pl     runtime.Pipeline
}

// NewSavingsPlanClient creates a new instance of SavingsPlanClient with the specified values.
// expand - May be used to expand the detail information of some properties.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSavingsPlanClient(expand *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SavingsPlanClient, error) {
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
	client := &SavingsPlanClient{
		expand: expand,
		host:   ep,
		pl:     pl,
	}
	return client, nil
}

// Get - Get savings plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// savingsPlanOrderID - Order ID of the savings plan
// savingsPlanID - ID of the savings plan
// options - SavingsPlanClientGetOptions contains the optional parameters for the SavingsPlanClient.Get method.
func (client *SavingsPlanClient) Get(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, options *SavingsPlanClientGetOptions) (SavingsPlanClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, savingsPlanOrderID, savingsPlanID, options)
	if err != nil {
		return SavingsPlanClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavingsPlanClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavingsPlanClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SavingsPlanClient) getCreateRequest(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, options *SavingsPlanClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans/{savingsPlanId}"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if savingsPlanID == "" {
		return nil, errors.New("parameter savingsPlanID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanId}", url.PathEscape(savingsPlanID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	if client.expand != nil {
		reqQP.Set("$expand", *client.expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SavingsPlanClient) getHandleResponse(resp *http.Response) (SavingsPlanClientGetResponse, error) {
	result := SavingsPlanClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanModel); err != nil {
		return SavingsPlanClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List savings plans in an order.
// Generated from API version 2022-11-01
// savingsPlanOrderID - Order ID of the savings plan
// options - SavingsPlanClientListOptions contains the optional parameters for the SavingsPlanClient.List method.
func (client *SavingsPlanClient) NewListPager(savingsPlanOrderID string, options *SavingsPlanClientListOptions) *runtime.Pager[SavingsPlanClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SavingsPlanClientListResponse]{
		More: func(page SavingsPlanClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SavingsPlanClientListResponse) (SavingsPlanClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, savingsPlanOrderID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SavingsPlanClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SavingsPlanClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SavingsPlanClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SavingsPlanClient) listCreateRequest(ctx context.Context, savingsPlanOrderID string, options *SavingsPlanClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SavingsPlanClient) listHandleResponse(resp *http.Response) (SavingsPlanClientListResponse, error) {
	result := SavingsPlanClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanModelList); err != nil {
		return SavingsPlanClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - List savings plans.
// Generated from API version 2022-11-01
// options - SavingsPlanClientListAllOptions contains the optional parameters for the SavingsPlanClient.ListAll method.
func (client *SavingsPlanClient) NewListAllPager(options *SavingsPlanClientListAllOptions) *runtime.Pager[SavingsPlanClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[SavingsPlanClientListAllResponse]{
		More: func(page SavingsPlanClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SavingsPlanClientListAllResponse) (SavingsPlanClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SavingsPlanClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SavingsPlanClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SavingsPlanClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *SavingsPlanClient) listAllCreateRequest(ctx context.Context, options *SavingsPlanClientListAllOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlans"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.RefreshSummary != nil {
		reqQP.Set("refreshSummary", *options.RefreshSummary)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", strconv.FormatFloat(float64(*options.Skiptoken), 'f', -1, 32))
	}
	if options != nil && options.SelectedState != nil {
		reqQP.Set("selectedState", *options.SelectedState)
	}
	if options != nil && options.Take != nil {
		reqQP.Set("take", strconv.FormatFloat(float64(*options.Take), 'f', -1, 32))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *SavingsPlanClient) listAllHandleResponse(resp *http.Response) (SavingsPlanClientListAllResponse, error) {
	result := SavingsPlanClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanModelListResult); err != nil {
		return SavingsPlanClientListAllResponse{}, err
	}
	return result, nil
}

// Update - Update savings plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// savingsPlanOrderID - Order ID of the savings plan
// savingsPlanID - ID of the savings plan
// body - Request body for patching a savings plan order alias
// options - SavingsPlanClientUpdateOptions contains the optional parameters for the SavingsPlanClient.Update method.
func (client *SavingsPlanClient) Update(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, body SavingsPlanUpdateRequest, options *SavingsPlanClientUpdateOptions) (SavingsPlanClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, savingsPlanOrderID, savingsPlanID, body, options)
	if err != nil {
		return SavingsPlanClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavingsPlanClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return SavingsPlanClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SavingsPlanClient) updateCreateRequest(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, body SavingsPlanUpdateRequest, options *SavingsPlanClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans/{savingsPlanId}"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if savingsPlanID == "" {
		return nil, errors.New("parameter savingsPlanID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanId}", url.PathEscape(savingsPlanID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *SavingsPlanClient) updateHandleResponse(resp *http.Response) (SavingsPlanClientUpdateResponse, error) {
	result := SavingsPlanClientUpdateResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanModel); err != nil {
		return SavingsPlanClientUpdateResponse{}, err
	}
	return result, nil
}

// ValidateUpdate - Validate savings plan patch.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// savingsPlanOrderID - Order ID of the savings plan
// savingsPlanID - ID of the savings plan
// body - Request body for validating a savings plan patch request
// options - SavingsPlanClientValidateUpdateOptions contains the optional parameters for the SavingsPlanClient.ValidateUpdate
// method.
func (client *SavingsPlanClient) ValidateUpdate(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, body SavingsPlanUpdateValidateRequest, options *SavingsPlanClientValidateUpdateOptions) (SavingsPlanClientValidateUpdateResponse, error) {
	req, err := client.validateUpdateCreateRequest(ctx, savingsPlanOrderID, savingsPlanID, body, options)
	if err != nil {
		return SavingsPlanClientValidateUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavingsPlanClientValidateUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavingsPlanClientValidateUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateUpdateHandleResponse(resp)
}

// validateUpdateCreateRequest creates the ValidateUpdate request.
func (client *SavingsPlanClient) validateUpdateCreateRequest(ctx context.Context, savingsPlanOrderID string, savingsPlanID string, body SavingsPlanUpdateValidateRequest, options *SavingsPlanClientValidateUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{savingsPlanOrderId}/savingsPlans/{savingsPlanId}/validate"
	if savingsPlanOrderID == "" {
		return nil, errors.New("parameter savingsPlanOrderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderId}", url.PathEscape(savingsPlanOrderID))
	if savingsPlanID == "" {
		return nil, errors.New("parameter savingsPlanID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanId}", url.PathEscape(savingsPlanID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// validateUpdateHandleResponse handles the ValidateUpdate response.
func (client *SavingsPlanClient) validateUpdateHandleResponse(resp *http.Response) (SavingsPlanClientValidateUpdateResponse, error) {
	result := SavingsPlanClientValidateUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanValidateResponse); err != nil {
		return SavingsPlanClientValidateUpdateResponse{}, err
	}
	return result, nil
}
