// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 listener resource within OpenStack.
//
// > **Note:** This resource has attributes that depend on octavia minor versions.
// Please ensure your Openstack cloud supports the required minor version.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loadbalancer.NewListener(ctx, "listener1", &loadbalancer.ListenerArgs{
// 			InsertHeaders: pulumi.AnyMap{
// 				"X-Forwarded-For": pulumi.Any("true"),
// 			},
// 			LoadbalancerId: pulumi.String("d9415786-5f1a-428b-b35f-2f1523e146d2"),
// 			Protocol:       pulumi.String("HTTP"),
// 			ProtocolPort:   pulumi.Int(8080),
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
// Load Balancer Listener can be imported using the Listener ID, e.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/listener:Listener listener_1 b67ce64e-8b26-405d-afeb-4a078901f15a
// ```
type Listener struct {
	pulumi.CustomResourceState

	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// A list of CIDR blocks that are permitted to connect to this listener, denying
	// all other source addresses. If not present, defaults to allow all.
	AllowedCidrs pulumi.StringArrayOutput `pulumi:"allowedCidrs"`
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit pulumi.IntOutput `pulumi:"connectionLimit"`
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId pulumi.StringOutput `pulumi:"defaultPoolId"`
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef pulumi.StringPtrOutput `pulumi:"defaultTlsContainerRef"`
	// Human-readable description for the Listener.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The list of key value pairs representing headers to insert
	// into the request before it is sent to the backend members. Changing this updates the headers of the
	// existing listener.
	InsertHeaders pulumi.MapOutput `pulumi:"insertHeaders"`
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId pulumi.StringOutput `pulumi:"loadbalancerId"`
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol - can either be TCP, HTTP, HTTPS,
	// TERMINATED_HTTPS, UDP (supported only in Octavia) or SCTP (supported only
	// in **Octavia minor version >= 2.23**). Changing this creates a new Listener.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort pulumi.IntOutput `pulumi:"protocolPort"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs pulumi.StringArrayOutput `pulumi:"sniContainerRefs"`
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The client inactivity timeout in milliseconds.
	TimeoutClientData pulumi.IntOutput `pulumi:"timeoutClientData"`
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect pulumi.IntOutput `pulumi:"timeoutMemberConnect"`
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData pulumi.IntOutput `pulumi:"timeoutMemberData"`
	// The time in milliseconds, to wait for additional
	// TCP packets for content inspection.
	TimeoutTcpInspect pulumi.IntOutput `pulumi:"timeoutTcpInspect"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadbalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadbalancerId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ProtocolPort == nil {
		return nil, errors.New("invalid value for required argument 'ProtocolPort'")
	}
	var resource Listener
	err := ctx.RegisterResource("openstack:loadbalancer/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("openstack:loadbalancer/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// A list of CIDR blocks that are permitted to connect to this listener, denying
	// all other source addresses. If not present, defaults to allow all.
	AllowedCidrs []string `pulumi:"allowedCidrs"`
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId *string `pulumi:"defaultPoolId"`
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef *string `pulumi:"defaultTlsContainerRef"`
	// Human-readable description for the Listener.
	Description *string `pulumi:"description"`
	// The list of key value pairs representing headers to insert
	// into the request before it is sent to the backend members. Changing this updates the headers of the
	// existing listener.
	InsertHeaders map[string]interface{} `pulumi:"insertHeaders"`
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The protocol - can either be TCP, HTTP, HTTPS,
	// TERMINATED_HTTPS, UDP (supported only in Octavia) or SCTP (supported only
	// in **Octavia minor version >= 2.23**). Changing this creates a new Listener.
	Protocol *string `pulumi:"protocol"`
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort *int `pulumi:"protocolPort"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region *string `pulumi:"region"`
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs []string `pulumi:"sniContainerRefs"`
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId *string `pulumi:"tenantId"`
	// The client inactivity timeout in milliseconds.
	TimeoutClientData *int `pulumi:"timeoutClientData"`
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect *int `pulumi:"timeoutMemberConnect"`
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData *int `pulumi:"timeoutMemberData"`
	// The time in milliseconds, to wait for additional
	// TCP packets for content inspection.
	TimeoutTcpInspect *int `pulumi:"timeoutTcpInspect"`
}

type ListenerState struct {
	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// A list of CIDR blocks that are permitted to connect to this listener, denying
	// all other source addresses. If not present, defaults to allow all.
	AllowedCidrs pulumi.StringArrayInput
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit pulumi.IntPtrInput
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId pulumi.StringPtrInput
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef pulumi.StringPtrInput
	// Human-readable description for the Listener.
	Description pulumi.StringPtrInput
	// The list of key value pairs representing headers to insert
	// into the request before it is sent to the backend members. Changing this updates the headers of the
	// existing listener.
	InsertHeaders pulumi.MapInput
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The protocol - can either be TCP, HTTP, HTTPS,
	// TERMINATED_HTTPS, UDP (supported only in Octavia) or SCTP (supported only
	// in **Octavia minor version >= 2.23**). Changing this creates a new Listener.
	Protocol pulumi.StringPtrInput
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort pulumi.IntPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region pulumi.StringPtrInput
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs pulumi.StringArrayInput
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId pulumi.StringPtrInput
	// The client inactivity timeout in milliseconds.
	TimeoutClientData pulumi.IntPtrInput
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect pulumi.IntPtrInput
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData pulumi.IntPtrInput
	// The time in milliseconds, to wait for additional
	// TCP packets for content inspection.
	TimeoutTcpInspect pulumi.IntPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// A list of CIDR blocks that are permitted to connect to this listener, denying
	// all other source addresses. If not present, defaults to allow all.
	AllowedCidrs []string `pulumi:"allowedCidrs"`
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId *string `pulumi:"defaultPoolId"`
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef *string `pulumi:"defaultTlsContainerRef"`
	// Human-readable description for the Listener.
	Description *string `pulumi:"description"`
	// The list of key value pairs representing headers to insert
	// into the request before it is sent to the backend members. Changing this updates the headers of the
	// existing listener.
	InsertHeaders map[string]interface{} `pulumi:"insertHeaders"`
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId string `pulumi:"loadbalancerId"`
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The protocol - can either be TCP, HTTP, HTTPS,
	// TERMINATED_HTTPS, UDP (supported only in Octavia) or SCTP (supported only
	// in **Octavia minor version >= 2.23**). Changing this creates a new Listener.
	Protocol string `pulumi:"protocol"`
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort int `pulumi:"protocolPort"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region *string `pulumi:"region"`
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs []string `pulumi:"sniContainerRefs"`
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId *string `pulumi:"tenantId"`
	// The client inactivity timeout in milliseconds.
	TimeoutClientData *int `pulumi:"timeoutClientData"`
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect *int `pulumi:"timeoutMemberConnect"`
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData *int `pulumi:"timeoutMemberData"`
	// The time in milliseconds, to wait for additional
	// TCP packets for content inspection.
	TimeoutTcpInspect *int `pulumi:"timeoutTcpInspect"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// A list of CIDR blocks that are permitted to connect to this listener, denying
	// all other source addresses. If not present, defaults to allow all.
	AllowedCidrs pulumi.StringArrayInput
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit pulumi.IntPtrInput
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId pulumi.StringPtrInput
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef pulumi.StringPtrInput
	// Human-readable description for the Listener.
	Description pulumi.StringPtrInput
	// The list of key value pairs representing headers to insert
	// into the request before it is sent to the backend members. Changing this updates the headers of the
	// existing listener.
	InsertHeaders pulumi.MapInput
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId pulumi.StringInput
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The protocol - can either be TCP, HTTP, HTTPS,
	// TERMINATED_HTTPS, UDP (supported only in Octavia) or SCTP (supported only
	// in **Octavia minor version >= 2.23**). Changing this creates a new Listener.
	Protocol pulumi.StringInput
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort pulumi.IntInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region pulumi.StringPtrInput
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs pulumi.StringArrayInput
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId pulumi.StringPtrInput
	// The client inactivity timeout in milliseconds.
	TimeoutClientData pulumi.IntPtrInput
	// The member connection timeout in milliseconds.
	TimeoutMemberConnect pulumi.IntPtrInput
	// The member inactivity timeout in milliseconds.
	TimeoutMemberData pulumi.IntPtrInput
	// The time in milliseconds, to wait for additional
	// TCP packets for content inspection.
	TimeoutTcpInspect pulumi.IntPtrInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

// ListenerArrayInput is an input type that accepts ListenerArray and ListenerArrayOutput values.
// You can construct a concrete instance of `ListenerArrayInput` via:
//
//          ListenerArray{ ListenerArgs{...} }
type ListenerArrayInput interface {
	pulumi.Input

	ToListenerArrayOutput() ListenerArrayOutput
	ToListenerArrayOutputWithContext(context.Context) ListenerArrayOutput
}

type ListenerArray []ListenerInput

func (ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (i ListenerArray) ToListenerArrayOutput() ListenerArrayOutput {
	return i.ToListenerArrayOutputWithContext(context.Background())
}

func (i ListenerArray) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerArrayOutput)
}

// ListenerMapInput is an input type that accepts ListenerMap and ListenerMapOutput values.
// You can construct a concrete instance of `ListenerMapInput` via:
//
//          ListenerMap{ "key": ListenerArgs{...} }
type ListenerMapInput interface {
	pulumi.Input

	ToListenerMapOutput() ListenerMapOutput
	ToListenerMapOutputWithContext(context.Context) ListenerMapOutput
}

type ListenerMap map[string]ListenerInput

func (ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (i ListenerMap) ToListenerMapOutput() ListenerMapOutput {
	return i.ToListenerMapOutputWithContext(context.Background())
}

func (i ListenerMap) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerMapOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

// The administrative state of the Listener.
// A valid value is true (UP) or false (DOWN).
func (o ListenerOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// A list of CIDR blocks that are permitted to connect to this listener, denying
// all other source addresses. If not present, defaults to allow all.
func (o ListenerOutput) AllowedCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringArrayOutput { return v.AllowedCidrs }).(pulumi.StringArrayOutput)
}

// The maximum number of connections allowed
// for the Listener.
func (o ListenerOutput) ConnectionLimit() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.ConnectionLimit }).(pulumi.IntOutput)
}

// The ID of the default pool with which the
// Listener is associated.
func (o ListenerOutput) DefaultPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.DefaultPoolId }).(pulumi.StringOutput)
}

// A reference to a Barbican Secrets
// container which stores TLS information. This is required if the protocol
// is `TERMINATED_HTTPS`. See
// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
// for more information.
func (o ListenerOutput) DefaultTlsContainerRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.DefaultTlsContainerRef }).(pulumi.StringPtrOutput)
}

// Human-readable description for the Listener.
func (o ListenerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The list of key value pairs representing headers to insert
// into the request before it is sent to the backend members. Changing this updates the headers of the
// existing listener.
func (o ListenerOutput) InsertHeaders() pulumi.MapOutput {
	return o.ApplyT(func(v *Listener) pulumi.MapOutput { return v.InsertHeaders }).(pulumi.MapOutput)
}

// The load balancer on which to provision this
// Listener. Changing this creates a new Listener.
func (o ListenerOutput) LoadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.LoadbalancerId }).(pulumi.StringOutput)
}

// Human-readable name for the Listener. Does not have
// to be unique.
func (o ListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol - can either be TCP, HTTP, HTTPS,
// TERMINATED_HTTPS, UDP (supported only in Octavia) or SCTP (supported only
// in **Octavia minor version >= 2.23**). Changing this creates a new Listener.
func (o ListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The port on which to listen for client traffic.
// Changing this creates a new Listener.
func (o ListenerOutput) ProtocolPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.ProtocolPort }).(pulumi.IntOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an . If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// Listener.
func (o ListenerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of references to Barbican Secrets
// containers which store SNI information. See
// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
// for more information.
func (o ListenerOutput) SniContainerRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringArrayOutput { return v.SniContainerRefs }).(pulumi.StringArrayOutput)
}

// Required for admins. The UUID of the tenant who owns
// the Listener.  Only administrative users can specify a tenant UUID
// other than their own. Changing this creates a new Listener.
func (o ListenerOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The client inactivity timeout in milliseconds.
func (o ListenerOutput) TimeoutClientData() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.TimeoutClientData }).(pulumi.IntOutput)
}

// The member connection timeout in milliseconds.
func (o ListenerOutput) TimeoutMemberConnect() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.TimeoutMemberConnect }).(pulumi.IntOutput)
}

// The member inactivity timeout in milliseconds.
func (o ListenerOutput) TimeoutMemberData() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.TimeoutMemberData }).(pulumi.IntOutput)
}

// The time in milliseconds, to wait for additional
// TCP packets for content inspection.
func (o ListenerOutput) TimeoutTcpInspect() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.TimeoutTcpInspect }).(pulumi.IntOutput)
}

type ListenerArrayOutput struct{ *pulumi.OutputState }

func (ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (o ListenerArrayOutput) ToListenerArrayOutput() ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) Index(i pulumi.IntInput) ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].([]*Listener)[vs[1].(int)]
	}).(ListenerOutput)
}

type ListenerMapOutput struct{ *pulumi.OutputState }

func (ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (o ListenerMapOutput) ToListenerMapOutput() ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) MapIndex(k pulumi.StringInput) ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].(map[string]*Listener)[vs[1].(string)]
	}).(ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerArrayInput)(nil)).Elem(), ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerMapInput)(nil)).Elem(), ListenerMap{})
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerArrayOutput{})
	pulumi.RegisterOutputType(ListenerMapOutput{})
}
