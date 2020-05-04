// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 pool resource within OpenStack.
type Pool struct {
	pulumi.CustomResourceState

	// The administrative state of the pool.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// Human-readable description for the pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The load balancing algorithm to
	// distribute traffic to the pool's members. Must be one of
	// ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
	// in Octavia).
	LbMethod pulumi.StringOutput `pulumi:"lbMethod"`
	// The Listener on which the members of the pool
	// will be associated with. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	ListenerId pulumi.StringPtrOutput `pulumi:"listenerId"`
	// The load balancer on which to provision this
	// pool. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId pulumi.StringPtrOutput `pulumi:"loadbalancerId"`
	// Human-readable name for the pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Omit this field to prevent session persistence.  Indicates
	// whether connections in the same session will be processed by the same Pool
	// member or not. Changing this creates a new pool.
	Persistence PoolPersistenceOutput `pulumi:"persistence"`
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY
	// or UDP (supported only in Octavia). Changing this creates a new pool.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// pool.
	Region pulumi.StringOutput `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the pool.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new pool.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil || args.LbMethod == nil {
		return nil, errors.New("missing required argument 'LbMethod'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil {
		args = &PoolArgs{}
	}
	var resource Pool
	err := ctx.RegisterResource("openstack:loadbalancer/pool:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("openstack:loadbalancer/pool:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
	// The administrative state of the pool.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description for the pool.
	Description *string `pulumi:"description"`
	// The load balancing algorithm to
	// distribute traffic to the pool's members. Must be one of
	// ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
	// in Octavia).
	LbMethod *string `pulumi:"lbMethod"`
	// The Listener on which the members of the pool
	// will be associated with. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	ListenerId *string `pulumi:"listenerId"`
	// The load balancer on which to provision this
	// pool. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the pool.
	Name *string `pulumi:"name"`
	// Omit this field to prevent session persistence.  Indicates
	// whether connections in the same session will be processed by the same Pool
	// member or not. Changing this creates a new pool.
	Persistence *PoolPersistence `pulumi:"persistence"`
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY
	// or UDP (supported only in Octavia). Changing this creates a new pool.
	Protocol *string `pulumi:"protocol"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// pool.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the pool.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new pool.
	TenantId *string `pulumi:"tenantId"`
}

type PoolState struct {
	// The administrative state of the pool.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable description for the pool.
	Description pulumi.StringPtrInput
	// The load balancing algorithm to
	// distribute traffic to the pool's members. Must be one of
	// ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
	// in Octavia).
	LbMethod pulumi.StringPtrInput
	// The Listener on which the members of the pool
	// will be associated with. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	ListenerId pulumi.StringPtrInput
	// The load balancer on which to provision this
	// pool. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the pool.
	Name pulumi.StringPtrInput
	// Omit this field to prevent session persistence.  Indicates
	// whether connections in the same session will be processed by the same Pool
	// member or not. Changing this creates a new pool.
	Persistence PoolPersistencePtrInput
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY
	// or UDP (supported only in Octavia). Changing this creates a new pool.
	Protocol pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// pool.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the pool.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new pool.
	TenantId pulumi.StringPtrInput
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// The administrative state of the pool.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description for the pool.
	Description *string `pulumi:"description"`
	// The load balancing algorithm to
	// distribute traffic to the pool's members. Must be one of
	// ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
	// in Octavia).
	LbMethod string `pulumi:"lbMethod"`
	// The Listener on which the members of the pool
	// will be associated with. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	ListenerId *string `pulumi:"listenerId"`
	// The load balancer on which to provision this
	// pool. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Human-readable name for the pool.
	Name *string `pulumi:"name"`
	// Omit this field to prevent session persistence.  Indicates
	// whether connections in the same session will be processed by the same Pool
	// member or not. Changing this creates a new pool.
	Persistence *PoolPersistence `pulumi:"persistence"`
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY
	// or UDP (supported only in Octavia). Changing this creates a new pool.
	Protocol string `pulumi:"protocol"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// pool.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the pool.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new pool.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// The administrative state of the pool.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable description for the pool.
	Description pulumi.StringPtrInput
	// The load balancing algorithm to
	// distribute traffic to the pool's members. Must be one of
	// ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
	// in Octavia).
	LbMethod pulumi.StringInput
	// The Listener on which the members of the pool
	// will be associated with. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	ListenerId pulumi.StringPtrInput
	// The load balancer on which to provision this
	// pool. Changing this creates a new pool.
	// Note:  One of LoadbalancerID or ListenerID must be provided.
	LoadbalancerId pulumi.StringPtrInput
	// Human-readable name for the pool.
	Name pulumi.StringPtrInput
	// Omit this field to prevent session persistence.  Indicates
	// whether connections in the same session will be processed by the same Pool
	// member or not. Changing this creates a new pool.
	Persistence PoolPersistencePtrInput
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY
	// or UDP (supported only in Octavia). Changing this creates a new pool.
	Protocol pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// pool.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the pool.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new pool.
	TenantId pulumi.StringPtrInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}
