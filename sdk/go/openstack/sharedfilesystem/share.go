// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to configure a share.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/networking"
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/sharedfilesystem"
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
//			_, err = sharedfilesystem.NewShare(ctx, "share_1", &sharedfilesystem.ShareArgs{
//				Name:           pulumi.String("nfs_share"),
//				Description:    pulumi.String("test share description"),
//				ShareProto:     pulumi.String("NFS"),
//				Size:           pulumi.Int(1),
//				ShareNetworkId: sharenetwork1.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource can be imported by specifying the ID of the share:
//
// ```sh
// $ pulumi import openstack:sharedfilesystem/share:Share share_1 id
// ```
type Share struct {
	pulumi.CustomResourceState

	// The map of metadata, assigned on the share, which has been
	// explicitly and implicitly added.
	AllMetadata pulumi.StringMapOutput `pulumi:"allMetadata"`
	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of export locations. For example, when a share server
	// has more than one network interface, it can have multiple export locations.
	ExportLocations ShareExportLocationArrayOutput `pulumi:"exportLocations"`
	// Indicates whether a share has replicas or not.
	HasReplicas pulumi.BoolOutput `pulumi:"hasReplicas"`
	// The share host name.
	Host pulumi.StringOutput `pulumi:"host"`
	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the share. Changing this updates the name
	// of the existing share.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the Share.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Shared File System
	// client. A Shared File System client is needed to create a share. Changing
	// this creates a new share.
	Region pulumi.StringOutput `pulumi:"region"`
	// The share replication type.
	ReplicationType pulumi.StringOutput `pulumi:"replicationType"`
	// The UUID of a share network where the share server exists
	// or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
	// the shareNetworkId value from the snapshot is used. Changing this creates a new share.
	ShareNetworkId pulumi.StringOutput `pulumi:"shareNetworkId"`
	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto pulumi.StringOutput `pulumi:"shareProto"`
	// The UUID of the share server.
	ShareServerId pulumi.StringOutput `pulumi:"shareServerId"`
	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType pulumi.StringOutput `pulumi:"shareType"`
	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size pulumi.IntOutput `pulumi:"size"`
	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
}

// NewShare registers a new resource with the given unique name, arguments, and options.
func NewShare(ctx *pulumi.Context,
	name string, args *ShareArgs, opts ...pulumi.ResourceOption) (*Share, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ShareProto == nil {
		return nil, errors.New("invalid value for required argument 'ShareProto'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Share
	err := ctx.RegisterResource("openstack:sharedfilesystem/share:Share", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShare gets an existing Share resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShareState, opts ...pulumi.ResourceOption) (*Share, error) {
	var resource Share
	err := ctx.ReadResource("openstack:sharedfilesystem/share:Share", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Share resources.
type shareState struct {
	// The map of metadata, assigned on the share, which has been
	// explicitly and implicitly added.
	AllMetadata map[string]string `pulumi:"allMetadata"`
	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description *string `pulumi:"description"`
	// A list of export locations. For example, when a share server
	// has more than one network interface, it can have multiple export locations.
	ExportLocations []ShareExportLocation `pulumi:"exportLocations"`
	// Indicates whether a share has replicas or not.
	HasReplicas *bool `pulumi:"hasReplicas"`
	// The share host name.
	Host *string `pulumi:"host"`
	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic *bool `pulumi:"isPublic"`
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the share. Changing this updates the name
	// of the existing share.
	Name *string `pulumi:"name"`
	// The owner of the Share.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Shared File System
	// client. A Shared File System client is needed to create a share. Changing
	// this creates a new share.
	Region *string `pulumi:"region"`
	// The share replication type.
	ReplicationType *string `pulumi:"replicationType"`
	// The UUID of a share network where the share server exists
	// or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
	// the shareNetworkId value from the snapshot is used. Changing this creates a new share.
	ShareNetworkId *string `pulumi:"shareNetworkId"`
	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto *string `pulumi:"shareProto"`
	// The UUID of the share server.
	ShareServerId *string `pulumi:"shareServerId"`
	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType *string `pulumi:"shareType"`
	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size *int `pulumi:"size"`
	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotId *string `pulumi:"snapshotId"`
}

type ShareState struct {
	// The map of metadata, assigned on the share, which has been
	// explicitly and implicitly added.
	AllMetadata pulumi.StringMapInput
	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone pulumi.StringPtrInput
	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description pulumi.StringPtrInput
	// A list of export locations. For example, when a share server
	// has more than one network interface, it can have multiple export locations.
	ExportLocations ShareExportLocationArrayInput
	// Indicates whether a share has replicas or not.
	HasReplicas pulumi.BoolPtrInput
	// The share host name.
	Host pulumi.StringPtrInput
	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic pulumi.BoolPtrInput
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata pulumi.StringMapInput
	// The name of the share. Changing this updates the name
	// of the existing share.
	Name pulumi.StringPtrInput
	// The owner of the Share.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Shared File System
	// client. A Shared File System client is needed to create a share. Changing
	// this creates a new share.
	Region pulumi.StringPtrInput
	// The share replication type.
	ReplicationType pulumi.StringPtrInput
	// The UUID of a share network where the share server exists
	// or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
	// the shareNetworkId value from the snapshot is used. Changing this creates a new share.
	ShareNetworkId pulumi.StringPtrInput
	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto pulumi.StringPtrInput
	// The UUID of the share server.
	ShareServerId pulumi.StringPtrInput
	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType pulumi.StringPtrInput
	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size pulumi.IntPtrInput
	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotId pulumi.StringPtrInput
}

func (ShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*shareState)(nil)).Elem()
}

type shareArgs struct {
	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description *string `pulumi:"description"`
	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic *bool `pulumi:"isPublic"`
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the share. Changing this updates the name
	// of the existing share.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V2 Shared File System
	// client. A Shared File System client is needed to create a share. Changing
	// this creates a new share.
	Region *string `pulumi:"region"`
	// The UUID of a share network where the share server exists
	// or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
	// the shareNetworkId value from the snapshot is used. Changing this creates a new share.
	ShareNetworkId *string `pulumi:"shareNetworkId"`
	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto string `pulumi:"shareProto"`
	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType *string `pulumi:"shareType"`
	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size int `pulumi:"size"`
	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotId *string `pulumi:"snapshotId"`
}

// The set of arguments for constructing a Share resource.
type ShareArgs struct {
	// The share availability zone. Changing this creates a
	// new share.
	AvailabilityZone pulumi.StringPtrInput
	// The human-readable description for the share.
	// Changing this updates the description of the existing share.
	Description pulumi.StringPtrInput
	// The level of visibility for the share. Set to true to make
	// share public. Set to false to make it private. Default value is false. Changing this
	// updates the existing share.
	IsPublic pulumi.BoolPtrInput
	// One or more metadata key and value pairs as a dictionary of
	// strings.
	Metadata pulumi.StringMapInput
	// The name of the share. Changing this updates the name
	// of the existing share.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V2 Shared File System
	// client. A Shared File System client is needed to create a share. Changing
	// this creates a new share.
	Region pulumi.StringPtrInput
	// The UUID of a share network where the share server exists
	// or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
	// the shareNetworkId value from the snapshot is used. Changing this creates a new share.
	ShareNetworkId pulumi.StringPtrInput
	// The share protocol - can either be NFS, CIFS,
	// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
	ShareProto pulumi.StringInput
	// The share type name. If you omit this parameter, the default
	// share type is used.
	ShareType pulumi.StringPtrInput
	// The share size, in GBs. The requested share size cannot be greater
	// than the allowed GB quota. Changing this resizes the existing share.
	Size pulumi.IntInput
	// The UUID of the share's base snapshot. Changing this creates
	// a new share.
	SnapshotId pulumi.StringPtrInput
}

func (ShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shareArgs)(nil)).Elem()
}

type ShareInput interface {
	pulumi.Input

	ToShareOutput() ShareOutput
	ToShareOutputWithContext(ctx context.Context) ShareOutput
}

func (*Share) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (i *Share) ToShareOutput() ShareOutput {
	return i.ToShareOutputWithContext(context.Background())
}

func (i *Share) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareOutput)
}

// ShareArrayInput is an input type that accepts ShareArray and ShareArrayOutput values.
// You can construct a concrete instance of `ShareArrayInput` via:
//
//	ShareArray{ ShareArgs{...} }
type ShareArrayInput interface {
	pulumi.Input

	ToShareArrayOutput() ShareArrayOutput
	ToShareArrayOutputWithContext(context.Context) ShareArrayOutput
}

type ShareArray []ShareInput

func (ShareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Share)(nil)).Elem()
}

func (i ShareArray) ToShareArrayOutput() ShareArrayOutput {
	return i.ToShareArrayOutputWithContext(context.Background())
}

func (i ShareArray) ToShareArrayOutputWithContext(ctx context.Context) ShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareArrayOutput)
}

// ShareMapInput is an input type that accepts ShareMap and ShareMapOutput values.
// You can construct a concrete instance of `ShareMapInput` via:
//
//	ShareMap{ "key": ShareArgs{...} }
type ShareMapInput interface {
	pulumi.Input

	ToShareMapOutput() ShareMapOutput
	ToShareMapOutputWithContext(context.Context) ShareMapOutput
}

type ShareMap map[string]ShareInput

func (ShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Share)(nil)).Elem()
}

func (i ShareMap) ToShareMapOutput() ShareMapOutput {
	return i.ToShareMapOutputWithContext(context.Background())
}

func (i ShareMap) ToShareMapOutputWithContext(ctx context.Context) ShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareMapOutput)
}

type ShareOutput struct{ *pulumi.OutputState }

func (ShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Share)(nil)).Elem()
}

func (o ShareOutput) ToShareOutput() ShareOutput {
	return o
}

func (o ShareOutput) ToShareOutputWithContext(ctx context.Context) ShareOutput {
	return o
}

// The map of metadata, assigned on the share, which has been
// explicitly and implicitly added.
func (o ShareOutput) AllMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Share) pulumi.StringMapOutput { return v.AllMetadata }).(pulumi.StringMapOutput)
}

// The share availability zone. Changing this creates a
// new share.
func (o ShareOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The human-readable description for the share.
// Changing this updates the description of the existing share.
func (o ShareOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of export locations. For example, when a share server
// has more than one network interface, it can have multiple export locations.
func (o ShareOutput) ExportLocations() ShareExportLocationArrayOutput {
	return o.ApplyT(func(v *Share) ShareExportLocationArrayOutput { return v.ExportLocations }).(ShareExportLocationArrayOutput)
}

// Indicates whether a share has replicas or not.
func (o ShareOutput) HasReplicas() pulumi.BoolOutput {
	return o.ApplyT(func(v *Share) pulumi.BoolOutput { return v.HasReplicas }).(pulumi.BoolOutput)
}

// The share host name.
func (o ShareOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The level of visibility for the share. Set to true to make
// share public. Set to false to make it private. Default value is false. Changing this
// updates the existing share.
func (o ShareOutput) IsPublic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.BoolPtrOutput { return v.IsPublic }).(pulumi.BoolPtrOutput)
}

// One or more metadata key and value pairs as a dictionary of
// strings.
func (o ShareOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Share) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the share. Changing this updates the name
// of the existing share.
func (o ShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The owner of the Share.
func (o ShareOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Shared File System
// client. A Shared File System client is needed to create a share. Changing
// this creates a new share.
func (o ShareOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The share replication type.
func (o ShareOutput) ReplicationType() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ReplicationType }).(pulumi.StringOutput)
}

// The UUID of a share network where the share server exists
// or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
// the shareNetworkId value from the snapshot is used. Changing this creates a new share.
func (o ShareOutput) ShareNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ShareNetworkId }).(pulumi.StringOutput)
}

// The share protocol - can either be NFS, CIFS,
// CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
func (o ShareOutput) ShareProto() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ShareProto }).(pulumi.StringOutput)
}

// The UUID of the share server.
func (o ShareOutput) ShareServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ShareServerId }).(pulumi.StringOutput)
}

// The share type name. If you omit this parameter, the default
// share type is used.
func (o ShareOutput) ShareType() pulumi.StringOutput {
	return o.ApplyT(func(v *Share) pulumi.StringOutput { return v.ShareType }).(pulumi.StringOutput)
}

// The share size, in GBs. The requested share size cannot be greater
// than the allowed GB quota. Changing this resizes the existing share.
func (o ShareOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Share) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The UUID of the share's base snapshot. Changing this creates
// a new share.
func (o ShareOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Share) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

type ShareArrayOutput struct{ *pulumi.OutputState }

func (ShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Share)(nil)).Elem()
}

func (o ShareArrayOutput) ToShareArrayOutput() ShareArrayOutput {
	return o
}

func (o ShareArrayOutput) ToShareArrayOutputWithContext(ctx context.Context) ShareArrayOutput {
	return o
}

func (o ShareArrayOutput) Index(i pulumi.IntInput) ShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Share {
		return vs[0].([]*Share)[vs[1].(int)]
	}).(ShareOutput)
}

type ShareMapOutput struct{ *pulumi.OutputState }

func (ShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Share)(nil)).Elem()
}

func (o ShareMapOutput) ToShareMapOutput() ShareMapOutput {
	return o
}

func (o ShareMapOutput) ToShareMapOutputWithContext(ctx context.Context) ShareMapOutput {
	return o
}

func (o ShareMapOutput) MapIndex(k pulumi.StringInput) ShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Share {
		return vs[0].(map[string]*Share)[vs[1].(string)]
	}).(ShareOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShareInput)(nil)).Elem(), &Share{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareArrayInput)(nil)).Elem(), ShareArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareMapInput)(nil)).Elem(), ShareMap{})
	pulumi.RegisterOutputType(ShareOutput{})
	pulumi.RegisterOutputType(ShareArrayOutput{})
	pulumi.RegisterOutputType(ShareMapOutput{})
}
