// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

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
	Payload pulumi.StringOutput `pulumi:"payload"`
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
