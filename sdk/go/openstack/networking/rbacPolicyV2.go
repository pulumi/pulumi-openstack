// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The RBAC policy resource contains functionality for working with Neutron RBAC
// Policies. Role-Based Access Control (RBAC) policy framework enables both
// operators and users to grant access to resources for specific projects.
//
// Sharing an object with a specific project is accomplished by creating a
// policy entry that permits the target project the `accessAsShared` action
// on that object.
//
// To make a network available as an external network for specific projects
// rather than all projects, use the `accessAsExternal` action.
// If a network is marked as external during creation, it now implicitly creates
// a wildcard RBAC policy granting everyone access to preserve previous behavior
// before this feature was added.
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
//			network1, err := networking.NewNetwork(ctx, "network_1", &networking.NetworkArgs{
//				Name:         pulumi.String("network_1"),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewRbacPolicyV2(ctx, "rbac_policy_1", &networking.RbacPolicyV2Args{
//				Action:       pulumi.String("access_as_shared"),
//				ObjectId:     network1.ID(),
//				ObjectType:   pulumi.String("network"),
//				TargetTenant: pulumi.String("20415a973c9e45d3917f078950644697"),
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
// RBAC policies can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:networking/rbacPolicyV2:RbacPolicyV2 rbac_policy_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
// ```
type RbacPolicyV2 struct {
	pulumi.CustomResourceState

	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The type of the object that the RBAC policy
	// affects. Can be one of the following: `addressScope`, `addressGroup`,
	// `network`, `qosPolicy`, `securityGroup`, `subnetpool` or `bgpvpn`.
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	ProjectId  pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant pulumi.StringOutput `pulumi:"targetTenant"`
}

// NewRbacPolicyV2 registers a new resource with the given unique name, arguments, and options.
func NewRbacPolicyV2(ctx *pulumi.Context,
	name string, args *RbacPolicyV2Args, opts ...pulumi.ResourceOption) (*RbacPolicyV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ObjectId == nil {
		return nil, errors.New("invalid value for required argument 'ObjectId'")
	}
	if args.ObjectType == nil {
		return nil, errors.New("invalid value for required argument 'ObjectType'")
	}
	if args.TargetTenant == nil {
		return nil, errors.New("invalid value for required argument 'TargetTenant'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RbacPolicyV2
	err := ctx.RegisterResource("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRbacPolicyV2 gets an existing RbacPolicyV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRbacPolicyV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RbacPolicyV2State, opts ...pulumi.ResourceOption) (*RbacPolicyV2, error) {
	var resource RbacPolicyV2
	err := ctx.ReadResource("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RbacPolicyV2 resources.
type rbacPolicyV2State struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action *string `pulumi:"action"`
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId *string `pulumi:"objectId"`
	// The type of the object that the RBAC policy
	// affects. Can be one of the following: `addressScope`, `addressGroup`,
	// `network`, `qosPolicy`, `securityGroup`, `subnetpool` or `bgpvpn`.
	ObjectType *string `pulumi:"objectType"`
	ProjectId  *string `pulumi:"projectId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant *string `pulumi:"targetTenant"`
}

type RbacPolicyV2State struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action pulumi.StringPtrInput
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId pulumi.StringPtrInput
	// The type of the object that the RBAC policy
	// affects. Can be one of the following: `addressScope`, `addressGroup`,
	// `network`, `qosPolicy`, `securityGroup`, `subnetpool` or `bgpvpn`.
	ObjectType pulumi.StringPtrInput
	ProjectId  pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant pulumi.StringPtrInput
}

func (RbacPolicyV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacPolicyV2State)(nil)).Elem()
}

type rbacPolicyV2Args struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action string `pulumi:"action"`
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId string `pulumi:"objectId"`
	// The type of the object that the RBAC policy
	// affects. Can be one of the following: `addressScope`, `addressGroup`,
	// `network`, `qosPolicy`, `securityGroup`, `subnetpool` or `bgpvpn`.
	ObjectType string `pulumi:"objectType"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant string `pulumi:"targetTenant"`
}

// The set of arguments for constructing a RbacPolicyV2 resource.
type RbacPolicyV2Args struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action pulumi.StringInput
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId pulumi.StringInput
	// The type of the object that the RBAC policy
	// affects. Can be one of the following: `addressScope`, `addressGroup`,
	// `network`, `qosPolicy`, `securityGroup`, `subnetpool` or `bgpvpn`.
	ObjectType pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant pulumi.StringInput
}

func (RbacPolicyV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacPolicyV2Args)(nil)).Elem()
}

type RbacPolicyV2Input interface {
	pulumi.Input

	ToRbacPolicyV2Output() RbacPolicyV2Output
	ToRbacPolicyV2OutputWithContext(ctx context.Context) RbacPolicyV2Output
}

func (*RbacPolicyV2) ElementType() reflect.Type {
	return reflect.TypeOf((**RbacPolicyV2)(nil)).Elem()
}

func (i *RbacPolicyV2) ToRbacPolicyV2Output() RbacPolicyV2Output {
	return i.ToRbacPolicyV2OutputWithContext(context.Background())
}

func (i *RbacPolicyV2) ToRbacPolicyV2OutputWithContext(ctx context.Context) RbacPolicyV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(RbacPolicyV2Output)
}

// RbacPolicyV2ArrayInput is an input type that accepts RbacPolicyV2Array and RbacPolicyV2ArrayOutput values.
// You can construct a concrete instance of `RbacPolicyV2ArrayInput` via:
//
//	RbacPolicyV2Array{ RbacPolicyV2Args{...} }
type RbacPolicyV2ArrayInput interface {
	pulumi.Input

	ToRbacPolicyV2ArrayOutput() RbacPolicyV2ArrayOutput
	ToRbacPolicyV2ArrayOutputWithContext(context.Context) RbacPolicyV2ArrayOutput
}

type RbacPolicyV2Array []RbacPolicyV2Input

func (RbacPolicyV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RbacPolicyV2)(nil)).Elem()
}

func (i RbacPolicyV2Array) ToRbacPolicyV2ArrayOutput() RbacPolicyV2ArrayOutput {
	return i.ToRbacPolicyV2ArrayOutputWithContext(context.Background())
}

func (i RbacPolicyV2Array) ToRbacPolicyV2ArrayOutputWithContext(ctx context.Context) RbacPolicyV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacPolicyV2ArrayOutput)
}

// RbacPolicyV2MapInput is an input type that accepts RbacPolicyV2Map and RbacPolicyV2MapOutput values.
// You can construct a concrete instance of `RbacPolicyV2MapInput` via:
//
//	RbacPolicyV2Map{ "key": RbacPolicyV2Args{...} }
type RbacPolicyV2MapInput interface {
	pulumi.Input

	ToRbacPolicyV2MapOutput() RbacPolicyV2MapOutput
	ToRbacPolicyV2MapOutputWithContext(context.Context) RbacPolicyV2MapOutput
}

type RbacPolicyV2Map map[string]RbacPolicyV2Input

func (RbacPolicyV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RbacPolicyV2)(nil)).Elem()
}

func (i RbacPolicyV2Map) ToRbacPolicyV2MapOutput() RbacPolicyV2MapOutput {
	return i.ToRbacPolicyV2MapOutputWithContext(context.Background())
}

func (i RbacPolicyV2Map) ToRbacPolicyV2MapOutputWithContext(ctx context.Context) RbacPolicyV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RbacPolicyV2MapOutput)
}

type RbacPolicyV2Output struct{ *pulumi.OutputState }

func (RbacPolicyV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RbacPolicyV2)(nil)).Elem()
}

func (o RbacPolicyV2Output) ToRbacPolicyV2Output() RbacPolicyV2Output {
	return o
}

func (o RbacPolicyV2Output) ToRbacPolicyV2OutputWithContext(ctx context.Context) RbacPolicyV2Output {
	return o
}

// Action for the RBAC policy. Can either be
// `accessAsExternal` or `accessAsShared`.
func (o RbacPolicyV2Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPolicyV2) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The ID of the `objectType` resource. An
// `objectType` of `network` returns a network ID and an `objectType` of
// `qosPolicy` returns a QoS ID.
func (o RbacPolicyV2Output) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPolicyV2) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

// The type of the object that the RBAC policy
// affects. Can be one of the following: `addressScope`, `addressGroup`,
// `network`, `qosPolicy`, `securityGroup`, `subnetpool` or `bgpvpn`.
func (o RbacPolicyV2Output) ObjectType() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPolicyV2) pulumi.StringOutput { return v.ObjectType }).(pulumi.StringOutput)
}

func (o RbacPolicyV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPolicyV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to configure a routing entry on a subnet. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// routing entry.
func (o RbacPolicyV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPolicyV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the tenant to which the RBAC policy
// will be enforced.
func (o RbacPolicyV2Output) TargetTenant() pulumi.StringOutput {
	return o.ApplyT(func(v *RbacPolicyV2) pulumi.StringOutput { return v.TargetTenant }).(pulumi.StringOutput)
}

type RbacPolicyV2ArrayOutput struct{ *pulumi.OutputState }

func (RbacPolicyV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RbacPolicyV2)(nil)).Elem()
}

func (o RbacPolicyV2ArrayOutput) ToRbacPolicyV2ArrayOutput() RbacPolicyV2ArrayOutput {
	return o
}

func (o RbacPolicyV2ArrayOutput) ToRbacPolicyV2ArrayOutputWithContext(ctx context.Context) RbacPolicyV2ArrayOutput {
	return o
}

func (o RbacPolicyV2ArrayOutput) Index(i pulumi.IntInput) RbacPolicyV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RbacPolicyV2 {
		return vs[0].([]*RbacPolicyV2)[vs[1].(int)]
	}).(RbacPolicyV2Output)
}

type RbacPolicyV2MapOutput struct{ *pulumi.OutputState }

func (RbacPolicyV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RbacPolicyV2)(nil)).Elem()
}

func (o RbacPolicyV2MapOutput) ToRbacPolicyV2MapOutput() RbacPolicyV2MapOutput {
	return o
}

func (o RbacPolicyV2MapOutput) ToRbacPolicyV2MapOutputWithContext(ctx context.Context) RbacPolicyV2MapOutput {
	return o
}

func (o RbacPolicyV2MapOutput) MapIndex(k pulumi.StringInput) RbacPolicyV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RbacPolicyV2 {
		return vs[0].(map[string]*RbacPolicyV2)[vs[1].(string)]
	}).(RbacPolicyV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RbacPolicyV2Input)(nil)).Elem(), &RbacPolicyV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*RbacPolicyV2ArrayInput)(nil)).Elem(), RbacPolicyV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RbacPolicyV2MapInput)(nil)).Elem(), RbacPolicyV2Map{})
	pulumi.RegisterOutputType(RbacPolicyV2Output{})
	pulumi.RegisterOutputType(RbacPolicyV2ArrayOutput{})
	pulumi.RegisterOutputType(RbacPolicyV2MapOutput{})
}
