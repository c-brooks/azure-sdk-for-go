//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices_test

import (
	"context"
	"log"

	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListDomainServicesBySubscription.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListDomainServicesByResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("TestResourceGroup",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/CreateDomainService.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestResourceGroup",
		"TestDomainService.com",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				DomainName: to.Ptr("TestDomainService.com"),
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            to.Ptr(armdomainservices.NtlmV1Enabled),
					SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
					TLSV1:             to.Ptr(armdomainservices.TLSV1Disabled),
				},
				FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         to.Ptr(armdomainservices.ExternalAccessEnabled),
					Ldaps:                  to.Ptr(armdomainservices.LdapsEnabled),
					PfxCertificate:         to.Ptr("MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w..."),
					PfxCertificatePassword: to.Ptr("<pfxCertificatePassword>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.Ptr("jicha@microsoft.com"),
						to.Ptr("caalmont@microsoft.com")},
					NotifyDcAdmins:     to.Ptr(armdomainservices.NotifyDcAdminsEnabled),
					NotifyGlobalAdmins: to.Ptr(armdomainservices.NotifyGlobalAdminsEnabled),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.Ptr("West US"),
						SubnetID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS"),
					}},
			},
		},
		nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetDomainService.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"TestResourceGroup",
		"TestDomainService.com",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/DeleteDomainService.json
func ExampleClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"TestResourceGroup",
		"TestDomainService.com",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateDomainService.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("1639790a-76a2-4ac4-98d9-8562f5dfcb4d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"TestResourceGroup",
		"TestDomainService.com",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				ConfigDiagnostics: &armdomainservices.ConfigDiagnostics{
					LastExecuted: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2021-05-05T12:00:23Z;"); return t }()),
					ValidatorResults: []*armdomainservices.ConfigDiagnosticsValidatorResult{
						{
							Issues: []*armdomainservices.ConfigDiagnosticsValidatorResultIssue{
								{
									DescriptionParams: []*string{},
									ID:                to.Ptr("AADDS-CFG-DIAG-I20"),
								}},
							ReplicaSetSubnetDisplayName: to.Ptr("West US/aadds-subnet"),
							Status:                      to.Ptr(armdomainservices.StatusWarning),
							ValidatorID:                 to.Ptr("AADDS-CFG-DIAG-V06"),
						}},
				},
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            to.Ptr(armdomainservices.NtlmV1Enabled),
					SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
					TLSV1:             to.Ptr(armdomainservices.TLSV1Disabled),
				},
				FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         to.Ptr(armdomainservices.ExternalAccessEnabled),
					Ldaps:                  to.Ptr(armdomainservices.LdapsEnabled),
					PfxCertificate:         to.Ptr("MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w..."),
					PfxCertificatePassword: to.Ptr("<pfxCertificatePassword>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.Ptr("jicha@microsoft.com"),
						to.Ptr("caalmont@microsoft.com")},
					NotifyDcAdmins:     to.Ptr(armdomainservices.NotifyDcAdminsEnabled),
					NotifyGlobalAdmins: to.Ptr(armdomainservices.NotifyGlobalAdminsEnabled),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.Ptr("West US"),
						SubnetID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS"),
					},
					{
						Location: to.Ptr("East US"),
						SubnetID: to.Ptr("/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetEUS/subnets/TestSubnetEUS"),
					}},
			},
		},
		nil)
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
