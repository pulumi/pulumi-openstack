// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack"
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
	case "openstack:firewall/policy:Policy":
		r = &Policy{}
	case "openstack:firewall/rule:Rule":
		r = &Rule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := openstack.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/firewall",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"firewall/rule",
		&module{version},
	)
}
