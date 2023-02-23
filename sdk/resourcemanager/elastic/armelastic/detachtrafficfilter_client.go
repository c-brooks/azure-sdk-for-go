//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armelastic

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

// DetachTrafficFilterClient contains the methods for the DetachTrafficFilter group.
// Don't use this type directly, use NewDetachTrafficFilterClient() instead.
type DetachTrafficFilterClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDetachTrafficFilterClient creates a new instance of DetachTrafficFilterClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDetachTrafficFilterClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DetachTrafficFilterClient, error) {
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
	client := &DetachTrafficFilterClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginUpdate - Detach traffic filter for the given deployment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// options - DetachTrafficFilterClientBeginUpdateOptions contains the optional parameters for the DetachTrafficFilterClient.BeginUpdate
// method.
func (client *DetachTrafficFilterClient) BeginUpdate(ctx context.Context, resourceGroupName string, monitorName string, options *DetachTrafficFilterClientBeginUpdateOptions) (*runtime.Poller[DetachTrafficFilterClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, monitorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DetachTrafficFilterClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DetachTrafficFilterClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Detach traffic filter for the given deployment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *DetachTrafficFilterClient) update(ctx context.Context, resourceGroupName string, monitorName string, options *DetachTrafficFilterClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DetachTrafficFilterClient) updateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *DetachTrafficFilterClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/detachTrafficFilter"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	if options != nil && options.RulesetID != nil {
		reqQP.Set("rulesetId", *options.RulesetID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
