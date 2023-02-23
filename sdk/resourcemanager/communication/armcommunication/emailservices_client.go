//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcommunication

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

// EmailServicesClient contains the methods for the EmailServices group.
// Don't use this type directly, use NewEmailServicesClient() instead.
type EmailServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEmailServicesClient creates a new instance of EmailServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEmailServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EmailServicesClient, error) {
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
	client := &EmailServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new EmailService or update an existing EmailService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// parameters - Parameters for the create or update operation
// options - EmailServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the EmailServicesClient.BeginCreateOrUpdate
// method.
func (client *EmailServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, parameters EmailServiceResource, options *EmailServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[EmailServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, emailServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EmailServicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EmailServicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create a new EmailService or update an existing EmailService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *EmailServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, parameters EmailServiceResource, options *EmailServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, emailServiceName, parameters, options)
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
func (client *EmailServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, parameters EmailServiceResource, options *EmailServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Operation to delete a EmailService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// options - EmailServicesClientBeginDeleteOptions contains the optional parameters for the EmailServicesClient.BeginDelete
// method.
func (client *EmailServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, emailServiceName string, options *EmailServicesClientBeginDeleteOptions) (*runtime.Poller[EmailServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, emailServiceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EmailServicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EmailServicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Operation to delete a EmailService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *EmailServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, emailServiceName string, options *EmailServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, emailServiceName, options)
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
func (client *EmailServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, options *EmailServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the EmailService and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// options - EmailServicesClientGetOptions contains the optional parameters for the EmailServicesClient.Get method.
func (client *EmailServicesClient) Get(ctx context.Context, resourceGroupName string, emailServiceName string, options *EmailServicesClientGetOptions) (EmailServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, emailServiceName, options)
	if err != nil {
		return EmailServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EmailServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EmailServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, options *EmailServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmailServicesClient) getHandleResponse(resp *http.Response) (EmailServicesClientGetResponse, error) {
	result := EmailServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailServiceResource); err != nil {
		return EmailServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Handles requests to list all resources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - EmailServicesClientListByResourceGroupOptions contains the optional parameters for the EmailServicesClient.ListByResourceGroup
// method.
func (client *EmailServicesClient) NewListByResourceGroupPager(resourceGroupName string, options *EmailServicesClientListByResourceGroupOptions) *runtime.Pager[EmailServicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[EmailServicesClientListByResourceGroupResponse]{
		More: func(page EmailServicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EmailServicesClientListByResourceGroupResponse) (EmailServicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EmailServicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EmailServicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EmailServicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *EmailServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *EmailServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices"
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
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *EmailServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (EmailServicesClientListByResourceGroupResponse, error) {
	result := EmailServicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailServiceResourceList); err != nil {
		return EmailServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Handles requests to list all resources in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// options - EmailServicesClientListBySubscriptionOptions contains the optional parameters for the EmailServicesClient.ListBySubscription
// method.
func (client *EmailServicesClient) NewListBySubscriptionPager(options *EmailServicesClientListBySubscriptionOptions) *runtime.Pager[EmailServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[EmailServicesClientListBySubscriptionResponse]{
		More: func(page EmailServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EmailServicesClientListBySubscriptionResponse) (EmailServicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EmailServicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EmailServicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EmailServicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *EmailServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *EmailServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Communication/emailServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *EmailServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (EmailServicesClientListBySubscriptionResponse, error) {
	result := EmailServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmailServiceResourceList); err != nil {
		return EmailServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListVerifiedExchangeOnlineDomains - Get a list of domains that are fully verified in Exchange Online.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// options - EmailServicesClientListVerifiedExchangeOnlineDomainsOptions contains the optional parameters for the EmailServicesClient.ListVerifiedExchangeOnlineDomains
// method.
func (client *EmailServicesClient) ListVerifiedExchangeOnlineDomains(ctx context.Context, options *EmailServicesClientListVerifiedExchangeOnlineDomainsOptions) (EmailServicesClientListVerifiedExchangeOnlineDomainsResponse, error) {
	req, err := client.listVerifiedExchangeOnlineDomainsCreateRequest(ctx, options)
	if err != nil {
		return EmailServicesClientListVerifiedExchangeOnlineDomainsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EmailServicesClientListVerifiedExchangeOnlineDomainsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmailServicesClientListVerifiedExchangeOnlineDomainsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listVerifiedExchangeOnlineDomainsHandleResponse(resp)
}

// listVerifiedExchangeOnlineDomainsCreateRequest creates the ListVerifiedExchangeOnlineDomains request.
func (client *EmailServicesClient) listVerifiedExchangeOnlineDomainsCreateRequest(ctx context.Context, options *EmailServicesClientListVerifiedExchangeOnlineDomainsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Communication/listVerifiedExchangeOnlineDomains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listVerifiedExchangeOnlineDomainsHandleResponse handles the ListVerifiedExchangeOnlineDomains response.
func (client *EmailServicesClient) listVerifiedExchangeOnlineDomainsHandleResponse(resp *http.Response) (EmailServicesClientListVerifiedExchangeOnlineDomainsResponse, error) {
	result := EmailServicesClientListVerifiedExchangeOnlineDomainsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringArray); err != nil {
		return EmailServicesClientListVerifiedExchangeOnlineDomainsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Operation to update an existing EmailService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// emailServiceName - The name of the EmailService resource.
// parameters - Parameters for the update operation
// options - EmailServicesClientBeginUpdateOptions contains the optional parameters for the EmailServicesClient.BeginUpdate
// method.
func (client *EmailServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, emailServiceName string, parameters EmailServiceResourceUpdate, options *EmailServicesClientBeginUpdateOptions) (*runtime.Poller[EmailServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, emailServiceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[EmailServicesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[EmailServicesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Operation to update an existing EmailService.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
func (client *EmailServicesClient) update(ctx context.Context, resourceGroupName string, emailServiceName string, parameters EmailServiceResourceUpdate, options *EmailServicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, emailServiceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *EmailServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, emailServiceName string, parameters EmailServiceResourceUpdate, options *EmailServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Communication/emailServices/{emailServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if emailServiceName == "" {
		return nil, errors.New("parameter emailServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{emailServiceName}", url.PathEscape(emailServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
