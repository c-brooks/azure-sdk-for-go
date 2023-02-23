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

// ContentTypeClient contains the methods for the ContentType group.
// Don't use this type directly, use NewContentTypeClient() instead.
type ContentTypeClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContentTypeClient creates a new instance of ContentTypeClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContentTypeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContentTypeClient, error) {
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
	client := &ContentTypeClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the developer portal's content type. Content types describe content items' properties,
// validation rules, and constraints. Custom content types' identifiers need to start with the c-
// prefix. Built-in content types can't be modified.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// options - ContentTypeClientCreateOrUpdateOptions contains the optional parameters for the ContentTypeClient.CreateOrUpdate
// method.
func (client *ContentTypeClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentTypeClientCreateOrUpdateOptions) (ContentTypeClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, options)
	if err != nil {
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ContentTypeClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContentTypeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentTypeClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContentTypeClient) createOrUpdateHandleResponse(resp *http.Response) (ContentTypeClientCreateOrUpdateResponse, error) {
	result := ContentTypeClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentTypeContract); err != nil {
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the specified developer portal's content type. Content types describe content items' properties, validation
// rules, and constraints. Built-in content types (with identifiers starting with the
// c- prefix) can't be removed.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - ContentTypeClientDeleteOptions contains the optional parameters for the ContentTypeClient.Delete method.
func (client *ContentTypeClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, ifMatch string, options *ContentTypeClientDeleteOptions) (ContentTypeClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, ifMatch, options)
	if err != nil {
		return ContentTypeClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentTypeClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ContentTypeClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ContentTypeClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContentTypeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, ifMatch string, options *ContentTypeClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the developer portal's content type. Content types describe content items' properties, validation
// rules, and constraints.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// contentTypeID - Content type identifier.
// options - ContentTypeClientGetOptions contains the optional parameters for the ContentTypeClient.Get method.
func (client *ContentTypeClient) Get(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentTypeClientGetOptions) (ContentTypeClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, options)
	if err != nil {
		return ContentTypeClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentTypeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContentTypeClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContentTypeClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentTypeClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContentTypeClient) getHandleResponse(resp *http.Response) (ContentTypeClientGetResponse, error) {
	result := ContentTypeClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentTypeContract); err != nil {
		return ContentTypeClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Lists the developer portal's content types. Content types describe content items' properties, validation
// rules, and constraints.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - ContentTypeClientListByServiceOptions contains the optional parameters for the ContentTypeClient.ListByService
// method.
func (client *ContentTypeClient) NewListByServicePager(resourceGroupName string, serviceName string, options *ContentTypeClientListByServiceOptions) *runtime.Pager[ContentTypeClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContentTypeClientListByServiceResponse]{
		More: func(page ContentTypeClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContentTypeClientListByServiceResponse) (ContentTypeClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContentTypeClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContentTypeClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContentTypeClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *ContentTypeClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ContentTypeClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *ContentTypeClient) listByServiceHandleResponse(resp *http.Response) (ContentTypeClientListByServiceResponse, error) {
	result := ContentTypeClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentTypeCollection); err != nil {
		return ContentTypeClientListByServiceResponse{}, err
	}
	return result, nil
}
