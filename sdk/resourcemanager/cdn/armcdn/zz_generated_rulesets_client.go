//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// RuleSetsClient contains the methods for the RuleSets group.
// Don't use this type directly, use NewRuleSetsClient() instead.
type RuleSetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRuleSetsClient creates a new instance of RuleSetsClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRuleSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RuleSetsClient, error) {
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
	client := &RuleSetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates a new rule set within the specified profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// ruleSetName - Name of the rule set under the profile which is unique globally
// options - RuleSetsClientCreateOptions contains the optional parameters for the RuleSetsClient.Create method.
func (client *RuleSetsClient) Create(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientCreateOptions) (RuleSetsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, options)
	if err != nil {
		return RuleSetsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RuleSetsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return RuleSetsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RuleSetsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *RuleSetsClient) createHandleResponse(resp *http.Response) (RuleSetsClientCreateResponse, error) {
	result := RuleSetsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleSet); err != nil {
		return RuleSetsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription,
// resource group and profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// ruleSetName - Name of the rule set under the profile which is unique globally.
// options - RuleSetsClientBeginDeleteOptions contains the optional parameters for the RuleSetsClient.BeginDelete method.
func (client *RuleSetsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientBeginDeleteOptions) (*runtime.Poller[RuleSetsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, ruleSetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[RuleSetsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[RuleSetsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription,
// resource group and profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *RuleSetsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RuleSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription, resource
// group and profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// ruleSetName - Name of the rule set under the profile which is unique globally.
// options - RuleSetsClientGetOptions contains the optional parameters for the RuleSetsClient.Get method.
func (client *RuleSetsClient) Get(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientGetOptions) (RuleSetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, options)
	if err != nil {
		return RuleSetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RuleSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RuleSetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RuleSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RuleSetsClient) getHandleResponse(resp *http.Response) (RuleSetsClientGetResponse, error) {
	result := RuleSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleSet); err != nil {
		return RuleSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProfilePager - Lists existing AzureFrontDoor rule sets within a profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// options - RuleSetsClientListByProfileOptions contains the optional parameters for the RuleSetsClient.ListByProfile method.
func (client *RuleSetsClient) NewListByProfilePager(resourceGroupName string, profileName string, options *RuleSetsClientListByProfileOptions) *runtime.Pager[RuleSetsClientListByProfileResponse] {
	return runtime.NewPager(runtime.PagingHandler[RuleSetsClientListByProfileResponse]{
		More: func(page RuleSetsClientListByProfileResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RuleSetsClientListByProfileResponse) (RuleSetsClientListByProfileResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProfileCreateRequest(ctx, resourceGroupName, profileName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RuleSetsClientListByProfileResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RuleSetsClientListByProfileResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RuleSetsClientListByProfileResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProfileHandleResponse(resp)
		},
	})
}

// listByProfileCreateRequest creates the ListByProfile request.
func (client *RuleSetsClient) listByProfileCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *RuleSetsClientListByProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProfileHandleResponse handles the ListByProfile response.
func (client *RuleSetsClient) listByProfileHandleResponse(resp *http.Response) (RuleSetsClientListByProfileResponse, error) {
	result := RuleSetsClientListByProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RuleSetListResult); err != nil {
		return RuleSetsClientListByProfileResponse{}, err
	}
	return result, nil
}

// NewListResourceUsagePager - Checks the quota and actual usage of the given AzureFrontDoor rule set under the given CDN
// profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// profileName - Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource
// group.
// ruleSetName - Name of the rule set under the profile which is unique globally.
// options - RuleSetsClientListResourceUsageOptions contains the optional parameters for the RuleSetsClient.ListResourceUsage
// method.
func (client *RuleSetsClient) NewListResourceUsagePager(resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientListResourceUsageOptions) *runtime.Pager[RuleSetsClientListResourceUsageResponse] {
	return runtime.NewPager(runtime.PagingHandler[RuleSetsClientListResourceUsageResponse]{
		More: func(page RuleSetsClientListResourceUsageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RuleSetsClientListResourceUsageResponse) (RuleSetsClientListResourceUsageResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listResourceUsageCreateRequest(ctx, resourceGroupName, profileName, ruleSetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RuleSetsClientListResourceUsageResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RuleSetsClientListResourceUsageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RuleSetsClientListResourceUsageResponse{}, runtime.NewResponseError(resp)
			}
			return client.listResourceUsageHandleResponse(resp)
		},
	})
}

// listResourceUsageCreateRequest creates the ListResourceUsage request.
func (client *RuleSetsClient) listResourceUsageCreateRequest(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *RuleSetsClientListResourceUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/ruleSets/{ruleSetName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listResourceUsageHandleResponse handles the ListResourceUsage response.
func (client *RuleSetsClient) listResourceUsageHandleResponse(resp *http.Response) (RuleSetsClientListResourceUsageResponse, error) {
	result := RuleSetsClientListResourceUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsagesListResult); err != nil {
		return RuleSetsClientListResourceUsageResponse{}, err
	}
	return result, nil
}
