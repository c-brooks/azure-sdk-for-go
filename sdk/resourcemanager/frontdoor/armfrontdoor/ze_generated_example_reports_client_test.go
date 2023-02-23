//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetLatencyScorecard.json
func ExampleReportsClient_GetLatencyScorecards() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewReportsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetLatencyScorecards(ctx,
		"MyResourceGroup",
		"MyProfile",
		"MyExperiment",
		armfrontdoor.LatencyScorecardAggregationIntervalDaily,
		&armfrontdoor.ReportsClientGetLatencyScorecardsOptions{EndDateTimeUTC: nil,
			Country: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
func ExampleReportsClient_GetTimeseries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewReportsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetTimeseries(ctx,
		"MyResourceGroup",
		"MyProfile",
		"MyExperiment",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T17:32:28Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T17:32:28Z"); return t }(),
		armfrontdoor.TimeseriesAggregationIntervalHourly,
		armfrontdoor.TimeseriesTypeMeasurementCounts,
		&armfrontdoor.ReportsClientGetTimeseriesOptions{Endpoint: nil,
			Country: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
