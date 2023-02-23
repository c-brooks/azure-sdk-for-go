//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// PredictionsClient contains the methods for the Predictions group.
// Don't use this type directly, use NewPredictionsClient() instead.
type PredictionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPredictionsClient creates a new instance of PredictionsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPredictionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PredictionsClient, error) {
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
	client := &PredictionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a Prediction or updates an existing Prediction in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// predictionName - The name of the Prediction.
// parameters - Parameters supplied to the create/update Prediction operation.
// options - PredictionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PredictionsClient.BeginCreateOrUpdate
// method.
func (client *PredictionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionResourceFormat, options *PredictionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PredictionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, predictionName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PredictionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PredictionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a Prediction or updates an existing Prediction in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *PredictionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionResourceFormat, options *PredictionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, predictionName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PredictionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionResourceFormat, options *PredictionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if predictionName == "" {
		return nil, errors.New("parameter predictionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predictionName}", url.PathEscape(predictionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a Prediction in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// predictionName - The name of the Prediction.
// options - PredictionsClientBeginDeleteOptions contains the optional parameters for the PredictionsClient.BeginDelete method.
func (client *PredictionsClient) BeginDelete(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientBeginDeleteOptions) (*runtime.Poller[PredictionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hubName, predictionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PredictionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PredictionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Prediction in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *PredictionsClient) deleteOperation(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, predictionName, options)
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

// deleteCreateRequest creates the Delete request.
func (client *PredictionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if predictionName == "" {
		return nil, errors.New("parameter predictionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predictionName}", url.PathEscape(predictionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a Prediction in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// predictionName - The name of the Prediction.
// options - PredictionsClientGetOptions contains the optional parameters for the PredictionsClient.Get method.
func (client *PredictionsClient) Get(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientGetOptions) (PredictionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, predictionName, options)
	if err != nil {
		return PredictionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PredictionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PredictionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PredictionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if predictionName == "" {
		return nil, errors.New("parameter predictionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predictionName}", url.PathEscape(predictionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PredictionsClient) getHandleResponse(resp *http.Response) (PredictionsClientGetResponse, error) {
	result := PredictionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredictionResourceFormat); err != nil {
		return PredictionsClientGetResponse{}, err
	}
	return result, nil
}

// GetModelStatus - Gets model status of the prediction.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// predictionName - The name of the Prediction.
// options - PredictionsClientGetModelStatusOptions contains the optional parameters for the PredictionsClient.GetModelStatus
// method.
func (client *PredictionsClient) GetModelStatus(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientGetModelStatusOptions) (PredictionsClientGetModelStatusResponse, error) {
	req, err := client.getModelStatusCreateRequest(ctx, resourceGroupName, hubName, predictionName, options)
	if err != nil {
		return PredictionsClientGetModelStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PredictionsClientGetModelStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PredictionsClientGetModelStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getModelStatusHandleResponse(resp)
}

// getModelStatusCreateRequest creates the GetModelStatus request.
func (client *PredictionsClient) getModelStatusCreateRequest(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientGetModelStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}/getModelStatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if predictionName == "" {
		return nil, errors.New("parameter predictionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predictionName}", url.PathEscape(predictionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getModelStatusHandleResponse handles the GetModelStatus response.
func (client *PredictionsClient) getModelStatusHandleResponse(resp *http.Response) (PredictionsClientGetModelStatusResponse, error) {
	result := PredictionsClientGetModelStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredictionModelStatus); err != nil {
		return PredictionsClientGetModelStatusResponse{}, err
	}
	return result, nil
}

// GetTrainingResults - Gets training results.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// predictionName - The name of the Prediction.
// options - PredictionsClientGetTrainingResultsOptions contains the optional parameters for the PredictionsClient.GetTrainingResults
// method.
func (client *PredictionsClient) GetTrainingResults(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientGetTrainingResultsOptions) (PredictionsClientGetTrainingResultsResponse, error) {
	req, err := client.getTrainingResultsCreateRequest(ctx, resourceGroupName, hubName, predictionName, options)
	if err != nil {
		return PredictionsClientGetTrainingResultsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PredictionsClientGetTrainingResultsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PredictionsClientGetTrainingResultsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getTrainingResultsHandleResponse(resp)
}

// getTrainingResultsCreateRequest creates the GetTrainingResults request.
func (client *PredictionsClient) getTrainingResultsCreateRequest(ctx context.Context, resourceGroupName string, hubName string, predictionName string, options *PredictionsClientGetTrainingResultsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}/getTrainingResults"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if predictionName == "" {
		return nil, errors.New("parameter predictionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predictionName}", url.PathEscape(predictionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTrainingResultsHandleResponse handles the GetTrainingResults response.
func (client *PredictionsClient) getTrainingResultsHandleResponse(resp *http.Response) (PredictionsClientGetTrainingResultsResponse, error) {
	result := PredictionsClientGetTrainingResultsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredictionTrainingResults); err != nil {
		return PredictionsClientGetTrainingResultsResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all the predictions in the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - PredictionsClientListByHubOptions contains the optional parameters for the PredictionsClient.ListByHub method.
func (client *PredictionsClient) NewListByHubPager(resourceGroupName string, hubName string, options *PredictionsClientListByHubOptions) *runtime.Pager[PredictionsClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[PredictionsClientListByHubResponse]{
		More: func(page PredictionsClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PredictionsClientListByHubResponse) (PredictionsClientListByHubResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PredictionsClientListByHubResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PredictionsClientListByHubResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PredictionsClientListByHubResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHubHandleResponse(resp)
		},
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *PredictionsClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *PredictionsClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *PredictionsClient) listByHubHandleResponse(resp *http.Response) (PredictionsClientListByHubResponse, error) {
	result := PredictionsClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PredictionListResult); err != nil {
		return PredictionsClientListByHubResponse{}, err
	}
	return result, nil
}

// ModelStatus - Creates or updates the model status of prediction.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// predictionName - The name of the Prediction.
// parameters - Parameters supplied to the create/update prediction model status operation.
// options - PredictionsClientModelStatusOptions contains the optional parameters for the PredictionsClient.ModelStatus method.
func (client *PredictionsClient) ModelStatus(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionModelStatus, options *PredictionsClientModelStatusOptions) (PredictionsClientModelStatusResponse, error) {
	req, err := client.modelStatusCreateRequest(ctx, resourceGroupName, hubName, predictionName, parameters, options)
	if err != nil {
		return PredictionsClientModelStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PredictionsClientModelStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PredictionsClientModelStatusResponse{}, runtime.NewResponseError(resp)
	}
	return PredictionsClientModelStatusResponse{}, nil
}

// modelStatusCreateRequest creates the ModelStatus request.
func (client *PredictionsClient) modelStatusCreateRequest(ctx context.Context, resourceGroupName string, hubName string, predictionName string, parameters PredictionModelStatus, options *PredictionsClientModelStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}/modelStatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if predictionName == "" {
		return nil, errors.New("parameter predictionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{predictionName}", url.PathEscape(predictionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}
