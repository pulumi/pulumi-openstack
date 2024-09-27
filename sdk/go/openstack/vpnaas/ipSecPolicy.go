// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpnaas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron IPSec policy resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/vpnaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpnaas.NewIpSecPolicy(ctx, "policy_1", &vpnaas.IpSecPolicyArgs{
//				Name: pulumi.String("my_policy"),
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
// Policies can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:vpnaas/ipSecPolicy:IpSecPolicy policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
// ```
type IpSecPolicy struct {
	pulumi.CustomResourceState

	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm pulumi.StringOutput `pulumi:"authAlgorithm"`
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode pulumi.StringOutput `pulumi:"encapsulationMode"`
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm pulumi.StringOutput `pulumi:"encryptionAlgorithm"`
	// The lifetime of the security association. Consists of Unit and Value.
	Lifetimes IpSecPolicyLifetimeArrayOutput `pulumi:"lifetimes"`
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs pulumi.StringOutput `pulumi:"pfs"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// policy.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol pulumi.StringOutput `pulumi:"transformProtocol"`
	// Map of additional options.
	ValueSpecs pulumi.StringMapOutput `pulumi:"valueSpecs"`
}

// NewIpSecPolicy registers a new resource with the given unique name, arguments, and options.
func NewIpSecPolicy(ctx *pulumi.Context,
	name string, args *IpSecPolicyArgs, opts ...pulumi.ResourceOption) (*IpSecPolicy, error) {
	if args == nil {
		args = &IpSecPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpSecPolicy
	err := ctx.RegisterResource("openstack:vpnaas/ipSecPolicy:IpSecPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSecPolicy gets an existing IpSecPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSecPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSecPolicyState, opts ...pulumi.ResourceOption) (*IpSecPolicy, error) {
	var resource IpSecPolicy
	err := ctx.ReadResource("openstack:vpnaas/ipSecPolicy:IpSecPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSecPolicy resources.
type ipSecPolicyState struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm *string `pulumi:"authAlgorithm"`
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description *string `pulumi:"description"`
	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode *string `pulumi:"encapsulationMode"`
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// The lifetime of the security association. Consists of Unit and Value.
	Lifetimes []IpSecPolicyLifetime `pulumi:"lifetimes"`
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name *string `pulumi:"name"`
	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs *string `pulumi:"pfs"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// policy.
	Region *string `pulumi:"region"`
	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantId *string `pulumi:"tenantId"`
	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol *string `pulumi:"transformProtocol"`
	// Map of additional options.
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
}

type IpSecPolicyState struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm pulumi.StringPtrInput
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description pulumi.StringPtrInput
	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode pulumi.StringPtrInput
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm pulumi.StringPtrInput
	// The lifetime of the security association. Consists of Unit and Value.
	Lifetimes IpSecPolicyLifetimeArrayInput
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name pulumi.StringPtrInput
	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// policy.
	Region pulumi.StringPtrInput
	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantId pulumi.StringPtrInput
	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.StringMapInput
}

func (IpSecPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSecPolicyState)(nil)).Elem()
}

type ipSecPolicyArgs struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm *string `pulumi:"authAlgorithm"`
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description *string `pulumi:"description"`
	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode *string `pulumi:"encapsulationMode"`
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// The lifetime of the security association. Consists of Unit and Value.
	Lifetimes []IpSecPolicyLifetime `pulumi:"lifetimes"`
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name *string `pulumi:"name"`
	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs *string `pulumi:"pfs"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// policy.
	Region *string `pulumi:"region"`
	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantId *string `pulumi:"tenantId"`
	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol *string `pulumi:"transformProtocol"`
	// Map of additional options.
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a IpSecPolicy resource.
type IpSecPolicyArgs struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm pulumi.StringPtrInput
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description pulumi.StringPtrInput
	// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
	// Changing this updates the existing policy.
	EncapsulationMode pulumi.StringPtrInput
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm pulumi.StringPtrInput
	// The lifetime of the security association. Consists of Unit and Value.
	Lifetimes IpSecPolicyLifetimeArrayInput
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name pulumi.StringPtrInput
	// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
	// is group5. Changing this updates the existing policy.
	Pfs pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create an IPSec policy. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// policy.
	Region pulumi.StringPtrInput
	// The owner of the policy. Required if admin wants to
	// create a policy for another project. Changing this creates a new policy.
	TenantId pulumi.StringPtrInput
	// The transform protocol. Valid values are esp, ah and ah-esp.
	// Changing this updates the existing policy. Default is ESP.
	TransformProtocol pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.StringMapInput
}

func (IpSecPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSecPolicyArgs)(nil)).Elem()
}

type IpSecPolicyInput interface {
	pulumi.Input

	ToIpSecPolicyOutput() IpSecPolicyOutput
	ToIpSecPolicyOutputWithContext(ctx context.Context) IpSecPolicyOutput
}

func (*IpSecPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSecPolicy)(nil)).Elem()
}

func (i *IpSecPolicy) ToIpSecPolicyOutput() IpSecPolicyOutput {
	return i.ToIpSecPolicyOutputWithContext(context.Background())
}

func (i *IpSecPolicy) ToIpSecPolicyOutputWithContext(ctx context.Context) IpSecPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecPolicyOutput)
}

// IpSecPolicyArrayInput is an input type that accepts IpSecPolicyArray and IpSecPolicyArrayOutput values.
// You can construct a concrete instance of `IpSecPolicyArrayInput` via:
//
//	IpSecPolicyArray{ IpSecPolicyArgs{...} }
type IpSecPolicyArrayInput interface {
	pulumi.Input

	ToIpSecPolicyArrayOutput() IpSecPolicyArrayOutput
	ToIpSecPolicyArrayOutputWithContext(context.Context) IpSecPolicyArrayOutput
}

type IpSecPolicyArray []IpSecPolicyInput

func (IpSecPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSecPolicy)(nil)).Elem()
}

func (i IpSecPolicyArray) ToIpSecPolicyArrayOutput() IpSecPolicyArrayOutput {
	return i.ToIpSecPolicyArrayOutputWithContext(context.Background())
}

func (i IpSecPolicyArray) ToIpSecPolicyArrayOutputWithContext(ctx context.Context) IpSecPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecPolicyArrayOutput)
}

// IpSecPolicyMapInput is an input type that accepts IpSecPolicyMap and IpSecPolicyMapOutput values.
// You can construct a concrete instance of `IpSecPolicyMapInput` via:
//
//	IpSecPolicyMap{ "key": IpSecPolicyArgs{...} }
type IpSecPolicyMapInput interface {
	pulumi.Input

	ToIpSecPolicyMapOutput() IpSecPolicyMapOutput
	ToIpSecPolicyMapOutputWithContext(context.Context) IpSecPolicyMapOutput
}

type IpSecPolicyMap map[string]IpSecPolicyInput

func (IpSecPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSecPolicy)(nil)).Elem()
}

func (i IpSecPolicyMap) ToIpSecPolicyMapOutput() IpSecPolicyMapOutput {
	return i.ToIpSecPolicyMapOutputWithContext(context.Background())
}

func (i IpSecPolicyMap) ToIpSecPolicyMapOutputWithContext(ctx context.Context) IpSecPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecPolicyMapOutput)
}

type IpSecPolicyOutput struct{ *pulumi.OutputState }

func (IpSecPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSecPolicy)(nil)).Elem()
}

func (o IpSecPolicyOutput) ToIpSecPolicyOutput() IpSecPolicyOutput {
	return o
}

func (o IpSecPolicyOutput) ToIpSecPolicyOutputWithContext(ctx context.Context) IpSecPolicyOutput {
	return o
}

// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
// Default is sha1. Changing this updates the algorithm of the existing policy.
func (o IpSecPolicyOutput) AuthAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.AuthAlgorithm }).(pulumi.StringOutput)
}

// The human-readable description for the policy.
// Changing this updates the description of the existing policy.
func (o IpSecPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
// Changing this updates the existing policy.
func (o IpSecPolicyOutput) EncapsulationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.EncapsulationMode }).(pulumi.StringOutput)
}

// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
// The default value is aes-128. Changing this updates the existing policy.
func (o IpSecPolicyOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

// The lifetime of the security association. Consists of Unit and Value.
func (o IpSecPolicyOutput) Lifetimes() IpSecPolicyLifetimeArrayOutput {
	return o.ApplyT(func(v *IpSecPolicy) IpSecPolicyLifetimeArrayOutput { return v.Lifetimes }).(IpSecPolicyLifetimeArrayOutput)
}

// The name of the policy. Changing this updates the name of
// the existing policy.
func (o IpSecPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
// is group5. Changing this updates the existing policy.
func (o IpSecPolicyOutput) Pfs() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.Pfs }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create an IPSec policy. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// policy.
func (o IpSecPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The owner of the policy. Required if admin wants to
// create a policy for another project. Changing this creates a new policy.
func (o IpSecPolicyOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The transform protocol. Valid values are esp, ah and ah-esp.
// Changing this updates the existing policy. Default is ESP.
func (o IpSecPolicyOutput) TransformProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringOutput { return v.TransformProtocol }).(pulumi.StringOutput)
}

// Map of additional options.
func (o IpSecPolicyOutput) ValueSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpSecPolicy) pulumi.StringMapOutput { return v.ValueSpecs }).(pulumi.StringMapOutput)
}

type IpSecPolicyArrayOutput struct{ *pulumi.OutputState }

func (IpSecPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSecPolicy)(nil)).Elem()
}

func (o IpSecPolicyArrayOutput) ToIpSecPolicyArrayOutput() IpSecPolicyArrayOutput {
	return o
}

func (o IpSecPolicyArrayOutput) ToIpSecPolicyArrayOutputWithContext(ctx context.Context) IpSecPolicyArrayOutput {
	return o
}

func (o IpSecPolicyArrayOutput) Index(i pulumi.IntInput) IpSecPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpSecPolicy {
		return vs[0].([]*IpSecPolicy)[vs[1].(int)]
	}).(IpSecPolicyOutput)
}

type IpSecPolicyMapOutput struct{ *pulumi.OutputState }

func (IpSecPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSecPolicy)(nil)).Elem()
}

func (o IpSecPolicyMapOutput) ToIpSecPolicyMapOutput() IpSecPolicyMapOutput {
	return o
}

func (o IpSecPolicyMapOutput) ToIpSecPolicyMapOutputWithContext(ctx context.Context) IpSecPolicyMapOutput {
	return o
}

func (o IpSecPolicyMapOutput) MapIndex(k pulumi.StringInput) IpSecPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpSecPolicy {
		return vs[0].(map[string]*IpSecPolicy)[vs[1].(string)]
	}).(IpSecPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpSecPolicyInput)(nil)).Elem(), &IpSecPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSecPolicyArrayInput)(nil)).Elem(), IpSecPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSecPolicyMapInput)(nil)).Elem(), IpSecPolicyMap{})
	pulumi.RegisterOutputType(IpSecPolicyOutput{})
	pulumi.RegisterOutputType(IpSecPolicyArrayOutput{})
	pulumi.RegisterOutputType(IpSecPolicyMapOutput{})
}
