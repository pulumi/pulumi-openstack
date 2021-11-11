// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

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
	case "openstack:loadbalancer/l7PolicyV2:L7PolicyV2":
		r = &L7PolicyV2{}
	case "openstack:loadbalancer/l7RuleV2:L7RuleV2":
		r = &L7RuleV2{}
	case "openstack:loadbalancer/listener:Listener":
		r = &Listener{}
	case "openstack:loadbalancer/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "openstack:loadbalancer/member:Member":
		r = &Member{}
	case "openstack:loadbalancer/memberV1:MemberV1":
		r = &MemberV1{}
	case "openstack:loadbalancer/members:Members":
		r = &Members{}
	case "openstack:loadbalancer/monitor:Monitor":
		r = &Monitor{}
	case "openstack:loadbalancer/monitorV1:MonitorV1":
		r = &MonitorV1{}
	case "openstack:loadbalancer/pool:Pool":
		r = &Pool{}
	case "openstack:loadbalancer/poolV1:PoolV1":
		r = &PoolV1{}
	case "openstack:loadbalancer/quota:Quota":
		r = &Quota{}
	case "openstack:loadbalancer/vip:Vip":
		r = &Vip{}
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
		"loadbalancer/l7PolicyV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/l7RuleV2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/member",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/memberV1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/members",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/monitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/monitorV1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/pool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/poolV1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/quota",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"openstack",
		"loadbalancer/vip",
		&module{version},
	)
}
