// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 networking quota resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
//     in case of delete call.
//
// > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
//     created with zero value.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project1, err := identity.NewProject(ctx, "project1", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = networking.NewQuotaV2(ctx, "quota1", &networking.QuotaV2Args{
// 			ProjectId:         project1.ID(),
// 			Floatingip:        pulumi.Int(10),
// 			Network:           pulumi.Int(4),
// 			Port:              pulumi.Int(100),
// 			RbacPolicy:        pulumi.Int(10),
// 			Router:            pulumi.Int(4),
// 			SecurityGroup:     pulumi.Int(10),
// 			SecurityGroupRule: pulumi.Int(100),
// 			Subnet:            pulumi.Int(8),
// 			Subnetpool:        pulumi.Int(2),
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
// Quotas can be imported using the `project_id/region_name`, e.g.
//
// ```sh
//  $ pulumi import openstack:networking/quotaV2:QuotaV2 quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
// ```
type QuotaV2 struct {
	pulumi.CustomResourceState

	// Quota value for floating IPs. Changing this updates the
	// existing quota.
	Floatingip pulumi.IntOutput `pulumi:"floatingip"`
	// Quota value for networks. Changing this updates the
	// existing quota.
	Network pulumi.IntOutput `pulumi:"network"`
	// Quota value for ports. Changing this updates the
	// existing quota.
	Port pulumi.IntOutput `pulumi:"port"`
	// ID of the project to manage quota. Changing this
	// creates new quota.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Quota value for RBAC policies.
	// Changing this updates the existing quota.
	RbacPolicy pulumi.IntOutput `pulumi:"rbacPolicy"`
	// The region in which to create the quota. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates new quota.
	Region pulumi.StringOutput `pulumi:"region"`
	// Quota value for routers. Changing this updates the
	// existing quota.
	Router pulumi.IntOutput `pulumi:"router"`
	// Quota value for security groups. Changing
	// this updates the existing quota.
	SecurityGroup pulumi.IntOutput `pulumi:"securityGroup"`
	// Quota value for security group rules.
	// Changing this updates the existing quota.
	SecurityGroupRule pulumi.IntOutput `pulumi:"securityGroupRule"`
	// Quota value for subnets. Changing
	// this updates the existing quota.
	Subnet pulumi.IntOutput `pulumi:"subnet"`
	// Quota value for subnetpools.
	// Changing this updates the existing quota.
	Subnetpool pulumi.IntOutput `pulumi:"subnetpool"`
}

// NewQuotaV2 registers a new resource with the given unique name, arguments, and options.
func NewQuotaV2(ctx *pulumi.Context,
	name string, args *QuotaV2Args, opts ...pulumi.ResourceOption) (*QuotaV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource QuotaV2
	err := ctx.RegisterResource("openstack:networking/quotaV2:QuotaV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaV2 gets an existing QuotaV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaV2State, opts ...pulumi.ResourceOption) (*QuotaV2, error) {
	var resource QuotaV2
	err := ctx.ReadResource("openstack:networking/quotaV2:QuotaV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaV2 resources.
type quotaV2State struct {
	// Quota value for floating IPs. Changing this updates the
	// existing quota.
	Floatingip *int `pulumi:"floatingip"`
	// Quota value for networks. Changing this updates the
	// existing quota.
	Network *int `pulumi:"network"`
	// Quota value for ports. Changing this updates the
	// existing quota.
	Port *int `pulumi:"port"`
	// ID of the project to manage quota. Changing this
	// creates new quota.
	ProjectId *string `pulumi:"projectId"`
	// Quota value for RBAC policies.
	// Changing this updates the existing quota.
	RbacPolicy *int `pulumi:"rbacPolicy"`
	// The region in which to create the quota. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates new quota.
	Region *string `pulumi:"region"`
	// Quota value for routers. Changing this updates the
	// existing quota.
	Router *int `pulumi:"router"`
	// Quota value for security groups. Changing
	// this updates the existing quota.
	SecurityGroup *int `pulumi:"securityGroup"`
	// Quota value for security group rules.
	// Changing this updates the existing quota.
	SecurityGroupRule *int `pulumi:"securityGroupRule"`
	// Quota value for subnets. Changing
	// this updates the existing quota.
	Subnet *int `pulumi:"subnet"`
	// Quota value for subnetpools.
	// Changing this updates the existing quota.
	Subnetpool *int `pulumi:"subnetpool"`
}

type QuotaV2State struct {
	// Quota value for floating IPs. Changing this updates the
	// existing quota.
	Floatingip pulumi.IntPtrInput
	// Quota value for networks. Changing this updates the
	// existing quota.
	Network pulumi.IntPtrInput
	// Quota value for ports. Changing this updates the
	// existing quota.
	Port pulumi.IntPtrInput
	// ID of the project to manage quota. Changing this
	// creates new quota.
	ProjectId pulumi.StringPtrInput
	// Quota value for RBAC policies.
	// Changing this updates the existing quota.
	RbacPolicy pulumi.IntPtrInput
	// The region in which to create the quota. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates new quota.
	Region pulumi.StringPtrInput
	// Quota value for routers. Changing this updates the
	// existing quota.
	Router pulumi.IntPtrInput
	// Quota value for security groups. Changing
	// this updates the existing quota.
	SecurityGroup pulumi.IntPtrInput
	// Quota value for security group rules.
	// Changing this updates the existing quota.
	SecurityGroupRule pulumi.IntPtrInput
	// Quota value for subnets. Changing
	// this updates the existing quota.
	Subnet pulumi.IntPtrInput
	// Quota value for subnetpools.
	// Changing this updates the existing quota.
	Subnetpool pulumi.IntPtrInput
}

func (QuotaV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaV2State)(nil)).Elem()
}

type quotaV2Args struct {
	// Quota value for floating IPs. Changing this updates the
	// existing quota.
	Floatingip *int `pulumi:"floatingip"`
	// Quota value for networks. Changing this updates the
	// existing quota.
	Network *int `pulumi:"network"`
	// Quota value for ports. Changing this updates the
	// existing quota.
	Port *int `pulumi:"port"`
	// ID of the project to manage quota. Changing this
	// creates new quota.
	ProjectId string `pulumi:"projectId"`
	// Quota value for RBAC policies.
	// Changing this updates the existing quota.
	RbacPolicy *int `pulumi:"rbacPolicy"`
	// The region in which to create the quota. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates new quota.
	Region *string `pulumi:"region"`
	// Quota value for routers. Changing this updates the
	// existing quota.
	Router *int `pulumi:"router"`
	// Quota value for security groups. Changing
	// this updates the existing quota.
	SecurityGroup *int `pulumi:"securityGroup"`
	// Quota value for security group rules.
	// Changing this updates the existing quota.
	SecurityGroupRule *int `pulumi:"securityGroupRule"`
	// Quota value for subnets. Changing
	// this updates the existing quota.
	Subnet *int `pulumi:"subnet"`
	// Quota value for subnetpools.
	// Changing this updates the existing quota.
	Subnetpool *int `pulumi:"subnetpool"`
}

// The set of arguments for constructing a QuotaV2 resource.
type QuotaV2Args struct {
	// Quota value for floating IPs. Changing this updates the
	// existing quota.
	Floatingip pulumi.IntPtrInput
	// Quota value for networks. Changing this updates the
	// existing quota.
	Network pulumi.IntPtrInput
	// Quota value for ports. Changing this updates the
	// existing quota.
	Port pulumi.IntPtrInput
	// ID of the project to manage quota. Changing this
	// creates new quota.
	ProjectId pulumi.StringInput
	// Quota value for RBAC policies.
	// Changing this updates the existing quota.
	RbacPolicy pulumi.IntPtrInput
	// The region in which to create the quota. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates new quota.
	Region pulumi.StringPtrInput
	// Quota value for routers. Changing this updates the
	// existing quota.
	Router pulumi.IntPtrInput
	// Quota value for security groups. Changing
	// this updates the existing quota.
	SecurityGroup pulumi.IntPtrInput
	// Quota value for security group rules.
	// Changing this updates the existing quota.
	SecurityGroupRule pulumi.IntPtrInput
	// Quota value for subnets. Changing
	// this updates the existing quota.
	Subnet pulumi.IntPtrInput
	// Quota value for subnetpools.
	// Changing this updates the existing quota.
	Subnetpool pulumi.IntPtrInput
}

func (QuotaV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaV2Args)(nil)).Elem()
}

type QuotaV2Input interface {
	pulumi.Input

	ToQuotaV2Output() QuotaV2Output
	ToQuotaV2OutputWithContext(ctx context.Context) QuotaV2Output
}

func (*QuotaV2) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaV2)(nil))
}

func (i *QuotaV2) ToQuotaV2Output() QuotaV2Output {
	return i.ToQuotaV2OutputWithContext(context.Background())
}

func (i *QuotaV2) ToQuotaV2OutputWithContext(ctx context.Context) QuotaV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaV2Output)
}

func (i *QuotaV2) ToQuotaV2PtrOutput() QuotaV2PtrOutput {
	return i.ToQuotaV2PtrOutputWithContext(context.Background())
}

func (i *QuotaV2) ToQuotaV2PtrOutputWithContext(ctx context.Context) QuotaV2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaV2PtrOutput)
}

type QuotaV2PtrInput interface {
	pulumi.Input

	ToQuotaV2PtrOutput() QuotaV2PtrOutput
	ToQuotaV2PtrOutputWithContext(ctx context.Context) QuotaV2PtrOutput
}

type quotaV2PtrType QuotaV2Args

func (*quotaV2PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaV2)(nil))
}

func (i *quotaV2PtrType) ToQuotaV2PtrOutput() QuotaV2PtrOutput {
	return i.ToQuotaV2PtrOutputWithContext(context.Background())
}

func (i *quotaV2PtrType) ToQuotaV2PtrOutputWithContext(ctx context.Context) QuotaV2PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaV2PtrOutput)
}

// QuotaV2ArrayInput is an input type that accepts QuotaV2Array and QuotaV2ArrayOutput values.
// You can construct a concrete instance of `QuotaV2ArrayInput` via:
//
//          QuotaV2Array{ QuotaV2Args{...} }
type QuotaV2ArrayInput interface {
	pulumi.Input

	ToQuotaV2ArrayOutput() QuotaV2ArrayOutput
	ToQuotaV2ArrayOutputWithContext(context.Context) QuotaV2ArrayOutput
}

type QuotaV2Array []QuotaV2Input

func (QuotaV2Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*QuotaV2)(nil))
}

func (i QuotaV2Array) ToQuotaV2ArrayOutput() QuotaV2ArrayOutput {
	return i.ToQuotaV2ArrayOutputWithContext(context.Background())
}

func (i QuotaV2Array) ToQuotaV2ArrayOutputWithContext(ctx context.Context) QuotaV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaV2ArrayOutput)
}

// QuotaV2MapInput is an input type that accepts QuotaV2Map and QuotaV2MapOutput values.
// You can construct a concrete instance of `QuotaV2MapInput` via:
//
//          QuotaV2Map{ "key": QuotaV2Args{...} }
type QuotaV2MapInput interface {
	pulumi.Input

	ToQuotaV2MapOutput() QuotaV2MapOutput
	ToQuotaV2MapOutputWithContext(context.Context) QuotaV2MapOutput
}

type QuotaV2Map map[string]QuotaV2Input

func (QuotaV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*QuotaV2)(nil))
}

func (i QuotaV2Map) ToQuotaV2MapOutput() QuotaV2MapOutput {
	return i.ToQuotaV2MapOutputWithContext(context.Background())
}

func (i QuotaV2Map) ToQuotaV2MapOutputWithContext(ctx context.Context) QuotaV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaV2MapOutput)
}

type QuotaV2Output struct {
	*pulumi.OutputState
}

func (QuotaV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaV2)(nil))
}

func (o QuotaV2Output) ToQuotaV2Output() QuotaV2Output {
	return o
}

func (o QuotaV2Output) ToQuotaV2OutputWithContext(ctx context.Context) QuotaV2Output {
	return o
}

func (o QuotaV2Output) ToQuotaV2PtrOutput() QuotaV2PtrOutput {
	return o.ToQuotaV2PtrOutputWithContext(context.Background())
}

func (o QuotaV2Output) ToQuotaV2PtrOutputWithContext(ctx context.Context) QuotaV2PtrOutput {
	return o.ApplyT(func(v QuotaV2) *QuotaV2 {
		return &v
	}).(QuotaV2PtrOutput)
}

type QuotaV2PtrOutput struct {
	*pulumi.OutputState
}

func (QuotaV2PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaV2)(nil))
}

func (o QuotaV2PtrOutput) ToQuotaV2PtrOutput() QuotaV2PtrOutput {
	return o
}

func (o QuotaV2PtrOutput) ToQuotaV2PtrOutputWithContext(ctx context.Context) QuotaV2PtrOutput {
	return o
}

type QuotaV2ArrayOutput struct{ *pulumi.OutputState }

func (QuotaV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuotaV2)(nil))
}

func (o QuotaV2ArrayOutput) ToQuotaV2ArrayOutput() QuotaV2ArrayOutput {
	return o
}

func (o QuotaV2ArrayOutput) ToQuotaV2ArrayOutputWithContext(ctx context.Context) QuotaV2ArrayOutput {
	return o
}

func (o QuotaV2ArrayOutput) Index(i pulumi.IntInput) QuotaV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QuotaV2 {
		return vs[0].([]QuotaV2)[vs[1].(int)]
	}).(QuotaV2Output)
}

type QuotaV2MapOutput struct{ *pulumi.OutputState }

func (QuotaV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QuotaV2)(nil))
}

func (o QuotaV2MapOutput) ToQuotaV2MapOutput() QuotaV2MapOutput {
	return o
}

func (o QuotaV2MapOutput) ToQuotaV2MapOutputWithContext(ctx context.Context) QuotaV2MapOutput {
	return o
}

func (o QuotaV2MapOutput) MapIndex(k pulumi.StringInput) QuotaV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QuotaV2 {
		return vs[0].(map[string]QuotaV2)[vs[1].(string)]
	}).(QuotaV2Output)
}

func init() {
	pulumi.RegisterOutputType(QuotaV2Output{})
	pulumi.RegisterOutputType(QuotaV2PtrOutput{})
	pulumi.RegisterOutputType(QuotaV2ArrayOutput{})
	pulumi.RegisterOutputType(QuotaV2MapOutput{})
}
