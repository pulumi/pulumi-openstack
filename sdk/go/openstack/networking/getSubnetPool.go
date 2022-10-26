// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack subnetpool.
//
// ## Example Usage
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
//			_, err := networking.LookupSubnetPool(ctx, &networking.LookupSubnetPoolArgs{
//				Name: pulumi.StringRef("subnetpool_1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSubnetPool(ctx *pulumi.Context, args *LookupSubnetPoolArgs, opts ...pulumi.InvokeOption) (*LookupSubnetPoolResult, error) {
	var rv LookupSubnetPoolResult
	err := ctx.Invoke("openstack:networking/getSubnetPool:getSubnetPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnetPool.
type LookupSubnetPoolArgs struct {
	// The Neutron address scope that subnetpools
	// is assigned to.
	AddressScopeId *string `pulumi:"addressScopeId"`
	// The size of the subnetpool default prefix
	// length.
	DefaultPrefixlen *int `pulumi:"defaultPrefixlen"`
	// The per-project quota on the prefix space that
	// can be allocated from the subnetpool for project subnets.
	DefaultQuota *int `pulumi:"defaultQuota"`
	// The human-readable description for the subnetpool.
	Description *string `pulumi:"description"`
	// The IP protocol version.
	IpVersion *int `pulumi:"ipVersion"`
	// Whether the subnetpool is default subnetpool or not.
	IsDefault *bool `pulumi:"isDefault"`
	// The size of the subnetpool max prefix length.
	MaxPrefixlen *int `pulumi:"maxPrefixlen"`
	// The size of the subnetpool min prefix length.
	MinPrefixlen *int `pulumi:"minPrefixlen"`
	// The name of the subnetpool.
	Name *string `pulumi:"name"`
	// The owner of the subnetpool.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to retrieve a subnetpool id. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// Whether this subnetpool is shared across all projects.
	Shared *bool `pulumi:"shared"`
	// The list of subnetpool tags to filter.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getSubnetPool.
type LookupSubnetPoolResult struct {
	// See Argument Reference above.
	// * `ipVersion` -The IP protocol version.
	AddressScopeId string `pulumi:"addressScopeId"`
	// The set of string tags applied on the subnetpool.
	AllTags []string `pulumi:"allTags"`
	// The time at which subnetpool was created.
	CreatedAt string `pulumi:"createdAt"`
	// See Argument Reference above.
	DefaultPrefixlen int `pulumi:"defaultPrefixlen"`
	// See Argument Reference above.
	DefaultQuota int `pulumi:"defaultQuota"`
	// See Argument Reference above.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	IpVersion int    `pulumi:"ipVersion"`
	// See Argument Reference above.
	IsDefault bool `pulumi:"isDefault"`
	// See Argument Reference above.
	MaxPrefixlen int `pulumi:"maxPrefixlen"`
	// See Argument Reference above.
	MinPrefixlen int `pulumi:"minPrefixlen"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	Prefixes []string `pulumi:"prefixes"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The revision number of the subnetpool.
	RevisionNumber int `pulumi:"revisionNumber"`
	// See Argument Reference above.
	Shared bool     `pulumi:"shared"`
	Tags   []string `pulumi:"tags"`
	// The time at which subnetpool was created.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupSubnetPoolOutput(ctx *pulumi.Context, args LookupSubnetPoolOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetPoolResult, error) {
			args := v.(LookupSubnetPoolArgs)
			r, err := LookupSubnetPool(ctx, &args, opts...)
			var s LookupSubnetPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetPoolResultOutput)
}

// A collection of arguments for invoking getSubnetPool.
type LookupSubnetPoolOutputArgs struct {
	// The Neutron address scope that subnetpools
	// is assigned to.
	AddressScopeId pulumi.StringPtrInput `pulumi:"addressScopeId"`
	// The size of the subnetpool default prefix
	// length.
	DefaultPrefixlen pulumi.IntPtrInput `pulumi:"defaultPrefixlen"`
	// The per-project quota on the prefix space that
	// can be allocated from the subnetpool for project subnets.
	DefaultQuota pulumi.IntPtrInput `pulumi:"defaultQuota"`
	// The human-readable description for the subnetpool.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The IP protocol version.
	IpVersion pulumi.IntPtrInput `pulumi:"ipVersion"`
	// Whether the subnetpool is default subnetpool or not.
	IsDefault pulumi.BoolPtrInput `pulumi:"isDefault"`
	// The size of the subnetpool max prefix length.
	MaxPrefixlen pulumi.IntPtrInput `pulumi:"maxPrefixlen"`
	// The size of the subnetpool min prefix length.
	MinPrefixlen pulumi.IntPtrInput `pulumi:"minPrefixlen"`
	// The name of the subnetpool.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The owner of the subnetpool.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to retrieve a subnetpool id. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Whether this subnetpool is shared across all projects.
	Shared pulumi.BoolPtrInput `pulumi:"shared"`
	// The list of subnetpool tags to filter.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (LookupSubnetPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetPoolArgs)(nil)).Elem()
}

// A collection of values returned by getSubnetPool.
type LookupSubnetPoolResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetPoolResult)(nil)).Elem()
}

func (o LookupSubnetPoolResultOutput) ToLookupSubnetPoolResultOutput() LookupSubnetPoolResultOutput {
	return o
}

func (o LookupSubnetPoolResultOutput) ToLookupSubnetPoolResultOutputWithContext(ctx context.Context) LookupSubnetPoolResultOutput {
	return o
}

// See Argument Reference above.
// * `ipVersion` -The IP protocol version.
func (o LookupSubnetPoolResultOutput) AddressScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.AddressScopeId }).(pulumi.StringOutput)
}

// The set of string tags applied on the subnetpool.
func (o LookupSubnetPoolResultOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) []string { return v.AllTags }).(pulumi.StringArrayOutput)
}

// The time at which subnetpool was created.
func (o LookupSubnetPoolResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) DefaultPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) int { return v.DefaultPrefixlen }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) DefaultQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) int { return v.DefaultQuota }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSubnetPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubnetPoolResultOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) int { return v.IpVersion }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) MaxPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) int { return v.MaxPrefixlen }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) MinPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) int { return v.MinPrefixlen }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) Prefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) []string { return v.Prefixes }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.Region }).(pulumi.StringOutput)
}

// The revision number of the subnetpool.
func (o LookupSubnetPoolResultOutput) RevisionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) int { return v.RevisionNumber }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o LookupSubnetPoolResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) bool { return v.Shared }).(pulumi.BoolOutput)
}

func (o LookupSubnetPoolResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The time at which subnetpool was created.
func (o LookupSubnetPoolResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetPoolResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetPoolResultOutput{})
}
