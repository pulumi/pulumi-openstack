// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 load balancer vip resource within OpenStack.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_vip_v1.html.markdown.
type Vip struct {
	s *pulumi.ResourceState
}

// NewVip registers a new resource with the given unique name, arguments, and options.
func NewVip(ctx *pulumi.Context,
	name string, args *VipArgs, opts ...pulumi.ResourceOpt) (*Vip, error) {
	if args == nil || args.PoolId == nil {
		return nil, errors.New("missing required argument 'PoolId'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["address"] = nil
		inputs["adminStateUp"] = nil
		inputs["connLimit"] = nil
		inputs["description"] = nil
		inputs["floatingIp"] = nil
		inputs["name"] = nil
		inputs["persistence"] = nil
		inputs["poolId"] = nil
		inputs["port"] = nil
		inputs["protocol"] = nil
		inputs["region"] = nil
		inputs["subnetId"] = nil
		inputs["tenantId"] = nil
	} else {
		inputs["address"] = args.Address
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["connLimit"] = args.ConnLimit
		inputs["description"] = args.Description
		inputs["floatingIp"] = args.FloatingIp
		inputs["name"] = args.Name
		inputs["persistence"] = args.Persistence
		inputs["poolId"] = args.PoolId
		inputs["port"] = args.Port
		inputs["protocol"] = args.Protocol
		inputs["region"] = args.Region
		inputs["subnetId"] = args.SubnetId
		inputs["tenantId"] = args.TenantId
	}
	inputs["portId"] = nil
	s, err := ctx.RegisterResource("openstack:loadbalancer/vip:Vip", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vip{s: s}, nil
}

// GetVip gets an existing Vip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVip(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VipState, opts ...pulumi.ResourceOpt) (*Vip, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["address"] = state.Address
		inputs["adminStateUp"] = state.AdminStateUp
		inputs["connLimit"] = state.ConnLimit
		inputs["description"] = state.Description
		inputs["floatingIp"] = state.FloatingIp
		inputs["name"] = state.Name
		inputs["persistence"] = state.Persistence
		inputs["poolId"] = state.PoolId
		inputs["port"] = state.Port
		inputs["portId"] = state.PortId
		inputs["protocol"] = state.Protocol
		inputs["region"] = state.Region
		inputs["subnetId"] = state.SubnetId
		inputs["tenantId"] = state.TenantId
	}
	s, err := ctx.ReadResource("openstack:loadbalancer/vip:Vip", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vip{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Vip) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Vip) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The IP address of the vip. Changing this creates a new
// vip.
func (r *Vip) Address() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["address"])
}

// The administrative state of the vip.
// Acceptable values are "true" and "false". Changing this value updates the
// state of the existing vip.
func (r *Vip) AdminStateUp() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["adminStateUp"])
}

// The maximum number of connections allowed for the
// vip. Default is -1, meaning no limit. Changing this updates the connLimit
// of the existing vip.
func (r *Vip) ConnLimit() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["connLimit"])
}

// Human-readable description for the vip. Changing
// this updates the description of the existing vip.
func (r *Vip) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// A *Networking* Floating IP that will be associated
// with the vip. The Floating IP must be provisioned already.
func (r *Vip) FloatingIp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["floatingIp"])
}

// The name of the vip. Changing this updates the name of
// the existing vip.
func (r *Vip) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Omit this field to prevent session persistence.
// The persistence object structure is documented below. Changing this updates
// the persistence of the existing vip.
func (r *Vip) Persistence() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["persistence"])
}

// The ID of the pool with which the vip is associated.
// Changing this updates the poolId of the existing vip.
func (r *Vip) PoolId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["poolId"])
}

// The port on which to listen for client traffic. Changing
// this creates a new vip.
func (r *Vip) Port() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["port"])
}

// Port UUID for this VIP at associated floating IP (if any).
func (r *Vip) PortId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["portId"])
}

// The protocol - can be either 'TCP, 'HTTP', or
// HTTPS'. Changing this creates a new vip.
func (r *Vip) Protocol() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["protocol"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a VIP. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// VIP.
func (r *Vip) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The network on which to allocate the vip's address. A
// tenant can only create vips on networks authorized by policy (e.g. networks
// that belong to them or networks that are shared). Changing this creates a
// new vip.
func (r *Vip) SubnetId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["subnetId"])
}

// The owner of the vip. Required if admin wants to
// create a vip member for another tenant. Changing this creates a new vip.
func (r *Vip) TenantId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["tenantId"])
}

// Input properties used for looking up and filtering Vip resources.
type VipState struct {
	// The IP address of the vip. Changing this creates a new
	// vip.
	Address interface{}
	// The administrative state of the vip.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing vip.
	AdminStateUp interface{}
	// The maximum number of connections allowed for the
	// vip. Default is -1, meaning no limit. Changing this updates the connLimit
	// of the existing vip.
	ConnLimit interface{}
	// Human-readable description for the vip. Changing
	// this updates the description of the existing vip.
	Description interface{}
	// A *Networking* Floating IP that will be associated
	// with the vip. The Floating IP must be provisioned already.
	FloatingIp interface{}
	// The name of the vip. Changing this updates the name of
	// the existing vip.
	Name interface{}
	// Omit this field to prevent session persistence.
	// The persistence object structure is documented below. Changing this updates
	// the persistence of the existing vip.
	Persistence interface{}
	// The ID of the pool with which the vip is associated.
	// Changing this updates the poolId of the existing vip.
	PoolId interface{}
	// The port on which to listen for client traffic. Changing
	// this creates a new vip.
	Port interface{}
	// Port UUID for this VIP at associated floating IP (if any).
	PortId interface{}
	// The protocol - can be either 'TCP, 'HTTP', or
	// HTTPS'. Changing this creates a new vip.
	Protocol interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VIP. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// VIP.
	Region interface{}
	// The network on which to allocate the vip's address. A
	// tenant can only create vips on networks authorized by policy (e.g. networks
	// that belong to them or networks that are shared). Changing this creates a
	// new vip.
	SubnetId interface{}
	// The owner of the vip. Required if admin wants to
	// create a vip member for another tenant. Changing this creates a new vip.
	TenantId interface{}
}

// The set of arguments for constructing a Vip resource.
type VipArgs struct {
	// The IP address of the vip. Changing this creates a new
	// vip.
	Address interface{}
	// The administrative state of the vip.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing vip.
	AdminStateUp interface{}
	// The maximum number of connections allowed for the
	// vip. Default is -1, meaning no limit. Changing this updates the connLimit
	// of the existing vip.
	ConnLimit interface{}
	// Human-readable description for the vip. Changing
	// this updates the description of the existing vip.
	Description interface{}
	// A *Networking* Floating IP that will be associated
	// with the vip. The Floating IP must be provisioned already.
	FloatingIp interface{}
	// The name of the vip. Changing this updates the name of
	// the existing vip.
	Name interface{}
	// Omit this field to prevent session persistence.
	// The persistence object structure is documented below. Changing this updates
	// the persistence of the existing vip.
	Persistence interface{}
	// The ID of the pool with which the vip is associated.
	// Changing this updates the poolId of the existing vip.
	PoolId interface{}
	// The port on which to listen for client traffic. Changing
	// this creates a new vip.
	Port interface{}
	// The protocol - can be either 'TCP, 'HTTP', or
	// HTTPS'. Changing this creates a new vip.
	Protocol interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VIP. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// VIP.
	Region interface{}
	// The network on which to allocate the vip's address. A
	// tenant can only create vips on networks authorized by policy (e.g. networks
	// that belong to them or networks that are shared). Changing this creates a
	// new vip.
	SubnetId interface{}
	// The owner of the vip. Required if admin wants to
	// create a vip member for another tenant. Changing this creates a new vip.
	TenantId interface{}
}
