//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization_test

import (
	"context"
	"log"

	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PutRoleEligibilityScheduleRequest.json
func ExampleRoleEligibilityScheduleRequestsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleEligibilityScheduleRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			Condition:        to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
			ConditionVersion: to.Ptr("1.0"),
			PrincipalID:      to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
			RequestType:      to.Ptr(armauthorization.RequestTypeAdminAssign),
			RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
			ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: to.Ptr("P365D"),
				},
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleRequestByName.json
func ExampleRoleEligibilityScheduleRequestsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleEligibilityScheduleRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleRequestByScope.json
func ExampleRoleEligibilityScheduleRequestsClient_NewListForScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleEligibilityScheduleRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListForScopePager("providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", &armauthorization.RoleEligibilityScheduleRequestsClientListForScopeOptions{Filter: to.Ptr("assignedTo('A3BB8764-CB92-4276-9D2A-CA1E895E55EA')")})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/CancelRoleEligibilityScheduleRequestByName.json
func ExampleRoleEligibilityScheduleRequestsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleEligibilityScheduleRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Cancel(ctx, "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/ValidateRoleEligibilityScheduleRequestByName.json
func ExampleRoleEligibilityScheduleRequestsClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armauthorization.NewRoleEligibilityScheduleRequestsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Validate(ctx, "subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f", "64caffb6-55c0-4deb-a585-68e948ea1ad6", armauthorization.RoleEligibilityScheduleRequest{
		Properties: &armauthorization.RoleEligibilityScheduleRequestProperties{
			Condition:        to.Ptr("@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'"),
			ConditionVersion: to.Ptr("1.0"),
			PrincipalID:      to.Ptr("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
			RequestType:      to.Ptr(armauthorization.RequestTypeAdminAssign),
			RoleDefinitionID: to.Ptr("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
			ScheduleInfo: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfo{
				Expiration: &armauthorization.RoleEligibilityScheduleRequestPropertiesScheduleInfoExpiration{
					Type:     to.Ptr(armauthorization.TypeAfterDuration),
					Duration: to.Ptr("P365D"),
				},
				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-09-09T21:31:27.91Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
