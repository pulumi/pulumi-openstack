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

// Manages a V3 Endpoint resource within OpenStack Keystone.
//
// > **Note:** This usually requires admin privileges.
//
// ## Example Usage
//
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
//			service1, err := identity.NewServiceV3(ctx, "service_1", &identity.ServiceV3Args{
//				Name: pulumi.String("my-service"),
//				Type: pulumi.String("my-service-type"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewEndpointV3(ctx, "endpoint_1", &identity.EndpointV3Args{
//				Name:           pulumi.String("my-endpoint"),
//				ServiceId:      service1.ID(),
//				EndpointRegion: service1.Region,
//				Url:            pulumi.String("http://my-endpoint"),
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
// Endpoints can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:identity/endpointV3:EndpointV3 endpoint_1 5392472b-106a-4845-90c6-7c8445f18770
// ```
type EndpointV3 struct {
	pulumi.CustomResourceState

	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion pulumi.StringOutput `pulumi:"endpointRegion"`
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface pulumi.StringPtrOutput `pulumi:"interface"`
	// The endpoint name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The endpoint service ID.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The service name of the endpoint.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The service type of the endpoint.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// The endpoint url.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewEndpointV3 registers a new resource with the given unique name, arguments, and options.
func NewEndpointV3(ctx *pulumi.Context,
	name string, args *EndpointV3Args, opts ...pulumi.ResourceOption) (*EndpointV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointRegion == nil {
		return nil, errors.New("invalid value for required argument 'EndpointRegion'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointV3
	err := ctx.RegisterResource("openstack:identity/endpointV3:EndpointV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointV3 gets an existing EndpointV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointV3State, opts ...pulumi.ResourceOption) (*EndpointV3, error) {
	var resource EndpointV3
	err := ctx.ReadResource("openstack:identity/endpointV3:EndpointV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointV3 resources.
type endpointV3State struct {
	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion *string `pulumi:"endpointRegion"`
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface *string `pulumi:"interface"`
	// The endpoint name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The endpoint service ID.
	ServiceId *string `pulumi:"serviceId"`
	// The service name of the endpoint.
	ServiceName *string `pulumi:"serviceName"`
	// The service type of the endpoint.
	ServiceType *string `pulumi:"serviceType"`
	// The endpoint url.
	Url *string `pulumi:"url"`
}

type EndpointV3State struct {
	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion pulumi.StringPtrInput
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface pulumi.StringPtrInput
	// The endpoint name.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput
	// The endpoint service ID.
	ServiceId pulumi.StringPtrInput
	// The service name of the endpoint.
	ServiceName pulumi.StringPtrInput
	// The service type of the endpoint.
	ServiceType pulumi.StringPtrInput
	// The endpoint url.
	Url pulumi.StringPtrInput
}

func (EndpointV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointV3State)(nil)).Elem()
}

type endpointV3Args struct {
	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion string `pulumi:"endpointRegion"`
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface *string `pulumi:"interface"`
	// The endpoint name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The endpoint service ID.
	ServiceId string `pulumi:"serviceId"`
	// The endpoint url.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a EndpointV3 resource.
type EndpointV3Args struct {
	// The endpoint region. The `region` and
	// `endpointRegion` can be different.
	EndpointRegion pulumi.StringInput
	// The endpoint interface. Valid values are `public`,
	// `internal` and `admin`. Default value is `public`
	Interface pulumi.StringPtrInput
	// The endpoint name.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput
	// The endpoint service ID.
	ServiceId pulumi.StringInput
	// The endpoint url.
	Url pulumi.StringInput
}

func (EndpointV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointV3Args)(nil)).Elem()
}

type EndpointV3Input interface {
	pulumi.Input

	ToEndpointV3Output() EndpointV3Output
	ToEndpointV3OutputWithContext(ctx context.Context) EndpointV3Output
}

func (*EndpointV3) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointV3)(nil)).Elem()
}

func (i *EndpointV3) ToEndpointV3Output() EndpointV3Output {
	return i.ToEndpointV3OutputWithContext(context.Background())
}

func (i *EndpointV3) ToEndpointV3OutputWithContext(ctx context.Context) EndpointV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointV3Output)
}

// EndpointV3ArrayInput is an input type that accepts EndpointV3Array and EndpointV3ArrayOutput values.
// You can construct a concrete instance of `EndpointV3ArrayInput` via:
//
//	EndpointV3Array{ EndpointV3Args{...} }
type EndpointV3ArrayInput interface {
	pulumi.Input

	ToEndpointV3ArrayOutput() EndpointV3ArrayOutput
	ToEndpointV3ArrayOutputWithContext(context.Context) EndpointV3ArrayOutput
}

type EndpointV3Array []EndpointV3Input

func (EndpointV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointV3)(nil)).Elem()
}

func (i EndpointV3Array) ToEndpointV3ArrayOutput() EndpointV3ArrayOutput {
	return i.ToEndpointV3ArrayOutputWithContext(context.Background())
}

func (i EndpointV3Array) ToEndpointV3ArrayOutputWithContext(ctx context.Context) EndpointV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointV3ArrayOutput)
}

// EndpointV3MapInput is an input type that accepts EndpointV3Map and EndpointV3MapOutput values.
// You can construct a concrete instance of `EndpointV3MapInput` via:
//
//	EndpointV3Map{ "key": EndpointV3Args{...} }
type EndpointV3MapInput interface {
	pulumi.Input

	ToEndpointV3MapOutput() EndpointV3MapOutput
	ToEndpointV3MapOutputWithContext(context.Context) EndpointV3MapOutput
}

type EndpointV3Map map[string]EndpointV3Input

func (EndpointV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointV3)(nil)).Elem()
}

func (i EndpointV3Map) ToEndpointV3MapOutput() EndpointV3MapOutput {
	return i.ToEndpointV3MapOutputWithContext(context.Background())
}

func (i EndpointV3Map) ToEndpointV3MapOutputWithContext(ctx context.Context) EndpointV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointV3MapOutput)
}

type EndpointV3Output struct{ *pulumi.OutputState }

func (EndpointV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointV3)(nil)).Elem()
}

func (o EndpointV3Output) ToEndpointV3Output() EndpointV3Output {
	return o
}

func (o EndpointV3Output) ToEndpointV3OutputWithContext(ctx context.Context) EndpointV3Output {
	return o
}

// The endpoint region. The `region` and
// `endpointRegion` can be different.
func (o EndpointV3Output) EndpointRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.EndpointRegion }).(pulumi.StringOutput)
}

// The endpoint interface. Valid values are `public`,
// `internal` and `admin`. Default value is `public`
func (o EndpointV3Output) Interface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringPtrOutput { return v.Interface }).(pulumi.StringPtrOutput)
}

// The endpoint name.
func (o EndpointV3Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used.
func (o EndpointV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The endpoint service ID.
func (o EndpointV3Output) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The service name of the endpoint.
func (o EndpointV3Output) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The service type of the endpoint.
func (o EndpointV3Output) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

// The endpoint url.
func (o EndpointV3Output) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointV3) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type EndpointV3ArrayOutput struct{ *pulumi.OutputState }

func (EndpointV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointV3)(nil)).Elem()
}

func (o EndpointV3ArrayOutput) ToEndpointV3ArrayOutput() EndpointV3ArrayOutput {
	return o
}

func (o EndpointV3ArrayOutput) ToEndpointV3ArrayOutputWithContext(ctx context.Context) EndpointV3ArrayOutput {
	return o
}

func (o EndpointV3ArrayOutput) Index(i pulumi.IntInput) EndpointV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointV3 {
		return vs[0].([]*EndpointV3)[vs[1].(int)]
	}).(EndpointV3Output)
}

type EndpointV3MapOutput struct{ *pulumi.OutputState }

func (EndpointV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointV3)(nil)).Elem()
}

func (o EndpointV3MapOutput) ToEndpointV3MapOutput() EndpointV3MapOutput {
	return o
}

func (o EndpointV3MapOutput) ToEndpointV3MapOutputWithContext(ctx context.Context) EndpointV3MapOutput {
	return o
}

func (o EndpointV3MapOutput) MapIndex(k pulumi.StringInput) EndpointV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointV3 {
		return vs[0].(map[string]*EndpointV3)[vs[1].(string)]
	}).(EndpointV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointV3Input)(nil)).Elem(), &EndpointV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointV3ArrayInput)(nil)).Elem(), EndpointV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointV3MapInput)(nil)).Elem(), EndpointV3Map{})
	pulumi.RegisterOutputType(EndpointV3Output{})
	pulumi.RegisterOutputType(EndpointV3ArrayOutput{})
	pulumi.RegisterOutputType(EndpointV3MapOutput{})
}
