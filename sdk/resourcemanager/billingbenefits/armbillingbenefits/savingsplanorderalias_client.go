//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbillingbenefits

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

// SavingsPlanOrderAliasClient contains the methods for the SavingsPlanOrderAlias group.
// Don't use this type directly, use NewSavingsPlanOrderAliasClient() instead.
type SavingsPlanOrderAliasClient struct {
	host string
	pl   runtime.Pipeline
}

// NewSavingsPlanOrderAliasClient creates a new instance of SavingsPlanOrderAliasClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSavingsPlanOrderAliasClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SavingsPlanOrderAliasClient, error) {
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
	client := &SavingsPlanOrderAliasClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreate - Create a savings plan. Learn more about permissions needed at https://go.microsoft.com/fwlink/?linkid=2215851
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// savingsPlanOrderAliasName - Name of the savings plan order alias
// body - Request body for creating a savings plan order alias
// options - SavingsPlanOrderAliasClientBeginCreateOptions contains the optional parameters for the SavingsPlanOrderAliasClient.BeginCreate
// method.
func (client *SavingsPlanOrderAliasClient) BeginCreate(ctx context.Context, savingsPlanOrderAliasName string, body SavingsPlanOrderAliasModel, options *SavingsPlanOrderAliasClientBeginCreateOptions) (*runtime.Poller[SavingsPlanOrderAliasClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, savingsPlanOrderAliasName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SavingsPlanOrderAliasClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SavingsPlanOrderAliasClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a savings plan. Learn more about permissions needed at https://go.microsoft.com/fwlink/?linkid=2215851
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
func (client *SavingsPlanOrderAliasClient) create(ctx context.Context, savingsPlanOrderAliasName string, body SavingsPlanOrderAliasModel, options *SavingsPlanOrderAliasClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, savingsPlanOrderAliasName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *SavingsPlanOrderAliasClient) createCreateRequest(ctx context.Context, savingsPlanOrderAliasName string, body SavingsPlanOrderAliasModel, options *SavingsPlanOrderAliasClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrderAliases/{savingsPlanOrderAliasName}"
	if savingsPlanOrderAliasName == "" {
		return nil, errors.New("parameter savingsPlanOrderAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderAliasName}", url.PathEscape(savingsPlanOrderAliasName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Get a savings plan.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-01
// savingsPlanOrderAliasName - Name of the savings plan order alias
// options - SavingsPlanOrderAliasClientGetOptions contains the optional parameters for the SavingsPlanOrderAliasClient.Get
// method.
func (client *SavingsPlanOrderAliasClient) Get(ctx context.Context, savingsPlanOrderAliasName string, options *SavingsPlanOrderAliasClientGetOptions) (SavingsPlanOrderAliasClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, savingsPlanOrderAliasName, options)
	if err != nil {
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SavingsPlanOrderAliasClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SavingsPlanOrderAliasClient) getCreateRequest(ctx context.Context, savingsPlanOrderAliasName string, options *SavingsPlanOrderAliasClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.BillingBenefits/savingsPlanOrderAliases/{savingsPlanOrderAliasName}"
	if savingsPlanOrderAliasName == "" {
		return nil, errors.New("parameter savingsPlanOrderAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{savingsPlanOrderAliasName}", url.PathEscape(savingsPlanOrderAliasName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SavingsPlanOrderAliasClient) getHandleResponse(resp *http.Response) (SavingsPlanOrderAliasClientGetResponse, error) {
	result := SavingsPlanOrderAliasClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SavingsPlanOrderAliasModel); err != nil {
		return SavingsPlanOrderAliasClientGetResponse{}, err
	}
	return result, nil
}
