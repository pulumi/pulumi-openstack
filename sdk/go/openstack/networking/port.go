// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V2 port resource within OpenStack.
//
// ## Example Usage
// ### Simple port
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network1, err := networking.NewNetwork(ctx, "network1", &networking.NetworkArgs{
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewPort(ctx, "port1", &networking.PortArgs{
//				NetworkId:    network1.ID(),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Port with physical binding information
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network1, err := networking.NewNetwork(ctx, "network1", &networking.NetworkArgs{
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewPort(ctx, "port1", &networking.PortArgs{
//				NetworkId:    network1.ID(),
//				DeviceId:     pulumi.String("cdf70fcf-c161-4f24-9c70-96b3f5a54b71"),
//				DeviceOwner:  pulumi.String("baremetal:none"),
//				AdminStateUp: pulumi.Bool(true),
//				Binding: &networking.PortBindingArgs{
//					HostId:   pulumi.String("b080b9cf-46e0-4ce8-ad47-0fd4accc872b"),
//					VnicType: pulumi.String("baremetal"),
//					Profile: pulumi.String(`{
//	  "local_link_information": [
//	    {
//	      "switch_info": "info1",
//	      "port_id": "Ethernet3/4",
//	      "switch_id": "12:34:56:78:9A:BC"
//	    },
//	    {
//	      "switch_info": "info2",
//	      "port_id": "Ethernet3/4",
//	      "switch_id": "12:34:56:78:9A:BD"
//	    }
//	  ],
//	  "vlan_type": "allowed"
//	}
//
// `),
//
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Notes
//
// ### Ports and Instances
//
// There are some notes to consider when connecting Instances to networks using
// Ports. Please see the `compute.Instance` documentation for further
// documentation.
//
// ## Import
//
// Ports can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:networking/port:Port port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
//
// ```
type Port struct {
	pulumi.CustomResourceState

	// Administrative up/down status for the port
	// (must be `true` or `false` if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp pulumi.BoolOutput `pulumi:"adminStateUp"`
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps pulumi.StringArrayOutput `pulumi:"allFixedIps"`
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds pulumi.StringArrayOutput `pulumi:"allSecurityGroupIds"`
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs PortAllowedAddressPairArrayOutput `pulumi:"allowedAddressPairs"`
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding PortBindingOutput `pulumi:"binding"`
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing port.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner pulumi.StringOutput `pulumi:"deviceOwner"`
	// The list of maps representing port DNS assignments.
	DnsAssignments pulumi.MapArrayOutput `pulumi:"dnsAssignments"`
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions PortExtraDhcpOptionArrayOutput `pulumi:"extraDhcpOptions"`
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps PortFixedIpArrayOutput `pulumi:"fixedIps"`
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp pulumi.BoolPtrOutput `pulumi:"noFixedIp"`
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups pulumi.BoolPtrOutput `pulumi:"noSecurityGroups"`
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of `true`. Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled pulumi.BoolOutput `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringOutput `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A set of string tags for the port.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewPort registers a new resource with the given unique name, arguments, and options.
func NewPort(ctx *pulumi.Context,
	name string, args *PortArgs, opts ...pulumi.ResourceOption) (*Port, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Port
	err := ctx.RegisterResource("openstack:networking/port:Port", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPort gets an existing Port resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortState, opts ...pulumi.ResourceOption) (*Port, error) {
	var resource Port
	err := ctx.ReadResource("openstack:networking/port:Port", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Port resources.
type portState struct {
	// Administrative up/down status for the port
	// (must be `true` or `false` if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps []string `pulumi:"allFixedIps"`
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds []string `pulumi:"allSecurityGroupIds"`
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []PortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding *PortBinding `pulumi:"binding"`
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing port.
	Description *string `pulumi:"description"`
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// The list of maps representing port DNS assignments.
	DnsAssignments []map[string]interface{} `pulumi:"dnsAssignments"`
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName *string `pulumi:"dnsName"`
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions []PortExtraDhcpOption `pulumi:"extraDhcpOptions"`
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps []PortFixedIp `pulumi:"fixedIps"`
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress *string `pulumi:"macAddress"`
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name *string `pulumi:"name"`
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId *string `pulumi:"networkId"`
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp *bool `pulumi:"noFixedIp"`
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups *bool `pulumi:"noSecurityGroups"`
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of `true`. Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region *string `pulumi:"region"`
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A set of string tags for the port.
	Tags []string `pulumi:"tags"`
	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type PortState struct {
	// Administrative up/down status for the port
	// (must be `true` or `false` if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp pulumi.BoolPtrInput
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps pulumi.StringArrayInput
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds pulumi.StringArrayInput
	// The collection of tags assigned on the port, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs PortAllowedAddressPairArrayInput
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding PortBindingPtrInput
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing port.
	Description pulumi.StringPtrInput
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId pulumi.StringPtrInput
	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner pulumi.StringPtrInput
	// The list of maps representing port DNS assignments.
	DnsAssignments pulumi.MapArrayInput
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName pulumi.StringPtrInput
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions PortExtraDhcpOptionArrayInput
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps PortFixedIpArrayInput
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress pulumi.StringPtrInput
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name pulumi.StringPtrInput
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId pulumi.StringPtrInput
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp pulumi.BoolPtrInput
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups pulumi.BoolPtrInput
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of `true`. Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region pulumi.StringPtrInput
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
	// A set of string tags for the port.
	Tags pulumi.StringArrayInput
	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (PortState) ElementType() reflect.Type {
	return reflect.TypeOf((*portState)(nil)).Elem()
}

type portArgs struct {
	// Administrative up/down status for the port
	// (must be `true` or `false` if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []PortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding *PortBinding `pulumi:"binding"`
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing port.
	Description *string `pulumi:"description"`
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName *string `pulumi:"dnsName"`
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions []PortExtraDhcpOption `pulumi:"extraDhcpOptions"`
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps []PortFixedIp `pulumi:"fixedIps"`
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress *string `pulumi:"macAddress"`
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name *string `pulumi:"name"`
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId string `pulumi:"networkId"`
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp *bool `pulumi:"noFixedIp"`
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups *bool `pulumi:"noSecurityGroups"`
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of `true`. Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled *bool `pulumi:"portSecurityEnabled"`
	// Reference to the associated QoS policy.
	QosPolicyId *string `pulumi:"qosPolicyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region *string `pulumi:"region"`
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A set of string tags for the port.
	Tags []string `pulumi:"tags"`
	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Port resource.
type PortArgs struct {
	// Administrative up/down status for the port
	// (must be `true` or `false` if provided). Changing this updates the
	// `adminStateUp` of an existing port.
	AdminStateUp pulumi.BoolPtrInput
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs PortAllowedAddressPairArrayInput
	// The port binding allows to specify binding information
	// for the port. The structure is described below.
	Binding PortBindingPtrInput
	// Human-readable description of the port. Changing
	// this updates the `description` of an existing port.
	Description pulumi.StringPtrInput
	// The ID of the device attached to the port. Changing this
	// creates a new port.
	DeviceId pulumi.StringPtrInput
	// The device owner of the port. Changing this creates
	// a new port.
	DeviceOwner pulumi.StringPtrInput
	// The port DNS name. Available, when Neutron DNS extension
	// is enabled.
	DnsName pulumi.StringPtrInput
	// An extra DHCP option that needs to be configured
	// on the port. The structure is described below. Can be specified multiple
	// times.
	ExtraDhcpOptions PortExtraDhcpOptionArrayInput
	// An array of desired IPs for
	// this port. The structure is described below.
	FixedIps PortFixedIpArrayInput
	// Specify a specific MAC address for the port. Changing
	// this creates a new port.
	MacAddress pulumi.StringPtrInput
	// A unique name for the port. Changing this
	// updates the `name` of an existing port.
	Name pulumi.StringPtrInput
	// The ID of the network to attach the port to. Changing
	// this creates a new port.
	NetworkId pulumi.StringInput
	// Create a port with no fixed
	// IP address. This will also remove any fixed IPs previously set on a port. `true`
	// is the only valid value for this argument.
	NoFixedIp pulumi.BoolPtrInput
	// If set to
	// `true`, then no security groups are applied to the port. If set to `false` and
	// no `securityGroupIds` are specified, then the port will yield to the default
	// behavior of the Networking service, which is to usually apply the "default"
	// security group.
	NoSecurityGroups pulumi.BoolPtrInput
	// Whether to explicitly enable or disable
	// port security on the port. Port Security is usually enabled by default, so
	// omitting argument will usually result in a value of `true`. Setting this
	// explicitly to `false` will disable port security. In order to disable port
	// security, the port must not have any security groups. Valid values are `true`
	// and `false`.
	PortSecurityEnabled pulumi.BoolPtrInput
	// Reference to the associated QoS policy.
	QosPolicyId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// port.
	Region pulumi.StringPtrInput
	// A list
	// of security group IDs to apply to the port. The security groups must be
	// specified by ID and not name (as opposed to how they are configured with
	// the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
	// A set of string tags for the port.
	Tags pulumi.StringArrayInput
	// The owner of the port. Required if admin wants
	// to create a port for another tenant. Changing this creates a new port.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (PortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portArgs)(nil)).Elem()
}

type PortInput interface {
	pulumi.Input

	ToPortOutput() PortOutput
	ToPortOutputWithContext(ctx context.Context) PortOutput
}

func (*Port) ElementType() reflect.Type {
	return reflect.TypeOf((**Port)(nil)).Elem()
}

func (i *Port) ToPortOutput() PortOutput {
	return i.ToPortOutputWithContext(context.Background())
}

func (i *Port) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortOutput)
}

func (i *Port) ToOutput(ctx context.Context) pulumix.Output[*Port] {
	return pulumix.Output[*Port]{
		OutputState: i.ToPortOutputWithContext(ctx).OutputState,
	}
}

// PortArrayInput is an input type that accepts PortArray and PortArrayOutput values.
// You can construct a concrete instance of `PortArrayInput` via:
//
//	PortArray{ PortArgs{...} }
type PortArrayInput interface {
	pulumi.Input

	ToPortArrayOutput() PortArrayOutput
	ToPortArrayOutputWithContext(context.Context) PortArrayOutput
}

type PortArray []PortInput

func (PortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Port)(nil)).Elem()
}

func (i PortArray) ToPortArrayOutput() PortArrayOutput {
	return i.ToPortArrayOutputWithContext(context.Background())
}

func (i PortArray) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortArrayOutput)
}

func (i PortArray) ToOutput(ctx context.Context) pulumix.Output[[]*Port] {
	return pulumix.Output[[]*Port]{
		OutputState: i.ToPortArrayOutputWithContext(ctx).OutputState,
	}
}

// PortMapInput is an input type that accepts PortMap and PortMapOutput values.
// You can construct a concrete instance of `PortMapInput` via:
//
//	PortMap{ "key": PortArgs{...} }
type PortMapInput interface {
	pulumi.Input

	ToPortMapOutput() PortMapOutput
	ToPortMapOutputWithContext(context.Context) PortMapOutput
}

type PortMap map[string]PortInput

func (PortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Port)(nil)).Elem()
}

func (i PortMap) ToPortMapOutput() PortMapOutput {
	return i.ToPortMapOutputWithContext(context.Background())
}

func (i PortMap) ToPortMapOutputWithContext(ctx context.Context) PortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortMapOutput)
}

func (i PortMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Port] {
	return pulumix.Output[map[string]*Port]{
		OutputState: i.ToPortMapOutputWithContext(ctx).OutputState,
	}
}

type PortOutput struct{ *pulumi.OutputState }

func (PortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Port)(nil)).Elem()
}

func (o PortOutput) ToPortOutput() PortOutput {
	return o
}

func (o PortOutput) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return o
}

func (o PortOutput) ToOutput(ctx context.Context) pulumix.Output[*Port] {
	return pulumix.Output[*Port]{
		OutputState: o.OutputState,
	}
}

// Administrative up/down status for the port
// (must be `true` or `false` if provided). Changing this updates the
// `adminStateUp` of an existing port.
func (o PortOutput) AdminStateUp() pulumi.BoolOutput {
	return o.ApplyT(func(v *Port) pulumi.BoolOutput { return v.AdminStateUp }).(pulumi.BoolOutput)
}

// The collection of Fixed IP addresses on the port in the
// order returned by the Network v2 API.
func (o PortOutput) AllFixedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.AllFixedIps }).(pulumi.StringArrayOutput)
}

// The collection of Security Group IDs on the port
// which have been explicitly and implicitly added.
func (o PortOutput) AllSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.AllSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The collection of tags assigned on the port, which have been
// explicitly and implicitly added.
func (o PortOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.AllTags }).(pulumi.StringArrayOutput)
}

// An IP/MAC Address pair of additional IP
// addresses that can be active on this port. The structure is described
// below.
func (o PortOutput) AllowedAddressPairs() PortAllowedAddressPairArrayOutput {
	return o.ApplyT(func(v *Port) PortAllowedAddressPairArrayOutput { return v.AllowedAddressPairs }).(PortAllowedAddressPairArrayOutput)
}

// The port binding allows to specify binding information
// for the port. The structure is described below.
func (o PortOutput) Binding() PortBindingOutput {
	return o.ApplyT(func(v *Port) PortBindingOutput { return v.Binding }).(PortBindingOutput)
}

// Human-readable description of the port. Changing
// this updates the `description` of an existing port.
func (o PortOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Port) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the device attached to the port. Changing this
// creates a new port.
func (o PortOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// The device owner of the port. Changing this creates
// a new port.
func (o PortOutput) DeviceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.DeviceOwner }).(pulumi.StringOutput)
}

// The list of maps representing port DNS assignments.
func (o PortOutput) DnsAssignments() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.MapArrayOutput { return v.DnsAssignments }).(pulumi.MapArrayOutput)
}

// The port DNS name. Available, when Neutron DNS extension
// is enabled.
func (o PortOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// An extra DHCP option that needs to be configured
// on the port. The structure is described below. Can be specified multiple
// times.
func (o PortOutput) ExtraDhcpOptions() PortExtraDhcpOptionArrayOutput {
	return o.ApplyT(func(v *Port) PortExtraDhcpOptionArrayOutput { return v.ExtraDhcpOptions }).(PortExtraDhcpOptionArrayOutput)
}

// An array of desired IPs for
// this port. The structure is described below.
func (o PortOutput) FixedIps() PortFixedIpArrayOutput {
	return o.ApplyT(func(v *Port) PortFixedIpArrayOutput { return v.FixedIps }).(PortFixedIpArrayOutput)
}

// Specify a specific MAC address for the port. Changing
// this creates a new port.
func (o PortOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// A unique name for the port. Changing this
// updates the `name` of an existing port.
func (o PortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the network to attach the port to. Changing
// this creates a new port.
func (o PortOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Create a port with no fixed
// IP address. This will also remove any fixed IPs previously set on a port. `true`
// is the only valid value for this argument.
func (o PortOutput) NoFixedIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Port) pulumi.BoolPtrOutput { return v.NoFixedIp }).(pulumi.BoolPtrOutput)
}

// If set to
// `true`, then no security groups are applied to the port. If set to `false` and
// no `securityGroupIds` are specified, then the port will yield to the default
// behavior of the Networking service, which is to usually apply the "default"
// security group.
func (o PortOutput) NoSecurityGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Port) pulumi.BoolPtrOutput { return v.NoSecurityGroups }).(pulumi.BoolPtrOutput)
}

// Whether to explicitly enable or disable
// port security on the port. Port Security is usually enabled by default, so
// omitting argument will usually result in a value of `true`. Setting this
// explicitly to `false` will disable port security. In order to disable port
// security, the port must not have any security groups. Valid values are `true`
// and `false`.
func (o PortOutput) PortSecurityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Port) pulumi.BoolOutput { return v.PortSecurityEnabled }).(pulumi.BoolOutput)
}

// Reference to the associated QoS policy.
func (o PortOutput) QosPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.QosPolicyId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a port. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// port.
func (o PortOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list
// of security group IDs to apply to the port. The security groups must be
// specified by ID and not name (as opposed to how they are configured with
// the Compute Instance).
func (o PortOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A set of string tags for the port.
func (o PortOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Port) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The owner of the port. Required if admin wants
// to create a port for another tenant. Changing this creates a new port.
func (o PortOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Port) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Map of additional options.
func (o PortOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Port) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type PortArrayOutput struct{ *pulumi.OutputState }

func (PortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Port)(nil)).Elem()
}

func (o PortArrayOutput) ToPortArrayOutput() PortArrayOutput {
	return o
}

func (o PortArrayOutput) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return o
}

func (o PortArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Port] {
	return pulumix.Output[[]*Port]{
		OutputState: o.OutputState,
	}
}

func (o PortArrayOutput) Index(i pulumi.IntInput) PortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Port {
		return vs[0].([]*Port)[vs[1].(int)]
	}).(PortOutput)
}

type PortMapOutput struct{ *pulumi.OutputState }

func (PortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Port)(nil)).Elem()
}

func (o PortMapOutput) ToPortMapOutput() PortMapOutput {
	return o
}

func (o PortMapOutput) ToPortMapOutputWithContext(ctx context.Context) PortMapOutput {
	return o
}

func (o PortMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Port] {
	return pulumix.Output[map[string]*Port]{
		OutputState: o.OutputState,
	}
}

func (o PortMapOutput) MapIndex(k pulumi.StringInput) PortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Port {
		return vs[0].(map[string]*Port)[vs[1].(string)]
	}).(PortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortInput)(nil)).Elem(), &Port{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortArrayInput)(nil)).Elem(), PortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortMapInput)(nil)).Elem(), PortMap{})
	pulumi.RegisterOutputType(PortOutput{})
	pulumi.RegisterOutputType(PortArrayOutput{})
	pulumi.RegisterOutputType(PortMapOutput{})
}
