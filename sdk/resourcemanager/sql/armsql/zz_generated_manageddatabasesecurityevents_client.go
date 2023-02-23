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
	"strconv"
	"strings"
)

// ManagedDatabaseSecurityEventsClient contains the methods for the ManagedDatabaseSecurityEvents group.
// Don't use this type directly, use NewManagedDatabaseSecurityEventsClient() instead.
type ManagedDatabaseSecurityEventsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedDatabaseSecurityEventsClient creates a new instance of ManagedDatabaseSecurityEventsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedDatabaseSecurityEventsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseSecurityEventsClient, error) {
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
	client := &ManagedDatabaseSecurityEventsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByDatabasePager - Gets a list of security events.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// databaseName - The name of the managed database for which the security events are retrieved.
// options - ManagedDatabaseSecurityEventsClientListByDatabaseOptions contains the optional parameters for the ManagedDatabaseSecurityEventsClient.ListByDatabase
// method.
func (client *ManagedDatabaseSecurityEventsClient) NewListByDatabasePager(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSecurityEventsClientListByDatabaseOptions) *runtime.Pager[ManagedDatabaseSecurityEventsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedDatabaseSecurityEventsClientListByDatabaseResponse]{
		More: func(page ManagedDatabaseSecurityEventsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseSecurityEventsClientListByDatabaseResponse) (ManagedDatabaseSecurityEventsClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabaseSecurityEventsClientListByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ManagedDatabaseSecurityEventsClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabaseSecurityEventsClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedDatabaseSecurityEventsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSecurityEventsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/securityEvents"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedDatabaseSecurityEventsClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseSecurityEventsClientListByDatabaseResponse, error) {
	result := ManagedDatabaseSecurityEventsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityEventCollection); err != nil {
		return ManagedDatabaseSecurityEventsClientListByDatabaseResponse{}, err
	}
	return result, nil
}
