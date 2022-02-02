// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 member resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loadbalancer.NewMember(ctx, "member1", &loadbalancer.MemberArgs{
// 			Address:      pulumi.String("192.168.199.23"),
// 			PoolId:       pulumi.String("935685fb-a896-40f9-9ff4-ae531a3a00fe"),
// 			ProtocolPort: pulumi.Int(8080),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Load Balancer Pool Member can be imported using the Pool ID and Member ID separated by a slash, e.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/member:Member member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f
// ```
type Member struct {
	pulumi.CustomResourceState

	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address pulumi.StringOutput `pulumi:"address"`
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// Human-readable name for the member.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort pulumi.IntOutput `pulumi:"protocolPort"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new member.
	Region pulumi.StringOutput `pulumi:"region"`
	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight pulumi.IntOutput `pulumi:"weight"`
}

// NewMember registers a new resource with the given unique name, arguments, and options.
func NewMember(ctx *pulumi.Context,
	name string, args *MemberArgs, opts ...pulumi.ResourceOption) (*Member, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.PoolId == nil {
		return nil, errors.New("invalid value for required argument 'PoolId'")
	}
	if args.ProtocolPort == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolPort'")
	}
	var resource Member
	err := ctx.RegisterResource("openstack:loadbalancer/member:Member", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMember gets an existing Member resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberState, opts ...pulumi.ResourceOption) (*Member, error) {
	var resource Member
	err := ctx.ReadResource("openstack:loadbalancer/member:Member", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Member resources.
type memberState struct {
	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address *string `pulumi:"address"`
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable name for the member.
	Name *string `pulumi:"name"`
	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolId *string `pulumi:"poolId"`
	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort *int `pulumi:"protocolPort"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new member.
	Region *string `pulumi:"region"`
	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetId *string `pulumi:"subnetId"`
	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantId *string `pulumi:"tenantId"`
	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *int `pulumi:"weight"`
}

type MemberState struct {
	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address pulumi.StringPtrInput
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable name for the member.
	Name pulumi.StringPtrInput
	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolId pulumi.StringPtrInput
	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort pulumi.IntPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new member.
	Region pulumi.StringPtrInput
	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetId pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantId pulumi.StringPtrInput
	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight pulumi.IntPtrInput
}

func (MemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*memberState)(nil)).Elem()
}

type memberArgs struct {
	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address string `pulumi:"address"`
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable name for the member.
	Name *string `pulumi:"name"`
	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolId string `pulumi:"poolId"`
	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort int `pulumi:"protocolPort"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new member.
	Region *string `pulumi:"region"`
	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetId *string `pulumi:"subnetId"`
	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantId *string `pulumi:"tenantId"`
	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Member resource.
type MemberArgs struct {
	// The IP address of the member to receive traffic from
	// the load balancer. Changing this creates a new member.
	Address pulumi.StringInput
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable name for the member.
	Name pulumi.StringPtrInput
	// The id of the pool that this member will be assigned
	// to. Changing this creates a new member.
	PoolId pulumi.StringInput
	// The port on which to listen for client traffic.
	// Changing this creates a new member.
	ProtocolPort pulumi.IntInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a member. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new member.
	Region pulumi.StringPtrInput
	// The subnet in which to access the member. Changing
	// this creates a new member.
	SubnetId pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the member.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new member.
	TenantId pulumi.StringPtrInput
	// A positive integer value that indicates the relative
	// portion of traffic that this member should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight pulumi.IntPtrInput
}

func (MemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*memberArgs)(nil)).Elem()
}

type MemberInput interface {
	pulumi.Input

	ToMemberOutput() MemberOutput
	ToMemberOutputWithContext(ctx context.Context) MemberOutput
}

func (*Member) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (i *Member) ToMemberOutput() MemberOutput {
	return i.ToMemberOutputWithContext(context.Background())
}

func (i *Member) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberOutput)
}

// MemberArrayInput is an input type that accepts MemberArray and MemberArrayOutput values.
// You can construct a concrete instance of `MemberArrayInput` via:
//
//          MemberArray{ MemberArgs{...} }
type MemberArrayInput interface {
	pulumi.Input

	ToMemberArrayOutput() MemberArrayOutput
	ToMemberArrayOutputWithContext(context.Context) MemberArrayOutput
}

type MemberArray []MemberInput

func (MemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (i MemberArray) ToMemberArrayOutput() MemberArrayOutput {
	return i.ToMemberArrayOutputWithContext(context.Background())
}

func (i MemberArray) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberArrayOutput)
}

// MemberMapInput is an input type that accepts MemberMap and MemberMapOutput values.
// You can construct a concrete instance of `MemberMapInput` via:
//
//          MemberMap{ "key": MemberArgs{...} }
type MemberMapInput interface {
	pulumi.Input

	ToMemberMapOutput() MemberMapOutput
	ToMemberMapOutputWithContext(context.Context) MemberMapOutput
}

type MemberMap map[string]MemberInput

func (MemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (i MemberMap) ToMemberMapOutput() MemberMapOutput {
	return i.ToMemberMapOutputWithContext(context.Background())
}

func (i MemberMap) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MemberMapOutput)
}

type MemberOutput struct{ *pulumi.OutputState }

func (MemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Member)(nil)).Elem()
}

func (o MemberOutput) ToMemberOutput() MemberOutput {
	return o
}

func (o MemberOutput) ToMemberOutputWithContext(ctx context.Context) MemberOutput {
	return o
}

type MemberArrayOutput struct{ *pulumi.OutputState }

func (MemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Member)(nil)).Elem()
}

func (o MemberArrayOutput) ToMemberArrayOutput() MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) ToMemberArrayOutputWithContext(ctx context.Context) MemberArrayOutput {
	return o
}

func (o MemberArrayOutput) Index(i pulumi.IntInput) MemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Member {
		return vs[0].([]*Member)[vs[1].(int)]
	}).(MemberOutput)
}

type MemberMapOutput struct{ *pulumi.OutputState }

func (MemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Member)(nil)).Elem()
}

func (o MemberMapOutput) ToMemberMapOutput() MemberMapOutput {
	return o
}

func (o MemberMapOutput) ToMemberMapOutputWithContext(ctx context.Context) MemberMapOutput {
	return o
}

func (o MemberMapOutput) MapIndex(k pulumi.StringInput) MemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Member {
		return vs[0].(map[string]*Member)[vs[1].(string)]
	}).(MemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MemberInput)(nil)).Elem(), &Member{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberArrayInput)(nil)).Elem(), MemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MemberMapInput)(nil)).Elem(), MemberMap{})
	pulumi.RegisterOutputType(MemberOutput{})
	pulumi.RegisterOutputType(MemberArrayOutput{})
	pulumi.RegisterOutputType(MemberMapOutput{})
}
