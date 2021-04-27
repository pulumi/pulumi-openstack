// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNS record set in the OpenStack DNS Service.
//
// ## Example Usage
// ### Automatically detect the correct network
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
// 		exampleZone, err := dns.NewZone(ctx, "exampleZone", &dns.ZoneArgs{
// 			Description: pulumi.String("a zone"),
// 			Email:       pulumi.String("email2@example.com"),
// 			Ttl:         pulumi.Int(6000),
// 			Type:        pulumi.String("PRIMARY"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewRecordSet(ctx, "rsExampleCom", &dns.RecordSetArgs{
// 			Description: pulumi.String("An example record set"),
// 			Records: pulumi.StringArray{
// 				pulumi.String("10.0.0.1"),
// 			},
// 			Ttl:    pulumi.Int(3000),
// 			Type:   pulumi.String("A"),
// 			ZoneId: exampleZone.ID(),
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
// This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash.
//
// ```sh
//  $ pulumi import openstack:dns/recordSet:RecordSet recordset_1 <zone_id>/<recordset_id>
// ```
type RecordSet struct {
	pulumi.CustomResourceState

	// A description of the  record set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrOutput `pulumi:"disableStatusCheck"`
	// The name of the record set. Note the `.` at the end of the name.
	// Changing this creates a new DNS  record set.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of DNS records. _Note:_ if an IPv6 address
	// contains brackets (`[ ]`), the brackets will be stripped and the modified
	// address will be recorded in the state.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS  record set.
	Region pulumi.StringOutput `pulumi:"region"`
	// The time to live (TTL) of the record set.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	Type pulumi.StringOutput `pulumi:"type"`
	// Map of additional options. Changing this creates a
	// new record set.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource RecordSet
	err := ctx.RegisterResource("openstack:dns/recordSet:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("openstack:dns/recordSet:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSet resources.
type recordSetState struct {
	// A description of the  record set.
	Description *string `pulumi:"description"`
	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	DisableStatusCheck *bool `pulumi:"disableStatusCheck"`
	// The name of the record set. Note the `.` at the end of the name.
	// Changing this creates a new DNS  record set.
	Name *string `pulumi:"name"`
	// An array of DNS records. _Note:_ if an IPv6 address
	// contains brackets (`[ ]`), the brackets will be stripped and the modified
	// address will be recorded in the state.
	Records []string `pulumi:"records"`
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS  record set.
	Region *string `pulumi:"region"`
	// The time to live (TTL) of the record set.
	Ttl *int `pulumi:"ttl"`
	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	Type *string `pulumi:"type"`
	// Map of additional options. Changing this creates a
	// new record set.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	ZoneId *string `pulumi:"zoneId"`
}

type RecordSetState struct {
	// A description of the  record set.
	Description pulumi.StringPtrInput
	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	// The name of the record set. Note the `.` at the end of the name.
	// Changing this creates a new DNS  record set.
	Name pulumi.StringPtrInput
	// An array of DNS records. _Note:_ if an IPv6 address
	// contains brackets (`[ ]`), the brackets will be stripped and the modified
	// address will be recorded in the state.
	Records pulumi.StringArrayInput
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS  record set.
	Region pulumi.StringPtrInput
	// The time to live (TTL) of the record set.
	Ttl pulumi.IntPtrInput
	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	Type pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new record set.
	ValueSpecs pulumi.MapInput
	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	ZoneId pulumi.StringPtrInput
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	// A description of the  record set.
	Description *string `pulumi:"description"`
	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	DisableStatusCheck *bool `pulumi:"disableStatusCheck"`
	// The name of the record set. Note the `.` at the end of the name.
	// Changing this creates a new DNS  record set.
	Name *string `pulumi:"name"`
	// An array of DNS records. _Note:_ if an IPv6 address
	// contains brackets (`[ ]`), the brackets will be stripped and the modified
	// address will be recorded in the state.
	Records []string `pulumi:"records"`
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS  record set.
	Region *string `pulumi:"region"`
	// The time to live (TTL) of the record set.
	Ttl *int `pulumi:"ttl"`
	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	Type *string `pulumi:"type"`
	// Map of additional options. Changing this creates a
	// new record set.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	// A description of the  record set.
	Description pulumi.StringPtrInput
	// Disable wait for recordset to reach ACTIVE
	// status. This argumen is disabled by default. If it is set to true, the recordset
	// will be considered as created/updated/deleted if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	// The name of the record set. Note the `.` at the end of the name.
	// Changing this creates a new DNS  record set.
	Name pulumi.StringPtrInput
	// An array of DNS records. _Note:_ if an IPv6 address
	// contains brackets (`[ ]`), the brackets will be stripped and the modified
	// address will be recorded in the state.
	Records pulumi.StringArrayInput
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS  record set.
	Region pulumi.StringPtrInput
	// The time to live (TTL) of the record set.
	Ttl pulumi.IntPtrInput
	// The type of record set. Examples: "A", "MX".
	// Changing this creates a new DNS  record set.
	Type pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new record set.
	ValueSpecs pulumi.MapInput
	// The ID of the zone in which to create the record set.
	// Changing this creates a new DNS  record set.
	ZoneId pulumi.StringInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}

type RecordSetInput interface {
	pulumi.Input

	ToRecordSetOutput() RecordSetOutput
	ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput
}

func (*RecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSet)(nil))
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

func (i *RecordSet) ToRecordSetPtrOutput() RecordSetPtrOutput {
	return i.ToRecordSetPtrOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetPtrOutputWithContext(ctx context.Context) RecordSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetPtrOutput)
}

type RecordSetPtrInput interface {
	pulumi.Input

	ToRecordSetPtrOutput() RecordSetPtrOutput
	ToRecordSetPtrOutputWithContext(ctx context.Context) RecordSetPtrOutput
}

type recordSetPtrType RecordSetArgs

func (*recordSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil))
}

func (i *recordSetPtrType) ToRecordSetPtrOutput() RecordSetPtrOutput {
	return i.ToRecordSetPtrOutputWithContext(context.Background())
}

func (i *recordSetPtrType) ToRecordSetPtrOutputWithContext(ctx context.Context) RecordSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetPtrOutput)
}

// RecordSetArrayInput is an input type that accepts RecordSetArray and RecordSetArrayOutput values.
// You can construct a concrete instance of `RecordSetArrayInput` via:
//
//          RecordSetArray{ RecordSetArgs{...} }
type RecordSetArrayInput interface {
	pulumi.Input

	ToRecordSetArrayOutput() RecordSetArrayOutput
	ToRecordSetArrayOutputWithContext(context.Context) RecordSetArrayOutput
}

type RecordSetArray []RecordSetInput

func (RecordSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RecordSet)(nil))
}

func (i RecordSetArray) ToRecordSetArrayOutput() RecordSetArrayOutput {
	return i.ToRecordSetArrayOutputWithContext(context.Background())
}

func (i RecordSetArray) ToRecordSetArrayOutputWithContext(ctx context.Context) RecordSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetArrayOutput)
}

// RecordSetMapInput is an input type that accepts RecordSetMap and RecordSetMapOutput values.
// You can construct a concrete instance of `RecordSetMapInput` via:
//
//          RecordSetMap{ "key": RecordSetArgs{...} }
type RecordSetMapInput interface {
	pulumi.Input

	ToRecordSetMapOutput() RecordSetMapOutput
	ToRecordSetMapOutputWithContext(context.Context) RecordSetMapOutput
}

type RecordSetMap map[string]RecordSetInput

func (RecordSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RecordSet)(nil))
}

func (i RecordSetMap) ToRecordSetMapOutput() RecordSetMapOutput {
	return i.ToRecordSetMapOutputWithContext(context.Background())
}

func (i RecordSetMap) ToRecordSetMapOutputWithContext(ctx context.Context) RecordSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetMapOutput)
}

type RecordSetOutput struct {
	*pulumi.OutputState
}

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSet)(nil))
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetPtrOutput() RecordSetPtrOutput {
	return o.ToRecordSetPtrOutputWithContext(context.Background())
}

func (o RecordSetOutput) ToRecordSetPtrOutputWithContext(ctx context.Context) RecordSetPtrOutput {
	return o.ApplyT(func(v RecordSet) *RecordSet {
		return &v
	}).(RecordSetPtrOutput)
}

type RecordSetPtrOutput struct {
	*pulumi.OutputState
}

func (RecordSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil))
}

func (o RecordSetPtrOutput) ToRecordSetPtrOutput() RecordSetPtrOutput {
	return o
}

func (o RecordSetPtrOutput) ToRecordSetPtrOutputWithContext(ctx context.Context) RecordSetPtrOutput {
	return o
}

type RecordSetArrayOutput struct{ *pulumi.OutputState }

func (RecordSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecordSet)(nil))
}

func (o RecordSetArrayOutput) ToRecordSetArrayOutput() RecordSetArrayOutput {
	return o
}

func (o RecordSetArrayOutput) ToRecordSetArrayOutputWithContext(ctx context.Context) RecordSetArrayOutput {
	return o
}

func (o RecordSetArrayOutput) Index(i pulumi.IntInput) RecordSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecordSet {
		return vs[0].([]RecordSet)[vs[1].(int)]
	}).(RecordSetOutput)
}

type RecordSetMapOutput struct{ *pulumi.OutputState }

func (RecordSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RecordSet)(nil))
}

func (o RecordSetMapOutput) ToRecordSetMapOutput() RecordSetMapOutput {
	return o
}

func (o RecordSetMapOutput) ToRecordSetMapOutputWithContext(ctx context.Context) RecordSetMapOutput {
	return o
}

func (o RecordSetMapOutput) MapIndex(k pulumi.StringInput) RecordSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RecordSet {
		return vs[0].(map[string]RecordSet)[vs[1].(string)]
	}).(RecordSetOutput)
}

func init() {
	pulumi.RegisterOutputType(RecordSetOutput{})
	pulumi.RegisterOutputType(RecordSetPtrOutput{})
	pulumi.RegisterOutputType(RecordSetArrayOutput{})
	pulumi.RegisterOutputType(RecordSetMapOutput{})
}
