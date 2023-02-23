//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// RelationshipLinksClient contains the methods for the RelationshipLinks group.
// Don't use this type directly, use NewRelationshipLinksClient() instead.
type RelationshipLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRelationshipLinksClient creates a new instance of RelationshipLinksClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRelationshipLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RelationshipLinksClient, error) {
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
	client := &RelationshipLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a relationship link or updates an existing relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// relationshipLinkName - The name of the relationship link.
// parameters - Parameters supplied to the CreateOrUpdate relationship link operation.
// options - RelationshipLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the RelationshipLinksClient.BeginCreateOrUpdate
// method.
func (client *RelationshipLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, options *RelationshipLinksClientBeginCreateOrUpdateOptions) (*runtime.Poller[RelationshipLinksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, relationshipLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RelationshipLinksClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[RelationshipLinksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a relationship link or updates an existing relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *RelationshipLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, options *RelationshipLinksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, relationshipLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RelationshipLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, parameters RelationshipLinkResourceFormat, options *RelationshipLinksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipLinkName == "" {
		return nil, errors.New("parameter relationshipLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipLinkName}", url.PathEscape(relationshipLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// relationshipLinkName - The name of the relationship.
// options - RelationshipLinksClientBeginDeleteOptions contains the optional parameters for the RelationshipLinksClient.BeginDelete
// method.
func (client *RelationshipLinksClient) BeginDelete(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientBeginDeleteOptions) (*runtime.Poller[RelationshipLinksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hubName, relationshipLinkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RelationshipLinksClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[RelationshipLinksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a relationship link within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *RelationshipLinksClient) deleteOperation(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, relationshipLinkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RelationshipLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipLinkName == "" {
		return nil, errors.New("parameter relationshipLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipLinkName}", url.PathEscape(relationshipLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about the specified relationship Link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// relationshipLinkName - The name of the relationship link.
// options - RelationshipLinksClientGetOptions contains the optional parameters for the RelationshipLinksClient.Get method.
func (client *RelationshipLinksClient) Get(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientGetOptions) (RelationshipLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, relationshipLinkName, options)
	if err != nil {
		return RelationshipLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RelationshipLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RelationshipLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RelationshipLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipLinkName string, options *RelationshipLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipLinkName == "" {
		return nil, errors.New("parameter relationshipLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipLinkName}", url.PathEscape(relationshipLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RelationshipLinksClient) getHandleResponse(resp *http.Response) (RelationshipLinksClientGetResponse, error) {
	result := RelationshipLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationshipLinkResourceFormat); err != nil {
		return RelationshipLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all relationship links in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - RelationshipLinksClientListByHubOptions contains the optional parameters for the RelationshipLinksClient.ListByHub
// method.
func (client *RelationshipLinksClient) NewListByHubPager(resourceGroupName string, hubName string, options *RelationshipLinksClientListByHubOptions) *runtime.Pager[RelationshipLinksClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[RelationshipLinksClientListByHubResponse]{
		More: func(page RelationshipLinksClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RelationshipLinksClientListByHubResponse) (RelationshipLinksClientListByHubResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RelationshipLinksClientListByHubResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RelationshipLinksClientListByHubResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RelationshipLinksClientListByHubResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHubHandleResponse(resp)
		},
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *RelationshipLinksClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *RelationshipLinksClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *RelationshipLinksClient) listByHubHandleResponse(resp *http.Response) (RelationshipLinksClientListByHubResponse, error) {
	result := RelationshipLinksClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationshipLinkListResult); err != nil {
		return RelationshipLinksClientListByHubResponse{}, err
	}
	return result, nil
}
