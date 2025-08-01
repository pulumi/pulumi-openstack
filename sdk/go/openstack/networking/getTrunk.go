// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack trunk.
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
//			_, err := networking.LookupTrunk(ctx, &networking.LookupTrunkArgs{
//				Name: pulumi.StringRef("trunk_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTrunk(ctx *pulumi.Context, args *LookupTrunkArgs, opts ...pulumi.InvokeOption) (*LookupTrunkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTrunkResult
	err := ctx.Invoke("openstack:networking/getTrunk:getTrunk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTrunk.
type LookupTrunkArgs struct {
	// The administrative state of the trunk.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description of the trunk.
	Description *string `pulumi:"description"`
	// The name of the trunk.
	Name *string `pulumi:"name"`
	// The ID of the trunk parent port.
	PortId *string `pulumi:"portId"`
	// The owner of the trunk.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve trunk ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the trunk.
	Status *string `pulumi:"status"`
	// The list of trunk tags to filter.
	Tags []string `pulumi:"tags"`
	// The ID of the trunk.
	TrunkId *string `pulumi:"trunkId"`
}

// A collection of values returned by getTrunk.
type LookupTrunkResult struct {
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The set of string tags applied on the trunk.
	AllTags     []string `pulumi:"allTags"`
	Description *string  `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The ID of the trunk subport.
	PortId    *string `pulumi:"portId"`
	ProjectId string  `pulumi:"projectId"`
	Region    string  `pulumi:"region"`
	Status    *string `pulumi:"status"`
	// The set of the trunk subports. The structure of each subport is
	// described below.
	SubPorts []GetTrunkSubPort `pulumi:"subPorts"`
	Tags     []string          `pulumi:"tags"`
	TrunkId  *string           `pulumi:"trunkId"`
}

func LookupTrunkOutput(ctx *pulumi.Context, args LookupTrunkOutputArgs, opts ...pulumi.InvokeOption) LookupTrunkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTrunkResultOutput, error) {
			args := v.(LookupTrunkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:networking/getTrunk:getTrunk", args, LookupTrunkResultOutput{}, options).(LookupTrunkResultOutput), nil
		}).(LookupTrunkResultOutput)
}

// A collection of arguments for invoking getTrunk.
type LookupTrunkOutputArgs struct {
	// The administrative state of the trunk.
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// Human-readable description of the trunk.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The name of the trunk.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the trunk parent port.
	PortId pulumi.StringPtrInput `pulumi:"portId"`
	// The owner of the trunk.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve trunk ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The status of the trunk.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The list of trunk tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The ID of the trunk.
	TrunkId pulumi.StringPtrInput `pulumi:"trunkId"`
}

func (LookupTrunkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrunkArgs)(nil)).Elem()
}

// A collection of values returned by getTrunk.
type LookupTrunkResultOutput struct{ *pulumi.OutputState }

func (LookupTrunkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrunkResult)(nil)).Elem()
}

func (o LookupTrunkResultOutput) ToLookupTrunkResultOutput() LookupTrunkResultOutput {
	return o
}

func (o LookupTrunkResultOutput) ToLookupTrunkResultOutputWithContext(ctx context.Context) LookupTrunkResultOutput {
	return o
}

func (o LookupTrunkResultOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTrunkResult) *bool { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// The set of string tags applied on the trunk.
func (o LookupTrunkResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrunkResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

func (o LookupTrunkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTrunkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTrunkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the trunk subport.
func (o LookupTrunkResultOutput) PortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkResult) *string { return v.PortId }).(pulumi.StringPtrOutput)
}

func (o LookupTrunkResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupTrunkResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrunkResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupTrunkResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The set of the trunk subports. The structure of each subport is
// described below.
func (o LookupTrunkResultOutput) SubPorts() GetTrunkSubPortArrayOutput {
	return o.ApplyT(func(v LookupTrunkResult) []GetTrunkSubPort { return v.SubPorts }).(GetTrunkSubPortArrayOutput)
}

func (o LookupTrunkResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrunkResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupTrunkResultOutput) TrunkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrunkResult) *string { return v.TrunkId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrunkResultOutput{})
}
