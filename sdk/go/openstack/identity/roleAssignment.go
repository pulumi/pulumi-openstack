// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V3 Role assignment within OpenStack Keystone.
// 
// Note: You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
type RoleAssignment struct {
	s *pulumi.ResourceState
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOpt) (*RoleAssignment, error) {
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["domainId"] = nil
		inputs["groupId"] = nil
		inputs["projectId"] = nil
		inputs["region"] = nil
		inputs["roleId"] = nil
		inputs["userId"] = nil
	} else {
		inputs["domainId"] = args.DomainId
		inputs["groupId"] = args.GroupId
		inputs["projectId"] = args.ProjectId
		inputs["region"] = args.Region
		inputs["roleId"] = args.RoleId
		inputs["userId"] = args.UserId
	}
	s, err := ctx.RegisterResource("openstack:identity/roleAssignment:RoleAssignment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RoleAssignment{s: s}, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RoleAssignmentState, opts ...pulumi.ResourceOpt) (*RoleAssignment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["domainId"] = state.DomainId
		inputs["groupId"] = state.GroupId
		inputs["projectId"] = state.ProjectId
		inputs["region"] = state.Region
		inputs["roleId"] = state.RoleId
		inputs["userId"] = state.UserId
	}
	s, err := ctx.ReadResource("openstack:identity/roleAssignment:RoleAssignment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RoleAssignment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RoleAssignment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RoleAssignment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The domain to assign the role in.
func (r *RoleAssignment) DomainId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainId"])
}

// The group to assign the role to.
func (r *RoleAssignment) GroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["groupId"])
}

// The project to assign the role in.
func (r *RoleAssignment) ProjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["projectId"])
}

func (r *RoleAssignment) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The role to assign.
func (r *RoleAssignment) RoleId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleId"])
}

// The user to assign the role to.
func (r *RoleAssignment) UserId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userId"])
}

// Input properties used for looking up and filtering RoleAssignment resources.
type RoleAssignmentState struct {
	// The domain to assign the role in.
	DomainId interface{}
	// The group to assign the role to.
	GroupId interface{}
	// The project to assign the role in.
	ProjectId interface{}
	Region interface{}
	// The role to assign.
	RoleId interface{}
	// The user to assign the role to.
	UserId interface{}
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The domain to assign the role in.
	DomainId interface{}
	// The group to assign the role to.
	GroupId interface{}
	// The project to assign the role in.
	ProjectId interface{}
	Region interface{}
	// The role to assign.
	RoleId interface{}
	// The user to assign the role to.
	UserId interface{}
}
