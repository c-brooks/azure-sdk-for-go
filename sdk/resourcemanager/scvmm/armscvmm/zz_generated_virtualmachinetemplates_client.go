//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

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

// VirtualMachineTemplatesClient contains the methods for the VirtualMachineTemplates group.
// Don't use this type directly, use NewVirtualMachineTemplatesClient() instead.
type VirtualMachineTemplatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineTemplatesClient creates a new instance of VirtualMachineTemplatesClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineTemplatesClient, error) {
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
	client := &VirtualMachineTemplatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Onboards the ScVmm VM Template as an Azure VM Template resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// virtualMachineTemplateName - Name of the VirtualMachineTemplate.
// body - Request payload.
// options - VirtualMachineTemplatesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineTemplatesClient.BeginCreateOrUpdate
// method.
func (client *VirtualMachineTemplatesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body VirtualMachineTemplate, options *VirtualMachineTemplatesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineTemplatesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachineTemplatesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineTemplatesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Onboards the ScVmm VM Template as an Azure VM Template resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
func (client *VirtualMachineTemplatesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body VirtualMachineTemplate, options *VirtualMachineTemplatesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
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
func (client *VirtualMachineTemplatesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body VirtualMachineTemplate, options *VirtualMachineTemplatesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deregisters the ScVmm VM Template from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// virtualMachineTemplateName - Name of the VirtualMachineTemplate.
// options - VirtualMachineTemplatesClientBeginDeleteOptions contains the optional parameters for the VirtualMachineTemplatesClient.BeginDelete
// method.
func (client *VirtualMachineTemplatesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineTemplatesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualMachineTemplateName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachineTemplatesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineTemplatesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deregisters the ScVmm VM Template from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
func (client *VirtualMachineTemplatesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, options)
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
func (client *VirtualMachineTemplatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements VirtualMachineTemplate GET method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// virtualMachineTemplateName - Name of the VirtualMachineTemplate.
// options - VirtualMachineTemplatesClientGetOptions contains the optional parameters for the VirtualMachineTemplatesClient.Get
// method.
func (client *VirtualMachineTemplatesClient) Get(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientGetOptions) (VirtualMachineTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, options)
	if err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, options *VirtualMachineTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachineTemplates/{virtualMachineTemplateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineTemplatesClient) getHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientGetResponse, error) {
	result := VirtualMachineTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplate); err != nil {
		return VirtualMachineTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of VirtualMachineTemplates in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// options - VirtualMachineTemplatesClientListByResourceGroupOptions contains the optional parameters for the VirtualMachineTemplatesClient.ListByResourceGroup
// method.
func (client *VirtualMachineTemplatesClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualMachineTemplatesClientListByResourceGroupOptions) *runtime.Pager[VirtualMachineTemplatesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineTemplatesClientListByResourceGroupResponse]{
		More: func(page VirtualMachineTemplatesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineTemplatesClientListByResourceGroupResponse) (VirtualMachineTemplatesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineTemplatesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualMachineTemplatesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualMachineTemplatesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachineTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualMachineTemplatesClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientListByResourceGroupResponse, error) {
	result := VirtualMachineTemplatesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplateListResult); err != nil {
		return VirtualMachineTemplatesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List of VirtualMachineTemplates in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// options - VirtualMachineTemplatesClientListBySubscriptionOptions contains the optional parameters for the VirtualMachineTemplatesClient.ListBySubscription
// method.
func (client *VirtualMachineTemplatesClient) NewListBySubscriptionPager(options *VirtualMachineTemplatesClientListBySubscriptionOptions) *runtime.Pager[VirtualMachineTemplatesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineTemplatesClientListBySubscriptionResponse]{
		More: func(page VirtualMachineTemplatesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineTemplatesClientListBySubscriptionResponse) (VirtualMachineTemplatesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineTemplatesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineTemplatesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineTemplatesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VirtualMachineTemplatesClient) listBySubscriptionCreateRequest(ctx context.Context, options *VirtualMachineTemplatesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ScVmm/virtualMachineTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VirtualMachineTemplatesClient) listBySubscriptionHandleResponse(resp *http.Response) (VirtualMachineTemplatesClientListBySubscriptionResponse, error) {
	result := VirtualMachineTemplatesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineTemplateListResult); err != nil {
		return VirtualMachineTemplatesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the VirtualMachineTemplate resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
// resourceGroupName - The name of the resource group.
// virtualMachineTemplateName - Name of the VirtualMachineTemplate.
// body - VirtualMachineTemplates patch details.
// options - VirtualMachineTemplatesClientBeginUpdateOptions contains the optional parameters for the VirtualMachineTemplatesClient.BeginUpdate
// method.
func (client *VirtualMachineTemplatesClient) BeginUpdate(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body ResourcePatch, options *VirtualMachineTemplatesClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineTemplatesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VirtualMachineTemplatesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineTemplatesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates the VirtualMachineTemplate resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-06-05-preview
func (client *VirtualMachineTemplatesClient) update(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body ResourcePatch, options *VirtualMachineTemplatesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualMachineTemplateName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineTemplatesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineTemplateName string, body ResourcePatch, options *VirtualMachineTemplatesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualMachineTemplates/{virtualMachineTemplateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualMachineTemplateName == "" {
		return nil, errors.New("parameter virtualMachineTemplateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineTemplateName}", url.PathEscape(virtualMachineTemplateName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-05-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
