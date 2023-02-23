//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights

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
	"strconv"
	"strings"
)

// IncidentCommentsClient contains the methods for the IncidentComments group.
// Don't use this type directly, use NewIncidentCommentsClient() instead.
type IncidentCommentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIncidentCommentsClient creates a new instance of IncidentCommentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIncidentCommentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IncidentCommentsClient, error) {
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
	client := &IncidentCommentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the incident comment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// incidentCommentID - Incident comment ID
// incidentComment - The incident comment
// options - IncidentCommentsClientCreateOrUpdateOptions contains the optional parameters for the IncidentCommentsClient.CreateOrUpdate
// method.
func (client *IncidentCommentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, incidentComment IncidentComment, options *IncidentCommentsClientCreateOrUpdateOptions) (IncidentCommentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incidentCommentID, incidentComment, options)
	if err != nil {
		return IncidentCommentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentCommentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IncidentCommentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IncidentCommentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, incidentComment IncidentComment, options *IncidentCommentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/comments/{incidentCommentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if incidentCommentID == "" {
		return nil, errors.New("parameter incidentCommentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentCommentId}", url.PathEscape(incidentCommentID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, incidentComment)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IncidentCommentsClient) createOrUpdateHandleResponse(resp *http.Response) (IncidentCommentsClientCreateOrUpdateResponse, error) {
	result := IncidentCommentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentComment); err != nil {
		return IncidentCommentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the incident comment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// incidentCommentID - Incident comment ID
// options - IncidentCommentsClientDeleteOptions contains the optional parameters for the IncidentCommentsClient.Delete method.
func (client *IncidentCommentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, options *IncidentCommentsClientDeleteOptions) (IncidentCommentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incidentCommentID, options)
	if err != nil {
		return IncidentCommentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentCommentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IncidentCommentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IncidentCommentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IncidentCommentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, options *IncidentCommentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/comments/{incidentCommentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if incidentCommentID == "" {
		return nil, errors.New("parameter incidentCommentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentCommentId}", url.PathEscape(incidentCommentID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an incident comment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// incidentCommentID - Incident comment ID
// options - IncidentCommentsClientGetOptions contains the optional parameters for the IncidentCommentsClient.Get method.
func (client *IncidentCommentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, options *IncidentCommentsClientGetOptions) (IncidentCommentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incidentCommentID, options)
	if err != nil {
		return IncidentCommentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IncidentCommentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentCommentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IncidentCommentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, options *IncidentCommentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/comments/{incidentCommentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if incidentCommentID == "" {
		return nil, errors.New("parameter incidentCommentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentCommentId}", url.PathEscape(incidentCommentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IncidentCommentsClient) getHandleResponse(resp *http.Response) (IncidentCommentsClientGetResponse, error) {
	result := IncidentCommentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentComment); err != nil {
		return IncidentCommentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all incident comments.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// incidentID - Incident ID
// options - IncidentCommentsClientListOptions contains the optional parameters for the IncidentCommentsClient.List method.
func (client *IncidentCommentsClient) NewListPager(resourceGroupName string, workspaceName string, incidentID string, options *IncidentCommentsClientListOptions) *runtime.Pager[IncidentCommentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IncidentCommentsClientListResponse]{
		More: func(page IncidentCommentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IncidentCommentsClientListResponse) (IncidentCommentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IncidentCommentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IncidentCommentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IncidentCommentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IncidentCommentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentCommentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/comments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IncidentCommentsClient) listHandleResponse(resp *http.Response) (IncidentCommentsClientListResponse, error) {
	result := IncidentCommentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentCommentList); err != nil {
		return IncidentCommentsClientListResponse{}, err
	}
	return result, nil
}
