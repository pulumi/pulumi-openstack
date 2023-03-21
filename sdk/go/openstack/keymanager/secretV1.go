// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Simple secret
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/keymanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keymanager.NewSecretV1(ctx, "secret1", &keymanager.SecretV1Args{
//				Algorithm: pulumi.String("aes"),
//				BitLength: pulumi.Int(256),
//				Metadata: pulumi.AnyMap{
//					"key": pulumi.Any("foo"),
//				},
//				Mode:               pulumi.String("cbc"),
//				Payload:            pulumi.String("foobar"),
//				PayloadContentType: pulumi.String("text/plain"),
//				SecretType:         pulumi.String("passphrase"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Secret with the ACL
//
// > **Note** Only read ACLs are supported
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/keymanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keymanager.NewSecretV1(ctx, "secret1", &keymanager.SecretV1Args{
//				Acl: &keymanager.SecretV1AclArgs{
//					Read: &keymanager.SecretV1AclReadArgs{
//						ProjectAccess: pulumi.Bool(false),
//						Users: pulumi.StringArray{
//							pulumi.String("userid1"),
//							pulumi.String("userid2"),
//						},
//					},
//				},
//				Payload:            readFileOrPanic("certificate.pem"),
//				PayloadContentType: pulumi.String("text/plain"),
//				SecretType:         pulumi.String("certificate"),
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
// Secrets can be imported using the secret id (the last part of the secret reference), e.g.
//
// ```sh
//
//	$ pulumi import openstack:keymanager/secretV1:SecretV1 secret_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
//
// ```
type SecretV1 struct {
	pulumi.CustomResourceState

	// Allows to control an access to a secret. Currently only the
	// `read` operation is supported. If not specified, the secret is accessible
	// project wide.
	Acl SecretV1AclOutput `pulumi:"acl"`
	// Metadata provided by a user or system for informational purposes.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// The map of metadata, assigned on the secret, which has been
	// explicitly and implicitly added.
	AllMetadata pulumi.MapOutput `pulumi:"allMetadata"`
	// Metadata provided by a user or system for informational purposes.
	BitLength pulumi.IntOutput `pulumi:"bitLength"`
	// The map of the content types, assigned on the secret.
	ContentTypes pulumi.MapOutput `pulumi:"contentTypes"`
	// The date the secret ACL was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The creator of the secret.
	CreatorId pulumi.StringOutput `pulumi:"creatorId"`
	// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration pulumi.StringPtrOutput `pulumi:"expiration"`
	// Additional Metadata for the secret.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Metadata provided by a user or system for informational purposes.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name pulumi.StringOutput `pulumi:"name"`
	// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
	Payload pulumi.StringPtrOutput `pulumi:"payload"`
	// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
	PayloadContentEncoding pulumi.StringPtrOutput `pulumi:"payloadContentEncoding"`
	// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
	PayloadContentType pulumi.StringPtrOutput `pulumi:"payloadContentType"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region pulumi.StringOutput `pulumi:"region"`
	// The secret reference / where to find the secret.
	SecretRef pulumi.StringOutput `pulumi:"secretRef"`
	// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
	SecretType pulumi.StringOutput `pulumi:"secretType"`
	// The status of the secret.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date the secret ACL was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewSecretV1 registers a new resource with the given unique name, arguments, and options.
func NewSecretV1(ctx *pulumi.Context,
	name string, args *SecretV1Args, opts ...pulumi.ResourceOption) (*SecretV1, error) {
	if args == nil {
		args = &SecretV1Args{}
	}

	if args.Payload != nil {
		args.Payload = pulumi.ToSecret(args.Payload).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"payload",
	})
	opts = append(opts, secrets)
	var resource SecretV1
	err := ctx.RegisterResource("openstack:keymanager/secretV1:SecretV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretV1 gets an existing SecretV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretV1State, opts ...pulumi.ResourceOption) (*SecretV1, error) {
	var resource SecretV1
	err := ctx.ReadResource("openstack:keymanager/secretV1:SecretV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretV1 resources.
type secretV1State struct {
	// Allows to control an access to a secret. Currently only the
	// `read` operation is supported. If not specified, the secret is accessible
	// project wide.
	Acl *SecretV1Acl `pulumi:"acl"`
	// Metadata provided by a user or system for informational purposes.
	Algorithm *string `pulumi:"algorithm"`
	// The map of metadata, assigned on the secret, which has been
	// explicitly and implicitly added.
	AllMetadata map[string]interface{} `pulumi:"allMetadata"`
	// Metadata provided by a user or system for informational purposes.
	BitLength *int `pulumi:"bitLength"`
	// The map of the content types, assigned on the secret.
	ContentTypes map[string]interface{} `pulumi:"contentTypes"`
	// The date the secret ACL was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The creator of the secret.
	CreatorId *string `pulumi:"creatorId"`
	// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration *string `pulumi:"expiration"`
	// Additional Metadata for the secret.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata provided by a user or system for informational purposes.
	Mode *string `pulumi:"mode"`
	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
	Payload *string `pulumi:"payload"`
	// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
	PayloadContentEncoding *string `pulumi:"payloadContentEncoding"`
	// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
	PayloadContentType *string `pulumi:"payloadContentType"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region *string `pulumi:"region"`
	// The secret reference / where to find the secret.
	SecretRef *string `pulumi:"secretRef"`
	// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
	SecretType *string `pulumi:"secretType"`
	// The status of the secret.
	Status *string `pulumi:"status"`
	// The date the secret ACL was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type SecretV1State struct {
	// Allows to control an access to a secret. Currently only the
	// `read` operation is supported. If not specified, the secret is accessible
	// project wide.
	Acl SecretV1AclPtrInput
	// Metadata provided by a user or system for informational purposes.
	Algorithm pulumi.StringPtrInput
	// The map of metadata, assigned on the secret, which has been
	// explicitly and implicitly added.
	AllMetadata pulumi.MapInput
	// Metadata provided by a user or system for informational purposes.
	BitLength pulumi.IntPtrInput
	// The map of the content types, assigned on the secret.
	ContentTypes pulumi.MapInput
	// The date the secret ACL was created.
	CreatedAt pulumi.StringPtrInput
	// The creator of the secret.
	CreatorId pulumi.StringPtrInput
	// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration pulumi.StringPtrInput
	// Additional Metadata for the secret.
	Metadata pulumi.MapInput
	// Metadata provided by a user or system for informational purposes.
	Mode pulumi.StringPtrInput
	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
	Payload pulumi.StringPtrInput
	// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
	PayloadContentEncoding pulumi.StringPtrInput
	// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
	PayloadContentType pulumi.StringPtrInput
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region pulumi.StringPtrInput
	// The secret reference / where to find the secret.
	SecretRef pulumi.StringPtrInput
	// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
	SecretType pulumi.StringPtrInput
	// The status of the secret.
	Status pulumi.StringPtrInput
	// The date the secret ACL was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (SecretV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*secretV1State)(nil)).Elem()
}

type secretV1Args struct {
	// Allows to control an access to a secret. Currently only the
	// `read` operation is supported. If not specified, the secret is accessible
	// project wide.
	Acl *SecretV1Acl `pulumi:"acl"`
	// Metadata provided by a user or system for informational purposes.
	Algorithm *string `pulumi:"algorithm"`
	// Metadata provided by a user or system for informational purposes.
	BitLength *int `pulumi:"bitLength"`
	// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration *string `pulumi:"expiration"`
	// Additional Metadata for the secret.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Metadata provided by a user or system for informational purposes.
	Mode *string `pulumi:"mode"`
	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name *string `pulumi:"name"`
	// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
	Payload *string `pulumi:"payload"`
	// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
	PayloadContentEncoding *string `pulumi:"payloadContentEncoding"`
	// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
	PayloadContentType *string `pulumi:"payloadContentType"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region *string `pulumi:"region"`
	// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
	SecretType *string `pulumi:"secretType"`
}

// The set of arguments for constructing a SecretV1 resource.
type SecretV1Args struct {
	// Allows to control an access to a secret. Currently only the
	// `read` operation is supported. If not specified, the secret is accessible
	// project wide.
	Acl SecretV1AclPtrInput
	// Metadata provided by a user or system for informational purposes.
	Algorithm pulumi.StringPtrInput
	// Metadata provided by a user or system for informational purposes.
	BitLength pulumi.IntPtrInput
	// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
	Expiration pulumi.StringPtrInput
	// Additional Metadata for the secret.
	Metadata pulumi.MapInput
	// Metadata provided by a user or system for informational purposes.
	Mode pulumi.StringPtrInput
	// Human-readable name for the Secret. Does not have
	// to be unique.
	Name pulumi.StringPtrInput
	// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
	Payload pulumi.StringPtrInput
	// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
	PayloadContentEncoding pulumi.StringPtrInput
	// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
	PayloadContentType pulumi.StringPtrInput
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a secret. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 secret.
	Region pulumi.StringPtrInput
	// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
	SecretType pulumi.StringPtrInput
}

func (SecretV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*secretV1Args)(nil)).Elem()
}

type SecretV1Input interface {
	pulumi.Input

	ToSecretV1Output() SecretV1Output
	ToSecretV1OutputWithContext(ctx context.Context) SecretV1Output
}

func (*SecretV1) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretV1)(nil)).Elem()
}

func (i *SecretV1) ToSecretV1Output() SecretV1Output {
	return i.ToSecretV1OutputWithContext(context.Background())
}

func (i *SecretV1) ToSecretV1OutputWithContext(ctx context.Context) SecretV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV1Output)
}

// SecretV1ArrayInput is an input type that accepts SecretV1Array and SecretV1ArrayOutput values.
// You can construct a concrete instance of `SecretV1ArrayInput` via:
//
//	SecretV1Array{ SecretV1Args{...} }
type SecretV1ArrayInput interface {
	pulumi.Input

	ToSecretV1ArrayOutput() SecretV1ArrayOutput
	ToSecretV1ArrayOutputWithContext(context.Context) SecretV1ArrayOutput
}

type SecretV1Array []SecretV1Input

func (SecretV1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretV1)(nil)).Elem()
}

func (i SecretV1Array) ToSecretV1ArrayOutput() SecretV1ArrayOutput {
	return i.ToSecretV1ArrayOutputWithContext(context.Background())
}

func (i SecretV1Array) ToSecretV1ArrayOutputWithContext(ctx context.Context) SecretV1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV1ArrayOutput)
}

// SecretV1MapInput is an input type that accepts SecretV1Map and SecretV1MapOutput values.
// You can construct a concrete instance of `SecretV1MapInput` via:
//
//	SecretV1Map{ "key": SecretV1Args{...} }
type SecretV1MapInput interface {
	pulumi.Input

	ToSecretV1MapOutput() SecretV1MapOutput
	ToSecretV1MapOutputWithContext(context.Context) SecretV1MapOutput
}

type SecretV1Map map[string]SecretV1Input

func (SecretV1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretV1)(nil)).Elem()
}

func (i SecretV1Map) ToSecretV1MapOutput() SecretV1MapOutput {
	return i.ToSecretV1MapOutputWithContext(context.Background())
}

func (i SecretV1Map) ToSecretV1MapOutputWithContext(ctx context.Context) SecretV1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretV1MapOutput)
}

type SecretV1Output struct{ *pulumi.OutputState }

func (SecretV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretV1)(nil)).Elem()
}

func (o SecretV1Output) ToSecretV1Output() SecretV1Output {
	return o
}

func (o SecretV1Output) ToSecretV1OutputWithContext(ctx context.Context) SecretV1Output {
	return o
}

// Allows to control an access to a secret. Currently only the
// `read` operation is supported. If not specified, the secret is accessible
// project wide.
func (o SecretV1Output) Acl() SecretV1AclOutput {
	return o.ApplyT(func(v *SecretV1) SecretV1AclOutput { return v.Acl }).(SecretV1AclOutput)
}

// Metadata provided by a user or system for informational purposes.
func (o SecretV1Output) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// The map of metadata, assigned on the secret, which has been
// explicitly and implicitly added.
func (o SecretV1Output) AllMetadata() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.MapOutput { return v.AllMetadata }).(pulumi.MapOutput)
}

// Metadata provided by a user or system for informational purposes.
func (o SecretV1Output) BitLength() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.IntOutput { return v.BitLength }).(pulumi.IntOutput)
}

// The map of the content types, assigned on the secret.
func (o SecretV1Output) ContentTypes() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.MapOutput { return v.ContentTypes }).(pulumi.MapOutput)
}

// The date the secret ACL was created.
func (o SecretV1Output) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The creator of the secret.
func (o SecretV1Output) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.CreatorId }).(pulumi.StringOutput)
}

// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
func (o SecretV1Output) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringPtrOutput { return v.Expiration }).(pulumi.StringPtrOutput)
}

// Additional Metadata for the secret.
func (o SecretV1Output) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// Metadata provided by a user or system for informational purposes.
func (o SecretV1Output) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Human-readable name for the Secret. Does not have
// to be unique.
func (o SecretV1Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
func (o SecretV1Output) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringPtrOutput { return v.Payload }).(pulumi.StringPtrOutput)
}

// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
func (o SecretV1Output) PayloadContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringPtrOutput { return v.PayloadContentEncoding }).(pulumi.StringPtrOutput)
}

// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
func (o SecretV1Output) PayloadContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringPtrOutput { return v.PayloadContentType }).(pulumi.StringPtrOutput)
}

// The region in which to obtain the V1 KeyManager client.
// A KeyManager client is needed to create a secret. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// V1 secret.
func (o SecretV1Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The secret reference / where to find the secret.
func (o SecretV1Output) SecretRef() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.SecretRef }).(pulumi.StringOutput)
}

// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
func (o SecretV1Output) SecretType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.SecretType }).(pulumi.StringOutput)
}

// The status of the secret.
func (o SecretV1Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date the secret ACL was last updated.
func (o SecretV1Output) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretV1) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type SecretV1ArrayOutput struct{ *pulumi.OutputState }

func (SecretV1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretV1)(nil)).Elem()
}

func (o SecretV1ArrayOutput) ToSecretV1ArrayOutput() SecretV1ArrayOutput {
	return o
}

func (o SecretV1ArrayOutput) ToSecretV1ArrayOutputWithContext(ctx context.Context) SecretV1ArrayOutput {
	return o
}

func (o SecretV1ArrayOutput) Index(i pulumi.IntInput) SecretV1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretV1 {
		return vs[0].([]*SecretV1)[vs[1].(int)]
	}).(SecretV1Output)
}

type SecretV1MapOutput struct{ *pulumi.OutputState }

func (SecretV1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretV1)(nil)).Elem()
}

func (o SecretV1MapOutput) ToSecretV1MapOutput() SecretV1MapOutput {
	return o
}

func (o SecretV1MapOutput) ToSecretV1MapOutputWithContext(ctx context.Context) SecretV1MapOutput {
	return o
}

func (o SecretV1MapOutput) MapIndex(k pulumi.StringInput) SecretV1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretV1 {
		return vs[0].(map[string]*SecretV1)[vs[1].(string)]
	}).(SecretV1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV1Input)(nil)).Elem(), &SecretV1{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV1ArrayInput)(nil)).Elem(), SecretV1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretV1MapInput)(nil)).Elem(), SecretV1Map{})
	pulumi.RegisterOutputType(SecretV1Output{})
	pulumi.RegisterOutputType(SecretV1ArrayOutput{})
	pulumi.RegisterOutputType(SecretV1MapOutput{})
}
