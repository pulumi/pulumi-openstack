// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "openstack:firewall/firewall:Firewall":
		r = &Firewall{}
	case "openstack:firewall/groupV2:GroupV2":
		r = &GroupV2{}
	case "openstack:firewall/policy:Policy":
		r = &Policy{}
	case "openstack:firewall/policyV2:PolicyV2":
		r = &PolicyV2{}
	case "openstack:firewall/rule:Rule":
		r = &Rule{}
	case "openstack:firewall/ruleV2:RuleV2":
		r = &RuleV2{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/firewall",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/groupV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/policyV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/rule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/ruleV2",
		&module{version},
	)
}
