//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// LocationsClient contains the methods for the PeeringLocations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
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
	client := &LocationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Lists all of the available peering locations for the specified kind of peering.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// kind - The kind of the peering.
// options - LocationsClientListOptions contains the optional parameters for the LocationsClient.List method.
func (client *LocationsClient) NewListPager(kind PeeringLocationsKind, options *LocationsClientListOptions) *runtime.Pager[LocationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationsClientListResponse]{
		More: func(page LocationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationsClientListResponse) (LocationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, kind, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LocationsClient) listCreateRequest(ctx context.Context, kind PeeringLocationsKind, options *LocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("kind", string(kind))
	if options != nil && options.DirectPeeringType != nil {
		reqQP.Set("directPeeringType", string(*options.DirectPeeringType))
	}
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationsClient) listHandleResponse(resp *http.Response) (LocationsClientListResponse, error) {
	result := LocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationListResult); err != nil {
		return LocationsClientListResponse{}, err
	}
	return result, nil
}
