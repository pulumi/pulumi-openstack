// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestExamples(t *testing.T) {
	region := os.Getenv("OS_AUTH_URL")
	if region == "" {
		t.Skipf("Skipping test due to missing OS_AUTH_URL environment variable")
	}
	cwd, err := os.Getwd()
	if !assert.NoError(t, err, "expected a valid working directory: %v", err) {
		return
	}

	base := integration.ProgramTestOptions{
		Tracing: "https://tracing.pulumi-engineering.com/collector/api/v1/spans",
	}

	examples := []integration.ProgramTestOptions{
		base.With(integration.ProgramTestOptions{
			Dir: path.Join(cwd, "webserver"),
			Dependencies: []string{
				"@pulumi/openstack",
			},
			// One change is known to occur during refresh of the resources in this example:
			// `~  openstack:compute:Instance test updated changes: + blockDevices,personalities,schedulerHints``
			ExpectRefreshChanges: true,
		}),
	}

	for _, ex := range examples {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, &example)
		})
	}
}
