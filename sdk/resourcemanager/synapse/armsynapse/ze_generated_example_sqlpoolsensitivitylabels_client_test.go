//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsSensitivityLabelsWithSourceCurrent.json
func ExampleSQLPoolSensitivityLabelsClient_NewListCurrentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListCurrentPager("myRG",
		"myServer",
		"myDatabase",
		&armsynapse.SQLPoolSensitivityLabelsClientListCurrentOptions{Filter: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsCurrentUpdate.json
func ExampleSQLPoolSensitivityLabelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Update(ctx,
		"myRG",
		"myWorkspace",
		"mySqlPool",
		armsynapse.SensitivityLabelUpdateList{
			Operations: []*armsynapse.SensitivityLabelUpdate{
				{
					Properties: &armsynapse.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column1"),
						Op:     to.Ptr(armsynapse.SensitivityLabelUpdateKindSet),
						SensitivityLabel: &armsynapse.SensitivityLabel{
							Properties: &armsynapse.SensitivityLabelProperties{
								InformationType:   to.Ptr("Financial"),
								InformationTypeID: to.Ptr("1D3652D6-422C-4115-82F1-65DAEBC665C8"),
								LabelID:           to.Ptr("3A477B16-9423-432B-AA97-6069B481CEC3"),
								LabelName:         to.Ptr("Highly Confidential"),
								Rank:              to.Ptr(armsynapse.SensitivityLabelRankLow),
							},
						},
						Table: to.Ptr("table1"),
					},
				},
				{
					Properties: &armsynapse.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("column2"),
						Op:     to.Ptr(armsynapse.SensitivityLabelUpdateKindSet),
						SensitivityLabel: &armsynapse.SensitivityLabel{
							Properties: &armsynapse.SensitivityLabelProperties{
								InformationType:   to.Ptr("PhoneNumber"),
								InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
								LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
								LabelName:         to.Ptr("PII"),
								Rank:              to.Ptr(armsynapse.SensitivityLabelRankCritical),
							},
						},
						Table: to.Ptr("table2"),
					},
				},
				{
					Properties: &armsynapse.SensitivityLabelUpdateProperties{
						Schema: to.Ptr("dbo"),
						Column: to.Ptr("Column3"),
						Op:     to.Ptr(armsynapse.SensitivityLabelUpdateKindRemove),
						Table:  to.Ptr("Table1"),
					},
				}},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolSensitivityLabelsWithSourceRecommended.json
func ExampleSQLPoolSensitivityLabelsClient_NewListRecommendedPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListRecommendedPager("myRG",
		"myServer",
		"myDatabase",
		&armsynapse.SQLPoolSensitivityLabelsClientListRecommendedOptions{IncludeDisabledRecommendations: nil,
			SkipToken: nil,
			Filter:    nil,
		})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolColumnSensitivityLabelWithAllParameters.json
func ExampleSQLPoolSensitivityLabelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"myRG",
		"myServer",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		armsynapse.SensitivityLabel{
			Properties: &armsynapse.SensitivityLabelProperties{
				InformationType:   to.Ptr("PhoneNumber"),
				InformationTypeID: to.Ptr("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
				LabelID:           to.Ptr("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
				LabelName:         to.Ptr("PII"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolColumnSensitivityLabel.json
func ExampleSQLPoolSensitivityLabelsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"myRG",
		"myServer",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolColumnSensitivityLabelGet.json
func ExampleSQLPoolSensitivityLabelsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myRG",
		"myServer",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		armsynapse.SensitivityLabelSourceCurrent,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RecommendedColumnSensitivityLabelEnable.json
func ExampleSQLPoolSensitivityLabelsClient_EnableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.EnableRecommendation(ctx,
		"myRG",
		"myServer",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RecommendedColumnSensitivityLabelDisable.json
func ExampleSQLPoolSensitivityLabelsClient_DisableRecommendation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewSQLPoolSensitivityLabelsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.DisableRecommendation(ctx,
		"myRG",
		"myServer",
		"myDatabase",
		"dbo",
		"myTable",
		"myColumn",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
