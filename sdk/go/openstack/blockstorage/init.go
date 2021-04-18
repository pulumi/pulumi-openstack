// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

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
	case "openstack:blockstorage/quoteSetV2:QuoteSetV2":
		r = &QuoteSetV2{}
	case "openstack:blockstorage/quoteSetV3:QuoteSetV3":
		r = &QuoteSetV3{}
	case "openstack:blockstorage/volume:Volume":
		r = &Volume{}
	case "openstack:blockstorage/volumeAttach:VolumeAttach":
		r = &VolumeAttach{}
	case "openstack:blockstorage/volumeAttachV2:VolumeAttachV2":
		r = &VolumeAttachV2{}
	case "openstack:blockstorage/volumeTypeV3:VolumeTypeV3":
		r = &VolumeTypeV3{}
	case "openstack:blockstorage/volumeV1:VolumeV1":
		r = &VolumeV1{}
	case "openstack:blockstorage/volumeV2:VolumeV2":
		r = &VolumeV2{}
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
		"blockstorage/quoteSetV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/quoteSetV3",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/volume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/volumeAttach",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/volumeAttachV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/volumeTypeV3",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/volumeV1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"blockstorage/volumeV2",
		&module{version},
	)
}