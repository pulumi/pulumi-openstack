// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IkePolicyLifetime struct {
	Units *string `pulumi:"units"`
	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Value *int `pulumi:"value"`
}

// IkePolicyLifetimeInput is an input type that accepts IkePolicyLifetimeArgs and IkePolicyLifetimeOutput values.
// You can construct a concrete instance of `IkePolicyLifetimeInput` via:
//
//          IkePolicyLifetimeArgs{...}
type IkePolicyLifetimeInput interface {
	pulumi.Input

	ToIkePolicyLifetimeOutput() IkePolicyLifetimeOutput
	ToIkePolicyLifetimeOutputWithContext(context.Context) IkePolicyLifetimeOutput
}

type IkePolicyLifetimeArgs struct {
	Units pulumi.StringPtrInput `pulumi:"units"`
	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Value pulumi.IntPtrInput `pulumi:"value"`
}

func (IkePolicyLifetimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IkePolicyLifetime)(nil)).Elem()
}

func (i IkePolicyLifetimeArgs) ToIkePolicyLifetimeOutput() IkePolicyLifetimeOutput {
	return i.ToIkePolicyLifetimeOutputWithContext(context.Background())
}

func (i IkePolicyLifetimeArgs) ToIkePolicyLifetimeOutputWithContext(ctx context.Context) IkePolicyLifetimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IkePolicyLifetimeOutput)
}

// IkePolicyLifetimeArrayInput is an input type that accepts IkePolicyLifetimeArray and IkePolicyLifetimeArrayOutput values.
// You can construct a concrete instance of `IkePolicyLifetimeArrayInput` via:
//
//          IkePolicyLifetimeArray{ IkePolicyLifetimeArgs{...} }
type IkePolicyLifetimeArrayInput interface {
	pulumi.Input

	ToIkePolicyLifetimeArrayOutput() IkePolicyLifetimeArrayOutput
	ToIkePolicyLifetimeArrayOutputWithContext(context.Context) IkePolicyLifetimeArrayOutput
}

type IkePolicyLifetimeArray []IkePolicyLifetimeInput

func (IkePolicyLifetimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IkePolicyLifetime)(nil)).Elem()
}

func (i IkePolicyLifetimeArray) ToIkePolicyLifetimeArrayOutput() IkePolicyLifetimeArrayOutput {
	return i.ToIkePolicyLifetimeArrayOutputWithContext(context.Background())
}

func (i IkePolicyLifetimeArray) ToIkePolicyLifetimeArrayOutputWithContext(ctx context.Context) IkePolicyLifetimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IkePolicyLifetimeArrayOutput)
}

type IkePolicyLifetimeOutput struct{ *pulumi.OutputState }

func (IkePolicyLifetimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IkePolicyLifetime)(nil)).Elem()
}

func (o IkePolicyLifetimeOutput) ToIkePolicyLifetimeOutput() IkePolicyLifetimeOutput {
	return o
}

func (o IkePolicyLifetimeOutput) ToIkePolicyLifetimeOutputWithContext(ctx context.Context) IkePolicyLifetimeOutput {
	return o
}

func (o IkePolicyLifetimeOutput) Units() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IkePolicyLifetime) *string { return v.Units }).(pulumi.StringPtrOutput)
}

// The value for the lifetime of the security association. Must be a positive integer.
// Default is 3600.
func (o IkePolicyLifetimeOutput) Value() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IkePolicyLifetime) *int { return v.Value }).(pulumi.IntPtrOutput)
}

type IkePolicyLifetimeArrayOutput struct{ *pulumi.OutputState }

func (IkePolicyLifetimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IkePolicyLifetime)(nil)).Elem()
}

func (o IkePolicyLifetimeArrayOutput) ToIkePolicyLifetimeArrayOutput() IkePolicyLifetimeArrayOutput {
	return o
}

func (o IkePolicyLifetimeArrayOutput) ToIkePolicyLifetimeArrayOutputWithContext(ctx context.Context) IkePolicyLifetimeArrayOutput {
	return o
}

func (o IkePolicyLifetimeArrayOutput) Index(i pulumi.IntInput) IkePolicyLifetimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IkePolicyLifetime {
		return vs[0].([]IkePolicyLifetime)[vs[1].(int)]
	}).(IkePolicyLifetimeOutput)
}

type IpSecPolicyLifetime struct {
	Units *string `pulumi:"units"`
	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Value *int `pulumi:"value"`
}

// IpSecPolicyLifetimeInput is an input type that accepts IpSecPolicyLifetimeArgs and IpSecPolicyLifetimeOutput values.
// You can construct a concrete instance of `IpSecPolicyLifetimeInput` via:
//
//          IpSecPolicyLifetimeArgs{...}
type IpSecPolicyLifetimeInput interface {
	pulumi.Input

	ToIpSecPolicyLifetimeOutput() IpSecPolicyLifetimeOutput
	ToIpSecPolicyLifetimeOutputWithContext(context.Context) IpSecPolicyLifetimeOutput
}

type IpSecPolicyLifetimeArgs struct {
	Units pulumi.StringPtrInput `pulumi:"units"`
	// The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Value pulumi.IntPtrInput `pulumi:"value"`
}

func (IpSecPolicyLifetimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecPolicyLifetime)(nil)).Elem()
}

func (i IpSecPolicyLifetimeArgs) ToIpSecPolicyLifetimeOutput() IpSecPolicyLifetimeOutput {
	return i.ToIpSecPolicyLifetimeOutputWithContext(context.Background())
}

func (i IpSecPolicyLifetimeArgs) ToIpSecPolicyLifetimeOutputWithContext(ctx context.Context) IpSecPolicyLifetimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecPolicyLifetimeOutput)
}

// IpSecPolicyLifetimeArrayInput is an input type that accepts IpSecPolicyLifetimeArray and IpSecPolicyLifetimeArrayOutput values.
// You can construct a concrete instance of `IpSecPolicyLifetimeArrayInput` via:
//
//          IpSecPolicyLifetimeArray{ IpSecPolicyLifetimeArgs{...} }
type IpSecPolicyLifetimeArrayInput interface {
	pulumi.Input

	ToIpSecPolicyLifetimeArrayOutput() IpSecPolicyLifetimeArrayOutput
	ToIpSecPolicyLifetimeArrayOutputWithContext(context.Context) IpSecPolicyLifetimeArrayOutput
}

type IpSecPolicyLifetimeArray []IpSecPolicyLifetimeInput

func (IpSecPolicyLifetimeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecPolicyLifetime)(nil)).Elem()
}

func (i IpSecPolicyLifetimeArray) ToIpSecPolicyLifetimeArrayOutput() IpSecPolicyLifetimeArrayOutput {
	return i.ToIpSecPolicyLifetimeArrayOutputWithContext(context.Background())
}

func (i IpSecPolicyLifetimeArray) ToIpSecPolicyLifetimeArrayOutputWithContext(ctx context.Context) IpSecPolicyLifetimeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecPolicyLifetimeArrayOutput)
}

type IpSecPolicyLifetimeOutput struct{ *pulumi.OutputState }

func (IpSecPolicyLifetimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecPolicyLifetime)(nil)).Elem()
}

func (o IpSecPolicyLifetimeOutput) ToIpSecPolicyLifetimeOutput() IpSecPolicyLifetimeOutput {
	return o
}

func (o IpSecPolicyLifetimeOutput) ToIpSecPolicyLifetimeOutputWithContext(ctx context.Context) IpSecPolicyLifetimeOutput {
	return o
}

func (o IpSecPolicyLifetimeOutput) Units() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecPolicyLifetime) *string { return v.Units }).(pulumi.StringPtrOutput)
}

// The value for the lifetime of the security association. Must be a positive integer.
// Default is 3600.
func (o IpSecPolicyLifetimeOutput) Value() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecPolicyLifetime) *int { return v.Value }).(pulumi.IntPtrOutput)
}

type IpSecPolicyLifetimeArrayOutput struct{ *pulumi.OutputState }

func (IpSecPolicyLifetimeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecPolicyLifetime)(nil)).Elem()
}

func (o IpSecPolicyLifetimeArrayOutput) ToIpSecPolicyLifetimeArrayOutput() IpSecPolicyLifetimeArrayOutput {
	return o
}

func (o IpSecPolicyLifetimeArrayOutput) ToIpSecPolicyLifetimeArrayOutputWithContext(ctx context.Context) IpSecPolicyLifetimeArrayOutput {
	return o
}

func (o IpSecPolicyLifetimeArrayOutput) Index(i pulumi.IntInput) IpSecPolicyLifetimeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecPolicyLifetime {
		return vs[0].([]IpSecPolicyLifetime)[vs[1].(int)]
	}).(IpSecPolicyLifetimeOutput)
}

type SiteConnectionDpd struct {
	// The dead peer detection (DPD) action.
	// A valid value is clear, hold, restart, disabled, or restart-by-peer.
	// Default value is hold.
	Action *string `pulumi:"action"`
	// The dead peer detection (DPD) interval, in seconds.
	// A valid value is a positive integer.
	// Default is 30.
	Interval *int `pulumi:"interval"`
	// The dead peer detection (DPD) timeout in seconds.
	// A valid value is a positive integer that is greater than the DPD interval value.
	// Default is 120.
	Timeout *int `pulumi:"timeout"`
}

// SiteConnectionDpdInput is an input type that accepts SiteConnectionDpdArgs and SiteConnectionDpdOutput values.
// You can construct a concrete instance of `SiteConnectionDpdInput` via:
//
//          SiteConnectionDpdArgs{...}
type SiteConnectionDpdInput interface {
	pulumi.Input

	ToSiteConnectionDpdOutput() SiteConnectionDpdOutput
	ToSiteConnectionDpdOutputWithContext(context.Context) SiteConnectionDpdOutput
}

type SiteConnectionDpdArgs struct {
	// The dead peer detection (DPD) action.
	// A valid value is clear, hold, restart, disabled, or restart-by-peer.
	// Default value is hold.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// The dead peer detection (DPD) interval, in seconds.
	// A valid value is a positive integer.
	// Default is 30.
	Interval pulumi.IntPtrInput `pulumi:"interval"`
	// The dead peer detection (DPD) timeout in seconds.
	// A valid value is a positive integer that is greater than the DPD interval value.
	// Default is 120.
	Timeout pulumi.IntPtrInput `pulumi:"timeout"`
}

func (SiteConnectionDpdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConnectionDpd)(nil)).Elem()
}

func (i SiteConnectionDpdArgs) ToSiteConnectionDpdOutput() SiteConnectionDpdOutput {
	return i.ToSiteConnectionDpdOutputWithContext(context.Background())
}

func (i SiteConnectionDpdArgs) ToSiteConnectionDpdOutputWithContext(ctx context.Context) SiteConnectionDpdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionDpdOutput)
}

// SiteConnectionDpdArrayInput is an input type that accepts SiteConnectionDpdArray and SiteConnectionDpdArrayOutput values.
// You can construct a concrete instance of `SiteConnectionDpdArrayInput` via:
//
//          SiteConnectionDpdArray{ SiteConnectionDpdArgs{...} }
type SiteConnectionDpdArrayInput interface {
	pulumi.Input

	ToSiteConnectionDpdArrayOutput() SiteConnectionDpdArrayOutput
	ToSiteConnectionDpdArrayOutputWithContext(context.Context) SiteConnectionDpdArrayOutput
}

type SiteConnectionDpdArray []SiteConnectionDpdInput

func (SiteConnectionDpdArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteConnectionDpd)(nil)).Elem()
}

func (i SiteConnectionDpdArray) ToSiteConnectionDpdArrayOutput() SiteConnectionDpdArrayOutput {
	return i.ToSiteConnectionDpdArrayOutputWithContext(context.Background())
}

func (i SiteConnectionDpdArray) ToSiteConnectionDpdArrayOutputWithContext(ctx context.Context) SiteConnectionDpdArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionDpdArrayOutput)
}

type SiteConnectionDpdOutput struct{ *pulumi.OutputState }

func (SiteConnectionDpdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConnectionDpd)(nil)).Elem()
}

func (o SiteConnectionDpdOutput) ToSiteConnectionDpdOutput() SiteConnectionDpdOutput {
	return o
}

func (o SiteConnectionDpdOutput) ToSiteConnectionDpdOutputWithContext(ctx context.Context) SiteConnectionDpdOutput {
	return o
}

// The dead peer detection (DPD) action.
// A valid value is clear, hold, restart, disabled, or restart-by-peer.
// Default value is hold.
func (o SiteConnectionDpdOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConnectionDpd) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// The dead peer detection (DPD) interval, in seconds.
// A valid value is a positive integer.
// Default is 30.
func (o SiteConnectionDpdOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConnectionDpd) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

// The dead peer detection (DPD) timeout in seconds.
// A valid value is a positive integer that is greater than the DPD interval value.
// Default is 120.
func (o SiteConnectionDpdOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConnectionDpd) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

type SiteConnectionDpdArrayOutput struct{ *pulumi.OutputState }

func (SiteConnectionDpdArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteConnectionDpd)(nil)).Elem()
}

func (o SiteConnectionDpdArrayOutput) ToSiteConnectionDpdArrayOutput() SiteConnectionDpdArrayOutput {
	return o
}

func (o SiteConnectionDpdArrayOutput) ToSiteConnectionDpdArrayOutputWithContext(ctx context.Context) SiteConnectionDpdArrayOutput {
	return o
}

func (o SiteConnectionDpdArrayOutput) Index(i pulumi.IntInput) SiteConnectionDpdOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SiteConnectionDpd {
		return vs[0].([]SiteConnectionDpd)[vs[1].(int)]
	}).(SiteConnectionDpdOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IkePolicyLifetimeInput)(nil)).Elem(), IkePolicyLifetimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IkePolicyLifetimeArrayInput)(nil)).Elem(), IkePolicyLifetimeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSecPolicyLifetimeInput)(nil)).Elem(), IpSecPolicyLifetimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSecPolicyLifetimeArrayInput)(nil)).Elem(), IpSecPolicyLifetimeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteConnectionDpdInput)(nil)).Elem(), SiteConnectionDpdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SiteConnectionDpdArrayInput)(nil)).Elem(), SiteConnectionDpdArray{})
	pulumi.RegisterOutputType(IkePolicyLifetimeOutput{})
	pulumi.RegisterOutputType(IkePolicyLifetimeArrayOutput{})
	pulumi.RegisterOutputType(IpSecPolicyLifetimeOutput{})
	pulumi.RegisterOutputType(IpSecPolicyLifetimeArrayOutput{})
	pulumi.RegisterOutputType(SiteConnectionDpdOutput{})
	pulumi.RegisterOutputType(SiteConnectionDpdArrayOutput{})
}
