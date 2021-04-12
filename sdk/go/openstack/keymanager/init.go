// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "openstack:keymanager/containerV1:ContainerV1":
		r = &ContainerV1{}
	case "openstack:keymanager/orderV1:OrderV1":
		r = &OrderV1{}
	case "openstack:keymanager/secretV1:SecretV1":
		r = &SecretV1{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := openstack.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"openstack",
		"keymanager/containerV1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"keymanager/orderV1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"keymanager/secretV1",
		&module{version},
	)
}
