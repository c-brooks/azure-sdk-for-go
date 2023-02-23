//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// ComponentFeatureCapabilitiesClient contains the methods for the ComponentFeatureCapabilities group.
// Don't use this type directly, use NewComponentFeatureCapabilitiesClient() instead.
type ComponentFeatureCapabilitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewComponentFeatureCapabilitiesClient creates a new instance of ComponentFeatureCapabilitiesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewComponentFeatureCapabilitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComponentFeatureCapabilitiesClient, error) {
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
	client := &ComponentFeatureCapabilitiesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns feature capabilities of the application insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - ComponentFeatureCapabilitiesClientGetOptions contains the optional parameters for the ComponentFeatureCapabilitiesClient.Get
// method.
func (client *ComponentFeatureCapabilitiesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentFeatureCapabilitiesClientGetOptions) (ComponentFeatureCapabilitiesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ComponentFeatureCapabilitiesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentFeatureCapabilitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentFeatureCapabilitiesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ComponentFeatureCapabilitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentFeatureCapabilitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/featurecapabilities"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComponentFeatureCapabilitiesClient) getHandleResponse(resp *http.Response) (ComponentFeatureCapabilitiesClientGetResponse, error) {
	result := ComponentFeatureCapabilitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentFeatureCapabilities); err != nil {
		return ComponentFeatureCapabilitiesClientGetResponse{}, err
	}
	return result, nil
}
