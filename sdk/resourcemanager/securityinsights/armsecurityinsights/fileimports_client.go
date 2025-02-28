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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FileImportsClient contains the methods for the FileImports group.
// Don't use this type directly, use NewFileImportsClient() instead.
type FileImportsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFileImportsClient creates a new instance of FileImportsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFileImportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FileImportsClient, error) {
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
	client := &FileImportsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates the file import.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// fileImportID - File import ID
// fileImport - The file import
// options - FileImportsClientCreateOptions contains the optional parameters for the FileImportsClient.Create method.
func (client *FileImportsClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, fileImport FileImport, options *FileImportsClientCreateOptions) (FileImportsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, fileImportID, fileImport, options)
	if err != nil {
		return FileImportsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileImportsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return FileImportsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *FileImportsClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, fileImport FileImport, options *FileImportsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/fileImports/{fileImportId}"
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
	if fileImportID == "" {
		return nil, errors.New("parameter fileImportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileImportId}", url.PathEscape(fileImportID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, fileImport)
}

// createHandleResponse handles the Create response.
func (client *FileImportsClient) createHandleResponse(resp *http.Response) (FileImportsClientCreateResponse, error) {
	result := FileImportsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileImport); err != nil {
		return FileImportsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete the file import.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// fileImportID - File import ID
// options - FileImportsClientBeginDeleteOptions contains the optional parameters for the FileImportsClient.BeginDelete method.
func (client *FileImportsClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, options *FileImportsClientBeginDeleteOptions) (*runtime.Poller[FileImportsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, fileImportID, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[FileImportsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FileImportsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete the file import.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *FileImportsClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, options *FileImportsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, fileImportID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FileImportsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, options *FileImportsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/fileImports/{fileImportId}"
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
	if fileImportID == "" {
		return nil, errors.New("parameter fileImportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileImportId}", url.PathEscape(fileImportID))
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

// Get - Gets a file import.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// fileImportID - File import ID
// options - FileImportsClientGetOptions contains the optional parameters for the FileImportsClient.Get method.
func (client *FileImportsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, options *FileImportsClientGetOptions) (FileImportsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, fileImportID, options)
	if err != nil {
		return FileImportsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FileImportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FileImportsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FileImportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, fileImportID string, options *FileImportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/fileImports/{fileImportId}"
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
	if fileImportID == "" {
		return nil, errors.New("parameter fileImportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fileImportId}", url.PathEscape(fileImportID))
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
func (client *FileImportsClient) getHandleResponse(resp *http.Response) (FileImportsClientGetResponse, error) {
	result := FileImportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileImport); err != nil {
		return FileImportsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all file imports.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - FileImportsClientListOptions contains the optional parameters for the FileImportsClient.List method.
func (client *FileImportsClient) NewListPager(resourceGroupName string, workspaceName string, options *FileImportsClientListOptions) *runtime.Pager[FileImportsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FileImportsClientListResponse]{
		More: func(page FileImportsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FileImportsClientListResponse) (FileImportsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FileImportsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FileImportsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FileImportsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FileImportsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *FileImportsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/fileImports"
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
func (client *FileImportsClient) listHandleResponse(resp *http.Response) (FileImportsClientListResponse, error) {
	result := FileImportsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileImportList); err != nil {
		return FileImportsClientListResponse{}, err
	}
	return result, nil
}
