// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Server Group resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewServerGroup(ctx, "test-sg", &compute.ServerGroupArgs{
//				Policies: pulumi.StringArray{
//					pulumi.String("anti-affinity"),
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
// ## Policies
//
//   - `affinity` - All instances/servers launched in this group will be hosted on
//     the same compute node.
//
//   - `anti-affinity` - All instances/servers launched in this group will be
//     hosted on different compute nodes.
//
//   - `soft-affinity` - All instances/servers launched in this group will be hosted
//     on the same compute node if possible, but if not possible they
//     still will be scheduled instead of failure. To use this policy your
//     OpenStack environment should support Compute service API 2.15 or above.
//
//   - `soft-anti-affinity` - All instances/servers launched in this group will be
//     hosted on different compute nodes if possible, but if not possible they
//     still will be scheduled instead of failure. To use this policy your
//     OpenStack environment should support Compute service API 2.15 or above.
//
// ## Import
//
// Server Groups can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:compute/serverGroup:ServerGroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
//
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// The instances that are part of this server group.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		args = &ServerGroupArgs{}
	}

	var resource ServerGroup
	err := ctx.RegisterResource("openstack:compute/serverGroup:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("openstack:compute/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// The instances that are part of this server group.
	Members []string `pulumi:"members"`
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name *string `pulumi:"name"`
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies []string `pulumi:"policies"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region *string `pulumi:"region"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type ServerGroupState struct {
	// The instances that are part of this server group.
	Members pulumi.StringArrayInput
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name pulumi.StringPtrInput
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies pulumi.StringArrayInput
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name *string `pulumi:"name"`
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies []string `pulumi:"policies"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region *string `pulumi:"region"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name pulumi.StringPtrInput
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies pulumi.StringArrayInput
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}

type ServerGroupInput interface {
	pulumi.Input

	ToServerGroupOutput() ServerGroupOutput
	ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput
}

func (*ServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

// ServerGroupArrayInput is an input type that accepts ServerGroupArray and ServerGroupArrayOutput values.
// You can construct a concrete instance of `ServerGroupArrayInput` via:
//
//	ServerGroupArray{ ServerGroupArgs{...} }
type ServerGroupArrayInput interface {
	pulumi.Input

	ToServerGroupArrayOutput() ServerGroupArrayOutput
	ToServerGroupArrayOutputWithContext(context.Context) ServerGroupArrayOutput
}

type ServerGroupArray []ServerGroupInput

func (ServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupArray) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return i.ToServerGroupArrayOutputWithContext(context.Background())
}

func (i ServerGroupArray) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupArrayOutput)
}

// ServerGroupMapInput is an input type that accepts ServerGroupMap and ServerGroupMapOutput values.
// You can construct a concrete instance of `ServerGroupMapInput` via:
//
//	ServerGroupMap{ "key": ServerGroupArgs{...} }
type ServerGroupMapInput interface {
	pulumi.Input

	ToServerGroupMapOutput() ServerGroupMapOutput
	ToServerGroupMapOutputWithContext(context.Context) ServerGroupMapOutput
}

type ServerGroupMap map[string]ServerGroupInput

func (ServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupMap) ToServerGroupMapOutput() ServerGroupMapOutput {
	return i.ToServerGroupMapOutputWithContext(context.Background())
}

func (i ServerGroupMap) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupMapOutput)
}

type ServerGroupOutput struct{ *pulumi.OutputState }

func (ServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

// The instances that are part of this server group.
func (o ServerGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// A unique name for the server group. Changing this creates
// a new server group.
func (o ServerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of policies for the server group. All policies
// are mutually exclusive. See the Policies section for more information.
// Changing this creates a new server group.
func (o ServerGroupOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// The region in which to obtain the V2 Compute client.
// If omitted, the `region` argument of the provider is used. Changing
// this creates a new server group.
func (o ServerGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of additional options.
func (o ServerGroupOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type ServerGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) Index(i pulumi.IntInput) ServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].([]*ServerGroup)[vs[1].(int)]
	}).(ServerGroupOutput)
}

type ServerGroupMapOutput struct{ *pulumi.OutputState }

func (ServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupMapOutput) ToServerGroupMapOutput() ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) MapIndex(k pulumi.StringInput) ServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].(map[string]*ServerGroup)[vs[1].(string)]
	}).(ServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupArrayInput)(nil)).Elem(), ServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupMapInput)(nil)).Elem(), ServerGroupMap{})
	pulumi.RegisterOutputType(ServerGroupOutput{})
	pulumi.RegisterOutputType(ServerGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupMapOutput{})
}
