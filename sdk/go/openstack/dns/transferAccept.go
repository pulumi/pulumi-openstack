// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNS zone transfer accept in the OpenStack DNS Service.
//
// ## Example Usage
//
// ### Automatically detect the correct network
//
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
//			exampleZone, err := dns.NewZone(ctx, "exampleZone", &dns.ZoneArgs{
//				Email:       pulumi.String("jdoe@example.com"),
//				Description: pulumi.String("An example zone"),
//				Ttl:         pulumi.Int(3000),
//				Type:        pulumi.String("PRIMARY"),
//			})
//			if err != nil {
//				return err
//			}
//			request1, err := dns.NewTransferRequest(ctx, "request1", &dns.TransferRequestArgs{
//				ZoneId:      exampleZone.ID(),
//				Description: pulumi.String("a transfer accept"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewTransferAccept(ctx, "accept1", &dns.TransferAcceptArgs{
//				ZoneTransferRequestId: request1.ID(),
//				Key:                   request1.Key,
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
// This resource can be imported by specifying the transferAccept ID:
//
// ```sh
// $ pulumi import openstack:dns/transferAccept:TransferAccept accept_1 accept_id
// ```
type TransferAccept struct {
	pulumi.CustomResourceState

	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck pulumi.BoolPtrOutput `pulumi:"disableStatusCheck"`
	// The transfer key.
	Key pulumi.StringOutput `pulumi:"key"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of additional options. Changing this creates a
	// new transfer accept.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
	// The ID of the zone transfer request.
	ZoneTransferRequestId pulumi.StringOutput `pulumi:"zoneTransferRequestId"`
}

// NewTransferAccept registers a new resource with the given unique name, arguments, and options.
func NewTransferAccept(ctx *pulumi.Context,
	name string, args *TransferAcceptArgs, opts ...pulumi.ResourceOption) (*TransferAccept, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ZoneTransferRequestId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneTransferRequestId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransferAccept
	err := ctx.RegisterResource("openstack:dns/transferAccept:TransferAccept", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferAccept gets an existing TransferAccept resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferAccept(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferAcceptState, opts ...pulumi.ResourceOption) (*TransferAccept, error) {
	var resource TransferAccept
	err := ctx.ReadResource("openstack:dns/transferAccept:TransferAccept", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferAccept resources.
type transferAcceptState struct {
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck *bool `pulumi:"disableStatusCheck"`
	// The transfer key.
	Key *string `pulumi:"key"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region *string `pulumi:"region"`
	// Map of additional options. Changing this creates a
	// new transfer accept.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// The ID of the zone transfer request.
	ZoneTransferRequestId *string `pulumi:"zoneTransferRequestId"`
}

type TransferAcceptState struct {
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	// The transfer key.
	Key pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new transfer accept.
	ValueSpecs pulumi.MapInput
	// The ID of the zone transfer request.
	ZoneTransferRequestId pulumi.StringPtrInput
}

func (TransferAcceptState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferAcceptState)(nil)).Elem()
}

type transferAcceptArgs struct {
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck *bool `pulumi:"disableStatusCheck"`
	// The transfer key.
	Key string `pulumi:"key"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region *string `pulumi:"region"`
	// Map of additional options. Changing this creates a
	// new transfer accept.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
	// The ID of the zone transfer request.
	ZoneTransferRequestId string `pulumi:"zoneTransferRequestId"`
}

// The set of arguments for constructing a TransferAccept resource.
type TransferAcceptArgs struct {
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack accept returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	// The transfer key.
	Key pulumi.StringInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone.
	Region pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new transfer accept.
	ValueSpecs pulumi.MapInput
	// The ID of the zone transfer request.
	ZoneTransferRequestId pulumi.StringInput
}

func (TransferAcceptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferAcceptArgs)(nil)).Elem()
}

type TransferAcceptInput interface {
	pulumi.Input

	ToTransferAcceptOutput() TransferAcceptOutput
	ToTransferAcceptOutputWithContext(ctx context.Context) TransferAcceptOutput
}

func (*TransferAccept) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAccept)(nil)).Elem()
}

func (i *TransferAccept) ToTransferAcceptOutput() TransferAcceptOutput {
	return i.ToTransferAcceptOutputWithContext(context.Background())
}

func (i *TransferAccept) ToTransferAcceptOutputWithContext(ctx context.Context) TransferAcceptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAcceptOutput)
}

// TransferAcceptArrayInput is an input type that accepts TransferAcceptArray and TransferAcceptArrayOutput values.
// You can construct a concrete instance of `TransferAcceptArrayInput` via:
//
//	TransferAcceptArray{ TransferAcceptArgs{...} }
type TransferAcceptArrayInput interface {
	pulumi.Input

	ToTransferAcceptArrayOutput() TransferAcceptArrayOutput
	ToTransferAcceptArrayOutputWithContext(context.Context) TransferAcceptArrayOutput
}

type TransferAcceptArray []TransferAcceptInput

func (TransferAcceptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferAccept)(nil)).Elem()
}

func (i TransferAcceptArray) ToTransferAcceptArrayOutput() TransferAcceptArrayOutput {
	return i.ToTransferAcceptArrayOutputWithContext(context.Background())
}

func (i TransferAcceptArray) ToTransferAcceptArrayOutputWithContext(ctx context.Context) TransferAcceptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAcceptArrayOutput)
}

// TransferAcceptMapInput is an input type that accepts TransferAcceptMap and TransferAcceptMapOutput values.
// You can construct a concrete instance of `TransferAcceptMapInput` via:
//
//	TransferAcceptMap{ "key": TransferAcceptArgs{...} }
type TransferAcceptMapInput interface {
	pulumi.Input

	ToTransferAcceptMapOutput() TransferAcceptMapOutput
	ToTransferAcceptMapOutputWithContext(context.Context) TransferAcceptMapOutput
}

type TransferAcceptMap map[string]TransferAcceptInput

func (TransferAcceptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferAccept)(nil)).Elem()
}

func (i TransferAcceptMap) ToTransferAcceptMapOutput() TransferAcceptMapOutput {
	return i.ToTransferAcceptMapOutputWithContext(context.Background())
}

func (i TransferAcceptMap) ToTransferAcceptMapOutputWithContext(ctx context.Context) TransferAcceptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferAcceptMapOutput)
}

type TransferAcceptOutput struct{ *pulumi.OutputState }

func (TransferAcceptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferAccept)(nil)).Elem()
}

func (o TransferAcceptOutput) ToTransferAcceptOutput() TransferAcceptOutput {
	return o
}

func (o TransferAcceptOutput) ToTransferAcceptOutputWithContext(ctx context.Context) TransferAcceptOutput {
	return o
}

// Disable wait for zone to reach ACTIVE
// status. The check is enabled by default. If this argument is true, zone
// will be considered as created/updated if OpenStack accept returned success.
func (o TransferAcceptOutput) DisableStatusCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransferAccept) pulumi.BoolPtrOutput { return v.DisableStatusCheck }).(pulumi.BoolPtrOutput)
}

// The transfer key.
func (o TransferAcceptOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferAccept) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Compute client.
// Keypairs are associated with accounts, but a Compute client is needed to
// create one. If omitted, the `region` argument of the provider is used.
// Changing this creates a new DNS zone.
func (o TransferAcceptOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferAccept) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of additional options. Changing this creates a
// new transfer accept.
func (o TransferAcceptOutput) ValueSpecs() pulumi.MapOutput {
	return o.ApplyT(func(v *TransferAccept) pulumi.MapOutput { return v.ValueSpecs }).(pulumi.MapOutput)
}

// The ID of the zone transfer request.
func (o TransferAcceptOutput) ZoneTransferRequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferAccept) pulumi.StringOutput { return v.ZoneTransferRequestId }).(pulumi.StringOutput)
}

type TransferAcceptArrayOutput struct{ *pulumi.OutputState }

func (TransferAcceptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferAccept)(nil)).Elem()
}

func (o TransferAcceptArrayOutput) ToTransferAcceptArrayOutput() TransferAcceptArrayOutput {
	return o
}

func (o TransferAcceptArrayOutput) ToTransferAcceptArrayOutputWithContext(ctx context.Context) TransferAcceptArrayOutput {
	return o
}

func (o TransferAcceptArrayOutput) Index(i pulumi.IntInput) TransferAcceptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransferAccept {
		return vs[0].([]*TransferAccept)[vs[1].(int)]
	}).(TransferAcceptOutput)
}

type TransferAcceptMapOutput struct{ *pulumi.OutputState }

func (TransferAcceptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferAccept)(nil)).Elem()
}

func (o TransferAcceptMapOutput) ToTransferAcceptMapOutput() TransferAcceptMapOutput {
	return o
}

func (o TransferAcceptMapOutput) ToTransferAcceptMapOutputWithContext(ctx context.Context) TransferAcceptMapOutput {
	return o
}

func (o TransferAcceptMapOutput) MapIndex(k pulumi.StringInput) TransferAcceptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransferAccept {
		return vs[0].(map[string]*TransferAccept)[vs[1].(string)]
	}).(TransferAcceptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransferAcceptInput)(nil)).Elem(), &TransferAccept{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferAcceptArrayInput)(nil)).Elem(), TransferAcceptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferAcceptMapInput)(nil)).Elem(), TransferAcceptMap{})
	pulumi.RegisterOutputType(TransferAcceptOutput{})
	pulumi.RegisterOutputType(TransferAcceptArrayOutput{})
	pulumi.RegisterOutputType(TransferAcceptMapOutput{})
}
