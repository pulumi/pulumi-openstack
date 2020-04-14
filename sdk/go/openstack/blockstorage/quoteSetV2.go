// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 block storage quotaset resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
//     in case of delete call.
type QuoteSetV2 struct {
	pulumi.CustomResourceState

	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes pulumi.IntOutput `pulumi:"backupGigabytes"`
	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups pulumi.IntOutput `pulumi:"backups"`
	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes pulumi.IntOutput `pulumi:"gigabytes"`
	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups pulumi.IntOutput `pulumi:"groups"`
	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes pulumi.IntOutput `pulumi:"perVolumeGigabytes"`
	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringOutput `pulumi:"region"`
	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots pulumi.IntOutput `pulumi:"snapshots"`
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes pulumi.IntOutput `pulumi:"volumes"`
}

// NewQuoteSetV2 registers a new resource with the given unique name, arguments, and options.
func NewQuoteSetV2(ctx *pulumi.Context,
	name string, args *QuoteSetV2Args, opts ...pulumi.ResourceOption) (*QuoteSetV2, error) {
	if args == nil || args.ProjectId == nil {
		return nil, errors.New("missing required argument 'ProjectId'")
	}
	if args == nil {
		args = &QuoteSetV2Args{}
	}
	var resource QuoteSetV2
	err := ctx.RegisterResource("openstack:blockstorage/quoteSetV2:QuoteSetV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuoteSetV2 gets an existing QuoteSetV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuoteSetV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuoteSetV2State, opts ...pulumi.ResourceOption) (*QuoteSetV2, error) {
	var resource QuoteSetV2
	err := ctx.ReadResource("openstack:blockstorage/quoteSetV2:QuoteSetV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuoteSetV2 resources.
type quoteSetV2State struct {
	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes *int `pulumi:"backupGigabytes"`
	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups *int `pulumi:"backups"`
	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes *int `pulumi:"gigabytes"`
	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups *int `pulumi:"groups"`
	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes *int `pulumi:"perVolumeGigabytes"`
	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots *int `pulumi:"snapshots"`
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes *int `pulumi:"volumes"`
}

type QuoteSetV2State struct {
	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes pulumi.IntPtrInput
	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups pulumi.IntPtrInput
	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes pulumi.IntPtrInput
	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups pulumi.IntPtrInput
	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes pulumi.IntPtrInput
	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	ProjectId pulumi.StringPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots pulumi.IntPtrInput
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes pulumi.IntPtrInput
}

func (QuoteSetV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*quoteSetV2State)(nil)).Elem()
}

type quoteSetV2Args struct {
	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes *int `pulumi:"backupGigabytes"`
	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups *int `pulumi:"backups"`
	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes *int `pulumi:"gigabytes"`
	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups *int `pulumi:"groups"`
	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes *int `pulumi:"perVolumeGigabytes"`
	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	ProjectId string `pulumi:"projectId"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots *int `pulumi:"snapshots"`
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes *int `pulumi:"volumes"`
}

// The set of arguments for constructing a QuoteSetV2 resource.
type QuoteSetV2Args struct {
	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes pulumi.IntPtrInput
	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups pulumi.IntPtrInput
	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes pulumi.IntPtrInput
	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups pulumi.IntPtrInput
	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes pulumi.IntPtrInput
	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	ProjectId pulumi.StringInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots pulumi.IntPtrInput
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes pulumi.IntPtrInput
}

func (QuoteSetV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*quoteSetV2Args)(nil)).Elem()
}
