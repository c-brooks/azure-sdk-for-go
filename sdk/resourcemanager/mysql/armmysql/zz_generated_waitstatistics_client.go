//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// WaitStatisticsClient contains the methods for the WaitStatistics group.
// Don't use this type directly, use NewWaitStatisticsClient() instead.
type WaitStatisticsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWaitStatisticsClient creates a new instance of WaitStatisticsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWaitStatisticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WaitStatisticsClient, error) {
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
	client := &WaitStatisticsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Retrieve wait statistics for specified identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// waitStatisticsID - The Wait Statistic identifier.
// options - WaitStatisticsClientGetOptions contains the optional parameters for the WaitStatisticsClient.Get method.
func (client *WaitStatisticsClient) Get(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string, options *WaitStatisticsClientGetOptions) (WaitStatisticsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, waitStatisticsID, options)
	if err != nil {
		return WaitStatisticsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WaitStatisticsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WaitStatisticsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WaitStatisticsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string, options *WaitStatisticsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics/{waitStatisticsId}"
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
	if waitStatisticsID == "" {
		return nil, errors.New("parameter waitStatisticsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{waitStatisticsId}", url.PathEscape(waitStatisticsID))
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
func (client *WaitStatisticsClient) getHandleResponse(resp *http.Response) (WaitStatisticsClientGetResponse, error) {
	result := WaitStatisticsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WaitStatistic); err != nil {
		return WaitStatisticsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Retrieve wait statistics for specified aggregation window.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// parameters - The required parameters for retrieving wait statistics.
// options - WaitStatisticsClientListByServerOptions contains the optional parameters for the WaitStatisticsClient.ListByServer
// method.
func (client *WaitStatisticsClient) NewListByServerPager(resourceGroupName string, serverName string, parameters WaitStatisticsInput, options *WaitStatisticsClientListByServerOptions) *runtime.Pager[WaitStatisticsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[WaitStatisticsClientListByServerResponse]{
		More: func(page WaitStatisticsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WaitStatisticsClientListByServerResponse) (WaitStatisticsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WaitStatisticsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WaitStatisticsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WaitStatisticsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *WaitStatisticsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters WaitStatisticsInput, options *WaitStatisticsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics"
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
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listByServerHandleResponse handles the ListByServer response.
func (client *WaitStatisticsClient) listByServerHandleResponse(resp *http.Response) (WaitStatisticsClientListByServerResponse, error) {
	result := WaitStatisticsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WaitStatisticsResultList); err != nil {
		return WaitStatisticsClientListByServerResponse{}, err
	}
	return result, nil
}
