// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 Service resource within OpenStack Keystone.
//
// > **Note:** This usually requires admin privileges.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewServiceV3(ctx, "service_1", &identity.ServiceV3Args{
//				Name: pulumi.String("custom"),
//				Type: pulumi.String("custom"),
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
// Services can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:identity/serviceV3:ServiceV3 service_1 6688e967-158a-496f-a224-cae3414e6b61
// ```
type ServiceV3 struct {
	pulumi.CustomResourceState

	// The service description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The service status. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The service type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceV3 registers a new resource with the given unique name, arguments, and options.
func NewServiceV3(ctx *pulumi.Context,
	name string, args *ServiceV3Args, opts ...pulumi.ResourceOption) (*ServiceV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceV3
	err := ctx.RegisterResource("openstack:identity/serviceV3:ServiceV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceV3 gets an existing ServiceV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceV3State, opts ...pulumi.ResourceOption) (*ServiceV3, error) {
	var resource ServiceV3
	err := ctx.ReadResource("openstack:identity/serviceV3:ServiceV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceV3 resources.
type serviceV3State struct {
	// The service description.
	Description *string `pulumi:"description"`
	// The service status. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The service name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The service type.
	Type *string `pulumi:"type"`
}

type ServiceV3State struct {
	// The service description.
	Description pulumi.StringPtrInput
	// The service status. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The service name.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput
	// The service type.
	Type pulumi.StringPtrInput
}

func (ServiceV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceV3State)(nil)).Elem()
}

type serviceV3Args struct {
	// The service description.
	Description *string `pulumi:"description"`
	// The service status. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The service name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The service type.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ServiceV3 resource.
type ServiceV3Args struct {
	// The service description.
	Description pulumi.StringPtrInput
	// The service status. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The service name.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput
	// The service type.
	Type pulumi.StringInput
}

func (ServiceV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceV3Args)(nil)).Elem()
}

type ServiceV3Input interface {
	pulumi.Input

	ToServiceV3Output() ServiceV3Output
	ToServiceV3OutputWithContext(ctx context.Context) ServiceV3Output
}

func (*ServiceV3) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceV3)(nil)).Elem()
}

func (i *ServiceV3) ToServiceV3Output() ServiceV3Output {
	return i.ToServiceV3OutputWithContext(context.Background())
}

func (i *ServiceV3) ToServiceV3OutputWithContext(ctx context.Context) ServiceV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceV3Output)
}

// ServiceV3ArrayInput is an input type that accepts ServiceV3Array and ServiceV3ArrayOutput values.
// You can construct a concrete instance of `ServiceV3ArrayInput` via:
//
//	ServiceV3Array{ ServiceV3Args{...} }
type ServiceV3ArrayInput interface {
	pulumi.Input

	ToServiceV3ArrayOutput() ServiceV3ArrayOutput
	ToServiceV3ArrayOutputWithContext(context.Context) ServiceV3ArrayOutput
}

type ServiceV3Array []ServiceV3Input

func (ServiceV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceV3)(nil)).Elem()
}

func (i ServiceV3Array) ToServiceV3ArrayOutput() ServiceV3ArrayOutput {
	return i.ToServiceV3ArrayOutputWithContext(context.Background())
}

func (i ServiceV3Array) ToServiceV3ArrayOutputWithContext(ctx context.Context) ServiceV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceV3ArrayOutput)
}

// ServiceV3MapInput is an input type that accepts ServiceV3Map and ServiceV3MapOutput values.
// You can construct a concrete instance of `ServiceV3MapInput` via:
//
//	ServiceV3Map{ "key": ServiceV3Args{...} }
type ServiceV3MapInput interface {
	pulumi.Input

	ToServiceV3MapOutput() ServiceV3MapOutput
	ToServiceV3MapOutputWithContext(context.Context) ServiceV3MapOutput
}

type ServiceV3Map map[string]ServiceV3Input

func (ServiceV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceV3)(nil)).Elem()
}

func (i ServiceV3Map) ToServiceV3MapOutput() ServiceV3MapOutput {
	return i.ToServiceV3MapOutputWithContext(context.Background())
}

func (i ServiceV3Map) ToServiceV3MapOutputWithContext(ctx context.Context) ServiceV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceV3MapOutput)
}

type ServiceV3Output struct{ *pulumi.OutputState }

func (ServiceV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceV3)(nil)).Elem()
}

func (o ServiceV3Output) ToServiceV3Output() ServiceV3Output {
	return o
}

func (o ServiceV3Output) ToServiceV3OutputWithContext(ctx context.Context) ServiceV3Output {
	return o
}

// The service description.
func (o ServiceV3Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceV3) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The service status. Defaults to `true`.
func (o ServiceV3Output) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceV3) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The service name.
func (o ServiceV3Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceV3) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used.
func (o ServiceV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The service type.
func (o ServiceV3Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceV3) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ServiceV3ArrayOutput struct{ *pulumi.OutputState }

func (ServiceV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceV3)(nil)).Elem()
}

func (o ServiceV3ArrayOutput) ToServiceV3ArrayOutput() ServiceV3ArrayOutput {
	return o
}

func (o ServiceV3ArrayOutput) ToServiceV3ArrayOutputWithContext(ctx context.Context) ServiceV3ArrayOutput {
	return o
}

func (o ServiceV3ArrayOutput) Index(i pulumi.IntInput) ServiceV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceV3 {
		return vs[0].([]*ServiceV3)[vs[1].(int)]
	}).(ServiceV3Output)
}

type ServiceV3MapOutput struct{ *pulumi.OutputState }

func (ServiceV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceV3)(nil)).Elem()
}

func (o ServiceV3MapOutput) ToServiceV3MapOutput() ServiceV3MapOutput {
	return o
}

func (o ServiceV3MapOutput) ToServiceV3MapOutputWithContext(ctx context.Context) ServiceV3MapOutput {
	return o
}

func (o ServiceV3MapOutput) MapIndex(k pulumi.StringInput) ServiceV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceV3 {
		return vs[0].(map[string]*ServiceV3)[vs[1].(string)]
	}).(ServiceV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceV3Input)(nil)).Elem(), &ServiceV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceV3ArrayInput)(nil)).Elem(), ServiceV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceV3MapInput)(nil)).Elem(), ServiceV3Map{})
	pulumi.RegisterOutputType(ServiceV3Output{})
	pulumi.RegisterOutputType(ServiceV3ArrayOutput{})
	pulumi.RegisterOutputType(ServiceV3MapOutput{})
}
