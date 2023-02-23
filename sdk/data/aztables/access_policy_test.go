// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztables

import (
	"strconv"
	"testing"
	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/internal/recording"
	"github.com/stretchr/testify/require"
)

func TestSetEmptyAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	_, err := client.SetAccessPolicy(ctx, nil)
	require.NoError(t, err)
}

func TestSetAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	expiration := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	permission := "r"
	id := "1"

	signedIdentifiers := make([]*SignedIdentifier, 0)

	signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
		AccessPolicy: &AccessPolicy{
			Expiry:     &expiration,
			Start:      &start,
			Permission: &permission,
		},
		ID: &id,
	})

	param := SetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NoError(t, err)

	recording.Sleep(60 * time.Second)

	resp, err := client.GetAccessPolicy(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, 1, len(resp.SignedIdentifiers))

	// set nil access policy
	_, err = client.SetAccessPolicy(ctx, nil)
	require.NoError(t, err)

	recording.Sleep(60 * time.Second)

	resp, err = client.GetAccessPolicy(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, 0, len(resp.SignedIdentifiers))
}

func TestSetMultipleAccessPolicies(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	id := "empty"

	signedIdentifiers := make([]*SignedIdentifier, 0)
	signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
		ID: &id,
	})

	permission2 := "r"
	id2 := "partial"

	signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
		ID: &id2,
		AccessPolicy: &AccessPolicy{
			Permission: &permission2,
		},
	})

	id3 := "full"
	permission3 := "r"
	start := time.Date(2021, 6, 8, 2, 10, 9, 0, time.UTC)
	expiry := time.Date(2021, 6, 8, 2, 10, 9, 0, time.UTC)

	signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
		ID: &id3,
		AccessPolicy: &AccessPolicy{
			Start:      &start,
			Expiry:     &expiry,
			Permission: &permission3,
		},
	})

	param := SetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NoError(t, err)

	// Make a Get to assert two access policies
	resp, err := client.GetAccessPolicy(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, len(resp.SignedIdentifiers), 3)
}

func TestSetTooManyAccessPolicies(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	expiration := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	permission := "r"
	id := "1"
	signedIdentifiers := make([]*SignedIdentifier, 0)

	for i := 0; i < 6; i++ {
		expiration = time.Date(2024+i, 1, 1, 0, 0, 0, 0, time.UTC)
		id = strconv.Itoa(i)

		signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
			AccessPolicy: &AccessPolicy{
				Expiry:     &expiration,
				Start:      &start,
				Permission: &permission,
			},
			ID: &id,
		})

	}

	param := SetAccessPolicyOptions{TableACL: signedIdentifiers}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.Error(t, err)
	require.Contains(t, err.Error(), errTooManyAccessPoliciesError.Error())
}

func TestSetNullAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	id := "null"

	signedIdentifiers := make([]*SignedIdentifier, 0)
	signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
		ID: &id,
	})

	param := SetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.NoError(t, err)

	resp, err := client.GetAccessPolicy(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, len(resp.SignedIdentifiers), 1)
}

func TestSetInvalidAccessPolicy(t *testing.T) {
	client, delete := initClientTest(t, "storage", true)
	defer delete()

	signedIdentifiers := make([]*SignedIdentifier, 0)
	signedIdentifiers = append(signedIdentifiers, &SignedIdentifier{
		AccessPolicy: &AccessPolicy{
			Expiry: to.Ptr(time.Date(2042, 1, 1, 1, 1, 1, 1, time.UTC)),
		},
	})

	param := SetAccessPolicyOptions{
		TableACL: signedIdentifiers,
	}

	_, err := client.SetAccessPolicy(ctx, &param)
	require.Error(t, err)
	var httpErr *azcore.ResponseError
	require.ErrorAs(t, err, &httpErr)
	require.Equal(t, "InvalidXmlDocument", httpErr.ErrorCode)
	require.Contains(t, PossibleTableErrorCodeValues(), TableErrorCode(httpErr.ErrorCode))
}
