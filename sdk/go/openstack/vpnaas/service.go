// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V2 Neutron VPN service resource within OpenStack.
//
// ## Import
//
// Services can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:vpnaas/service:Service service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
//
// ```
type Service struct {
	pulumi.CustomResourceState

	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The read-only external (public) IPv4 address that is used for the VPN service.
	ExternalV4Ip pulumi.StringOutput `pulumi:"externalV4Ip"`
	// The read-only external (public) IPv6 address that is used for the VPN service.
	ExternalV6Ip pulumi.StringOutput `pulumi:"externalV6Ip"`
	// The name of the service. Changing this updates the name of
	// the existing service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the router. Changing this creates a new service.
	RouterId pulumi.StringOutput `pulumi:"routerId"`
	// Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
	Status pulumi.StringOutput `pulumi:"status"`
	// SubnetID is the ID of the subnet. Default is null.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouterId == nil {
		return nil, errors.New("invalid value for required argument 'RouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("openstack:vpnaas/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("openstack:vpnaas/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description *string `pulumi:"description"`
	// The read-only external (public) IPv4 address that is used for the VPN service.
	ExternalV4Ip *string `pulumi:"externalV4Ip"`
	// The read-only external (public) IPv6 address that is used for the VPN service.
	ExternalV6Ip *string `pulumi:"externalV6Ip"`
	// The name of the service. Changing this updates the name of
	// the existing service.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region *string `pulumi:"region"`
	// The ID of the router. Changing this creates a new service.
	RouterId *string `pulumi:"routerId"`
	// Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
	Status *string `pulumi:"status"`
	// SubnetID is the ID of the subnet. Default is null.
	SubnetId *string `pulumi:"subnetId"`
	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type ServiceState struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp pulumi.BoolPtrInput
	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description pulumi.StringPtrInput
	// The read-only external (public) IPv4 address that is used for the VPN service.
	ExternalV4Ip pulumi.StringPtrInput
	// The read-only external (public) IPv6 address that is used for the VPN service.
	ExternalV6Ip pulumi.StringPtrInput
	// The name of the service. Changing this updates the name of
	// the existing service.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region pulumi.StringPtrInput
	// The ID of the router. Changing this creates a new service.
	RouterId pulumi.StringPtrInput
	// Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
	Status pulumi.StringPtrInput
	// SubnetID is the ID of the subnet. Default is null.
	SubnetId pulumi.StringPtrInput
	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description *string `pulumi:"description"`
	// The name of the service. Changing this updates the name of
	// the existing service.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region *string `pulumi:"region"`
	// The ID of the router. Changing this creates a new service.
	RouterId string `pulumi:"routerId"`
	// SubnetID is the ID of the subnet. Default is null.
	SubnetId *string `pulumi:"subnetId"`
	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The administrative state of the resource. Can either be up(true) or down(false).
	// Changing this updates the administrative state of the existing service.
	AdminStateUp pulumi.BoolPtrInput
	// The human-readable description for the service.
	// Changing this updates the description of the existing service.
	Description pulumi.StringPtrInput
	// The name of the service. Changing this updates the name of
	// the existing service.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region pulumi.StringPtrInput
	// The ID of the router. Changing this creates a new service.
	RouterId pulumi.StringInput
	// SubnetID is the ID of the subnet. Default is null.
	SubnetId pulumi.StringPtrInput
	// The owner of the service. Required if admin wants to
	// create a service for another project. Changing this creates a new service.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

func (i *Service) ToOutput(ctx context.Context) pulumix.Output[*Service] {
	return pulumix.Output[*Service]{
		OutputState: i.ToServiceOutputWithContext(ctx).OutputState,
	}
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

func (i ServiceArray) ToOutput(ctx context.Context) pulumix.Output[[]*Service] {
	return pulumix.Output[[]*Service]{
		OutputState: i.ToServiceArrayOutputWithContext(ctx).OutputState,
	}
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

func (i ServiceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Service] {
	return pulumix.Output[map[string]*Service]{
		OutputState: i.ToServiceMapOutputWithContext(ctx).OutputState,
	}
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

func (o ServiceOutput) ToOutput(ctx context.Context) pulumix.Output[*Service] {
	return pulumix.Output[*Service]{
		OutputState: o.OutputState,
	}
}

// The administrative state of the resource. Can either be up(true) or down(false).
// Changing this updates the administrative state of the existing service.
func (o ServiceOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// The human-readable description for the service.
// Changing this updates the description of the existing service.
func (o ServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The read-only external (public) IPv4 address that is used for the VPN service.
func (o ServiceOutput) ExternalV4Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ExternalV4Ip }).(pulumi.StringOutput)
}

// The read-only external (public) IPv6 address that is used for the VPN service.
func (o ServiceOutput) ExternalV6Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ExternalV6Ip }).(pulumi.StringOutput)
}

// The name of the service. Changing this updates the name of
// the existing service.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a VPN service. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// service.
func (o ServiceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the router. Changing this creates a new service.
func (o ServiceOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.RouterId }).(pulumi.StringOutput)
}

// Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
func (o ServiceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// SubnetID is the ID of the subnet. Default is null.
func (o ServiceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The owner of the service. Required if admin wants to
// create a service for another project. Changing this creates a new service.
func (o ServiceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Map of additional options.
func (o ServiceOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Service) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Service] {
	return pulumix.Output[[]*Service]{
		OutputState: o.OutputState,
	}
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Service] {
	return pulumix.Output[map[string]*Service]{
		OutputState: o.OutputState,
	}
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
