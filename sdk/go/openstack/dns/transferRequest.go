// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DNS zone transfer request in the OpenStack DNS Service.
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
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/dns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleZone, err := dns.NewZone(ctx, "example_zone", &dns.ZoneArgs{
//				Name:        pulumi.String("example.com."),
//				Email:       pulumi.String("jdoe@example.com"),
//				Description: pulumi.String("An example zone"),
//				Ttl:         pulumi.Int(3000),
//				Type:        pulumi.String("PRIMARY"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = dns.NewTransferRequest(ctx, "request_1", &dns.TransferRequestArgs{
//				ZoneId:      exampleZone.ID(),
//				Description: pulumi.String("a transfer request"),
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
// This resource can be imported by specifying the transferRequest ID:
//
// ```sh
// $ pulumi import openstack:dns/transferRequest:TransferRequest request_1 request_id
// ```
type TransferRequest struct {
	pulumi.CustomResourceState

	// A description of the zone tranfer request.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrOutput `pulumi:"disableStatusCheck"`
	Key                pulumi.StringOutput  `pulumi:"key"`
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone zone transfer accept.
	Region pulumi.StringOutput `pulumi:"region"`
	// The target Project ID to transfer to.
	TargetProjectId pulumi.StringOutput `pulumi:"targetProjectId"`
	// Map of additional options. Changing this creates a
	// new transfer request.
	ValueSpecs pulumi.StringMapOutput `pulumi:"valueSpecs"`
	// The ID of the zone for which to create the transfer
	// request.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewTransferRequest registers a new resource with the given unique name, arguments, and options.
func NewTransferRequest(ctx *pulumi.Context,
	name string, args *TransferRequestArgs, opts ...pulumi.ResourceOption) (*TransferRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TransferRequest
	err := ctx.RegisterResource("openstack:dns/transferRequest:TransferRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferRequest gets an existing TransferRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferRequestState, opts ...pulumi.ResourceOption) (*TransferRequest, error) {
	var resource TransferRequest
	err := ctx.ReadResource("openstack:dns/transferRequest:TransferRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferRequest resources.
type transferRequestState struct {
	// A description of the zone tranfer request.
	Description *string `pulumi:"description"`
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck *bool   `pulumi:"disableStatusCheck"`
	Key                *string `pulumi:"key"`
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone zone transfer accept.
	Region *string `pulumi:"region"`
	// The target Project ID to transfer to.
	TargetProjectId *string `pulumi:"targetProjectId"`
	// Map of additional options. Changing this creates a
	// new transfer request.
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
	// The ID of the zone for which to create the transfer
	// request.
	ZoneId *string `pulumi:"zoneId"`
}

type TransferRequestState struct {
	// A description of the zone tranfer request.
	Description pulumi.StringPtrInput
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	Key                pulumi.StringPtrInput
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone zone transfer accept.
	Region pulumi.StringPtrInput
	// The target Project ID to transfer to.
	TargetProjectId pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new transfer request.
	ValueSpecs pulumi.StringMapInput
	// The ID of the zone for which to create the transfer
	// request.
	ZoneId pulumi.StringPtrInput
}

func (TransferRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferRequestState)(nil)).Elem()
}

type transferRequestArgs struct {
	// A description of the zone tranfer request.
	Description *string `pulumi:"description"`
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck *bool   `pulumi:"disableStatusCheck"`
	Key                *string `pulumi:"key"`
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone zone transfer accept.
	Region *string `pulumi:"region"`
	// The target Project ID to transfer to.
	TargetProjectId *string `pulumi:"targetProjectId"`
	// Map of additional options. Changing this creates a
	// new transfer request.
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
	// The ID of the zone for which to create the transfer
	// request.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a TransferRequest resource.
type TransferRequestArgs struct {
	// A description of the zone tranfer request.
	Description pulumi.StringPtrInput
	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	DisableStatusCheck pulumi.BoolPtrInput
	Key                pulumi.StringPtrInput
	// The region in which to obtain the V2 DNS client.
	// If omitted, the `region` argument of the provider is used.
	// Changing this creates a new DNS zone zone transfer accept.
	Region pulumi.StringPtrInput
	// The target Project ID to transfer to.
	TargetProjectId pulumi.StringPtrInput
	// Map of additional options. Changing this creates a
	// new transfer request.
	ValueSpecs pulumi.StringMapInput
	// The ID of the zone for which to create the transfer
	// request.
	ZoneId pulumi.StringInput
}

func (TransferRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferRequestArgs)(nil)).Elem()
}

type TransferRequestInput interface {
	pulumi.Input

	ToTransferRequestOutput() TransferRequestOutput
	ToTransferRequestOutputWithContext(ctx context.Context) TransferRequestOutput
}

func (*TransferRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferRequest)(nil)).Elem()
}

func (i *TransferRequest) ToTransferRequestOutput() TransferRequestOutput {
	return i.ToTransferRequestOutputWithContext(context.Background())
}

func (i *TransferRequest) ToTransferRequestOutputWithContext(ctx context.Context) TransferRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferRequestOutput)
}

// TransferRequestArrayInput is an input type that accepts TransferRequestArray and TransferRequestArrayOutput values.
// You can construct a concrete instance of `TransferRequestArrayInput` via:
//
//	TransferRequestArray{ TransferRequestArgs{...} }
type TransferRequestArrayInput interface {
	pulumi.Input

	ToTransferRequestArrayOutput() TransferRequestArrayOutput
	ToTransferRequestArrayOutputWithContext(context.Context) TransferRequestArrayOutput
}

type TransferRequestArray []TransferRequestInput

func (TransferRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferRequest)(nil)).Elem()
}

func (i TransferRequestArray) ToTransferRequestArrayOutput() TransferRequestArrayOutput {
	return i.ToTransferRequestArrayOutputWithContext(context.Background())
}

func (i TransferRequestArray) ToTransferRequestArrayOutputWithContext(ctx context.Context) TransferRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferRequestArrayOutput)
}

// TransferRequestMapInput is an input type that accepts TransferRequestMap and TransferRequestMapOutput values.
// You can construct a concrete instance of `TransferRequestMapInput` via:
//
//	TransferRequestMap{ "key": TransferRequestArgs{...} }
type TransferRequestMapInput interface {
	pulumi.Input

	ToTransferRequestMapOutput() TransferRequestMapOutput
	ToTransferRequestMapOutputWithContext(context.Context) TransferRequestMapOutput
}

type TransferRequestMap map[string]TransferRequestInput

func (TransferRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferRequest)(nil)).Elem()
}

func (i TransferRequestMap) ToTransferRequestMapOutput() TransferRequestMapOutput {
	return i.ToTransferRequestMapOutputWithContext(context.Background())
}

func (i TransferRequestMap) ToTransferRequestMapOutputWithContext(ctx context.Context) TransferRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferRequestMapOutput)
}

type TransferRequestOutput struct{ *pulumi.OutputState }

func (TransferRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferRequest)(nil)).Elem()
}

func (o TransferRequestOutput) ToTransferRequestOutput() TransferRequestOutput {
	return o
}

func (o TransferRequestOutput) ToTransferRequestOutputWithContext(ctx context.Context) TransferRequestOutput {
	return o
}

// A description of the zone tranfer request.
func (o TransferRequestOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Disable wait for zone to reach ACTIVE
// status. The check is enabled by default. If this argument is true, zone
// will be considered as created/updated if OpenStack request returned success.
func (o TransferRequestOutput) DisableStatusCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.BoolPtrOutput { return v.DisableStatusCheck }).(pulumi.BoolPtrOutput)
}

func (o TransferRequestOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 DNS client.
// If omitted, the `region` argument of the provider is used.
// Changing this creates a new DNS zone zone transfer accept.
func (o TransferRequestOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The target Project ID to transfer to.
func (o TransferRequestOutput) TargetProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.StringOutput { return v.TargetProjectId }).(pulumi.StringOutput)
}

// Map of additional options. Changing this creates a
// new transfer request.
func (o TransferRequestOutput) ValueSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.StringMapOutput { return v.ValueSpecs }).(pulumi.StringMapOutput)
}

// The ID of the zone for which to create the transfer
// request.
func (o TransferRequestOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferRequest) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type TransferRequestArrayOutput struct{ *pulumi.OutputState }

func (TransferRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferRequest)(nil)).Elem()
}

func (o TransferRequestArrayOutput) ToTransferRequestArrayOutput() TransferRequestArrayOutput {
	return o
}

func (o TransferRequestArrayOutput) ToTransferRequestArrayOutputWithContext(ctx context.Context) TransferRequestArrayOutput {
	return o
}

func (o TransferRequestArrayOutput) Index(i pulumi.IntInput) TransferRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransferRequest {
		return vs[0].([]*TransferRequest)[vs[1].(int)]
	}).(TransferRequestOutput)
}

type TransferRequestMapOutput struct{ *pulumi.OutputState }

func (TransferRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferRequest)(nil)).Elem()
}

func (o TransferRequestMapOutput) ToTransferRequestMapOutput() TransferRequestMapOutput {
	return o
}

func (o TransferRequestMapOutput) ToTransferRequestMapOutputWithContext(ctx context.Context) TransferRequestMapOutput {
	return o
}

func (o TransferRequestMapOutput) MapIndex(k pulumi.StringInput) TransferRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransferRequest {
		return vs[0].(map[string]*TransferRequest)[vs[1].(string)]
	}).(TransferRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransferRequestInput)(nil)).Elem(), &TransferRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferRequestArrayInput)(nil)).Elem(), TransferRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferRequestMapInput)(nil)).Elem(), TransferRequestMap{})
	pulumi.RegisterOutputType(TransferRequestOutput{})
	pulumi.RegisterOutputType(TransferRequestArrayOutput{})
	pulumi.RegisterOutputType(TransferRequestMapOutput{})
}
