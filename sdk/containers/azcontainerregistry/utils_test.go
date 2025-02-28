//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcontainerregistry

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

// FakeCredential is an empty credential for testing.
type FakeCredential struct {
}

// GetToken provide a fake access token.
func (c *FakeCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{Token: "Sanitized", ExpiresOn: time.Now().Add(time.Hour * 24).UTC()}, nil
}

// getCredAndClientOptions will create a credential and a client options for test application.
// The client options will initialize the transport for recording client add recording policy to the pipeline.
// In the record mode, the credential will be a DefaultAzureCredential which combines several common credentials.
// In the playback mode, the credential will be a fake credential which will bypass truly authorization.
func getCredAndClientOptions(t *testing.T) (azcore.TokenCredential, azcore.ClientOptions) {
	transport, err := recording.NewRecordingHTTPClient(t, nil)
	require.NoError(t, err)

	options := azcore.ClientOptions{
		Transport: transport,
	}

	var cred azcore.TokenCredential
	if recording.GetRecordMode() != recording.PlaybackMode {
		cred, err = azidentity.NewDefaultAzureCredential(nil)
		require.NoError(t, err)
	} else {
		cred = &FakeCredential{}
	}

	return cred, options
}

// startRecording starts the recording.
func startRecording(t *testing.T) {
	err := recording.Start(t, "sdk/containers/azcontainerregistry/testdata", nil)
	require.NoError(t, err)
	t.Cleanup(func() {
		err := recording.Stop(t, nil)
		require.NoError(t, err)
	})
}

func TestMain(m *testing.M) {
	err := recording.ResetProxy(nil)
	if err != nil {
		panic(err)
	}
	if recording.GetRecordMode() == recording.RecordingMode {
		defer func() {
			err := recording.ResetProxy(nil)
			if err != nil {
				panic(err)
			}
		}()
		// sanitizer for any uuid string, e.g., subscriptionID
		err = recording.AddGeneralRegexSanitizer("00000000-0000-0000-0000-000000000000", `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`, nil)
		if err != nil {
			panic(err)
		}
		// sanitizer for authentication
		err = recording.AddBodyRegexSanitizer("access_token=Sanitized&", "access_token=[^&]+&", nil)
		if err != nil {
			panic(err)
		}
		err = recording.AddBodyRegexSanitizer("\"refresh_token\":\".eyJqdGkiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJzdWIiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJuYmYiOjQ2NzA0MTEyMTIsImV4cCI6NDY3MDQyMjkxMiwiaWF0Ijo0NjcwNDExMjEyLCJpc3MiOiJBenVyZSBDb250YWluZXIgUmVnaXN0cnkiLCJhdWQiOiJhemFjcmxpdmV0ZXN0LmF6dXJlY3IuaW8iLCJ2ZXJzaW9uIjoiMS4wIiwicmlkIjoiMDAwMCIsImdyYW50X3R5cGUiOiJyZWZyZXNoX3Rva2VuIiwiYXBwaWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJwZXJtaXNzaW9ucyI6eyJBY3Rpb25zIjpbInJlYWQiLCJ3cml0ZSIsImRlbGV0ZSIsImRlbGV0ZWQvcmVhZCIsImRlbGV0ZWQvcmVzdG9yZS9hY3Rpb24iXSwiTm90QWN0aW9ucyI6bnVsbH0sInJvbGVzIjpbXX0=.\"", "\"refresh_token\":\".+\"", nil)
		if err != nil {
			panic(err)
		}
		err = recording.AddBodyRegexSanitizer("refresh_token=.eyJqdGkiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJzdWIiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJuYmYiOjQ2NzA0MTEyMTIsImV4cCI6NDY3MDQyMjkxMiwiaWF0Ijo0NjcwNDExMjEyLCJpc3MiOiJBenVyZSBDb250YWluZXIgUmVnaXN0cnkiLCJhdWQiOiJhemFjcmxpdmV0ZXN0LmF6dXJlY3IuaW8iLCJ2ZXJzaW9uIjoiMS4wIiwicmlkIjoiMDAwMCIsImdyYW50X3R5cGUiOiJyZWZyZXNoX3Rva2VuIiwiYXBwaWQiOiIwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAiLCJwZXJtaXNzaW9ucyI6eyJBY3Rpb25zIjpbInJlYWQiLCJ3cml0ZSIsImRlbGV0ZSIsImRlbGV0ZWQvcmVhZCIsImRlbGV0ZWQvcmVzdG9yZS9hY3Rpb24iXSwiTm90QWN0aW9ucyI6bnVsbH0sInJvbGVzIjpbXX0%3D.&", "refresh_token=[^&]+&", nil)
		if err != nil {
			panic(err)
		}
	}
	code := m.Run()
	os.Exit(code)
}
