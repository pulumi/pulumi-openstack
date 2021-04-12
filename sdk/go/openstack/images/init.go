// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

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
	case "openstack:images/image:Image":
		r = &Image{}
	case "openstack:images/imageAccess:ImageAccess":
		r = &ImageAccess{}
	case "openstack:images/imageAccessAccept:ImageAccessAccept":
		r = &ImageAccessAccept{}
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
		"images/image",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"images/imageAccess",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"images/imageAccessAccept",
		&module{version},
	)
}
