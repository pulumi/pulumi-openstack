// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 monitor resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loadbalancer.NewMonitor(ctx, "monitor1", &loadbalancer.MonitorArgs{
// 			Delay:      pulumi.Int(20),
// 			MaxRetries: pulumi.Int(5),
// 			PoolId:     pulumi.Any(openstack_lb_pool_v2.Pool_1.Id),
// 			Timeout:    pulumi.Int(10),
// 			Type:       pulumi.String("PING"),
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
// Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2
// ```
//
//  In case of using OpenContrail, the import may not work properly. If you face an issue, try to import the monitor providing its parent pool ID
//
// ```sh
//  $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2/708bc224-0f8c-4981-ac82-97095fe051b6
// ```
type Monitor struct {
	pulumi.CustomResourceState

	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The time, in seconds, between sending probes to members.
	Delay pulumi.IntOutput `pulumi:"delay"`
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202".
	ExpectedCodes pulumi.StringOutput `pulumi:"expectedCodes"`
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to "GET".
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1
	// and 10.
	MaxRetries pulumi.IntOutput `pulumi:"maxRetries"`
	// Number of permissible ping failures befor changing the member's
	// status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
	// Changing this updates the maxRetriesDown of the existing monitor.
	MaxRetriesDown pulumi.IntOutput `pulumi:"maxRetriesDown"`
	// The Name of the Monitor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the pool that this monitor will be assigned to.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// monitor.
	Region pulumi.StringOutput `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the monitor.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay
	// value.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// The type of probe, which is PING, TCP, HTTP, HTTPS,
	// TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
	// balancer to verify the member state. Changing this creates a new monitor.
	Type pulumi.StringOutput `pulumi:"type"`
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS.
	UrlPath pulumi.StringOutput `pulumi:"urlPath"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Delay == nil {
		return nil, errors.New("invalid value for required argument 'Delay'")
	}
	if args.MaxRetries == nil {
		return nil, errors.New("invalid value for required argument 'MaxRetries'")
	}
	if args.PoolId == nil {
		return nil, errors.New("invalid value for required argument 'PoolId'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Monitor
	err := ctx.RegisterResource("openstack:loadbalancer/monitor:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("openstack:loadbalancer/monitor:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The time, in seconds, between sending probes to members.
	Delay *int `pulumi:"delay"`
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202".
	ExpectedCodes *string `pulumi:"expectedCodes"`
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to "GET".
	HttpMethod *string `pulumi:"httpMethod"`
	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1
	// and 10.
	MaxRetries *int `pulumi:"maxRetries"`
	// Number of permissible ping failures befor changing the member's
	// status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
	// Changing this updates the maxRetriesDown of the existing monitor.
	MaxRetriesDown *int `pulumi:"maxRetriesDown"`
	// The Name of the Monitor.
	Name *string `pulumi:"name"`
	// The id of the pool that this monitor will be assigned to.
	PoolId *string `pulumi:"poolId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// monitor.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the monitor.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantId *string `pulumi:"tenantId"`
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay
	// value.
	Timeout *int `pulumi:"timeout"`
	// The type of probe, which is PING, TCP, HTTP, HTTPS,
	// TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
	// balancer to verify the member state. Changing this creates a new monitor.
	Type *string `pulumi:"type"`
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS.
	UrlPath *string `pulumi:"urlPath"`
}

type MonitorState struct {
	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// The time, in seconds, between sending probes to members.
	Delay pulumi.IntPtrInput
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202".
	ExpectedCodes pulumi.StringPtrInput
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to "GET".
	HttpMethod pulumi.StringPtrInput
	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1
	// and 10.
	MaxRetries pulumi.IntPtrInput
	// Number of permissible ping failures befor changing the member's
	// status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
	// Changing this updates the maxRetriesDown of the existing monitor.
	MaxRetriesDown pulumi.IntPtrInput
	// The Name of the Monitor.
	Name pulumi.StringPtrInput
	// The id of the pool that this monitor will be assigned to.
	PoolId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// monitor.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the monitor.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantId pulumi.StringPtrInput
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay
	// value.
	Timeout pulumi.IntPtrInput
	// The type of probe, which is PING, TCP, HTTP, HTTPS,
	// TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
	// balancer to verify the member state. Changing this creates a new monitor.
	Type pulumi.StringPtrInput
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS.
	UrlPath pulumi.StringPtrInput
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The time, in seconds, between sending probes to members.
	Delay int `pulumi:"delay"`
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202".
	ExpectedCodes *string `pulumi:"expectedCodes"`
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to "GET".
	HttpMethod *string `pulumi:"httpMethod"`
	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1
	// and 10.
	MaxRetries int `pulumi:"maxRetries"`
	// Number of permissible ping failures befor changing the member's
	// status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
	// Changing this updates the maxRetriesDown of the existing monitor.
	MaxRetriesDown *int `pulumi:"maxRetriesDown"`
	// The Name of the Monitor.
	Name *string `pulumi:"name"`
	// The id of the pool that this monitor will be assigned to.
	PoolId string `pulumi:"poolId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// monitor.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the monitor.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantId *string `pulumi:"tenantId"`
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay
	// value.
	Timeout int `pulumi:"timeout"`
	// The type of probe, which is PING, TCP, HTTP, HTTPS,
	// TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
	// balancer to verify the member state. Changing this creates a new monitor.
	Type string `pulumi:"type"`
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS.
	UrlPath *string `pulumi:"urlPath"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// The administrative state of the monitor.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// The time, in seconds, between sending probes to members.
	Delay pulumi.IntInput
	// Required for HTTP(S) types. Expected HTTP codes
	// for a passing HTTP(S) monitor. You can either specify a single status like
	// "200", or a range like "200-202".
	ExpectedCodes pulumi.StringPtrInput
	// Required for HTTP(S) types. The HTTP method used
	// for requests by the monitor. If this attribute is not specified, it
	// defaults to "GET".
	HttpMethod pulumi.StringPtrInput
	// Number of permissible ping failures before
	// changing the member's status to INACTIVE. Must be a number between 1
	// and 10.
	MaxRetries pulumi.IntInput
	// Number of permissible ping failures befor changing the member's
	// status to ERROR. Must be a number between 1 and 10 (supported only in Octavia).
	// Changing this updates the maxRetriesDown of the existing monitor.
	MaxRetriesDown pulumi.IntPtrInput
	// The Name of the Monitor.
	Name pulumi.StringPtrInput
	// The id of the pool that this monitor will be assigned to.
	PoolId pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// monitor.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the monitor.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new monitor.
	TenantId pulumi.StringPtrInput
	// Maximum number of seconds for a monitor to wait for a
	// ping reply before it times out. The value must be less than the delay
	// value.
	Timeout pulumi.IntInput
	// The type of probe, which is PING, TCP, HTTP, HTTPS,
	// TLS-HELLO or UDP-CONNECT (supported only in Octavia), that is sent by the load
	// balancer to verify the member state. Changing this creates a new monitor.
	Type pulumi.StringInput
	// Required for HTTP(S) types. URI path that will be
	// accessed if monitor type is HTTP or HTTPS.
	UrlPath pulumi.StringPtrInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((*Monitor)(nil))
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

func (i *Monitor) ToMonitorPtrOutput() MonitorPtrOutput {
	return i.ToMonitorPtrOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorPtrOutputWithContext(ctx context.Context) MonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPtrOutput)
}

type MonitorPtrInput interface {
	pulumi.Input

	ToMonitorPtrOutput() MonitorPtrOutput
	ToMonitorPtrOutputWithContext(ctx context.Context) MonitorPtrOutput
}

type monitorPtrType MonitorArgs

func (*monitorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil))
}

func (i *monitorPtrType) ToMonitorPtrOutput() MonitorPtrOutput {
	return i.ToMonitorPtrOutputWithContext(context.Background())
}

func (i *monitorPtrType) ToMonitorPtrOutputWithContext(ctx context.Context) MonitorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorPtrOutput)
}

// MonitorArrayInput is an input type that accepts MonitorArray and MonitorArrayOutput values.
// You can construct a concrete instance of `MonitorArrayInput` via:
//
//          MonitorArray{ MonitorArgs{...} }
type MonitorArrayInput interface {
	pulumi.Input

	ToMonitorArrayOutput() MonitorArrayOutput
	ToMonitorArrayOutputWithContext(context.Context) MonitorArrayOutput
}

type MonitorArray []MonitorInput

func (MonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Monitor)(nil))
}

func (i MonitorArray) ToMonitorArrayOutput() MonitorArrayOutput {
	return i.ToMonitorArrayOutputWithContext(context.Background())
}

func (i MonitorArray) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorArrayOutput)
}

// MonitorMapInput is an input type that accepts MonitorMap and MonitorMapOutput values.
// You can construct a concrete instance of `MonitorMapInput` via:
//
//          MonitorMap{ "key": MonitorArgs{...} }
type MonitorMapInput interface {
	pulumi.Input

	ToMonitorMapOutput() MonitorMapOutput
	ToMonitorMapOutputWithContext(context.Context) MonitorMapOutput
}

type MonitorMap map[string]MonitorInput

func (MonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Monitor)(nil))
}

func (i MonitorMap) ToMonitorMapOutput() MonitorMapOutput {
	return i.ToMonitorMapOutputWithContext(context.Background())
}

func (i MonitorMap) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorMapOutput)
}

type MonitorOutput struct {
	*pulumi.OutputState
}

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Monitor)(nil))
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorPtrOutput() MonitorPtrOutput {
	return o.ToMonitorPtrOutputWithContext(context.Background())
}

func (o MonitorOutput) ToMonitorPtrOutputWithContext(ctx context.Context) MonitorPtrOutput {
	return o.ApplyT(func(v Monitor) *Monitor {
		return &v
	}).(MonitorPtrOutput)
}

type MonitorPtrOutput struct {
	*pulumi.OutputState
}

func (MonitorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil))
}

func (o MonitorPtrOutput) ToMonitorPtrOutput() MonitorPtrOutput {
	return o
}

func (o MonitorPtrOutput) ToMonitorPtrOutputWithContext(ctx context.Context) MonitorPtrOutput {
	return o
}

type MonitorArrayOutput struct{ *pulumi.OutputState }

func (MonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Monitor)(nil))
}

func (o MonitorArrayOutput) ToMonitorArrayOutput() MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) Index(i pulumi.IntInput) MonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Monitor {
		return vs[0].([]Monitor)[vs[1].(int)]
	}).(MonitorOutput)
}

type MonitorMapOutput struct{ *pulumi.OutputState }

func (MonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Monitor)(nil))
}

func (o MonitorMapOutput) ToMonitorMapOutput() MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) MapIndex(k pulumi.StringInput) MonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Monitor {
		return vs[0].(map[string]Monitor)[vs[1].(string)]
	}).(MonitorOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitorOutput{})
	pulumi.RegisterOutputType(MonitorPtrOutput{})
	pulumi.RegisterOutputType(MonitorArrayOutput{})
	pulumi.RegisterOutputType(MonitorMapOutput{})
}
