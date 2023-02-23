//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

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

// ProjectEnvironmentTypesClient contains the methods for the ProjectEnvironmentTypes group.
// Don't use this type directly, use NewProjectEnvironmentTypesClient() instead.
type ProjectEnvironmentTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProjectEnvironmentTypesClient creates a new instance of ProjectEnvironmentTypesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProjectEnvironmentTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectEnvironmentTypesClient, error) {
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
	client := &ProjectEnvironmentTypesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a project environment type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// environmentTypeName - The name of the environment type.
// body - Represents a Project Environment Type.
// options - ProjectEnvironmentTypesClientCreateOrUpdateOptions contains the optional parameters for the ProjectEnvironmentTypesClient.CreateOrUpdate
// method.
func (client *ProjectEnvironmentTypesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, body ProjectEnvironmentType, options *ProjectEnvironmentTypesClientCreateOrUpdateOptions) (ProjectEnvironmentTypesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, projectName, environmentTypeName, body, options)
	if err != nil {
		return ProjectEnvironmentTypesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProjectEnvironmentTypesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProjectEnvironmentTypesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProjectEnvironmentTypesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, body ProjectEnvironmentType, options *ProjectEnvironmentTypesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProjectEnvironmentTypesClient) createOrUpdateHandleResponse(resp *http.Response) (ProjectEnvironmentTypesClientCreateOrUpdateResponse, error) {
	result := ProjectEnvironmentTypesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectEnvironmentType); err != nil {
		return ProjectEnvironmentTypesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a project environment type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// environmentTypeName - The name of the environment type.
// options - ProjectEnvironmentTypesClientDeleteOptions contains the optional parameters for the ProjectEnvironmentTypesClient.Delete
// method.
func (client *ProjectEnvironmentTypesClient) Delete(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *ProjectEnvironmentTypesClientDeleteOptions) (ProjectEnvironmentTypesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, environmentTypeName, options)
	if err != nil {
		return ProjectEnvironmentTypesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProjectEnvironmentTypesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProjectEnvironmentTypesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProjectEnvironmentTypesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProjectEnvironmentTypesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *ProjectEnvironmentTypesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a project environment type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// environmentTypeName - The name of the environment type.
// options - ProjectEnvironmentTypesClientGetOptions contains the optional parameters for the ProjectEnvironmentTypesClient.Get
// method.
func (client *ProjectEnvironmentTypesClient) Get(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *ProjectEnvironmentTypesClientGetOptions) (ProjectEnvironmentTypesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, environmentTypeName, options)
	if err != nil {
		return ProjectEnvironmentTypesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProjectEnvironmentTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProjectEnvironmentTypesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProjectEnvironmentTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, options *ProjectEnvironmentTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectEnvironmentTypesClient) getHandleResponse(resp *http.Response) (ProjectEnvironmentTypesClientGetResponse, error) {
	result := ProjectEnvironmentTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectEnvironmentType); err != nil {
		return ProjectEnvironmentTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists environment types for a project.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// options - ProjectEnvironmentTypesClientListOptions contains the optional parameters for the ProjectEnvironmentTypesClient.List
// method.
func (client *ProjectEnvironmentTypesClient) NewListPager(resourceGroupName string, projectName string, options *ProjectEnvironmentTypesClientListOptions) *runtime.Pager[ProjectEnvironmentTypesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectEnvironmentTypesClientListResponse]{
		More: func(page ProjectEnvironmentTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectEnvironmentTypesClientListResponse) (ProjectEnvironmentTypesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProjectEnvironmentTypesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProjectEnvironmentTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProjectEnvironmentTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProjectEnvironmentTypesClient) listCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *ProjectEnvironmentTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/environmentTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProjectEnvironmentTypesClient) listHandleResponse(resp *http.Response) (ProjectEnvironmentTypesClientListResponse, error) {
	result := ProjectEnvironmentTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectEnvironmentTypeListResult); err != nil {
		return ProjectEnvironmentTypesClientListResponse{}, err
	}
	return result, nil
}

// Update - Partially updates a project environment type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-11-11-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// projectName - The name of the project.
// environmentTypeName - The name of the environment type.
// body - Updatable project environment type properties.
// options - ProjectEnvironmentTypesClientUpdateOptions contains the optional parameters for the ProjectEnvironmentTypesClient.Update
// method.
func (client *ProjectEnvironmentTypesClient) Update(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, body ProjectEnvironmentTypeUpdate, options *ProjectEnvironmentTypesClientUpdateOptions) (ProjectEnvironmentTypesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, projectName, environmentTypeName, body, options)
	if err != nil {
		return ProjectEnvironmentTypesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProjectEnvironmentTypesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProjectEnvironmentTypesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProjectEnvironmentTypesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, environmentTypeName string, body ProjectEnvironmentTypeUpdate, options *ProjectEnvironmentTypesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/environmentTypes/{environmentTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if environmentTypeName == "" {
		return nil, errors.New("parameter environmentTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentTypeName}", url.PathEscape(environmentTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *ProjectEnvironmentTypesClient) updateHandleResponse(resp *http.Response) (ProjectEnvironmentTypesClientUpdateResponse, error) {
	result := ProjectEnvironmentTypesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectEnvironmentType); err != nil {
		return ProjectEnvironmentTypesClientUpdateResponse{}, err
	}
	return result, nil
}
