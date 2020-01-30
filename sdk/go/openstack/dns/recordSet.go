// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a DNS record set in the OpenStack DNS Service.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/dns_recordset_v2.html.markdown.
type RecordSet struct {
	pulumi.CustomResourceState

	// A description of the  record set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
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
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &RecordSetArgs{}
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

