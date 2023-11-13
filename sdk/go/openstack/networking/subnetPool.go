// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron subnetpool resource within OpenStack.
//
// ## Example Usage
// ### Create a Subnet Pool
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
//			_, err := networking.NewSubnetPool(ctx, "subnetpool1", &networking.SubnetPoolArgs{
//				IpVersion: pulumi.Int(6),
//				Prefixes: pulumi.StringArray{
//					pulumi.String("fdf7:b13d:dead:beef::/64"),
//					pulumi.String("fd65:86cc:a334:39b7::/64"),
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
// ### Create a Subnet from a Subnet Pool
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
//			subnetpool1, err := networking.NewSubnetPool(ctx, "subnetpool1", &networking.SubnetPoolArgs{
//				Prefixes: pulumi.StringArray{
//					pulumi.String("10.11.12.0/24"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewSubnet(ctx, "subnet1", &networking.SubnetArgs{
//				Cidr:         pulumi.String("10.11.12.0/25"),
//				NetworkId:    network1.ID(),
//				SubnetpoolId: subnetpool1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Subnetpools can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:networking/subnetPool:SubnetPool subnetpool_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
//
// ```
type SubnetPool struct {
	pulumi.CustomResourceState

	// The Neutron address scope to assign to the
	// subnetpool. Changing this updates the address scope id of the existing
	// subnetpool.
	AddressScopeId pulumi.StringPtrOutput `pulumi:"addressScopeId"`
	// The collection of tags assigned on the subnetpool, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// The time at which subnetpool was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The size of the prefix to allocate when the cidr
	// or prefixlen attributes are omitted when you create the subnet. Defaults to the
	// MinPrefixLen. Changing this updates the default prefixlen of the existing
	// subnetpool.
	DefaultPrefixlen pulumi.IntOutput `pulumi:"defaultPrefixlen"`
	// The per-project quota on the prefix space that can be
	// allocated from the subnetpool for project subnets. Changing this updates the
	// default quota of the existing subnetpool.
	DefaultQuota pulumi.IntPtrOutput `pulumi:"defaultQuota"`
	// The human-readable description for the subnetpool.
	// Changing this updates the description of the existing subnetpool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP protocol version.
	IpVersion pulumi.IntOutput `pulumi:"ipVersion"`
	// Indicates whether the subnetpool is default
	// subnetpool or not. Changing this updates the default status of the existing
	// subnetpool.
	IsDefault pulumi.BoolPtrOutput `pulumi:"isDefault"`
	// The maximum prefix size that can be allocated from
	// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
	// default is 128. Changing this updates the max prefixlen of the existing
	// subnetpool.
	MaxPrefixlen pulumi.IntOutput `pulumi:"maxPrefixlen"`
	// The smallest prefix that can be allocated from a
	// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
	// is 64. Changing this updates the min prefixlen of the existing subnetpool.
	MinPrefixlen pulumi.IntOutput `pulumi:"minPrefixlen"`
	// The name of the subnetpool. Changing this updates the name of
	// the existing subnetpool.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of subnet prefixes to assign to the subnetpool.
	// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
	// subnet prefix must be unique among all subnet prefixes in all subnetpools that
	// are associated with the address scope. Changing this updates the prefixes list
	// of the existing subnetpool.
	Prefixes pulumi.StringArrayOutput `pulumi:"prefixes"`
	// The owner of the subnetpool. Required if admin wants to
	// create a subnetpool for another project. Changing this creates a new subnetpool.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnetpool. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnetpool.
	Region pulumi.StringOutput `pulumi:"region"`
	// The revision number of the subnetpool.
	RevisionNumber pulumi.IntOutput `pulumi:"revisionNumber"`
	// Indicates whether this subnetpool is shared across
	// all projects. Changing this updates the shared status of the existing
	// subnetpool.
	Shared pulumi.BoolPtrOutput `pulumi:"shared"`
	// A set of string tags for the subnetpool.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The time at which subnetpool was created.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewSubnetPool registers a new resource with the given unique name, arguments, and options.
func NewSubnetPool(ctx *pulumi.Context,
	name string, args *SubnetPoolArgs, opts ...pulumi.ResourceOption) (*SubnetPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Prefixes == nil {
		return nil, errors.New("invalid value for required argument 'Prefixes'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SubnetPool
	err := ctx.RegisterResource("openstack:networking/subnetPool:SubnetPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetPool gets an existing SubnetPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetPoolState, opts ...pulumi.ResourceOption) (*SubnetPool, error) {
	var resource SubnetPool
	err := ctx.ReadResource("openstack:networking/subnetPool:SubnetPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetPool resources.
type subnetPoolState struct {
	// The Neutron address scope to assign to the
	// subnetpool. Changing this updates the address scope id of the existing
	// subnetpool.
	AddressScopeId *string `pulumi:"addressScopeId"`
	// The collection of tags assigned on the subnetpool, which have been
	// explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// The time at which subnetpool was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The size of the prefix to allocate when the cidr
	// or prefixlen attributes are omitted when you create the subnet. Defaults to the
	// MinPrefixLen. Changing this updates the default prefixlen of the existing
	// subnetpool.
	DefaultPrefixlen *int `pulumi:"defaultPrefixlen"`
	// The per-project quota on the prefix space that can be
	// allocated from the subnetpool for project subnets. Changing this updates the
	// default quota of the existing subnetpool.
	DefaultQuota *int `pulumi:"defaultQuota"`
	// The human-readable description for the subnetpool.
	// Changing this updates the description of the existing subnetpool.
	Description *string `pulumi:"description"`
	// The IP protocol version.
	IpVersion *int `pulumi:"ipVersion"`
	// Indicates whether the subnetpool is default
	// subnetpool or not. Changing this updates the default status of the existing
	// subnetpool.
	IsDefault *bool `pulumi:"isDefault"`
	// The maximum prefix size that can be allocated from
	// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
	// default is 128. Changing this updates the max prefixlen of the existing
	// subnetpool.
	MaxPrefixlen *int `pulumi:"maxPrefixlen"`
	// The smallest prefix that can be allocated from a
	// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
	// is 64. Changing this updates the min prefixlen of the existing subnetpool.
	MinPrefixlen *int `pulumi:"minPrefixlen"`
	// The name of the subnetpool. Changing this updates the name of
	// the existing subnetpool.
	Name *string `pulumi:"name"`
	// A list of subnet prefixes to assign to the subnetpool.
	// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
	// subnet prefix must be unique among all subnet prefixes in all subnetpools that
	// are associated with the address scope. Changing this updates the prefixes list
	// of the existing subnetpool.
	Prefixes []string `pulumi:"prefixes"`
	// The owner of the subnetpool. Required if admin wants to
	// create a subnetpool for another project. Changing this creates a new subnetpool.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnetpool. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnetpool.
	Region *string `pulumi:"region"`
	// The revision number of the subnetpool.
	RevisionNumber *int `pulumi:"revisionNumber"`
	// Indicates whether this subnetpool is shared across
	// all projects. Changing this updates the shared status of the existing
	// subnetpool.
	Shared *bool `pulumi:"shared"`
	// A set of string tags for the subnetpool.
	Tags []string `pulumi:"tags"`
	// The time at which subnetpool was created.
	UpdatedAt *string `pulumi:"updatedAt"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type SubnetPoolState struct {
	// The Neutron address scope to assign to the
	// subnetpool. Changing this updates the address scope id of the existing
	// subnetpool.
	AddressScopeId pulumi.StringPtrInput
	// The collection of tags assigned on the subnetpool, which have been
	// explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// The time at which subnetpool was created.
	CreatedAt pulumi.StringPtrInput
	// The size of the prefix to allocate when the cidr
	// or prefixlen attributes are omitted when you create the subnet. Defaults to the
	// MinPrefixLen. Changing this updates the default prefixlen of the existing
	// subnetpool.
	DefaultPrefixlen pulumi.IntPtrInput
	// The per-project quota on the prefix space that can be
	// allocated from the subnetpool for project subnets. Changing this updates the
	// default quota of the existing subnetpool.
	DefaultQuota pulumi.IntPtrInput
	// The human-readable description for the subnetpool.
	// Changing this updates the description of the existing subnetpool.
	Description pulumi.StringPtrInput
	// The IP protocol version.
	IpVersion pulumi.IntPtrInput
	// Indicates whether the subnetpool is default
	// subnetpool or not. Changing this updates the default status of the existing
	// subnetpool.
	IsDefault pulumi.BoolPtrInput
	// The maximum prefix size that can be allocated from
	// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
	// default is 128. Changing this updates the max prefixlen of the existing
	// subnetpool.
	MaxPrefixlen pulumi.IntPtrInput
	// The smallest prefix that can be allocated from a
	// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
	// is 64. Changing this updates the min prefixlen of the existing subnetpool.
	MinPrefixlen pulumi.IntPtrInput
	// The name of the subnetpool. Changing this updates the name of
	// the existing subnetpool.
	Name pulumi.StringPtrInput
	// A list of subnet prefixes to assign to the subnetpool.
	// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
	// subnet prefix must be unique among all subnet prefixes in all subnetpools that
	// are associated with the address scope. Changing this updates the prefixes list
	// of the existing subnetpool.
	Prefixes pulumi.StringArrayInput
	// The owner of the subnetpool. Required if admin wants to
	// create a subnetpool for another project. Changing this creates a new subnetpool.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnetpool. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnetpool.
	Region pulumi.StringPtrInput
	// The revision number of the subnetpool.
	RevisionNumber pulumi.IntPtrInput
	// Indicates whether this subnetpool is shared across
	// all projects. Changing this updates the shared status of the existing
	// subnetpool.
	Shared pulumi.BoolPtrInput
	// A set of string tags for the subnetpool.
	Tags pulumi.StringArrayInput
	// The time at which subnetpool was created.
	UpdatedAt pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (SubnetPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetPoolState)(nil)).Elem()
}

type subnetPoolArgs struct {
	// The Neutron address scope to assign to the
	// subnetpool. Changing this updates the address scope id of the existing
	// subnetpool.
	AddressScopeId *string `pulumi:"addressScopeId"`
	// The size of the prefix to allocate when the cidr
	// or prefixlen attributes are omitted when you create the subnet. Defaults to the
	// MinPrefixLen. Changing this updates the default prefixlen of the existing
	// subnetpool.
	DefaultPrefixlen *int `pulumi:"defaultPrefixlen"`
	// The per-project quota on the prefix space that can be
	// allocated from the subnetpool for project subnets. Changing this updates the
	// default quota of the existing subnetpool.
	DefaultQuota *int `pulumi:"defaultQuota"`
	// The human-readable description for the subnetpool.
	// Changing this updates the description of the existing subnetpool.
	Description *string `pulumi:"description"`
	// The IP protocol version.
	IpVersion *int `pulumi:"ipVersion"`
	// Indicates whether the subnetpool is default
	// subnetpool or not. Changing this updates the default status of the existing
	// subnetpool.
	IsDefault *bool `pulumi:"isDefault"`
	// The maximum prefix size that can be allocated from
	// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
	// default is 128. Changing this updates the max prefixlen of the existing
	// subnetpool.
	MaxPrefixlen *int `pulumi:"maxPrefixlen"`
	// The smallest prefix that can be allocated from a
	// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
	// is 64. Changing this updates the min prefixlen of the existing subnetpool.
	MinPrefixlen *int `pulumi:"minPrefixlen"`
	// The name of the subnetpool. Changing this updates the name of
	// the existing subnetpool.
	Name *string `pulumi:"name"`
	// A list of subnet prefixes to assign to the subnetpool.
	// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
	// subnet prefix must be unique among all subnet prefixes in all subnetpools that
	// are associated with the address scope. Changing this updates the prefixes list
	// of the existing subnetpool.
	Prefixes []string `pulumi:"prefixes"`
	// The owner of the subnetpool. Required if admin wants to
	// create a subnetpool for another project. Changing this creates a new subnetpool.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnetpool. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnetpool.
	Region *string `pulumi:"region"`
	// Indicates whether this subnetpool is shared across
	// all projects. Changing this updates the shared status of the existing
	// subnetpool.
	Shared *bool `pulumi:"shared"`
	// A set of string tags for the subnetpool.
	Tags []string `pulumi:"tags"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a SubnetPool resource.
type SubnetPoolArgs struct {
	// The Neutron address scope to assign to the
	// subnetpool. Changing this updates the address scope id of the existing
	// subnetpool.
	AddressScopeId pulumi.StringPtrInput
	// The size of the prefix to allocate when the cidr
	// or prefixlen attributes are omitted when you create the subnet. Defaults to the
	// MinPrefixLen. Changing this updates the default prefixlen of the existing
	// subnetpool.
	DefaultPrefixlen pulumi.IntPtrInput
	// The per-project quota on the prefix space that can be
	// allocated from the subnetpool for project subnets. Changing this updates the
	// default quota of the existing subnetpool.
	DefaultQuota pulumi.IntPtrInput
	// The human-readable description for the subnetpool.
	// Changing this updates the description of the existing subnetpool.
	Description pulumi.StringPtrInput
	// The IP protocol version.
	IpVersion pulumi.IntPtrInput
	// Indicates whether the subnetpool is default
	// subnetpool or not. Changing this updates the default status of the existing
	// subnetpool.
	IsDefault pulumi.BoolPtrInput
	// The maximum prefix size that can be allocated from
	// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
	// default is 128. Changing this updates the max prefixlen of the existing
	// subnetpool.
	MaxPrefixlen pulumi.IntPtrInput
	// The smallest prefix that can be allocated from a
	// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
	// is 64. Changing this updates the min prefixlen of the existing subnetpool.
	MinPrefixlen pulumi.IntPtrInput
	// The name of the subnetpool. Changing this updates the name of
	// the existing subnetpool.
	Name pulumi.StringPtrInput
	// A list of subnet prefixes to assign to the subnetpool.
	// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
	// subnet prefix must be unique among all subnet prefixes in all subnetpools that
	// are associated with the address scope. Changing this updates the prefixes list
	// of the existing subnetpool.
	Prefixes pulumi.StringArrayInput
	// The owner of the subnetpool. Required if admin wants to
	// create a subnetpool for another project. Changing this creates a new subnetpool.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnetpool. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// subnetpool.
	Region pulumi.StringPtrInput
	// Indicates whether this subnetpool is shared across
	// all projects. Changing this updates the shared status of the existing
	// subnetpool.
	Shared pulumi.BoolPtrInput
	// A set of string tags for the subnetpool.
	Tags pulumi.StringArrayInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (SubnetPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetPoolArgs)(nil)).Elem()
}

type SubnetPoolInput interface {
	pulumi.Input

	ToSubnetPoolOutput() SubnetPoolOutput
	ToSubnetPoolOutputWithContext(ctx context.Context) SubnetPoolOutput
}

func (*SubnetPool) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetPool)(nil)).Elem()
}

func (i *SubnetPool) ToSubnetPoolOutput() SubnetPoolOutput {
	return i.ToSubnetPoolOutputWithContext(context.Background())
}

func (i *SubnetPool) ToSubnetPoolOutputWithContext(ctx context.Context) SubnetPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPoolOutput)
}

// SubnetPoolArrayInput is an input type that accepts SubnetPoolArray and SubnetPoolArrayOutput values.
// You can construct a concrete instance of `SubnetPoolArrayInput` via:
//
//	SubnetPoolArray{ SubnetPoolArgs{...} }
type SubnetPoolArrayInput interface {
	pulumi.Input

	ToSubnetPoolArrayOutput() SubnetPoolArrayOutput
	ToSubnetPoolArrayOutputWithContext(context.Context) SubnetPoolArrayOutput
}

type SubnetPoolArray []SubnetPoolInput

func (SubnetPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubnetPool)(nil)).Elem()
}

func (i SubnetPoolArray) ToSubnetPoolArrayOutput() SubnetPoolArrayOutput {
	return i.ToSubnetPoolArrayOutputWithContext(context.Background())
}

func (i SubnetPoolArray) ToSubnetPoolArrayOutputWithContext(ctx context.Context) SubnetPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPoolArrayOutput)
}

// SubnetPoolMapInput is an input type that accepts SubnetPoolMap and SubnetPoolMapOutput values.
// You can construct a concrete instance of `SubnetPoolMapInput` via:
//
//	SubnetPoolMap{ "key": SubnetPoolArgs{...} }
type SubnetPoolMapInput interface {
	pulumi.Input

	ToSubnetPoolMapOutput() SubnetPoolMapOutput
	ToSubnetPoolMapOutputWithContext(context.Context) SubnetPoolMapOutput
}

type SubnetPoolMap map[string]SubnetPoolInput

func (SubnetPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubnetPool)(nil)).Elem()
}

func (i SubnetPoolMap) ToSubnetPoolMapOutput() SubnetPoolMapOutput {
	return i.ToSubnetPoolMapOutputWithContext(context.Background())
}

func (i SubnetPoolMap) ToSubnetPoolMapOutputWithContext(ctx context.Context) SubnetPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPoolMapOutput)
}

type SubnetPoolOutput struct{ *pulumi.OutputState }

func (SubnetPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetPool)(nil)).Elem()
}

func (o SubnetPoolOutput) ToSubnetPoolOutput() SubnetPoolOutput {
	return o
}

func (o SubnetPoolOutput) ToSubnetPoolOutputWithContext(ctx context.Context) SubnetPoolOutput {
	return o
}

// The Neutron address scope to assign to the
// subnetpool. Changing this updates the address scope id of the existing
// subnetpool.
func (o SubnetPoolOutput) AddressScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringPtrOutput { return v.AddressScopeId }).(pulumi.StringPtrOutput)
}

// The collection of tags assigned on the subnetpool, which have been
// explicitly and implicitly added.
func (o SubnetPoolOutput) AllTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringArrayOutput { return v.AllTags }).(pulumi.StringArrayOutput)
}

// The time at which subnetpool was created.
func (o SubnetPoolOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The size of the prefix to allocate when the cidr
// or prefixlen attributes are omitted when you create the subnet. Defaults to the
// MinPrefixLen. Changing this updates the default prefixlen of the existing
// subnetpool.
func (o SubnetPoolOutput) DefaultPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.IntOutput { return v.DefaultPrefixlen }).(pulumi.IntOutput)
}

// The per-project quota on the prefix space that can be
// allocated from the subnetpool for project subnets. Changing this updates the
// default quota of the existing subnetpool.
func (o SubnetPoolOutput) DefaultQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.IntPtrOutput { return v.DefaultQuota }).(pulumi.IntPtrOutput)
}

// The human-readable description for the subnetpool.
// Changing this updates the description of the existing subnetpool.
func (o SubnetPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IP protocol version.
func (o SubnetPoolOutput) IpVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.IntOutput { return v.IpVersion }).(pulumi.IntOutput)
}

// Indicates whether the subnetpool is default
// subnetpool or not. Changing this updates the default status of the existing
// subnetpool.
func (o SubnetPoolOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.BoolPtrOutput { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The maximum prefix size that can be allocated from
// the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
// default is 128. Changing this updates the max prefixlen of the existing
// subnetpool.
func (o SubnetPoolOutput) MaxPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.IntOutput { return v.MaxPrefixlen }).(pulumi.IntOutput)
}

// The smallest prefix that can be allocated from a
// subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
// is 64. Changing this updates the min prefixlen of the existing subnetpool.
func (o SubnetPoolOutput) MinPrefixlen() pulumi.IntOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.IntOutput { return v.MinPrefixlen }).(pulumi.IntOutput)
}

// The name of the subnetpool. Changing this updates the name of
// the existing subnetpool.
func (o SubnetPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of subnet prefixes to assign to the subnetpool.
// Neutron API merges adjacent prefixes and treats them as a single prefix. Each
// subnet prefix must be unique among all subnet prefixes in all subnetpools that
// are associated with the address scope. Changing this updates the prefixes list
// of the existing subnetpool.
func (o SubnetPoolOutput) Prefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringArrayOutput { return v.Prefixes }).(pulumi.StringArrayOutput)
}

// The owner of the subnetpool. Required if admin wants to
// create a subnetpool for another project. Changing this creates a new subnetpool.
func (o SubnetPoolOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron subnetpool. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// subnetpool.
func (o SubnetPoolOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The revision number of the subnetpool.
func (o SubnetPoolOutput) RevisionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.IntOutput { return v.RevisionNumber }).(pulumi.IntOutput)
}

// Indicates whether this subnetpool is shared across
// all projects. Changing this updates the shared status of the existing
// subnetpool.
func (o SubnetPoolOutput) Shared() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.BoolPtrOutput { return v.Shared }).(pulumi.BoolPtrOutput)
}

// A set of string tags for the subnetpool.
func (o SubnetPoolOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The time at which subnetpool was created.
func (o SubnetPoolOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Map of additional options.
func (o SubnetPoolOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *SubnetPool) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type SubnetPoolArrayOutput struct{ *pulumi.OutputState }

func (SubnetPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubnetPool)(nil)).Elem()
}

func (o SubnetPoolArrayOutput) ToSubnetPoolArrayOutput() SubnetPoolArrayOutput {
	return o
}

func (o SubnetPoolArrayOutput) ToSubnetPoolArrayOutputWithContext(ctx context.Context) SubnetPoolArrayOutput {
	return o
}

func (o SubnetPoolArrayOutput) Index(i pulumi.IntInput) SubnetPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubnetPool {
		return vs[0].([]*SubnetPool)[vs[1].(int)]
	}).(SubnetPoolOutput)
}

type SubnetPoolMapOutput struct{ *pulumi.OutputState }

func (SubnetPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubnetPool)(nil)).Elem()
}

func (o SubnetPoolMapOutput) ToSubnetPoolMapOutput() SubnetPoolMapOutput {
	return o
}

func (o SubnetPoolMapOutput) ToSubnetPoolMapOutputWithContext(ctx context.Context) SubnetPoolMapOutput {
	return o
}

func (o SubnetPoolMapOutput) MapIndex(k pulumi.StringInput) SubnetPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubnetPool {
		return vs[0].(map[string]*SubnetPool)[vs[1].(string)]
	}).(SubnetPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetPoolInput)(nil)).Elem(), &SubnetPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetPoolArrayInput)(nil)).Elem(), SubnetPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetPoolMapInput)(nil)).Elem(), SubnetPoolMap{})
	pulumi.RegisterOutputType(SubnetPoolOutput{})
	pulumi.RegisterOutputType(SubnetPoolArrayOutput{})
	pulumi.RegisterOutputType(SubnetPoolMapOutput{})
}
