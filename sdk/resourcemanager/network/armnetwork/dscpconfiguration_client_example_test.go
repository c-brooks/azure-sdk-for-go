//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DscpConfigurationCreate.json
func ExampleDscpConfigurationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewDscpConfigurationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "mydscpconfig", armnetwork.DscpConfiguration{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.DscpConfigurationPropertiesFormat{
			QosDefinitionCollection: []*armnetwork.QosDefinition{
				{
					DestinationIPRanges: []*armnetwork.QosIPRange{
						{
							EndIP:   to.Ptr("127.0.10.2"),
							StartIP: to.Ptr("127.0.10.1"),
						}},
					DestinationPortRanges: []*armnetwork.QosPortRange{
						{
							End:   to.Ptr[int32](15),
							Start: to.Ptr[int32](15),
						}},
					Markings: []*int32{
						to.Ptr[int32](1)},
					SourceIPRanges: []*armnetwork.QosIPRange{
						{
							EndIP:   to.Ptr("127.0.0.2"),
							StartIP: to.Ptr("127.0.0.1"),
						}},
					SourcePortRanges: []*armnetwork.QosPortRange{
						{
							End:   to.Ptr[int32](11),
							Start: to.Ptr[int32](10),
						},
						{
							End:   to.Ptr[int32](21),
							Start: to.Ptr[int32](20),
						}},
					Protocol: to.Ptr(armnetwork.ProtocolTypeTCP),
				},
				{
					DestinationIPRanges: []*armnetwork.QosIPRange{
						{
							EndIP:   to.Ptr("12.0.10.2"),
							StartIP: to.Ptr("12.0.10.1"),
						}},
					DestinationPortRanges: []*armnetwork.QosPortRange{
						{
							End:   to.Ptr[int32](52),
							Start: to.Ptr[int32](51),
						}},
					Markings: []*int32{
						to.Ptr[int32](2)},
					SourceIPRanges: []*armnetwork.QosIPRange{
						{
							EndIP:   to.Ptr("12.0.0.2"),
							StartIP: to.Ptr("12.0.0.1"),
						}},
					SourcePortRanges: []*armnetwork.QosPortRange{
						{
							End:   to.Ptr[int32](12),
							Start: to.Ptr[int32](11),
						}},
					Protocol: to.Ptr(armnetwork.ProtocolTypeUDP),
				}},
		},
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DscpConfigurationDelete.json
func ExampleDscpConfigurationClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewDscpConfigurationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "mydscpConfig", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DscpConfigurationGet.json
func ExampleDscpConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewDscpConfigurationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "mydscpConfig", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DscpConfigurationList.json
func ExampleDscpConfigurationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewDscpConfigurationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DscpConfigurationListAll.json
func ExampleDscpConfigurationClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewDscpConfigurationClient("subid", cred, nil)
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
