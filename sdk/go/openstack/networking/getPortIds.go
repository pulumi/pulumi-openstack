// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of Openstack Port IDs matching the
// specified criteria.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networking.GetPortIds(ctx, &networking.GetPortIdsArgs{
// 			Name: pulumi.StringRef("port"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetPortIds(ctx *pulumi.Context, args *GetPortIdsArgs, opts ...pulumi.InvokeOption) (*GetPortIdsResult, error) {
	var rv GetPortIdsResult
	err := ctx.Invoke("openstack:networking/getPortIds:getPortIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPortIds.
type GetPortIdsArgs struct {
	// The administrative state of the port.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description of the port.
	Description *string `pulumi:"description"`
	// The ID of the device the port belongs to.
	DeviceId *string `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner *string `pulumi:"deviceOwner"`
	DnsName     *string `pulumi:"dnsName"`
	// The port IP address filter.
	FixedIp *string `pulumi:"fixedIp"`
	// The MAC address of the port.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the port.
	Name *string `pulumi:"name"`
	// The ID of the network the port belongs to.
	NetworkId *string `pulumi:"networkId"`
	// The owner of the port.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve port ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The list of port security group IDs to filter.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Order the results in either `asc` or `desc`.
	// Defaults to none.
	SortDirection *string `pulumi:"sortDirection"`
	// Sort ports based on a certain key. Defaults to none.
	SortKey *string `pulumi:"sortKey"`
	// The status of the port.
	Status *string `pulumi:"status"`
	// The list of port tags to filter.
	Tags     []string `pulumi:"tags"`
	TenantId *string  `pulumi:"tenantId"`
}

// A collection of values returned by getPortIds.
type GetPortIdsResult struct {
	AdminStateUp *bool   `pulumi:"adminStateUp"`
	Description  *string `pulumi:"description"`
	DeviceId     *string `pulumi:"deviceId"`
	DeviceOwner  *string `pulumi:"deviceOwner"`
	DnsName      *string `pulumi:"dnsName"`
	FixedIp      *string `pulumi:"fixedIp"`
	// The provider-assigned unique ID for this managed resource.
	Id               string   `pulumi:"id"`
	Ids              []string `pulumi:"ids"`
	MacAddress       *string  `pulumi:"macAddress"`
	Name             *string  `pulumi:"name"`
	NetworkId        *string  `pulumi:"networkId"`
	ProjectId        *string  `pulumi:"projectId"`
	Region           *string  `pulumi:"region"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	SortDirection    *string  `pulumi:"sortDirection"`
	SortKey          *string  `pulumi:"sortKey"`
	Status           *string  `pulumi:"status"`
	Tags             []string `pulumi:"tags"`
	TenantId         *string  `pulumi:"tenantId"`
}

func GetPortIdsOutput(ctx *pulumi.Context, args GetPortIdsOutputArgs, opts ...pulumi.InvokeOption) GetPortIdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPortIdsResult, error) {
			args := v.(GetPortIdsArgs)
			r, err := GetPortIds(ctx, &args, opts...)
			return *r, err
		}).(GetPortIdsResultOutput)
}

// A collection of arguments for invoking getPortIds.
type GetPortIdsOutputArgs struct {
	// The administrative state of the port.
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// Human-readable description of the port.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The ID of the device the port belongs to.
	DeviceId pulumi.StringPtrInput `pulumi:"deviceId"`
	// The device owner of the port.
	DeviceOwner pulumi.StringPtrInput `pulumi:"deviceOwner"`
	DnsName     pulumi.StringPtrInput `pulumi:"dnsName"`
	// The port IP address filter.
	FixedIp pulumi.StringPtrInput `pulumi:"fixedIp"`
	// The MAC address of the port.
	MacAddress pulumi.StringPtrInput `pulumi:"macAddress"`
	// The name of the port.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the network the port belongs to.
	NetworkId pulumi.StringPtrInput `pulumi:"networkId"`
	// The owner of the port.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve port ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The list of port security group IDs to filter.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// Order the results in either `asc` or `desc`.
	// Defaults to none.
	SortDirection pulumi.StringPtrInput `pulumi:"sortDirection"`
	// Sort ports based on a certain key. Defaults to none.
	SortKey pulumi.StringPtrInput `pulumi:"sortKey"`
	// The status of the port.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The list of port tags to filter.
	Tags     pulumi.StringArrayInput `pulumi:"tags"`
	TenantId pulumi.StringPtrInput   `pulumi:"tenantId"`
}

func (GetPortIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPortIdsArgs)(nil)).Elem()
}

// A collection of values returned by getPortIds.
type GetPortIdsResultOutput struct{ *pulumi.OutputState }

func (GetPortIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPortIdsResult)(nil)).Elem()
}

func (o GetPortIdsResultOutput) ToGetPortIdsResultOutput() GetPortIdsResultOutput {
	return o
}

func (o GetPortIdsResultOutput) ToGetPortIdsResultOutputWithContext(ctx context.Context) GetPortIdsResultOutput {
	return o
}

func (o GetPortIdsResultOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *bool { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

func (o GetPortIdsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.DeviceId }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) DeviceOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.DeviceOwner }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) FixedIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.FixedIp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPortIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPortIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPortIdsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPortIdsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetPortIdsResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) NetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.NetworkId }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPortIdsResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o GetPortIdsResultOutput) SortDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.SortDirection }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) SortKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.SortKey }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetPortIdsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPortIdsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GetPortIdsResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPortIdsResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPortIdsResultOutput{})
}
