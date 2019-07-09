// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ShareAccess struct {
	s *pulumi.ResourceState
}

// NewShareAccess registers a new resource with the given unique name, arguments, and options.
func NewShareAccess(ctx *pulumi.Context,
	name string, args *ShareAccessArgs, opts ...pulumi.ResourceOpt) (*ShareAccess, error) {
	if args == nil || args.AccessLevel == nil {
		return nil, errors.New("missing required argument 'AccessLevel'")
	}
	if args == nil || args.AccessTo == nil {
		return nil, errors.New("missing required argument 'AccessTo'")
	}
	if args == nil || args.AccessType == nil {
		return nil, errors.New("missing required argument 'AccessType'")
	}
	if args == nil || args.ShareId == nil {
		return nil, errors.New("missing required argument 'ShareId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessLevel"] = nil
		inputs["accessTo"] = nil
		inputs["accessType"] = nil
		inputs["region"] = nil
		inputs["shareId"] = nil
	} else {
		inputs["accessLevel"] = args.AccessLevel
		inputs["accessTo"] = args.AccessTo
		inputs["accessType"] = args.AccessType
		inputs["region"] = args.Region
		inputs["shareId"] = args.ShareId
	}
	inputs["accessKey"] = nil
	s, err := ctx.RegisterResource("openstack:sharedfilesystem/shareAccess:ShareAccess", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ShareAccess{s: s}, nil
}

// GetShareAccess gets an existing ShareAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShareAccess(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ShareAccessState, opts ...pulumi.ResourceOpt) (*ShareAccess, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessKey"] = state.AccessKey
		inputs["accessLevel"] = state.AccessLevel
		inputs["accessTo"] = state.AccessTo
		inputs["accessType"] = state.AccessType
		inputs["region"] = state.Region
		inputs["shareId"] = state.ShareId
	}
	s, err := ctx.ReadResource("openstack:sharedfilesystem/shareAccess:ShareAccess", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ShareAccess{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ShareAccess) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ShareAccess) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The access credential of the entity granted access.
func (r *ShareAccess) AccessKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessKey"])
}

// The access level to the share. Can either be `rw` or `ro`.
func (r *ShareAccess) AccessLevel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessLevel"])
}

// The value that defines the access. Can either be an IP
// address or a username verified by configured Security Service of the Share Network.
func (r *ShareAccess) AccessTo() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessTo"])
}

// The access rule type. Can either be an ip, user,
// cert, or cephx. cephx support requires an OpenStack environment that supports
// Shared Filesystem microversion 2.13 (Mitaka) or later.
func (r *ShareAccess) AccessType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessType"])
}

// The region in which to obtain the V2 Shared File System client.
// A Shared File System client is needed to create a share access. Changing this
// creates a new share access.
func (r *ShareAccess) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The UUID of the share to which you are granted access.
func (r *ShareAccess) ShareId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["shareId"])
}

// Input properties used for looking up and filtering ShareAccess resources.
type ShareAccessState struct {
	// The access credential of the entity granted access.
	AccessKey interface{}
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel interface{}
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo interface{}
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType interface{}
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region interface{}
	// The UUID of the share to which you are granted access.
	ShareId interface{}
}

// The set of arguments for constructing a ShareAccess resource.
type ShareAccessArgs struct {
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel interface{}
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo interface{}
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType interface{}
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region interface{}
	// The UUID of the share to which you are granted access.
	ShareId interface{}
}
