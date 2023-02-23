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

// JobClient contains the methods for the Job group.
// Don't use this type directly, use NewJobClient() instead.
type JobClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobClient creates a new instance of JobClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobClient, error) {
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
	client := &JobClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create a job of the runbook.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// parameters - The parameters supplied to the create job operation.
// options - JobClientCreateOptions contains the optional parameters for the JobClient.Create method.
func (client *JobClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, parameters JobCreateParameters, options *JobClientCreateOptions) (JobClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, parameters, options)
	if err != nil {
		return JobClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return JobClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *JobClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, parameters JobCreateParameters, options *JobClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
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
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *JobClient) createHandleResponse(resp *http.Response) (JobClientCreateResponse, error) {
	result := JobClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return JobClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Retrieve the job identified by job name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// options - JobClientGetOptions contains the optional parameters for the JobClient.Get method.
func (client *JobClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientGetOptions) (JobClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
	if err != nil {
		return JobClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}"
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
func (client *JobClient) getHandleResponse(resp *http.Response) (JobClientGetResponse, error) {
	result := JobClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return JobClientGetResponse{}, err
	}
	return result, nil
}

// GetOutput - Retrieve the job output identified by job name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The name of the job to be created.
// options - JobClientGetOutputOptions contains the optional parameters for the JobClient.GetOutput method.
func (client *JobClient) GetOutput(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientGetOutputOptions) (JobClientGetOutputResponse, error) {
	req, err := client.getOutputCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
	if err != nil {
		return JobClientGetOutputResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientGetOutputResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobClientGetOutputResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOutputHandleResponse(resp)
}

// getOutputCreateRequest creates the GetOutput request.
func (client *JobClient) getOutputCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientGetOutputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/output"
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
	req.Raw().Header["Accept"] = []string{"text/plain"}
	return req, nil
}

// getOutputHandleResponse handles the GetOutput response.
func (client *JobClient) getOutputHandleResponse(resp *http.Response) (JobClientGetOutputResponse, error) {
	result := JobClientGetOutputResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return JobClientGetOutputResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// GetRunbookContent - Retrieve the runbook content of the job identified by job name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// options - JobClientGetRunbookContentOptions contains the optional parameters for the JobClient.GetRunbookContent method.
func (client *JobClient) GetRunbookContent(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientGetRunbookContentOptions) (JobClientGetRunbookContentResponse, error) {
	req, err := client.getRunbookContentCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
	if err != nil {
		return JobClientGetRunbookContentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientGetRunbookContentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobClientGetRunbookContentResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRunbookContentHandleResponse(resp)
}

// getRunbookContentCreateRequest creates the GetRunbookContent request.
func (client *JobClient) getRunbookContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientGetRunbookContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/runbookContent"
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
	req.Raw().Header["Accept"] = []string{"text/powershell"}
	return req, nil
}

// getRunbookContentHandleResponse handles the GetRunbookContent response.
func (client *JobClient) getRunbookContentHandleResponse(resp *http.Response) (JobClientGetRunbookContentResponse, error) {
	result := JobClientGetRunbookContentResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return JobClientGetRunbookContentResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of jobs.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - JobClientListByAutomationAccountOptions contains the optional parameters for the JobClient.ListByAutomationAccount
// method.
func (client *JobClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *JobClientListByAutomationAccountOptions) *runtime.Pager[JobClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobClientListByAutomationAccountResponse]{
		More: func(page JobClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobClientListByAutomationAccountResponse) (JobClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *JobClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *JobClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
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

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *JobClient) listByAutomationAccountHandleResponse(resp *http.Response) (JobClientListByAutomationAccountResponse, error) {
	result := JobClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobListResultV2); err != nil {
		return JobClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Resume - Resume the job identified by jobName.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// options - JobClientResumeOptions contains the optional parameters for the JobClient.Resume method.
func (client *JobClient) Resume(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientResumeOptions) (JobClientResumeResponse, error) {
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
	if err != nil {
		return JobClientResumeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientResumeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobClientResumeResponse{}, runtime.NewResponseError(resp)
	}
	return JobClientResumeResponse{}, nil
}

// resumeCreateRequest creates the Resume request.
func (client *JobClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/resume"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
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

// Stop - Stop the job identified by jobName.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// options - JobClientStopOptions contains the optional parameters for the JobClient.Stop method.
func (client *JobClient) Stop(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientStopOptions) (JobClientStopResponse, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
	if err != nil {
		return JobClientStopResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobClientStopResponse{}, runtime.NewResponseError(resp)
	}
	return JobClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *JobClient) stopCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/stop"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
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

// Suspend - Suspend the job identified by job name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// jobName - The job name.
// options - JobClientSuspendOptions contains the optional parameters for the JobClient.Suspend method.
func (client *JobClient) Suspend(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientSuspendOptions) (JobClientSuspendResponse, error) {
	req, err := client.suspendCreateRequest(ctx, resourceGroupName, automationAccountName, jobName, options)
	if err != nil {
		return JobClientSuspendResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobClientSuspendResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobClientSuspendResponse{}, runtime.NewResponseError(resp)
	}
	return JobClientSuspendResponse{}, nil
}

// suspendCreateRequest creates the Suspend request.
func (client *JobClient) suspendCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *JobClientSuspendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobName}/suspend"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
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
