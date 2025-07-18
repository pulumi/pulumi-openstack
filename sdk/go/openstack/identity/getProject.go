// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an OpenStack project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.LookupProject(ctx, &identity.LookupProjectArgs{
//				Name: pulumi.StringRef("demo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("openstack:identity/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The domain this project belongs to.
	DomainId *string `pulumi:"domainId"`
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain *bool `pulumi:"isDomain"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The parent of this project.
	ParentId *string `pulumi:"parentId"`
	// The id of the project. Conflicts with any of the
	// above arguments.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// The description of the project.
	Description string `pulumi:"description"`
	// See Argument Reference above.
	DomainId string `pulumi:"domainId"`
	// See Argument Reference above.
	Enabled *bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IsDomain *bool `pulumi:"isDomain"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	ParentId *string `pulumi:"parentId"`
	// See Argument Reference above.
	ProjectId *string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	Tags []string `pulumi:"tags"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:identity/getProject:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The domain this project belongs to.
	DomainId pulumi.StringPtrInput `pulumi:"domainId"`
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain pulumi.BoolPtrInput `pulumi:"isDomain"`
	// The name of the project.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The parent of this project.
	ParentId pulumi.StringPtrInput `pulumi:"parentId"`
	// The id of the project. Conflicts with any of the
	// above arguments.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// The description of the project.
func (o LookupProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DomainId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) IsDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *bool { return v.IsDomain }).(pulumi.BoolPtrOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) ParentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ParentId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupProjectResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
