// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccWebserver(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "webserver"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	checkAuthUrl(t)
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/openstack",
		},
	})

	return baseJS
}
