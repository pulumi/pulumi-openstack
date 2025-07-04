// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 members resource within OpenStack (batch members update).
//
// > **Note:** This resource has attributes that depend on octavia minor versions.
// Please ensure your Openstack cloud supports the required minor version.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/loadbalancer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := loadbalancer.NewMembers(ctx, "members_1", &loadbalancer.MembersArgs{
//				PoolId: pulumi.String("935685fb-a896-40f9-9ff4-ae531a3a00fe"),
//				Members: loadbalancer.MembersMemberArray{
//					&loadbalancer.MembersMemberArgs{
//						Address:      pulumi.String("192.168.199.23"),
//						ProtocolPort: pulumi.Int(8080),
//					},
//					&loadbalancer.MembersMemberArgs{
//						Address:      pulumi.String("192.168.199.24"),
//						ProtocolPort: pulumi.Int(8080),
//					},
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
//
// ## Import
//
// Load Balancer Pool Members can be imported using the Pool ID, e.g.:
//
// ```sh
// $ pulumi import openstack:loadbalancer/members:Members members_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5
// ```
type Members struct {
	pulumi.CustomResourceState

	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Members MembersMemberArrayOutput `pulumi:"members"`
	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// members resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewMembers registers a new resource with the given unique name, arguments, and options.
func NewMembers(ctx *pulumi.Context,
	name string, args *MembersArgs, opts ...pulumi.ResourceOption) (*Members, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PoolId == nil {
		return nil, errors.New("invalid value for required argument 'PoolId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Members
	err := ctx.RegisterResource("openstack:loadbalancer/members:Members", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembers gets an existing Members resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembersState, opts ...pulumi.ResourceOption) (*Members, error) {
	var resource Members
	err := ctx.ReadResource("openstack:loadbalancer/members:Members", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Members resources.
type membersState struct {
	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Members []MembersMember `pulumi:"members"`
	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolId *string `pulumi:"poolId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// members resource.
	Region *string `pulumi:"region"`
}

type MembersState struct {
	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Members MembersMemberArrayInput
	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// members resource.
	Region pulumi.StringPtrInput
}

func (MembersState) ElementType() reflect.Type {
	return reflect.TypeOf((*membersState)(nil)).Elem()
}

type membersArgs struct {
	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Members []MembersMember `pulumi:"members"`
	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolId string `pulumi:"poolId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// members resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Members resource.
type MembersArgs struct {
	// A set of dictionaries containing member parameters. The
	// structure is described below.
	Members MembersMemberArrayInput
	// The id of the pool that members will be assigned to.
	// Changing this creates a new members resource.
	PoolId pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create pool members. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// members resource.
	Region pulumi.StringPtrInput
}

func (MembersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membersArgs)(nil)).Elem()
}

type MembersInput interface {
	pulumi.Input

	ToMembersOutput() MembersOutput
	ToMembersOutputWithContext(ctx context.Context) MembersOutput
}

func (*Members) ElementType() reflect.Type {
	return reflect.TypeOf((**Members)(nil)).Elem()
}

func (i *Members) ToMembersOutput() MembersOutput {
	return i.ToMembersOutputWithContext(context.Background())
}

func (i *Members) ToMembersOutputWithContext(ctx context.Context) MembersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembersOutput)
}

// MembersArrayInput is an input type that accepts MembersArray and MembersArrayOutput values.
// You can construct a concrete instance of `MembersArrayInput` via:
//
//	MembersArray{ MembersArgs{...} }
type MembersArrayInput interface {
	pulumi.Input

	ToMembersArrayOutput() MembersArrayOutput
	ToMembersArrayOutputWithContext(context.Context) MembersArrayOutput
}

type MembersArray []MembersInput

func (MembersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Members)(nil)).Elem()
}

func (i MembersArray) ToMembersArrayOutput() MembersArrayOutput {
	return i.ToMembersArrayOutputWithContext(context.Background())
}

func (i MembersArray) ToMembersArrayOutputWithContext(ctx context.Context) MembersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembersArrayOutput)
}

// MembersMapInput is an input type that accepts MembersMap and MembersMapOutput values.
// You can construct a concrete instance of `MembersMapInput` via:
//
//	MembersMap{ "key": MembersArgs{...} }
type MembersMapInput interface {
	pulumi.Input

	ToMembersMapOutput() MembersMapOutput
	ToMembersMapOutputWithContext(context.Context) MembersMapOutput
}

type MembersMap map[string]MembersInput

func (MembersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Members)(nil)).Elem()
}

func (i MembersMap) ToMembersMapOutput() MembersMapOutput {
	return i.ToMembersMapOutputWithContext(context.Background())
}

func (i MembersMap) ToMembersMapOutputWithContext(ctx context.Context) MembersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembersMapOutput)
}

type MembersOutput struct{ *pulumi.OutputState }

func (MembersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Members)(nil)).Elem()
}

func (o MembersOutput) ToMembersOutput() MembersOutput {
	return o
}

func (o MembersOutput) ToMembersOutputWithContext(ctx context.Context) MembersOutput {
	return o
}

// A set of dictionaries containing member parameters. The
// structure is described below.
func (o MembersOutput) Members() MembersMemberArrayOutput {
	return o.ApplyT(func(v *Members) MembersMemberArrayOutput { return v.Members }).(MembersMemberArrayOutput)
}

// The id of the pool that members will be assigned to.
// Changing this creates a new members resource.
func (o MembersOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Members) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create pool members. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// members resource.
func (o MembersOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Members) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type MembersArrayOutput struct{ *pulumi.OutputState }

func (MembersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Members)(nil)).Elem()
}

func (o MembersArrayOutput) ToMembersArrayOutput() MembersArrayOutput {
	return o
}

func (o MembersArrayOutput) ToMembersArrayOutputWithContext(ctx context.Context) MembersArrayOutput {
	return o
}

func (o MembersArrayOutput) Index(i pulumi.IntInput) MembersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Members {
		return vs[0].([]*Members)[vs[1].(int)]
	}).(MembersOutput)
}

type MembersMapOutput struct{ *pulumi.OutputState }

func (MembersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Members)(nil)).Elem()
}

func (o MembersMapOutput) ToMembersMapOutput() MembersMapOutput {
	return o
}

func (o MembersMapOutput) ToMembersMapOutputWithContext(ctx context.Context) MembersMapOutput {
	return o
}

func (o MembersMapOutput) MapIndex(k pulumi.StringInput) MembersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Members {
		return vs[0].(map[string]*Members)[vs[1].(string)]
	}).(MembersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembersInput)(nil)).Elem(), &Members{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembersArrayInput)(nil)).Elem(), MembersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembersMapInput)(nil)).Elem(), MembersMap{})
	pulumi.RegisterOutputType(MembersOutput{})
	pulumi.RegisterOutputType(MembersArrayOutput{})
	pulumi.RegisterOutputType(MembersMapOutput{})
}
