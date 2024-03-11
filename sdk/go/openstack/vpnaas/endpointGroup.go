// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron Endpoint Group resource within OpenStack.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/vpnaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpnaas.NewEndpointGroup(ctx, "group1", &vpnaas.EndpointGroupArgs{
//				Endpoints: pulumi.StringArray{
//					pulumi.String("10.2.0.0/24"),
//					pulumi.String("10.3.0.0/24"),
//				},
//				Type: pulumi.String("cidr"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Groups can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:vpnaas/endpointGroup:EndpointGroup group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
// ```
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints pulumi.StringArrayOutput `pulumi:"endpoints"`
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type pulumi.StringOutput `pulumi:"type"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil {
		args = &EndpointGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointGroup
	err := ctx.RegisterResource("openstack:vpnaas/endpointGroup:EndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	var resource EndpointGroup
	err := ctx.ReadResource("openstack:vpnaas/endpointGroup:EndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type endpointGroupState struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description *string `pulumi:"description"`
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints []string `pulumi:"endpoints"`
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region *string `pulumi:"region"`
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId *string `pulumi:"tenantId"`
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type *string `pulumi:"type"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type EndpointGroupState struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description pulumi.StringPtrInput
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints pulumi.StringArrayInput
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region pulumi.StringPtrInput
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId pulumi.StringPtrInput
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (EndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupState)(nil)).Elem()
}

type endpointGroupArgs struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description *string `pulumi:"description"`
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints []string `pulumi:"endpoints"`
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region *string `pulumi:"region"`
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId *string `pulumi:"tenantId"`
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type *string `pulumi:"type"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The human-readable description for the group.
	// Changing this updates the description of the existing group.
	Description pulumi.StringPtrInput
	// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
	// Changing this creates a new group.
	Endpoints pulumi.StringArrayInput
	// The name of the group. Changing this updates the name of
	// the existing group.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an endpoint group. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// group.
	Region pulumi.StringPtrInput
	// The owner of the group. Required if admin wants to
	// create an endpoint group for another project. Changing this creates a new group.
	TenantId pulumi.StringPtrInput
	// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
	// Changing this creates a new group.
	Type pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (EndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupArgs)(nil)).Elem()
}

type EndpointGroupInput interface {
	pulumi.Input

	ToEndpointGroupOutput() EndpointGroupOutput
	ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput
}

func (*EndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil)).Elem()
}

func (i *EndpointGroup) ToEndpointGroupOutput() EndpointGroupOutput {
	return i.ToEndpointGroupOutputWithContext(context.Background())
}

func (i *EndpointGroup) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupOutput)
}

// EndpointGroupArrayInput is an input type that accepts EndpointGroupArray and EndpointGroupArrayOutput values.
// You can construct a concrete instance of `EndpointGroupArrayInput` via:
//
//	EndpointGroupArray{ EndpointGroupArgs{...} }
type EndpointGroupArrayInput interface {
	pulumi.Input

	ToEndpointGroupArrayOutput() EndpointGroupArrayOutput
	ToEndpointGroupArrayOutputWithContext(context.Context) EndpointGroupArrayOutput
}

type EndpointGroupArray []EndpointGroupInput

func (EndpointGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointGroup)(nil)).Elem()
}

func (i EndpointGroupArray) ToEndpointGroupArrayOutput() EndpointGroupArrayOutput {
	return i.ToEndpointGroupArrayOutputWithContext(context.Background())
}

func (i EndpointGroupArray) ToEndpointGroupArrayOutputWithContext(ctx context.Context) EndpointGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupArrayOutput)
}

// EndpointGroupMapInput is an input type that accepts EndpointGroupMap and EndpointGroupMapOutput values.
// You can construct a concrete instance of `EndpointGroupMapInput` via:
//
//	EndpointGroupMap{ "key": EndpointGroupArgs{...} }
type EndpointGroupMapInput interface {
	pulumi.Input

	ToEndpointGroupMapOutput() EndpointGroupMapOutput
	ToEndpointGroupMapOutputWithContext(context.Context) EndpointGroupMapOutput
}

type EndpointGroupMap map[string]EndpointGroupInput

func (EndpointGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointGroup)(nil)).Elem()
}

func (i EndpointGroupMap) ToEndpointGroupMapOutput() EndpointGroupMapOutput {
	return i.ToEndpointGroupMapOutputWithContext(context.Background())
}

func (i EndpointGroupMap) ToEndpointGroupMapOutputWithContext(ctx context.Context) EndpointGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupMapOutput)
}

type EndpointGroupOutput struct{ *pulumi.OutputState }

func (EndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupOutput) ToEndpointGroupOutput() EndpointGroupOutput {
	return o
}

func (o EndpointGroupOutput) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return o
}

// The human-readable description for the group.
// Changing this updates the description of the existing group.
func (o EndpointGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of endpoints of the same type, for the endpoint group. The values will depend on the type.
// Changing this creates a new group.
func (o EndpointGroupOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringArrayOutput { return v.Endpoints }).(pulumi.StringArrayOutput)
}

// The name of the group. Changing this updates the name of
// the existing group.
func (o EndpointGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an endpoint group. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// group.
func (o EndpointGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The owner of the group. Required if admin wants to
// create an endpoint group for another project. Changing this creates a new group.
func (o EndpointGroupOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
// Changing this creates a new group.
func (o EndpointGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Map of additional options.
func (o EndpointGroupOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *EndpointGroup) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type EndpointGroupArrayOutput struct{ *pulumi.OutputState }

func (EndpointGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupArrayOutput) ToEndpointGroupArrayOutput() EndpointGroupArrayOutput {
	return o
}

func (o EndpointGroupArrayOutput) ToEndpointGroupArrayOutputWithContext(ctx context.Context) EndpointGroupArrayOutput {
	return o
}

func (o EndpointGroupArrayOutput) Index(i pulumi.IntInput) EndpointGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointGroup {
		return vs[0].([]*EndpointGroup)[vs[1].(int)]
	}).(EndpointGroupOutput)
}

type EndpointGroupMapOutput struct{ *pulumi.OutputState }

func (EndpointGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointGroup)(nil)).Elem()
}

func (o EndpointGroupMapOutput) ToEndpointGroupMapOutput() EndpointGroupMapOutput {
	return o
}

func (o EndpointGroupMapOutput) ToEndpointGroupMapOutputWithContext(ctx context.Context) EndpointGroupMapOutput {
	return o
}

func (o EndpointGroupMapOutput) MapIndex(k pulumi.StringInput) EndpointGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointGroup {
		return vs[0].(map[string]*EndpointGroup)[vs[1].(string)]
	}).(EndpointGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupInput)(nil)).Elem(), &EndpointGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupArrayInput)(nil)).Elem(), EndpointGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointGroupMapInput)(nil)).Elem(), EndpointGroupMap{})
	pulumi.RegisterOutputType(EndpointGroupOutput{})
	pulumi.RegisterOutputType(EndpointGroupArrayOutput{})
	pulumi.RegisterOutputType(EndpointGroupMapOutput{})
}
