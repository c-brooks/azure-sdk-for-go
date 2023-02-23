//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_ListByRuleSet.json
func ExampleRulesClient_NewListByRuleSetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByRuleSetPager("RG",
		"profile1",
		"ruleSet1",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Get.json
func ExampleRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"RG",
		"profile1",
		"ruleSet1",
		"rule1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Create.json
func ExampleRulesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"RG",
		"profile1",
		"ruleSet1",
		"rule1",
		armcdn.Rule{
			Properties: &armcdn.RuleProperties{
				Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
					&armcdn.DeliveryRuleResponseHeaderAction{
						Name: to.Ptr(armcdn.DeliveryRuleActionModifyResponseHeader),
						Parameters: &armcdn.HeaderActionParameters{
							HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
							HeaderName:   to.Ptr("X-CDN"),
							TypeName:     to.Ptr(armcdn.HeaderActionParametersTypeNameDeliveryRuleHeaderActionParameters),
							Value:        to.Ptr("MSFT"),
						},
					}},
				Conditions: []armcdn.DeliveryRuleConditionClassification{
					&armcdn.DeliveryRuleRequestMethodCondition{
						Name: to.Ptr(armcdn.MatchVariableRequestMethod),
						Parameters: &armcdn.RequestMethodMatchConditionParameters{
							MatchValues: []*armcdn.RequestMethodMatchConditionParametersMatchValuesItem{
								to.Ptr(armcdn.RequestMethodMatchConditionParametersMatchValuesItemGET)},
							NegateCondition: to.Ptr(false),
							Operator:        to.Ptr(armcdn.RequestMethodOperatorEqual),
							TypeName:        to.Ptr(armcdn.RequestMethodMatchConditionParametersTypeNameDeliveryRuleRequestMethodConditionParameters),
						},
					}},
				Order: to.Ptr[int32](1),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Update.json
func ExampleRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"RG",
		"profile1",
		"ruleSet1",
		"rule1",
		armcdn.RuleUpdateParameters{
			Properties: &armcdn.RuleUpdatePropertiesParameters{
				Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
					&armcdn.DeliveryRuleResponseHeaderAction{
						Name: to.Ptr(armcdn.DeliveryRuleActionModifyResponseHeader),
						Parameters: &armcdn.HeaderActionParameters{
							HeaderAction: to.Ptr(armcdn.HeaderActionOverwrite),
							HeaderName:   to.Ptr("X-CDN"),
							TypeName:     to.Ptr(armcdn.HeaderActionParametersTypeNameDeliveryRuleHeaderActionParameters),
							Value:        to.Ptr("MSFT"),
						},
					}},
				Order: to.Ptr[int32](1),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Delete.json
func ExampleRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcdn.NewRulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"RG",
		"profile1",
		"ruleSet1",
		"rule1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
