//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azquery

import (
	"context"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// NewListDefinitionsPager - Lists the metric definitions for the resource.
// Generated from API version 2018-01-01
// resourceURI - The identifier of the resource.
// options - MetricsClientListDefinitionsOptions contains the optional parameters for the MetricsClient.ListDefinitions method.
func (client *MetricsClient) NewListDefinitionsPager(resourceURI string, options *MetricsClientListDefinitionsOptions) *runtime.Pager[MetricsClientListDefinitionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetricsClientListDefinitionsResponse]{
		More: func(page MetricsClientListDefinitionsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *MetricsClientListDefinitionsResponse) (MetricsClientListDefinitionsResponse, error) {
			req, err := client.listDefinitionsCreateRequest(ctx, resourceURI, options)
			if err != nil {
				return MetricsClientListDefinitionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MetricsClientListDefinitionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MetricsClientListDefinitionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listDefinitionsHandleResponse(resp)
		},
	})
}

// listDefinitionsCreateRequest creates the ListDefinitions request.
func (client *MetricsClient) listDefinitionsCreateRequest(ctx context.Context, resourceURI string, options *MetricsClientListDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metricDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-01-01")
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDefinitionsHandleResponse handles the ListDefinitions response.
func (client *MetricsClient) listDefinitionsHandleResponse(resp *http.Response) (MetricsClientListDefinitionsResponse, error) {
	result := MetricsClientListDefinitionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricDefinitionCollection); err != nil {
		return MetricsClientListDefinitionsResponse{}, err
	}
	return result, nil
}

// NewListNamespacesPager - Lists the metric namespaces for the resource.
// Generated from API version 2017-12-01-preview
// resourceURI - The identifier of the resource.
// options - MetricsClientListNamespacesOptions contains the optional parameters for the MetricsClient.ListNamespaces method.
func (client *MetricsClient) NewListNamespacesPager(resourceURI string, options *MetricsClientListNamespacesOptions) *runtime.Pager[MetricsClientListNamespacesResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetricsClientListNamespacesResponse]{
		More: func(page MetricsClientListNamespacesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *MetricsClientListNamespacesResponse) (MetricsClientListNamespacesResponse, error) {
			req, err := client.listNamespacesCreateRequest(ctx, resourceURI, options)
			if err != nil {
				return MetricsClientListNamespacesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MetricsClientListNamespacesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MetricsClientListNamespacesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listNamespacesHandleResponse(resp)
		},
	})
}

// listNamespacesCreateRequest creates the ListNamespaces request.
func (client *MetricsClient) listNamespacesCreateRequest(ctx context.Context, resourceURI string, options *MetricsClientListNamespacesOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/metricNamespaces"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01-preview")
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listNamespacesHandleResponse handles the ListNamespaces response.
func (client *MetricsClient) listNamespacesHandleResponse(resp *http.Response) (MetricsClientListNamespacesResponse, error) {
	result := MetricsClientListNamespacesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetricNamespaceCollection); err != nil {
		return MetricsClientListNamespacesResponse{}, err
	}
	return result, nil
}

// QueryResource - Lists the metric values for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-01-01
// resourceURI - The identifier of the resource.
// options - MetricsClientQueryResourceOptions contains the optional parameters for the MetricsClient.QueryResource method.
func (client *MetricsClient) QueryResource(ctx context.Context, resourceURI string, options *MetricsClientQueryResourceOptions) (MetricsClientQueryResourceResponse, error) {
	req, err := client.queryResourceCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return MetricsClientQueryResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricsClientQueryResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricsClientQueryResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryResourceHandleResponse(resp)
}

// queryResourceCreateRequest creates the QueryResource request.
func (client *MetricsClient) queryResourceCreateRequest(ctx context.Context, resourceURI string, options *MetricsClientQueryResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metrics"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", string(*options.Timespan))
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2018-01-01")
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// queryResourceHandleResponse handles the QueryResource response.
func (client *MetricsClient) queryResourceHandleResponse(resp *http.Response) (MetricsClientQueryResourceResponse, error) {
	result := MetricsClientQueryResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Response); err != nil {
		return MetricsClientQueryResourceResponse{}, err
	}
	return result, nil
}
