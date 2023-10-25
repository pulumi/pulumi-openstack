// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V2 block storage quotaset resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
//
//	in case of delete call.
//
// > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
//
//	created with zero value. This excludes volume type quota.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project1, err := identity.NewProject(ctx, "project1", nil)
//			if err != nil {
//				return err
//			}
//			_, err = blockstorage.NewQuoteSetV2(ctx, "quotaset1", &blockstorage.QuoteSetV2Args{
//				ProjectId:          project1.ID(),
//				Volumes:            pulumi.Int(10),
//				Snapshots:          pulumi.Int(4),
//				Gigabytes:          pulumi.Int(100),
//				PerVolumeGigabytes: pulumi.Int(10),
//				Backups:            pulumi.Int(4),
//				BackupGigabytes:    pulumi.Int(10),
//				Groups:             pulumi.Int(100),
//				VolumeTypeQuota: pulumi.Map{
//					"volumes_ssd":   pulumi.Any(30),
//					"gigabytes_ssd": pulumi.Any(500),
//					"snapshots_ssd": pulumi.Any(10),
//				},
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
// Quotasets can be imported using the `project_id/region`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:blockstorage/quoteSetV2:QuoteSetV2 quotaset_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
//
// ```
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
	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are `snapshots_<volume_type_name>`,
	// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
	VolumeTypeQuota pulumi.MapOutput `pulumi:"volumeTypeQuota"`
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes pulumi.IntOutput `pulumi:"volumes"`
}

// NewQuoteSetV2 registers a new resource with the given unique name, arguments, and options.
func NewQuoteSetV2(ctx *pulumi.Context,
	name string, args *QuoteSetV2Args, opts ...pulumi.ResourceOption) (*QuoteSetV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are `snapshots_<volume_type_name>`,
	// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
	VolumeTypeQuota map[string]interface{} `pulumi:"volumeTypeQuota"`
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
	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are `snapshots_<volume_type_name>`,
	// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
	VolumeTypeQuota pulumi.MapInput
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
	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are `snapshots_<volume_type_name>`,
	// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
	VolumeTypeQuota map[string]interface{} `pulumi:"volumeTypeQuota"`
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
	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are `snapshots_<volume_type_name>`,
	// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
	VolumeTypeQuota pulumi.MapInput
	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes pulumi.IntPtrInput
}

func (QuoteSetV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*quoteSetV2Args)(nil)).Elem()
}

type QuoteSetV2Input interface {
	pulumi.Input

	ToQuoteSetV2Output() QuoteSetV2Output
	ToQuoteSetV2OutputWithContext(ctx context.Context) QuoteSetV2Output
}

func (*QuoteSetV2) ElementType() reflect.Type {
	return reflect.TypeOf((**QuoteSetV2)(nil)).Elem()
}

func (i *QuoteSetV2) ToQuoteSetV2Output() QuoteSetV2Output {
	return i.ToQuoteSetV2OutputWithContext(context.Background())
}

func (i *QuoteSetV2) ToQuoteSetV2OutputWithContext(ctx context.Context) QuoteSetV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(QuoteSetV2Output)
}

func (i *QuoteSetV2) ToOutput(ctx context.Context) pulumix.Output[*QuoteSetV2] {
	return pulumix.Output[*QuoteSetV2]{
		OutputState: i.ToQuoteSetV2OutputWithContext(ctx).OutputState,
	}
}

// QuoteSetV2ArrayInput is an input type that accepts QuoteSetV2Array and QuoteSetV2ArrayOutput values.
// You can construct a concrete instance of `QuoteSetV2ArrayInput` via:
//
//	QuoteSetV2Array{ QuoteSetV2Args{...} }
type QuoteSetV2ArrayInput interface {
	pulumi.Input

	ToQuoteSetV2ArrayOutput() QuoteSetV2ArrayOutput
	ToQuoteSetV2ArrayOutputWithContext(context.Context) QuoteSetV2ArrayOutput
}

type QuoteSetV2Array []QuoteSetV2Input

func (QuoteSetV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuoteSetV2)(nil)).Elem()
}

func (i QuoteSetV2Array) ToQuoteSetV2ArrayOutput() QuoteSetV2ArrayOutput {
	return i.ToQuoteSetV2ArrayOutputWithContext(context.Background())
}

func (i QuoteSetV2Array) ToQuoteSetV2ArrayOutputWithContext(ctx context.Context) QuoteSetV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuoteSetV2ArrayOutput)
}

func (i QuoteSetV2Array) ToOutput(ctx context.Context) pulumix.Output[[]*QuoteSetV2] {
	return pulumix.Output[[]*QuoteSetV2]{
		OutputState: i.ToQuoteSetV2ArrayOutputWithContext(ctx).OutputState,
	}
}

// QuoteSetV2MapInput is an input type that accepts QuoteSetV2Map and QuoteSetV2MapOutput values.
// You can construct a concrete instance of `QuoteSetV2MapInput` via:
//
//	QuoteSetV2Map{ "key": QuoteSetV2Args{...} }
type QuoteSetV2MapInput interface {
	pulumi.Input

	ToQuoteSetV2MapOutput() QuoteSetV2MapOutput
	ToQuoteSetV2MapOutputWithContext(context.Context) QuoteSetV2MapOutput
}

type QuoteSetV2Map map[string]QuoteSetV2Input

func (QuoteSetV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuoteSetV2)(nil)).Elem()
}

func (i QuoteSetV2Map) ToQuoteSetV2MapOutput() QuoteSetV2MapOutput {
	return i.ToQuoteSetV2MapOutputWithContext(context.Background())
}

func (i QuoteSetV2Map) ToQuoteSetV2MapOutputWithContext(ctx context.Context) QuoteSetV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuoteSetV2MapOutput)
}

func (i QuoteSetV2Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*QuoteSetV2] {
	return pulumix.Output[map[string]*QuoteSetV2]{
		OutputState: i.ToQuoteSetV2MapOutputWithContext(ctx).OutputState,
	}
}

type QuoteSetV2Output struct{ *pulumi.OutputState }

func (QuoteSetV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**QuoteSetV2)(nil)).Elem()
}

func (o QuoteSetV2Output) ToQuoteSetV2Output() QuoteSetV2Output {
	return o
}

func (o QuoteSetV2Output) ToQuoteSetV2OutputWithContext(ctx context.Context) QuoteSetV2Output {
	return o
}

func (o QuoteSetV2Output) ToOutput(ctx context.Context) pulumix.Output[*QuoteSetV2] {
	return pulumix.Output[*QuoteSetV2]{
		OutputState: o.OutputState,
	}
}

// Quota value for backup gigabytes. Changing
// this updates the existing quotaset.
func (o QuoteSetV2Output) BackupGigabytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.BackupGigabytes }).(pulumi.IntOutput)
}

// Quota value for backups. Changing this updates the
// existing quotaset.
func (o QuoteSetV2Output) Backups() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.Backups }).(pulumi.IntOutput)
}

// Quota value for gigabytes. Changing this updates the
// existing quotaset.
func (o QuoteSetV2Output) Gigabytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.Gigabytes }).(pulumi.IntOutput)
}

// Quota value for groups. Changing this updates the
// existing quotaset.
func (o QuoteSetV2Output) Groups() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.Groups }).(pulumi.IntOutput)
}

// Quota value for gigabytes per volume .
// Changing this updates the existing quotaset.
func (o QuoteSetV2Output) PerVolumeGigabytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.PerVolumeGigabytes }).(pulumi.IntOutput)
}

// ID of the project to manage quotas. Changing this
// creates a new quotaset.
func (o QuoteSetV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to create the volume. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new quotaset.
func (o QuoteSetV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Quota value for snapshots. Changing this updates the
// existing quotaset.
func (o QuoteSetV2Output) Snapshots() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.Snapshots }).(pulumi.IntOutput)
}

// Key/Value pairs for setting quota for
// volumes types. Possible keys are `snapshots_<volume_type_name>`,
// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
func (o QuoteSetV2Output) VolumeTypeQuota() pulumi.MapOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.MapOutput { return v.VolumeTypeQuota }).(pulumi.MapOutput)
}

// Quota value for volumes. Changing this updates the
// existing quotaset.
func (o QuoteSetV2Output) Volumes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV2) pulumi.IntOutput { return v.Volumes }).(pulumi.IntOutput)
}

type QuoteSetV2ArrayOutput struct{ *pulumi.OutputState }

func (QuoteSetV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuoteSetV2)(nil)).Elem()
}

func (o QuoteSetV2ArrayOutput) ToQuoteSetV2ArrayOutput() QuoteSetV2ArrayOutput {
	return o
}

func (o QuoteSetV2ArrayOutput) ToQuoteSetV2ArrayOutputWithContext(ctx context.Context) QuoteSetV2ArrayOutput {
	return o
}

func (o QuoteSetV2ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*QuoteSetV2] {
	return pulumix.Output[[]*QuoteSetV2]{
		OutputState: o.OutputState,
	}
}

func (o QuoteSetV2ArrayOutput) Index(i pulumi.IntInput) QuoteSetV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuoteSetV2 {
		return vs[0].([]*QuoteSetV2)[vs[1].(int)]
	}).(QuoteSetV2Output)
}

type QuoteSetV2MapOutput struct{ *pulumi.OutputState }

func (QuoteSetV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuoteSetV2)(nil)).Elem()
}

func (o QuoteSetV2MapOutput) ToQuoteSetV2MapOutput() QuoteSetV2MapOutput {
	return o
}

func (o QuoteSetV2MapOutput) ToQuoteSetV2MapOutputWithContext(ctx context.Context) QuoteSetV2MapOutput {
	return o
}

func (o QuoteSetV2MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*QuoteSetV2] {
	return pulumix.Output[map[string]*QuoteSetV2]{
		OutputState: o.OutputState,
	}
}

func (o QuoteSetV2MapOutput) MapIndex(k pulumi.StringInput) QuoteSetV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuoteSetV2 {
		return vs[0].(map[string]*QuoteSetV2)[vs[1].(string)]
	}).(QuoteSetV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuoteSetV2Input)(nil)).Elem(), &QuoteSetV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuoteSetV2ArrayInput)(nil)).Elem(), QuoteSetV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuoteSetV2MapInput)(nil)).Elem(), QuoteSetV2Map{})
	pulumi.RegisterOutputType(QuoteSetV2Output{})
	pulumi.RegisterOutputType(QuoteSetV2ArrayOutput{})
	pulumi.RegisterOutputType(QuoteSetV2MapOutput{})
}
