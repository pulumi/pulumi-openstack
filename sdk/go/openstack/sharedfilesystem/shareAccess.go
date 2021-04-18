// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// This resource can be imported by specifying the ID of the share and the ID of the share access, separated by a slash, e.g.
//
// ```sh
//  $ pulumi import openstack:sharedfilesystem/shareAccess:ShareAccess share_access_1 <share id>/<share access id>
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevel == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevel'")
	}
	if args.AccessTo == nil {
		return nil, errors.New("invalid value for required argument 'AccessTo'")
	}
	if args.AccessType == nil {
		return nil, errors.New("invalid value for required argument 'AccessType'")
	}
	if args.ShareId == nil {
		return nil, errors.New("invalid value for required argument 'ShareId'")
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

type ShareAccessInput interface {
	pulumi.Input

	ToShareAccessOutput() ShareAccessOutput
	ToShareAccessOutputWithContext(ctx context.Context) ShareAccessOutput
}

func (*ShareAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccess)(nil))
}

func (i *ShareAccess) ToShareAccessOutput() ShareAccessOutput {
	return i.ToShareAccessOutputWithContext(context.Background())
}

func (i *ShareAccess) ToShareAccessOutputWithContext(ctx context.Context) ShareAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessOutput)
}

func (i *ShareAccess) ToShareAccessPtrOutput() ShareAccessPtrOutput {
	return i.ToShareAccessPtrOutputWithContext(context.Background())
}

func (i *ShareAccess) ToShareAccessPtrOutputWithContext(ctx context.Context) ShareAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessPtrOutput)
}

type ShareAccessPtrInput interface {
	pulumi.Input

	ToShareAccessPtrOutput() ShareAccessPtrOutput
	ToShareAccessPtrOutputWithContext(ctx context.Context) ShareAccessPtrOutput
}

type shareAccessPtrType ShareAccessArgs

func (*shareAccessPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareAccess)(nil))
}

func (i *shareAccessPtrType) ToShareAccessPtrOutput() ShareAccessPtrOutput {
	return i.ToShareAccessPtrOutputWithContext(context.Background())
}

func (i *shareAccessPtrType) ToShareAccessPtrOutputWithContext(ctx context.Context) ShareAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessPtrOutput)
}

// ShareAccessArrayInput is an input type that accepts ShareAccessArray and ShareAccessArrayOutput values.
// You can construct a concrete instance of `ShareAccessArrayInput` via:
//
//          ShareAccessArray{ ShareAccessArgs{...} }
type ShareAccessArrayInput interface {
	pulumi.Input

	ToShareAccessArrayOutput() ShareAccessArrayOutput
	ToShareAccessArrayOutputWithContext(context.Context) ShareAccessArrayOutput
}

type ShareAccessArray []ShareAccessInput

func (ShareAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ShareAccess)(nil))
}

func (i ShareAccessArray) ToShareAccessArrayOutput() ShareAccessArrayOutput {
	return i.ToShareAccessArrayOutputWithContext(context.Background())
}

func (i ShareAccessArray) ToShareAccessArrayOutputWithContext(ctx context.Context) ShareAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessArrayOutput)
}

// ShareAccessMapInput is an input type that accepts ShareAccessMap and ShareAccessMapOutput values.
// You can construct a concrete instance of `ShareAccessMapInput` via:
//
//          ShareAccessMap{ "key": ShareAccessArgs{...} }
type ShareAccessMapInput interface {
	pulumi.Input

	ToShareAccessMapOutput() ShareAccessMapOutput
	ToShareAccessMapOutputWithContext(context.Context) ShareAccessMapOutput
}

type ShareAccessMap map[string]ShareAccessInput

func (ShareAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ShareAccess)(nil))
}

func (i ShareAccessMap) ToShareAccessMapOutput() ShareAccessMapOutput {
	return i.ToShareAccessMapOutputWithContext(context.Background())
}

func (i ShareAccessMap) ToShareAccessMapOutputWithContext(ctx context.Context) ShareAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessMapOutput)
}

type ShareAccessOutput struct {
	*pulumi.OutputState
}

func (ShareAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareAccess)(nil))
}

func (o ShareAccessOutput) ToShareAccessOutput() ShareAccessOutput {
	return o
}

func (o ShareAccessOutput) ToShareAccessOutputWithContext(ctx context.Context) ShareAccessOutput {
	return o
}

func (o ShareAccessOutput) ToShareAccessPtrOutput() ShareAccessPtrOutput {
	return o.ToShareAccessPtrOutputWithContext(context.Background())
}

func (o ShareAccessOutput) ToShareAccessPtrOutputWithContext(ctx context.Context) ShareAccessPtrOutput {
	return o.ApplyT(func(v ShareAccess) *ShareAccess {
		return &v
	}).(ShareAccessPtrOutput)
}

type ShareAccessPtrOutput struct {
	*pulumi.OutputState
}

func (ShareAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareAccess)(nil))
}

func (o ShareAccessPtrOutput) ToShareAccessPtrOutput() ShareAccessPtrOutput {
	return o
}

func (o ShareAccessPtrOutput) ToShareAccessPtrOutputWithContext(ctx context.Context) ShareAccessPtrOutput {
	return o
}

type ShareAccessArrayOutput struct{ *pulumi.OutputState }

func (ShareAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareAccess)(nil))
}

func (o ShareAccessArrayOutput) ToShareAccessArrayOutput() ShareAccessArrayOutput {
	return o
}

func (o ShareAccessArrayOutput) ToShareAccessArrayOutputWithContext(ctx context.Context) ShareAccessArrayOutput {
	return o
}

func (o ShareAccessArrayOutput) Index(i pulumi.IntInput) ShareAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareAccess {
		return vs[0].([]ShareAccess)[vs[1].(int)]
	}).(ShareAccessOutput)
}

type ShareAccessMapOutput struct{ *pulumi.OutputState }

func (ShareAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ShareAccess)(nil))
}

func (o ShareAccessMapOutput) ToShareAccessMapOutput() ShareAccessMapOutput {
	return o
}

func (o ShareAccessMapOutput) ToShareAccessMapOutputWithContext(ctx context.Context) ShareAccessMapOutput {
	return o
}

func (o ShareAccessMapOutput) MapIndex(k pulumi.StringInput) ShareAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ShareAccess {
		return vs[0].(map[string]ShareAccess)[vs[1].(string)]
	}).(ShareAccessOutput)
}

func init() {
	pulumi.RegisterOutputType(ShareAccessOutput{})
	pulumi.RegisterOutputType(ShareAccessPtrOutput{})
	pulumi.RegisterOutputType(ShareAccessArrayOutput{})
	pulumi.RegisterOutputType(ShareAccessMapOutput{})
}
