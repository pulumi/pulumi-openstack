// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openstack

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 load balancer flavorprofile resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// ## Example Usage
//
// ### Using jsonencode
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"loadbalancer_topology": "SINGLE",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = openstack.NewLbFlavorprofileV2(ctx, "flavorprofile_1", &openstack.LbFlavorprofileV2Args{
//				Name:         pulumi.String("amphora-single-profile"),
//				ProviderName: pulumi.String("amphora"),
//				FlavorData:   pulumi.String(json0),
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
// ### Using plain string
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := openstack.NewLbFlavorprofileV2(ctx, "flavorprofile_1", &openstack.LbFlavorprofileV2Args{
//				Name:         pulumi.String("amphora-single-profile"),
//				ProviderName: pulumi.String("amphora"),
//				FlavorData:   pulumi.String("{\"loadbalancer_topology\": \"SINGLE\"}"),
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
// flavorprofiles can be imported using their `id`. Example:
//
// ```sh
// $ pulumi import openstack:index/lbFlavorprofileV2:LbFlavorprofileV2 flavorprofile_1 2a0f2240-c5e6-41de-896d-e80d97428d6b
// ```
type LbFlavorprofileV2 struct {
	pulumi.CustomResourceState

	// String that passes the flavorData for the flavorprofile.
	// The data that are allowed depend on the `providerName` that is passed. jsonencode
	// can be used for readability as shown in the example above.
	// Changing this updates the existing flavorprofile.
	FlavorData pulumi.StringOutput `pulumi:"flavorData"`
	// Name of the flavorprofile. Changing this updates the existing
	// flavorprofile.
	Name pulumi.StringOutput `pulumi:"name"`
	// The providerName that the flavorProfile will use.
	// Changing this updates the existing flavorprofile.
	ProviderName pulumi.StringOutput `pulumi:"providerName"`
	Region       pulumi.StringOutput `pulumi:"region"`
}

// NewLbFlavorprofileV2 registers a new resource with the given unique name, arguments, and options.
func NewLbFlavorprofileV2(ctx *pulumi.Context,
	name string, args *LbFlavorprofileV2Args, opts ...pulumi.ResourceOption) (*LbFlavorprofileV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FlavorData == nil {
		return nil, errors.New("invalid value for required argument 'FlavorData'")
	}
	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbFlavorprofileV2
	err := ctx.RegisterResource("openstack:index/lbFlavorprofileV2:LbFlavorprofileV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbFlavorprofileV2 gets an existing LbFlavorprofileV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbFlavorprofileV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbFlavorprofileV2State, opts ...pulumi.ResourceOption) (*LbFlavorprofileV2, error) {
	var resource LbFlavorprofileV2
	err := ctx.ReadResource("openstack:index/lbFlavorprofileV2:LbFlavorprofileV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbFlavorprofileV2 resources.
type lbFlavorprofileV2State struct {
	// String that passes the flavorData for the flavorprofile.
	// The data that are allowed depend on the `providerName` that is passed. jsonencode
	// can be used for readability as shown in the example above.
	// Changing this updates the existing flavorprofile.
	FlavorData *string `pulumi:"flavorData"`
	// Name of the flavorprofile. Changing this updates the existing
	// flavorprofile.
	Name *string `pulumi:"name"`
	// The providerName that the flavorProfile will use.
	// Changing this updates the existing flavorprofile.
	ProviderName *string `pulumi:"providerName"`
	Region       *string `pulumi:"region"`
}

type LbFlavorprofileV2State struct {
	// String that passes the flavorData for the flavorprofile.
	// The data that are allowed depend on the `providerName` that is passed. jsonencode
	// can be used for readability as shown in the example above.
	// Changing this updates the existing flavorprofile.
	FlavorData pulumi.StringPtrInput
	// Name of the flavorprofile. Changing this updates the existing
	// flavorprofile.
	Name pulumi.StringPtrInput
	// The providerName that the flavorProfile will use.
	// Changing this updates the existing flavorprofile.
	ProviderName pulumi.StringPtrInput
	Region       pulumi.StringPtrInput
}

func (LbFlavorprofileV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*lbFlavorprofileV2State)(nil)).Elem()
}

type lbFlavorprofileV2Args struct {
	// String that passes the flavorData for the flavorprofile.
	// The data that are allowed depend on the `providerName` that is passed. jsonencode
	// can be used for readability as shown in the example above.
	// Changing this updates the existing flavorprofile.
	FlavorData string `pulumi:"flavorData"`
	// Name of the flavorprofile. Changing this updates the existing
	// flavorprofile.
	Name *string `pulumi:"name"`
	// The providerName that the flavorProfile will use.
	// Changing this updates the existing flavorprofile.
	ProviderName string  `pulumi:"providerName"`
	Region       *string `pulumi:"region"`
}

// The set of arguments for constructing a LbFlavorprofileV2 resource.
type LbFlavorprofileV2Args struct {
	// String that passes the flavorData for the flavorprofile.
	// The data that are allowed depend on the `providerName` that is passed. jsonencode
	// can be used for readability as shown in the example above.
	// Changing this updates the existing flavorprofile.
	FlavorData pulumi.StringInput
	// Name of the flavorprofile. Changing this updates the existing
	// flavorprofile.
	Name pulumi.StringPtrInput
	// The providerName that the flavorProfile will use.
	// Changing this updates the existing flavorprofile.
	ProviderName pulumi.StringInput
	Region       pulumi.StringPtrInput
}

func (LbFlavorprofileV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*lbFlavorprofileV2Args)(nil)).Elem()
}

type LbFlavorprofileV2Input interface {
	pulumi.Input

	ToLbFlavorprofileV2Output() LbFlavorprofileV2Output
	ToLbFlavorprofileV2OutputWithContext(ctx context.Context) LbFlavorprofileV2Output
}

func (*LbFlavorprofileV2) ElementType() reflect.Type {
	return reflect.TypeOf((**LbFlavorprofileV2)(nil)).Elem()
}

func (i *LbFlavorprofileV2) ToLbFlavorprofileV2Output() LbFlavorprofileV2Output {
	return i.ToLbFlavorprofileV2OutputWithContext(context.Background())
}

func (i *LbFlavorprofileV2) ToLbFlavorprofileV2OutputWithContext(ctx context.Context) LbFlavorprofileV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(LbFlavorprofileV2Output)
}

// LbFlavorprofileV2ArrayInput is an input type that accepts LbFlavorprofileV2Array and LbFlavorprofileV2ArrayOutput values.
// You can construct a concrete instance of `LbFlavorprofileV2ArrayInput` via:
//
//	LbFlavorprofileV2Array{ LbFlavorprofileV2Args{...} }
type LbFlavorprofileV2ArrayInput interface {
	pulumi.Input

	ToLbFlavorprofileV2ArrayOutput() LbFlavorprofileV2ArrayOutput
	ToLbFlavorprofileV2ArrayOutputWithContext(context.Context) LbFlavorprofileV2ArrayOutput
}

type LbFlavorprofileV2Array []LbFlavorprofileV2Input

func (LbFlavorprofileV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbFlavorprofileV2)(nil)).Elem()
}

func (i LbFlavorprofileV2Array) ToLbFlavorprofileV2ArrayOutput() LbFlavorprofileV2ArrayOutput {
	return i.ToLbFlavorprofileV2ArrayOutputWithContext(context.Background())
}

func (i LbFlavorprofileV2Array) ToLbFlavorprofileV2ArrayOutputWithContext(ctx context.Context) LbFlavorprofileV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbFlavorprofileV2ArrayOutput)
}

// LbFlavorprofileV2MapInput is an input type that accepts LbFlavorprofileV2Map and LbFlavorprofileV2MapOutput values.
// You can construct a concrete instance of `LbFlavorprofileV2MapInput` via:
//
//	LbFlavorprofileV2Map{ "key": LbFlavorprofileV2Args{...} }
type LbFlavorprofileV2MapInput interface {
	pulumi.Input

	ToLbFlavorprofileV2MapOutput() LbFlavorprofileV2MapOutput
	ToLbFlavorprofileV2MapOutputWithContext(context.Context) LbFlavorprofileV2MapOutput
}

type LbFlavorprofileV2Map map[string]LbFlavorprofileV2Input

func (LbFlavorprofileV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbFlavorprofileV2)(nil)).Elem()
}

func (i LbFlavorprofileV2Map) ToLbFlavorprofileV2MapOutput() LbFlavorprofileV2MapOutput {
	return i.ToLbFlavorprofileV2MapOutputWithContext(context.Background())
}

func (i LbFlavorprofileV2Map) ToLbFlavorprofileV2MapOutputWithContext(ctx context.Context) LbFlavorprofileV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbFlavorprofileV2MapOutput)
}

type LbFlavorprofileV2Output struct{ *pulumi.OutputState }

func (LbFlavorprofileV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**LbFlavorprofileV2)(nil)).Elem()
}

func (o LbFlavorprofileV2Output) ToLbFlavorprofileV2Output() LbFlavorprofileV2Output {
	return o
}

func (o LbFlavorprofileV2Output) ToLbFlavorprofileV2OutputWithContext(ctx context.Context) LbFlavorprofileV2Output {
	return o
}

// String that passes the flavorData for the flavorprofile.
// The data that are allowed depend on the `providerName` that is passed. jsonencode
// can be used for readability as shown in the example above.
// Changing this updates the existing flavorprofile.
func (o LbFlavorprofileV2Output) FlavorData() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFlavorprofileV2) pulumi.StringOutput { return v.FlavorData }).(pulumi.StringOutput)
}

// Name of the flavorprofile. Changing this updates the existing
// flavorprofile.
func (o LbFlavorprofileV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFlavorprofileV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The providerName that the flavorProfile will use.
// Changing this updates the existing flavorprofile.
func (o LbFlavorprofileV2Output) ProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFlavorprofileV2) pulumi.StringOutput { return v.ProviderName }).(pulumi.StringOutput)
}

func (o LbFlavorprofileV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *LbFlavorprofileV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type LbFlavorprofileV2ArrayOutput struct{ *pulumi.OutputState }

func (LbFlavorprofileV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbFlavorprofileV2)(nil)).Elem()
}

func (o LbFlavorprofileV2ArrayOutput) ToLbFlavorprofileV2ArrayOutput() LbFlavorprofileV2ArrayOutput {
	return o
}

func (o LbFlavorprofileV2ArrayOutput) ToLbFlavorprofileV2ArrayOutputWithContext(ctx context.Context) LbFlavorprofileV2ArrayOutput {
	return o
}

func (o LbFlavorprofileV2ArrayOutput) Index(i pulumi.IntInput) LbFlavorprofileV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbFlavorprofileV2 {
		return vs[0].([]*LbFlavorprofileV2)[vs[1].(int)]
	}).(LbFlavorprofileV2Output)
}

type LbFlavorprofileV2MapOutput struct{ *pulumi.OutputState }

func (LbFlavorprofileV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbFlavorprofileV2)(nil)).Elem()
}

func (o LbFlavorprofileV2MapOutput) ToLbFlavorprofileV2MapOutput() LbFlavorprofileV2MapOutput {
	return o
}

func (o LbFlavorprofileV2MapOutput) ToLbFlavorprofileV2MapOutputWithContext(ctx context.Context) LbFlavorprofileV2MapOutput {
	return o
}

func (o LbFlavorprofileV2MapOutput) MapIndex(k pulumi.StringInput) LbFlavorprofileV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbFlavorprofileV2 {
		return vs[0].(map[string]*LbFlavorprofileV2)[vs[1].(string)]
	}).(LbFlavorprofileV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbFlavorprofileV2Input)(nil)).Elem(), &LbFlavorprofileV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbFlavorprofileV2ArrayInput)(nil)).Elem(), LbFlavorprofileV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbFlavorprofileV2MapInput)(nil)).Elem(), LbFlavorprofileV2Map{})
	pulumi.RegisterOutputType(LbFlavorprofileV2Output{})
	pulumi.RegisterOutputType(LbFlavorprofileV2ArrayOutput{})
	pulumi.RegisterOutputType(LbFlavorprofileV2MapOutput{})
}