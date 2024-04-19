// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 Barbican order resource within OpenStack.
//
// ## Example Usage
//
// ### Symmetric key order
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/keymanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keymanager.NewOrderV1(ctx, "order_1", &keymanager.OrderV1Args{
//				Type: pulumi.String("key"),
//				Meta: &keymanager.OrderV1MetaArgs{
//					Algorithm: pulumi.String("aes"),
//					BitLength: pulumi.Int(256),
//					Name:      pulumi.String("mysecret"),
//					Mode:      pulumi.String("cbc"),
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
// <!--End PulumiCodeChooser -->
//
// ### Asymmetric key pair order
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/keymanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keymanager.NewOrderV1(ctx, "order_1", &keymanager.OrderV1Args{
//				Type: pulumi.String("asymmetric"),
//				Meta: &keymanager.OrderV1MetaArgs{
//					Algorithm: pulumi.String("rsa"),
//					BitLength: pulumi.Int(4096),
//					Name:      pulumi.String("mysecret"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Orders can be imported using the order id (the last part of the order reference), e.g.:
//
// ```sh
// $ pulumi import openstack:keymanager/orderV1:OrderV1 order_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
// ```
type OrderV1 struct {
	pulumi.CustomResourceState

	// The container reference / where to find the container.
	ContainerRef pulumi.StringOutput `pulumi:"containerRef"`
	// The date the order was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The creator of the order.
	CreatorId pulumi.StringOutput `pulumi:"creatorId"`
	// Dictionary containing the order metadata used to generate the order. The structure is described below.
	Meta OrderV1MetaOutput `pulumi:"meta"`
	// The order reference / where to find the order.
	OrderRef pulumi.StringOutput `pulumi:"orderRef"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a order. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 order.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret reference / where to find the secret.
	SecretRef pulumi.StringOutput `pulumi:"secretRef"`
	// The status of the order.
	Status pulumi.StringOutput `pulumi:"status"`
	// The sub status of the order.
	SubStatus pulumi.StringOutput `pulumi:"subStatus"`
	// The sub status message of the order.
	SubStatusMessage pulumi.StringOutput `pulumi:"subStatusMessage"`
	// The type of key to be generated. Must be one of `asymmetric`, `key`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date the order was last updated.
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewOrderV1 registers a new resource with the given unique name, arguments, and options.
func NewOrderV1(ctx *pulumi.Context,
	name string, args *OrderV1Args, opts ...pulumi.ResourceOption) (*OrderV1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Meta == nil {
		return nil, errors.New("invalid value for required argument 'Meta'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrderV1
	err := ctx.RegisterResource("openstack:keymanager/orderV1:OrderV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrderV1 gets an existing OrderV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrderV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderV1State, opts ...pulumi.ResourceOption) (*OrderV1, error) {
	var resource OrderV1
	err := ctx.ReadResource("openstack:keymanager/orderV1:OrderV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrderV1 resources.
type orderV1State struct {
	// The container reference / where to find the container.
	ContainerRef *string `pulumi:"containerRef"`
	// The date the order was created.
	Created *string `pulumi:"created"`
	// The creator of the order.
	CreatorId *string `pulumi:"creatorId"`
	// Dictionary containing the order metadata used to generate the order. The structure is described below.
	Meta *OrderV1Meta `pulumi:"meta"`
	// The order reference / where to find the order.
	OrderRef *string `pulumi:"orderRef"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a order. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 order.
	Region *string `pulumi:"region"`
	// The secret reference / where to find the secret.
	SecretRef *string `pulumi:"secretRef"`
	// The status of the order.
	Status *string `pulumi:"status"`
	// The sub status of the order.
	SubStatus *string `pulumi:"subStatus"`
	// The sub status message of the order.
	SubStatusMessage *string `pulumi:"subStatusMessage"`
	// The type of key to be generated. Must be one of `asymmetric`, `key`.
	Type *string `pulumi:"type"`
	// The date the order was last updated.
	Updated *string `pulumi:"updated"`
}

type OrderV1State struct {
	// The container reference / where to find the container.
	ContainerRef pulumi.StringPtrInput
	// The date the order was created.
	Created pulumi.StringPtrInput
	// The creator of the order.
	CreatorId pulumi.StringPtrInput
	// Dictionary containing the order metadata used to generate the order. The structure is described below.
	Meta OrderV1MetaPtrInput
	// The order reference / where to find the order.
	OrderRef pulumi.StringPtrInput
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a order. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 order.
	Region pulumi.StringPtrInput
	// The secret reference / where to find the secret.
	SecretRef pulumi.StringPtrInput
	// The status of the order.
	Status pulumi.StringPtrInput
	// The sub status of the order.
	SubStatus pulumi.StringPtrInput
	// The sub status message of the order.
	SubStatusMessage pulumi.StringPtrInput
	// The type of key to be generated. Must be one of `asymmetric`, `key`.
	Type pulumi.StringPtrInput
	// The date the order was last updated.
	Updated pulumi.StringPtrInput
}

func (OrderV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*orderV1State)(nil)).Elem()
}

type orderV1Args struct {
	// Dictionary containing the order metadata used to generate the order. The structure is described below.
	Meta OrderV1Meta `pulumi:"meta"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a order. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 order.
	Region *string `pulumi:"region"`
	// The type of key to be generated. Must be one of `asymmetric`, `key`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a OrderV1 resource.
type OrderV1Args struct {
	// Dictionary containing the order metadata used to generate the order. The structure is described below.
	Meta OrderV1MetaInput
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a order. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 order.
	Region pulumi.StringPtrInput
	// The type of key to be generated. Must be one of `asymmetric`, `key`.
	Type pulumi.StringInput
}

func (OrderV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*orderV1Args)(nil)).Elem()
}

type OrderV1Input interface {
	pulumi.Input

	ToOrderV1Output() OrderV1Output
	ToOrderV1OutputWithContext(ctx context.Context) OrderV1Output
}

func (*OrderV1) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderV1)(nil)).Elem()
}

func (i *OrderV1) ToOrderV1Output() OrderV1Output {
	return i.ToOrderV1OutputWithContext(context.Background())
}

func (i *OrderV1) ToOrderV1OutputWithContext(ctx context.Context) OrderV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(OrderV1Output)
}

// OrderV1ArrayInput is an input type that accepts OrderV1Array and OrderV1ArrayOutput values.
// You can construct a concrete instance of `OrderV1ArrayInput` via:
//
//	OrderV1Array{ OrderV1Args{...} }
type OrderV1ArrayInput interface {
	pulumi.Input

	ToOrderV1ArrayOutput() OrderV1ArrayOutput
	ToOrderV1ArrayOutputWithContext(context.Context) OrderV1ArrayOutput
}

type OrderV1Array []OrderV1Input

func (OrderV1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrderV1)(nil)).Elem()
}

func (i OrderV1Array) ToOrderV1ArrayOutput() OrderV1ArrayOutput {
	return i.ToOrderV1ArrayOutputWithContext(context.Background())
}

func (i OrderV1Array) ToOrderV1ArrayOutputWithContext(ctx context.Context) OrderV1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderV1ArrayOutput)
}

// OrderV1MapInput is an input type that accepts OrderV1Map and OrderV1MapOutput values.
// You can construct a concrete instance of `OrderV1MapInput` via:
//
//	OrderV1Map{ "key": OrderV1Args{...} }
type OrderV1MapInput interface {
	pulumi.Input

	ToOrderV1MapOutput() OrderV1MapOutput
	ToOrderV1MapOutputWithContext(context.Context) OrderV1MapOutput
}

type OrderV1Map map[string]OrderV1Input

func (OrderV1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrderV1)(nil)).Elem()
}

func (i OrderV1Map) ToOrderV1MapOutput() OrderV1MapOutput {
	return i.ToOrderV1MapOutputWithContext(context.Background())
}

func (i OrderV1Map) ToOrderV1MapOutputWithContext(ctx context.Context) OrderV1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderV1MapOutput)
}

type OrderV1Output struct{ *pulumi.OutputState }

func (OrderV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**OrderV1)(nil)).Elem()
}

func (o OrderV1Output) ToOrderV1Output() OrderV1Output {
	return o
}

func (o OrderV1Output) ToOrderV1OutputWithContext(ctx context.Context) OrderV1Output {
	return o
}

// The container reference / where to find the container.
func (o OrderV1Output) ContainerRef() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.ContainerRef }).(pulumi.StringOutput)
}

// The date the order was created.
func (o OrderV1Output) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The creator of the order.
func (o OrderV1Output) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.CreatorId }).(pulumi.StringOutput)
}

// Dictionary containing the order metadata used to generate the order. The structure is described below.
func (o OrderV1Output) Meta() OrderV1MetaOutput {
	return o.ApplyT(func(v *OrderV1) OrderV1MetaOutput { return v.Meta }).(OrderV1MetaOutput)
}

// The order reference / where to find the order.
func (o OrderV1Output) OrderRef() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.OrderRef }).(pulumi.StringOutput)
}

// The region in which to obtain the V1 KeyManager client.
// A KeyManager client is needed to create a order. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// V1 order.
func (o OrderV1Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret reference / where to find the secret.
func (o OrderV1Output) SecretRef() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.SecretRef }).(pulumi.StringOutput)
}

// The status of the order.
func (o OrderV1Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The sub status of the order.
func (o OrderV1Output) SubStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.SubStatus }).(pulumi.StringOutput)
}

// The sub status message of the order.
func (o OrderV1Output) SubStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.SubStatusMessage }).(pulumi.StringOutput)
}

// The type of key to be generated. Must be one of `asymmetric`, `key`.
func (o OrderV1Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The date the order was last updated.
func (o OrderV1Output) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v *OrderV1) pulumi.StringOutput { return v.Updated }).(pulumi.StringOutput)
}

type OrderV1ArrayOutput struct{ *pulumi.OutputState }

func (OrderV1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrderV1)(nil)).Elem()
}

func (o OrderV1ArrayOutput) ToOrderV1ArrayOutput() OrderV1ArrayOutput {
	return o
}

func (o OrderV1ArrayOutput) ToOrderV1ArrayOutputWithContext(ctx context.Context) OrderV1ArrayOutput {
	return o
}

func (o OrderV1ArrayOutput) Index(i pulumi.IntInput) OrderV1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrderV1 {
		return vs[0].([]*OrderV1)[vs[1].(int)]
	}).(OrderV1Output)
}

type OrderV1MapOutput struct{ *pulumi.OutputState }

func (OrderV1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrderV1)(nil)).Elem()
}

func (o OrderV1MapOutput) ToOrderV1MapOutput() OrderV1MapOutput {
	return o
}

func (o OrderV1MapOutput) ToOrderV1MapOutputWithContext(ctx context.Context) OrderV1MapOutput {
	return o
}

func (o OrderV1MapOutput) MapIndex(k pulumi.StringInput) OrderV1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrderV1 {
		return vs[0].(map[string]*OrderV1)[vs[1].(string)]
	}).(OrderV1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrderV1Input)(nil)).Elem(), &OrderV1{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderV1ArrayInput)(nil)).Elem(), OrderV1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrderV1MapInput)(nil)).Elem(), OrderV1Map{})
	pulumi.RegisterOutputType(OrderV1Output{})
	pulumi.RegisterOutputType(OrderV1ArrayOutput{})
	pulumi.RegisterOutputType(OrderV1MapOutput{})
}
