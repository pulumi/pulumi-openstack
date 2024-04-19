// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 load balancer monitor resource within OpenStack.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/loadbalancer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := loadbalancer.NewMonitorV1(ctx, "monitor_1", &loadbalancer.MonitorV1Args{
//				Type:         pulumi.String("PING"),
//				Delay:        pulumi.Int(30),
//				Timeout:      pulumi.Int(5),
//				MaxRetries:   pulumi.Int(3),
//				AdminStateUp: pulumi.String("true"),
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
// Load Balancer Members can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:loadbalancer/monitorV1:MonitorV1 monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c
// ```
type MonitorV1 struct {
	pulumi.CustomResourceState

	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp pulumi.StringOutput `pulumi:"adminStateUp"`
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay pulumi.IntOutput `pulumi:"delay"`
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expectedCodes
	// of the existing monitor.
	ExpectedCodes pulumi.StringPtrOutput `pulumi:"expectedCodes"`
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the httpMethod of the existing monitor.
	HttpMethod pulumi.StringPtrOutput `pulumi:"httpMethod"`
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the maxRetries of the existing monitor.
	MaxRetries pulumi.IntOutput `pulumi:"maxRetries"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type pulumi.StringOutput `pulumi:"type"`
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// urlPath of the existing monitor.
	UrlPath pulumi.StringPtrOutput `pulumi:"urlPath"`
}

// NewMonitorV1 registers a new resource with the given unique name, arguments, and options.
func NewMonitorV1(ctx *pulumi.Context,
	name string, args *MonitorV1Args, opts ...pulumi.ResourceOption) (*MonitorV1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Delay == nil {
		return nil, errors.New("invalid value for required argument 'Delay'")
	}
	if args.MaxRetries == nil {
		return nil, errors.New("invalid value for required argument 'MaxRetries'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MonitorV1
	err := ctx.RegisterResource("openstack:loadbalancer/monitorV1:MonitorV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitorV1 gets an existing MonitorV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitorV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorV1State, opts ...pulumi.ResourceOption) (*MonitorV1, error) {
	var resource MonitorV1
	err := ctx.ReadResource("openstack:loadbalancer/monitorV1:MonitorV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitorV1 resources.
type monitorV1State struct {
	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp *string `pulumi:"adminStateUp"`
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay *int `pulumi:"delay"`
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expectedCodes
	// of the existing monitor.
	ExpectedCodes *string `pulumi:"expectedCodes"`
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the httpMethod of the existing monitor.
	HttpMethod *string `pulumi:"httpMethod"`
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the maxRetries of the existing monitor.
	MaxRetries *int `pulumi:"maxRetries"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region *string `pulumi:"region"`
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId *string `pulumi:"tenantId"`
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout *int `pulumi:"timeout"`
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type *string `pulumi:"type"`
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// urlPath of the existing monitor.
	UrlPath *string `pulumi:"urlPath"`
}

type MonitorV1State struct {
	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp pulumi.StringPtrInput
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay pulumi.IntPtrInput
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expectedCodes
	// of the existing monitor.
	ExpectedCodes pulumi.StringPtrInput
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the httpMethod of the existing monitor.
	HttpMethod pulumi.StringPtrInput
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the maxRetries of the existing monitor.
	MaxRetries pulumi.IntPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region pulumi.StringPtrInput
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId pulumi.StringPtrInput
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout pulumi.IntPtrInput
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type pulumi.StringPtrInput
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// urlPath of the existing monitor.
	UrlPath pulumi.StringPtrInput
}

func (MonitorV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorV1State)(nil)).Elem()
}

type monitorV1Args struct {
	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp *string `pulumi:"adminStateUp"`
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay int `pulumi:"delay"`
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expectedCodes
	// of the existing monitor.
	ExpectedCodes *string `pulumi:"expectedCodes"`
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the httpMethod of the existing monitor.
	HttpMethod *string `pulumi:"httpMethod"`
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the maxRetries of the existing monitor.
	MaxRetries int `pulumi:"maxRetries"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region *string `pulumi:"region"`
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId *string `pulumi:"tenantId"`
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout int `pulumi:"timeout"`
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type string `pulumi:"type"`
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// urlPath of the existing monitor.
	UrlPath *string `pulumi:"urlPath"`
}

// The set of arguments for constructing a MonitorV1 resource.
type MonitorV1Args struct {
	// The administrative state of the monitor.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing monitor.
	AdminStateUp pulumi.StringPtrInput
	// The time, in seconds, between sending probes to members.
	// Changing this creates a new monitor.
	Delay pulumi.IntInput
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202". Changing this updates the expectedCodes
	// of the existing monitor.
	ExpectedCodes pulumi.StringPtrInput
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it defaults
	// to "GET". Changing this updates the httpMethod of the existing monitor.
	HttpMethod pulumi.StringPtrInput
	// Number of permissible ping failures before changing
	// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
	// this updates the maxRetries of the existing monitor.
	MaxRetries pulumi.IntInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB monitor. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB monitor.
	Region pulumi.StringPtrInput
	// The owner of the monitor. Required if admin wants to
	// create a monitor for another tenant. Changing this creates a new monitor.
	TenantId pulumi.StringPtrInput
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay value.
	// Changing this updates the timeout of the existing monitor.
	Timeout pulumi.IntInput
	// The type of probe, which is PING, TCP, HTTP, or HTTPS,
	// that is sent by the monitor to verify the member state. Changing this
	// creates a new monitor.
	Type pulumi.StringInput
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS. Changing this updates the
	// urlPath of the existing monitor.
	UrlPath pulumi.StringPtrInput
}

func (MonitorV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorV1Args)(nil)).Elem()
}

type MonitorV1Input interface {
	pulumi.Input

	ToMonitorV1Output() MonitorV1Output
	ToMonitorV1OutputWithContext(ctx context.Context) MonitorV1Output
}

func (*MonitorV1) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorV1)(nil)).Elem()
}

func (i *MonitorV1) ToMonitorV1Output() MonitorV1Output {
	return i.ToMonitorV1OutputWithContext(context.Background())
}

func (i *MonitorV1) ToMonitorV1OutputWithContext(ctx context.Context) MonitorV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorV1Output)
}

// MonitorV1ArrayInput is an input type that accepts MonitorV1Array and MonitorV1ArrayOutput values.
// You can construct a concrete instance of `MonitorV1ArrayInput` via:
//
//	MonitorV1Array{ MonitorV1Args{...} }
type MonitorV1ArrayInput interface {
	pulumi.Input

	ToMonitorV1ArrayOutput() MonitorV1ArrayOutput
	ToMonitorV1ArrayOutputWithContext(context.Context) MonitorV1ArrayOutput
}

type MonitorV1Array []MonitorV1Input

func (MonitorV1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitorV1)(nil)).Elem()
}

func (i MonitorV1Array) ToMonitorV1ArrayOutput() MonitorV1ArrayOutput {
	return i.ToMonitorV1ArrayOutputWithContext(context.Background())
}

func (i MonitorV1Array) ToMonitorV1ArrayOutputWithContext(ctx context.Context) MonitorV1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorV1ArrayOutput)
}

// MonitorV1MapInput is an input type that accepts MonitorV1Map and MonitorV1MapOutput values.
// You can construct a concrete instance of `MonitorV1MapInput` via:
//
//	MonitorV1Map{ "key": MonitorV1Args{...} }
type MonitorV1MapInput interface {
	pulumi.Input

	ToMonitorV1MapOutput() MonitorV1MapOutput
	ToMonitorV1MapOutputWithContext(context.Context) MonitorV1MapOutput
}

type MonitorV1Map map[string]MonitorV1Input

func (MonitorV1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitorV1)(nil)).Elem()
}

func (i MonitorV1Map) ToMonitorV1MapOutput() MonitorV1MapOutput {
	return i.ToMonitorV1MapOutputWithContext(context.Background())
}

func (i MonitorV1Map) ToMonitorV1MapOutputWithContext(ctx context.Context) MonitorV1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorV1MapOutput)
}

type MonitorV1Output struct{ *pulumi.OutputState }

func (MonitorV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorV1)(nil)).Elem()
}

func (o MonitorV1Output) ToMonitorV1Output() MonitorV1Output {
	return o
}

func (o MonitorV1Output) ToMonitorV1OutputWithContext(ctx context.Context) MonitorV1Output {
	return o
}

// The administrative state of the monitor.
// Acceptable values are "true" and "false". Changing this value updates the
// state of the existing monitor.
func (o MonitorV1Output) AdminStateUp() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringOutput { return v.AdminStateUp }).(pulumi.StringOutput)
}

// The time, in seconds, between sending probes to members.
// Changing this creates a new monitor.
func (o MonitorV1Output) Delay() pulumi.IntOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.IntOutput { return v.Delay }).(pulumi.IntOutput)
}

// Required for HTTP(S) types. Expected HTTP codes
// for a passing HTTP(S) monitor. You can either specify a single status like
// "200", or a range like "200-202". Changing this updates the expectedCodes
// of the existing monitor.
func (o MonitorV1Output) ExpectedCodes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringPtrOutput { return v.ExpectedCodes }).(pulumi.StringPtrOutput)
}

// Required for HTTP(S) types. The HTTP method used
// for requests by the monitor. If this attribute is not specified, it defaults
// to "GET". Changing this updates the httpMethod of the existing monitor.
func (o MonitorV1Output) HttpMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringPtrOutput { return v.HttpMethod }).(pulumi.StringPtrOutput)
}

// Number of permissible ping failures before changing
// the member's status to INACTIVE. Must be a number between 1 and 10. Changing
// this updates the maxRetries of the existing monitor.
func (o MonitorV1Output) MaxRetries() pulumi.IntOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.IntOutput { return v.MaxRetries }).(pulumi.IntOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an LB monitor. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// LB monitor.
func (o MonitorV1Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The owner of the monitor. Required if admin wants to
// create a monitor for another tenant. Changing this creates a new monitor.
func (o MonitorV1Output) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Maximum number of seconds for a monitor to wait for a
// ping reply before it times out. The value must be less than the delay value.
// Changing this updates the timeout of the existing monitor.
func (o MonitorV1Output) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

// The type of probe, which is PING, TCP, HTTP, or HTTPS,
// that is sent by the monitor to verify the member state. Changing this
// creates a new monitor.
func (o MonitorV1Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Required for HTTP(S) types. URI path that will be
// accessed if monitor type is HTTP or HTTPS. Changing this updates the
// urlPath of the existing monitor.
func (o MonitorV1Output) UrlPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorV1) pulumi.StringPtrOutput { return v.UrlPath }).(pulumi.StringPtrOutput)
}

type MonitorV1ArrayOutput struct{ *pulumi.OutputState }

func (MonitorV1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitorV1)(nil)).Elem()
}

func (o MonitorV1ArrayOutput) ToMonitorV1ArrayOutput() MonitorV1ArrayOutput {
	return o
}

func (o MonitorV1ArrayOutput) ToMonitorV1ArrayOutputWithContext(ctx context.Context) MonitorV1ArrayOutput {
	return o
}

func (o MonitorV1ArrayOutput) Index(i pulumi.IntInput) MonitorV1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MonitorV1 {
		return vs[0].([]*MonitorV1)[vs[1].(int)]
	}).(MonitorV1Output)
}

type MonitorV1MapOutput struct{ *pulumi.OutputState }

func (MonitorV1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitorV1)(nil)).Elem()
}

func (o MonitorV1MapOutput) ToMonitorV1MapOutput() MonitorV1MapOutput {
	return o
}

func (o MonitorV1MapOutput) ToMonitorV1MapOutputWithContext(ctx context.Context) MonitorV1MapOutput {
	return o
}

func (o MonitorV1MapOutput) MapIndex(k pulumi.StringInput) MonitorV1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MonitorV1 {
		return vs[0].(map[string]*MonitorV1)[vs[1].(string)]
	}).(MonitorV1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorV1Input)(nil)).Elem(), &MonitorV1{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorV1ArrayInput)(nil)).Elem(), MonitorV1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorV1MapInput)(nil)).Elem(), MonitorV1Map{})
	pulumi.RegisterOutputType(MonitorV1Output{})
	pulumi.RegisterOutputType(MonitorV1ArrayOutput{})
	pulumi.RegisterOutputType(MonitorV1MapOutput{})
}
