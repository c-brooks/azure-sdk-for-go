//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStoreOffers.json
func ExamplePrivateStoreCollectionOfferClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/GetPrivateStoreCollectionOffer.json
func ExamplePrivateStoreCollectionOfferClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"marketplacetestthirdparty.md-test-third-party-2",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/PrivateStoreOffer_update.json
func ExamplePrivateStoreCollectionOfferClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"marketplacetestthirdparty.md-test-third-party-2",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		&armmarketplace.PrivateStoreCollectionOfferClientCreateOrUpdateOptions{Payload: &armmarketplace.Offer{
			Properties: &armmarketplace.OfferProperties{
				ETag: to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350666\""),
				SpecificPlanIDsLimitation: []*string{
					to.Ptr("0001"),
					to.Ptr("0002")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/DeletePrivateStoreOffer.json
func ExamplePrivateStoreCollectionOfferClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"marketplacetestthirdparty.md-test-third-party-2",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/PostPrivateStoreCollectionOffer.json
func ExamplePrivateStoreCollectionOfferClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Post(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"marketplacetestthirdparty.md-test-third-party-2",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		&armmarketplace.PrivateStoreCollectionOfferClientPostOptions{Payload: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
