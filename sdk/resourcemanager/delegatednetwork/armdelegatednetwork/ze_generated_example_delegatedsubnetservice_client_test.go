//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/getDelegatedSubnet.json
func ExampleDelegatedSubnetServiceClient_GetDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetDetails(ctx,
		"TestRG",
		"delegated1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/putDelegatedSubnet.json
func ExampleDelegatedSubnetServiceClient_BeginPutDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPutDetails(ctx,
		"TestRG",
		"delegated1",
		armdelegatednetwork.DelegatedSubnet{
			Location: to.Ptr("West US"),
			Properties: &armdelegatednetwork.DelegatedSubnetProperties{
				ControllerDetails: &armdelegatednetwork.ControllerDetails{
					ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/dnctestcontroller"),
				},
				SubnetDetails: &armdelegatednetwork.SubnetDetails{
					ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/patchDelegatedSubnet.json
func ExampleDelegatedSubnetServiceClient_BeginPatchDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginPatchDetails(ctx,
		"TestRG",
		"delegated1",
		armdelegatednetwork.ResourceUpdateParameters{
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/deleteDelegatedSubnet.json
func ExampleDelegatedSubnetServiceClient_BeginDeleteDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeleteDetails(ctx,
		"TestRG",
		"delegated1",
		&armdelegatednetwork.DelegatedSubnetServiceClientBeginDeleteDetailsOptions{ForceDelete: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/delegatedSubnetListBySub.json
func ExampleDelegatedSubnetServiceClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/delegatedSubnetListByRG.json
func ExampleDelegatedSubnetServiceClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdelegatednetwork.NewDelegatedSubnetServiceClient("50016170-c839-41ba-a724-51e9df440b9e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("testRG",
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
