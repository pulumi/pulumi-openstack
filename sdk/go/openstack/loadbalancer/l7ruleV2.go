// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 L7 Rule resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/loadbalancer"
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network1, err := networking.NewNetwork(ctx, "network1", &networking.NetworkArgs{
// 			AdminStateUp: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		subnet1, err := networking.NewSubnet(ctx, "subnet1", &networking.SubnetArgs{
// 			Cidr:      pulumi.String("192.168.199.0/24"),
// 			IpVersion: pulumi.Int(4),
// 			NetworkId: network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		loadbalancer1, err := loadbalancer.NewLoadBalancer(ctx, "loadbalancer1", &loadbalancer.LoadBalancerArgs{
// 			VipSubnetId: subnet1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		listener1, err := loadbalancer.NewListener(ctx, "listener1", &loadbalancer.ListenerArgs{
// 			LoadbalancerId: loadbalancer1.ID(),
// 			Protocol:       pulumi.String("HTTP"),
// 			ProtocolPort:   pulumi.Int(8080),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loadbalancer.NewPool(ctx, "pool1", &loadbalancer.PoolArgs{
// 			LbMethod:       pulumi.String("ROUND_ROBIN"),
// 			LoadbalancerId: loadbalancer1.ID(),
// 			Protocol:       pulumi.String("HTTP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		l7policy1, err := loadbalancer.NewL7PolicyV2(ctx, "l7policy1", &loadbalancer.L7PolicyV2Args{
// 			Action:      pulumi.String("REDIRECT_TO_URL"),
// 			Description: pulumi.String("test description"),
// 			ListenerId:  listener1.ID(),
// 			Position:    pulumi.Int(1),
// 			RedirectUrl: pulumi.String("http://www.example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loadbalancer.NewL7RuleV2(ctx, "l7rule1", &loadbalancer.L7RuleV2Args{
// 			CompareType: pulumi.String("EQUAL_TO"),
// 			L7policyId:  l7policy1.ID(),
// 			Type:        pulumi.String("PATH"),
// 			Value:       pulumi.String("/api"),
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
// Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/l7RuleV2:L7RuleV2 l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
// ```
type L7RuleV2 struct {
	pulumi.CustomResourceState

	// The administrative state of the L7 Rule.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// The comparison type for the L7 rule - can either be
	// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
	CompareType pulumi.StringOutput `pulumi:"compareType"`
	// When true the logic of the rule is inverted. For example, with invert
	// true, equal to would become not equal to. Default is false.
	Invert pulumi.BoolPtrOutput `pulumi:"invert"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate. Valid when `type` is set to COOKIE or HEADER.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// The ID of the L7 Policy to query. Changing this creates a new
	// L7 Rule.
	L7policyId pulumi.StringOutput `pulumi:"l7policyId"`
	// The ID of the Listener owning this resource.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Rule.
	Region pulumi.StringOutput `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the L7 Rule.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Rule.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
	// HOST\_NAME or PATH.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value to use for the comparison. For example, the file type to
	// compare.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewL7RuleV2 registers a new resource with the given unique name, arguments, and options.
func NewL7RuleV2(ctx *pulumi.Context,
	name string, args *L7RuleV2Args, opts ...pulumi.ResourceOption) (*L7RuleV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompareType == nil {
		return nil, errors.New("invalid value for required argument 'CompareType'")
	}
	if args.L7policyId == nil {
		return nil, errors.New("invalid value for required argument 'L7policyId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource L7RuleV2
	err := ctx.RegisterResource("openstack:loadbalancer/l7RuleV2:L7RuleV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL7RuleV2 gets an existing L7RuleV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL7RuleV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L7RuleV2State, opts ...pulumi.ResourceOption) (*L7RuleV2, error) {
	var resource L7RuleV2
	err := ctx.ReadResource("openstack:loadbalancer/l7RuleV2:L7RuleV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L7RuleV2 resources.
type l7ruleV2State struct {
	// The administrative state of the L7 Rule.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The comparison type for the L7 rule - can either be
	// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
	CompareType *string `pulumi:"compareType"`
	// When true the logic of the rule is inverted. For example, with invert
	// true, equal to would become not equal to. Default is false.
	Invert *bool `pulumi:"invert"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate. Valid when `type` is set to COOKIE or HEADER.
	Key *string `pulumi:"key"`
	// The ID of the L7 Policy to query. Changing this creates a new
	// L7 Rule.
	L7policyId *string `pulumi:"l7policyId"`
	// The ID of the Listener owning this resource.
	ListenerId *string `pulumi:"listenerId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Rule.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the L7 Rule.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Rule.
	TenantId *string `pulumi:"tenantId"`
	// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
	// HOST\_NAME or PATH.
	Type *string `pulumi:"type"`
	// The value to use for the comparison. For example, the file type to
	// compare.
	Value *string `pulumi:"value"`
}

type L7RuleV2State struct {
	// The administrative state of the L7 Rule.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// The comparison type for the L7 rule - can either be
	// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
	CompareType pulumi.StringPtrInput
	// When true the logic of the rule is inverted. For example, with invert
	// true, equal to would become not equal to. Default is false.
	Invert pulumi.BoolPtrInput
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate. Valid when `type` is set to COOKIE or HEADER.
	Key pulumi.StringPtrInput
	// The ID of the L7 Policy to query. Changing this creates a new
	// L7 Rule.
	L7policyId pulumi.StringPtrInput
	// The ID of the Listener owning this resource.
	ListenerId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Rule.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the L7 Rule.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Rule.
	TenantId pulumi.StringPtrInput
	// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
	// HOST\_NAME or PATH.
	Type pulumi.StringPtrInput
	// The value to use for the comparison. For example, the file type to
	// compare.
	Value pulumi.StringPtrInput
}

func (L7RuleV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleV2State)(nil)).Elem()
}

type l7ruleV2Args struct {
	// The administrative state of the L7 Rule.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// The comparison type for the L7 rule - can either be
	// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
	CompareType string `pulumi:"compareType"`
	// When true the logic of the rule is inverted. For example, with invert
	// true, equal to would become not equal to. Default is false.
	Invert *bool `pulumi:"invert"`
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate. Valid when `type` is set to COOKIE or HEADER.
	Key *string `pulumi:"key"`
	// The ID of the L7 Policy to query. Changing this creates a new
	// L7 Rule.
	L7policyId string `pulumi:"l7policyId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Rule.
	Region *string `pulumi:"region"`
	// Required for admins. The UUID of the tenant who owns
	// the L7 Rule.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Rule.
	TenantId *string `pulumi:"tenantId"`
	// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
	// HOST\_NAME or PATH.
	Type string `pulumi:"type"`
	// The value to use for the comparison. For example, the file type to
	// compare.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a L7RuleV2 resource.
type L7RuleV2Args struct {
	// The administrative state of the L7 Rule.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp pulumi.BoolPtrInput
	// The comparison type for the L7 rule - can either be
	// CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
	CompareType pulumi.StringInput
	// When true the logic of the rule is inverted. For example, with invert
	// true, equal to would become not equal to. Default is false.
	Invert pulumi.BoolPtrInput
	// The key to use for the comparison. For example, the name of the cookie to
	// evaluate. Valid when `type` is set to COOKIE or HEADER.
	Key pulumi.StringPtrInput
	// The ID of the L7 Policy to query. Changing this creates a new
	// L7 Rule.
	L7policyId pulumi.StringInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an . If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// L7 Rule.
	Region pulumi.StringPtrInput
	// Required for admins. The UUID of the tenant who owns
	// the L7 Rule.  Only administrative users can specify a tenant UUID
	// other than their own. Changing this creates a new L7 Rule.
	TenantId pulumi.StringPtrInput
	// The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
	// HOST\_NAME or PATH.
	Type pulumi.StringInput
	// The value to use for the comparison. For example, the file type to
	// compare.
	Value pulumi.StringInput
}

func (L7RuleV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleV2Args)(nil)).Elem()
}

type L7RuleV2Input interface {
	pulumi.Input

	ToL7RuleV2Output() L7RuleV2Output
	ToL7RuleV2OutputWithContext(ctx context.Context) L7RuleV2Output
}

func (L7RuleV2) ElementType() reflect.Type {
	return reflect.TypeOf((*L7RuleV2)(nil)).Elem()
}

func (i L7RuleV2) ToL7RuleV2Output() L7RuleV2Output {
	return i.ToL7RuleV2OutputWithContext(context.Background())
}

func (i L7RuleV2) ToL7RuleV2OutputWithContext(ctx context.Context) L7RuleV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleV2Output)
}

type L7RuleV2Output struct {
	*pulumi.OutputState
}

func (L7RuleV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*L7RuleV2Output)(nil)).Elem()
}

func (o L7RuleV2Output) ToL7RuleV2Output() L7RuleV2Output {
	return o
}

func (o L7RuleV2Output) ToL7RuleV2OutputWithContext(ctx context.Context) L7RuleV2Output {
	return o
}

func init() {
	pulumi.RegisterOutputType(L7RuleV2Output{})
}
