//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// InformationProtectionPoliciesClient contains the methods for the InformationProtectionPolicies group.
// Don't use this type directly, use NewInformationProtectionPoliciesClient() instead.
type InformationProtectionPoliciesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewInformationProtectionPoliciesClient creates a new instance of InformationProtectionPoliciesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInformationProtectionPoliciesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*InformationProtectionPoliciesClient, error) {
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
	client := &InformationProtectionPoliciesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// CreateOrUpdate - Details of the information protection policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-08-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// informationProtectionPolicyName - Name of the information protection policy.
// informationProtectionPolicy - Information protection policy.
// options - InformationProtectionPoliciesClientCreateOrUpdateOptions contains the optional parameters for the InformationProtectionPoliciesClient.CreateOrUpdate
// method.
func (client *InformationProtectionPoliciesClient) CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, informationProtectionPolicy InformationProtectionPolicy, options *InformationProtectionPoliciesClientCreateOrUpdateOptions) (InformationProtectionPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, informationProtectionPolicyName, informationProtectionPolicy, options)
	if err != nil {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InformationProtectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, informationProtectionPolicy InformationProtectionPolicy, options *InformationProtectionPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if informationProtectionPolicyName == "" {
		return nil, errors.New("parameter informationProtectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{informationProtectionPolicyName}", url.PathEscape(string(informationProtectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, informationProtectionPolicy)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *InformationProtectionPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (InformationProtectionPoliciesClientCreateOrUpdateResponse, error) {
	result := InformationProtectionPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicy); err != nil {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Details of the information protection policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-08-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// informationProtectionPolicyName - Name of the information protection policy.
// options - InformationProtectionPoliciesClientGetOptions contains the optional parameters for the InformationProtectionPoliciesClient.Get
// method.
func (client *InformationProtectionPoliciesClient) Get(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, options *InformationProtectionPoliciesClientGetOptions) (InformationProtectionPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, informationProtectionPolicyName, options)
	if err != nil {
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InformationProtectionPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InformationProtectionPoliciesClient) getCreateRequest(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, options *InformationProtectionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if informationProtectionPolicyName == "" {
		return nil, errors.New("parameter informationProtectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{informationProtectionPolicyName}", url.PathEscape(string(informationProtectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InformationProtectionPoliciesClient) getHandleResponse(resp *http.Response) (InformationProtectionPoliciesClientGetResponse, error) {
	result := InformationProtectionPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicy); err != nil {
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Information protection policies of a specific management group.
// Generated from API version 2017-08-01-preview
// scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
// (/providers/Microsoft.Management/managementGroups/mgName).
// options - InformationProtectionPoliciesClientListOptions contains the optional parameters for the InformationProtectionPoliciesClient.List
// method.
func (client *InformationProtectionPoliciesClient) NewListPager(scope string, options *InformationProtectionPoliciesClientListOptions) *runtime.Pager[InformationProtectionPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InformationProtectionPoliciesClientListResponse]{
		More: func(page InformationProtectionPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InformationProtectionPoliciesClientListResponse) (InformationProtectionPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InformationProtectionPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InformationProtectionPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InformationProtectionPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *InformationProtectionPoliciesClient) listCreateRequest(ctx context.Context, scope string, options *InformationProtectionPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InformationProtectionPoliciesClient) listHandleResponse(resp *http.Response) (InformationProtectionPoliciesClientListResponse, error) {
	result := InformationProtectionPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicyList); err != nil {
		return InformationProtectionPoliciesClientListResponse{}, err
	}
	return result, nil
}
