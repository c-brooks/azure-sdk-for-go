//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

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

// ResourceGuardProxiesClient contains the methods for the ResourceGuardProxies group.
// Don't use this type directly, use NewResourceGuardProxiesClient() instead.
type ResourceGuardProxiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceGuardProxiesClient creates a new instance of ResourceGuardProxiesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceGuardProxiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceGuardProxiesClient, error) {
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
	client := &ResourceGuardProxiesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewGetPager - List the ResourceGuardProxies under vault
//
// Generated from API version 2023-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ResourceGuardProxiesClientGetOptions contains the optional parameters for the ResourceGuardProxiesClient.NewGetPager
//     method.
func (client *ResourceGuardProxiesClient) NewGetPager(vaultName string, resourceGroupName string, options *ResourceGuardProxiesClientGetOptions) *runtime.Pager[ResourceGuardProxiesClientGetResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceGuardProxiesClientGetResponse]{
		More: func(page ResourceGuardProxiesClientGetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceGuardProxiesClientGetResponse) (ResourceGuardProxiesClientGetResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getCreateRequest(ctx, vaultName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceGuardProxiesClientGetResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceGuardProxiesClientGetResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceGuardProxiesClientGetResponse{}, runtime.NewResponseError(resp)
			}
			return client.getHandleResponse(resp)
		},
	})
}

// getCreateRequest creates the Get request.
func (client *ResourceGuardProxiesClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *ResourceGuardProxiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceGuardProxiesClient) getHandleResponse(resp *http.Response) (ResourceGuardProxiesClientGetResponse, error) {
	result := ResourceGuardProxiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResourceList); err != nil {
		return ResourceGuardProxiesClientGetResponse{}, err
	}
	return result, nil
}
