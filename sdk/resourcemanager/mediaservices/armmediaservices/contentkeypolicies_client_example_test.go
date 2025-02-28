//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-by-lastModified.json
func ExampleContentKeyPoliciesClient_NewListPager_listsContentKeyPoliciesOrderedByLastModified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("contoso", "contosomedia", &armmediaservices.ContentKeyPoliciesClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: to.Ptr("properties/lastModified"),
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-in-date-range.json
func ExampleContentKeyPoliciesClient_NewListPager_listsContentKeyPoliciesWithCreatedAndLastModifiedFilters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("contoso", "contosomedia", &armmediaservices.ContentKeyPoliciesClientListOptions{Filter: to.Ptr("properties/lastModified gt 2016-06-01 and properties/created lt 2013-07-01"),
		Top:     nil,
		Orderby: nil,
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-all.json
func ExampleContentKeyPoliciesClient_NewListPager_listsAllContentKeyPolicies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("contoso", "contosomedia", &armmediaservices.ContentKeyPoliciesClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: nil,
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-get-by-name.json
func ExampleContentKeyPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "contoso", "contosomedia", "PolicyWithMultipleOptions", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-nodrm-token.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithClearKeyOptionAndTokenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithClearKeyOptionAndSwtTokenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ClearKeyOption"),
					Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
						Audience:  to.Ptr("urn:audience"),
						Issuer:    to.Ptr("urn:issuer"),
						PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
							KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
						},
						RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeSwt),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-playready-open.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithPlayReadyOptionAndOpenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithPlayReadyOptionAndOpenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ArmPolicyOptionName"),
					Configuration: &armmediaservices.ContentKeyPolicyPlayReadyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration"),
						Licenses: []*armmediaservices.ContentKeyPolicyPlayReadyLicense{
							{
								AllowTestDevices: to.Ptr(true),
								BeginDate:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-16T18:22:53.46Z"); return t }()),
								ContentKeyLocation: &armmediaservices.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader{
									ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader"),
								},
								ContentType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload),
								LicenseType: to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyLicenseTypePersistent),
								PlayRight: &armmediaservices.ContentKeyPolicyPlayReadyPlayRight{
									AllowPassingVideoContentToUnknownOutput:            to.Ptr(armmediaservices.ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed),
									DigitalVideoOnlyContentRestriction:                 to.Ptr(false),
									ImageConstraintForAnalogComponentVideoRestriction:  to.Ptr(true),
									ImageConstraintForAnalogComputerMonitorRestriction: to.Ptr(false),
									ScmsRestriction: to.Ptr[int32](2),
								},
								SecurityLevel: to.Ptr(armmediaservices.SecurityLevelSL150),
							}},
					},
					Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-widevine-token.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithWidevineOptionAndTokenRestriction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyWithWidevineOptionAndJwtTokenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("widevineoption"),
					Configuration: &armmediaservices.ContentKeyPolicyWidevineConfiguration{
						ODataType:        to.Ptr("#Microsoft.Media.ContentKeyPolicyWidevineConfiguration"),
						WidevineTemplate: to.Ptr("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
						AlternateVerificationKeys: []armmediaservices.ContentKeyPolicyRestrictionTokenKeyClassification{
							&armmediaservices.ContentKeyPolicySymmetricTokenKey{
								ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
								KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
							}},
						Audience: to.Ptr("urn:audience"),
						Issuer:   to.Ptr("urn:issuer"),
						PrimaryVerificationKey: &armmediaservices.ContentKeyPolicyRsaTokenKey{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyRsaTokenKey"),
							Exponent:  []byte("AQAB"),
							Modulus:   []byte("AQAD"),
						},
						RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeJwt),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-multiple-options.json
func ExampleContentKeyPoliciesClient_CreateOrUpdate_createsAContentKeyPolicyWithMultipleOptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "contoso", "contosomedia", "PolicyCreatedWithMultipleOptions", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("ArmPolicyDescription"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ClearKeyOption"),
					Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyTokenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyTokenRestriction"),
						Audience:  to.Ptr("urn:audience"),
						Issuer:    to.Ptr("urn:issuer"),
						PrimaryVerificationKey: &armmediaservices.ContentKeyPolicySymmetricTokenKey{
							ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicySymmetricTokenKey"),
							KeyValue:  []byte("AAAAAAAAAAAAAAAAAAAAAA=="),
						},
						RestrictionTokenType: to.Ptr(armmediaservices.ContentKeyPolicyRestrictionTokenTypeSwt),
					},
				},
				{
					Name: to.Ptr("widevineoption"),
					Configuration: &armmediaservices.ContentKeyPolicyWidevineConfiguration{
						ODataType:        to.Ptr("#Microsoft.Media.ContentKeyPolicyWidevineConfiguration"),
						WidevineTemplate: to.Ptr("{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-delete.json
func ExampleContentKeyPoliciesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "contoso", "contosomedia", "PolicyWithPlayReadyOptionAndOpenRestriction", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-update.json
func ExampleContentKeyPoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "contoso", "contosomedia", "PolicyWithClearKeyOptionAndTokenRestriction", armmediaservices.ContentKeyPolicy{
		Properties: &armmediaservices.ContentKeyPolicyProperties{
			Description: to.Ptr("Updated Policy"),
			Options: []*armmediaservices.ContentKeyPolicyOption{
				{
					Name: to.Ptr("ClearKeyOption"),
					Configuration: &armmediaservices.ContentKeyPolicyClearKeyConfiguration{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration"),
					},
					Restriction: &armmediaservices.ContentKeyPolicyOpenRestriction{
						ODataType: to.Ptr("#Microsoft.Media.ContentKeyPolicyOpenRestriction"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-get-with-secrets.json
func ExampleContentKeyPoliciesClient_GetPolicyPropertiesWithSecrets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewContentKeyPoliciesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetPolicyPropertiesWithSecrets(ctx, "contoso", "contosomedia", "PolicyWithMultipleOptions", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
