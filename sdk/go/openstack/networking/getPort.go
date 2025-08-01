// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack port.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.LookupPort(ctx, &networking.LookupPortArgs{
//				Name: pulumi.StringRef("port_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPort(ctx *pulumi.Context, args *LookupPortArgs, opts ...pulumi.InvokeOption) (*LookupPortResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPortResult
	err := ctx.Invoke("openstack:networking/getPort:getPort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPort.
type LookupPortArgs struct {
	// The administrative state of the port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description of the port.
	Description *string `pulumi:"description"`
	// The ID of the device the port belongs to.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// The port DNS name to filter. Available, when Neutron
	// DNS extension is enabled.
	DnsName *string `pulumi:"dnsName"`
	// The port IP address filter.
	FixedIp *string `pulumi:"fixedIp"`
	// The MAC address of the port.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the port.
	Name *string `pulumi:"name"`
	// The ID of the network the port belongs to.
	NetworkId *string `pulumi:"networkId"`
	// The ID of the port.
	PortId *string `pulumi:"portId"`
	// The owner of the port.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve port ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The list of port security group IDs to filter.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The status of the port.
	Status *string `pulumi:"status"`
	// The list of port tags to filter.
	Tags     []string `pulumi:"tags"`
	TenantId *string  `pulumi:"tenantId"`
}

// A collection of values returned by getPort.
type LookupPortResult struct {
	// See Argument Reference above.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps []string `pulumi:"allFixedIps"`
	// The set of security group IDs applied on the port.
	AllSecurityGroupIds []string `pulumi:"allSecurityGroupIds"`
	// The set of string tags applied on the port.
	AllTags []string `pulumi:"allTags"`
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs []GetPortAllowedAddressPair `pulumi:"allowedAddressPairs"`
	// The port binding information. The structure is described below.
	Bindings []GetPortBinding `pulumi:"bindings"`
	// See Argument Reference above.
	Description *string `pulumi:"description"`
	// See Argument Reference above.
	DeviceId *string `pulumi:"deviceId"`
	// See Argument Reference above.
	DeviceOwner *string `pulumi:"deviceOwner"`
	// The list of maps representing port DNS assignments.
	DnsAssignments []map[string]string `pulumi:"dnsAssignments"`
	// See Argument Reference above.
	DnsName *string `pulumi:"dnsName"`
	// An extra DHCP option configured on the port.
	// The structure is described below.
	ExtraDhcpOptions []GetPortExtraDhcpOption `pulumi:"extraDhcpOptions"`
	FixedIp          *string                  `pulumi:"fixedIp"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The additional MAC address.
	MacAddress *string `pulumi:"macAddress"`
	// Name of the DHCP option.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	NetworkId *string `pulumi:"networkId"`
	// See Argument Reference above.
	PortId *string `pulumi:"portId"`
	// See Argument Reference above.
	ProjectId *string `pulumi:"projectId"`
	// See Argument Reference above.
	Region           *string  `pulumi:"region"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	Status           *string  `pulumi:"status"`
	Tags             []string `pulumi:"tags"`
	TenantId         *string  `pulumi:"tenantId"`
}

func LookupPortOutput(ctx *pulumi.Context, args LookupPortOutputArgs, opts ...pulumi.InvokeOption) LookupPortResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPortResultOutput, error) {
			args := v.(LookupPortArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:networking/getPort:getPort", args, LookupPortResultOutput{}, options).(LookupPortResultOutput), nil
		}).(LookupPortResultOutput)
}

// A collection of arguments for invoking getPort.
type LookupPortOutputArgs struct {
	// The administrative state of the port.
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// Human-readable description of the port.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The ID of the device the port belongs to.
	DeviceId pulumi.StringPtrInput `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner pulumi.StringPtrInput `pulumi:"deviceOwner"`
	// The port DNS name to filter. Available, when Neutron
	// DNS extension is enabled.
	DnsName pulumi.StringPtrInput `pulumi:"dnsName"`
	// The port IP address filter.
	FixedIp pulumi.StringPtrInput `pulumi:"fixedIp"`
	// The MAC address of the port.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The name of the port.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the network the port belongs to.
	NetworkId pulumi.StringPtrInput `pulumi:"networkId"`
	// The ID of the port.
	PortId pulumi.StringPtrInput `pulumi:"portId"`
	// The owner of the port.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve port ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The list of port security group IDs to filter.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// The status of the port.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The list of port tags to filter.
	Tags     pulumi.StringArrayInput `pulumi:"tags"`
	TenantId pulumi.StringPtrInput   `pulumi:"tenantId"`
}

func (LookupPortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortArgs)(nil)).Elem()
}

// A collection of values returned by getPort.
type LookupPortResultOutput struct{ *pulumi.OutputState }

func (LookupPortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortResult)(nil)).Elem()
}

func (o LookupPortResultOutput) ToLookupPortResultOutput() LookupPortResultOutput {
	return o
}

func (o LookupPortResultOutput) ToLookupPortResultOutputWithContext(ctx context.Context) LookupPortResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupPortResultOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *bool { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// The collection of Fixed IP addresses on the port in the
// order returned by the Network v2 API.
func (o LookupPortResultOutput) AllFixedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.AllFixedIps }).(pulumi.StringArrayOutput)
}

// The set of security group IDs applied on the port.
func (o LookupPortResultOutput) AllSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.AllSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The set of string tags applied on the port.
func (o LookupPortResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

// An IP/MAC Address pair of additional IP
// addresses that can be active on this port. The structure is described
// below.
func (o LookupPortResultOutput) AllowedAddressPairs() GetPortAllowedAddressPairArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []GetPortAllowedAddressPair { return v.AllowedAddressPairs }).(GetPortAllowedAddressPairArrayOutput)
}

// The port binding information. The structure is described below.
func (o LookupPortResultOutput) Bindings() GetPortBindingArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []GetPortBinding { return v.Bindings }).(GetPortBindingArrayOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.DeviceId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) DeviceOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.DeviceOwner }).(pulumi.StringPtrOutput)
}

// The list of maps representing port DNS assignments.
func (o LookupPortResultOutput) DnsAssignments() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []map[string]string { return v.DnsAssignments }).(pulumi.StringMapArrayOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

// An extra DHCP option configured on the port.
// The structure is described below.
func (o LookupPortResultOutput) ExtraDhcpOptions() GetPortExtraDhcpOptionArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []GetPortExtraDhcpOption { return v.ExtraDhcpOptions }).(GetPortExtraDhcpOptionArrayOutput)
}

func (o LookupPortResultOutput) FixedIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.FixedIp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPortResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortResult) string { return v.Id }).(pulumi.StringOutput)
}

// The additional MAC address.
func (o LookupPortResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// Name of the DHCP option.
func (o LookupPortResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) PortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.PortId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupPortResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupPortResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o LookupPortResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupPortResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPortResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupPortResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPortResultOutput{})
}
