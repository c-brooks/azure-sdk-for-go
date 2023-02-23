//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeducation

import (
	"context"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/c-brooks/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ManagementClient contains the methods for the EducationManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	host string
	pl   runtime.Pipeline
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagementClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
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
	client := &ManagementClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// RedeemInvitationCode - Redeem invite code to join a redeemable lab
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// parameters - Request parameters to provide redeem code.
// options - ManagementClientRedeemInvitationCodeOptions contains the optional parameters for the ManagementClient.RedeemInvitationCode
// method.
func (client *ManagementClient) RedeemInvitationCode(ctx context.Context, parameters RedeemRequest, options *ManagementClientRedeemInvitationCodeOptions) (ManagementClientRedeemInvitationCodeResponse, error) {
	req, err := client.redeemInvitationCodeCreateRequest(ctx, parameters, options)
	if err != nil {
		return ManagementClientRedeemInvitationCodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementClientRedeemInvitationCodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientRedeemInvitationCodeResponse{}, runtime.NewResponseError(resp)
	}
	return ManagementClientRedeemInvitationCodeResponse{}, nil
}

// redeemInvitationCodeCreateRequest creates the RedeemInvitationCode request.
func (client *ManagementClient) redeemInvitationCodeCreateRequest(ctx context.Context, parameters RedeemRequest, options *ManagementClientRedeemInvitationCodeOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Education/redeemInvitationCode"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
