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

	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherCreate.json
func ExampleWatchersClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg1", "nw1", armnetwork.Watcher{
		Location:   to.Ptr("eastus"),
		Properties: &armnetwork.WatcherPropertiesFormat{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherGet.json
func ExampleWatchersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "nw1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherDelete.json
func ExampleWatchersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "nw1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherUpdateTags.json
func ExampleWatchersClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx, "rg1", "nw1", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherList.json
func ExampleWatchersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherListAll.json
func ExampleWatchersClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherTopologyGet.json
func ExampleWatchersClient_GetTopology() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetTopology(ctx, "rg1", "nw1", armnetwork.TopologyParameters{
		TargetResourceGroupName: to.Ptr("rg2"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherIpFlowVerify.json
func ExampleWatchersClient_BeginVerifyIPFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginVerifyIPFlow(ctx, "rg1", "nw1", armnetwork.VerificationIPFlowParameters{
		Direction:        to.Ptr(armnetwork.DirectionOutbound),
		LocalIPAddress:   to.Ptr("10.2.0.4"),
		LocalPort:        to.Ptr("80"),
		RemoteIPAddress:  to.Ptr("121.10.1.1"),
		RemotePort:       to.Ptr("80"),
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
		Protocol:         to.Ptr(armnetwork.IPFlowProtocolTCP),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherNextHopGet.json
func ExampleWatchersClient_BeginGetNextHop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetNextHop(ctx, "rg1", "nw1", armnetwork.NextHopParameters{
		DestinationIPAddress: to.Ptr("10.0.0.10"),
		SourceIPAddress:      to.Ptr("10.0.0.5"),
		TargetNicResourceID:  to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1"),
		TargetResourceID:     to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherSecurityGroupViewGet.json
func ExampleWatchersClient_BeginGetVMSecurityRules() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetVMSecurityRules(ctx, "rg1", "nw1", armnetwork.SecurityGroupViewParameters{
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherTroubleshootGet.json
func ExampleWatchersClient_BeginGetTroubleshooting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetTroubleshooting(ctx, "rg1", "nw1", armnetwork.TroubleshootingParameters{
		Properties: &armnetwork.TroubleshootingProperties{
			StorageID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1"),
			StoragePath: to.Ptr("https://st1.blob.core.windows.net/cn1"),
		},
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherTroubleshootResultQuery.json
func ExampleWatchersClient_BeginGetTroubleshootingResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetTroubleshootingResult(ctx, "rg1", "nw1", armnetwork.QueryTroubleshootingParameters{
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherFlowLogConfigure.json
func ExampleWatchersClient_BeginSetFlowLogConfiguration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSetFlowLogConfiguration(ctx, "rg1", "nw1", armnetwork.FlowLogInformation{
		Properties: &armnetwork.FlowLogProperties{
			Enabled:   to.Ptr(true),
			StorageID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1"),
		},
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherFlowLogStatusQuery.json
func ExampleWatchersClient_BeginGetFlowLogStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetFlowLogStatus(ctx, "rg1", "nw1", armnetwork.FlowLogStatusParameters{
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherConnectivityCheck.json
func ExampleWatchersClient_BeginCheckConnectivity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCheckConnectivity(ctx, "rg1", "nw1", armnetwork.ConnectivityParameters{
		Destination: &armnetwork.ConnectivityDestination{
			Address: to.Ptr("192.168.100.4"),
			Port:    to.Ptr[int32](3389),
		},
		PreferredIPVersion: to.Ptr(armnetwork.IPVersionIPv4),
		Source: &armnetwork.ConnectivitySource{
			ResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherAzureReachabilityReportGet.json
func ExampleWatchersClient_BeginGetAzureReachabilityReport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetAzureReachabilityReport(ctx, "rg1", "nw1", armnetwork.AzureReachabilityReportParameters{
		AzureLocations: []*string{
			to.Ptr("West US")},
		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-10T00:00:00Z"); return t }()),
		ProviderLocation: &armnetwork.AzureReachabilityReportLocation{
			Country: to.Ptr("United States"),
			State:   to.Ptr("washington"),
		},
		Providers: []*string{
			to.Ptr("Frontier Communications of America, Inc. - ASN 5650")},
		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-09-07T00:00:00Z"); return t }()),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherAvailableProvidersListGet.json
func ExampleWatchersClient_BeginListAvailableProviders() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginListAvailableProviders(ctx, "rg1", "nw1", armnetwork.AvailableProvidersListParameters{
		AzureLocations: []*string{
			to.Ptr("West US")},
		City:    to.Ptr("seattle"),
		Country: to.Ptr("United States"),
		State:   to.Ptr("washington"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NetworkWatcherNetworkConfigurationDiagnostic.json
func ExampleWatchersClient_BeginGetNetworkConfigurationDiagnostic() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetNetworkConfigurationDiagnostic(ctx, "rg1", "nw1", armnetwork.ConfigurationDiagnosticParameters{
		Profiles: []*armnetwork.ConfigurationDiagnosticProfile{
			{
				Destination:     to.Ptr("12.11.12.14"),
				DestinationPort: to.Ptr("12100"),
				Direction:       to.Ptr(armnetwork.DirectionInbound),
				Source:          to.Ptr("10.1.0.4"),
				Protocol:        to.Ptr("TCP"),
			}},
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
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
