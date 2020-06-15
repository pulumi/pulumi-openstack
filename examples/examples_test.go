// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func checkAuthUrl(t *testing.T) {
	authUrl := os.Getenv("OS_AUTH_URL")
	if authUrl == "" {
		t.Skipf("Skipping test due to missing OS_AUTH_URL environment variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Tracing: "https://tracing.pulumi-engineering.com/collector/api/v1/spans",
		// One change is known to occur during refresh of the resources in this example:
		// `~  openstack:compute:Instance test updated changes: + blockDevices,personalities,schedulerHints``
		ExpectRefreshChanges: true,
	}
}
