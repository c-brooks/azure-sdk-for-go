//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_List.json
func ExampleProjectEnvironmentTypesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("0ac520ee-14c0-480f-b6c9-0a90c58ffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", "ContosoProj", &armdevcenter.ProjectEnvironmentTypesClientListOptions{Top: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_Get.json
func ExampleProjectEnvironmentTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("0ac520ee-14c0-480f-b6c9-0a90c58ffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "ContosoProj", "DevTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_Put.json
func ExampleProjectEnvironmentTypesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("0ac520ee-14c0-480f-b6c9-0a90c58ffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "rg1", "ContosoProj", "DevTest", armdevcenter.ProjectEnvironmentType{
		Identity: &armdevcenter.ManagedServiceIdentity{
			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {},
			},
		},
		Properties: &armdevcenter.ProjectEnvironmentTypeProperties{
			CreatorRoleAssignment: &armdevcenter.ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment{
				Roles: map[string]*armdevcenter.EnvironmentRole{
					"4cbf0b6c-e750-441c-98a7-10da8387e4d6": {},
				},
			},
			DeploymentTargetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
			Status:             to.Ptr(armdevcenter.EnableStatusEnabled),
			UserRoleAssignments: map[string]*armdevcenter.UserRoleAssignmentValue{
				"e45e3m7c-176e-416a-b466-0c5ec8298f8a": {
					Roles: map[string]*armdevcenter.EnvironmentRole{
						"4cbf0b6c-e750-441c-98a7-10da8387e4d6": {},
					},
				},
			},
		},
		Tags: map[string]*string{
			"CostCenter": to.Ptr("RnD"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_Patch.json
func ExampleProjectEnvironmentTypesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("0ac520ee-14c0-480f-b6c9-0a90c58ffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "rg1", "ContosoProj", "DevTest", armdevcenter.ProjectEnvironmentTypeUpdate{
		Identity: &armdevcenter.ManagedServiceIdentity{
			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {},
			},
		},
		Properties: &armdevcenter.ProjectEnvironmentTypeUpdateProperties{
			DeploymentTargetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
			Status:             to.Ptr(armdevcenter.EnableStatusEnabled),
			UserRoleAssignments: map[string]*armdevcenter.UserRoleAssignmentValue{
				"e45e3m7c-176e-416a-b466-0c5ec8298f8a": {
					Roles: map[string]*armdevcenter.EnvironmentRole{
						"4cbf0b6c-e750-441c-98a7-10da8387e4d6": {},
					},
				},
			},
		},
		Tags: map[string]*string{
			"CostCenter": to.Ptr("RnD"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_Delete.json
func ExampleProjectEnvironmentTypesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("0ac520ee-14c0-480f-b6c9-0a90c58ffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "rg1", "ContosoProj", "DevTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
