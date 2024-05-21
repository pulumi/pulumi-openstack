// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 load balancer member resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/loadbalancer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := loadbalancer.NewMemberV1(ctx, "member_1", &loadbalancer.MemberV1Args{
//				PoolId:  pulumi.String("d9415786-5f1a-428b-b35f-2f1523e146d2"),
//				Address: pulumi.String("192.168.0.10"),
//				Port:    pulumi.Int(80),
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
// Load Balancer Members can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:loadbalancer/memberV1:MemberV1 member_1 a7498676-4fe4-4243-a864-2eaaf18c73df
// ```
type MemberV1 struct {
	pulumi.CustomResourceState

	// The IP address of the member. Changing this creates a
	// new member.
	Address pulumi.StringOutput `pulumi:"address"`
	// The administrative state of the member.
	// Acceptable values are 'true' and 'false'. Changing this value updates the
	// state of the existing member.
	AdminStateUp pulumi.BoolOutput `pulumi:"adminStateUp"`
	// The ID of the LB pool. Changing this creates a new
	// member.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// An integer representing the port on which the member is
	// hosted. Changing this creates a new member.
	Port pulumi.IntOutput `pulumi:"port"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the member. Required if admin wants to
	// create a member for another tenant. Changing this creates a new member.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	Weight   pulumi.IntOutput       `pulumi:"weight"`
}

// NewMemberV1 registers a new resource with the given unique name, arguments, and options.
func NewMemberV1(ctx *pulumi.Context,
	name string, args *MemberV1Args, opts ...pulumi.ResourceOption) (*MemberV1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.PoolId == nil {
		return nil, errors.New("invalid value for required argument 'PoolId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MemberV1
	err := ctx.RegisterResource("openstack:loadbalancer/memberV1:MemberV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMemberV1 gets an existing MemberV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMemberV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberV1State, opts ...pulumi.ResourceOption) (*MemberV1, error) {
	var resource MemberV1
	err := ctx.ReadResource("openstack:loadbalancer/memberV1:MemberV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MemberV1 resources.
type memberV1State struct {
	// The IP address of the member. Changing this creates a
	// new member.
	Address *string `pulumi:"address"`
	// The administrative state of the member.
	// Acceptable values are 'true' and 'false'. Changing this value updates the
	// state of the existing member.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The ID of the LB pool. Changing this creates a new
	// member.
	PoolId *string `pulumi:"poolId"`
	// An integer representing the port on which the member is
	// hosted. Changing this creates a new member.
	Port *int `pulumi:"port"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region *string `pulumi:"region"`
	// The owner of the member. Required if admin wants to
	// create a member for another tenant. Changing this creates a new member.
	TenantId *string `pulumi:"tenantId"`
	Weight   *int    `pulumi:"weight"`
}

type MemberV1State struct {
	// The IP address of the member. Changing this creates a
	// new member.
	Address pulumi.StringPtrInput
	// The administrative state of the member.
	// Acceptable values are 'true' and 'false'. Changing this value updates the
	// state of the existing member.
	AdminStateUp pulumi.BoolPtrInput
	// The ID of the LB pool. Changing this creates a new
	// member.
	PoolId pulumi.StringPtrInput
	// An integer representing the port on which the member is
	// hosted. Changing this creates a new member.
	Port pulumi.IntPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region pulumi.StringPtrInput
	// The owner of the member. Required if admin wants to
	// create a member for another tenant. Changing this creates a new member.
	TenantId pulumi.StringPtrInput
	Weight   pulumi.IntPtrInput
}

func (MemberV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*memberV1State)(nil)).Elem()
}

type memberV1Args struct {
	// The IP address of the member. Changing this creates a
	// new member.
	Address string `pulumi:"address"`
	// The administrative state of the member.
	// Acceptable values are 'true' and 'false'. Changing this value updates the
	// state of the existing member.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The ID of the LB pool. Changing this creates a new
	// member.
	PoolId string `pulumi:"poolId"`
	// An integer representing the port on which the member is
	// hosted. Changing this creates a new member.
	Port int `pulumi:"port"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region *string `pulumi:"region"`
	// The owner of the member. Required if admin wants to
	// create a member for another tenant. Changing this creates a new member.
	TenantId *string `pulumi:"tenantId"`
	Weight   *int    `pulumi:"weight"`
}

// The set of arguments for constructing a MemberV1 resource.
type MemberV1Args struct {
	// The IP address of the member. Changing this creates a
	// new member.
	Address pulumi.StringInput
	// The administrative state of the member.
	// Acceptable values are 'true' and 'false'. Changing this value updates the
	// state of the existing member.
	AdminStateUp pulumi.BoolPtrInput
	// The ID of the LB pool. Changing this creates a new
	// member.
	PoolId pulumi.StringInput
	// An integer representing the port on which the member is
	// hosted. Changing this creates a new member.
	Port pulumi.IntInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region pulumi.StringPtrInput
	// The owner of the member. Required if admin wants to
	// create a member for another tenant. Changing this creates a new member.
	TenantId pulumi.StringPtrInput
	Weight   pulumi.IntPtrInput
}

func (MemberV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*memberV1Args)(nil)).Elem()
}

type MemberV1Input interface {
	pulumi.Input

	ToMemberV1Output() MemberV1Output
	ToMemberV1OutputWithContext(ctx context.Context) MemberV1Output
}

func (*MemberV1) ElementType() reflect.Type {
	return reflect.TypeOf((**MemberV1)(nil)).Elem()
}

func (i *MemberV1) ToMemberV1Output() MemberV1Output {
	return i.ToMemberV1OutputWithContext(context.Background())
}

func (i *MemberV1) ToMemberV1OutputWithContext(ctx context.Context) MemberV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(MemberV1Output)
}

// MemberV1ArrayInput is an input type that accepts MemberV1Array and MemberV1ArrayOutput values.
// You can construct a concrete instance of `MemberV1ArrayInput` via:
//
//	MemberV1Array{ MemberV1Args{...} }
type MemberV1ArrayInput interface {
	pulumi.Input

	ToMemberV1ArrayOutput() MemberV1ArrayOutput
	ToMemberV1ArrayOutputWithContext(context.Context) MemberV1ArrayOutput
}

type MemberV1Array []MemberV1Input

func (MemberV1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MemberV1)(nil)).Elem()
}

func (i MemberV1Array) ToMemberV1ArrayOutput() MemberV1ArrayOutput {
	return i.ToMemberV1ArrayOutputWithContext(context.Background())
}

func (i MemberV1Array) ToMemberV1ArrayOutputWithContext(ctx context.Context) MemberV1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberV1ArrayOutput)
}

// MemberV1MapInput is an input type that accepts MemberV1Map and MemberV1MapOutput values.
// You can construct a concrete instance of `MemberV1MapInput` via:
//
//	MemberV1Map{ "key": MemberV1Args{...} }
type MemberV1MapInput interface {
	pulumi.Input

	ToMemberV1MapOutput() MemberV1MapOutput
	ToMemberV1MapOutputWithContext(context.Context) MemberV1MapOutput
}

type MemberV1Map map[string]MemberV1Input

func (MemberV1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MemberV1)(nil)).Elem()
}

func (i MemberV1Map) ToMemberV1MapOutput() MemberV1MapOutput {
	return i.ToMemberV1MapOutputWithContext(context.Background())
}

func (i MemberV1Map) ToMemberV1MapOutputWithContext(ctx context.Context) MemberV1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberV1MapOutput)
}

type MemberV1Output struct{ *pulumi.OutputState }

func (MemberV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**MemberV1)(nil)).Elem()
}

func (o MemberV1Output) ToMemberV1Output() MemberV1Output {
	return o
}

func (o MemberV1Output) ToMemberV1OutputWithContext(ctx context.Context) MemberV1Output {
	return o
}

// The IP address of the member. Changing this creates a
// new member.
func (o MemberV1Output) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The administrative state of the member.
// Acceptable values are 'true' and 'false'. Changing this value updates the
// state of the existing member.
func (o MemberV1Output) AdminStateUp() pulumi.BoolOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.BoolOutput { return v.AdminStateUp }).(pulumi.BoolOutput)
}

// The ID of the LB pool. Changing this creates a new
// member.
func (o MemberV1Output) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

// An integer representing the port on which the member is
// hosted. Changing this creates a new member.
func (o MemberV1Output) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an LB member. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// LB member.
func (o MemberV1Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The owner of the member. Required if admin wants to
// create a member for another tenant. Changing this creates a new member.
func (o MemberV1Output) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o MemberV1Output) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v *MemberV1) pulumi.IntOutput { return v.Weight }).(pulumi.IntOutput)
}

type MemberV1ArrayOutput struct{ *pulumi.OutputState }

func (MemberV1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MemberV1)(nil)).Elem()
}

func (o MemberV1ArrayOutput) ToMemberV1ArrayOutput() MemberV1ArrayOutput {
	return o
}

func (o MemberV1ArrayOutput) ToMemberV1ArrayOutputWithContext(ctx context.Context) MemberV1ArrayOutput {
	return o
}

func (o MemberV1ArrayOutput) Index(i pulumi.IntInput) MemberV1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MemberV1 {
		return vs[0].([]*MemberV1)[vs[1].(int)]
	}).(MemberV1Output)
}

type MemberV1MapOutput struct{ *pulumi.OutputState }

func (MemberV1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MemberV1)(nil)).Elem()
}

func (o MemberV1MapOutput) ToMemberV1MapOutput() MemberV1MapOutput {
	return o
}

func (o MemberV1MapOutput) ToMemberV1MapOutputWithContext(ctx context.Context) MemberV1MapOutput {
	return o
}

func (o MemberV1MapOutput) MapIndex(k pulumi.StringInput) MemberV1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MemberV1 {
		return vs[0].(map[string]*MemberV1)[vs[1].(string)]
	}).(MemberV1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MemberV1Input)(nil)).Elem(), &MemberV1{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberV1ArrayInput)(nil)).Elem(), MemberV1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberV1MapInput)(nil)).Elem(), MemberV1Map{})
	pulumi.RegisterOutputType(MemberV1Output{})
	pulumi.RegisterOutputType(MemberV1ArrayOutput{})
	pulumi.RegisterOutputType(MemberV1MapOutput{})
}
