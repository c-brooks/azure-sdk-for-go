// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs_test

import (
	"context"
	"fmt"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

var consumerClient *azeventhubs.ConsumerClient
var err error

func ExampleNewConsumerClient() {
	defaultAzureCred, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		panic(err)
	}

	consumerClient, err = azeventhubs.NewConsumerClient("<ex: myeventhubnamespace.servicebus.windows.net>", "eventhub-name", azeventhubs.DefaultConsumerGroup, defaultAzureCred, nil)

	if err != nil {
		panic(err)
	}
}

func ExampleNewConsumerClientFromConnectionString() {
	// if the connection string contains an EntityPath
	//
	connectionString := "Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;EntityPath=<entity path>"
	consumerClient, err = azeventhubs.NewConsumerClientFromConnectionString(connectionString, "", azeventhubs.DefaultConsumerGroup, nil)

	// or

	// if the connection string does not contain an EntityPath
	connectionString = "Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>"
	consumerClient, err = azeventhubs.NewConsumerClientFromConnectionString(connectionString, "eventhub-name", azeventhubs.DefaultConsumerGroup, nil)

	if err != nil {
		panic(err)
	}
}

func ExampleConsumerClient_NewPartitionClient_receiveEvents() {
	const partitionID = "0"

	partitionClient, err := consumerClient.NewPartitionClient(partitionID, nil)

	if err != nil {
		panic(err)
	}

	defer partitionClient.Close(context.TODO())

	events, err := partitionClient.ReceiveEvents(context.TODO(), 100, nil)

	if err != nil {
		panic(err)
	}

	for _, evt := range events {
		fmt.Printf("Body: %s\n", string(evt.Body))
	}
}

func ExampleConsumerClient_GetEventHubProperties() {
	eventHubProps, err := consumerClient.GetEventHubProperties(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	for _, partitionID := range eventHubProps.PartitionIDs {
		fmt.Printf("Partition ID: %s\n", partitionID)
	}
}

func ExampleConsumerClient_GetPartitionProperties() {
	partitionProps, err := consumerClient.GetPartitionProperties(context.TODO(), "partition-id", nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("First sequence number for partition ID %s: %d\n", partitionProps.PartitionID, partitionProps.BeginningSequenceNumber)
	fmt.Printf("Last sequence number for partition ID %s: %d\n", partitionProps.PartitionID, partitionProps.LastEnqueuedSequenceNumber)
}

func ExampleConsumerClient_NewPartitionClient_configuringPrefetch() {
	const partitionID = "0"

	// Prefetching configures the Event Hubs client to continually cache events, up to the configured size
	// in PartitionClientOptions.Prefetch. PartitionClient.ReceiveEvents will read from the cache first,
	// which can speed up throughput in situations where you might normally be forced to request and wait
	// for more events.

	// By default, prefetch is enabled.
	partitionClient, err := consumerClient.NewPartitionClient(partitionID, nil)

	if err != nil {
		panic(err)
	}

	defer partitionClient.Close(context.TODO())

	// You can configure the prefetch buffer size as well. The default is 300.
	partitionClientWithCustomPrefetch, err := consumerClient.NewPartitionClient(partitionID, &azeventhubs.PartitionClientOptions{
		Prefetch: 301,
	})

	if err != nil {
		panic(err)
	}

	defer partitionClientWithCustomPrefetch.Close(context.TODO())

	// And prefetch can be disabled if you prefer to manually control the flow of events. Excess
	// events (that arrive after your ReceiveEvents() call has completed) will still be
	// buffered internally, but they will not be automatically replenished.
	partitionClientWithPrefetchDisabled, err := consumerClient.NewPartitionClient(partitionID, &azeventhubs.PartitionClientOptions{
		Prefetch: -1,
	})

	if err != nil {
		panic(err)
	}

	defer partitionClientWithPrefetchDisabled.Close(context.TODO())

	events, err := partitionClient.ReceiveEvents(context.TODO(), 100, nil)

	if err != nil {
		panic(err)
	}

	for _, evt := range events {
		fmt.Printf("Body: %s\n", string(evt.Body))
	}
}
