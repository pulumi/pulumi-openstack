// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V2 security group resource within OpenStack.
//
// Please note that managing security groups through the OpenStack Compute API
// has been deprecated. Unless you are using an older OpenStack environment, it is
// recommended to use the `networking.SecGroup`
// and `networking.SecGroupRule`
// resources instead, which uses the OpenStack Networking API.
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
//			_, err := compute.NewSecGroup(ctx, "secgroup1", &compute.SecGroupArgs{
//				Description: pulumi.String("my security group"),
//				Rules: compute.SecGroupRuleArray{
//					&compute.SecGroupRuleArgs{
//						Cidr:       pulumi.String("0.0.0.0/0"),
//						FromPort:   pulumi.Int(22),
//						IpProtocol: pulumi.String("tcp"),
//						ToPort:     pulumi.Int(22),
//					},
//					&compute.SecGroupRuleArgs{
//						Cidr:       pulumi.String("0.0.0.0/0"),
//						FromPort:   pulumi.Int(80),
//						IpProtocol: pulumi.String("tcp"),
//						ToPort:     pulumi.Int(80),
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
// ## Notes
//
// ### ICMP Rules
//
// When using ICMP as the `ipProtocol`, the `fromPort` sets the ICMP _type_ and the `toPort` sets the ICMP _code_. To allow all ICMP types, set each value to `-1`, like so:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
//
// A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages).
//
// ### Referencing Security Groups
//
// When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:
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
//			_, err := compute.NewInstance(ctx, "test-server", &compute.InstanceArgs{
//				ImageId:  pulumi.String("ad091b52-742f-469e-8f3c-fd81cadf0743"),
//				FlavorId: pulumi.String("3"),
//				KeyPair:  pulumi.String("my_key_pair_name"),
//				SecurityGroups: pulumi.StringArray{
//					openstack_compute_secgroup_v2.Secgroup_1.Name,
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
// Security Groups can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:compute/secGroup:SecGroup my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
//
// ```
type SecGroup struct {
	pulumi.CustomResourceState

	// A description for the security group. Changing this
	// updates the `description` of an existing security group.
	Description pulumi.StringOutput `pulumi:"description"`
	// A unique name for the security group. Changing this
	// updates the `name` of an existing security group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a security group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region pulumi.StringOutput `pulumi:"region"`
	// A rule describing how the security group operates. The
	// rule object structure is documented below. Changing this updates the
	// security group rules. As shown in the example above, multiple rule blocks
	// may be used.
	Rules SecGroupRuleArrayOutput `pulumi:"rules"`
}

// NewSecGroup registers a new resource with the given unique name, arguments, and options.
func NewSecGroup(ctx *pulumi.Context,
	name string, args *SecGroupArgs, opts ...pulumi.ResourceOption) (*SecGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecGroup
	err := ctx.RegisterResource("openstack:compute/secGroup:SecGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecGroup gets an existing SecGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecGroupState, opts ...pulumi.ResourceOption) (*SecGroup, error) {
	var resource SecGroup
	err := ctx.ReadResource("openstack:compute/secGroup:SecGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecGroup resources.
type secGroupState struct {
	// A description for the security group. Changing this
	// updates the `description` of an existing security group.
	Description *string `pulumi:"description"`
	// A unique name for the security group. Changing this
	// updates the `name` of an existing security group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a security group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region *string `pulumi:"region"`
	// A rule describing how the security group operates. The
	// rule object structure is documented below. Changing this updates the
	// security group rules. As shown in the example above, multiple rule blocks
	// may be used.
	Rules []SecGroupRule `pulumi:"rules"`
}

type SecGroupState struct {
	// A description for the security group. Changing this
	// updates the `description` of an existing security group.
	Description pulumi.StringPtrInput
	// A unique name for the security group. Changing this
	// updates the `name` of an existing security group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a security group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region pulumi.StringPtrInput
	// A rule describing how the security group operates. The
	// rule object structure is documented below. Changing this updates the
	// security group rules. As shown in the example above, multiple rule blocks
	// may be used.
	Rules SecGroupRuleArrayInput
}

func (SecGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*secGroupState)(nil)).Elem()
}

type secGroupArgs struct {
	// A description for the security group. Changing this
	// updates the `description` of an existing security group.
	Description string `pulumi:"description"`
	// A unique name for the security group. Changing this
	// updates the `name` of an existing security group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a security group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region *string `pulumi:"region"`
	// A rule describing how the security group operates. The
	// rule object structure is documented below. Changing this updates the
	// security group rules. As shown in the example above, multiple rule blocks
	// may be used.
	Rules []SecGroupRule `pulumi:"rules"`
}

// The set of arguments for constructing a SecGroup resource.
type SecGroupArgs struct {
	// A description for the security group. Changing this
	// updates the `description` of an existing security group.
	Description pulumi.StringInput
	// A unique name for the security group. Changing this
	// updates the `name` of an existing security group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a security group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security group.
	Region pulumi.StringPtrInput
	// A rule describing how the security group operates. The
	// rule object structure is documented below. Changing this updates the
	// security group rules. As shown in the example above, multiple rule blocks
	// may be used.
	Rules SecGroupRuleArrayInput
}

func (SecGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secGroupArgs)(nil)).Elem()
}

type SecGroupInput interface {
	pulumi.Input

	ToSecGroupOutput() SecGroupOutput
	ToSecGroupOutputWithContext(ctx context.Context) SecGroupOutput
}

func (*SecGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecGroup)(nil)).Elem()
}

func (i *SecGroup) ToSecGroupOutput() SecGroupOutput {
	return i.ToSecGroupOutputWithContext(context.Background())
}

func (i *SecGroup) ToSecGroupOutputWithContext(ctx context.Context) SecGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupOutput)
}

func (i *SecGroup) ToOutput(ctx context.Context) pulumix.Output[*SecGroup] {
	return pulumix.Output[*SecGroup]{
		OutputState: i.ToSecGroupOutputWithContext(ctx).OutputState,
	}
}

// SecGroupArrayInput is an input type that accepts SecGroupArray and SecGroupArrayOutput values.
// You can construct a concrete instance of `SecGroupArrayInput` via:
//
//	SecGroupArray{ SecGroupArgs{...} }
type SecGroupArrayInput interface {
	pulumi.Input

	ToSecGroupArrayOutput() SecGroupArrayOutput
	ToSecGroupArrayOutputWithContext(context.Context) SecGroupArrayOutput
}

type SecGroupArray []SecGroupInput

func (SecGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecGroup)(nil)).Elem()
}

func (i SecGroupArray) ToSecGroupArrayOutput() SecGroupArrayOutput {
	return i.ToSecGroupArrayOutputWithContext(context.Background())
}

func (i SecGroupArray) ToSecGroupArrayOutputWithContext(ctx context.Context) SecGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupArrayOutput)
}

func (i SecGroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*SecGroup] {
	return pulumix.Output[[]*SecGroup]{
		OutputState: i.ToSecGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// SecGroupMapInput is an input type that accepts SecGroupMap and SecGroupMapOutput values.
// You can construct a concrete instance of `SecGroupMapInput` via:
//
//	SecGroupMap{ "key": SecGroupArgs{...} }
type SecGroupMapInput interface {
	pulumi.Input

	ToSecGroupMapOutput() SecGroupMapOutput
	ToSecGroupMapOutputWithContext(context.Context) SecGroupMapOutput
}

type SecGroupMap map[string]SecGroupInput

func (SecGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecGroup)(nil)).Elem()
}

func (i SecGroupMap) ToSecGroupMapOutput() SecGroupMapOutput {
	return i.ToSecGroupMapOutputWithContext(context.Background())
}

func (i SecGroupMap) ToSecGroupMapOutputWithContext(ctx context.Context) SecGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecGroupMapOutput)
}

func (i SecGroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecGroup] {
	return pulumix.Output[map[string]*SecGroup]{
		OutputState: i.ToSecGroupMapOutputWithContext(ctx).OutputState,
	}
}

type SecGroupOutput struct{ *pulumi.OutputState }

func (SecGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecGroup)(nil)).Elem()
}

func (o SecGroupOutput) ToSecGroupOutput() SecGroupOutput {
	return o
}

func (o SecGroupOutput) ToSecGroupOutputWithContext(ctx context.Context) SecGroupOutput {
	return o
}

func (o SecGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*SecGroup] {
	return pulumix.Output[*SecGroup]{
		OutputState: o.OutputState,
	}
}

// A description for the security group. Changing this
// updates the `description` of an existing security group.
func (o SecGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A unique name for the security group. Changing this
// updates the `name` of an existing security group.
func (o SecGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Compute client.
// A Compute client is needed to create a security group. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// security group.
func (o SecGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SecGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A rule describing how the security group operates. The
// rule object structure is documented below. Changing this updates the
// security group rules. As shown in the example above, multiple rule blocks
// may be used.
func (o SecGroupOutput) Rules() SecGroupRuleArrayOutput {
	return o.ApplyT(func(v *SecGroup) SecGroupRuleArrayOutput { return v.Rules }).(SecGroupRuleArrayOutput)
}

type SecGroupArrayOutput struct{ *pulumi.OutputState }

func (SecGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecGroup)(nil)).Elem()
}

func (o SecGroupArrayOutput) ToSecGroupArrayOutput() SecGroupArrayOutput {
	return o
}

func (o SecGroupArrayOutput) ToSecGroupArrayOutputWithContext(ctx context.Context) SecGroupArrayOutput {
	return o
}

func (o SecGroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SecGroup] {
	return pulumix.Output[[]*SecGroup]{
		OutputState: o.OutputState,
	}
}

func (o SecGroupArrayOutput) Index(i pulumi.IntInput) SecGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecGroup {
		return vs[0].([]*SecGroup)[vs[1].(int)]
	}).(SecGroupOutput)
}

type SecGroupMapOutput struct{ *pulumi.OutputState }

func (SecGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecGroup)(nil)).Elem()
}

func (o SecGroupMapOutput) ToSecGroupMapOutput() SecGroupMapOutput {
	return o
}

func (o SecGroupMapOutput) ToSecGroupMapOutputWithContext(ctx context.Context) SecGroupMapOutput {
	return o
}

func (o SecGroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecGroup] {
	return pulumix.Output[map[string]*SecGroup]{
		OutputState: o.OutputState,
	}
}

func (o SecGroupMapOutput) MapIndex(k pulumi.StringInput) SecGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecGroup {
		return vs[0].(map[string]*SecGroup)[vs[1].(string)]
	}).(SecGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupInput)(nil)).Elem(), &SecGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupArrayInput)(nil)).Elem(), SecGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecGroupMapInput)(nil)).Elem(), SecGroupMap{})
	pulumi.RegisterOutputType(SecGroupOutput{})
	pulumi.RegisterOutputType(SecGroupArrayOutput{})
	pulumi.RegisterOutputType(SecGroupMapOutput{})
}
