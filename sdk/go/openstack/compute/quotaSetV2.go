// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 compute quotaset resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
//
//	in case of delete call.
//
// > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
//
//	created with zero value.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/compute"
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
//			_, err = compute.NewQuotaSetV2(ctx, "quotaset1", &compute.QuotaSetV2Args{
//				ProjectId:          project1.ID(),
//				KeyPairs:           pulumi.Int(10),
//				Ram:                pulumi.Int(40960),
//				Cores:              pulumi.Int(32),
//				Instances:          pulumi.Int(20),
//				ServerGroups:       pulumi.Int(4),
//				ServerGroupMembers: pulumi.Int(8),
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
// Quotasets can be imported using the `project_id/region_name`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:compute/quotaSetV2:QuotaSetV2 quotaset_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
//
// ```
type QuotaSetV2 struct {
	pulumi.CustomResourceState

	// Quota value for cores.
	// Changing this updates the existing quotaset.
	Cores pulumi.IntOutput `pulumi:"cores"`
	// Quota value for fixed IPs.
	// Changing this updates the existing quotaset.
	FixedIps pulumi.IntOutput `pulumi:"fixedIps"`
	// Quota value for floating IPs.
	// Changing this updates the existing quotaset.
	FloatingIps pulumi.IntOutput `pulumi:"floatingIps"`
	// Quota value for content bytes
	// of injected files. Changing this updates the existing quotaset.
	InjectedFileContentBytes pulumi.IntOutput `pulumi:"injectedFileContentBytes"`
	// Quota value for path bytes of
	// injected files. Changing this updates the existing quotaset.
	InjectedFilePathBytes pulumi.IntOutput `pulumi:"injectedFilePathBytes"`
	// Quota value for injected files.
	// Changing this updates the existing quotaset.
	InjectedFiles pulumi.IntOutput `pulumi:"injectedFiles"`
	// Quota value for instances.
	// Changing this updates the existing quotaset.
	Instances pulumi.IntOutput `pulumi:"instances"`
	// Quota value for key pairs.
	// Changing this updates the existing quotaset.
	KeyPairs pulumi.IntOutput `pulumi:"keyPairs"`
	// Quota value for metadata items.
	// Changing this updates the existing quotaset.
	MetadataItems pulumi.IntOutput `pulumi:"metadataItems"`
	// ID of the project to manage quotas.
	// Changing this creates a new quotaset.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Quota value for RAM.
	// Changing this updates the existing quotaset.
	Ram pulumi.IntOutput `pulumi:"ram"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringOutput `pulumi:"region"`
	// Quota value for security group rules.
	// Changing this updates the existing quotaset.
	SecurityGroupRules pulumi.IntOutput `pulumi:"securityGroupRules"`
	// Quota value for security groups.
	// Changing this updates the existing quotaset.
	SecurityGroups pulumi.IntOutput `pulumi:"securityGroups"`
	// Quota value for server groups members.
	// Changing this updates the existing quotaset.
	ServerGroupMembers pulumi.IntOutput `pulumi:"serverGroupMembers"`
	// Quota value for server groups.
	// Changing this updates the existing quotaset.
	ServerGroups pulumi.IntOutput `pulumi:"serverGroups"`
}

// NewQuotaSetV2 registers a new resource with the given unique name, arguments, and options.
func NewQuotaSetV2(ctx *pulumi.Context,
	name string, args *QuotaSetV2Args, opts ...pulumi.ResourceOption) (*QuotaSetV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource QuotaSetV2
	err := ctx.RegisterResource("openstack:compute/quotaSetV2:QuotaSetV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaSetV2 gets an existing QuotaSetV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaSetV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaSetV2State, opts ...pulumi.ResourceOption) (*QuotaSetV2, error) {
	var resource QuotaSetV2
	err := ctx.ReadResource("openstack:compute/quotaSetV2:QuotaSetV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaSetV2 resources.
type quotaSetV2State struct {
	// Quota value for cores.
	// Changing this updates the existing quotaset.
	Cores *int `pulumi:"cores"`
	// Quota value for fixed IPs.
	// Changing this updates the existing quotaset.
	FixedIps *int `pulumi:"fixedIps"`
	// Quota value for floating IPs.
	// Changing this updates the existing quotaset.
	FloatingIps *int `pulumi:"floatingIps"`
	// Quota value for content bytes
	// of injected files. Changing this updates the existing quotaset.
	InjectedFileContentBytes *int `pulumi:"injectedFileContentBytes"`
	// Quota value for path bytes of
	// injected files. Changing this updates the existing quotaset.
	InjectedFilePathBytes *int `pulumi:"injectedFilePathBytes"`
	// Quota value for injected files.
	// Changing this updates the existing quotaset.
	InjectedFiles *int `pulumi:"injectedFiles"`
	// Quota value for instances.
	// Changing this updates the existing quotaset.
	Instances *int `pulumi:"instances"`
	// Quota value for key pairs.
	// Changing this updates the existing quotaset.
	KeyPairs *int `pulumi:"keyPairs"`
	// Quota value for metadata items.
	// Changing this updates the existing quotaset.
	MetadataItems *int `pulumi:"metadataItems"`
	// ID of the project to manage quotas.
	// Changing this creates a new quotaset.
	ProjectId *string `pulumi:"projectId"`
	// Quota value for RAM.
	// Changing this updates the existing quotaset.
	Ram *int `pulumi:"ram"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
	// Quota value for security group rules.
	// Changing this updates the existing quotaset.
	SecurityGroupRules *int `pulumi:"securityGroupRules"`
	// Quota value for security groups.
	// Changing this updates the existing quotaset.
	SecurityGroups *int `pulumi:"securityGroups"`
	// Quota value for server groups members.
	// Changing this updates the existing quotaset.
	ServerGroupMembers *int `pulumi:"serverGroupMembers"`
	// Quota value for server groups.
	// Changing this updates the existing quotaset.
	ServerGroups *int `pulumi:"serverGroups"`
}

type QuotaSetV2State struct {
	// Quota value for cores.
	// Changing this updates the existing quotaset.
	Cores pulumi.IntPtrInput
	// Quota value for fixed IPs.
	// Changing this updates the existing quotaset.
	FixedIps pulumi.IntPtrInput
	// Quota value for floating IPs.
	// Changing this updates the existing quotaset.
	FloatingIps pulumi.IntPtrInput
	// Quota value for content bytes
	// of injected files. Changing this updates the existing quotaset.
	InjectedFileContentBytes pulumi.IntPtrInput
	// Quota value for path bytes of
	// injected files. Changing this updates the existing quotaset.
	InjectedFilePathBytes pulumi.IntPtrInput
	// Quota value for injected files.
	// Changing this updates the existing quotaset.
	InjectedFiles pulumi.IntPtrInput
	// Quota value for instances.
	// Changing this updates the existing quotaset.
	Instances pulumi.IntPtrInput
	// Quota value for key pairs.
	// Changing this updates the existing quotaset.
	KeyPairs pulumi.IntPtrInput
	// Quota value for metadata items.
	// Changing this updates the existing quotaset.
	MetadataItems pulumi.IntPtrInput
	// ID of the project to manage quotas.
	// Changing this creates a new quotaset.
	ProjectId pulumi.StringPtrInput
	// Quota value for RAM.
	// Changing this updates the existing quotaset.
	Ram pulumi.IntPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
	// Quota value for security group rules.
	// Changing this updates the existing quotaset.
	SecurityGroupRules pulumi.IntPtrInput
	// Quota value for security groups.
	// Changing this updates the existing quotaset.
	SecurityGroups pulumi.IntPtrInput
	// Quota value for server groups members.
	// Changing this updates the existing quotaset.
	ServerGroupMembers pulumi.IntPtrInput
	// Quota value for server groups.
	// Changing this updates the existing quotaset.
	ServerGroups pulumi.IntPtrInput
}

func (QuotaSetV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaSetV2State)(nil)).Elem()
}

type quotaSetV2Args struct {
	// Quota value for cores.
	// Changing this updates the existing quotaset.
	Cores *int `pulumi:"cores"`
	// Quota value for fixed IPs.
	// Changing this updates the existing quotaset.
	FixedIps *int `pulumi:"fixedIps"`
	// Quota value for floating IPs.
	// Changing this updates the existing quotaset.
	FloatingIps *int `pulumi:"floatingIps"`
	// Quota value for content bytes
	// of injected files. Changing this updates the existing quotaset.
	InjectedFileContentBytes *int `pulumi:"injectedFileContentBytes"`
	// Quota value for path bytes of
	// injected files. Changing this updates the existing quotaset.
	InjectedFilePathBytes *int `pulumi:"injectedFilePathBytes"`
	// Quota value for injected files.
	// Changing this updates the existing quotaset.
	InjectedFiles *int `pulumi:"injectedFiles"`
	// Quota value for instances.
	// Changing this updates the existing quotaset.
	Instances *int `pulumi:"instances"`
	// Quota value for key pairs.
	// Changing this updates the existing quotaset.
	KeyPairs *int `pulumi:"keyPairs"`
	// Quota value for metadata items.
	// Changing this updates the existing quotaset.
	MetadataItems *int `pulumi:"metadataItems"`
	// ID of the project to manage quotas.
	// Changing this creates a new quotaset.
	ProjectId string `pulumi:"projectId"`
	// Quota value for RAM.
	// Changing this updates the existing quotaset.
	Ram *int `pulumi:"ram"`
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `pulumi:"region"`
	// Quota value for security group rules.
	// Changing this updates the existing quotaset.
	SecurityGroupRules *int `pulumi:"securityGroupRules"`
	// Quota value for security groups.
	// Changing this updates the existing quotaset.
	SecurityGroups *int `pulumi:"securityGroups"`
	// Quota value for server groups members.
	// Changing this updates the existing quotaset.
	ServerGroupMembers *int `pulumi:"serverGroupMembers"`
	// Quota value for server groups.
	// Changing this updates the existing quotaset.
	ServerGroups *int `pulumi:"serverGroups"`
}

// The set of arguments for constructing a QuotaSetV2 resource.
type QuotaSetV2Args struct {
	// Quota value for cores.
	// Changing this updates the existing quotaset.
	Cores pulumi.IntPtrInput
	// Quota value for fixed IPs.
	// Changing this updates the existing quotaset.
	FixedIps pulumi.IntPtrInput
	// Quota value for floating IPs.
	// Changing this updates the existing quotaset.
	FloatingIps pulumi.IntPtrInput
	// Quota value for content bytes
	// of injected files. Changing this updates the existing quotaset.
	InjectedFileContentBytes pulumi.IntPtrInput
	// Quota value for path bytes of
	// injected files. Changing this updates the existing quotaset.
	InjectedFilePathBytes pulumi.IntPtrInput
	// Quota value for injected files.
	// Changing this updates the existing quotaset.
	InjectedFiles pulumi.IntPtrInput
	// Quota value for instances.
	// Changing this updates the existing quotaset.
	Instances pulumi.IntPtrInput
	// Quota value for key pairs.
	// Changing this updates the existing quotaset.
	KeyPairs pulumi.IntPtrInput
	// Quota value for metadata items.
	// Changing this updates the existing quotaset.
	MetadataItems pulumi.IntPtrInput
	// ID of the project to manage quotas.
	// Changing this creates a new quotaset.
	ProjectId pulumi.StringInput
	// Quota value for RAM.
	// Changing this updates the existing quotaset.
	Ram pulumi.IntPtrInput
	// The region in which to create the volume. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new quotaset.
	Region pulumi.StringPtrInput
	// Quota value for security group rules.
	// Changing this updates the existing quotaset.
	SecurityGroupRules pulumi.IntPtrInput
	// Quota value for security groups.
	// Changing this updates the existing quotaset.
	SecurityGroups pulumi.IntPtrInput
	// Quota value for server groups members.
	// Changing this updates the existing quotaset.
	ServerGroupMembers pulumi.IntPtrInput
	// Quota value for server groups.
	// Changing this updates the existing quotaset.
	ServerGroups pulumi.IntPtrInput
}

func (QuotaSetV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaSetV2Args)(nil)).Elem()
}

type QuotaSetV2Input interface {
	pulumi.Input

	ToQuotaSetV2Output() QuotaSetV2Output
	ToQuotaSetV2OutputWithContext(ctx context.Context) QuotaSetV2Output
}

func (*QuotaSetV2) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaSetV2)(nil)).Elem()
}

func (i *QuotaSetV2) ToQuotaSetV2Output() QuotaSetV2Output {
	return i.ToQuotaSetV2OutputWithContext(context.Background())
}

func (i *QuotaSetV2) ToQuotaSetV2OutputWithContext(ctx context.Context) QuotaSetV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaSetV2Output)
}

// QuotaSetV2ArrayInput is an input type that accepts QuotaSetV2Array and QuotaSetV2ArrayOutput values.
// You can construct a concrete instance of `QuotaSetV2ArrayInput` via:
//
//	QuotaSetV2Array{ QuotaSetV2Args{...} }
type QuotaSetV2ArrayInput interface {
	pulumi.Input

	ToQuotaSetV2ArrayOutput() QuotaSetV2ArrayOutput
	ToQuotaSetV2ArrayOutputWithContext(context.Context) QuotaSetV2ArrayOutput
}

type QuotaSetV2Array []QuotaSetV2Input

func (QuotaSetV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaSetV2)(nil)).Elem()
}

func (i QuotaSetV2Array) ToQuotaSetV2ArrayOutput() QuotaSetV2ArrayOutput {
	return i.ToQuotaSetV2ArrayOutputWithContext(context.Background())
}

func (i QuotaSetV2Array) ToQuotaSetV2ArrayOutputWithContext(ctx context.Context) QuotaSetV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaSetV2ArrayOutput)
}

// QuotaSetV2MapInput is an input type that accepts QuotaSetV2Map and QuotaSetV2MapOutput values.
// You can construct a concrete instance of `QuotaSetV2MapInput` via:
//
//	QuotaSetV2Map{ "key": QuotaSetV2Args{...} }
type QuotaSetV2MapInput interface {
	pulumi.Input

	ToQuotaSetV2MapOutput() QuotaSetV2MapOutput
	ToQuotaSetV2MapOutputWithContext(context.Context) QuotaSetV2MapOutput
}

type QuotaSetV2Map map[string]QuotaSetV2Input

func (QuotaSetV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaSetV2)(nil)).Elem()
}

func (i QuotaSetV2Map) ToQuotaSetV2MapOutput() QuotaSetV2MapOutput {
	return i.ToQuotaSetV2MapOutputWithContext(context.Background())
}

func (i QuotaSetV2Map) ToQuotaSetV2MapOutputWithContext(ctx context.Context) QuotaSetV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaSetV2MapOutput)
}

type QuotaSetV2Output struct{ *pulumi.OutputState }

func (QuotaSetV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaSetV2)(nil)).Elem()
}

func (o QuotaSetV2Output) ToQuotaSetV2Output() QuotaSetV2Output {
	return o
}

func (o QuotaSetV2Output) ToQuotaSetV2OutputWithContext(ctx context.Context) QuotaSetV2Output {
	return o
}

// Quota value for cores.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.Cores }).(pulumi.IntOutput)
}

// Quota value for fixed IPs.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) FixedIps() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.FixedIps }).(pulumi.IntOutput)
}

// Quota value for floating IPs.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) FloatingIps() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.FloatingIps }).(pulumi.IntOutput)
}

// Quota value for content bytes
// of injected files. Changing this updates the existing quotaset.
func (o QuotaSetV2Output) InjectedFileContentBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.InjectedFileContentBytes }).(pulumi.IntOutput)
}

// Quota value for path bytes of
// injected files. Changing this updates the existing quotaset.
func (o QuotaSetV2Output) InjectedFilePathBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.InjectedFilePathBytes }).(pulumi.IntOutput)
}

// Quota value for injected files.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) InjectedFiles() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.InjectedFiles }).(pulumi.IntOutput)
}

// Quota value for instances.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) Instances() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.Instances }).(pulumi.IntOutput)
}

// Quota value for key pairs.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) KeyPairs() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.KeyPairs }).(pulumi.IntOutput)
}

// Quota value for metadata items.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) MetadataItems() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.MetadataItems }).(pulumi.IntOutput)
}

// ID of the project to manage quotas.
// Changing this creates a new quotaset.
func (o QuotaSetV2Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Quota value for RAM.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.Ram }).(pulumi.IntOutput)
}

// The region in which to create the volume. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new quotaset.
func (o QuotaSetV2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Quota value for security group rules.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) SecurityGroupRules() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.SecurityGroupRules }).(pulumi.IntOutput)
}

// Quota value for security groups.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) SecurityGroups() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.SecurityGroups }).(pulumi.IntOutput)
}

// Quota value for server groups members.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) ServerGroupMembers() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.ServerGroupMembers }).(pulumi.IntOutput)
}

// Quota value for server groups.
// Changing this updates the existing quotaset.
func (o QuotaSetV2Output) ServerGroups() pulumi.IntOutput {
	return o.ApplyT(func(v *QuotaSetV2) pulumi.IntOutput { return v.ServerGroups }).(pulumi.IntOutput)
}

type QuotaSetV2ArrayOutput struct{ *pulumi.OutputState }

func (QuotaSetV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuotaSetV2)(nil)).Elem()
}

func (o QuotaSetV2ArrayOutput) ToQuotaSetV2ArrayOutput() QuotaSetV2ArrayOutput {
	return o
}

func (o QuotaSetV2ArrayOutput) ToQuotaSetV2ArrayOutputWithContext(ctx context.Context) QuotaSetV2ArrayOutput {
	return o
}

func (o QuotaSetV2ArrayOutput) Index(i pulumi.IntInput) QuotaSetV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuotaSetV2 {
		return vs[0].([]*QuotaSetV2)[vs[1].(int)]
	}).(QuotaSetV2Output)
}

type QuotaSetV2MapOutput struct{ *pulumi.OutputState }

func (QuotaSetV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuotaSetV2)(nil)).Elem()
}

func (o QuotaSetV2MapOutput) ToQuotaSetV2MapOutput() QuotaSetV2MapOutput {
	return o
}

func (o QuotaSetV2MapOutput) ToQuotaSetV2MapOutputWithContext(ctx context.Context) QuotaSetV2MapOutput {
	return o
}

func (o QuotaSetV2MapOutput) MapIndex(k pulumi.StringInput) QuotaSetV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuotaSetV2 {
		return vs[0].(map[string]*QuotaSetV2)[vs[1].(string)]
	}).(QuotaSetV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaSetV2Input)(nil)).Elem(), &QuotaSetV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaSetV2ArrayInput)(nil)).Elem(), QuotaSetV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuotaSetV2MapInput)(nil)).Elem(), QuotaSetV2Map{})
	pulumi.RegisterOutputType(QuotaSetV2Output{})
	pulumi.RegisterOutputType(QuotaSetV2ArrayOutput{})
	pulumi.RegisterOutputType(QuotaSetV2MapOutput{})
}
