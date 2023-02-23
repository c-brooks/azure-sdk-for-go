//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

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

// PublishedArtifactsClient contains the methods for the PublishedArtifacts group.
// Don't use this type directly, use NewPublishedArtifactsClient() instead.
type PublishedArtifactsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPublishedArtifactsClient creates a new instance of PublishedArtifactsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPublishedArtifactsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PublishedArtifactsClient, error) {
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
	client := &PublishedArtifactsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Get an artifact for a published blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// versionID - Version of the published blueprint definition.
// artifactName - Name of the blueprint artifact.
// options - PublishedArtifactsClientGetOptions contains the optional parameters for the PublishedArtifactsClient.Get method.
func (client *PublishedArtifactsClient) Get(ctx context.Context, resourceScope string, blueprintName string, versionID string, artifactName string, options *PublishedArtifactsClientGetOptions) (PublishedArtifactsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceScope, blueprintName, versionID, artifactName, options)
	if err != nil {
		return PublishedArtifactsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PublishedArtifactsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PublishedArtifactsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PublishedArtifactsClient) getCreateRequest(ctx context.Context, resourceScope string, blueprintName string, versionID string, artifactName string, options *PublishedArtifactsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}/artifacts/{artifactName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	if artifactName == "" {
		return nil, errors.New("parameter artifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactName}", url.PathEscape(artifactName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PublishedArtifactsClient) getHandleResponse(resp *http.Response) (PublishedArtifactsClientGetResponse, error) {
	result := PublishedArtifactsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PublishedArtifactsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List artifacts for a version of a published blueprint definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01-preview
// resourceScope - The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
// subscription (format: '/subscriptions/{subscriptionId}').
// blueprintName - Name of the blueprint definition.
// versionID - Version of the published blueprint definition.
// options - PublishedArtifactsClientListOptions contains the optional parameters for the PublishedArtifactsClient.List method.
func (client *PublishedArtifactsClient) NewListPager(resourceScope string, blueprintName string, versionID string, options *PublishedArtifactsClientListOptions) *runtime.Pager[PublishedArtifactsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PublishedArtifactsClientListResponse]{
		More: func(page PublishedArtifactsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PublishedArtifactsClientListResponse) (PublishedArtifactsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceScope, blueprintName, versionID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PublishedArtifactsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PublishedArtifactsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PublishedArtifactsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PublishedArtifactsClient) listCreateRequest(ctx context.Context, resourceScope string, blueprintName string, versionID string, options *PublishedArtifactsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceScope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}/artifacts"
	urlPath = strings.ReplaceAll(urlPath, "{resourceScope}", resourceScope)
	if blueprintName == "" {
		return nil, errors.New("parameter blueprintName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blueprintName}", url.PathEscape(blueprintName))
	if versionID == "" {
		return nil, errors.New("parameter versionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionId}", url.PathEscape(versionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PublishedArtifactsClient) listHandleResponse(resp *http.Response) (PublishedArtifactsClientListResponse, error) {
	result := PublishedArtifactsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactList); err != nil {
		return PublishedArtifactsClientListResponse{}, err
	}
	return result, nil
}
