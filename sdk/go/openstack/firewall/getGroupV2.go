// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information of an available OpenStack firewall group v2.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/firewall"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := firewall.LookupGroupV2(ctx, &firewall.LookupGroupV2Args{
//				Name: pulumi.StringRef("tf_test_group"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGroupV2(ctx *pulumi.Context, args *LookupGroupV2Args, opts ...pulumi.InvokeOption) (*LookupGroupV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupV2Result
	err := ctx.Invoke("openstack:firewall/getGroupV2:getGroupV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupV2.
type LookupGroupV2Args struct {
	// Administrative up/down status for the firewall group.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Human-readable description of the firewall group.
	Description *string `pulumi:"description"`
	// The egress policy ID of the firewall group.
	EgressFirewallPolicyId *string `pulumi:"egressFirewallPolicyId"`
	// The ID of the firewall group.
	GroupId *string `pulumi:"groupId"`
	// The ingress policy ID of the firewall group.
	IngressFirewallPolicyId *string `pulumi:"ingressFirewallPolicyId"`
	// The name of the firewall group.
	Name *string `pulumi:"name"`
	// This argument conflicts and is interchangeable
	// with `tenantId`. The owner of the firewall group.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve firewall group ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The sharing status of the firewall group.
	Shared *bool `pulumi:"shared"`
	// Enabled status for the firewall group.
	Status *string `pulumi:"status"`
	// This argument conflicts and is interchangeable
	// with `projectId`. The owner of the firewall group.
	TenantId *string `pulumi:"tenantId"`
}

// A collection of values returned by getGroupV2.
type LookupGroupV2Result struct {
	// See Argument Reference above.
	AdminStateUp bool `pulumi:"adminStateUp"`
	// See Argument Reference above.
	Description *string `pulumi:"description"`
	// See Argument Reference above.
	EgressFirewallPolicyId *string `pulumi:"egressFirewallPolicyId"`
	// See Argument Reference above.
	GroupId *string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IngressFirewallPolicyId *string `pulumi:"ingressFirewallPolicyId"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// Ports associated with the firewall group.
	Ports []string `pulumi:"ports"`
	// See Argument Reference above.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	Shared bool `pulumi:"shared"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
	// See Argument Reference above.
	TenantId string `pulumi:"tenantId"`
}

func LookupGroupV2Output(ctx *pulumi.Context, args LookupGroupV2OutputArgs, opts ...pulumi.InvokeOption) LookupGroupV2ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGroupV2ResultOutput, error) {
			args := v.(LookupGroupV2Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:firewall/getGroupV2:getGroupV2", args, LookupGroupV2ResultOutput{}, options).(LookupGroupV2ResultOutput), nil
		}).(LookupGroupV2ResultOutput)
}

// A collection of arguments for invoking getGroupV2.
type LookupGroupV2OutputArgs struct {
	// Administrative up/down status for the firewall group.
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// Human-readable description of the firewall group.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The egress policy ID of the firewall group.
	EgressFirewallPolicyId pulumi.StringPtrInput `pulumi:"egressFirewallPolicyId"`
	// The ID of the firewall group.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// The ingress policy ID of the firewall group.
	IngressFirewallPolicyId pulumi.StringPtrInput `pulumi:"ingressFirewallPolicyId"`
	// The name of the firewall group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// This argument conflicts and is interchangeable
	// with `tenantId`. The owner of the firewall group.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve firewall group ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The sharing status of the firewall group.
	Shared pulumi.BoolPtrInput `pulumi:"shared"`
	// Enabled status for the firewall group.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// This argument conflicts and is interchangeable
	// with `projectId`. The owner of the firewall group.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupGroupV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupV2Args)(nil)).Elem()
}

// A collection of values returned by getGroupV2.
type LookupGroupV2ResultOutput struct{ *pulumi.OutputState }

func (LookupGroupV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupV2Result)(nil)).Elem()
}

func (o LookupGroupV2ResultOutput) ToLookupGroupV2ResultOutput() LookupGroupV2ResultOutput {
	return o
}

func (o LookupGroupV2ResultOutput) ToLookupGroupV2ResultOutputWithContext(ctx context.Context) LookupGroupV2ResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) AdminStateUp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupV2Result) bool { return v.AdminStateUp }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupV2Result) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) EgressFirewallPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupV2Result) *string { return v.EgressFirewallPolicyId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupV2Result) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGroupV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) IngressFirewallPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupV2Result) *string { return v.IngressFirewallPolicyId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Ports associated with the firewall group.
func (o LookupGroupV2ResultOutput) Ports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupV2Result) []string { return v.Ports }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupV2Result) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupV2Result) bool { return v.Shared }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupV2Result) string { return v.Status }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupGroupV2ResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupV2Result) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupV2ResultOutput{})
}
