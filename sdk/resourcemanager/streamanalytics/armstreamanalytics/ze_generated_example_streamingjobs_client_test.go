//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics_test

import (
	"context"
	"log"

	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Create_CompleteJob.json
func ExampleStreamingJobsClient_BeginCreateOrReplace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrReplace(ctx,
		"sjrg3276",
		"sj7804",
		armstreamanalytics.StreamingJob{
			Location: to.Ptr("West US"),
			Tags: map[string]*string{
				"key1":      to.Ptr("value1"),
				"key3":      to.Ptr("value3"),
				"randomKey": to.Ptr("randomValue"),
			},
			Properties: &armstreamanalytics.StreamingJobProperties{
				CompatibilityLevel:                 to.Ptr(armstreamanalytics.CompatibilityLevelOne0),
				DataLocale:                         to.Ptr("en-US"),
				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](5),
				EventsOutOfOrderMaxDelayInSeconds:  to.Ptr[int32](0),
				EventsOutOfOrderPolicy:             to.Ptr(armstreamanalytics.EventsOutOfOrderPolicyDrop),
				Functions:                          []*armstreamanalytics.Function{},
				Inputs: []*armstreamanalytics.Input{
					{
						Name: to.Ptr("inputtest"),
						Properties: &armstreamanalytics.StreamInputProperties{
							Type: to.Ptr("Stream"),
							Serialization: &armstreamanalytics.JSONSerialization{
								Type: to.Ptr(armstreamanalytics.EventSerializationTypeJSON),
								Properties: &armstreamanalytics.JSONSerializationProperties{
									Encoding: to.Ptr(armstreamanalytics.EncodingUTF8),
								},
							},
							Datasource: &armstreamanalytics.BlobStreamInputDataSource{
								Type: to.Ptr("Microsoft.Storage/Blob"),
								Properties: &armstreamanalytics.BlobStreamInputDataSourceProperties{
									Container:   to.Ptr("containerName"),
									PathPattern: to.Ptr(""),
									StorageAccounts: []*armstreamanalytics.StorageAccount{
										{
											AccountKey:  to.Ptr("yourAccountKey=="),
											AccountName: to.Ptr("yourAccountName"),
										}},
								},
							},
						},
					}},
				OutputErrorPolicy: to.Ptr(armstreamanalytics.OutputErrorPolicyDrop),
				Outputs: []*armstreamanalytics.Output{
					{
						Name: to.Ptr("outputtest"),
						Properties: &armstreamanalytics.OutputProperties{
							Datasource: &armstreamanalytics.AzureSQLDatabaseOutputDataSource{
								Type: to.Ptr("Microsoft.Sql/Server/Database"),
								Properties: &armstreamanalytics.AzureSQLDatabaseOutputDataSourceProperties{
									Database: to.Ptr("databaseName"),
									Password: to.Ptr("userPassword"),
									Server:   to.Ptr("serverName"),
									Table:    to.Ptr("tableName"),
									User:     to.Ptr("<user>"),
								},
							},
						},
					}},
				SKU: &armstreamanalytics.SKU{
					Name: to.Ptr(armstreamanalytics.SKUNameStandard),
				},
				Transformation: &armstreamanalytics.Transformation{
					Name: to.Ptr("transformationtest"),
					Properties: &armstreamanalytics.TransformationProperties{
						Query:          to.Ptr("Select Id, Name from inputtest"),
						StreamingUnits: to.Ptr[int32](1),
					},
				},
			},
		},
		&armstreamanalytics.StreamingJobsClientBeginCreateOrReplaceOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Update.json
func ExampleStreamingJobsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"sjrg6936",
		"sj59",
		armstreamanalytics.StreamingJob{
			Properties: &armstreamanalytics.StreamingJobProperties{
				EventsLateArrivalMaxDelayInSeconds: to.Ptr[int32](13),
				EventsOutOfOrderMaxDelayInSeconds:  to.Ptr[int32](21),
			},
		},
		&armstreamanalytics.StreamingJobsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Delete.json
func ExampleStreamingJobsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"sjrg6936",
		"sj59",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Get_NoExpand.json
func ExampleStreamingJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"sjrg6936",
		"sj59",
		&armstreamanalytics.StreamingJobsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_List_ByResourceGroup_NoExpand.json
func ExampleStreamingJobsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("sjrg6936",
		&armstreamanalytics.StreamingJobsClientListByResourceGroupOptions{Expand: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_List_BySubscription_NoExpand.json
func ExampleStreamingJobsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armstreamanalytics.StreamingJobsClientListOptions{Expand: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Start_CustomTime.json
func ExampleStreamingJobsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStart(ctx,
		"sjrg6936",
		"sj59",
		&armstreamanalytics.StreamingJobsClientBeginStartOptions{StartJobParameters: &armstreamanalytics.StartStreamingJobParameters{
			OutputStartMode: to.Ptr(armstreamanalytics.OutputStartModeCustomTime),
			OutputStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2012-12-12T12:12:12Z"); return t }()),
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Stop.json
func ExampleStreamingJobsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStop(ctx,
		"sjrg6936",
		"sj59",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Scale.json
func ExampleStreamingJobsClient_BeginScale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewStreamingJobsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginScale(ctx,
		"sjrg6936",
		"sj59",
		&armstreamanalytics.StreamingJobsClientBeginScaleOptions{ScaleJobParameters: &armstreamanalytics.ScaleStreamingJobParameters{
			StreamingUnits: to.Ptr[int32](36),
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
