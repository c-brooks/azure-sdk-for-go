//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetListConnectorSubscription_example.json
func ExampleAccountConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/GetConnectorSubscription_example.json
func ExampleAccountConnectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "aws_dev1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate_awsAssumeRoleCreateACloudAccountConnectorForASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "aws_dev2", armsecurity.ConnectorSetting{
		Properties: &armsecurity.ConnectorSettingProperties{
			AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
				AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsAssumeRole),
				AwsAssumeRoleArn:   to.Ptr("arn:aws:iam::81231569658:role/AscConnector"),
				AwsExternalID:      to.Ptr("20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
			},
			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
				AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
				ProxyServer: &armsecurity.ProxyServerProperties{
					IP:   to.Ptr("167.220.197.140"),
					Port: to.Ptr("34"),
				},
				Region:            to.Ptr("West US 2"),
				ResourceGroupName: to.Ptr("AwsConnectorRG"),
				ServicePrincipal: &armsecurity.ServicePrincipalProperties{
					ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
					Secret:        to.Ptr("<secret>"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsCredConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate_awsCredCreateACloudAccountConnectorForASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "aws_dev1", armsecurity.ConnectorSetting{
		Properties: &armsecurity.ConnectorSettingProperties{
			AuthenticationDetails: &armsecurity.AwsCredsAuthenticationDetailsProperties{
				AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsCreds),
				AwsAccessKeyID:     to.Ptr("AKIARPZCNODDNAEQFSOE"),
				AwsSecretAccessKey: to.Ptr("<awsSecretAccessKey>"),
			},
			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
				AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
				ProxyServer: &armsecurity.ProxyServerProperties{
					IP:   to.Ptr("167.220.197.140"),
					Port: to.Ptr("34"),
				},
				Region:            to.Ptr("West US 2"),
				ResourceGroupName: to.Ptr("AwsConnectorRG"),
				ServicePrincipal: &armsecurity.ServicePrincipalProperties{
					ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
					Secret:        to.Ptr("<secret>"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateGcpCredentialsConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate_gcpCredentialsCreateACloudAccountConnectorForASubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "gcp_dev", armsecurity.ConnectorSetting{
		Properties: &armsecurity.ConnectorSettingProperties{
			AuthenticationDetails: &armsecurity.GcpCredentialsDetailsProperties{
				AuthenticationType:      to.Ptr(armsecurity.AuthenticationTypeGcpCredentials),
				Type:                    to.Ptr("service_account"),
				AuthProviderX509CertURL: to.Ptr("https://www.googleapis.com/oauth2/v1/certs"),
				AuthURI:                 to.Ptr("https://accounts.google.com/o/oauth2/auth"),
				ClientEmail:             to.Ptr("asc-135@asc-project-1234.iam.gserviceaccount.com"),
				ClientID:                to.Ptr("105889053725632919854"),
				ClientX509CertURL:       to.Ptr("https://www.googleapis.com/robot/v1/metadata/x509/asc-135%40asc-project-1234.iam.gserviceaccount.com"),
				OrganizationID:          to.Ptr("AscDemoOrg"),
				PrivateKey:              to.Ptr("-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCpxYHcLzcDZ6/Q\nAeQZnQXM5GTb3p09Xsbjo2T2F61b6I7FZiQXBrbw3Zf0CUCkkqTTpD5xifl82yQ6\n89V7SAe8hxI7esAcVDhm/aJMqzVjHLISAU2L3li1sn0jjY2oYtndwN6bRivP8O6t\n9F+W6E0zMlbCxtpZEHLbb6WxlJJrwEQ0MPH2yOCwZUQi6NHksAtEzX2nNKJNyUC7\nQyBVHHMm34H2bmZwsuQp3y2otpcJ9tJnVmYfC3k/w4x2L+DIK7JnQP/C1wQqu2du\nc0w6sydF6RhLoHButrVdYRJTdfK4k03SsSTyMqZ+f7LNnKw3xenzw1VmEpk8mvoQ\nt08tCBOrAgMBAAECggEAByzz6iyMtLYjNjV+QJ7kad6VbL2iA8AHxANZ9xTVHPdd\nYXaJu/dqsA+NpqDlfI8+LDva782XH/HbPCqmMUnAGfXTjXQIvqnIoIHD5F2wKfpC\nhIRNlMXXFgbvRxtqi11yO+80+XcjzuwuCmgzyhsTeEB+bkkdXXpWgHPdmv3emnM6\nMQM9Zgrug0UndPmiUwKOcJSU4PlmlTpHEV4vA6JfA4bvphy9m1jxO5qWeah5yym2\n6FP5BRIDF98kFrDnSXJjajwgLCQ+MypFQXyax6XkxDxuKXbng1bv7eZDjqazIChk\nm0y14X0s0jnWc+AX8vfeSf7d+EsGdVinEwR1aAawEQKBgQDqDB0qxcIQ1oI1Kww8\n9vXefTiuWsf47F+fJ/DIOEbiRfE8IdCgmOABvcqJIoxW/DFMBEdLCcx73Km7pOmd\nKg1ddScnaO8cOj2v/Ub+fAqVrA4ki4ViYP0A7/Nogga3Jr/x3ey5bitrIfFImteS\nCgBHBzZvoQpvO4lB2tKVgo2P9wKBgQC5sgTEq4sasRGSAY6lIoJno0I8w28a/16D\nes60XQeY1ger8uTGwlT02v/u/arDUmRLPClpujXq6gK29KvtRCHy7JkpGbqW2bZs\nPFKKWR7Tk3XPKYyjv94AIi5/xoFeDhS4lpAvy3Z5tQhYS6wqWKvT6yZQ3kM+Hfxs\npHgvu3mU7QKBgQC9/E1k3hj1cBtMK4CIsHPPQljTd4+iacYJPPPAo6YuoVX8WPqw\nksgrwbN59Fh1d8xQh5yTtgWOegYx8uFMGcm1lpbM7+pBQKm4hWGuzGQPMRZd5f/F\nZzOZIi61I+9tlv/yxxIVR+/ozCm/pSneO04UWi9/F/uPZYW6tnWAtfRR6wKBgGsZ\n8MQaCK4JaI/klAhMghgSQnbXZXKVzUZaA3Rln6cX8u7KtgapOOTMlwaZie8Dy1LV\nTTFstAJcm9o3/h1nyYjZy3C4JTUyNpPwqs6enjf7edxVI4eidwFutZD+xcigqHTa\naikW2atSrZB3fMIjyF7+5meH+hKOqvNiXOty3qn1AoGAZuVxYQy5FVq3YZxzr3Aa\nAm0ShoXTF6QYIbsaUiUGoa/NlHcw9V/lj4AqBRbxbaYMD+hz2J/od9cb268eJKY8\n3b6MvaUqdNhNnWodJXLhgtmGEHDKmTppz2JSTx/tVzCfhFdcOC79StZvcKLhtoFQ\n+/3lEw6NCIXzm5E4+dtJG4k=\n-----END PRIVATE KEY-----\n"),
				PrivateKeyID:            to.Ptr("6efg587hra2568as34d22326b044cc20dc2af"),
				ProjectID:               to.Ptr("asc-project-1234"),
				TokenURI:                to.Ptr("https://oauth2.googleapis.com/token"),
			},
			HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
				AutoProvision: to.Ptr(armsecurity.AutoProvisionOff),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/DeleteConnectorSubscription_example.json
func ExampleAccountConnectorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "aws_dev1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
