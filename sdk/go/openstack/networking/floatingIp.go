// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Floating IPs can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:networking/floatingIp:FloatingIp floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1
// ```
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
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A list of external subnet IDs to try over each to
	// allocate a floating IP address. If a subnet ID in a list has exhausted
	// floating IP pool, the next subnet ID will be tried. This argument is used only
	// during the resource creation. Conflicts with a `subnetId` argument.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Pool == nil {
		return nil, errors.New("invalid value for required argument 'Pool'")
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
	// A list of external subnet IDs to try over each to
	// allocate a floating IP address. If a subnet ID in a list has exhausted
	// floating IP pool, the next subnet ID will be tried. This argument is used only
	// during the resource creation. Conflicts with a `subnetId` argument.
	SubnetIds []string `pulumi:"subnetIds"`
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
	// A list of external subnet IDs to try over each to
	// allocate a floating IP address. If a subnet ID in a list has exhausted
	// floating IP pool, the next subnet ID will be tried. This argument is used only
	// during the resource creation. Conflicts with a `subnetId` argument.
	SubnetIds pulumi.StringArrayInput
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
	// A list of external subnet IDs to try over each to
	// allocate a floating IP address. If a subnet ID in a list has exhausted
	// floating IP pool, the next subnet ID will be tried. This argument is used only
	// during the resource creation. Conflicts with a `subnetId` argument.
	SubnetIds []string `pulumi:"subnetIds"`
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
	// A list of external subnet IDs to try over each to
	// allocate a floating IP address. If a subnet ID in a list has exhausted
	// floating IP pool, the next subnet ID will be tried. This argument is used only
	// during the resource creation. Conflicts with a `subnetId` argument.
	SubnetIds pulumi.StringArrayInput
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

type FloatingIpInput interface {
	pulumi.Input

	ToFloatingIpOutput() FloatingIpOutput
	ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput
}

func (*FloatingIp) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIp)(nil)).Elem()
}

func (i *FloatingIp) ToFloatingIpOutput() FloatingIpOutput {
	return i.ToFloatingIpOutputWithContext(context.Background())
}

func (i *FloatingIp) ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpOutput)
}

// FloatingIpArrayInput is an input type that accepts FloatingIpArray and FloatingIpArrayOutput values.
// You can construct a concrete instance of `FloatingIpArrayInput` via:
//
//          FloatingIpArray{ FloatingIpArgs{...} }
type FloatingIpArrayInput interface {
	pulumi.Input

	ToFloatingIpArrayOutput() FloatingIpArrayOutput
	ToFloatingIpArrayOutputWithContext(context.Context) FloatingIpArrayOutput
}

type FloatingIpArray []FloatingIpInput

func (FloatingIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FloatingIp)(nil)).Elem()
}

func (i FloatingIpArray) ToFloatingIpArrayOutput() FloatingIpArrayOutput {
	return i.ToFloatingIpArrayOutputWithContext(context.Background())
}

func (i FloatingIpArray) ToFloatingIpArrayOutputWithContext(ctx context.Context) FloatingIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpArrayOutput)
}

// FloatingIpMapInput is an input type that accepts FloatingIpMap and FloatingIpMapOutput values.
// You can construct a concrete instance of `FloatingIpMapInput` via:
//
//          FloatingIpMap{ "key": FloatingIpArgs{...} }
type FloatingIpMapInput interface {
	pulumi.Input

	ToFloatingIpMapOutput() FloatingIpMapOutput
	ToFloatingIpMapOutputWithContext(context.Context) FloatingIpMapOutput
}

type FloatingIpMap map[string]FloatingIpInput

func (FloatingIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FloatingIp)(nil)).Elem()
}

func (i FloatingIpMap) ToFloatingIpMapOutput() FloatingIpMapOutput {
	return i.ToFloatingIpMapOutputWithContext(context.Background())
}

func (i FloatingIpMap) ToFloatingIpMapOutputWithContext(ctx context.Context) FloatingIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpMapOutput)
}

type FloatingIpOutput struct{ *pulumi.OutputState }

func (FloatingIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIp)(nil)).Elem()
}

func (o FloatingIpOutput) ToFloatingIpOutput() FloatingIpOutput {
	return o
}

func (o FloatingIpOutput) ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput {
	return o
}

type FloatingIpArrayOutput struct{ *pulumi.OutputState }

func (FloatingIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FloatingIp)(nil)).Elem()
}

func (o FloatingIpArrayOutput) ToFloatingIpArrayOutput() FloatingIpArrayOutput {
	return o
}

func (o FloatingIpArrayOutput) ToFloatingIpArrayOutputWithContext(ctx context.Context) FloatingIpArrayOutput {
	return o
}

func (o FloatingIpArrayOutput) Index(i pulumi.IntInput) FloatingIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FloatingIp {
		return vs[0].([]*FloatingIp)[vs[1].(int)]
	}).(FloatingIpOutput)
}

type FloatingIpMapOutput struct{ *pulumi.OutputState }

func (FloatingIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FloatingIp)(nil)).Elem()
}

func (o FloatingIpMapOutput) ToFloatingIpMapOutput() FloatingIpMapOutput {
	return o
}

func (o FloatingIpMapOutput) ToFloatingIpMapOutputWithContext(ctx context.Context) FloatingIpMapOutput {
	return o
}

func (o FloatingIpMapOutput) MapIndex(k pulumi.StringInput) FloatingIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FloatingIp {
		return vs[0].(map[string]*FloatingIp)[vs[1].(string)]
	}).(FloatingIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpInput)(nil)).Elem(), &FloatingIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpArrayInput)(nil)).Elem(), FloatingIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpMapInput)(nil)).Elem(), FloatingIpMap{})
	pulumi.RegisterOutputType(FloatingIpOutput{})
	pulumi.RegisterOutputType(FloatingIpArrayOutput{})
	pulumi.RegisterOutputType(FloatingIpMapOutput{})
}
