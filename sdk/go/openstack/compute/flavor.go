// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V2 flavor resource within OpenStack.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_flavor_v2.html.markdown.
type Flavor struct {
	pulumi.CustomResourceState

	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk      pulumi.IntOutput    `pulumi:"disk"`
	Ephemeral pulumi.IntPtrOutput `pulumi:"ephemeral"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs pulumi.MapOutput `pulumi:"extraSpecs"`
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic pulumi.BoolPtrOutput `pulumi:"isPublic"`
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram pulumi.IntOutput `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region pulumi.StringOutput `pulumi:"region"`
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor pulumi.Float64PtrOutput `pulumi:"rxTxFactor"`
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap pulumi.IntPtrOutput `pulumi:"swap"`
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus pulumi.IntOutput `pulumi:"vcpus"`
}

// NewFlavor registers a new resource with the given unique name, arguments, and options.
func NewFlavor(ctx *pulumi.Context,
	name string, args *FlavorArgs, opts ...pulumi.ResourceOption) (*Flavor, error) {
	if args == nil || args.Disk == nil {
		return nil, errors.New("missing required argument 'Disk'")
	}
	if args == nil || args.Ram == nil {
		return nil, errors.New("missing required argument 'Ram'")
	}
	if args == nil || args.Vcpus == nil {
		return nil, errors.New("missing required argument 'Vcpus'")
	}
	if args == nil {
		args = &FlavorArgs{}
	}
	var resource Flavor
	err := ctx.RegisterResource("openstack:compute/flavor:Flavor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlavor gets an existing Flavor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlavor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlavorState, opts ...pulumi.ResourceOption) (*Flavor, error) {
	var resource Flavor
	err := ctx.ReadResource("openstack:compute/flavor:Flavor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flavor resources.
type flavorState struct {
	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk      *int `pulumi:"disk"`
	Ephemeral *int `pulumi:"ephemeral"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic *bool `pulumi:"isPublic"`
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name *string `pulumi:"name"`
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram *int `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region *string `pulumi:"region"`
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor *float64 `pulumi:"rxTxFactor"`
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap *int `pulumi:"swap"`
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus *int `pulumi:"vcpus"`
}

type FlavorState struct {
	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk      pulumi.IntPtrInput
	Ephemeral pulumi.IntPtrInput
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs pulumi.MapInput
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic pulumi.BoolPtrInput
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name pulumi.StringPtrInput
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram pulumi.IntPtrInput
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region pulumi.StringPtrInput
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor pulumi.Float64PtrInput
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap pulumi.IntPtrInput
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus pulumi.IntPtrInput
}

func (FlavorState) ElementType() reflect.Type {
	return reflect.TypeOf((*flavorState)(nil)).Elem()
}

type flavorArgs struct {
	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk      int  `pulumi:"disk"`
	Ephemeral *int `pulumi:"ephemeral"`
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs map[string]interface{} `pulumi:"extraSpecs"`
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic *bool `pulumi:"isPublic"`
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name *string `pulumi:"name"`
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram int `pulumi:"ram"`
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region *string `pulumi:"region"`
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor *float64 `pulumi:"rxTxFactor"`
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap *int `pulumi:"swap"`
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus int `pulumi:"vcpus"`
}

// The set of arguments for constructing a Flavor resource.
type FlavorArgs struct {
	// The amount of disk space in gigabytes to use for the root
	// (/) partition. Changing this creates a new flavor.
	Disk      pulumi.IntInput
	Ephemeral pulumi.IntPtrInput
	// Key/Value pairs of metadata for the flavor.
	ExtraSpecs pulumi.MapInput
	// Whether the flavor is public. Changing this creates
	// a new flavor.
	IsPublic pulumi.BoolPtrInput
	// A unique name for the flavor. Changing this creates a new
	// flavor.
	Name pulumi.StringPtrInput
	// The amount of RAM to use, in megabytes. Changing this
	// creates a new flavor.
	Ram pulumi.IntInput
	// The region in which to obtain the V2 Compute client.
	// Flavors are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new flavor.
	Region pulumi.StringPtrInput
	// RX/TX bandwith factor. The default is 1. Changing
	// this creates a new flavor.
	RxTxFactor pulumi.Float64PtrInput
	// The amount of disk space in megabytes to use. If
	// unspecified, the default is 0. Changing this creates a new flavor.
	Swap pulumi.IntPtrInput
	// The number of virtual CPUs to use. Changing this creates
	// a new flavor.
	Vcpus pulumi.IntInput
}

func (FlavorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flavorArgs)(nil)).Elem()
}
