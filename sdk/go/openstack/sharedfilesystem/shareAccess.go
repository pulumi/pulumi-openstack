// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sharedfilesystem

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/sharedfilesystem_share_access_v2.html.markdown.
type ShareAccess struct {
	pulumi.CustomResourceState

	// The access credential of the entity granted access.
	AccessKey pulumi.StringOutput `pulumi:"accessKey"`
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo pulumi.StringOutput `pulumi:"accessTo"`
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType pulumi.StringOutput `pulumi:"accessType"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region pulumi.StringOutput `pulumi:"region"`
	// The UUID of the share to which you are granted access.
	ShareId pulumi.StringOutput `pulumi:"shareId"`
}

// NewShareAccess registers a new resource with the given unique name, arguments, and options.
func NewShareAccess(ctx *pulumi.Context,
	name string, args *ShareAccessArgs, opts ...pulumi.ResourceOption) (*ShareAccess, error) {
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
	if args == nil {
		args = &ShareAccessArgs{}
	}
	var resource ShareAccess
	err := ctx.RegisterResource("openstack:sharedfilesystem/shareAccess:ShareAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShareAccess gets an existing ShareAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShareAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareAccessState, opts ...pulumi.ResourceOption) (*ShareAccess, error) {
	var resource ShareAccess
	err := ctx.ReadResource("openstack:sharedfilesystem/shareAccess:ShareAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShareAccess resources.
type shareAccessState struct {
	// The access credential of the entity granted access.
	AccessKey *string `pulumi:"accessKey"`
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel *string `pulumi:"accessLevel"`
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo *string `pulumi:"accessTo"`
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType *string `pulumi:"accessType"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region *string `pulumi:"region"`
	// The UUID of the share to which you are granted access.
	ShareId *string `pulumi:"shareId"`
}

type ShareAccessState struct {
	// The access credential of the entity granted access.
	AccessKey pulumi.StringPtrInput
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel pulumi.StringPtrInput
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo pulumi.StringPtrInput
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType pulumi.StringPtrInput
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region pulumi.StringPtrInput
	// The UUID of the share to which you are granted access.
	ShareId pulumi.StringPtrInput
}

func (ShareAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareAccessState)(nil)).Elem()
}

type shareAccessArgs struct {
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel string `pulumi:"accessLevel"`
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo string `pulumi:"accessTo"`
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType string `pulumi:"accessType"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region *string `pulumi:"region"`
	// The UUID of the share to which you are granted access.
	ShareId string `pulumi:"shareId"`
}

// The set of arguments for constructing a ShareAccess resource.
type ShareAccessArgs struct {
	// The access level to the share. Can either be `rw` or `ro`.
	AccessLevel pulumi.StringInput
	// The value that defines the access. Can either be an IP
	// address or a username verified by configured Security Service of the Share Network.
	AccessTo pulumi.StringInput
	// The access rule type. Can either be an ip, user,
	// cert, or cephx. cephx support requires an OpenStack environment that supports
	// Shared Filesystem microversion 2.13 (Mitaka) or later.
	AccessType pulumi.StringInput
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a share access. Changing this
	// creates a new share access.
	Region pulumi.StringPtrInput
	// The UUID of the share to which you are granted access.
	ShareId pulumi.StringInput
}

func (ShareAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareAccessArgs)(nil)).Elem()
}

