// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNS zone in the OpenStack DNS Service.
//
// ## Example Usage
//
// ### Automatically detect the correct network
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dns.NewZone(ctx, "example_com", &dns.ZoneArgs{
//				Name:        pulumi.String("example.com."),
//				Email:       pulumi.String("jdoe@example.com"),
//				Description: pulumi.String("An example zone"),
//				Ttl:         pulumi.Int(3000),
//				Type:        pulumi.String("PRIMARY"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// This resource can be imported by specifying the zone ID with optional project ID:
//
// ```sh
// $ pulumi import openstack:dns/zone:Zone zone_1 zone_id
// ```
//
// ```sh
// $ pulumi import openstack:dns/zone:Zone zone_1 zone_id/project_id
// ```
type Zone struct {
	pulumi.CustomResourceState

	// Attributes for the DNS Service scheduler.
	// Changing this creates a new zone.
	Attributes pulumi.MapOutput `pulumi:"attributes"`
	// A description of the zone.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrOutput `pulumi:"disableStatusCheck"`
	// The email contact for the zone record.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// An array of master DNS servers. For when `type` is
	// `SECONDARY`.
	Masters pulumi.StringArrayOutput `pulumi:"masters"`
	// The name of the zone. Note the `.` at the end of the name.
	// Changing this creates a new DNS zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project DNS zone is created
	// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
	// user role in target project)
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region pulumi.StringOutput `pulumi:"region"`
	// The time to live (TTL) of the zone.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
	// Changing this creates a new zone.
	Type pulumi.StringOutput `pulumi:"type"`
	// Map of additional options. Changing this creates a
	// new zone.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		args = &ZoneArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Zone
	err := ctx.RegisterResource("openstack:dns/zone:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("openstack:dns/zone:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
	// Attributes for the DNS Service scheduler.
	// Changing this creates a new zone.
	Attributes map[string]interface{} `pulumi:"attributes"`
	// A description of the zone.
	Description *string `pulumi:"description"`
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck *bool `pulumi:"disableStatusCheck"`
	// The email contact for the zone record.
	Email *string `pulumi:"email"`
	// An array of master DNS servers. For when `type` is
	// `SECONDARY`.
	Masters []string `pulumi:"masters"`
	// The name of the zone. Note the `.` at the end of the name.
	// Changing this creates a new DNS zone.
	Name *string `pulumi:"name"`
	// The ID of the project DNS zone is created
	// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
	// user role in target project)
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region *string `pulumi:"region"`
	// The time to live (TTL) of the zone.
	Ttl *int `pulumi:"ttl"`
	// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
	// Changing this creates a new zone.
	Type *string `pulumi:"type"`
	// Map of additional options. Changing this creates a
	// new zone.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type ZoneState struct {
	// Attributes for the DNS Service scheduler.
	// Changing this creates a new zone.
	Attributes pulumi.MapInput
	// A description of the zone.
	Description pulumi.StringPtrInput
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	// The email contact for the zone record.
	Email pulumi.StringPtrInput
	// An array of master DNS servers. For when `type` is
	// `SECONDARY`.
	Masters pulumi.StringArrayInput
	// The name of the zone. Note the `.` at the end of the name.
	// Changing this creates a new DNS zone.
	Name pulumi.StringPtrInput
	// The ID of the project DNS zone is created
	// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
	// user role in target project)
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region pulumi.StringPtrInput
	// The time to live (TTL) of the zone.
	Ttl pulumi.IntPtrInput
	// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
	// Changing this creates a new zone.
	Type pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new zone.
	ValueSpecs pulumi.MapInput
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// Attributes for the DNS Service scheduler.
	// Changing this creates a new zone.
	Attributes map[string]interface{} `pulumi:"attributes"`
	// A description of the zone.
	Description *string `pulumi:"description"`
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck *bool `pulumi:"disableStatusCheck"`
	// The email contact for the zone record.
	Email *string `pulumi:"email"`
	// An array of master DNS servers. For when `type` is
	// `SECONDARY`.
	Masters []string `pulumi:"masters"`
	// The name of the zone. Note the `.` at the end of the name.
	// Changing this creates a new DNS zone.
	Name *string `pulumi:"name"`
	// The ID of the project DNS zone is created
	// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
	// user role in target project)
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region *string `pulumi:"region"`
	// The time to live (TTL) of the zone.
	Ttl *int `pulumi:"ttl"`
	// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
	// Changing this creates a new zone.
	Type *string `pulumi:"type"`
	// Map of additional options. Changing this creates a
	// new zone.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// Attributes for the DNS Service scheduler.
	// Changing this creates a new zone.
	Attributes pulumi.MapInput
	// A description of the zone.
	Description pulumi.StringPtrInput
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	// The email contact for the zone record.
	Email pulumi.StringPtrInput
	// An array of master DNS servers. For when `type` is
	// `SECONDARY`.
	Masters pulumi.StringArrayInput
	// The name of the zone. Note the `.` at the end of the name.
	// Changing this creates a new DNS zone.
	Name pulumi.StringPtrInput
	// The ID of the project DNS zone is created
	// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
	// user role in target project)
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region pulumi.StringPtrInput
	// The time to live (TTL) of the zone.
	Ttl pulumi.IntPtrInput
	// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
	// Changing this creates a new zone.
	Type pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new zone.
	ValueSpecs pulumi.MapInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

// ZoneArrayInput is an input type that accepts ZoneArray and ZoneArrayOutput values.
// You can construct a concrete instance of `ZoneArrayInput` via:
//
//	ZoneArray{ ZoneArgs{...} }
type ZoneArrayInput interface {
	pulumi.Input

	ToZoneArrayOutput() ZoneArrayOutput
	ToZoneArrayOutputWithContext(context.Context) ZoneArrayOutput
}

type ZoneArray []ZoneInput

func (ZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (i ZoneArray) ToZoneArrayOutput() ZoneArrayOutput {
	return i.ToZoneArrayOutputWithContext(context.Background())
}

func (i ZoneArray) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneArrayOutput)
}

// ZoneMapInput is an input type that accepts ZoneMap and ZoneMapOutput values.
// You can construct a concrete instance of `ZoneMapInput` via:
//
//	ZoneMap{ "key": ZoneArgs{...} }
type ZoneMapInput interface {
	pulumi.Input

	ToZoneMapOutput() ZoneMapOutput
	ToZoneMapOutputWithContext(context.Context) ZoneMapOutput
}

type ZoneMap map[string]ZoneInput

func (ZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (i ZoneMap) ToZoneMapOutput() ZoneMapOutput {
	return i.ToZoneMapOutputWithContext(context.Background())
}

func (i ZoneMap) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneMapOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

// Attributes for the DNS Service scheduler.
// Changing this creates a new zone.
func (o ZoneOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v *Zone) pulumi.MapOutput { return v.Attributes }).(pulumi.MapOutput)
}

// A description of the zone.
func (o ZoneOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Disable wait for zone to reach ACTIVE
// status. The check is enabled by default. If this argument is true, zone
// will be considered as created/updated if OpenStack request returned success.
func (o ZoneOutput) DisableStatusCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.BoolPtrOutput { return v.DisableStatusCheck }).(pulumi.BoolPtrOutput)
}

// The email contact for the zone record.
func (o ZoneOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// An array of master DNS servers. For when `type` is
// `SECONDARY`.
func (o ZoneOutput) Masters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.Masters }).(pulumi.StringArrayOutput)
}

// The name of the zone. Note the `.` at the end of the name.
// Changing this creates a new DNS zone.
func (o ZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project DNS zone is created
// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
// user role in target project)
func (o ZoneOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Compute client.
// Keypairs are associated with accounts, but a Compute client is needed to
// create one. If omitted, the `region` argument of the provider is used.
// Changing this creates a new DNS zone.
func (o ZoneOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The time to live (TTL) of the zone.
func (o ZoneOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *Zone) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
// Changing this creates a new zone.
func (o ZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Map of additional options. Changing this creates a
// new zone.
func (o ZoneOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *Zone) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

type ZoneArrayOutput struct{ *pulumi.OutputState }

func (ZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Zone)(nil)).Elem()
}

func (o ZoneArrayOutput) ToZoneArrayOutput() ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) ToZoneArrayOutputWithContext(ctx context.Context) ZoneArrayOutput {
	return o
}

func (o ZoneArrayOutput) Index(i pulumi.IntInput) ZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].([]*Zone)[vs[1].(int)]
	}).(ZoneOutput)
}

type ZoneMapOutput struct{ *pulumi.OutputState }

func (ZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Zone)(nil)).Elem()
}

func (o ZoneMapOutput) ToZoneMapOutput() ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) ToZoneMapOutputWithContext(ctx context.Context) ZoneMapOutput {
	return o
}

func (o ZoneMapOutput) MapIndex(k pulumi.StringInput) ZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Zone {
		return vs[0].(map[string]*Zone)[vs[1].(string)]
	}).(ZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneInput)(nil)).Elem(), &Zone{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneArrayInput)(nil)).Elem(), ZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneMapInput)(nil)).Elem(), ZoneMap{})
	pulumi.RegisterOutputType(ZoneOutput{})
	pulumi.RegisterOutputType(ZoneArrayOutput{})
	pulumi.RegisterOutputType(ZoneMapOutput{})
}
