// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 Server Group resource within OpenStack.
//
// ## Policies
//
// * `affinity` - All instances/servers launched in this group will be hosted on
//     the same compute node.
//
// * `anti-affinity` - All instances/servers launched in this group will be
//     hosted on different compute nodes.
//
// * `soft-affinity` - All instances/servers launched in this group will be hosted
//     on the same compute node if possible, but if not possible they
//     still will be scheduled instead of failure. To use this policy your
//     OpenStack environment should support Compute service API 2.15 or above.
//
// * `soft-anti-affinity` - All instances/servers launched in this group will be
//     hosted on different compute nodes if possible, but if not possible they
//     still will be scheduled instead of failure. To use this policy your
//     OpenStack environment should support Compute service API 2.15 or above.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_servergroup_v2.html.markdown.
type ServerGroup struct {
	pulumi.CustomResourceState

	// The instances that are part of this server group.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		args = &ServerGroupArgs{}
	}
	var resource ServerGroup
	err := ctx.RegisterResource("openstack:compute/serverGroup:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("openstack:compute/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// The instances that are part of this server group.
	Members []string `pulumi:"members"`
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name *string `pulumi:"name"`
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies []string `pulumi:"policies"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region *string `pulumi:"region"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type ServerGroupState struct {
	// The instances that are part of this server group.
	Members pulumi.StringArrayInput
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name pulumi.StringPtrInput
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies pulumi.StringArrayInput
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name *string `pulumi:"name"`
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies []string `pulumi:"policies"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region *string `pulumi:"region"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// A unique name for the server group. Changing this creates
	// a new server group.
	Name pulumi.StringPtrInput
	// The set of policies for the server group. All policies
	// are mutually exclusive. See the Policies section for more information.
	// Changing this creates a new server group.
	Policies pulumi.StringArrayInput
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used. Changing
	// this creates a new server group.
	Region pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}
