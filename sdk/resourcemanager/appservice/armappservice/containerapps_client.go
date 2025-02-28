//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice

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
	"strings"
)

// ContainerAppsClient contains the methods for the ContainerApps group.
// Don't use this type directly, use NewContainerAppsClient() instead.
type ContainerAppsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContainerAppsClient creates a new instance of ContainerAppsClient with the specified values.
// subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContainerAppsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerAppsClient, error) {
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
	client := &ContainerAppsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Description for Create or update a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of the Container App.
// options - ContainerAppsClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerAppsClient.BeginCreateOrUpdate
// method.
func (client *ContainerAppsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainerAppsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, name, containerAppEnvelope, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerAppsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Description for Create or update a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
func (client *ContainerAppsClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, containerAppEnvelope, options)
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
func (client *ContainerAppsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, containerAppEnvelope ContainerApp, options *ContainerAppsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, containerAppEnvelope)
}

// BeginDelete - Description for Delete a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of the Container App.
// options - ContainerAppsClientBeginDeleteOptions contains the optional parameters for the ContainerAppsClient.BeginDelete
// method.
func (client *ContainerAppsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *ContainerAppsClientBeginDeleteOptions) (*runtime.Poller[ContainerAppsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ContainerAppsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ContainerAppsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Description for Delete a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
func (client *ContainerAppsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *ContainerAppsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *ContainerAppsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ContainerAppsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a Container App.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// name - Name of the Container App.
// options - ContainerAppsClientGetOptions contains the optional parameters for the ContainerAppsClient.Get method.
func (client *ContainerAppsClient) Get(ctx context.Context, resourceGroupName string, name string, options *ContainerAppsClientGetOptions) (ContainerAppsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return ContainerAppsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerAppsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerAppsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContainerAppsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *ContainerAppsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerAppsClient) getHandleResponse(resp *http.Response) (ContainerAppsClientGetResponse, error) {
	result := ContainerAppsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerApp); err != nil {
		return ContainerAppsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get the Container Apps in a given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// resourceGroupName - Name of the resource group to which the resource belongs.
// options - ContainerAppsClientListByResourceGroupOptions contains the optional parameters for the ContainerAppsClient.ListByResourceGroup
// method.
func (client *ContainerAppsClient) NewListByResourceGroupPager(resourceGroupName string, options *ContainerAppsClientListByResourceGroupOptions) *runtime.Pager[ContainerAppsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsClientListByResourceGroupResponse]{
		More: func(page ContainerAppsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsClientListByResourceGroupResponse) (ContainerAppsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerAppsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainerAppsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerAppsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerAppsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerAppsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/containerApps"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerAppsClient) listByResourceGroupHandleResponse(resp *http.Response) (ContainerAppsClientListByResourceGroupResponse, error) {
	result := ContainerAppsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppCollection); err != nil {
		return ContainerAppsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get the Container Apps in a given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// options - ContainerAppsClientListBySubscriptionOptions contains the optional parameters for the ContainerAppsClient.ListBySubscription
// method.
func (client *ContainerAppsClient) NewListBySubscriptionPager(options *ContainerAppsClientListBySubscriptionOptions) *runtime.Pager[ContainerAppsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerAppsClientListBySubscriptionResponse]{
		More: func(page ContainerAppsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerAppsClientListBySubscriptionResponse) (ContainerAppsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ContainerAppsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContainerAppsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContainerAppsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ContainerAppsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ContainerAppsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/containerApps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ContainerAppsClient) listBySubscriptionHandleResponse(resp *http.Response) (ContainerAppsClientListBySubscriptionResponse, error) {
	result := ContainerAppsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerAppCollection); err != nil {
		return ContainerAppsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListSecrets - List secrets for a container app
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01
// name - Name of the Container App.
// options - ContainerAppsClientListSecretsOptions contains the optional parameters for the ContainerAppsClient.ListSecrets
// method.
func (client *ContainerAppsClient) ListSecrets(ctx context.Context, name string, options *ContainerAppsClientListSecretsOptions) (ContainerAppsClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, name, options)
	if err != nil {
		return ContainerAppsClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContainerAppsClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContainerAppsClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *ContainerAppsClient) listSecretsCreateRequest(ctx context.Context, name string, options *ContainerAppsClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/containerApps/{name}/listSecrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *ContainerAppsClient) listSecretsHandleResponse(resp *http.Response) (ContainerAppsClientListSecretsResponse, error) {
	result := ContainerAppsClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecretsCollection); err != nil {
		return ContainerAppsClientListSecretsResponse{}, err
	}
	return result, nil
}
