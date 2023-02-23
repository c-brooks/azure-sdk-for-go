//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

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

// AdvisorsClient contains the methods for the Advisors group.
// Don't use this type directly, use NewAdvisorsClient() instead.
type AdvisorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAdvisorsClient creates a new instance of AdvisorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAdvisorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AdvisorsClient, error) {
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
	client := &AdvisorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a recommendation action advisor.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// advisorName - The advisor name for recommendation action.
// options - AdvisorsClientGetOptions contains the optional parameters for the AdvisorsClient.Get method.
func (client *AdvisorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *AdvisorsClientGetOptions) (AdvisorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, advisorName, options)
	if err != nil {
		return AdvisorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AdvisorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AdvisorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AdvisorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *AdvisorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/advisors/{advisorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdvisorsClient) getHandleResponse(resp *http.Response) (AdvisorsClientGetResponse, error) {
	result := AdvisorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Advisor); err != nil {
		return AdvisorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - List recommendation action advisors.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - AdvisorsClientListByServerOptions contains the optional parameters for the AdvisorsClient.ListByServer method.
func (client *AdvisorsClient) NewListByServerPager(resourceGroupName string, serverName string, options *AdvisorsClientListByServerOptions) *runtime.Pager[AdvisorsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[AdvisorsClientListByServerResponse]{
		More: func(page AdvisorsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AdvisorsClientListByServerResponse) (AdvisorsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AdvisorsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AdvisorsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AdvisorsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *AdvisorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *AdvisorsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/advisors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *AdvisorsClient) listByServerHandleResponse(resp *http.Response) (AdvisorsClientListByServerResponse, error) {
	result := AdvisorsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdvisorsResultList); err != nil {
		return AdvisorsClientListByServerResponse{}, err
	}
	return result, nil
}
