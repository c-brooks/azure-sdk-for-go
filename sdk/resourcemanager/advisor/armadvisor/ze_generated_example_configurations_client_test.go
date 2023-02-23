//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListConfigurations.json
func ExampleConfigurationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armadvisor.NewConfigurationsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
func ExampleConfigurationsClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armadvisor.NewConfigurationsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateInSubscription(ctx,
		armadvisor.ConfigurationNameDefault,
		armadvisor.ConfigData{
			Properties: &armadvisor.ConfigDataProperties{
				Digests: []*armadvisor.DigestConfig{
					{
						Name:                  to.Ptr("digestConfigName"),
						ActionGroupResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName"),
						Categories: []*armadvisor.Category{
							to.Ptr(armadvisor.CategoryHighAvailability),
							to.Ptr(armadvisor.CategorySecurity),
							to.Ptr(armadvisor.CategoryPerformance),
							to.Ptr(armadvisor.CategoryCost),
							to.Ptr(armadvisor.CategoryOperationalExcellence)},
						Frequency: to.Ptr[int32](30),
						State:     to.Ptr(armadvisor.DigestConfigStateActive),
						Language:  to.Ptr("en"),
					}},
				Exclude:         to.Ptr(true),
				LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListConfigurations.json
func ExampleConfigurationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armadvisor.NewConfigurationsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("resourceGroup",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
func ExampleConfigurationsClient_CreateInResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armadvisor.NewConfigurationsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateInResourceGroup(ctx,
		armadvisor.ConfigurationNameDefault,
		"resourceGroup",
		armadvisor.ConfigData{
			Properties: &armadvisor.ConfigDataProperties{
				Digests: []*armadvisor.DigestConfig{
					{
						Name:                  to.Ptr("digestConfigName"),
						ActionGroupResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName"),
						Categories: []*armadvisor.Category{
							to.Ptr(armadvisor.CategoryHighAvailability),
							to.Ptr(armadvisor.CategorySecurity),
							to.Ptr(armadvisor.CategoryPerformance),
							to.Ptr(armadvisor.CategoryCost),
							to.Ptr(armadvisor.CategoryOperationalExcellence)},
						Frequency: to.Ptr[int32](30),
						State:     to.Ptr(armadvisor.DigestConfigStateActive),
						Language:  to.Ptr("en"),
					}},
				Exclude:         to.Ptr(true),
				LowCPUThreshold: to.Ptr(armadvisor.CPUThresholdFive),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
