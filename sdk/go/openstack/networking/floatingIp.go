// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 floating IP resource within OpenStack Neutron (networking)
// that can be used for load balancers.
// These are similar to Nova (compute) floating IP resources,
// but only compute floating IPs can be used with compute instances.
type FloatingIp struct {
	pulumi.CustomResourceState

	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address pulumi.StringOutput `pulumi:"address"`
	// The collection of tags assigned on the floating IP, which have
	// been explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// Human-readable description for the floating IP.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The floating IP DNS domain. Available, when Neutron
	// DNS extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsDomain pulumi.StringOutput `pulumi:"dnsDomain"`
	// The floating IP DNS name. Available, when Neutron DNS
	// extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp pulumi.StringOutput `pulumi:"fixedIp"`
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool pulumi.StringOutput `pulumi:"pool"`
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region pulumi.StringOutput `pulumi:"region"`
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// A set of string tags for the floating IP.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewFloatingIp registers a new resource with the given unique name, arguments, and options.
func NewFloatingIp(ctx *pulumi.Context,
	name string, args *FloatingIpArgs, opts ...pulumi.ResourceOption) (*FloatingIp, error) {
	if args == nil || args.Pool == nil {
		return nil, errors.New("missing required argument 'Pool'")
	}
	if args == nil {
		args = &FloatingIpArgs{}
	}
	var resource FloatingIp
	err := ctx.RegisterResource("openstack:networking/floatingIp:FloatingIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFloatingIp gets an existing FloatingIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FloatingIpState, opts ...pulumi.ResourceOption) (*FloatingIp, error) {
	var resource FloatingIp
	err := ctx.ReadResource("openstack:networking/floatingIp:FloatingIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FloatingIp resources.
type floatingIpState struct {
	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address *string `pulumi:"address"`
	// The collection of tags assigned on the floating IP, which have
	// been explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// Human-readable description for the floating IP.
	Description *string `pulumi:"description"`
	// The floating IP DNS domain. Available, when Neutron
	// DNS extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsDomain *string `pulumi:"dnsDomain"`
	// The floating IP DNS name. Available, when Neutron DNS
	// extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsName *string `pulumi:"dnsName"`
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp *string `pulumi:"fixedIp"`
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool *string `pulumi:"pool"`
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region *string `pulumi:"region"`
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId *string `pulumi:"subnetId"`
	// A set of string tags for the floating IP.
	Tags []string `pulumi:"tags"`
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type FloatingIpState struct {
	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address pulumi.StringPtrInput
	// The collection of tags assigned on the floating IP, which have
	// been explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// Human-readable description for the floating IP.
	Description pulumi.StringPtrInput
	// The floating IP DNS domain. Available, when Neutron
	// DNS extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsDomain pulumi.StringPtrInput
	// The floating IP DNS name. Available, when Neutron DNS
	// extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsName pulumi.StringPtrInput
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp pulumi.StringPtrInput
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool pulumi.StringPtrInput
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region pulumi.StringPtrInput
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId pulumi.StringPtrInput
	// A set of string tags for the floating IP.
	Tags pulumi.StringArrayInput
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (FloatingIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpState)(nil)).Elem()
}

type floatingIpArgs struct {
	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address *string `pulumi:"address"`
	// Human-readable description for the floating IP.
	Description *string `pulumi:"description"`
	// The floating IP DNS domain. Available, when Neutron
	// DNS extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsDomain *string `pulumi:"dnsDomain"`
	// The floating IP DNS name. Available, when Neutron DNS
	// extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsName *string `pulumi:"dnsName"`
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp *string `pulumi:"fixedIp"`
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool string `pulumi:"pool"`
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region *string `pulumi:"region"`
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId *string `pulumi:"subnetId"`
	// A set of string tags for the floating IP.
	Tags []string `pulumi:"tags"`
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a FloatingIp resource.
type FloatingIpArgs struct {
	// The actual/specific floating IP to obtain. By default,
	// non-admin users are not able to specify a floating IP, so you must either be
	// an admin user or have had a custom policy or role applied to your OpenStack
	// user or project.
	Address pulumi.StringPtrInput
	// Human-readable description for the floating IP.
	Description pulumi.StringPtrInput
	// The floating IP DNS domain. Available, when Neutron
	// DNS extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsDomain pulumi.StringPtrInput
	// The floating IP DNS name. Available, when Neutron DNS
	// extension is enabled. The data in this attribute will be published in an
	// external DNS service when Neutron is configured to integrate with such a
	// service. Changing this creates a new floating IP.
	DnsName pulumi.StringPtrInput
	// Fixed IP of the port to associate with this floating IP. Required if
	// the port has multiple fixed IPs.
	FixedIp pulumi.StringPtrInput
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool pulumi.StringInput
	// ID of an existing port with at least one IP address to
	// associate with this floating IP.
	PortId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a floating IP that can be used with
	// another networking resource, such as a load balancer. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// floating IP (which may or may not have a different address).
	Region pulumi.StringPtrInput
	// The subnet ID of the floating IP pool. Specify this if
	// the floating IP network has multiple subnets.
	SubnetId pulumi.StringPtrInput
	// A set of string tags for the floating IP.
	Tags pulumi.StringArrayInput
	// The target tenant ID in which to allocate the floating
	// IP, if you specify this together with a port_id, make sure the target port
	// belongs to the same tenant. Changing this creates a new floating IP (which
	// may or may not have a different address)
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (FloatingIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpArgs)(nil)).Elem()
}
