//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

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

// MetricAlertsStatusClient contains the methods for the MetricAlertsStatus group.
// Don't use this type directly, use NewMetricAlertsStatusClient() instead.
type MetricAlertsStatusClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMetricAlertsStatusClient creates a new instance of MetricAlertsStatusClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMetricAlertsStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricAlertsStatusClient, error) {
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
	client := &MetricAlertsStatusClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Retrieve an alert rule status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ruleName - The name of the rule.
// options - MetricAlertsStatusClientListOptions contains the optional parameters for the MetricAlertsStatusClient.List method.
func (client *MetricAlertsStatusClient) List(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsStatusClientListOptions) (MetricAlertsStatusClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, ruleName, options)
	if err != nil {
		return MetricAlertsStatusClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricAlertsStatusClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricAlertsStatusClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MetricAlertsStatusClient) listCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, options *MetricAlertsStatusClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}/status"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricAlertsStatusClient) listHandleResponse(resp *http.Response) (MetricAlertsStatusClientListResponse, error) {
	result := MetricAlertsStatusClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertStatusCollection); err != nil {
		return MetricAlertsStatusClientListResponse{}, err
	}
	return result, nil
}

// ListByName - Retrieve an alert rule status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-03-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// ruleName - The name of the rule.
// statusName - The name of the status.
// options - MetricAlertsStatusClientListByNameOptions contains the optional parameters for the MetricAlertsStatusClient.ListByName
// method.
func (client *MetricAlertsStatusClient) ListByName(ctx context.Context, resourceGroupName string, ruleName string, statusName string, options *MetricAlertsStatusClientListByNameOptions) (MetricAlertsStatusClientListByNameResponse, error) {
	req, err := client.listByNameCreateRequest(ctx, resourceGroupName, ruleName, statusName, options)
	if err != nil {
		return MetricAlertsStatusClientListByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricAlertsStatusClientListByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricAlertsStatusClientListByNameResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByNameHandleResponse(resp)
}

// listByNameCreateRequest creates the ListByName request.
func (client *MetricAlertsStatusClient) listByNameCreateRequest(ctx context.Context, resourceGroupName string, ruleName string, statusName string, options *MetricAlertsStatusClientListByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/metricAlerts/{ruleName}/status/{statusName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ruleName == "" {
		return nil, errors.New("parameter ruleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleName}", url.PathEscape(ruleName))
	if statusName == "" {
		return nil, errors.New("parameter statusName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{statusName}", url.PathEscape(statusName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNameHandleResponse handles the ListByName response.
func (client *MetricAlertsStatusClient) listByNameHandleResponse(resp *http.Response) (MetricAlertsStatusClientListByNameResponse, error) {
	result := MetricAlertsStatusClientListByNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricAlertStatusCollection); err != nil {
		return MetricAlertsStatusClientListByNameResponse{}, err
	}
	return result, nil
}
