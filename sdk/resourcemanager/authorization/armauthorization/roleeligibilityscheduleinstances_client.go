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

// RoleEligibilityScheduleInstancesClient contains the methods for the RoleEligibilityScheduleInstances group.
// Don't use this type directly, use NewRoleEligibilityScheduleInstancesClient() instead.
type RoleEligibilityScheduleInstancesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewRoleEligibilityScheduleInstancesClient creates a new instance of RoleEligibilityScheduleInstancesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleEligibilityScheduleInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleEligibilityScheduleInstancesClient, error) {
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
	client := &RoleEligibilityScheduleInstancesClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Gets the specified role eligibility schedule instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility schedules.
// roleEligibilityScheduleInstanceName - The name (hash of schedule name + time) of the role eligibility schedule to get.
// options - RoleEligibilityScheduleInstancesClientGetOptions contains the optional parameters for the RoleEligibilityScheduleInstancesClient.Get
// method.
func (client *RoleEligibilityScheduleInstancesClient) Get(ctx context.Context, scope string, roleEligibilityScheduleInstanceName string, options *RoleEligibilityScheduleInstancesClientGetOptions) (RoleEligibilityScheduleInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleInstanceName, options)
	if err != nil {
		return RoleEligibilityScheduleInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilityScheduleInstancesClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleInstanceName string, options *RoleEligibilityScheduleInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/{roleEligibilityScheduleInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleInstanceName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleInstanceName}", url.PathEscape(roleEligibilityScheduleInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilityScheduleInstancesClient) getHandleResponse(resp *http.Response) (RoleEligibilityScheduleInstancesClientGetResponse, error) {
	result := RoleEligibilityScheduleInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleInstance); err != nil {
		return RoleEligibilityScheduleInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role eligibility schedule instances of a role eligibility schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-10-01
// scope - The scope of the role eligibility schedule.
// options - RoleEligibilityScheduleInstancesClientListForScopeOptions contains the optional parameters for the RoleEligibilityScheduleInstancesClient.ListForScope
// method.
func (client *RoleEligibilityScheduleInstancesClient) NewListForScopePager(scope string, options *RoleEligibilityScheduleInstancesClientListForScopeOptions) *runtime.Pager[RoleEligibilityScheduleInstancesClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleEligibilityScheduleInstancesClientListForScopeResponse]{
		More: func(page RoleEligibilityScheduleInstancesClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleEligibilityScheduleInstancesClientListForScopeResponse) (RoleEligibilityScheduleInstancesClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleEligibilityScheduleInstancesClientListForScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RoleEligibilityScheduleInstancesClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleEligibilityScheduleInstancesClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilityScheduleInstancesClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeHandleResponse(resp *http.Response) (RoleEligibilityScheduleInstancesClientListForScopeResponse, error) {
	result := RoleEligibilityScheduleInstancesClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleInstanceListResult); err != nil {
		return RoleEligibilityScheduleInstancesClientListForScopeResponse{}, err
	}
	return result, nil
}
