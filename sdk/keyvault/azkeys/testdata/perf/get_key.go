// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/to"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azidentity"
	"github.com/c-brooks/azure-sdk-for-go/sdk/internal/perf"
	"github.com/c-brooks/azure-sdk-for-go/sdk/keyvault/azkeys"
)

type getKeyTestOptions struct{}

var getKeyTestOpts getKeyTestOptions = getKeyTestOptions{}

type getKeyTest struct {
	perf.PerfTestOptions
	keyName string
	client  *azkeys.Client
}

// newGetKeyTest is called once per process
func newGetKeyTest(ctx context.Context, options perf.PerfTestOptions) (perf.GlobalPerfTest, error) {
	d := &getKeyTest{
		PerfTestOptions: options,
		keyName:         "livekvtestgetkeyperfkey",
	}

	vaultURL, ok := os.LookupEnv("AZURE_KEYVAULT_URL")
	if !ok {
		return nil, fmt.Errorf("the environment variable 'AZURE_KEYVAULT_URL' could not be found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}

	client, err := azkeys.NewClient(vaultURL, cred, &azkeys.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: options.Transporter,
		},
	})

	_, err = client.CreateRSAKey(ctx, d.keyName, &azkeys.CreateRSAKeyOptions{Size: to.Ptr(int32(2048))})
	if err != nil {
		return nil, err
	}

	d.client = client
	return d, nil
}

func (gct *getKeyTest) GlobalCleanup(ctx context.Context) error {
	poller, err := gct.client.BeginDeleteKey(ctx, gct.keyName, nil)
	if err != nil {
		return err
	}

	_, err = poller.PollUntilDone(ctx, 500*time.Millisecond)
	if err != nil {
		return err
	}

	_, err = gct.client.PurgeDeletedKey(ctx, gct.keyName, nil)
	return err
}

type getKeyPerfTest struct {
	client  *azkeys.Client
	keyName string
}

// NewPerfTest is called once per goroutine
func (gct *getKeyTest) NewPerfTest(ctx context.Context, options *perf.PerfTestOptions) (perf.PerfTest, error) {
	return &getKeyPerfTest{
		client:  gct.client,
		keyName: gct.keyName,
	}, nil
}

func (gcpt *getKeyPerfTest) Run(ctx context.Context) error {
	_, err := gcpt.client.GetKey(ctx, gcpt.keyName, nil)
	return err
}

func (*getKeyPerfTest) Cleanup(ctx context.Context) error {
	return nil
}
