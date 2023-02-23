//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

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

// SubscriptionsClient contains the methods for the Subscriptions group.
// Don't use this type directly, use NewSubscriptionsClient() instead.
type SubscriptionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSubscriptionsClient creates a new instance of SubscriptionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubscriptionsClient, error) {
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
	client := &SubscriptionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// ListQuotas - Retrieves the subscription's current quota information in a particular region.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// location - The region in which to retrieve the subscription's quota information. You can find out which regions Azure Stream
// Analytics is supported in here: https://azure.microsoft.com/en-us/regions/
// options - SubscriptionsClientListQuotasOptions contains the optional parameters for the SubscriptionsClient.ListQuotas
// method.
func (client *SubscriptionsClient) ListQuotas(ctx context.Context, location string, options *SubscriptionsClientListQuotasOptions) (SubscriptionsClientListQuotasResponse, error) {
	req, err := client.listQuotasCreateRequest(ctx, location, options)
	if err != nil {
		return SubscriptionsClientListQuotasResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionsClientListQuotasResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionsClientListQuotasResponse{}, runtime.NewResponseError(resp)
	}
	return client.listQuotasHandleResponse(resp)
}

// listQuotasCreateRequest creates the ListQuotas request.
func (client *SubscriptionsClient) listQuotasCreateRequest(ctx context.Context, location string, options *SubscriptionsClientListQuotasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/quotas"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQuotasHandleResponse handles the ListQuotas response.
func (client *SubscriptionsClient) listQuotasHandleResponse(resp *http.Response) (SubscriptionsClientListQuotasResponse, error) {
	result := SubscriptionsClientListQuotasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionQuotasListResult); err != nil {
		return SubscriptionsClientListQuotasResponse{}, err
	}
	return result, nil
}
