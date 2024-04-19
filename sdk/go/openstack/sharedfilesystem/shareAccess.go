// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### NFS
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/sharedfilesystem"
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
//			sharenetwork1, err := sharedfilesystem.NewShareNetwork(ctx, "sharenetwork_1", &sharedfilesystem.ShareNetworkArgs{
//				Name:            pulumi.String("test_sharenetwork"),
//				Description:     pulumi.String("test share network with security services"),
//				NeutronNetId:    network1.ID(),
//				NeutronSubnetId: subnet1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			share1, err := sharedfilesystem.NewShare(ctx, "share_1", &sharedfilesystem.ShareArgs{
//				Name:           pulumi.String("nfs_share"),
//				Description:    pulumi.String("test share description"),
//				ShareProto:     pulumi.String("NFS"),
//				Size:           pulumi.Int(1),
//				ShareNetworkId: sharenetwork1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sharedfilesystem.NewShareAccess(ctx, "share_access_1", &sharedfilesystem.ShareAccessArgs{
//				ShareId:     share1.ID(),
//				AccessType:  pulumi.String("ip"),
//				AccessTo:    pulumi.String("192.168.199.10"),
//				AccessLevel: pulumi.String("rw"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### CIFS
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/sharedfilesystem"
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
//			securityservice1, err := sharedfilesystem.NewSecurityService(ctx, "securityservice_1", &sharedfilesystem.SecurityServiceArgs{
//				Name:        pulumi.String("security"),
//				Description: pulumi.String("created by terraform"),
//				Type:        pulumi.String("active_directory"),
//				Server:      pulumi.String("192.168.199.10"),
//				DnsIp:       pulumi.String("192.168.199.10"),
//				Domain:      pulumi.String("example.com"),
//				Ou:          pulumi.String("CN=Computers,DC=example,DC=com"),
//				User:        pulumi.String("joinDomainUser"),
//				Password:    pulumi.String("s8cret"),
//			})
//			if err != nil {
//				return err
//			}
//			sharenetwork1, err := sharedfilesystem.NewShareNetwork(ctx, "sharenetwork_1", &sharedfilesystem.ShareNetworkArgs{
//				Name:            pulumi.String("test_sharenetwork_secure"),
//				Description:     pulumi.String("share the secure love"),
//				NeutronNetId:    network1.ID(),
//				NeutronSubnetId: subnet1.ID(),
//				SecurityServiceIds: pulumi.StringArray{
//					securityservice1.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			share1, err := sharedfilesystem.NewShare(ctx, "share_1", &sharedfilesystem.ShareArgs{
//				Name:           pulumi.String("cifs_share"),
//				ShareProto:     pulumi.String("CIFS"),
//				Size:           pulumi.Int(1),
//				ShareNetworkId: sharenetwork1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sharedfilesystem.NewShareAccess(ctx, "share_access_1", &sharedfilesystem.ShareAccessArgs{
//				ShareId:     share1.ID(),
//				AccessType:  pulumi.String("user"),
//				AccessTo:    pulumi.String("windows"),
//				AccessLevel: pulumi.String("ro"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = sharedfilesystem.NewShareAccess(ctx, "share_access_2", &sharedfilesystem.ShareAccessArgs{
//				ShareId:     share1.ID(),
//				AccessType:  pulumi.String("user"),
//				AccessTo:    pulumi.String("linux"),
//				AccessLevel: pulumi.String("rw"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("exportLocations", share1.ExportLocations)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// This resource can be imported by specifying the ID of the share and the ID of the
// share access, separated by a slash, e.g.:
//
// ```sh
// $ pulumi import openstack:sharedfilesystem/shareAccess:ShareAccess share_access_1 share_id/share_access_id
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
	// The share access state.
	State pulumi.StringOutput `pulumi:"state"`
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
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"accessKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// The share access state.
	State *string `pulumi:"state"`
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
	// The share access state.
	State pulumi.StringPtrInput
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
	return reflect.TypeOf((**ShareAccess)(nil)).Elem()
}

func (i *ShareAccess) ToShareAccessOutput() ShareAccessOutput {
	return i.ToShareAccessOutputWithContext(context.Background())
}

func (i *ShareAccess) ToShareAccessOutputWithContext(ctx context.Context) ShareAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessOutput)
}

// ShareAccessArrayInput is an input type that accepts ShareAccessArray and ShareAccessArrayOutput values.
// You can construct a concrete instance of `ShareAccessArrayInput` via:
//
//	ShareAccessArray{ ShareAccessArgs{...} }
type ShareAccessArrayInput interface {
	pulumi.Input

	ToShareAccessArrayOutput() ShareAccessArrayOutput
	ToShareAccessArrayOutputWithContext(context.Context) ShareAccessArrayOutput
}

type ShareAccessArray []ShareAccessInput

func (ShareAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShareAccess)(nil)).Elem()
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
//	ShareAccessMap{ "key": ShareAccessArgs{...} }
type ShareAccessMapInput interface {
	pulumi.Input

	ToShareAccessMapOutput() ShareAccessMapOutput
	ToShareAccessMapOutputWithContext(context.Context) ShareAccessMapOutput
}

type ShareAccessMap map[string]ShareAccessInput

func (ShareAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShareAccess)(nil)).Elem()
}

func (i ShareAccessMap) ToShareAccessMapOutput() ShareAccessMapOutput {
	return i.ToShareAccessMapOutputWithContext(context.Background())
}

func (i ShareAccessMap) ToShareAccessMapOutputWithContext(ctx context.Context) ShareAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareAccessMapOutput)
}

type ShareAccessOutput struct{ *pulumi.OutputState }

func (ShareAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShareAccess)(nil)).Elem()
}

func (o ShareAccessOutput) ToShareAccessOutput() ShareAccessOutput {
	return o
}

func (o ShareAccessOutput) ToShareAccessOutputWithContext(ctx context.Context) ShareAccessOutput {
	return o
}

// The access credential of the entity granted access.
func (o ShareAccessOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.AccessKey }).(pulumi.StringOutput)
}

// The access level to the share. Can either be `rw` or `ro`.
func (o ShareAccessOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// The value that defines the access. Can either be an IP
// address or a username verified by configured Security Service of the Share Network.
func (o ShareAccessOutput) AccessTo() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.AccessTo }).(pulumi.StringOutput)
}

// The access rule type. Can either be an ip, user,
// cert, or cephx. cephx support requires an OpenStack environment that supports
// Shared Filesystem microversion 2.13 (Mitaka) or later.
func (o ShareAccessOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.AccessType }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Shared File System client.
// A Shared File System client is needed to create a share access. Changing this
// creates a new share access.
func (o ShareAccessOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The UUID of the share to which you are granted access.
func (o ShareAccessOutput) ShareId() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.ShareId }).(pulumi.StringOutput)
}

// The share access state.
func (o ShareAccessOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ShareAccess) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type ShareAccessArrayOutput struct{ *pulumi.OutputState }

func (ShareAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShareAccess)(nil)).Elem()
}

func (o ShareAccessArrayOutput) ToShareAccessArrayOutput() ShareAccessArrayOutput {
	return o
}

func (o ShareAccessArrayOutput) ToShareAccessArrayOutputWithContext(ctx context.Context) ShareAccessArrayOutput {
	return o
}

func (o ShareAccessArrayOutput) Index(i pulumi.IntInput) ShareAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShareAccess {
		return vs[0].([]*ShareAccess)[vs[1].(int)]
	}).(ShareAccessOutput)
}

type ShareAccessMapOutput struct{ *pulumi.OutputState }

func (ShareAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShareAccess)(nil)).Elem()
}

func (o ShareAccessMapOutput) ToShareAccessMapOutput() ShareAccessMapOutput {
	return o
}

func (o ShareAccessMapOutput) ToShareAccessMapOutputWithContext(ctx context.Context) ShareAccessMapOutput {
	return o
}

func (o ShareAccessMapOutput) MapIndex(k pulumi.StringInput) ShareAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShareAccess {
		return vs[0].(map[string]*ShareAccess)[vs[1].(string)]
	}).(ShareAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShareAccessInput)(nil)).Elem(), &ShareAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareAccessArrayInput)(nil)).Elem(), ShareAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareAccessMapInput)(nil)).Elem(), ShareAccessMap{})
	pulumi.RegisterOutputType(ShareAccessOutput{})
	pulumi.RegisterOutputType(ShareAccessArrayOutput{})
	pulumi.RegisterOutputType(ShareAccessMapOutput{})
}
