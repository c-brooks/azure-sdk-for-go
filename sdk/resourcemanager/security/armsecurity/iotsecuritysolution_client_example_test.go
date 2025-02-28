//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByIotHub.json
func ExampleIotSecuritySolutionClient_NewListBySubscriptionPager_listIoTSecuritySolutionsByIoTHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(&armsecurity.IotSecuritySolutionClientListBySubscriptionOptions{Filter: to.Ptr("properties.iotHubs/any(i eq \"/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub\")")})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsList.json
func ExampleIotSecuritySolutionClient_NewListBySubscriptionPager_listIoTSecuritySolutionsBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(&armsecurity.IotSecuritySolutionClientListBySubscriptionOptions{Filter: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByRg.json
func ExampleIotSecuritySolutionClient_NewListByResourceGroupPager_listIoTSecuritySolutionsByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("MyGroup", &armsecurity.IotSecuritySolutionClientListByResourceGroupOptions{Filter: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByIotHubAndRg.json
func ExampleIotSecuritySolutionClient_NewListByResourceGroupPager_listIoTSecuritySolutionsByResourceGroupAndIoTHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("MyRg", &armsecurity.IotSecuritySolutionClientListByResourceGroupOptions{Filter: to.Ptr("properties.iotHubs/any(i eq \"/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub\")")})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "MyGroup", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/CreateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyGroup", "default", armsecurity.IoTSecuritySolutionModel{
		Tags:     map[string]*string{},
		Location: to.Ptr("East Us"),
		Properties: &armsecurity.IoTSecuritySolutionProperties{
			DisabledDataSources: []*armsecurity.DataSource{},
			DisplayName:         to.Ptr("Solution Default"),
			Export:              []*armsecurity.ExportData{},
			IotHubs: []*string{
				to.Ptr("/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub")},
			RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
				{
					RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTOpenPorts),
					Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
				},
				{
					RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTSharedCredentials),
					Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
				}},
			Status:                  to.Ptr(armsecurity.SecuritySolutionStatusEnabled),
			UnmaskedIPLoggingStatus: to.Ptr(armsecurity.UnmaskedIPLoggingStatusEnabled),
			UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
				Query: to.Ptr("where type != \"microsoft.devices/iothubs\" | where name contains \"iot\""),
				QuerySubscriptions: []*string{
					to.Ptr("075423e9-7d33-4166-8bdf-3920b04e3735")},
			},
			Workspace: to.Ptr("/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "myRg", "default", armsecurity.UpdateIotSecuritySolutionData{
		Tags: map[string]*string{
			"foo": to.Ptr("bar"),
		},
		Properties: &armsecurity.UpdateIoTSecuritySolutionProperties{
			RecommendationsConfiguration: []*armsecurity.RecommendationConfigurationProperties{
				{
					RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTOpenPorts),
					Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
				},
				{
					RecommendationType: to.Ptr(armsecurity.RecommendationTypeIoTSharedCredentials),
					Status:             to.Ptr(armsecurity.RecommendationConfigStatusDisabled),
				}},
			UserDefinedResources: &armsecurity.UserDefinedResourcesProperties{
				Query: to.Ptr("where type != \"microsoft.devices/iothubs\" | where name contains \"v2\""),
				QuerySubscriptions: []*string{
					to.Ptr("075423e9-7d33-4166-8bdf-3920b04e3735")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/DeleteIoTSecuritySolution.json
func ExampleIotSecuritySolutionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewIotSecuritySolutionClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "MyGroup", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
