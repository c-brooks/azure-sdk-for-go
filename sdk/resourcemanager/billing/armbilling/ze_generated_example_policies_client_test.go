//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Policy.json
func ExamplePoliciesClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetByBillingProfile(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdatePolicy.json
func ExamplePoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		armbilling.Policy{
			Properties: &armbilling.PolicyProperties{
				MarketplacePurchases: to.Ptr(armbilling.MarketplacePurchasesPolicyOnlyFreeAllowed),
				ReservationPurchases: to.Ptr(armbilling.ReservationPurchasesPolicyNotAllowed),
				ViewCharges:          to.Ptr(armbilling.ViewChargesPolicyAllowed),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPolicy.json
func ExamplePoliciesClient_GetByCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetByCustomer(ctx,
		"{billingAccountName}",
		"{customerName}",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateCustomerPolicy.json
func ExamplePoliciesClient_UpdateCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateCustomer(ctx,
		"{billingAccountName}",
		"{customerName}",
		armbilling.CustomerPolicy{
			Properties: &armbilling.CustomerPolicyProperties{
				ViewCharges: to.Ptr(armbilling.ViewChargesNotAllowed),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
