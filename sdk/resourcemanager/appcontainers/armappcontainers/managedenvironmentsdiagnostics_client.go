//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// ManagedEnvironmentsDiagnosticsClient contains the methods for the ManagedEnvironmentsDiagnostics group.
// Don't use this type directly, use NewManagedEnvironmentsDiagnosticsClient() instead.
type ManagedEnvironmentsDiagnosticsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagedEnvironmentsDiagnosticsClient creates a new instance of ManagedEnvironmentsDiagnosticsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagedEnvironmentsDiagnosticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedEnvironmentsDiagnosticsClient, error) {
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
	client := &ManagedEnvironmentsDiagnosticsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetRoot - Get the properties of a Managed Environment used to host container apps.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// environmentName - Name of the Environment.
// options - ManagedEnvironmentsDiagnosticsClientGetRootOptions contains the optional parameters for the ManagedEnvironmentsDiagnosticsClient.GetRoot
// method.
func (client *ManagedEnvironmentsDiagnosticsClient) GetRoot(ctx context.Context, resourceGroupName string, environmentName string, options *ManagedEnvironmentsDiagnosticsClientGetRootOptions) (ManagedEnvironmentsDiagnosticsClientGetRootResponse, error) {
	req, err := client.getRootCreateRequest(ctx, resourceGroupName, environmentName, options)
	if err != nil {
		return ManagedEnvironmentsDiagnosticsClientGetRootResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedEnvironmentsDiagnosticsClientGetRootResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedEnvironmentsDiagnosticsClientGetRootResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRootHandleResponse(resp)
}

// getRootCreateRequest creates the GetRoot request.
func (client *ManagedEnvironmentsDiagnosticsClient) getRootCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *ManagedEnvironmentsDiagnosticsClientGetRootOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/detectorProperties/rootApi/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRootHandleResponse handles the GetRoot response.
func (client *ManagedEnvironmentsDiagnosticsClient) getRootHandleResponse(resp *http.Response) (ManagedEnvironmentsDiagnosticsClientGetRootResponse, error) {
	result := ManagedEnvironmentsDiagnosticsClientGetRootResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedEnvironment); err != nil {
		return ManagedEnvironmentsDiagnosticsClientGetRootResponse{}, err
	}
	return result, nil
}
