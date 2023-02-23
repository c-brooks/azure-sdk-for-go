//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// AnnotationsClient contains the methods for the Annotations group.
// Don't use this type directly, use NewAnnotationsClient() instead.
type AnnotationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAnnotationsClient creates a new instance of AnnotationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAnnotationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AnnotationsClient, error) {
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
	client := &AnnotationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Create an Annotation of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// annotationProperties - Properties that need to be specified to create an annotation of a Application Insights component.
// options - AnnotationsClientCreateOptions contains the optional parameters for the AnnotationsClient.Create method.
func (client *AnnotationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, annotationProperties Annotation, options *AnnotationsClientCreateOptions) (AnnotationsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, annotationProperties, options)
	if err != nil {
		return AnnotationsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AnnotationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AnnotationsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *AnnotationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, annotationProperties Annotation, options *AnnotationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, annotationProperties)
}

// createHandleResponse handles the Create response.
func (client *AnnotationsClient) createHandleResponse(resp *http.Response) (AnnotationsClientCreateResponse, error) {
	result := AnnotationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AnnotationArray); err != nil {
		return AnnotationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an Annotation of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// annotationID - The unique annotation ID. This is unique within a Application Insights component.
// options - AnnotationsClientDeleteOptions contains the optional parameters for the AnnotationsClient.Delete method.
func (client *AnnotationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, annotationID string, options *AnnotationsClientDeleteOptions) (AnnotationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, annotationID, options)
	if err != nil {
		return AnnotationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AnnotationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AnnotationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AnnotationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AnnotationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, annotationID string, options *AnnotationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if annotationID == "" {
		return nil, errors.New("parameter annotationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{annotationId}", url.PathEscape(annotationID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get the annotation for given id.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// annotationID - The unique annotation ID. This is unique within a Application Insights component.
// options - AnnotationsClientGetOptions contains the optional parameters for the AnnotationsClient.Get method.
func (client *AnnotationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, annotationID string, options *AnnotationsClientGetOptions) (AnnotationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, annotationID, options)
	if err != nil {
		return AnnotationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AnnotationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AnnotationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AnnotationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, annotationID string, options *AnnotationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if annotationID == "" {
		return nil, errors.New("parameter annotationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{annotationId}", url.PathEscape(annotationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AnnotationsClient) getHandleResponse(resp *http.Response) (AnnotationsClientGetResponse, error) {
	result := AnnotationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AnnotationArray); err != nil {
		return AnnotationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of annotations for a component for given time range
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// start - The start time to query from for annotations, cannot be older than 90 days from current date.
// end - The end time to query for annotations.
// options - AnnotationsClientListOptions contains the optional parameters for the AnnotationsClient.List method.
func (client *AnnotationsClient) NewListPager(resourceGroupName string, resourceName string, start string, end string, options *AnnotationsClientListOptions) *runtime.Pager[AnnotationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AnnotationsClientListResponse]{
		More: func(page AnnotationsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AnnotationsClientListResponse) (AnnotationsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, start, end, options)
			if err != nil {
				return AnnotationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AnnotationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AnnotationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AnnotationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, start string, end string, options *AnnotationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	reqQP.Set("start", start)
	reqQP.Set("end", end)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AnnotationsClient) listHandleResponse(resp *http.Response) (AnnotationsClientListResponse, error) {
	result := AnnotationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AnnotationsListResult); err != nil {
		return AnnotationsClientListResponse{}, err
	}
	return result, nil
}
