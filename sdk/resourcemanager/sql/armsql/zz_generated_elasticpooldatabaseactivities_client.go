//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ElasticPoolDatabaseActivitiesClient contains the methods for the ElasticPoolDatabaseActivities group.
// Don't use this type directly, use NewElasticPoolDatabaseActivitiesClient() instead.
type ElasticPoolDatabaseActivitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewElasticPoolDatabaseActivitiesClient creates a new instance of ElasticPoolDatabaseActivitiesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewElasticPoolDatabaseActivitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ElasticPoolDatabaseActivitiesClient, error) {
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
	client := &ElasticPoolDatabaseActivitiesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByElasticPoolPager - Returns activity on databases inside of an elastic pool.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// elasticPoolName - The name of the elastic pool.
// options - ElasticPoolDatabaseActivitiesClientListByElasticPoolOptions contains the optional parameters for the ElasticPoolDatabaseActivitiesClient.ListByElasticPool
// method.
func (client *ElasticPoolDatabaseActivitiesClient) NewListByElasticPoolPager(resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolDatabaseActivitiesClientListByElasticPoolOptions) *runtime.Pager[ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse]{
		More: func(page ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse) (ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse, error) {
			req, err := client.listByElasticPoolCreateRequest(ctx, resourceGroupName, serverName, elasticPoolName, options)
			if err != nil {
				return ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByElasticPoolHandleResponse(resp)
		},
	})
}

// listByElasticPoolCreateRequest creates the ListByElasticPool request.
func (client *ElasticPoolDatabaseActivitiesClient) listByElasticPoolCreateRequest(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, options *ElasticPoolDatabaseActivitiesClientListByElasticPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolDatabaseActivity"
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
	if elasticPoolName == "" {
		return nil, errors.New("parameter elasticPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elasticPoolName}", url.PathEscape(elasticPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByElasticPoolHandleResponse handles the ListByElasticPool response.
func (client *ElasticPoolDatabaseActivitiesClient) listByElasticPoolHandleResponse(resp *http.Response) (ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse, error) {
	result := ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ElasticPoolDatabaseActivityListResult); err != nil {
		return ElasticPoolDatabaseActivitiesClientListByElasticPoolResponse{}, err
	}
	return result, nil
}
