// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Instance struct {
	pulumi.CustomResourceState

	// The first detected Fixed IPv4 address.
	AccessIpV4 pulumi.StringOutput `pulumi:"accessIpV4"`
	// The first detected Fixed IPv6 address.
	AccessIpV6 pulumi.StringOutput `pulumi:"accessIpV6"`
	// The administrative password to assign to the server.
	// Changing this changes the root password on the existing server.
	AdminPass   pulumi.StringPtrOutput `pulumi:"adminPass"`
	AllMetadata pulumi.MapOutput       `pulumi:"allMetadata"`
	// The collection of tags assigned on the instance, which have
	// been explicitly and implicitly added.
	AllTags pulumi.StringArrayOutput `pulumi:"allTags"`
	// The availability zone in which to create
	// the server. Conflicts with `availabilityZoneHints`. Changing this creates
	// a new server.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The availability zone in which to
	// create the server. This argument is preferred to `availabilityZone`, when
	// scheduling the server on a
	// [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
	// host or node. Conflicts with `availabilityZone`. Changing this creates a
	// new server.
	AvailabilityZoneHints pulumi.StringPtrOutput `pulumi:"availabilityZoneHints"`
	// Configuration of block devices. The blockDevice
	// structure is documented below. Changing this creates a new server.
	// You can specify multiple block devices which will create an instance with
	// multiple disks. This configuration is very flexible, so please see the
	// following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
	// for more information.
	BlockDevices InstanceBlockDeviceArrayOutput `pulumi:"blockDevices"`
	// Whether to use the configDrive feature to
	// configure the instance. Changing this creates a new server.
	ConfigDrive pulumi.BoolPtrOutput `pulumi:"configDrive"`
	// The flavor ID of
	// the desired flavor for the server. Changing this resizes the existing server.
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// The name of the
	// desired flavor for the server. Changing this resizes the existing server.
	FlavorName pulumi.StringOutput `pulumi:"flavorName"`
	// Whether to force the OpenStack instance to be
	// forcefully deleted. This is useful for environments that have reclaim / soft
	// deletion enabled.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// (Optional; Required if `imageName` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The image ID of
	// the desired image for the server. Changing this creates a new server.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// (Optional; Required if `imageId` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The name of the
	// desired image for the server. Changing this creates a new server.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// Metadata key/value pairs to make available from
	// within the instance. Changing this updates the existing server metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The human-readable
	// name of the network. Changing this creates a new server.
	Name pulumi.StringOutput `pulumi:"name"`
	// Special string for `network` option to create
	// the server. `networkMode` can be `"auto"` or `"none"`.
	// Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
	NetworkMode pulumi.StringPtrOutput `pulumi:"networkMode"`
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new server.
	Networks InstanceNetworkArrayOutput `pulumi:"networks"`
	// Customize the personality of an instance by
	// defining one or more files and their contents. The personality structure
	// is described below.
	Personalities InstancePersonalityArrayOutput `pulumi:"personalities"`
	// Provide the VM state. Only 'active' and 'shutoff'
	// are supported values. *Note*: If the initial powerState is the shutoff
	// the VM will be stopped immediately after build and the provisioners like
	// remote-exec or files are not supported.
	PowerState pulumi.StringPtrOutput `pulumi:"powerState"`
	// The region in which to create the server instance. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new server.
	Region pulumi.StringOutput `pulumi:"region"`
	// Provide the Nova scheduler with hints on how
	// the instance should be launched. The available hints are described below.
	SchedulerHints InstanceSchedulerHintArrayOutput `pulumi:"schedulerHints"`
	// An array of one or more security group names
	// or ids to associate with the server. Changing this results in adding/removing
	// security groups from the existing server. *Note*: When attaching the
	// instance to networks using Ports, place the security groups on the Port
	// and not the instance.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Whether to try stop instance gracefully
	// before destroying it, thus giving chance for guest OS daemons to stop correctly.
	// If instance doesn't stop within timeout, it will be destroyed anyway.
	StopBeforeDestroy pulumi.BoolPtrOutput `pulumi:"stopBeforeDestroy"`
	// A set of string tags for the instance. Changing this
	// updates the existing instance tags.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions InstanceVendorOptionsPtrOutput `pulumi:"vendorOptions"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var resource Instance
	err := ctx.RegisterResource("openstack:compute/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("openstack:compute/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The first detected Fixed IPv4 address.
	AccessIpV4 *string `pulumi:"accessIpV4"`
	// The first detected Fixed IPv6 address.
	AccessIpV6 *string `pulumi:"accessIpV6"`
	// The administrative password to assign to the server.
	// Changing this changes the root password on the existing server.
	AdminPass   *string                `pulumi:"adminPass"`
	AllMetadata map[string]interface{} `pulumi:"allMetadata"`
	// The collection of tags assigned on the instance, which have
	// been explicitly and implicitly added.
	AllTags []string `pulumi:"allTags"`
	// The availability zone in which to create
	// the server. Conflicts with `availabilityZoneHints`. Changing this creates
	// a new server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The availability zone in which to
	// create the server. This argument is preferred to `availabilityZone`, when
	// scheduling the server on a
	// [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
	// host or node. Conflicts with `availabilityZone`. Changing this creates a
	// new server.
	AvailabilityZoneHints *string `pulumi:"availabilityZoneHints"`
	// Configuration of block devices. The blockDevice
	// structure is documented below. Changing this creates a new server.
	// You can specify multiple block devices which will create an instance with
	// multiple disks. This configuration is very flexible, so please see the
	// following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
	// for more information.
	BlockDevices []InstanceBlockDevice `pulumi:"blockDevices"`
	// Whether to use the configDrive feature to
	// configure the instance. Changing this creates a new server.
	ConfigDrive *bool `pulumi:"configDrive"`
	// The flavor ID of
	// the desired flavor for the server. Changing this resizes the existing server.
	FlavorId *string `pulumi:"flavorId"`
	// The name of the
	// desired flavor for the server. Changing this resizes the existing server.
	FlavorName *string `pulumi:"flavorName"`
	// Whether to force the OpenStack instance to be
	// forcefully deleted. This is useful for environments that have reclaim / soft
	// deletion enabled.
	ForceDelete *bool `pulumi:"forceDelete"`
	// (Optional; Required if `imageName` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The image ID of
	// the desired image for the server. Changing this creates a new server.
	ImageId *string `pulumi:"imageId"`
	// (Optional; Required if `imageId` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The name of the
	// desired image for the server. Changing this creates a new server.
	ImageName *string `pulumi:"imageName"`
	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	KeyPair *string `pulumi:"keyPair"`
	// Metadata key/value pairs to make available from
	// within the instance. Changing this updates the existing server metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The human-readable
	// name of the network. Changing this creates a new server.
	Name *string `pulumi:"name"`
	// Special string for `network` option to create
	// the server. `networkMode` can be `"auto"` or `"none"`.
	// Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
	NetworkMode *string `pulumi:"networkMode"`
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new server.
	Networks []InstanceNetwork `pulumi:"networks"`
	// Customize the personality of an instance by
	// defining one or more files and their contents. The personality structure
	// is described below.
	Personalities []InstancePersonality `pulumi:"personalities"`
	// Provide the VM state. Only 'active' and 'shutoff'
	// are supported values. *Note*: If the initial powerState is the shutoff
	// the VM will be stopped immediately after build and the provisioners like
	// remote-exec or files are not supported.
	PowerState *string `pulumi:"powerState"`
	// The region in which to create the server instance. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new server.
	Region *string `pulumi:"region"`
	// Provide the Nova scheduler with hints on how
	// the instance should be launched. The available hints are described below.
	SchedulerHints []InstanceSchedulerHint `pulumi:"schedulerHints"`
	// An array of one or more security group names
	// or ids to associate with the server. Changing this results in adding/removing
	// security groups from the existing server. *Note*: When attaching the
	// instance to networks using Ports, place the security groups on the Port
	// and not the instance.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Whether to try stop instance gracefully
	// before destroying it, thus giving chance for guest OS daemons to stop correctly.
	// If instance doesn't stop within timeout, it will be destroyed anyway.
	StopBeforeDestroy *bool `pulumi:"stopBeforeDestroy"`
	// A set of string tags for the instance. Changing this
	// updates the existing instance tags.
	Tags []string `pulumi:"tags"`
	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData *string `pulumi:"userData"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *InstanceVendorOptions `pulumi:"vendorOptions"`
}

type InstanceState struct {
	// The first detected Fixed IPv4 address.
	AccessIpV4 pulumi.StringPtrInput
	// The first detected Fixed IPv6 address.
	AccessIpV6 pulumi.StringPtrInput
	// The administrative password to assign to the server.
	// Changing this changes the root password on the existing server.
	AdminPass   pulumi.StringPtrInput
	AllMetadata pulumi.MapInput
	// The collection of tags assigned on the instance, which have
	// been explicitly and implicitly added.
	AllTags pulumi.StringArrayInput
	// The availability zone in which to create
	// the server. Conflicts with `availabilityZoneHints`. Changing this creates
	// a new server.
	AvailabilityZone pulumi.StringPtrInput
	// The availability zone in which to
	// create the server. This argument is preferred to `availabilityZone`, when
	// scheduling the server on a
	// [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
	// host or node. Conflicts with `availabilityZone`. Changing this creates a
	// new server.
	AvailabilityZoneHints pulumi.StringPtrInput
	// Configuration of block devices. The blockDevice
	// structure is documented below. Changing this creates a new server.
	// You can specify multiple block devices which will create an instance with
	// multiple disks. This configuration is very flexible, so please see the
	// following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
	// for more information.
	BlockDevices InstanceBlockDeviceArrayInput
	// Whether to use the configDrive feature to
	// configure the instance. Changing this creates a new server.
	ConfigDrive pulumi.BoolPtrInput
	// The flavor ID of
	// the desired flavor for the server. Changing this resizes the existing server.
	FlavorId pulumi.StringPtrInput
	// The name of the
	// desired flavor for the server. Changing this resizes the existing server.
	FlavorName pulumi.StringPtrInput
	// Whether to force the OpenStack instance to be
	// forcefully deleted. This is useful for environments that have reclaim / soft
	// deletion enabled.
	ForceDelete pulumi.BoolPtrInput
	// (Optional; Required if `imageName` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The image ID of
	// the desired image for the server. Changing this creates a new server.
	ImageId pulumi.StringPtrInput
	// (Optional; Required if `imageId` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The name of the
	// desired image for the server. Changing this creates a new server.
	ImageName pulumi.StringPtrInput
	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	KeyPair pulumi.StringPtrInput
	// Metadata key/value pairs to make available from
	// within the instance. Changing this updates the existing server metadata.
	Metadata pulumi.MapInput
	// The human-readable
	// name of the network. Changing this creates a new server.
	Name pulumi.StringPtrInput
	// Special string for `network` option to create
	// the server. `networkMode` can be `"auto"` or `"none"`.
	// Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
	NetworkMode pulumi.StringPtrInput
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new server.
	Networks InstanceNetworkArrayInput
	// Customize the personality of an instance by
	// defining one or more files and their contents. The personality structure
	// is described below.
	Personalities InstancePersonalityArrayInput
	// Provide the VM state. Only 'active' and 'shutoff'
	// are supported values. *Note*: If the initial powerState is the shutoff
	// the VM will be stopped immediately after build and the provisioners like
	// remote-exec or files are not supported.
	PowerState pulumi.StringPtrInput
	// The region in which to create the server instance. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new server.
	Region pulumi.StringPtrInput
	// Provide the Nova scheduler with hints on how
	// the instance should be launched. The available hints are described below.
	SchedulerHints InstanceSchedulerHintArrayInput
	// An array of one or more security group names
	// or ids to associate with the server. Changing this results in adding/removing
	// security groups from the existing server. *Note*: When attaching the
	// instance to networks using Ports, place the security groups on the Port
	// and not the instance.
	SecurityGroups pulumi.StringArrayInput
	// Whether to try stop instance gracefully
	// before destroying it, thus giving chance for guest OS daemons to stop correctly.
	// If instance doesn't stop within timeout, it will be destroyed anyway.
	StopBeforeDestroy pulumi.BoolPtrInput
	// A set of string tags for the instance. Changing this
	// updates the existing instance tags.
	Tags pulumi.StringArrayInput
	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData pulumi.StringPtrInput
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions InstanceVendorOptionsPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The first detected Fixed IPv4 address.
	AccessIpV4 *string `pulumi:"accessIpV4"`
	// The first detected Fixed IPv6 address.
	AccessIpV6 *string `pulumi:"accessIpV6"`
	// The administrative password to assign to the server.
	// Changing this changes the root password on the existing server.
	AdminPass *string `pulumi:"adminPass"`
	// The availability zone in which to create
	// the server. Conflicts with `availabilityZoneHints`. Changing this creates
	// a new server.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The availability zone in which to
	// create the server. This argument is preferred to `availabilityZone`, when
	// scheduling the server on a
	// [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
	// host or node. Conflicts with `availabilityZone`. Changing this creates a
	// new server.
	AvailabilityZoneHints *string `pulumi:"availabilityZoneHints"`
	// Configuration of block devices. The blockDevice
	// structure is documented below. Changing this creates a new server.
	// You can specify multiple block devices which will create an instance with
	// multiple disks. This configuration is very flexible, so please see the
	// following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
	// for more information.
	BlockDevices []InstanceBlockDevice `pulumi:"blockDevices"`
	// Whether to use the configDrive feature to
	// configure the instance. Changing this creates a new server.
	ConfigDrive *bool `pulumi:"configDrive"`
	// The flavor ID of
	// the desired flavor for the server. Changing this resizes the existing server.
	FlavorId *string `pulumi:"flavorId"`
	// The name of the
	// desired flavor for the server. Changing this resizes the existing server.
	FlavorName *string `pulumi:"flavorName"`
	// Whether to force the OpenStack instance to be
	// forcefully deleted. This is useful for environments that have reclaim / soft
	// deletion enabled.
	ForceDelete *bool `pulumi:"forceDelete"`
	// (Optional; Required if `imageName` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The image ID of
	// the desired image for the server. Changing this creates a new server.
	ImageId *string `pulumi:"imageId"`
	// (Optional; Required if `imageId` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The name of the
	// desired image for the server. Changing this creates a new server.
	ImageName *string `pulumi:"imageName"`
	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	KeyPair *string `pulumi:"keyPair"`
	// Metadata key/value pairs to make available from
	// within the instance. Changing this updates the existing server metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The human-readable
	// name of the network. Changing this creates a new server.
	Name *string `pulumi:"name"`
	// Special string for `network` option to create
	// the server. `networkMode` can be `"auto"` or `"none"`.
	// Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
	NetworkMode *string `pulumi:"networkMode"`
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new server.
	Networks []InstanceNetwork `pulumi:"networks"`
	// Customize the personality of an instance by
	// defining one or more files and their contents. The personality structure
	// is described below.
	Personalities []InstancePersonality `pulumi:"personalities"`
	// Provide the VM state. Only 'active' and 'shutoff'
	// are supported values. *Note*: If the initial powerState is the shutoff
	// the VM will be stopped immediately after build and the provisioners like
	// remote-exec or files are not supported.
	PowerState *string `pulumi:"powerState"`
	// The region in which to create the server instance. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new server.
	Region *string `pulumi:"region"`
	// Provide the Nova scheduler with hints on how
	// the instance should be launched. The available hints are described below.
	SchedulerHints []InstanceSchedulerHint `pulumi:"schedulerHints"`
	// An array of one or more security group names
	// or ids to associate with the server. Changing this results in adding/removing
	// security groups from the existing server. *Note*: When attaching the
	// instance to networks using Ports, place the security groups on the Port
	// and not the instance.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Whether to try stop instance gracefully
	// before destroying it, thus giving chance for guest OS daemons to stop correctly.
	// If instance doesn't stop within timeout, it will be destroyed anyway.
	StopBeforeDestroy *bool `pulumi:"stopBeforeDestroy"`
	// A set of string tags for the instance. Changing this
	// updates the existing instance tags.
	Tags []string `pulumi:"tags"`
	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData *string `pulumi:"userData"`
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions *InstanceVendorOptions `pulumi:"vendorOptions"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The first detected Fixed IPv4 address.
	AccessIpV4 pulumi.StringPtrInput
	// The first detected Fixed IPv6 address.
	AccessIpV6 pulumi.StringPtrInput
	// The administrative password to assign to the server.
	// Changing this changes the root password on the existing server.
	AdminPass pulumi.StringPtrInput
	// The availability zone in which to create
	// the server. Conflicts with `availabilityZoneHints`. Changing this creates
	// a new server.
	AvailabilityZone pulumi.StringPtrInput
	// The availability zone in which to
	// create the server. This argument is preferred to `availabilityZone`, when
	// scheduling the server on a
	// [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
	// host or node. Conflicts with `availabilityZone`. Changing this creates a
	// new server.
	AvailabilityZoneHints pulumi.StringPtrInput
	// Configuration of block devices. The blockDevice
	// structure is documented below. Changing this creates a new server.
	// You can specify multiple block devices which will create an instance with
	// multiple disks. This configuration is very flexible, so please see the
	// following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
	// for more information.
	BlockDevices InstanceBlockDeviceArrayInput
	// Whether to use the configDrive feature to
	// configure the instance. Changing this creates a new server.
	ConfigDrive pulumi.BoolPtrInput
	// The flavor ID of
	// the desired flavor for the server. Changing this resizes the existing server.
	FlavorId pulumi.StringPtrInput
	// The name of the
	// desired flavor for the server. Changing this resizes the existing server.
	FlavorName pulumi.StringPtrInput
	// Whether to force the OpenStack instance to be
	// forcefully deleted. This is useful for environments that have reclaim / soft
	// deletion enabled.
	ForceDelete pulumi.BoolPtrInput
	// (Optional; Required if `imageName` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The image ID of
	// the desired image for the server. Changing this creates a new server.
	ImageId pulumi.StringPtrInput
	// (Optional; Required if `imageId` is empty and not booting
	// from a volume. Do not specify if booting from a volume.) The name of the
	// desired image for the server. Changing this creates a new server.
	ImageName pulumi.StringPtrInput
	// The name of a key pair to put on the server. The key
	// pair must already be created and associated with the tenant's account.
	// Changing this creates a new server.
	KeyPair pulumi.StringPtrInput
	// Metadata key/value pairs to make available from
	// within the instance. Changing this updates the existing server metadata.
	Metadata pulumi.MapInput
	// The human-readable
	// name of the network. Changing this creates a new server.
	Name pulumi.StringPtrInput
	// Special string for `network` option to create
	// the server. `networkMode` can be `"auto"` or `"none"`.
	// Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
	NetworkMode pulumi.StringPtrInput
	// An array of one or more networks to attach to the
	// instance. The network object structure is documented below. Changing this
	// creates a new server.
	Networks InstanceNetworkArrayInput
	// Customize the personality of an instance by
	// defining one or more files and their contents. The personality structure
	// is described below.
	Personalities InstancePersonalityArrayInput
	// Provide the VM state. Only 'active' and 'shutoff'
	// are supported values. *Note*: If the initial powerState is the shutoff
	// the VM will be stopped immediately after build and the provisioners like
	// remote-exec or files are not supported.
	PowerState pulumi.StringPtrInput
	// The region in which to create the server instance. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new server.
	Region pulumi.StringPtrInput
	// Provide the Nova scheduler with hints on how
	// the instance should be launched. The available hints are described below.
	SchedulerHints InstanceSchedulerHintArrayInput
	// An array of one or more security group names
	// or ids to associate with the server. Changing this results in adding/removing
	// security groups from the existing server. *Note*: When attaching the
	// instance to networks using Ports, place the security groups on the Port
	// and not the instance.
	SecurityGroups pulumi.StringArrayInput
	// Whether to try stop instance gracefully
	// before destroying it, thus giving chance for guest OS daemons to stop correctly.
	// If instance doesn't stop within timeout, it will be destroyed anyway.
	StopBeforeDestroy pulumi.BoolPtrInput
	// A set of string tags for the instance. Changing this
	// updates the existing instance tags.
	Tags pulumi.StringArrayInput
	// The user data to provide when launching the instance.
	// Changing this creates a new server.
	UserData pulumi.StringPtrInput
	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions InstanceVendorOptionsPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil)).Elem()
}

func (i Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceOutput)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
}
