// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package loadbalancer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type MembersMember struct {
	// The IP address of the members to receive traffic from
	// the load balancer.
	Address string `pulumi:"address"`
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// A bool that indicates whether the member is
	// backup. **Requires octavia minor version 2.1 or later**.
	Backup *bool `pulumi:"backup"`
	// The unique ID for the members.
	Id *string `pulumi:"id"`
	// An alternate IP address used for health
	// monitoring a backend member.
	MonitorAddress *string `pulumi:"monitorAddress"`
	// An alternate protocol port used for health
	// monitoring a backend member.
	MonitorPort *int `pulumi:"monitorPort"`
	// Human-readable name for the member.
	Name *string `pulumi:"name"`
	// The port on which to listen for client traffic.
	ProtocolPort int `pulumi:"protocolPort"`
	// The subnet in which to access the member.
	SubnetId *string `pulumi:"subnetId"`
	// A positive integer value that indicates the relative
	// portion of traffic that this members should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight *int `pulumi:"weight"`
}

// MembersMemberInput is an input type that accepts MembersMemberArgs and MembersMemberOutput values.
// You can construct a concrete instance of `MembersMemberInput` via:
//
//	MembersMemberArgs{...}
type MembersMemberInput interface {
	pulumi.Input

	ToMembersMemberOutput() MembersMemberOutput
	ToMembersMemberOutputWithContext(context.Context) MembersMemberOutput
}

type MembersMemberArgs struct {
	// The IP address of the members to receive traffic from
	// the load balancer.
	Address pulumi.StringInput `pulumi:"address"`
	// The administrative state of the member.
	// A valid value is true (UP) or false (DOWN). Defaults to true.
	AdminStateUp pulumi.BoolPtrInput `pulumi:"adminStateUp"`
	// A bool that indicates whether the member is
	// backup. **Requires octavia minor version 2.1 or later**.
	Backup pulumi.BoolPtrInput `pulumi:"backup"`
	// The unique ID for the members.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// An alternate IP address used for health
	// monitoring a backend member.
	MonitorAddress pulumi.StringPtrInput `pulumi:"monitorAddress"`
	// An alternate protocol port used for health
	// monitoring a backend member.
	MonitorPort pulumi.IntPtrInput `pulumi:"monitorPort"`
	// Human-readable name for the member.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The port on which to listen for client traffic.
	ProtocolPort pulumi.IntInput `pulumi:"protocolPort"`
	// The subnet in which to access the member.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// A positive integer value that indicates the relative
	// portion of traffic that this members should receive from the pool. For
	// example, a member with a weight of 10 receives five times as much traffic
	// as a member with a weight of 2. Defaults to 1.
	Weight pulumi.IntPtrInput `pulumi:"weight"`
}

func (MembersMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MembersMember)(nil)).Elem()
}

func (i MembersMemberArgs) ToMembersMemberOutput() MembersMemberOutput {
	return i.ToMembersMemberOutputWithContext(context.Background())
}

func (i MembersMemberArgs) ToMembersMemberOutputWithContext(ctx context.Context) MembersMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembersMemberOutput)
}

// MembersMemberArrayInput is an input type that accepts MembersMemberArray and MembersMemberArrayOutput values.
// You can construct a concrete instance of `MembersMemberArrayInput` via:
//
//	MembersMemberArray{ MembersMemberArgs{...} }
type MembersMemberArrayInput interface {
	pulumi.Input

	ToMembersMemberArrayOutput() MembersMemberArrayOutput
	ToMembersMemberArrayOutputWithContext(context.Context) MembersMemberArrayOutput
}

type MembersMemberArray []MembersMemberInput

func (MembersMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MembersMember)(nil)).Elem()
}

func (i MembersMemberArray) ToMembersMemberArrayOutput() MembersMemberArrayOutput {
	return i.ToMembersMemberArrayOutputWithContext(context.Background())
}

func (i MembersMemberArray) ToMembersMemberArrayOutputWithContext(ctx context.Context) MembersMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembersMemberArrayOutput)
}

type MembersMemberOutput struct{ *pulumi.OutputState }

func (MembersMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MembersMember)(nil)).Elem()
}

func (o MembersMemberOutput) ToMembersMemberOutput() MembersMemberOutput {
	return o
}

func (o MembersMemberOutput) ToMembersMemberOutputWithContext(ctx context.Context) MembersMemberOutput {
	return o
}

// The IP address of the members to receive traffic from
// the load balancer.
func (o MembersMemberOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v MembersMember) string { return v.Address }).(pulumi.StringOutput)
}

// The administrative state of the member.
// A valid value is true (UP) or false (DOWN). Defaults to true.
func (o MembersMemberOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MembersMember) *bool { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// A bool that indicates whether the member is
// backup. **Requires octavia minor version 2.1 or later**.
func (o MembersMemberOutput) Backup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MembersMember) *bool { return v.Backup }).(pulumi.BoolPtrOutput)
}

// The unique ID for the members.
func (o MembersMemberOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MembersMember) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// An alternate IP address used for health
// monitoring a backend member.
func (o MembersMemberOutput) MonitorAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MembersMember) *string { return v.MonitorAddress }).(pulumi.StringPtrOutput)
}

// An alternate protocol port used for health
// monitoring a backend member.
func (o MembersMemberOutput) MonitorPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MembersMember) *int { return v.MonitorPort }).(pulumi.IntPtrOutput)
}

// Human-readable name for the member.
func (o MembersMemberOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MembersMember) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The port on which to listen for client traffic.
func (o MembersMemberOutput) ProtocolPort() pulumi.IntOutput {
	return o.ApplyT(func(v MembersMember) int { return v.ProtocolPort }).(pulumi.IntOutput)
}

// The subnet in which to access the member.
func (o MembersMemberOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MembersMember) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// A positive integer value that indicates the relative
// portion of traffic that this members should receive from the pool. For
// example, a member with a weight of 10 receives five times as much traffic
// as a member with a weight of 2. Defaults to 1.
func (o MembersMemberOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MembersMember) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type MembersMemberArrayOutput struct{ *pulumi.OutputState }

func (MembersMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MembersMember)(nil)).Elem()
}

func (o MembersMemberArrayOutput) ToMembersMemberArrayOutput() MembersMemberArrayOutput {
	return o
}

func (o MembersMemberArrayOutput) ToMembersMemberArrayOutputWithContext(ctx context.Context) MembersMemberArrayOutput {
	return o
}

func (o MembersMemberArrayOutput) Index(i pulumi.IntInput) MembersMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MembersMember {
		return vs[0].([]MembersMember)[vs[1].(int)]
	}).(MembersMemberOutput)
}

type PoolPersistence struct {
	// The name of the cookie if persistence mode is set
	// appropriately. Required if `type = APP_COOKIE`.
	CookieName *string `pulumi:"cookieName"`
	// The type of persistence mode. The current specification
	// supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
	Type string `pulumi:"type"`
}

// PoolPersistenceInput is an input type that accepts PoolPersistenceArgs and PoolPersistenceOutput values.
// You can construct a concrete instance of `PoolPersistenceInput` via:
//
//	PoolPersistenceArgs{...}
type PoolPersistenceInput interface {
	pulumi.Input

	ToPoolPersistenceOutput() PoolPersistenceOutput
	ToPoolPersistenceOutputWithContext(context.Context) PoolPersistenceOutput
}

type PoolPersistenceArgs struct {
	// The name of the cookie if persistence mode is set
	// appropriately. Required if `type = APP_COOKIE`.
	CookieName pulumi.StringPtrInput `pulumi:"cookieName"`
	// The type of persistence mode. The current specification
	// supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
	Type pulumi.StringInput `pulumi:"type"`
}

func (PoolPersistenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolPersistence)(nil)).Elem()
}

func (i PoolPersistenceArgs) ToPoolPersistenceOutput() PoolPersistenceOutput {
	return i.ToPoolPersistenceOutputWithContext(context.Background())
}

func (i PoolPersistenceArgs) ToPoolPersistenceOutputWithContext(ctx context.Context) PoolPersistenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPersistenceOutput)
}

func (i PoolPersistenceArgs) ToPoolPersistencePtrOutput() PoolPersistencePtrOutput {
	return i.ToPoolPersistencePtrOutputWithContext(context.Background())
}

func (i PoolPersistenceArgs) ToPoolPersistencePtrOutputWithContext(ctx context.Context) PoolPersistencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPersistenceOutput).ToPoolPersistencePtrOutputWithContext(ctx)
}

// PoolPersistencePtrInput is an input type that accepts PoolPersistenceArgs, PoolPersistencePtr and PoolPersistencePtrOutput values.
// You can construct a concrete instance of `PoolPersistencePtrInput` via:
//
//	        PoolPersistenceArgs{...}
//
//	or:
//
//	        nil
type PoolPersistencePtrInput interface {
	pulumi.Input

	ToPoolPersistencePtrOutput() PoolPersistencePtrOutput
	ToPoolPersistencePtrOutputWithContext(context.Context) PoolPersistencePtrOutput
}

type poolPersistencePtrType PoolPersistenceArgs

func PoolPersistencePtr(v *PoolPersistenceArgs) PoolPersistencePtrInput {
	return (*poolPersistencePtrType)(v)
}

func (*poolPersistencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolPersistence)(nil)).Elem()
}

func (i *poolPersistencePtrType) ToPoolPersistencePtrOutput() PoolPersistencePtrOutput {
	return i.ToPoolPersistencePtrOutputWithContext(context.Background())
}

func (i *poolPersistencePtrType) ToPoolPersistencePtrOutputWithContext(ctx context.Context) PoolPersistencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolPersistencePtrOutput)
}

type PoolPersistenceOutput struct{ *pulumi.OutputState }

func (PoolPersistenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolPersistence)(nil)).Elem()
}

func (o PoolPersistenceOutput) ToPoolPersistenceOutput() PoolPersistenceOutput {
	return o
}

func (o PoolPersistenceOutput) ToPoolPersistenceOutputWithContext(ctx context.Context) PoolPersistenceOutput {
	return o
}

func (o PoolPersistenceOutput) ToPoolPersistencePtrOutput() PoolPersistencePtrOutput {
	return o.ToPoolPersistencePtrOutputWithContext(context.Background())
}

func (o PoolPersistenceOutput) ToPoolPersistencePtrOutputWithContext(ctx context.Context) PoolPersistencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PoolPersistence) *PoolPersistence {
		return &v
	}).(PoolPersistencePtrOutput)
}

// The name of the cookie if persistence mode is set
// appropriately. Required if `type = APP_COOKIE`.
func (o PoolPersistenceOutput) CookieName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PoolPersistence) *string { return v.CookieName }).(pulumi.StringPtrOutput)
}

// The type of persistence mode. The current specification
// supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
func (o PoolPersistenceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PoolPersistence) string { return v.Type }).(pulumi.StringOutput)
}

type PoolPersistencePtrOutput struct{ *pulumi.OutputState }

func (PoolPersistencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolPersistence)(nil)).Elem()
}

func (o PoolPersistencePtrOutput) ToPoolPersistencePtrOutput() PoolPersistencePtrOutput {
	return o
}

func (o PoolPersistencePtrOutput) ToPoolPersistencePtrOutputWithContext(ctx context.Context) PoolPersistencePtrOutput {
	return o
}

func (o PoolPersistencePtrOutput) Elem() PoolPersistenceOutput {
	return o.ApplyT(func(v *PoolPersistence) PoolPersistence {
		if v != nil {
			return *v
		}
		var ret PoolPersistence
		return ret
	}).(PoolPersistenceOutput)
}

// The name of the cookie if persistence mode is set
// appropriately. Required if `type = APP_COOKIE`.
func (o PoolPersistencePtrOutput) CookieName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PoolPersistence) *string {
		if v == nil {
			return nil
		}
		return v.CookieName
	}).(pulumi.StringPtrOutput)
}

// The type of persistence mode. The current specification
// supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
func (o PoolPersistencePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PoolPersistence) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MembersMemberInput)(nil)).Elem(), MembersMemberArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MembersMemberArrayInput)(nil)).Elem(), MembersMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolPersistenceInput)(nil)).Elem(), PoolPersistenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PoolPersistencePtrInput)(nil)).Elem(), PoolPersistenceArgs{})
	pulumi.RegisterOutputType(MembersMemberOutput{})
	pulumi.RegisterOutputType(MembersMemberArrayOutput{})
	pulumi.RegisterOutputType(PoolPersistenceOutput{})
	pulumi.RegisterOutputType(PoolPersistencePtrOutput{})
}
