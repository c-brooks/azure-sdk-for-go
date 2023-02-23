//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armavs

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

// CloudLinksClient contains the methods for the CloudLinks group.
// Don't use this type directly, use NewCloudLinksClient() instead.
type CloudLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCloudLinksClient creates a new instance of CloudLinksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCloudLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudLinksClient, error) {
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
	client := &CloudLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a cloud link in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - The name of the private cloud.
// cloudLinkName - Name of the cloud link resource
// cloudLink - A cloud link in the private cloud
// options - CloudLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the CloudLinksClient.BeginCreateOrUpdate
// method.
func (client *CloudLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, cloudLink CloudLink, options *CloudLinksClientBeginCreateOrUpdateOptions) (*runtime.Poller[CloudLinksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, cloudLinkName, cloudLink, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CloudLinksClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CloudLinksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a cloud link in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *CloudLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, cloudLink CloudLink, options *CloudLinksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, cloudLinkName, cloudLink, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CloudLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, cloudLink CloudLink, options *CloudLinksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/cloudLinks/{cloudLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if cloudLinkName == "" {
		return nil, errors.New("parameter cloudLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudLinkName}", url.PathEscape(cloudLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, cloudLink)
}

// BeginDelete - Delete a cloud link in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// cloudLinkName - Name of the cloud link resource
// options - CloudLinksClientBeginDeleteOptions contains the optional parameters for the CloudLinksClient.BeginDelete method.
func (client *CloudLinksClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, options *CloudLinksClientBeginDeleteOptions) (*runtime.Poller[CloudLinksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, cloudLinkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CloudLinksClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[CloudLinksClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a cloud link in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *CloudLinksClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, options *CloudLinksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, cloudLinkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CloudLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, options *CloudLinksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/cloudLinks/{cloudLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if cloudLinkName == "" {
		return nil, errors.New("parameter cloudLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudLinkName}", url.PathEscape(cloudLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an cloud link by name in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// cloudLinkName - Name of the cloud link resource
// options - CloudLinksClientGetOptions contains the optional parameters for the CloudLinksClient.Get method.
func (client *CloudLinksClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, options *CloudLinksClientGetOptions) (CloudLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, cloudLinkName, options)
	if err != nil {
		return CloudLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CloudLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, cloudLinkName string, options *CloudLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/cloudLinks/{cloudLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if cloudLinkName == "" {
		return nil, errors.New("parameter cloudLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudLinkName}", url.PathEscape(cloudLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudLinksClient) getHandleResponse(resp *http.Response) (CloudLinksClientGetResponse, error) {
	result := CloudLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudLink); err != nil {
		return CloudLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List cloud link in a private cloud
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// privateCloudName - Name of the private cloud
// options - CloudLinksClientListOptions contains the optional parameters for the CloudLinksClient.List method.
func (client *CloudLinksClient) NewListPager(resourceGroupName string, privateCloudName string, options *CloudLinksClientListOptions) *runtime.Pager[CloudLinksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudLinksClientListResponse]{
		More: func(page CloudLinksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudLinksClientListResponse) (CloudLinksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CloudLinksClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CloudLinksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CloudLinksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CloudLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *CloudLinksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/cloudLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CloudLinksClient) listHandleResponse(resp *http.Response) (CloudLinksClientListResponse, error) {
	result := CloudLinksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudLinkList); err != nil {
		return CloudLinksClientListResponse{}, err
	}
	return result, nil
}
