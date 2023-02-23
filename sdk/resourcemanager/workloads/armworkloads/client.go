//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

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

// Client contains the methods for the WorkloadsClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
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
	client := &Client{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// SAPAvailabilityZoneDetails - Get SAP Availability Zone Details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// location - The name of Azure region.
// options - ClientSAPAvailabilityZoneDetailsOptions contains the optional parameters for the Client.SAPAvailabilityZoneDetails
// method.
func (client *Client) SAPAvailabilityZoneDetails(ctx context.Context, location string, options *ClientSAPAvailabilityZoneDetailsOptions) (ClientSAPAvailabilityZoneDetailsResponse, error) {
	req, err := client.sapAvailabilityZoneDetailsCreateRequest(ctx, location, options)
	if err != nil {
		return ClientSAPAvailabilityZoneDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientSAPAvailabilityZoneDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientSAPAvailabilityZoneDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.sapAvailabilityZoneDetailsHandleResponse(resp)
}

// sapAvailabilityZoneDetailsCreateRequest creates the SAPAvailabilityZoneDetails request.
func (client *Client) sapAvailabilityZoneDetailsCreateRequest(ctx context.Context, location string, options *ClientSAPAvailabilityZoneDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/locations/{location}/sapVirtualInstanceMetadata/default/getAvailabilityZoneDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SAPAvailabilityZoneDetails != nil {
		return req, runtime.MarshalAsJSON(req, *options.SAPAvailabilityZoneDetails)
	}
	return req, nil
}

// sapAvailabilityZoneDetailsHandleResponse handles the SAPAvailabilityZoneDetails response.
func (client *Client) sapAvailabilityZoneDetailsHandleResponse(resp *http.Response) (ClientSAPAvailabilityZoneDetailsResponse, error) {
	result := ClientSAPAvailabilityZoneDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPAvailabilityZoneDetailsResult); err != nil {
		return ClientSAPAvailabilityZoneDetailsResponse{}, err
	}
	return result, nil
}

// SAPDiskConfigurations - Get SAP Disk Configurations.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// location - The name of Azure region.
// options - ClientSAPDiskConfigurationsOptions contains the optional parameters for the Client.SAPDiskConfigurations method.
func (client *Client) SAPDiskConfigurations(ctx context.Context, location string, options *ClientSAPDiskConfigurationsOptions) (ClientSAPDiskConfigurationsResponse, error) {
	req, err := client.sapDiskConfigurationsCreateRequest(ctx, location, options)
	if err != nil {
		return ClientSAPDiskConfigurationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientSAPDiskConfigurationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientSAPDiskConfigurationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.sapDiskConfigurationsHandleResponse(resp)
}

// sapDiskConfigurationsCreateRequest creates the SAPDiskConfigurations request.
func (client *Client) sapDiskConfigurationsCreateRequest(ctx context.Context, location string, options *ClientSAPDiskConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/locations/{location}/sapVirtualInstanceMetadata/default/getDiskConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SAPDiskConfigurations != nil {
		return req, runtime.MarshalAsJSON(req, *options.SAPDiskConfigurations)
	}
	return req, nil
}

// sapDiskConfigurationsHandleResponse handles the SAPDiskConfigurations response.
func (client *Client) sapDiskConfigurationsHandleResponse(resp *http.Response) (ClientSAPDiskConfigurationsResponse, error) {
	result := ClientSAPDiskConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPDiskConfigurationsResult); err != nil {
		return ClientSAPDiskConfigurationsResponse{}, err
	}
	return result, nil
}

// SAPSizingRecommendations - Get SAP sizing recommendations.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// location - The name of Azure region.
// options - ClientSAPSizingRecommendationsOptions contains the optional parameters for the Client.SAPSizingRecommendations
// method.
func (client *Client) SAPSizingRecommendations(ctx context.Context, location string, options *ClientSAPSizingRecommendationsOptions) (ClientSAPSizingRecommendationsResponse, error) {
	req, err := client.sapSizingRecommendationsCreateRequest(ctx, location, options)
	if err != nil {
		return ClientSAPSizingRecommendationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientSAPSizingRecommendationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientSAPSizingRecommendationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.sapSizingRecommendationsHandleResponse(resp)
}

// sapSizingRecommendationsCreateRequest creates the SAPSizingRecommendations request.
func (client *Client) sapSizingRecommendationsCreateRequest(ctx context.Context, location string, options *ClientSAPSizingRecommendationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/locations/{location}/sapVirtualInstanceMetadata/default/getSizingRecommendations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SAPSizingRecommendation != nil {
		return req, runtime.MarshalAsJSON(req, *options.SAPSizingRecommendation)
	}
	return req, nil
}

// sapSizingRecommendationsHandleResponse handles the SAPSizingRecommendations response.
func (client *Client) sapSizingRecommendationsHandleResponse(resp *http.Response) (ClientSAPSizingRecommendationsResponse, error) {
	result := ClientSAPSizingRecommendationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ClientSAPSizingRecommendationsResponse{}, err
	}
	return result, nil
}

// SAPSupportedSKU - Get SAP supported SKUs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// location - The name of Azure region.
// options - ClientSAPSupportedSKUOptions contains the optional parameters for the Client.SAPSupportedSKU method.
func (client *Client) SAPSupportedSKU(ctx context.Context, location string, options *ClientSAPSupportedSKUOptions) (ClientSAPSupportedSKUResponse, error) {
	req, err := client.sapSupportedSKUCreateRequest(ctx, location, options)
	if err != nil {
		return ClientSAPSupportedSKUResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientSAPSupportedSKUResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientSAPSupportedSKUResponse{}, runtime.NewResponseError(resp)
	}
	return client.sapSupportedSKUHandleResponse(resp)
}

// sapSupportedSKUCreateRequest creates the SAPSupportedSKU request.
func (client *Client) sapSupportedSKUCreateRequest(ctx context.Context, location string, options *ClientSAPSupportedSKUOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/locations/{location}/sapVirtualInstanceMetadata/default/getSapSupportedSku"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SAPSupportedSKU != nil {
		return req, runtime.MarshalAsJSON(req, *options.SAPSupportedSKU)
	}
	return req, nil
}

// sapSupportedSKUHandleResponse handles the SAPSupportedSKU response.
func (client *Client) sapSupportedSKUHandleResponse(resp *http.Response) (ClientSAPSupportedSKUResponse, error) {
	result := ClientSAPSupportedSKUResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPSupportedResourceSKUsResult); err != nil {
		return ClientSAPSupportedSKUResponse{}, err
	}
	return result, nil
}
