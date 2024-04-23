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

// Manages a Load Balancer L7 Policy resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/loadbalancer"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			network1, err := networking.NewNetwork(ctx, "network_1", &networking.NetworkArgs{
//				Name:         pulumi.String("network_1"),
//				AdminStateUp: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			subnet1, err := networking.NewSubnet(ctx, "subnet_1", &networking.SubnetArgs{
//				Name:      pulumi.String("subnet_1"),
//				Cidr:      pulumi.String("192.168.199.0/24"),
//				IpVersion: pulumi.Int(4),
//				NetworkId: network1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			loadbalancer1, err := loadbalancer.NewLoadBalancer(ctx, "loadbalancer_1", &loadbalancer.LoadBalancerArgs{
//				Name:        pulumi.String("loadbalancer_1"),
//				VipSubnetId: subnet1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			listener1, err := loadbalancer.NewListener(ctx, "listener_1", &loadbalancer.ListenerArgs{
//				Name:           pulumi.String("listener_1"),
//				Protocol:       pulumi.String("HTTP"),
//				ProtocolPort:   pulumi.Int(8080),
//				LoadbalancerId: loadbalancer1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			pool1, err := loadbalancer.NewPool(ctx, "pool_1", &loadbalancer.PoolArgs{
//				Name:           pulumi.String("pool_1"),
//				Protocol:       pulumi.String("HTTP"),
//				LbMethod:       pulumi.String("ROUND_ROBIN"),
//				LoadbalancerId: loadbalancer1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = loadbalancer.NewL7PolicyV2(ctx, "l7policy_1", &loadbalancer.L7PolicyV2Args{
//				Name:           pulumi.String("test"),
//				Action:         pulumi.String("REDIRECT_TO_POOL"),
//				Description:    pulumi.String("test l7 policy"),
//				Position:       pulumi.Int(1),
//				ListenerId:     listener1.ID(),
//				RedirectPoolId: pool1.ID(),
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
// Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.:
//
// ```sh
// $ pulumi import openstack:loadbalancer/l7PolicyV2:L7PolicyV2 l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
// ```
type L7PolicyV2 struct {
	pulumi.CustomResourceState

	// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
	// REDIRECT\_TO\_URL or REJECT.
	Action pulumi.StringOutput `pulumi:"action"`
	// The administrative state of the L7 Policy.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// Human-readable description for the L7 Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Listener on which the L7 Policy will be associated with.
	// Changing this creates a new L7 Policy.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// Human-readable name for the L7 Policy. Does not have
	// to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The position of this policy on the listener. Positions start at 1.
	Position pulumi.IntOutput `pulumi:"position"`
	// Requests matching this policy will be redirected to the
	// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
	RedirectPoolId pulumi.StringPtrOutput `pulumi:"redirectPoolId"`
	// Requests matching this policy will be redirected to this URL.
	// Only valid if action is REDIRECT\_TO\_URL.
	RedirectUrl pulumi.StringPtrOutput `pulumi:"redirectUrl"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Policy.
	Region pulumi.StringOutput `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the L7 Policy.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Policy.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewL7PolicyV2 registers a new resource with the given unique name, arguments, and options.
func NewL7PolicyV2(ctx *pulumi.Context,
	name string, args *L7PolicyV2Args, opts ...pulumi.ResourceOption) (*L7PolicyV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource L7PolicyV2
	err := ctx.RegisterResource("openstack:loadbalancer/l7PolicyV2:L7PolicyV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL7PolicyV2 gets an existing L7PolicyV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL7PolicyV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L7PolicyV2State, opts ...pulumi.ResourceOption) (*L7PolicyV2, error) {
	var resource L7PolicyV2
	err := ctx.ReadResource("openstack:loadbalancer/l7PolicyV2:L7PolicyV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L7PolicyV2 resources.
type l7policyV2State struct {
	// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
	// REDIRECT\_TO\_URL or REJECT.
	Action *string `pulumi:"action"`
	// The administrative state of the L7 Policy.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description for the L7 Policy.
	Description *string `pulumi:"description"`
	// The Listener on which the L7 Policy will be associated with.
	// Changing this creates a new L7 Policy.
	ListenerId *string `pulumi:"listenerId"`
	// Human-readable name for the L7 Policy. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The position of this policy on the listener. Positions start at 1.
	Position *int `pulumi:"position"`
	// Requests matching this policy will be redirected to the
	// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
	RedirectPoolId *string `pulumi:"redirectPoolId"`
	// Requests matching this policy will be redirected to this URL.
	// Only valid if action is REDIRECT\_TO\_URL.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Policy.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the L7 Policy.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Policy.
	TenantId *string `pulumi:"tenantId"`
}

type L7PolicyV2State struct {
	// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
	// REDIRECT\_TO\_URL or REJECT.
	Action pulumi.StringPtrInput
	// The administrative state of the L7 Policy.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable description for the L7 Policy.
	Description pulumi.StringPtrInput
	// The Listener on which the L7 Policy will be associated with.
	// Changing this creates a new L7 Policy.
	ListenerId pulumi.StringPtrInput
	// Human-readable name for the L7 Policy. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The position of this policy on the listener. Positions start at 1.
	Position pulumi.IntPtrInput
	// Requests matching this policy will be redirected to the
	// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
	RedirectPoolId pulumi.StringPtrInput
	// Requests matching this policy will be redirected to this URL.
	// Only valid if action is REDIRECT\_TO\_URL.
	RedirectUrl pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Policy.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the L7 Policy.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Policy.
	TenantId pulumi.StringPtrInput
}

func (L7PolicyV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*l7policyV2State)(nil)).Elem()
}

type l7policyV2Args struct {
	// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
	// REDIRECT\_TO\_URL or REJECT.
	Action string `pulumi:"action"`
	// The administrative state of the L7 Policy.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description for the L7 Policy.
	Description *string `pulumi:"description"`
	// The Listener on which the L7 Policy will be associated with.
	// Changing this creates a new L7 Policy.
	ListenerId string `pulumi:"listenerId"`
	// Human-readable name for the L7 Policy. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The position of this policy on the listener. Positions start at 1.
	Position *int `pulumi:"position"`
	// Requests matching this policy will be redirected to the
	// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
	RedirectPoolId *string `pulumi:"redirectPoolId"`
	// Requests matching this policy will be redirected to this URL.
	// Only valid if action is REDIRECT\_TO\_URL.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Policy.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the L7 Policy.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Policy.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a L7PolicyV2 resource.
type L7PolicyV2Args struct {
	// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
	// REDIRECT\_TO\_URL or REJECT.
	Action pulumi.StringInput
	// The administrative state of the L7 Policy.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// Human-readable description for the L7 Policy.
	Description pulumi.StringPtrInput
	// The Listener on which the L7 Policy will be associated with.
	// Changing this creates a new L7 Policy.
	ListenerId pulumi.StringInput
	// Human-readable name for the L7 Policy. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The position of this policy on the listener. Positions start at 1.
	Position pulumi.IntPtrInput
	// Requests matching this policy will be redirected to the
	// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
	RedirectPoolId pulumi.StringPtrInput
	// Requests matching this policy will be redirected to this URL.
	// Only valid if action is REDIRECT\_TO\_URL.
	RedirectUrl pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Policy.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the L7 Policy.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Policy.
	TenantId pulumi.StringPtrInput
}

func (L7PolicyV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*l7policyV2Args)(nil)).Elem()
}

type L7PolicyV2Input interface {
	pulumi.Input

	ToL7PolicyV2Output() L7PolicyV2Output
	ToL7PolicyV2OutputWithContext(ctx context.Context) L7PolicyV2Output
}

func (*L7PolicyV2) ElementType() reflect.Type {
	return reflect.TypeOf((**L7PolicyV2)(nil)).Elem()
}

func (i *L7PolicyV2) ToL7PolicyV2Output() L7PolicyV2Output {
	return i.ToL7PolicyV2OutputWithContext(context.Background())
}

func (i *L7PolicyV2) ToL7PolicyV2OutputWithContext(ctx context.Context) L7PolicyV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(L7PolicyV2Output)
}

// L7PolicyV2ArrayInput is an input type that accepts L7PolicyV2Array and L7PolicyV2ArrayOutput values.
// You can construct a concrete instance of `L7PolicyV2ArrayInput` via:
//
//	L7PolicyV2Array{ L7PolicyV2Args{...} }
type L7PolicyV2ArrayInput interface {
	pulumi.Input

	ToL7PolicyV2ArrayOutput() L7PolicyV2ArrayOutput
	ToL7PolicyV2ArrayOutputWithContext(context.Context) L7PolicyV2ArrayOutput
}

type L7PolicyV2Array []L7PolicyV2Input

func (L7PolicyV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7PolicyV2)(nil)).Elem()
}

func (i L7PolicyV2Array) ToL7PolicyV2ArrayOutput() L7PolicyV2ArrayOutput {
	return i.ToL7PolicyV2ArrayOutputWithContext(context.Background())
}

func (i L7PolicyV2Array) ToL7PolicyV2ArrayOutputWithContext(ctx context.Context) L7PolicyV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7PolicyV2ArrayOutput)
}

// L7PolicyV2MapInput is an input type that accepts L7PolicyV2Map and L7PolicyV2MapOutput values.
// You can construct a concrete instance of `L7PolicyV2MapInput` via:
//
//	L7PolicyV2Map{ "key": L7PolicyV2Args{...} }
type L7PolicyV2MapInput interface {
	pulumi.Input

	ToL7PolicyV2MapOutput() L7PolicyV2MapOutput
	ToL7PolicyV2MapOutputWithContext(context.Context) L7PolicyV2MapOutput
}

type L7PolicyV2Map map[string]L7PolicyV2Input

func (L7PolicyV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7PolicyV2)(nil)).Elem()
}

func (i L7PolicyV2Map) ToL7PolicyV2MapOutput() L7PolicyV2MapOutput {
	return i.ToL7PolicyV2MapOutputWithContext(context.Background())
}

func (i L7PolicyV2Map) ToL7PolicyV2MapOutputWithContext(ctx context.Context) L7PolicyV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7PolicyV2MapOutput)
}

type L7PolicyV2Output struct{ *pulumi.OutputState }

func (L7PolicyV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**L7PolicyV2)(nil)).Elem()
}

func (o L7PolicyV2Output) ToL7PolicyV2Output() L7PolicyV2Output {
	return o
}

func (o L7PolicyV2Output) ToL7PolicyV2OutputWithContext(ctx context.Context) L7PolicyV2Output {
	return o
}

// The L7 Policy action - can either be REDIRECT\_TO\_POOL,
// REDIRECT\_TO\_URL or REJECT.
func (o L7PolicyV2Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The administrative state of the L7 Policy.
// A valid value is true (UP) or false (DOWN).
func (o L7PolicyV2Output) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// Human-readable description for the L7 Policy.
func (o L7PolicyV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Listener on which the L7 Policy will be associated with.
// Changing this creates a new L7 Policy.
func (o L7PolicyV2Output) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// Human-readable name for the L7 Policy. Does not have
// to be unique.
func (o L7PolicyV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The position of this policy on the listener. Positions start at 1.
func (o L7PolicyV2Output) Position() pulumi.IntOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.IntOutput { return v.Position }).(pulumi.IntOutput)
}

// Requests matching this policy will be redirected to the
// pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
func (o L7PolicyV2Output) RedirectPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringPtrOutput { return v.RedirectPoolId }).(pulumi.StringPtrOutput)
}

// Requests matching this policy will be redirected to this URL.
// Only valid if action is REDIRECT\_TO\_URL.
func (o L7PolicyV2Output) RedirectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringPtrOutput { return v.RedirectUrl }).(pulumi.StringPtrOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an . If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// L7 Policy.
func (o L7PolicyV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Required for admins. The UUID of the tenant who owns
// the L7 Policy.  Only administrative users can specify a tenant UUID
// other than their own. Changing this creates a new L7 Policy.
func (o L7PolicyV2Output) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *L7PolicyV2) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

type L7PolicyV2ArrayOutput struct{ *pulumi.OutputState }

func (L7PolicyV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7PolicyV2)(nil)).Elem()
}

func (o L7PolicyV2ArrayOutput) ToL7PolicyV2ArrayOutput() L7PolicyV2ArrayOutput {
	return o
}

func (o L7PolicyV2ArrayOutput) ToL7PolicyV2ArrayOutputWithContext(ctx context.Context) L7PolicyV2ArrayOutput {
	return o
}

func (o L7PolicyV2ArrayOutput) Index(i pulumi.IntInput) L7PolicyV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *L7PolicyV2 {
		return vs[0].([]*L7PolicyV2)[vs[1].(int)]
	}).(L7PolicyV2Output)
}

type L7PolicyV2MapOutput struct{ *pulumi.OutputState }

func (L7PolicyV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7PolicyV2)(nil)).Elem()
}

func (o L7PolicyV2MapOutput) ToL7PolicyV2MapOutput() L7PolicyV2MapOutput {
	return o
}

func (o L7PolicyV2MapOutput) ToL7PolicyV2MapOutputWithContext(ctx context.Context) L7PolicyV2MapOutput {
	return o
}

func (o L7PolicyV2MapOutput) MapIndex(k pulumi.StringInput) L7PolicyV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *L7PolicyV2 {
		return vs[0].(map[string]*L7PolicyV2)[vs[1].(string)]
	}).(L7PolicyV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*L7PolicyV2Input)(nil)).Elem(), &L7PolicyV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7PolicyV2ArrayInput)(nil)).Elem(), L7PolicyV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7PolicyV2MapInput)(nil)).Elem(), L7PolicyV2Map{})
	pulumi.RegisterOutputType(L7PolicyV2Output{})
	pulumi.RegisterOutputType(L7PolicyV2ArrayOutput{})
	pulumi.RegisterOutputType(L7PolicyV2MapOutput{})
}
