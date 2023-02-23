//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// UserConfirmationPasswordClient contains the methods for the UserConfirmationPassword group.
// Don't use this type directly, use NewUserConfirmationPasswordClient() instead.
type UserConfirmationPasswordClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewUserConfirmationPasswordClient creates a new instance of UserConfirmationPasswordClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUserConfirmationPasswordClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UserConfirmationPasswordClient, error) {
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
	client := &UserConfirmationPasswordClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Send - Sends confirmation
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// userID - User identifier. Must be unique in the current API Management service instance.
// options - UserConfirmationPasswordClientSendOptions contains the optional parameters for the UserConfirmationPasswordClient.Send
// method.
func (client *UserConfirmationPasswordClient) Send(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserConfirmationPasswordClientSendOptions) (UserConfirmationPasswordClientSendResponse, error) {
	req, err := client.sendCreateRequest(ctx, resourceGroupName, serviceName, userID, options)
	if err != nil {
		return UserConfirmationPasswordClientSendResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UserConfirmationPasswordClientSendResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return UserConfirmationPasswordClientSendResponse{}, runtime.NewResponseError(resp)
	}
	return UserConfirmationPasswordClientSendResponse{}, nil
}

// sendCreateRequest creates the Send request.
func (client *UserConfirmationPasswordClient) sendCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserConfirmationPasswordClientSendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/confirmations/password/send"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
