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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupDelete.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyNatRuleCollectionGroupGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyNatRuleCollectionGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyRuleCollectionGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyRuleCollectionGroupWithIpGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "firewallPolicy", "ruleGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesGet.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_Get_getFirewallPolicyRuleCollectionGroupWithWebCategories() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("e747cc13-97d4-4a79-b463-42d7f4e558f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyNatRuleCollectionGroupPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_createFirewallPolicyNatRuleCollectionGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](100),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyNatRuleCollection{
					Name:               to.Ptr("Example-Nat-Rule-Collection"),
					Priority:           to.Ptr[int32](100),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyNatRuleCollection),
					Action: &armnetwork.FirewallPolicyNatRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyNatRuleCollectionActionTypeDNAT),
					},
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.NatRule{
							Name:     to.Ptr("nat-rule1"),
							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNatRule),
							DestinationAddresses: []*string{
								to.Ptr("152.23.32.23")},
							DestinationPorts: []*string{
								to.Ptr("8080")},
							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP),
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolUDP)},
							SourceAddresses: []*string{
								to.Ptr("2.2.2.2")},
							SourceIPGroups: []*string{},
							TranslatedFqdn: to.Ptr("internalhttp.server.net"),
							TranslatedPort: to.Ptr("8080"),
						}},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_createFirewallPolicyRuleCollectionGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](100),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyFilterRuleCollection{
					Name:               to.Ptr("Example-Filter-Rule-Collection"),
					Priority:           to.Ptr[int32](100),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
					},
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.Rule{
							Name:     to.Ptr("network-rule1"),
							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
							DestinationAddresses: []*string{
								to.Ptr("*")},
							DestinationPorts: []*string{
								to.Ptr("*")},
							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP)},
							SourceAddresses: []*string{
								to.Ptr("10.1.25.0/24")},
						}},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_createFirewallPolicyRuleCollectionGroupWithIpGroups() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](110),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyFilterRuleCollection{
					Name:               to.Ptr("Example-Filter-Rule-Collection"),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
					},
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.Rule{
							Name:     to.Ptr("network-1"),
							RuleType: to.Ptr(armnetwork.FirewallPolicyRuleTypeNetworkRule),
							DestinationIPGroups: []*string{
								to.Ptr("/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups2")},
							DestinationPorts: []*string{
								to.Ptr("*")},
							IPProtocols: []*armnetwork.FirewallPolicyRuleNetworkProtocol{
								to.Ptr(armnetwork.FirewallPolicyRuleNetworkProtocolTCP)},
							SourceIPGroups: []*string{
								to.Ptr("/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups1")},
						}},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesPut.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_BeginCreateOrUpdate_createFirewallPolicyRuleCollectionGroupWithWebCategories() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("e747cc13-97d4-4a79-b463-42d7f4e558f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", "ruleCollectionGroup1", armnetwork.FirewallPolicyRuleCollectionGroup{
		Properties: &armnetwork.FirewallPolicyRuleCollectionGroupProperties{
			Priority: to.Ptr[int32](110),
			RuleCollections: []armnetwork.FirewallPolicyRuleCollectionClassification{
				&armnetwork.FirewallPolicyFilterRuleCollection{
					Name:               to.Ptr("Example-Filter-Rule-Collection"),
					RuleCollectionType: to.Ptr(armnetwork.FirewallPolicyRuleCollectionTypeFirewallPolicyFilterRuleCollection),
					Action: &armnetwork.FirewallPolicyFilterRuleCollectionAction{
						Type: to.Ptr(armnetwork.FirewallPolicyFilterRuleCollectionActionTypeDeny),
					},
					Rules: []armnetwork.FirewallPolicyRuleClassification{
						&armnetwork.ApplicationRule{
							Name:        to.Ptr("rule1"),
							Description: to.Ptr("Deny inbound rule"),
							RuleType:    to.Ptr(armnetwork.FirewallPolicyRuleTypeApplicationRule),
							Protocols: []*armnetwork.FirewallPolicyRuleApplicationProtocol{
								{
									Port:         to.Ptr[int32](443),
									ProtocolType: to.Ptr(armnetwork.FirewallPolicyRuleApplicationProtocolTypeHTTPS),
								}},
							SourceAddresses: []*string{
								to.Ptr("216.58.216.164"),
								to.Ptr("10.0.0.0/24")},
							WebCategories: []*string{
								to.Ptr("Hacking")},
						}},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupWithWebCategoriesList.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_NewListPager_listAllFirewallPolicyRuleCollectionGroupWithWebCategories() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("e747cc13-97d4-4a79-b463-42d7f4e558f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", "firewallPolicy", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupList.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_NewListPager_listAllFirewallPolicyRuleCollectionGroupsForAGivenFirewallPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", "firewallPolicy", nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsList.json
func ExampleFirewallPolicyRuleCollectionGroupsClient_NewListPager_listAllFirewallPolicyRuleCollectionGroupsWithIpGroupsForAGivenFirewallPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPolicyRuleCollectionGroupsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", "firewallPolicy", nil)
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
