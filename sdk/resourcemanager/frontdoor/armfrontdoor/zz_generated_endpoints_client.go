//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// EndpointsClient contains the methods for the Endpoints group.
// Don't use this type directly, use NewEndpointsClient() instead.
type EndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEndpointsClient creates a new instance of EndpointsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointsClient, error) {
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
	client := &EndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginPurgeContent - Removes a content from Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// contentFilePaths - The path to the content to be purged. Path can be a full URL, e.g. '/pictures/city.png' which removes
// a single file, or a directory with a wildcard, e.g. '/pictures/*' which removes all folders and
// files in the directory.
// options - EndpointsClientBeginPurgeContentOptions contains the optional parameters for the EndpointsClient.BeginPurgeContent
// method.
func (client *EndpointsClient) BeginPurgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters, options *EndpointsClientBeginPurgeContentOptions) (*runtime.Poller[EndpointsClientPurgeContentResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purgeContent(ctx, resourceGroupName, frontDoorName, contentFilePaths, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EndpointsClientPurgeContentResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EndpointsClientPurgeContentResponse](options.ResumeToken, client.pl, nil)
	}
}

// PurgeContent - Removes a content from Front Door.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-05-01
func (client *EndpointsClient) purgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters, options *EndpointsClientBeginPurgeContentOptions) (*http.Response, error) {
	req, err := client.purgeContentCreateRequest(ctx, resourceGroupName, frontDoorName, contentFilePaths, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// purgeContentCreateRequest creates the PurgeContent request.
func (client *EndpointsClient) purgeContentCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths PurgeParameters, options *EndpointsClientBeginPurgeContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/purge"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, contentFilePaths)
}
