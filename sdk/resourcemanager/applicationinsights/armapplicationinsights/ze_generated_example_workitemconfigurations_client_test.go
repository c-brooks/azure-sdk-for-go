//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigsGet.json
func ExampleWorkItemConfigurationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("my-resource-group",
		"my-component",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigCreate.json
func ExampleWorkItemConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"my-resource-group",
		"my-component",
		armapplicationinsights.WorkItemCreateConfiguration{
			ConnectorDataConfiguration: to.Ptr("{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"d0662b05-439a-4a1b-840b-33a7f8b42ebf\",\"Custom\":\"{\\\"/fields/System.WorkItemType\\\":\\\"Bug\\\",\\\"/fields/System.AreaPath\\\":\\\"todeletefirst\\\",\\\"/fields/System.AssignedTo\\\":\\\"\\\"}\"}"),
			ConnectorID:                to.Ptr("d334e2a4-6733-488e-8645-a9fdc1694f41"),
			ValidateOnly:               to.Ptr(true),
			WorkItemProperties: map[string]*string{
				"0": to.Ptr("[object Object]"),
				"1": to.Ptr("[object Object]"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigDefaultGet.json
func ExampleWorkItemConfigurationsClient_GetDefault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetDefault(ctx,
		"my-resource-group",
		"my-component",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigDelete.json
func ExampleWorkItemConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"my-resource-group",
		"my-component",
		"Visual Studio Team Services",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigGet.json
func ExampleWorkItemConfigurationsClient_GetItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetItem(ctx,
		"my-resource-group",
		"my-component",
		"Visual Studio Team Services",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigUpdate.json
func ExampleWorkItemConfigurationsClient_UpdateItem() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewWorkItemConfigurationsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateItem(ctx,
		"my-resource-group",
		"my-component",
		"Visual Studio Team Services",
		armapplicationinsights.WorkItemCreateConfiguration{
			ConnectorDataConfiguration: to.Ptr("{\"VSOAccountBaseUrl\":\"https://testtodelete.visualstudio.com\",\"ProjectCollection\":\"DefaultCollection\",\"Project\":\"todeletefirst\",\"ResourceId\":\"d0662b05-439a-4a1b-840b-33a7f8b42ebf\",\"Custom\":\"{\\\"/fields/System.WorkItemType\\\":\\\"Bug\\\",\\\"/fields/System.AreaPath\\\":\\\"todeletefirst\\\",\\\"/fields/System.AssignedTo\\\":\\\"\\\"}\"}"),
			ConnectorID:                to.Ptr("d334e2a4-6733-488e-8645-a9fdc1694f41"),
			ValidateOnly:               to.Ptr(true),
			WorkItemProperties: map[string]*string{
				"0": to.Ptr("[object Object]"),
				"1": to.Ptr("[object Object]"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
