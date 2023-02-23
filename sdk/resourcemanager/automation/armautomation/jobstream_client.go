//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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

// JobStreamClient contains the methods for the JobStream group.
// Don't use this type directly, use NewJobStreamClient() instead.
type JobStreamClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobStreamClient creates a new instance of JobStreamClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobStreamClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobStreamClient, error) {
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
	client := &JobStreamClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Retrieve the job stream identified by job stream id.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// jobStreamID - The job stream id.
// options - JobStreamClientGetOptions contains the optional parameters for the JobStreamClient.Get method.
func (client *JobStreamClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, jobStreamID string, options *JobStreamClientGetOptions) (JobStreamClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, jobStreamID, options)
	if err != nil {
		return JobStreamClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobStreamClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobStreamClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobStreamClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, jobStreamID string, options *JobStreamClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/streams/{jobStreamId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if jobStreamID == "" {
		return nil, errors.New("parameter jobStreamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobStreamId}", url.PathEscape(jobStreamID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobStreamClient) getHandleResponse(resp *http.Response) (JobStreamClientGetResponse, error) {
	result := JobStreamClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStream); err != nil {
		return JobStreamClientGetResponse{}, err
	}
	return result, nil
}

// NewListByJobPager - Retrieve a list of jobs streams identified by job name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// options - JobStreamClientListByJobOptions contains the optional parameters for the JobStreamClient.ListByJob method.
func (client *JobStreamClient) NewListByJobPager(resourceGroupName string, automationAccountName string, jobName string, options *JobStreamClientListByJobOptions) *runtime.Pager[JobStreamClientListByJobResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobStreamClientListByJobResponse]{
		More: func(page JobStreamClientListByJobResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobStreamClientListByJobResponse) (JobStreamClientListByJobResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByJobCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobStreamClientListByJobResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobStreamClientListByJobResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobStreamClientListByJobResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByJobHandleResponse(resp)
		},
	})
}

// listByJobCreateRequest creates the ListByJob request.
func (client *JobStreamClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobStreamClientListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/streams"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["clientRequestId"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *JobStreamClient) listByJobHandleResponse(resp *http.Response) (JobStreamClientListByJobResponse, error) {
	result := JobStreamClientListByJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStreamListResult); err != nil {
		return JobStreamClientListByJobResponse{}, err
	}
	return result, nil
}
