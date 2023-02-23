//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3751704f5318f1175875c94b66af769db917f2d3/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/AzureStorage/SoftDeletedContainers_List.json
func ExampleDeletedProtectionContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewDeletedProtectionContainersClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("testRg", "testVault", &armrecoveryservicesbackup.DeletedProtectionContainersClientListOptions{Filter: to.Ptr("backupManagementType eq 'AzureWorkload'")})
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
		// page.ProtectionContainerResourceList = armrecoveryservicesbackup.ProtectionContainerResourceList{
		// 	Value: []*armrecoveryservicesbackup.ProtectionContainerResource{
		// 		{
		// 			Name: to.Ptr("StorageContainer;Storage;testrg;suchandrtestsa125"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers"),
		// 			ID: to.Ptr("/subscriptions/172424a4-d65f-421e-a8de-197d98aabeba/resourcegroups/testrg/providers/microsoft.recoveryservices/vaults/suchandr-test-vault-wcus/backupFabrics/Azure/protectionContainers/StorageContainer;Storage;testrg;suchandrtestsa125"),
		// 			Properties: &armrecoveryservicesbackup.AzureStorageContainer{
		// 				BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 				ContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 				FriendlyName: to.Ptr("suchandrtestsa125"),
		// 				HealthStatus: to.Ptr("Healthy"),
		// 				RegistrationStatus: to.Ptr("SoftDeleted"),
		// 				ProtectedItemCount: to.Ptr[int64](2),
		// 				SourceResourceID: to.Ptr("/subscriptions/172424a4-d65f-421e-a8de-197d98aabeba/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/suchandrtestsa125"),
		// 			},
		// 	}},
		// }
	}
}
