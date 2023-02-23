//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbotservice_test

import (
	"context"
	"log"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DirectlineRegenerateKeys.json
func ExampleDirectLineClient_RegenerateKeys_regenerateKeysForDirectLineChannelSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewDirectLineClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RegenerateKeys(ctx, "OneResourceGroupName", "samplebotname", armbotservice.RegenerateKeysChannelNameDirectLineChannel, armbotservice.SiteInfo{
		Key:      to.Ptr(armbotservice.KeyKey1),
		SiteName: to.Ptr("testSiteName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/WebChatRegenerateKeys.json
func ExampleDirectLineClient_RegenerateKeys_regenerateKeysForWebChatChannelSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewDirectLineClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RegenerateKeys(ctx, "OneResourceGroupName", "samplebotname", armbotservice.RegenerateKeysChannelNameWebChatChannel, armbotservice.SiteInfo{
		Key:      to.Ptr(armbotservice.KeyKey1),
		SiteName: to.Ptr("testSiteName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
