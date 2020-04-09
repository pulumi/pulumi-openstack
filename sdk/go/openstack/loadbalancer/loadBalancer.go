// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package loadbalancer

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 loadbalancer resource within OpenStack.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_loadbalancer_v2.html.markdown.
type LoadBalancer struct {
	pulumi.CustomResourceState

	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// Human-readable description for the Loadbalancer.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The UUID of a flavor. Changing this creates a new
	// loadbalancer.
	FlavorId pulumi.StringPtrOutput `pulumi:"flavorId"`
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
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress pulumi.StringOutput `pulumi:"vipAddress"`
	// The Port ID of the Load Balancer IP.
	VipPortId pulumi.StringOutput `pulumi:"vipPortId"`
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	VipSubnetId pulumi.StringOutput `pulumi:"vipSubnetId"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil || args.VipSubnetId == nil {
		return nil, errors.New("missing required argument 'VipSubnetId'")
	}
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
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId *string `pulumi:"tenantId"`
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress *string `pulumi:"vipAddress"`
	// The Port ID of the Load Balancer IP.
	VipPortId *string `pulumi:"vipPortId"`
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	VipSubnetId *string `pulumi:"vipSubnetId"`
}

type LoadBalancerState struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
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
	// Required for admins. The UUID of the tenant who owns
	// the Loadbalancer.  Only administrative users can specify a tenant UUID
	// other than their own.  Changing this creates a new loadbalancer.
	TenantId pulumi.StringPtrInput
	// The ip address of the load balancer.
	// Changing this creates a new loadbalancer.
	VipAddress pulumi.StringPtrInput
	// The Port ID of the Load Balancer IP.
	VipPortId pulumi.StringPtrInput
	// The network on which to allocate the
	// Loadbalancer's address. A tenant can only create Loadbalancers on networks
	// authorized by policy (e.g. networks that belong to them or networks that
	// are shared).  Changing this creates a new loadbalancer.
	VipSubnetId pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
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
	VipSubnetId string `pulumi:"vipSubnetId"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
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
	VipSubnetId pulumi.StringInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}
