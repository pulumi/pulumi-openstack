// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 block storage quotaset resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
//     in case of delete call.
//
// > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
//     created with zero value. This excludes volume type quota.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/blockstorage"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project1, err := identity.NewProject(ctx, "project1", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = blockstorage.NewQuoteSetV3(ctx, "quotaset1", &blockstorage.QuoteSetV3Args{
// 			ProjectId:          project1.ID(),
// 			Volumes:            pulumi.Int(10),
// 			Snapshots:          pulumi.Int(4),
// 			Gigabytes:          pulumi.Int(100),
// 			PerVolumeGigabytes: pulumi.Int(10),
// 			Backups:            pulumi.Int(4),
// 			BackupGigabytes:    pulumi.Int(10),
// 			Groups:             pulumi.Int(100),
// 			VolumeTypeQuota: pulumi.AnyMap{
// 				"volumes_ssd":   pulumi.Any(30),
// 				"gigabytes_ssd": pulumi.Any(500),
// 				"snapshots_ssd": pulumi.Any(10),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Quotasets can be imported using the `project_id/region`, e.g.
//
// ```sh
//  $ pulumi import openstack:blockstorage/quoteSetV3:QuoteSetV3 quotaset_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
// ```
type QuoteSetV3 struct {
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

// NewQuoteSetV3 registers a new resource with the given unique name, arguments, and options.
func NewQuoteSetV3(ctx *pulumi.Context,
	name string, args *QuoteSetV3Args, opts ...pulumi.ResourceOption) (*QuoteSetV3, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource QuoteSetV3
	err := ctx.RegisterResource("openstack:blockstorage/quoteSetV3:QuoteSetV3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuoteSetV3 gets an existing QuoteSetV3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuoteSetV3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuoteSetV3State, opts ...pulumi.ResourceOption) (*QuoteSetV3, error) {
	var resource QuoteSetV3
	err := ctx.ReadResource("openstack:blockstorage/quoteSetV3:QuoteSetV3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuoteSetV3 resources.
type quoteSetV3State struct {
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

type QuoteSetV3State struct {
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

func (QuoteSetV3State) ElementType() reflect.Type {
	return reflect.TypeOf((*quoteSetV3State)(nil)).Elem()
}

type quoteSetV3Args struct {
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

// The set of arguments for constructing a QuoteSetV3 resource.
type QuoteSetV3Args struct {
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

func (QuoteSetV3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*quoteSetV3Args)(nil)).Elem()
}

type QuoteSetV3Input interface {
	pulumi.Input

	ToQuoteSetV3Output() QuoteSetV3Output
	ToQuoteSetV3OutputWithContext(ctx context.Context) QuoteSetV3Output
}

func (*QuoteSetV3) ElementType() reflect.Type {
	return reflect.TypeOf((**QuoteSetV3)(nil)).Elem()
}

func (i *QuoteSetV3) ToQuoteSetV3Output() QuoteSetV3Output {
	return i.ToQuoteSetV3OutputWithContext(context.Background())
}

func (i *QuoteSetV3) ToQuoteSetV3OutputWithContext(ctx context.Context) QuoteSetV3Output {
	return pulumi.ToOutputWithContext(ctx, i).(QuoteSetV3Output)
}

// QuoteSetV3ArrayInput is an input type that accepts QuoteSetV3Array and QuoteSetV3ArrayOutput values.
// You can construct a concrete instance of `QuoteSetV3ArrayInput` via:
//
//          QuoteSetV3Array{ QuoteSetV3Args{...} }
type QuoteSetV3ArrayInput interface {
	pulumi.Input

	ToQuoteSetV3ArrayOutput() QuoteSetV3ArrayOutput
	ToQuoteSetV3ArrayOutputWithContext(context.Context) QuoteSetV3ArrayOutput
}

type QuoteSetV3Array []QuoteSetV3Input

func (QuoteSetV3Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuoteSetV3)(nil)).Elem()
}

func (i QuoteSetV3Array) ToQuoteSetV3ArrayOutput() QuoteSetV3ArrayOutput {
	return i.ToQuoteSetV3ArrayOutputWithContext(context.Background())
}

func (i QuoteSetV3Array) ToQuoteSetV3ArrayOutputWithContext(ctx context.Context) QuoteSetV3ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuoteSetV3ArrayOutput)
}

// QuoteSetV3MapInput is an input type that accepts QuoteSetV3Map and QuoteSetV3MapOutput values.
// You can construct a concrete instance of `QuoteSetV3MapInput` via:
//
//          QuoteSetV3Map{ "key": QuoteSetV3Args{...} }
type QuoteSetV3MapInput interface {
	pulumi.Input

	ToQuoteSetV3MapOutput() QuoteSetV3MapOutput
	ToQuoteSetV3MapOutputWithContext(context.Context) QuoteSetV3MapOutput
}

type QuoteSetV3Map map[string]QuoteSetV3Input

func (QuoteSetV3Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuoteSetV3)(nil)).Elem()
}

func (i QuoteSetV3Map) ToQuoteSetV3MapOutput() QuoteSetV3MapOutput {
	return i.ToQuoteSetV3MapOutputWithContext(context.Background())
}

func (i QuoteSetV3Map) ToQuoteSetV3MapOutputWithContext(ctx context.Context) QuoteSetV3MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuoteSetV3MapOutput)
}

type QuoteSetV3Output struct{ *pulumi.OutputState }

func (QuoteSetV3Output) ElementType() reflect.Type {
	return reflect.TypeOf((**QuoteSetV3)(nil)).Elem()
}

func (o QuoteSetV3Output) ToQuoteSetV3Output() QuoteSetV3Output {
	return o
}

func (o QuoteSetV3Output) ToQuoteSetV3OutputWithContext(ctx context.Context) QuoteSetV3Output {
	return o
}

// Quota value for backup gigabytes. Changing
// this updates the existing quotaset.
func (o QuoteSetV3Output) BackupGigabytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.BackupGigabytes }).(pulumi.IntOutput)
}

// Quota value for backups. Changing this updates the
// existing quotaset.
func (o QuoteSetV3Output) Backups() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.Backups }).(pulumi.IntOutput)
}

// Quota value for gigabytes. Changing this updates the
// existing quotaset.
func (o QuoteSetV3Output) Gigabytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.Gigabytes }).(pulumi.IntOutput)
}

// Quota value for groups. Changing this updates the
// existing quotaset.
func (o QuoteSetV3Output) Groups() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.Groups }).(pulumi.IntOutput)
}

// Quota value for gigabytes per volume .
// Changing this updates the existing quotaset.
func (o QuoteSetV3Output) PerVolumeGigabytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.PerVolumeGigabytes }).(pulumi.IntOutput)
}

// ID of the project to manage quotas. Changing this
// creates a new quotaset.
func (o QuoteSetV3Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to create the volume. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new quotaset.
func (o QuoteSetV3Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Quota value for snapshots. Changing this updates the
// existing quotaset.
func (o QuoteSetV3Output) Snapshots() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.Snapshots }).(pulumi.IntOutput)
}

// Key/Value pairs for setting quota for
// volumes types. Possible keys are `snapshots_<volume_type_name>`,
// `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
func (o QuoteSetV3Output) VolumeTypeQuota() pulumi.MapOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.MapOutput { return v.VolumeTypeQuota }).(pulumi.MapOutput)
}

// Quota value for volumes. Changing this updates the
// existing quotaset.
func (o QuoteSetV3Output) Volumes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuoteSetV3) pulumi.IntOutput { return v.Volumes }).(pulumi.IntOutput)
}

type QuoteSetV3ArrayOutput struct{ *pulumi.OutputState }

func (QuoteSetV3ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuoteSetV3)(nil)).Elem()
}

func (o QuoteSetV3ArrayOutput) ToQuoteSetV3ArrayOutput() QuoteSetV3ArrayOutput {
	return o
}

func (o QuoteSetV3ArrayOutput) ToQuoteSetV3ArrayOutputWithContext(ctx context.Context) QuoteSetV3ArrayOutput {
	return o
}

func (o QuoteSetV3ArrayOutput) Index(i pulumi.IntInput) QuoteSetV3Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuoteSetV3 {
		return vs[0].([]*QuoteSetV3)[vs[1].(int)]
	}).(QuoteSetV3Output)
}

type QuoteSetV3MapOutput struct{ *pulumi.OutputState }

func (QuoteSetV3MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuoteSetV3)(nil)).Elem()
}

func (o QuoteSetV3MapOutput) ToQuoteSetV3MapOutput() QuoteSetV3MapOutput {
	return o
}

func (o QuoteSetV3MapOutput) ToQuoteSetV3MapOutputWithContext(ctx context.Context) QuoteSetV3MapOutput {
	return o
}

func (o QuoteSetV3MapOutput) MapIndex(k pulumi.StringInput) QuoteSetV3Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuoteSetV3 {
		return vs[0].(map[string]*QuoteSetV3)[vs[1].(string)]
	}).(QuoteSetV3Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuoteSetV3Input)(nil)).Elem(), &QuoteSetV3{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuoteSetV3ArrayInput)(nil)).Elem(), QuoteSetV3Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuoteSetV3MapInput)(nil)).Elem(), QuoteSetV3Map{})
	pulumi.RegisterOutputType(QuoteSetV3Output{})
	pulumi.RegisterOutputType(QuoteSetV3ArrayOutput{})
	pulumi.RegisterOutputType(QuoteSetV3MapOutput{})
}
