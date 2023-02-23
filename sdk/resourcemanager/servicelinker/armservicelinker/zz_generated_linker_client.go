//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

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

// LinkerClient contains the methods for the Linker group.
// Don't use this type directly, use NewLinkerClient() instead.
type LinkerClient struct {
	host string
	pl   runtime.Pipeline
}

// NewLinkerClient creates a new instance of LinkerClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLinkerClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkerClient, error) {
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
	client := &LinkerClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update linker resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// parameters - Linker details.
// options - LinkerClientBeginCreateOrUpdateOptions contains the optional parameters for the LinkerClient.BeginCreateOrUpdate
// method.
func (client *LinkerClient) BeginCreateOrUpdate(ctx context.Context, resourceURI string, linkerName string, parameters LinkerResource, options *LinkerClientBeginCreateOrUpdateOptions) (*runtime.Poller[LinkerClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceURI, linkerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LinkerClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LinkerClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update linker resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *LinkerClient) createOrUpdate(ctx context.Context, resourceURI string, linkerName string, parameters LinkerResource, options *LinkerClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, linkerName, parameters, options)
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
func (client *LinkerClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, linkerName string, parameters LinkerResource, options *LinkerClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// options - LinkerClientBeginDeleteOptions contains the optional parameters for the LinkerClient.BeginDelete method.
func (client *LinkerClient) BeginDelete(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginDeleteOptions) (*runtime.Poller[LinkerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceURI, linkerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LinkerClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LinkerClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *LinkerClient) deleteOperation(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceURI, linkerName, options)
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
func (client *LinkerClient) deleteCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
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

// Get - Returns Linker resource for a given name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// options - LinkerClientGetOptions contains the optional parameters for the LinkerClient.Get method.
func (client *LinkerClient) Get(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientGetOptions) (LinkerClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return LinkerClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkerClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LinkerClient) getCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
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
func (client *LinkerClient) getHandleResponse(resp *http.Response) (LinkerClientGetResponse, error) {
	result := LinkerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkerResource); err != nil {
		return LinkerClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of Linkers which connects to the resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// options - LinkerClientListOptions contains the optional parameters for the LinkerClient.List method.
func (client *LinkerClient) NewListPager(resourceURI string, options *LinkerClientListOptions) *runtime.Pager[LinkerClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LinkerClientListResponse]{
		More: func(page LinkerClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LinkerClientListResponse) (LinkerClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceURI, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LinkerClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LinkerClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LinkerClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LinkerClient) listCreateRequest(ctx context.Context, resourceURI string, options *LinkerClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
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
func (client *LinkerClient) listHandleResponse(resp *http.Response) (LinkerClientListResponse, error) {
	result := LinkerClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkerList); err != nil {
		return LinkerClientListResponse{}, err
	}
	return result, nil
}

// ListConfigurations - list source configurations for a linker.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// options - LinkerClientListConfigurationsOptions contains the optional parameters for the LinkerClient.ListConfigurations
// method.
func (client *LinkerClient) ListConfigurations(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientListConfigurationsOptions) (LinkerClientListConfigurationsResponse, error) {
	req, err := client.listConfigurationsCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LinkerClientListConfigurationsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listConfigurationsHandleResponse(resp)
}

// listConfigurationsCreateRequest creates the ListConfigurations request.
func (client *LinkerClient) listConfigurationsCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientListConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}/listConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConfigurationsHandleResponse handles the ListConfigurations response.
func (client *LinkerClient) listConfigurationsHandleResponse(resp *http.Response) (LinkerClientListConfigurationsResponse, error) {
	result := LinkerClientListConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SourceConfigurationResult); err != nil {
		return LinkerClientListConfigurationsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Operation to update an existing link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// parameters - Linker details.
// options - LinkerClientBeginUpdateOptions contains the optional parameters for the LinkerClient.BeginUpdate method.
func (client *LinkerClient) BeginUpdate(ctx context.Context, resourceURI string, linkerName string, parameters LinkerPatch, options *LinkerClientBeginUpdateOptions) (*runtime.Poller[LinkerClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceURI, linkerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LinkerClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LinkerClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Operation to update an existing link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *LinkerClient) update(ctx context.Context, resourceURI string, linkerName string, parameters LinkerPatch, options *LinkerClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceURI, linkerName, parameters, options)
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
func (client *LinkerClient) updateCreateRequest(ctx context.Context, resourceURI string, linkerName string, parameters LinkerPatch, options *LinkerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginValidate - Validate a link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
// linkerName - The name Linker resource.
// options - LinkerClientBeginValidateOptions contains the optional parameters for the LinkerClient.BeginValidate method.
func (client *LinkerClient) BeginValidate(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*runtime.Poller[LinkerClientValidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validate(ctx, resourceURI, linkerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[LinkerClientValidateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[LinkerClientValidateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Validate - Validate a link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *LinkerClient) validate(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*http.Response, error) {
	req, err := client.validateCreateRequest(ctx, resourceURI, linkerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// validateCreateRequest creates the Validate request.
func (client *LinkerClient) validateCreateRequest(ctx context.Context, resourceURI string, linkerName string, options *LinkerClientBeginValidateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ServiceLinker/linkers/{linkerName}/validateLinker"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if linkerName == "" {
		return nil, errors.New("parameter linkerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkerName}", url.PathEscape(linkerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
