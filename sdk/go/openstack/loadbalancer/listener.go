// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 listener resource within OpenStack.
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
// 		_, err := loadbalancer.NewListener(ctx, "listener1", &loadbalancer.ListenerArgs{
// 			InsertHeaders: pulumi.StringMap{
// 				"X-Forwarded-For": pulumi.String("true"),
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
type Listener struct {
	pulumi.CustomResourceState

	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
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
	// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
	// new Listener.
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
	if args == nil || args.LoadbalancerId == nil {
		return nil, errors.New("missing required argument 'LoadbalancerId'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ProtocolPort == nil {
		return nil, errors.New("missing required argument 'ProtocolPort'")
	}
	if args == nil {
		args = &ListenerArgs{}
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
	// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
	// new Listener.
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
	// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
	// new Listener.
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
	// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
	// new Listener.
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
	// TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
	// new Listener.
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
