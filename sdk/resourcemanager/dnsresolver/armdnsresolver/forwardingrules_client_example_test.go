//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Put.json
func ExampleForwardingRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", armdnsresolver.ForwardingRule{
		Properties: &armdnsresolver.ForwardingRuleProperties{
			DomainName:          to.Ptr("contoso.com."),
			ForwardingRuleState: to.Ptr(armdnsresolver.ForwardingRuleStateEnabled),
			Metadata: map[string]*string{
				"additionalProp1": to.Ptr("value1"),
			},
			TargetDNSServers: []*armdnsresolver.TargetDNSServer{
				{
					IPAddress: to.Ptr("10.0.0.1"),
					Port:      to.Ptr[int32](53),
				},
				{
					IPAddress: to.Ptr("10.0.0.2"),
					Port:      to.Ptr[int32](53),
				}},
		},
	}, &armdnsresolver.ForwardingRulesClientCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Patch.json
func ExampleForwardingRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", armdnsresolver.ForwardingRulePatch{
		Properties: &armdnsresolver.ForwardingRulePatchProperties{
			ForwardingRuleState: to.Ptr(armdnsresolver.ForwardingRuleStateDisabled),
			Metadata: map[string]*string{
				"additionalProp2": to.Ptr("value2"),
			},
		},
	}, &armdnsresolver.ForwardingRulesClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Delete.json
func ExampleForwardingRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", &armdnsresolver.ForwardingRulesClientDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_Get.json
func ExampleForwardingRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "sampleResourceGroup", "sampleDnsForwardingRuleset", "sampleForwardingRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_List.json
func ExampleForwardingRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewForwardingRulesClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("sampleResourceGroup", "sampleDnsForwardingRuleset", &armdnsresolver.ForwardingRulesClientListOptions{Top: nil})
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
