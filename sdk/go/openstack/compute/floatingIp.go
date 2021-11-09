// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 floating IP resource within OpenStack Nova (compute)
// that can be used for compute instances.
//
// Please note that managing floating IPs through the OpenStack Compute API has
// been deprecated. Unless you are using an older OpenStack environment, it is
// recommended to use the `networking.FloatingIp`
// resource instead, which uses the OpenStack Networking API.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewFloatingIp(ctx, "floatip1", &compute.FloatingIpArgs{
// 			Pool: pulumi.String("public"),
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
// Floating IPs can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:compute/floatingIp:FloatingIp floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2
// ```
type FloatingIp struct {
	pulumi.CustomResourceState

	// The actual floating IP address itself.
	Address pulumi.StringOutput `pulumi:"address"`
	// The fixed IP address corresponding to the floating IP.
	FixedIp pulumi.StringOutput `pulumi:"fixedIp"`
	// UUID of the compute instance associated with the floating IP.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool pulumi.StringOutput `pulumi:"pool"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a floating IP that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new floating IP (which may or may not
	// have a different address).
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewFloatingIp registers a new resource with the given unique name, arguments, and options.
func NewFloatingIp(ctx *pulumi.Context,
	name string, args *FloatingIpArgs, opts ...pulumi.ResourceOption) (*FloatingIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Pool == nil {
		return nil, errors.New("invalid value for required argument 'Pool'")
	}
	var resource FloatingIp
	err := ctx.RegisterResource("openstack:compute/floatingIp:FloatingIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFloatingIp gets an existing FloatingIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FloatingIpState, opts ...pulumi.ResourceOption) (*FloatingIp, error) {
	var resource FloatingIp
	err := ctx.ReadResource("openstack:compute/floatingIp:FloatingIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FloatingIp resources.
type floatingIpState struct {
	// The actual floating IP address itself.
	Address *string `pulumi:"address"`
	// The fixed IP address corresponding to the floating IP.
	FixedIp *string `pulumi:"fixedIp"`
	// UUID of the compute instance associated with the floating IP.
	InstanceId *string `pulumi:"instanceId"`
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool *string `pulumi:"pool"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a floating IP that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new floating IP (which may or may not
	// have a different address).
	Region *string `pulumi:"region"`
}

type FloatingIpState struct {
	// The actual floating IP address itself.
	Address pulumi.StringPtrInput
	// The fixed IP address corresponding to the floating IP.
	FixedIp pulumi.StringPtrInput
	// UUID of the compute instance associated with the floating IP.
	InstanceId pulumi.StringPtrInput
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a floating IP that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new floating IP (which may or may not
	// have a different address).
	Region pulumi.StringPtrInput
}

func (FloatingIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpState)(nil)).Elem()
}

type floatingIpArgs struct {
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool string `pulumi:"pool"`
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a floating IP that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new floating IP (which may or may not
	// have a different address).
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FloatingIp resource.
type FloatingIpArgs struct {
	// The name of the pool from which to obtain the floating
	// IP. Changing this creates a new floating IP.
	Pool pulumi.StringInput
	// The region in which to obtain the V2 Compute client.
	// A Compute client is needed to create a floating IP that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new floating IP (which may or may not
	// have a different address).
	Region pulumi.StringPtrInput
}

func (FloatingIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpArgs)(nil)).Elem()
}

type FloatingIpInput interface {
	pulumi.Input

	ToFloatingIpOutput() FloatingIpOutput
	ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput
}

func (*FloatingIp) ElementType() reflect.Type {
	return reflect.TypeOf((*FloatingIp)(nil))
}

func (i *FloatingIp) ToFloatingIpOutput() FloatingIpOutput {
	return i.ToFloatingIpOutputWithContext(context.Background())
}

func (i *FloatingIp) ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpOutput)
}

func (i *FloatingIp) ToFloatingIpPtrOutput() FloatingIpPtrOutput {
	return i.ToFloatingIpPtrOutputWithContext(context.Background())
}

func (i *FloatingIp) ToFloatingIpPtrOutputWithContext(ctx context.Context) FloatingIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpPtrOutput)
}

type FloatingIpPtrInput interface {
	pulumi.Input

	ToFloatingIpPtrOutput() FloatingIpPtrOutput
	ToFloatingIpPtrOutputWithContext(ctx context.Context) FloatingIpPtrOutput
}

type floatingIpPtrType FloatingIpArgs

func (*floatingIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIp)(nil))
}

func (i *floatingIpPtrType) ToFloatingIpPtrOutput() FloatingIpPtrOutput {
	return i.ToFloatingIpPtrOutputWithContext(context.Background())
}

func (i *floatingIpPtrType) ToFloatingIpPtrOutputWithContext(ctx context.Context) FloatingIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpPtrOutput)
}

// FloatingIpArrayInput is an input type that accepts FloatingIpArray and FloatingIpArrayOutput values.
// You can construct a concrete instance of `FloatingIpArrayInput` via:
//
//          FloatingIpArray{ FloatingIpArgs{...} }
type FloatingIpArrayInput interface {
	pulumi.Input

	ToFloatingIpArrayOutput() FloatingIpArrayOutput
	ToFloatingIpArrayOutputWithContext(context.Context) FloatingIpArrayOutput
}

type FloatingIpArray []FloatingIpInput

func (FloatingIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FloatingIp)(nil)).Elem()
}

func (i FloatingIpArray) ToFloatingIpArrayOutput() FloatingIpArrayOutput {
	return i.ToFloatingIpArrayOutputWithContext(context.Background())
}

func (i FloatingIpArray) ToFloatingIpArrayOutputWithContext(ctx context.Context) FloatingIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpArrayOutput)
}

// FloatingIpMapInput is an input type that accepts FloatingIpMap and FloatingIpMapOutput values.
// You can construct a concrete instance of `FloatingIpMapInput` via:
//
//          FloatingIpMap{ "key": FloatingIpArgs{...} }
type FloatingIpMapInput interface {
	pulumi.Input

	ToFloatingIpMapOutput() FloatingIpMapOutput
	ToFloatingIpMapOutputWithContext(context.Context) FloatingIpMapOutput
}

type FloatingIpMap map[string]FloatingIpInput

func (FloatingIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FloatingIp)(nil)).Elem()
}

func (i FloatingIpMap) ToFloatingIpMapOutput() FloatingIpMapOutput {
	return i.ToFloatingIpMapOutputWithContext(context.Background())
}

func (i FloatingIpMap) ToFloatingIpMapOutputWithContext(ctx context.Context) FloatingIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpMapOutput)
}

type FloatingIpOutput struct{ *pulumi.OutputState }

func (FloatingIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FloatingIp)(nil))
}

func (o FloatingIpOutput) ToFloatingIpOutput() FloatingIpOutput {
	return o
}

func (o FloatingIpOutput) ToFloatingIpOutputWithContext(ctx context.Context) FloatingIpOutput {
	return o
}

func (o FloatingIpOutput) ToFloatingIpPtrOutput() FloatingIpPtrOutput {
	return o.ToFloatingIpPtrOutputWithContext(context.Background())
}

func (o FloatingIpOutput) ToFloatingIpPtrOutputWithContext(ctx context.Context) FloatingIpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FloatingIp) *FloatingIp {
		return &v
	}).(FloatingIpPtrOutput)
}

type FloatingIpPtrOutput struct{ *pulumi.OutputState }

func (FloatingIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIp)(nil))
}

func (o FloatingIpPtrOutput) ToFloatingIpPtrOutput() FloatingIpPtrOutput {
	return o
}

func (o FloatingIpPtrOutput) ToFloatingIpPtrOutputWithContext(ctx context.Context) FloatingIpPtrOutput {
	return o
}

func (o FloatingIpPtrOutput) Elem() FloatingIpOutput {
	return o.ApplyT(func(v *FloatingIp) FloatingIp {
		if v != nil {
			return *v
		}
		var ret FloatingIp
		return ret
	}).(FloatingIpOutput)
}

type FloatingIpArrayOutput struct{ *pulumi.OutputState }

func (FloatingIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FloatingIp)(nil))
}

func (o FloatingIpArrayOutput) ToFloatingIpArrayOutput() FloatingIpArrayOutput {
	return o
}

func (o FloatingIpArrayOutput) ToFloatingIpArrayOutputWithContext(ctx context.Context) FloatingIpArrayOutput {
	return o
}

func (o FloatingIpArrayOutput) Index(i pulumi.IntInput) FloatingIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FloatingIp {
		return vs[0].([]FloatingIp)[vs[1].(int)]
	}).(FloatingIpOutput)
}

type FloatingIpMapOutput struct{ *pulumi.OutputState }

func (FloatingIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FloatingIp)(nil))
}

func (o FloatingIpMapOutput) ToFloatingIpMapOutput() FloatingIpMapOutput {
	return o
}

func (o FloatingIpMapOutput) ToFloatingIpMapOutputWithContext(ctx context.Context) FloatingIpMapOutput {
	return o
}

func (o FloatingIpMapOutput) MapIndex(k pulumi.StringInput) FloatingIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FloatingIp {
		return vs[0].(map[string]FloatingIp)[vs[1].(string)]
	}).(FloatingIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpInput)(nil)).Elem(), &FloatingIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpPtrInput)(nil)).Elem(), &FloatingIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpArrayInput)(nil)).Elem(), FloatingIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpMapInput)(nil)).Elem(), FloatingIpMap{})
	pulumi.RegisterOutputType(FloatingIpOutput{})
	pulumi.RegisterOutputType(FloatingIpPtrOutput{})
	pulumi.RegisterOutputType(FloatingIpArrayOutput{})
	pulumi.RegisterOutputType(FloatingIpMapOutput{})
}
