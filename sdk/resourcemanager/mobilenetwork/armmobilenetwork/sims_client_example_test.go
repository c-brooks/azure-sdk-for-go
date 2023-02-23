//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimDelete.json
func ExampleSimsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "testResourceGroupName", "testSimGroup", "testSim", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimGet.json
func ExampleSimsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "testResourceGroupName", "testSimGroup", "testSimName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimCreate.json
func ExampleSimsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "testSimGroup", "testSim", armmobilenetwork.Sim{
		Properties: &armmobilenetwork.SimPropertiesFormat{
			DeviceType:                            to.Ptr("Video camera"),
			IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
			InternationalMobileSubscriberIdentity: to.Ptr("00000"),
			SimPolicy: &armmobilenetwork.SimPolicyResourceID{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
			},
			StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
				{
					AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
					},
					Slice: &armmobilenetwork.SliceResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
					},
					StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
						IPv4Address: to.Ptr("2.4.0.1"),
					},
				}},
			AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
			OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimListBySimGroup.json
func ExampleSimsClient_NewListByGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByGroupPager("rg1", "testSimGroup", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimBulkUpload.json
func ExampleSimsClient_BeginBulkUpload() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginBulkUpload(ctx, "rg1", "testSimGroup", armmobilenetwork.SimUploadList{
		Sims: []*armmobilenetwork.SimNameAndProperties{
			{
				Name: to.Ptr("testSim"),
				Properties: &armmobilenetwork.SimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.1"),
							},
						}},
					AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
					OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
				},
			},
			{
				Name: to.Ptr("testSim2"),
				Properties: &armmobilenetwork.SimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000001"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.2"),
							},
						}},
					AuthenticationKey: to.Ptr("00000000000000000000000000000000"),
					OperatorKeyCode:   to.Ptr("00000000000000000000000000000000"),
				},
			}},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimBulkDelete.json
func ExampleSimsClient_BeginBulkDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginBulkDelete(ctx, "testResourceGroupName", "testSimGroup", armmobilenetwork.SimDeleteList{
		Sims: []*string{
			to.Ptr("testSim"),
			to.Ptr("testSim2")},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/SimBulkUploadEncrypted.json
func ExampleSimsClient_BeginBulkUploadEncrypted() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginBulkUploadEncrypted(ctx, "rg1", "testSimGroup", armmobilenetwork.EncryptedSimUploadList{
		AzureKeyIdentifier:    to.Ptr[int32](1),
		EncryptedTransportKey: to.Ptr("ABC123"),
		SignedTransportKey:    to.Ptr("ABC123"),
		Sims: []*armmobilenetwork.SimNameAndEncryptedProperties{
			{
				Name: to.Ptr("testSim"),
				Properties: &armmobilenetwork.EncryptedSimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000000"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.1"),
							},
						}},
					EncryptedCredentials: to.Ptr("ABC123"),
				},
			},
			{
				Name: to.Ptr("testSim2"),
				Properties: &armmobilenetwork.EncryptedSimPropertiesFormat{
					DeviceType:                            to.Ptr("Video camera"),
					IntegratedCircuitCardIdentifier:       to.Ptr("8900000000000000001"),
					InternationalMobileSubscriberIdentity: to.Ptr("00000"),
					SimPolicy: &armmobilenetwork.SimPolicyResourceID{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/simPolicies/MySimPolicy"),
					},
					StaticIPConfiguration: []*armmobilenetwork.SimStaticIPProperties{
						{
							AttachedDataNetwork: &armmobilenetwork.AttachedDataNetworkResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/TestPacketCoreCP/packetCoreDataPlanes/TestPacketCoreDP/attachedDataNetworks/TestAttachedDataNetwork"),
							},
							Slice: &armmobilenetwork.SliceResourceID{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork/slices/testSlice"),
							},
							StaticIP: &armmobilenetwork.SimStaticIPPropertiesStaticIP{
								IPv4Address: to.Ptr("2.4.0.2"),
							},
						}},
					EncryptedCredentials: to.Ptr("ABC123"),
				},
			}},
		VendorKeyFingerprint: to.Ptr("ABC123"),
		Version:              to.Ptr[int32](1),
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
