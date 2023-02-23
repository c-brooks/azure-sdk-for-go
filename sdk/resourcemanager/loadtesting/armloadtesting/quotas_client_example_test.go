//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armloadtesting_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_List.json
func ExampleQuotasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armloadtesting.NewQuotasClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("westus", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.QuotaResourceList = armloadtesting.QuotaResourceList{
		// 	Value: []*armloadtesting.QuotaResource{
		// 		{
		// 			Name: to.Ptr("testQuotaBucket"),
		// 			Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
		// 			Properties: &armloadtesting.QuotaResourceProperties{
		// 				Limit: to.Ptr[int32](50),
		// 				Usage: to.Ptr[int32](20),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_Get.json
func ExampleQuotasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armloadtesting.NewQuotasClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "westus", "testQuotaBucket", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QuotaResource = armloadtesting.QuotaResource{
	// 	Name: to.Ptr("testQuotaBucket"),
	// 	Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
	// 	Properties: &armloadtesting.QuotaResourceProperties{
	// 		Limit: to.Ptr[int32](50),
	// 		Usage: to.Ptr[int32](20),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_CheckAvailability.json
func ExampleQuotasClient_CheckAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armloadtesting.NewQuotasClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckAvailability(ctx, "westus", "testQuotaBucket", armloadtesting.QuotaBucketRequest{
		Properties: &armloadtesting.QuotaBucketRequestProperties{
			CurrentQuota: to.Ptr[int32](40),
			CurrentUsage: to.Ptr[int32](20),
			Dimensions: &armloadtesting.QuotaBucketRequestPropertiesDimensions{
				Location:       to.Ptr("westus"),
				SubscriptionID: to.Ptr("testsubscriptionId"),
			},
			NewQuota: to.Ptr[int32](50),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckQuotaAvailabilityResponse = armloadtesting.CheckQuotaAvailabilityResponse{
	// 	Name: to.Ptr("testQuotaBucket"),
	// 	Type: to.Ptr("Microsoft.LoadTestService/locations/quotas"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.LoadTestService/locations/westus/quotas/testQuotaBucket"),
	// 	Properties: &armloadtesting.CheckQuotaAvailabilityResponseProperties{
	// 		AvailabilityStatus: to.Ptr("The requested quota is currently unavailable. Please request for different quota, or upgrade subscription offer type and try again later."),
	// 		IsAvailable: to.Ptr(false),
	// 	},
	// }
}
