// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The RBAC policy resource contains functionality for working with Neutron RBAC
// Policies. Role-Based Access Control (RBAC) policy framework enables both
// operators and users to grant access to resources for specific projects.
//
// Sharing an object with a specific project is accomplished by creating a
// policy entry that permits the target project the `accessAsShared` action
// on that object.
//
// To make a network available as an external network for specific projects
// rather than all projects, use the `accessAsExternal` action.
// If a network is marked as external during creation, it now implicitly creates
// a wildcard RBAC policy granting everyone access to preserve previous behavior
// before this feature was added.
type RbacPolicyV2 struct {
	pulumi.CustomResourceState

	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// The type of the object that the RBAC policy
	// affects. Can either be `qos-policy` or `network`.
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	ProjectId  pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant pulumi.StringOutput `pulumi:"targetTenant"`
}

// NewRbacPolicyV2 registers a new resource with the given unique name, arguments, and options.
func NewRbacPolicyV2(ctx *pulumi.Context,
	name string, args *RbacPolicyV2Args, opts ...pulumi.ResourceOption) (*RbacPolicyV2, error) {
	if args == nil || args.Action == nil {
		return nil, errors.New("missing required argument 'Action'")
	}
	if args == nil || args.ObjectId == nil {
		return nil, errors.New("missing required argument 'ObjectId'")
	}
	if args == nil || args.ObjectType == nil {
		return nil, errors.New("missing required argument 'ObjectType'")
	}
	if args == nil || args.TargetTenant == nil {
		return nil, errors.New("missing required argument 'TargetTenant'")
	}
	if args == nil {
		args = &RbacPolicyV2Args{}
	}
	var resource RbacPolicyV2
	err := ctx.RegisterResource("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRbacPolicyV2 gets an existing RbacPolicyV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRbacPolicyV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RbacPolicyV2State, opts ...pulumi.ResourceOption) (*RbacPolicyV2, error) {
	var resource RbacPolicyV2
	err := ctx.ReadResource("openstack:networking/rbacPolicyV2:RbacPolicyV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RbacPolicyV2 resources.
type rbacPolicyV2State struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action *string `pulumi:"action"`
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId *string `pulumi:"objectId"`
	// The type of the object that the RBAC policy
	// affects. Can either be `qos-policy` or `network`.
	ObjectType *string `pulumi:"objectType"`
	ProjectId  *string `pulumi:"projectId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant *string `pulumi:"targetTenant"`
}

type RbacPolicyV2State struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action pulumi.StringPtrInput
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId pulumi.StringPtrInput
	// The type of the object that the RBAC policy
	// affects. Can either be `qos-policy` or `network`.
	ObjectType pulumi.StringPtrInput
	ProjectId  pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant pulumi.StringPtrInput
}

func (RbacPolicyV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacPolicyV2State)(nil)).Elem()
}

type rbacPolicyV2Args struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action string `pulumi:"action"`
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId string `pulumi:"objectId"`
	// The type of the object that the RBAC policy
	// affects. Can either be `qos-policy` or `network`.
	ObjectType string `pulumi:"objectType"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region *string `pulumi:"region"`
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant string `pulumi:"targetTenant"`
}

// The set of arguments for constructing a RbacPolicyV2 resource.
type RbacPolicyV2Args struct {
	// Action for the RBAC policy. Can either be
	// `accessAsExternal` or `accessAsShared`.
	Action pulumi.StringInput
	// The ID of the `objectType` resource. An
	// `objectType` of `network` returns a network ID and an `objectType` of
	// `qosPolicy` returns a QoS ID.
	ObjectId pulumi.StringInput
	// The type of the object that the RBAC policy
	// affects. Can either be `qos-policy` or `network`.
	ObjectType pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to configure a routing entry on a subnet. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// routing entry.
	Region pulumi.StringPtrInput
	// The ID of the tenant to which the RBAC policy
	// will be enforced.
	TargetTenant pulumi.StringInput
}

func (RbacPolicyV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*rbacPolicyV2Args)(nil)).Elem()
}
