//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Get.json
func ExampleMembersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"contosemember1",
		"mygroup",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Create.json
func ExampleMembersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"contosemember1",
		"mygroup",
		&armblockchain.MembersClientBeginCreateOptions{BlockchainMember: &armblockchain.Member{
			Location: to.Ptr("southeastasia"),
			Properties: &armblockchain.MemberProperties{
				Consortium:                          to.Ptr("ContoseConsortium"),
				ConsortiumManagementAccountPassword: to.Ptr("<consortiumManagementAccountPassword>"),
				Password:                            to.Ptr("<password>"),
				ValidatorNodesSKU: &armblockchain.MemberNodesSKU{
					Capacity: to.Ptr[int32](2),
				},
				Protocol: to.Ptr(armblockchain.BlockchainProtocolQuorum),
			},
		},
		})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Delete.json
func ExampleMembersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"contosemember1",
		"mygroup",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Update.json
func ExampleMembersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"ContoseMember1",
		"mygroup",
		&armblockchain.MembersClientUpdateOptions{BlockchainMember: &armblockchain.MemberUpdate{
			Properties: &armblockchain.MemberPropertiesUpdate{
				Password:                            to.Ptr("<password>"),
				ConsortiumManagementAccountPassword: to.Ptr("<consortiumManagementAccountPassword>"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_List.json
func ExampleMembersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("mygroup",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListAll.json
func ExampleMembersClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAllPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListConsortiumMembers.json
func ExampleMembersClient_NewListConsortiumMembersPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblockchain.NewMembersClient("51766542-3ed7-4a72-a187-0c8ab644ddab", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListConsortiumMembersPager("contosemember1",
		"mygroup",
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
