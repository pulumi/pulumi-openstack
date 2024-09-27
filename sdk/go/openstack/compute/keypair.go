// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### Import an Existing Public Key
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewKeypair(ctx, "test-keypair", &compute.KeypairArgs{
//				Name:      pulumi.String("my-keypair"),
//				PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB"),
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
// ### Generate a Public/Private Key Pair
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.NewKeypair(ctx, "test-keypair", &compute.KeypairArgs{
//				Name: pulumi.String("my-keypair"),
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
// Keypairs can be imported using the `name`, e.g.
//
// ```sh
// $ pulumi import openstack:compute/keypair:Keypair my-keypair test-keypair
// ```
type Keypair struct {
	pulumi.CustomResourceState

	// The fingerprint of the public key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// A unique name for the keypair. Changing this creates a new
	// keypair.
	Name pulumi.StringOutput `pulumi:"name"`
	// The generated private key when no public key is specified.
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new keypair.
	Region pulumi.StringOutput `pulumi:"region"`
	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Map of additional options.
	ValueSpecs pulumi.StringMapOutput `pulumi:"valueSpecs"`
}

// NewKeypair registers a new resource with the given unique name, arguments, and options.
func NewKeypair(ctx *pulumi.Context,
	name string, args *KeypairArgs, opts ...pulumi.ResourceOption) (*Keypair, error) {
	if args == nil {
		args = &KeypairArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Keypair
	err := ctx.RegisterResource("openstack:compute/keypair:Keypair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeypair gets an existing Keypair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeypair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeypairState, opts ...pulumi.ResourceOption) (*Keypair, error) {
	var resource Keypair
	err := ctx.ReadResource("openstack:compute/keypair:Keypair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Keypair resources.
type keypairState struct {
	// The fingerprint of the public key.
	Fingerprint *string `pulumi:"fingerprint"`
	// A unique name for the keypair. Changing this creates a new
	// keypair.
	Name *string `pulumi:"name"`
	// The generated private key when no public key is specified.
	PrivateKey *string `pulumi:"privateKey"`
	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	PublicKey *string `pulumi:"publicKey"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new keypair.
	Region *string `pulumi:"region"`
	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	UserId *string `pulumi:"userId"`
	// Map of additional options.
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
}

type KeypairState struct {
	// The fingerprint of the public key.
	Fingerprint pulumi.StringPtrInput
	// A unique name for the keypair. Changing this creates a new
	// keypair.
	Name pulumi.StringPtrInput
	// The generated private key when no public key is specified.
	PrivateKey pulumi.StringPtrInput
	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	PublicKey pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new keypair.
	Region pulumi.StringPtrInput
	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	UserId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.StringMapInput
}

func (KeypairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keypairState)(nil)).Elem()
}

type keypairArgs struct {
	// A unique name for the keypair. Changing this creates a new
	// keypair.
	Name *string `pulumi:"name"`
	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	PublicKey *string `pulumi:"publicKey"`
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new keypair.
	Region *string `pulumi:"region"`
	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	UserId *string `pulumi:"userId"`
	// Map of additional options.
	ValueSpecs map[string]string `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a Keypair resource.
type KeypairArgs struct {
	// A unique name for the keypair. Changing this creates a new
	// keypair.
	Name pulumi.StringPtrInput
	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	PublicKey pulumi.StringPtrInput
	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the `region` argument of the provider is used.
	// Changing this creates a new keypair.
	Region pulumi.StringPtrInput
	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	UserId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.StringMapInput
}

func (KeypairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keypairArgs)(nil)).Elem()
}

type KeypairInput interface {
	pulumi.Input

	ToKeypairOutput() KeypairOutput
	ToKeypairOutputWithContext(ctx context.Context) KeypairOutput
}

func (*Keypair) ElementType() reflect.Type {
	return reflect.TypeOf((**Keypair)(nil)).Elem()
}

func (i *Keypair) ToKeypairOutput() KeypairOutput {
	return i.ToKeypairOutputWithContext(context.Background())
}

func (i *Keypair) ToKeypairOutputWithContext(ctx context.Context) KeypairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeypairOutput)
}

// KeypairArrayInput is an input type that accepts KeypairArray and KeypairArrayOutput values.
// You can construct a concrete instance of `KeypairArrayInput` via:
//
//	KeypairArray{ KeypairArgs{...} }
type KeypairArrayInput interface {
	pulumi.Input

	ToKeypairArrayOutput() KeypairArrayOutput
	ToKeypairArrayOutputWithContext(context.Context) KeypairArrayOutput
}

type KeypairArray []KeypairInput

func (KeypairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Keypair)(nil)).Elem()
}

func (i KeypairArray) ToKeypairArrayOutput() KeypairArrayOutput {
	return i.ToKeypairArrayOutputWithContext(context.Background())
}

func (i KeypairArray) ToKeypairArrayOutputWithContext(ctx context.Context) KeypairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeypairArrayOutput)
}

// KeypairMapInput is an input type that accepts KeypairMap and KeypairMapOutput values.
// You can construct a concrete instance of `KeypairMapInput` via:
//
//	KeypairMap{ "key": KeypairArgs{...} }
type KeypairMapInput interface {
	pulumi.Input

	ToKeypairMapOutput() KeypairMapOutput
	ToKeypairMapOutputWithContext(context.Context) KeypairMapOutput
}

type KeypairMap map[string]KeypairInput

func (KeypairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Keypair)(nil)).Elem()
}

func (i KeypairMap) ToKeypairMapOutput() KeypairMapOutput {
	return i.ToKeypairMapOutputWithContext(context.Background())
}

func (i KeypairMap) ToKeypairMapOutputWithContext(ctx context.Context) KeypairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeypairMapOutput)
}

type KeypairOutput struct{ *pulumi.OutputState }

func (KeypairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Keypair)(nil)).Elem()
}

func (o KeypairOutput) ToKeypairOutput() KeypairOutput {
	return o
}

func (o KeypairOutput) ToKeypairOutputWithContext(ctx context.Context) KeypairOutput {
	return o
}

// The fingerprint of the public key.
func (o KeypairOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// A unique name for the keypair. Changing this creates a new
// keypair.
func (o KeypairOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The generated private key when no public key is specified.
func (o KeypairOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

// A pregenerated OpenSSH-formatted public key.
// Changing this creates a new keypair. If a public key is not specified, then
// a public/private key pair will be automatically generated. If a pair is
// created, then destroying this resource means you will lose access to that
// keypair forever.
func (o KeypairOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Compute client.
// Keypairs are associated with accounts, but a Compute client is needed to
// create one. If omitted, the `region` argument of the provider is used.
// Changing this creates a new keypair.
func (o KeypairOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// This allows administrative users to operate key-pairs
// of specified user ID. For this feature your need to have openstack microversion
// 2.10 (Liberty) or later.
func (o KeypairOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// Map of additional options.
func (o KeypairOutput) ValueSpecs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Keypair) pulumi.StringMapOutput { return v.ValueSpecs }).(pulumi.StringMapOutput)
}

type KeypairArrayOutput struct{ *pulumi.OutputState }

func (KeypairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Keypair)(nil)).Elem()
}

func (o KeypairArrayOutput) ToKeypairArrayOutput() KeypairArrayOutput {
	return o
}

func (o KeypairArrayOutput) ToKeypairArrayOutputWithContext(ctx context.Context) KeypairArrayOutput {
	return o
}

func (o KeypairArrayOutput) Index(i pulumi.IntInput) KeypairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Keypair {
		return vs[0].([]*Keypair)[vs[1].(int)]
	}).(KeypairOutput)
}

type KeypairMapOutput struct{ *pulumi.OutputState }

func (KeypairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Keypair)(nil)).Elem()
}

func (o KeypairMapOutput) ToKeypairMapOutput() KeypairMapOutput {
	return o
}

func (o KeypairMapOutput) ToKeypairMapOutputWithContext(ctx context.Context) KeypairMapOutput {
	return o
}

func (o KeypairMapOutput) MapIndex(k pulumi.StringInput) KeypairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Keypair {
		return vs[0].(map[string]*Keypair)[vs[1].(string)]
	}).(KeypairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeypairInput)(nil)).Elem(), &Keypair{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeypairArrayInput)(nil)).Elem(), KeypairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeypairMapInput)(nil)).Elem(), KeypairMap{})
	pulumi.RegisterOutputType(KeypairOutput{})
	pulumi.RegisterOutputType(KeypairArrayOutput{})
	pulumi.RegisterOutputType(KeypairMapOutput{})
}
