//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmaintenance

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

// ConfigurationsForResourceGroupClient contains the methods for the MaintenanceConfigurationsForResourceGroup group.
// Don't use this type directly, use NewConfigurationsForResourceGroupClient() instead.
type ConfigurationsForResourceGroupClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigurationsForResourceGroupClient creates a new instance of ConfigurationsForResourceGroupClient with the specified values.
// subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigurationsForResourceGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationsForResourceGroupClient, error) {
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
	client := &ConfigurationsForResourceGroupClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Get Configuration records within a subscription and resource group
// Generated from API version 2022-07-01-preview
// resourceGroupName - Resource Group Name
// options - ConfigurationsForResourceGroupClientListOptions contains the optional parameters for the ConfigurationsForResourceGroupClient.List
// method.
func (client *ConfigurationsForResourceGroupClient) NewListPager(resourceGroupName string, options *ConfigurationsForResourceGroupClientListOptions) *runtime.Pager[ConfigurationsForResourceGroupClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationsForResourceGroupClientListResponse]{
		More: func(page ConfigurationsForResourceGroupClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationsForResourceGroupClientListResponse) (ConfigurationsForResourceGroupClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ConfigurationsForResourceGroupClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConfigurationsForResourceGroupClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationsForResourceGroupClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConfigurationsForResourceGroupClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationsForResourceGroupClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maintenance/maintenanceConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationsForResourceGroupClient) listHandleResponse(resp *http.Response) (ConfigurationsForResourceGroupClientListResponse, error) {
	result := ConfigurationsForResourceGroupClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListMaintenanceConfigurationsResult); err != nil {
		return ConfigurationsForResourceGroupClientListResponse{}, err
	}
	return result, nil
}
