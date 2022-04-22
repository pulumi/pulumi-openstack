// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack DNS zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/dns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dns.GetDnsZone(ctx, &dns.GetDnsZoneArgs{
// 			Name: pulumi.StringRef("example.com"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDnsZone(ctx *pulumi.Context, args *GetDnsZoneArgs, opts ...pulumi.InvokeOption) (*GetDnsZoneResult, error) {
	var rv GetDnsZoneResult
	err := ctx.Invoke("openstack:dns/getDnsZone:getDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsZone.
type GetDnsZoneArgs struct {
	// Try to obtain zone ID by listing all projects
	// (requires admin role by default, depends on your policy configuration)
	AllProjects *string `pulumi:"allProjects"`
	// Attributes of the DNS Service scheduler.
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The time the zone was created.
	CreatedAt *string `pulumi:"createdAt"`
	// A description of the zone.
	Description *string `pulumi:"description"`
	// The email contact for the zone record.
	Email *string `pulumi:"email"`
	// An array of master DNS servers. When `type` is  `SECONDARY`.
	Masters []string `pulumi:"masters"`
	// The name of the zone.
	Name *string `pulumi:"name"`
	// The ID of the pool hosting the zone.
	PoolId *string `pulumi:"poolId"`
	// The ID of the project the DNS zone is obtained from,
	// sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned user role in target project)
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 DNS client.
	// A DNS client is needed to retrieve zone ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The serial number of the zone.
	Serial *int `pulumi:"serial"`
	// The zone's status.
	Status *string `pulumi:"status"`
	// The time the zone was transferred.
	TransferredAt *string `pulumi:"transferredAt"`
	// The time to live (TTL) of the zone.
	Ttl *int `pulumi:"ttl"`
	// The type of the zone. Can either be `PRIMARY` or `SECONDARY`.
	Type *string `pulumi:"type"`
	// The time the zone was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The version of the zone.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getDnsZone.
type GetDnsZoneResult struct {
	AllProjects *string `pulumi:"allProjects"`
	// Attributes of the DNS Service scheduler.
	Attributes map[string]interface{} `pulumi:"attributes"`
	// The time the zone was created.
	CreatedAt string `pulumi:"createdAt"`
	// See Argument Reference above.
	Description *string `pulumi:"description"`
	// See Argument Reference above.
	Email *string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An array of master DNS servers. When `type` is  `SECONDARY`.
	Masters []string `pulumi:"masters"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// The ID of the pool hosting the zone.
	PoolId string `pulumi:"poolId"`
	// The project ID that owns the zone.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The serial number of the zone.
	Serial int `pulumi:"serial"`
	// See Argument Reference above.
	Status *string `pulumi:"status"`
	// The time the zone was transferred.
	TransferredAt string `pulumi:"transferredAt"`
	// See Argument Reference above.
	Ttl *int `pulumi:"ttl"`
	// See Argument Reference above.
	Type *string `pulumi:"type"`
	// The time the zone was last updated.
	UpdatedAt string `pulumi:"updatedAt"`
	// The version of the zone.
	Version int `pulumi:"version"`
}

func GetDnsZoneOutput(ctx *pulumi.Context, args GetDnsZoneOutputArgs, opts ...pulumi.InvokeOption) GetDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsZoneResult, error) {
			args := v.(GetDnsZoneArgs)
			r, err := GetDnsZone(ctx, &args, opts...)
			var s GetDnsZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDnsZoneResultOutput)
}

// A collection of arguments for invoking getDnsZone.
type GetDnsZoneOutputArgs struct {
	// Try to obtain zone ID by listing all projects
	// (requires admin role by default, depends on your policy configuration)
	AllProjects pulumi.StringPtrInput `pulumi:"allProjects"`
	// Attributes of the DNS Service scheduler.
	Attributes pulumi.MapInput `pulumi:"attributes"`
	// The time the zone was created.
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// A description of the zone.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The email contact for the zone record.
	Email pulumi.StringPtrInput `pulumi:"email"`
	// An array of master DNS servers. When `type` is  `SECONDARY`.
	Masters pulumi.StringArrayInput `pulumi:"masters"`
	// The name of the zone.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the pool hosting the zone.
	PoolId pulumi.StringPtrInput `pulumi:"poolId"`
	// The ID of the project the DNS zone is obtained from,
	// sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned user role in target project)
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// The region in which to obtain the V2 DNS client.
	// A DNS client is needed to retrieve zone ids. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The serial number of the zone.
	Serial pulumi.IntPtrInput `pulumi:"serial"`
	// The zone's status.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The time the zone was transferred.
	TransferredAt pulumi.StringPtrInput `pulumi:"transferredAt"`
	// The time to live (TTL) of the zone.
	Ttl pulumi.IntPtrInput `pulumi:"ttl"`
	// The type of the zone. Can either be `PRIMARY` or `SECONDARY`.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The time the zone was last updated.
	UpdatedAt pulumi.StringPtrInput `pulumi:"updatedAt"`
	// The version of the zone.
	Version pulumi.IntPtrInput `pulumi:"version"`
}

func (GetDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsZoneArgs)(nil)).Elem()
}

// A collection of values returned by getDnsZone.
type GetDnsZoneResultOutput struct{ *pulumi.OutputState }

func (GetDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsZoneResult)(nil)).Elem()
}

func (o GetDnsZoneResultOutput) ToGetDnsZoneResultOutput() GetDnsZoneResultOutput {
	return o
}

func (o GetDnsZoneResultOutput) ToGetDnsZoneResultOutputWithContext(ctx context.Context) GetDnsZoneResultOutput {
	return o
}

func (o GetDnsZoneResultOutput) AllProjects() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *string { return v.AllProjects }).(pulumi.StringPtrOutput)
}

// Attributes of the DNS Service scheduler.
func (o GetDnsZoneResultOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v GetDnsZoneResult) map[string]interface{} { return v.Attributes }).(pulumi.MapOutput)
}

// The time the zone was created.
func (o GetDnsZoneResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDnsZoneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.Id }).(pulumi.StringOutput)
}

// An array of master DNS servers. When `type` is  `SECONDARY`.
func (o GetDnsZoneResultOutput) Masters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDnsZoneResult) []string { return v.Masters }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the pool hosting the zone.
func (o GetDnsZoneResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.PoolId }).(pulumi.StringOutput)
}

// The project ID that owns the zone.
func (o GetDnsZoneResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.Region }).(pulumi.StringOutput)
}

// The serial number of the zone.
func (o GetDnsZoneResultOutput) Serial() pulumi.IntOutput {
	return o.ApplyT(func(v GetDnsZoneResult) int { return v.Serial }).(pulumi.IntOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The time the zone was transferred.
func (o GetDnsZoneResultOutput) TransferredAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.TransferredAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *int { return v.Ttl }).(pulumi.IntPtrOutput)
}

// See Argument Reference above.
func (o GetDnsZoneResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDnsZoneResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The time the zone was last updated.
func (o GetDnsZoneResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsZoneResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The version of the zone.
func (o GetDnsZoneResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v GetDnsZoneResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsZoneResultOutput{})
}
