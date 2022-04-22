// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associate a floating IP to an instance.
//
// ## Example Usage
// ### Automatically detect the correct network
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance1, err := compute.NewInstance(ctx, "instance1", &compute.InstanceArgs{
// 			FlavorId: pulumi.String("3"),
// 			ImageId:  pulumi.String("ad091b52-742f-469e-8f3c-fd81cadf0743"),
// 			KeyPair:  pulumi.String("my_key_pair_name"),
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String("default"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fip1FloatingIp, err := networking.NewFloatingIp(ctx, "fip1FloatingIp", &networking.FloatingIpArgs{
// 			Pool: pulumi.String("my_pool"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewFloatingIpAssociate(ctx, "fip1FloatingIpAssociate", &compute.FloatingIpAssociateArgs{
// 			FloatingIp: fip1FloatingIp.Address,
// 			InstanceId: instance1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Explicitly set the network to attach to
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instance1, err := compute.NewInstance(ctx, "instance1", &compute.InstanceArgs{
// 			FlavorId: pulumi.String("3"),
// 			ImageId:  pulumi.String("ad091b52-742f-469e-8f3c-fd81cadf0743"),
// 			KeyPair:  pulumi.String("my_key_pair_name"),
// 			Networks: compute.InstanceNetworkArray{
// 				&compute.InstanceNetworkArgs{
// 					Name: pulumi.String("my_network"),
// 				},
// 				&compute.InstanceNetworkArgs{
// 					Name: pulumi.String("default"),
// 				},
// 			},
// 			SecurityGroups: pulumi.StringArray{
// 				pulumi.String("default"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fip1FloatingIp, err := networking.NewFloatingIp(ctx, "fip1FloatingIp", &networking.FloatingIpArgs{
// 			Pool: pulumi.String("my_pool"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewFloatingIpAssociate(ctx, "fip1FloatingIpAssociate", &compute.FloatingIpAssociateArgs{
// 			FixedIp: instance1.Networks.ApplyT(func(networks []compute.InstanceNetwork) (string, error) {
// 				return networks[1].FixedIpV4, nil
// 			}).(pulumi.StringOutput),
// 			FloatingIp: fip1FloatingIp.Address,
// 			InstanceId: instance1.ID(),
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
// This resource can be imported by specifying all three arguments, separated by a forward slash
//
// ```sh
//  $ pulumi import openstack:compute/floatingIpAssociate:FloatingIpAssociate fip_1 <floating_ip>/<instance_id>/<fixed_ip>
// ```
type FloatingIpAssociate struct {
	pulumi.CustomResourceState

	// The specific IP address to direct traffic to.
	FixedIp pulumi.StringPtrOutput `pulumi:"fixedIp"`
	// The floating IP to associate.
	FloatingIp pulumi.StringOutput `pulumi:"floatingIp"`
	// The instance to associte the floating IP with.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region              pulumi.StringOutput  `pulumi:"region"`
	WaitUntilAssociated pulumi.BoolPtrOutput `pulumi:"waitUntilAssociated"`
}

// NewFloatingIpAssociate registers a new resource with the given unique name, arguments, and options.
func NewFloatingIpAssociate(ctx *pulumi.Context,
	name string, args *FloatingIpAssociateArgs, opts ...pulumi.ResourceOption) (*FloatingIpAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FloatingIp == nil {
		return nil, errors.New("invalid value for required argument 'FloatingIp'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource FloatingIpAssociate
	err := ctx.RegisterResource("openstack:compute/floatingIpAssociate:FloatingIpAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFloatingIpAssociate gets an existing FloatingIpAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFloatingIpAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FloatingIpAssociateState, opts ...pulumi.ResourceOption) (*FloatingIpAssociate, error) {
	var resource FloatingIpAssociate
	err := ctx.ReadResource("openstack:compute/floatingIpAssociate:FloatingIpAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FloatingIpAssociate resources.
type floatingIpAssociateState struct {
	// The specific IP address to direct traffic to.
	FixedIp *string `pulumi:"fixedIp"`
	// The floating IP to associate.
	FloatingIp *string `pulumi:"floatingIp"`
	// The instance to associte the floating IP with.
	InstanceId *string `pulumi:"instanceId"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region              *string `pulumi:"region"`
	WaitUntilAssociated *bool   `pulumi:"waitUntilAssociated"`
}

type FloatingIpAssociateState struct {
	// The specific IP address to direct traffic to.
	FixedIp pulumi.StringPtrInput
	// The floating IP to associate.
	FloatingIp pulumi.StringPtrInput
	// The instance to associte the floating IP with.
	InstanceId pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region              pulumi.StringPtrInput
	WaitUntilAssociated pulumi.BoolPtrInput
}

func (FloatingIpAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpAssociateState)(nil)).Elem()
}

type floatingIpAssociateArgs struct {
	// The specific IP address to direct traffic to.
	FixedIp *string `pulumi:"fixedIp"`
	// The floating IP to associate.
	FloatingIp string `pulumi:"floatingIp"`
	// The instance to associte the floating IP with.
	InstanceId string `pulumi:"instanceId"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region              *string `pulumi:"region"`
	WaitUntilAssociated *bool   `pulumi:"waitUntilAssociated"`
}

// The set of arguments for constructing a FloatingIpAssociate resource.
type FloatingIpAssociateArgs struct {
	// The specific IP address to direct traffic to.
	FixedIp pulumi.StringPtrInput
	// The floating IP to associate.
	FloatingIp pulumi.StringInput
	// The instance to associte the floating IP with.
	InstanceId pulumi.StringInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new floatingip_associate.
	Region              pulumi.StringPtrInput
	WaitUntilAssociated pulumi.BoolPtrInput
}

func (FloatingIpAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*floatingIpAssociateArgs)(nil)).Elem()
}

type FloatingIpAssociateInput interface {
	pulumi.Input

	ToFloatingIpAssociateOutput() FloatingIpAssociateOutput
	ToFloatingIpAssociateOutputWithContext(ctx context.Context) FloatingIpAssociateOutput
}

func (*FloatingIpAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIpAssociate)(nil)).Elem()
}

func (i *FloatingIpAssociate) ToFloatingIpAssociateOutput() FloatingIpAssociateOutput {
	return i.ToFloatingIpAssociateOutputWithContext(context.Background())
}

func (i *FloatingIpAssociate) ToFloatingIpAssociateOutputWithContext(ctx context.Context) FloatingIpAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpAssociateOutput)
}

// FloatingIpAssociateArrayInput is an input type that accepts FloatingIpAssociateArray and FloatingIpAssociateArrayOutput values.
// You can construct a concrete instance of `FloatingIpAssociateArrayInput` via:
//
//          FloatingIpAssociateArray{ FloatingIpAssociateArgs{...} }
type FloatingIpAssociateArrayInput interface {
	pulumi.Input

	ToFloatingIpAssociateArrayOutput() FloatingIpAssociateArrayOutput
	ToFloatingIpAssociateArrayOutputWithContext(context.Context) FloatingIpAssociateArrayOutput
}

type FloatingIpAssociateArray []FloatingIpAssociateInput

func (FloatingIpAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FloatingIpAssociate)(nil)).Elem()
}

func (i FloatingIpAssociateArray) ToFloatingIpAssociateArrayOutput() FloatingIpAssociateArrayOutput {
	return i.ToFloatingIpAssociateArrayOutputWithContext(context.Background())
}

func (i FloatingIpAssociateArray) ToFloatingIpAssociateArrayOutputWithContext(ctx context.Context) FloatingIpAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpAssociateArrayOutput)
}

// FloatingIpAssociateMapInput is an input type that accepts FloatingIpAssociateMap and FloatingIpAssociateMapOutput values.
// You can construct a concrete instance of `FloatingIpAssociateMapInput` via:
//
//          FloatingIpAssociateMap{ "key": FloatingIpAssociateArgs{...} }
type FloatingIpAssociateMapInput interface {
	pulumi.Input

	ToFloatingIpAssociateMapOutput() FloatingIpAssociateMapOutput
	ToFloatingIpAssociateMapOutputWithContext(context.Context) FloatingIpAssociateMapOutput
}

type FloatingIpAssociateMap map[string]FloatingIpAssociateInput

func (FloatingIpAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FloatingIpAssociate)(nil)).Elem()
}

func (i FloatingIpAssociateMap) ToFloatingIpAssociateMapOutput() FloatingIpAssociateMapOutput {
	return i.ToFloatingIpAssociateMapOutputWithContext(context.Background())
}

func (i FloatingIpAssociateMap) ToFloatingIpAssociateMapOutputWithContext(ctx context.Context) FloatingIpAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FloatingIpAssociateMapOutput)
}

type FloatingIpAssociateOutput struct{ *pulumi.OutputState }

func (FloatingIpAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FloatingIpAssociate)(nil)).Elem()
}

func (o FloatingIpAssociateOutput) ToFloatingIpAssociateOutput() FloatingIpAssociateOutput {
	return o
}

func (o FloatingIpAssociateOutput) ToFloatingIpAssociateOutputWithContext(ctx context.Context) FloatingIpAssociateOutput {
	return o
}

type FloatingIpAssociateArrayOutput struct{ *pulumi.OutputState }

func (FloatingIpAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FloatingIpAssociate)(nil)).Elem()
}

func (o FloatingIpAssociateArrayOutput) ToFloatingIpAssociateArrayOutput() FloatingIpAssociateArrayOutput {
	return o
}

func (o FloatingIpAssociateArrayOutput) ToFloatingIpAssociateArrayOutputWithContext(ctx context.Context) FloatingIpAssociateArrayOutput {
	return o
}

func (o FloatingIpAssociateArrayOutput) Index(i pulumi.IntInput) FloatingIpAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FloatingIpAssociate {
		return vs[0].([]*FloatingIpAssociate)[vs[1].(int)]
	}).(FloatingIpAssociateOutput)
}

type FloatingIpAssociateMapOutput struct{ *pulumi.OutputState }

func (FloatingIpAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FloatingIpAssociate)(nil)).Elem()
}

func (o FloatingIpAssociateMapOutput) ToFloatingIpAssociateMapOutput() FloatingIpAssociateMapOutput {
	return o
}

func (o FloatingIpAssociateMapOutput) ToFloatingIpAssociateMapOutputWithContext(ctx context.Context) FloatingIpAssociateMapOutput {
	return o
}

func (o FloatingIpAssociateMapOutput) MapIndex(k pulumi.StringInput) FloatingIpAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FloatingIpAssociate {
		return vs[0].(map[string]*FloatingIpAssociate)[vs[1].(string)]
	}).(FloatingIpAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpAssociateInput)(nil)).Elem(), &FloatingIpAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpAssociateArrayInput)(nil)).Elem(), FloatingIpAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FloatingIpAssociateMapInput)(nil)).Elem(), FloatingIpAssociateMap{})
	pulumi.RegisterOutputType(FloatingIpAssociateOutput{})
	pulumi.RegisterOutputType(FloatingIpAssociateArrayOutput{})
	pulumi.RegisterOutputType(FloatingIpAssociateMapOutput{})
}
