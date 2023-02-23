//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// PropertyClient contains the methods for the BillingProperty group.
// Don't use this type directly, use NewPropertyClient() instead.
type PropertyClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPropertyClient creates a new instance of PropertyClient with the specified values.
// subscriptionID - The ID that uniquely identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPropertyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PropertyClient, error) {
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
	client := &PropertyClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get the billing properties for a subscription. This operation is not supported for billing accounts with agreement
// type Enterprise Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// options - PropertyClientGetOptions contains the optional parameters for the PropertyClient.Get method.
func (client *PropertyClient) Get(ctx context.Context, options *PropertyClientGetOptions) (PropertyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return PropertyClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PropertyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PropertyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PropertyClient) getCreateRequest(ctx context.Context, options *PropertyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingProperty/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PropertyClient) getHandleResponse(resp *http.Response) (PropertyClientGetResponse, error) {
	result := PropertyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Property); err != nil {
		return PropertyClientGetResponse{}, err
	}
	return result, nil
}

// Update - Updates the billing property of a subscription. Currently, cost center can be updated. The operation is supported
// only for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// parameters - Request parameters that are provided to the update billing property operation.
// options - PropertyClientUpdateOptions contains the optional parameters for the PropertyClient.Update method.
func (client *PropertyClient) Update(ctx context.Context, parameters Property, options *PropertyClientUpdateOptions) (PropertyClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, parameters, options)
	if err != nil {
		return PropertyClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PropertyClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PropertyClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PropertyClient) updateCreateRequest(ctx context.Context, parameters Property, options *PropertyClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingProperty/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *PropertyClient) updateHandleResponse(resp *http.Response) (PropertyClientUpdateResponse, error) {
	result := PropertyClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Property); err != nil {
		return PropertyClientUpdateResponse{}, err
	}
	return result, nil
}
