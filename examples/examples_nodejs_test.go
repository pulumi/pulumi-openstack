// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccWebserver(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "webserver"),
		})

	integration.ProgramTest(t, &test)
}

func TestKeyPair(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:   path.Join(getCwd(t), "keypair"),
			Quick: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				output, ok := stack.Outputs["keypair"]
				require.Truef(t, ok, "Expected a 'keypair' stack output, found none")
				o, ok := output.(map[string]interface{})
				if assert.Truef(t, ok, "Expected Secret 'keypair' stack output (map[string]any), found %T", output) {
					assert.Equalf(t, "1b47061264138c4ac30d75fd1eb44270",
						o["4dabf18193072939515e22adb298388d"],
						"Expected a Secret 'keypair' stack output")
				}
			},
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
