//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicyinsights_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryManagementGroupScope.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForManagementGroupPager_queryAtManagementGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForManagementGroupPager("myManagementGroup", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryManagementGroupScopeWithFilterAndTop.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForManagementGroupPager_queryAtManagementGroupScopeUsingQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForManagementGroupPager("myManagementGroup", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment' AND TrackedResourceId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/exampleTrackedResourceName'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QuerySubscriptionScope.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForSubscriptionPager_queryAtSubscriptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("fffedd8f-ffff-fffd-fffd-fffed2f84852", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QuerySubscriptionScopeWithFilterAndTop.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForSubscriptionPager_queryAtSubscriptionScopeUsingQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("fffedd8f-ffff-fffd-fffd-fffed2f84852", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForSubscriptionPager(armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment' AND TrackedResourceId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/exampleTrackedResourceName'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceGroupScope.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForResourceGroupPager_queryAtResourceGroupScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("fffedd8f-ffff-fffd-fffd-fffed2f84852", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForResourceGroupPager("myResourceGroup", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceGroupScopeWithFilterAndTop.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForResourceGroupPager_queryAtResourceGroupScopeUsingQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("fffedd8f-ffff-fffd-fffd-fffed2f84852", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForResourceGroupPager("myResourceGroup", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment' AND TrackedResourceId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource/nestedResourceType/TrackedResource1'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceScope.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForResourcePager_queryAtResourceScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForResourcePager("subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: nil,
		Filter:    nil,
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceScopeWithFilterAndTop.json
func ExamplePolicyTrackedResourcesClient_NewListQueryResultsForResourcePager_queryAtResourceScopeUsingQueryParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicyinsights.NewPolicyTrackedResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListQueryResultsForResourcePager("subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource", armpolicyinsights.PolicyTrackedResourcesResourceTypeDefault, &armpolicyinsights.QueryOptions{Top: to.Ptr[int32](1),
		Filter:    to.Ptr("PolicyAssignmentId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment' AND TrackedResourceId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource/nestedResourceType/TrackedResource1'"),
		OrderBy:   nil,
		Select:    nil,
		From:      nil,
		To:        nil,
		Apply:     nil,
		SkipToken: nil,
		Expand:    nil,
	}, nil)
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
