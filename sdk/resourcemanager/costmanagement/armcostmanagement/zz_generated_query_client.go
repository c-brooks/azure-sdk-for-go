//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

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

// QueryClient contains the methods for the Query group.
// Don't use this type directly, use NewQueryClient() instead.
type QueryClient struct {
	host string
	pl   runtime.Pipeline
}

// NewQueryClient creates a new instance of QueryClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewQueryClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*QueryClient, error) {
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
	client := &QueryClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Usage - Query the usage data for scope defined.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// scope - The scope associated with query and export operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope,
// '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// parameters - Parameters supplied to the CreateOrUpdate Query Config operation.
// options - QueryClientUsageOptions contains the optional parameters for the QueryClient.Usage method.
func (client *QueryClient) Usage(ctx context.Context, scope string, parameters QueryDefinition, options *QueryClientUsageOptions) (QueryClientUsageResponse, error) {
	req, err := client.usageCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return QueryClientUsageResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueryClientUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return QueryClientUsageResponse{}, runtime.NewResponseError(resp)
	}
	return client.usageHandleResponse(resp)
}

// usageCreateRequest creates the Usage request.
func (client *QueryClient) usageCreateRequest(ctx context.Context, scope string, parameters QueryDefinition, options *QueryClientUsageOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/query"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// usageHandleResponse handles the Usage response.
func (client *QueryClient) usageHandleResponse(resp *http.Response) (QueryClientUsageResponse, error) {
	result := QueryClientUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResult); err != nil {
		return QueryClientUsageResponse{}, err
	}
	return result, nil
}

// UsageByExternalCloudProviderType - Query the usage data for external cloud provider type defined.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// externalCloudProviderType - The external cloud provider type associated with dimension/query operations. This includes
// 'externalSubscriptions' for linked account and 'externalBillingAccounts' for consolidated account.
// externalCloudProviderID - This can be '{externalSubscriptionId}' for linked account or '{externalBillingAccountId}' for
// consolidated account used with dimension/query operations.
// parameters - Parameters supplied to the CreateOrUpdate Query Config operation.
// options - QueryClientUsageByExternalCloudProviderTypeOptions contains the optional parameters for the QueryClient.UsageByExternalCloudProviderType
// method.
func (client *QueryClient) UsageByExternalCloudProviderType(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, parameters QueryDefinition, options *QueryClientUsageByExternalCloudProviderTypeOptions) (QueryClientUsageByExternalCloudProviderTypeResponse, error) {
	req, err := client.usageByExternalCloudProviderTypeCreateRequest(ctx, externalCloudProviderType, externalCloudProviderID, parameters, options)
	if err != nil {
		return QueryClientUsageByExternalCloudProviderTypeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return QueryClientUsageByExternalCloudProviderTypeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueryClientUsageByExternalCloudProviderTypeResponse{}, runtime.NewResponseError(resp)
	}
	return client.usageByExternalCloudProviderTypeHandleResponse(resp)
}

// usageByExternalCloudProviderTypeCreateRequest creates the UsageByExternalCloudProviderType request.
func (client *QueryClient) usageByExternalCloudProviderTypeCreateRequest(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, parameters QueryDefinition, options *QueryClientUsageByExternalCloudProviderTypeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/query"
	if externalCloudProviderType == "" {
		return nil, errors.New("parameter externalCloudProviderType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderType}", url.PathEscape(string(externalCloudProviderType)))
	if externalCloudProviderID == "" {
		return nil, errors.New("parameter externalCloudProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderId}", url.PathEscape(externalCloudProviderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// usageByExternalCloudProviderTypeHandleResponse handles the UsageByExternalCloudProviderType response.
func (client *QueryClient) usageByExternalCloudProviderTypeHandleResponse(resp *http.Response) (QueryClientUsageByExternalCloudProviderTypeResponse, error) {
	result := QueryClientUsageByExternalCloudProviderTypeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResult); err != nil {
		return QueryClientUsageByExternalCloudProviderTypeResponse{}, err
	}
	return result, nil
}
