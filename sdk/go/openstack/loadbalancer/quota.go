// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 load balancer quota resource within OpenStack.
//
// > **Note:** This usually requires admin privileges.
//
// > **Note:** This resource is only available for Octavia.
//
// > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack
//    API in case of delete call.
//
// > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
//    created with zero value.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/loadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		project1, err := identity.NewProject(ctx, "project1", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = loadbalancer.NewQuota(ctx, "quota1", &loadbalancer.QuotaArgs{
// 			HealthMonitor: pulumi.Int(10),
// 			L7Policy:      pulumi.Int(11),
// 			L7Rule:        pulumi.Int(12),
// 			Listener:      pulumi.Int(7),
// 			Loadbalancer:  pulumi.Int(6),
// 			Member:        pulumi.Int(8),
// 			Pool:          pulumi.Int(9),
// 			ProjectId:     project1.ID(),
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
// Quotas can be imported using the `project_id/region_name`, where region_name is the one defined is the Openstack credentials that are in use. E.g.
//
// ```sh
//  $ pulumi import openstack:loadbalancer/quota:Quota quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
// ```
type Quota struct {
	pulumi.CustomResourceState

	// Quota value for health_monitors. Changing
	// this updates the existing quota. Omitting it sets it to 0.
	HealthMonitor pulumi.IntOutput `pulumi:"healthMonitor"`
	// Quota value for l7_policies. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Policy pulumi.IntOutput `pulumi:"l7Policy"`
	// Quota value for l7_rules. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Rule pulumi.IntOutput `pulumi:"l7Rule"`
	// Quota value for listeners. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Listener pulumi.IntOutput `pulumi:"listener"`
	// Quota value for loadbalancers. Changing this
	// updates the existing quota. Omitting it sets it to 0.
	Loadbalancer pulumi.IntOutput `pulumi:"loadbalancer"`
	// Quota value for members. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Member pulumi.IntOutput `pulumi:"member"`
	// Quota value for pools. Changing this updates the
	// the existing quota. Omitting it sets it to 0.
	Pool pulumi.IntOutput `pulumi:"pool"`
	// ID of the project to manage quotas. Changing this
	// creates a new quota.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Region in which to manage quotas. Changing this
	// creates a new quota. If ommited, the region of the credentials is used.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewQuota registers a new resource with the given unique name, arguments, and options.
func NewQuota(ctx *pulumi.Context,
	name string, args *QuotaArgs, opts ...pulumi.ResourceOption) (*Quota, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	var resource Quota
	err := ctx.RegisterResource("openstack:loadbalancer/quota:Quota", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuota gets an existing Quota resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuota(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaState, opts ...pulumi.ResourceOption) (*Quota, error) {
	var resource Quota
	err := ctx.ReadResource("openstack:loadbalancer/quota:Quota", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Quota resources.
type quotaState struct {
	// Quota value for health_monitors. Changing
	// this updates the existing quota. Omitting it sets it to 0.
	HealthMonitor *int `pulumi:"healthMonitor"`
	// Quota value for l7_policies. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Policy *int `pulumi:"l7Policy"`
	// Quota value for l7_rules. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Rule *int `pulumi:"l7Rule"`
	// Quota value for listeners. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Listener *int `pulumi:"listener"`
	// Quota value for loadbalancers. Changing this
	// updates the existing quota. Omitting it sets it to 0.
	Loadbalancer *int `pulumi:"loadbalancer"`
	// Quota value for members. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Member *int `pulumi:"member"`
	// Quota value for pools. Changing this updates the
	// the existing quota. Omitting it sets it to 0.
	Pool *int `pulumi:"pool"`
	// ID of the project to manage quotas. Changing this
	// creates a new quota.
	ProjectId *string `pulumi:"projectId"`
	// Region in which to manage quotas. Changing this
	// creates a new quota. If ommited, the region of the credentials is used.
	Region *string `pulumi:"region"`
}

type QuotaState struct {
	// Quota value for health_monitors. Changing
	// this updates the existing quota. Omitting it sets it to 0.
	HealthMonitor pulumi.IntPtrInput
	// Quota value for l7_policies. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Policy pulumi.IntPtrInput
	// Quota value for l7_rules. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Rule pulumi.IntPtrInput
	// Quota value for listeners. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Listener pulumi.IntPtrInput
	// Quota value for loadbalancers. Changing this
	// updates the existing quota. Omitting it sets it to 0.
	Loadbalancer pulumi.IntPtrInput
	// Quota value for members. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Member pulumi.IntPtrInput
	// Quota value for pools. Changing this updates the
	// the existing quota. Omitting it sets it to 0.
	Pool pulumi.IntPtrInput
	// ID of the project to manage quotas. Changing this
	// creates a new quota.
	ProjectId pulumi.StringPtrInput
	// Region in which to manage quotas. Changing this
	// creates a new quota. If ommited, the region of the credentials is used.
	Region pulumi.StringPtrInput
}

func (QuotaState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaState)(nil)).Elem()
}

type quotaArgs struct {
	// Quota value for health_monitors. Changing
	// this updates the existing quota. Omitting it sets it to 0.
	HealthMonitor *int `pulumi:"healthMonitor"`
	// Quota value for l7_policies. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Policy *int `pulumi:"l7Policy"`
	// Quota value for l7_rules. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Rule *int `pulumi:"l7Rule"`
	// Quota value for listeners. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Listener *int `pulumi:"listener"`
	// Quota value for loadbalancers. Changing this
	// updates the existing quota. Omitting it sets it to 0.
	Loadbalancer *int `pulumi:"loadbalancer"`
	// Quota value for members. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Member *int `pulumi:"member"`
	// Quota value for pools. Changing this updates the
	// the existing quota. Omitting it sets it to 0.
	Pool *int `pulumi:"pool"`
	// ID of the project to manage quotas. Changing this
	// creates a new quota.
	ProjectId string `pulumi:"projectId"`
	// Region in which to manage quotas. Changing this
	// creates a new quota. If ommited, the region of the credentials is used.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Quota resource.
type QuotaArgs struct {
	// Quota value for health_monitors. Changing
	// this updates the existing quota. Omitting it sets it to 0.
	HealthMonitor pulumi.IntPtrInput
	// Quota value for l7_policies. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Policy pulumi.IntPtrInput
	// Quota value for l7_rules. Changing this
	// updates the existing quota. Omitting it sets it to 0. Available in
	// Octavia 2.19.
	L7Rule pulumi.IntPtrInput
	// Quota value for listeners. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Listener pulumi.IntPtrInput
	// Quota value for loadbalancers. Changing this
	// updates the existing quota. Omitting it sets it to 0.
	Loadbalancer pulumi.IntPtrInput
	// Quota value for members. Changing this updates
	// the existing quota. Omitting it sets it to 0.
	Member pulumi.IntPtrInput
	// Quota value for pools. Changing this updates the
	// the existing quota. Omitting it sets it to 0.
	Pool pulumi.IntPtrInput
	// ID of the project to manage quotas. Changing this
	// creates a new quota.
	ProjectId pulumi.StringInput
	// Region in which to manage quotas. Changing this
	// creates a new quota. If ommited, the region of the credentials is used.
	Region pulumi.StringPtrInput
}

func (QuotaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaArgs)(nil)).Elem()
}

type QuotaInput interface {
	pulumi.Input

	ToQuotaOutput() QuotaOutput
	ToQuotaOutputWithContext(ctx context.Context) QuotaOutput
}

func (*Quota) ElementType() reflect.Type {
	return reflect.TypeOf((*Quota)(nil))
}

func (i *Quota) ToQuotaOutput() QuotaOutput {
	return i.ToQuotaOutputWithContext(context.Background())
}

func (i *Quota) ToQuotaOutputWithContext(ctx context.Context) QuotaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaOutput)
}

func (i *Quota) ToQuotaPtrOutput() QuotaPtrOutput {
	return i.ToQuotaPtrOutputWithContext(context.Background())
}

func (i *Quota) ToQuotaPtrOutputWithContext(ctx context.Context) QuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaPtrOutput)
}

type QuotaPtrInput interface {
	pulumi.Input

	ToQuotaPtrOutput() QuotaPtrOutput
	ToQuotaPtrOutputWithContext(ctx context.Context) QuotaPtrOutput
}

type quotaPtrType QuotaArgs

func (*quotaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Quota)(nil))
}

func (i *quotaPtrType) ToQuotaPtrOutput() QuotaPtrOutput {
	return i.ToQuotaPtrOutputWithContext(context.Background())
}

func (i *quotaPtrType) ToQuotaPtrOutputWithContext(ctx context.Context) QuotaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaPtrOutput)
}

// QuotaArrayInput is an input type that accepts QuotaArray and QuotaArrayOutput values.
// You can construct a concrete instance of `QuotaArrayInput` via:
//
//          QuotaArray{ QuotaArgs{...} }
type QuotaArrayInput interface {
	pulumi.Input

	ToQuotaArrayOutput() QuotaArrayOutput
	ToQuotaArrayOutputWithContext(context.Context) QuotaArrayOutput
}

type QuotaArray []QuotaInput

func (QuotaArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Quota)(nil))
}

func (i QuotaArray) ToQuotaArrayOutput() QuotaArrayOutput {
	return i.ToQuotaArrayOutputWithContext(context.Background())
}

func (i QuotaArray) ToQuotaArrayOutputWithContext(ctx context.Context) QuotaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaArrayOutput)
}

// QuotaMapInput is an input type that accepts QuotaMap and QuotaMapOutput values.
// You can construct a concrete instance of `QuotaMapInput` via:
//
//          QuotaMap{ "key": QuotaArgs{...} }
type QuotaMapInput interface {
	pulumi.Input

	ToQuotaMapOutput() QuotaMapOutput
	ToQuotaMapOutputWithContext(context.Context) QuotaMapOutput
}

type QuotaMap map[string]QuotaInput

func (QuotaMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Quota)(nil))
}

func (i QuotaMap) ToQuotaMapOutput() QuotaMapOutput {
	return i.ToQuotaMapOutputWithContext(context.Background())
}

func (i QuotaMap) ToQuotaMapOutputWithContext(ctx context.Context) QuotaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaMapOutput)
}

type QuotaOutput struct {
	*pulumi.OutputState
}

func (QuotaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Quota)(nil))
}

func (o QuotaOutput) ToQuotaOutput() QuotaOutput {
	return o
}

func (o QuotaOutput) ToQuotaOutputWithContext(ctx context.Context) QuotaOutput {
	return o
}

func (o QuotaOutput) ToQuotaPtrOutput() QuotaPtrOutput {
	return o.ToQuotaPtrOutputWithContext(context.Background())
}

func (o QuotaOutput) ToQuotaPtrOutputWithContext(ctx context.Context) QuotaPtrOutput {
	return o.ApplyT(func(v Quota) *Quota {
		return &v
	}).(QuotaPtrOutput)
}

type QuotaPtrOutput struct {
	*pulumi.OutputState
}

func (QuotaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Quota)(nil))
}

func (o QuotaPtrOutput) ToQuotaPtrOutput() QuotaPtrOutput {
	return o
}

func (o QuotaPtrOutput) ToQuotaPtrOutputWithContext(ctx context.Context) QuotaPtrOutput {
	return o
}

type QuotaArrayOutput struct{ *pulumi.OutputState }

func (QuotaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Quota)(nil))
}

func (o QuotaArrayOutput) ToQuotaArrayOutput() QuotaArrayOutput {
	return o
}

func (o QuotaArrayOutput) ToQuotaArrayOutputWithContext(ctx context.Context) QuotaArrayOutput {
	return o
}

func (o QuotaArrayOutput) Index(i pulumi.IntInput) QuotaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Quota {
		return vs[0].([]Quota)[vs[1].(int)]
	}).(QuotaOutput)
}

type QuotaMapOutput struct{ *pulumi.OutputState }

func (QuotaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Quota)(nil))
}

func (o QuotaMapOutput) ToQuotaMapOutput() QuotaMapOutput {
	return o
}

func (o QuotaMapOutput) ToQuotaMapOutputWithContext(ctx context.Context) QuotaMapOutput {
	return o
}

func (o QuotaMapOutput) MapIndex(k pulumi.StringInput) QuotaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Quota {
		return vs[0].(map[string]Quota)[vs[1].(string)]
	}).(QuotaOutput)
}

func init() {
	pulumi.RegisterOutputType(QuotaOutput{})
	pulumi.RegisterOutputType(QuotaPtrOutput{})
	pulumi.RegisterOutputType(QuotaArrayOutput{})
	pulumi.RegisterOutputType(QuotaMapOutput{})
}
