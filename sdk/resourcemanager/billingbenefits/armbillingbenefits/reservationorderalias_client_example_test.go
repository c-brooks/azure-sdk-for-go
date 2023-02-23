//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/ReservationOrderAliasCreate.json
func ExampleReservationOrderAliasClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbillingbenefits.NewReservationOrderAliasClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "reservationOrderAlias123", armbillingbenefits.ReservationOrderAliasRequest{
		Location: to.Ptr("eastus"),
		Properties: &armbillingbenefits.ReservationOrderAliasRequestProperties{
			AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
				ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
			},
			AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
			BillingPlan:      to.Ptr(armbillingbenefits.BillingPlanP1M),
			BillingScopeID:   to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
			DisplayName:      to.Ptr("ReservationOrder_2022-06-02"),
			Quantity:         to.Ptr[int32](5),
			Renew:            to.Ptr(true),
			ReservedResourceProperties: &armbillingbenefits.ReservationOrderAliasRequestPropertiesReservedResourceProperties{
				InstanceFlexibility: to.Ptr(armbillingbenefits.InstanceFlexibilityOn),
			},
			ReservedResourceType: to.Ptr(armbillingbenefits.ReservedResourceTypeVirtualMachines),
			Term:                 to.Ptr(armbillingbenefits.TermP1Y),
		},
		SKU: &armbillingbenefits.SKU{
			Name: to.Ptr("Standard_M64s_v2"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/ReservationOrderAliasGet.json
func ExampleReservationOrderAliasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbillingbenefits.NewReservationOrderAliasClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "reservationOrderAlias123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
