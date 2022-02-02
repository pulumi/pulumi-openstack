// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 loadbalancer resource within OpenStack.
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
// 		_, err := loadbalancer.NewLoadBalancer(ctx, "lb1", &loadbalancer.LoadBalancerArgs{
// 			VipSubnetId: pulumi.String("d9415786-5f1a-428b-b35f-2f1523e146d2"),
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
// Load Balancer can be imported using the Load Balancer ID, e.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/loadBalancer:LoadBalancer loadbalancer_1 19bcfdc7-c521-4a7e-9459-6750bd16df76
// ```
type LoadBalancer struct {
	pulumi.CustomResourceState

	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The availability zone of the Loadbalancer.
	// Changing this creates a new loadbalancer. Available only for Octavia
	// **minor version 2.14 or later**.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Human-readable description for the Loadbalancer.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The UUID of a flavor. Changing this creates a new
	// loadbalancer.
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// The name of the provider. Changing this
	// creates a new loadbalancer.
	LoadbalancerProvider pulumi.StringOutput `pulumi:"loadbalancerProvider"`
	// Human-readable name for the Loadbalancer. Does not have
	// to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of security group IDs to apply to the
	// loadbalancer. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A list of simple strings assigned to the loadbalancer.
	// Available only for Octavia **minor version 2.5 or later**.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress pulumi.StringOutput `pulumi:"vipAddress"`
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is available only for Octavia.
	VipNetworkId pulumi.StringOutput `pulumi:"vipNetworkId"`
	// The port UUID that the loadbalancer will use.
	// Changing this creates a new loadbalancer. It is available only for Octavia.
	VipPortId pulumi.StringOutput `pulumi:"vipPortId"`
	// The subnet on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is required to Neutron LBaaS but optional for Octavia.
	VipSubnetId pulumi.StringOutput `pulumi:"vipSubnetId"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		args = &LoadBalancerArgs{}
	}

	var resource LoadBalancer
	err := ctx.RegisterResource("openstack:loadbalancer/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("openstack:loadbalancer/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The availability zone of the Loadbalancer.
	// Changing this creates a new loadbalancer. Available only for Octavia
	// **minor version 2.14 or later**.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Human-readable description for the Loadbalancer.
	Description *string `pulumi:"description"`
	// The UUID of a flavor. Changing this creates a new
	// loadbalancer.
	FlavorId *string `pulumi:"flavorId"`
	// The name of the provider. Changing this
	// creates a new loadbalancer.
	LoadbalancerProvider *string `pulumi:"loadbalancerProvider"`
	// Human-readable name for the Loadbalancer. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region *string `pulumi:"region"`
	// A list of security group IDs to apply to the
	// loadbalancer. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A list of simple strings assigned to the loadbalancer.
	// Available only for Octavia **minor version 2.5 or later**.
	Tags []string `pulumi:"tags"`
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId *string `pulumi:"tenantId"`
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress *string `pulumi:"vipAddress"`
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is available only for Octavia.
	VipNetworkId *string `pulumi:"vipNetworkId"`
	// The port UUID that the loadbalancer will use.
	// Changing this creates a new loadbalancer. It is available only for Octavia.
	VipPortId *string `pulumi:"vipPortId"`
	// The subnet on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is required to Neutron LBaaS but optional for Octavia.
	VipSubnetId *string `pulumi:"vipSubnetId"`
}

type LoadBalancerState struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// The availability zone of the Loadbalancer.
	// Changing this creates a new loadbalancer. Available only for Octavia
	// **minor version 2.14 or later**.
	AvailabilityZone pulumi.StringPtrInput
	// Human-readable description for the Loadbalancer.
	Description pulumi.StringPtrInput
	// The UUID of a flavor. Changing this creates a new
	// loadbalancer.
	FlavorId pulumi.StringPtrInput
	// The name of the provider. Changing this
	// creates a new loadbalancer.
	LoadbalancerProvider pulumi.StringPtrInput
	// Human-readable name for the Loadbalancer. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region pulumi.StringPtrInput
	// A list of security group IDs to apply to the
	// loadbalancer. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
	// A list of simple strings assigned to the loadbalancer.
	// Available only for Octavia **minor version 2.5 or later**.
	Tags pulumi.StringArrayInput
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId pulumi.StringPtrInput
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress pulumi.StringPtrInput
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is available only for Octavia.
	VipNetworkId pulumi.StringPtrInput
	// The port UUID that the loadbalancer will use.
	// Changing this creates a new loadbalancer. It is available only for Octavia.
	VipPortId pulumi.StringPtrInput
	// The subnet on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is required to Neutron LBaaS but optional for Octavia.
	VipSubnetId pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The availability zone of the Loadbalancer.
	// Changing this creates a new loadbalancer. Available only for Octavia
	// **minor version 2.14 or later**.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Human-readable description for the Loadbalancer.
	Description *string `pulumi:"description"`
	// The UUID of a flavor. Changing this creates a new
	// loadbalancer.
	FlavorId *string `pulumi:"flavorId"`
	// The name of the provider. Changing this
	// creates a new loadbalancer.
	LoadbalancerProvider *string `pulumi:"loadbalancerProvider"`
	// Human-readable name for the Loadbalancer. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region *string `pulumi:"region"`
	// A list of security group IDs to apply to the
	// loadbalancer. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A list of simple strings assigned to the loadbalancer.
	// Available only for Octavia **minor version 2.5 or later**.
	Tags []string `pulumi:"tags"`
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId *string `pulumi:"tenantId"`
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress *string `pulumi:"vipAddress"`
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is available only for Octavia.
	VipNetworkId *string `pulumi:"vipNetworkId"`
	// The port UUID that the loadbalancer will use.
	// Changing this creates a new loadbalancer. It is available only for Octavia.
	VipPortId *string `pulumi:"vipPortId"`
	// The subnet on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is required to Neutron LBaaS but optional for Octavia.
	VipSubnetId *string `pulumi:"vipSubnetId"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// The availability zone of the Loadbalancer.
	// Changing this creates a new loadbalancer. Available only for Octavia
	// **minor version 2.14 or later**.
	AvailabilityZone pulumi.StringPtrInput
	// Human-readable description for the Loadbalancer.
	Description pulumi.StringPtrInput
	// The UUID of a flavor. Changing this creates a new
	// loadbalancer.
	FlavorId pulumi.StringPtrInput
	// The name of the provider. Changing this
	// creates a new loadbalancer.
	LoadbalancerProvider pulumi.StringPtrInput
	// Human-readable name for the Loadbalancer. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an LB member. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// LB member.
	Region pulumi.StringPtrInput
	// A list of security group IDs to apply to the
	// loadbalancer. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
	// A list of simple strings assigned to the loadbalancer.
	// Available only for Octavia **minor version 2.5 or later**.
	Tags pulumi.StringArrayInput
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId pulumi.StringPtrInput
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress pulumi.StringPtrInput
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is available only for Octavia.
	VipNetworkId pulumi.StringPtrInput
	// The port UUID that the loadbalancer will use.
	// Changing this creates a new loadbalancer. It is available only for Octavia.
	VipPortId pulumi.StringPtrInput
	// The subnet on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	// It is required to Neutron LBaaS but optional for Octavia.
	VipSubnetId pulumi.StringPtrInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

// LoadBalancerArrayInput is an input type that accepts LoadBalancerArray and LoadBalancerArrayOutput values.
// You can construct a concrete instance of `LoadBalancerArrayInput` via:
//
//          LoadBalancerArray{ LoadBalancerArgs{...} }
type LoadBalancerArrayInput interface {
	pulumi.Input

	ToLoadBalancerArrayOutput() LoadBalancerArrayOutput
	ToLoadBalancerArrayOutputWithContext(context.Context) LoadBalancerArrayOutput
}

type LoadBalancerArray []LoadBalancerInput

func (LoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return i.ToLoadBalancerArrayOutputWithContext(context.Background())
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerArrayOutput)
}

// LoadBalancerMapInput is an input type that accepts LoadBalancerMap and LoadBalancerMapOutput values.
// You can construct a concrete instance of `LoadBalancerMapInput` via:
//
//          LoadBalancerMap{ "key": LoadBalancerArgs{...} }
type LoadBalancerMapInput interface {
	pulumi.Input

	ToLoadBalancerMapOutput() LoadBalancerMapOutput
	ToLoadBalancerMapOutputWithContext(context.Context) LoadBalancerMapOutput
}

type LoadBalancerMap map[string]LoadBalancerInput

func (LoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerMap) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return i.ToLoadBalancerMapOutputWithContext(context.Background())
}

func (i LoadBalancerMap) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerMapOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

type LoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) Index(i pulumi.IntInput) LoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].([]*LoadBalancer)[vs[1].(int)]
	}).(LoadBalancerOutput)
}

type LoadBalancerMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].(map[string]*LoadBalancer)[vs[1].(string)]
	}).(LoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInput)(nil)).Elem(), &LoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerArrayInput)(nil)).Elem(), LoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerMapInput)(nil)).Elem(), LoadBalancerMap{})
	pulumi.RegisterOutputType(LoadBalancerOutput{})
	pulumi.RegisterOutputType(LoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerMapOutput{})
}
