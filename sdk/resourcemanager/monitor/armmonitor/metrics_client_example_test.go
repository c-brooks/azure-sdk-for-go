//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetric.json
func ExampleMetricsClient_List_getMetricForData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "subscriptions/b324c52b-4073-4807-93af-e07d289c093e/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/larryshoebox/blobServices/default", &armmonitor.MetricsClientListOptions{Timespan: to.Ptr("2017-04-14T02:20:00Z/2017-04-14T04:20:00Z"),
		Interval:        to.Ptr("PT1M"),
		Metricnames:     nil,
		Aggregation:     to.Ptr("Average,count"),
		Top:             to.Ptr[int32](3),
		Orderby:         to.Ptr("Average asc"),
		Filter:          to.Ptr("BlobType eq '*'"),
		ResultType:      nil,
		Metricnamespace: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricMetadata.json
func ExampleMetricsClient_List_getMetricForMetadata() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "subscriptions/b324c52b-4073-4807-93af-e07d289c093e/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/larryshoebox/blobServices/default", &armmonitor.MetricsClientListOptions{Timespan: to.Ptr("2017-04-14T02:20:00Z/2017-04-14T04:20:00Z"),
		Interval:        to.Ptr("PT1M"),
		Metricnames:     nil,
		Aggregation:     to.Ptr("Average,count"),
		Top:             to.Ptr[int32](3),
		Orderby:         to.Ptr("Average asc"),
		Filter:          to.Ptr("BlobType eq '*'"),
		ResultType:      nil,
		Metricnamespace: to.Ptr("Microsoft.Storage/storageAccounts/blobServices"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-01-01/examples/GetMetricError.json
func ExampleMetricsClient_List_getMetricWithError() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewMetricsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx, "subscriptions/ac41e21f-afd6-4a79-8070-f01eba278f97/resourceGroups/todking/providers/Microsoft.DocumentDb/databaseAccounts/tk-cosmos-mongo", &armmonitor.MetricsClientListOptions{Timespan: to.Ptr("2021-06-07T21:51:00Z/2021-06-08T01:51:00Z"),
		Interval:        to.Ptr("FULL"),
		Metricnames:     to.Ptr("MongoRequestsCount,MongoRequests"),
		Aggregation:     to.Ptr("average"),
		Top:             nil,
		Orderby:         nil,
		Filter:          nil,
		ResultType:      nil,
		Metricnamespace: to.Ptr("microsoft.documentdb/databaseaccounts"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
