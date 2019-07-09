// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 listener resource within OpenStack.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_listener_v2.html.markdown.
type Listener struct {
	s *pulumi.ResourceState
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOpt) (*Listener, error) {
	if args == nil || args.LoadbalancerId == nil {
		return nil, errors.New("missing required argument 'LoadbalancerId'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.ProtocolPort == nil {
		return nil, errors.New("missing required argument 'ProtocolPort'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminStateUp"] = nil
		inputs["connectionLimit"] = nil
		inputs["defaultPoolId"] = nil
		inputs["defaultTlsContainerRef"] = nil
		inputs["description"] = nil
		inputs["loadbalancerId"] = nil
		inputs["name"] = nil
		inputs["protocol"] = nil
		inputs["protocolPort"] = nil
		inputs["region"] = nil
		inputs["sniContainerRefs"] = nil
		inputs["tenantId"] = nil
	} else {
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["connectionLimit"] = args.ConnectionLimit
		inputs["defaultPoolId"] = args.DefaultPoolId
		inputs["defaultTlsContainerRef"] = args.DefaultTlsContainerRef
		inputs["description"] = args.Description
		inputs["loadbalancerId"] = args.LoadbalancerId
		inputs["name"] = args.Name
		inputs["protocol"] = args.Protocol
		inputs["protocolPort"] = args.ProtocolPort
		inputs["region"] = args.Region
		inputs["sniContainerRefs"] = args.SniContainerRefs
		inputs["tenantId"] = args.TenantId
	}
	s, err := ctx.RegisterResource("openstack:loadbalancer/listener:Listener", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Listener{s: s}, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ListenerState, opts ...pulumi.ResourceOpt) (*Listener, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminStateUp"] = state.AdminStateUp
		inputs["connectionLimit"] = state.ConnectionLimit
		inputs["defaultPoolId"] = state.DefaultPoolId
		inputs["defaultTlsContainerRef"] = state.DefaultTlsContainerRef
		inputs["description"] = state.Description
		inputs["loadbalancerId"] = state.LoadbalancerId
		inputs["name"] = state.Name
		inputs["protocol"] = state.Protocol
		inputs["protocolPort"] = state.ProtocolPort
		inputs["region"] = state.Region
		inputs["sniContainerRefs"] = state.SniContainerRefs
		inputs["tenantId"] = state.TenantId
	}
	s, err := ctx.ReadResource("openstack:loadbalancer/listener:Listener", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Listener{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Listener) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Listener) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The administrative state of the Listener.
// A valid value is true (UP) or false (DOWN).
func (r *Listener) AdminStateUp() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["adminStateUp"])
}

// The maximum number of connections allowed
// for the Listener.
func (r *Listener) ConnectionLimit() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["connectionLimit"])
}

// The ID of the default pool with which the
// Listener is associated.
func (r *Listener) DefaultPoolId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultPoolId"])
}

// A reference to a Barbican Secrets
// container which stores TLS information. This is required if the protocol
// is `TERMINATED_HTTPS`. See
// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
// for more information.
func (r *Listener) DefaultTlsContainerRef() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultTlsContainerRef"])
}

// Human-readable description for the Listener.
func (r *Listener) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The load balancer on which to provision this
// Listener. Changing this creates a new Listener.
func (r *Listener) LoadbalancerId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["loadbalancerId"])
}

// Human-readable name for the Listener. Does not have
// to be unique.
func (r *Listener) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
// Changing this creates a new Listener.
func (r *Listener) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// The port on which to listen for client traffic.
// Changing this creates a new Listener.
func (r *Listener) ProtocolPort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["protocolPort"])
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an . If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// Listener.
func (r *Listener) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// A list of references to Barbican Secrets
// containers which store SNI information. See
// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
// for more information.
func (r *Listener) SniContainerRefs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sniContainerRefs"])
}

// Required for admins. The UUID of the tenant who owns
// the Listener.  Only administrative users can specify a tenant UUID
// other than their own. Changing this creates a new Listener.
func (r *Listener) TenantId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tenantId"])
}

// Input properties used for looking up and filtering Listener resources.
type ListenerState struct {
	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp interface{}
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit interface{}
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId interface{}
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef interface{}
	// Human-readable description for the Listener.
	Description interface{}
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId interface{}
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name interface{}
	// The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
	// Changing this creates a new Listener.
	Protocol interface{}
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region interface{}
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs interface{}
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId interface{}
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The administrative state of the Listener.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp interface{}
	// The maximum number of connections allowed
	// for the Listener.
	ConnectionLimit interface{}
	// The ID of the default pool with which the
	// Listener is associated.
	DefaultPoolId interface{}
	// A reference to a Barbican Secrets
	// container which stores TLS information. This is required if the protocol
	// is `TERMINATED_HTTPS`. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	DefaultTlsContainerRef interface{}
	// Human-readable description for the Listener.
	Description interface{}
	// The load balancer on which to provision this
	// Listener. Changing this creates a new Listener.
	LoadbalancerId interface{}
	// Human-readable name for the Listener. Does not have
	// to be unique.
	Name interface{}
	// The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
	// Changing this creates a new Listener.
	Protocol interface{}
	// The port on which to listen for client traffic.
	// Changing this creates a new Listener.
	ProtocolPort interface{}
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// Listener.
	Region interface{}
	// A list of references to Barbican Secrets
	// containers which store SNI information. See
	// [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
	// for more information.
	SniContainerRefs interface{}
	// Required for admins. The UUID of the tenant who owns
	// the Listener.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new Listener.
	TenantId interface{}
}
