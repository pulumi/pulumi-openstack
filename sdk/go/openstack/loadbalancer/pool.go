// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 pool resource within OpenStack.
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
// 		_, err := loadbalancer.NewPool(ctx, "pool1", &loadbalancer.PoolArgs{
// 			LbMethod:   pulumi.String("ROUND_ROBIN"),
// 			ListenerId: pulumi.String("d9415786-5f1a-428b-b35f-2f1523e146d2"),
// 			Persistence: &loadbalancer.PoolPersistenceArgs{
// 				CookieName: pulumi.String("testCookie"),
// 				Type:       pulumi.String("APP_COOKIE"),
// 			},
// 			Protocol: pulumi.String("HTTP"),
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
// Load Balancer Pool can be imported using the Pool ID, e.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/pool:Pool pool_1 60ad9ee4-249a-4d60-a45b-aa60e046c513
// ```
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
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
	// UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
	// or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbMethod == nil {
		return nil, errors.New("invalid value for required argument 'LbMethod'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
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
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
	// UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
	// or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
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
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
	// UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
	// or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
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
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
	// UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
	// or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
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
	// The protocol - can either be TCP, HTTP, HTTPS, PROXY,
	// UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
	// or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
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

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

func (i *Pool) ToPoolPtrOutput() PoolPtrOutput {
	return i.ToPoolPtrOutputWithContext(context.Background())
}

func (i *Pool) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPtrOutput)
}

type PoolPtrInput interface {
	pulumi.Input

	ToPoolPtrOutput() PoolPtrOutput
	ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput
}

type poolPtrType PoolArgs

func (*poolPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil))
}

func (i *poolPtrType) ToPoolPtrOutput() PoolPtrOutput {
	return i.ToPoolPtrOutputWithContext(context.Background())
}

func (i *poolPtrType) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPtrOutput)
}

// PoolArrayInput is an input type that accepts PoolArray and PoolArrayOutput values.
// You can construct a concrete instance of `PoolArrayInput` via:
//
//          PoolArray{ PoolArgs{...} }
type PoolArrayInput interface {
	pulumi.Input

	ToPoolArrayOutput() PoolArrayOutput
	ToPoolArrayOutputWithContext(context.Context) PoolArrayOutput
}

type PoolArray []PoolInput

func (PoolArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Pool)(nil))
}

func (i PoolArray) ToPoolArrayOutput() PoolArrayOutput {
	return i.ToPoolArrayOutputWithContext(context.Background())
}

func (i PoolArray) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolArrayOutput)
}

// PoolMapInput is an input type that accepts PoolMap and PoolMapOutput values.
// You can construct a concrete instance of `PoolMapInput` via:
//
//          PoolMap{ "key": PoolArgs{...} }
type PoolMapInput interface {
	pulumi.Input

	ToPoolMapOutput() PoolMapOutput
	ToPoolMapOutputWithContext(context.Context) PoolMapOutput
}

type PoolMap map[string]PoolInput

func (PoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Pool)(nil))
}

func (i PoolMap) ToPoolMapOutput() PoolMapOutput {
	return i.ToPoolMapOutputWithContext(context.Background())
}

func (i PoolMap) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolMapOutput)
}

type PoolOutput struct {
	*pulumi.OutputState
}

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pool)(nil))
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

func (o PoolOutput) ToPoolPtrOutput() PoolPtrOutput {
	return o.ToPoolPtrOutputWithContext(context.Background())
}

func (o PoolOutput) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return o.ApplyT(func(v Pool) *Pool {
		return &v
	}).(PoolPtrOutput)
}

type PoolPtrOutput struct {
	*pulumi.OutputState
}

func (PoolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil))
}

func (o PoolPtrOutput) ToPoolPtrOutput() PoolPtrOutput {
	return o
}

func (o PoolPtrOutput) ToPoolPtrOutputWithContext(ctx context.Context) PoolPtrOutput {
	return o
}

type PoolArrayOutput struct{ *pulumi.OutputState }

func (PoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pool)(nil))
}

func (o PoolArrayOutput) ToPoolArrayOutput() PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) ToPoolArrayOutputWithContext(ctx context.Context) PoolArrayOutput {
	return o
}

func (o PoolArrayOutput) Index(i pulumi.IntInput) PoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pool {
		return vs[0].([]Pool)[vs[1].(int)]
	}).(PoolOutput)
}

type PoolMapOutput struct{ *pulumi.OutputState }

func (PoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pool)(nil))
}

func (o PoolMapOutput) ToPoolMapOutput() PoolMapOutput {
	return o
}

func (o PoolMapOutput) ToPoolMapOutputWithContext(ctx context.Context) PoolMapOutput {
	return o
}

func (o PoolMapOutput) MapIndex(k pulumi.StringInput) PoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pool {
		return vs[0].(map[string]Pool)[vs[1].(string)]
	}).(PoolOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
	pulumi.RegisterOutputType(PoolPtrOutput{})
	pulumi.RegisterOutputType(PoolArrayOutput{})
	pulumi.RegisterOutputType(PoolMapOutput{})
}
