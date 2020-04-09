// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vpnaas

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 Neutron IKE policy resource within OpenStack.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/vpnaas_ike_policy_v2.html.markdown.
type IkePolicy struct {
	pulumi.CustomResourceState

	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm pulumi.StringPtrOutput `pulumi:"authAlgorithm"`
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm pulumi.StringPtrOutput `pulumi:"encryptionAlgorithm"`
	// The IKE mode. A valid value is v1 or v2. Default is v1.
	// Changing this updates the existing policy.
	IkeVersion pulumi.StringPtrOutput `pulumi:"ikeVersion"`
	// The lifetime of the security association. Consists of Unit and Value.
	// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
	// Default is seconds.
	// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Lifetimes IkePolicyLifetimeArrayOutput `pulumi:"lifetimes"`
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
	// Changing this updates the existing policy.
	Pfs pulumi.StringPtrOutput `pulumi:"pfs"`
	// The IKE mode. A valid value is main, which is the default.
	// Changing this updates the existing policy.
	Phase1NegotiationMode pulumi.StringPtrOutput `pulumi:"phase1NegotiationMode"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region pulumi.StringOutput `pulumi:"region"`
	// The owner of the policy. Required if admin wants to
	// create a service for another policy. Changing this creates a new policy.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs pulumi.MapOutput `pulumi:"valueSpecs"`
}

// NewIkePolicy registers a new resource with the given unique name, arguments, and options.
func NewIkePolicy(ctx *pulumi.Context,
	name string, args *IkePolicyArgs, opts ...pulumi.ResourceOption) (*IkePolicy, error) {
	if args == nil {
		args = &IkePolicyArgs{}
	}
	var resource IkePolicy
	err := ctx.RegisterResource("openstack:vpnaas/ikePolicy:IkePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIkePolicy gets an existing IkePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIkePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IkePolicyState, opts ...pulumi.ResourceOption) (*IkePolicy, error) {
	var resource IkePolicy
	err := ctx.ReadResource("openstack:vpnaas/ikePolicy:IkePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IkePolicy resources.
type ikePolicyState struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm *string `pulumi:"authAlgorithm"`
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description *string `pulumi:"description"`
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// The IKE mode. A valid value is v1 or v2. Default is v1.
	// Changing this updates the existing policy.
	IkeVersion *string `pulumi:"ikeVersion"`
	// The lifetime of the security association. Consists of Unit and Value.
	// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
	// Default is seconds.
	// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Lifetimes []IkePolicyLifetime `pulumi:"lifetimes"`
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name *string `pulumi:"name"`
	// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
	// Changing this updates the existing policy.
	Pfs *string `pulumi:"pfs"`
	// The IKE mode. A valid value is main, which is the default.
	// Changing this updates the existing policy.
	Phase1NegotiationMode *string `pulumi:"phase1NegotiationMode"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region *string `pulumi:"region"`
	// The owner of the policy. Required if admin wants to
	// create a service for another policy. Changing this creates a new policy.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

type IkePolicyState struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm pulumi.StringPtrInput
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description pulumi.StringPtrInput
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm pulumi.StringPtrInput
	// The IKE mode. A valid value is v1 or v2. Default is v1.
	// Changing this updates the existing policy.
	IkeVersion pulumi.StringPtrInput
	// The lifetime of the security association. Consists of Unit and Value.
	// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
	// Default is seconds.
	// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Lifetimes IkePolicyLifetimeArrayInput
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name pulumi.StringPtrInput
	// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
	// Changing this updates the existing policy.
	Pfs pulumi.StringPtrInput
	// The IKE mode. A valid value is main, which is the default.
	// Changing this updates the existing policy.
	Phase1NegotiationMode pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region pulumi.StringPtrInput
	// The owner of the policy. Required if admin wants to
	// create a service for another policy. Changing this creates a new policy.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (IkePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ikePolicyState)(nil)).Elem()
}

type ikePolicyArgs struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm *string `pulumi:"authAlgorithm"`
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description *string `pulumi:"description"`
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// The IKE mode. A valid value is v1 or v2. Default is v1.
	// Changing this updates the existing policy.
	IkeVersion *string `pulumi:"ikeVersion"`
	// The lifetime of the security association. Consists of Unit and Value.
	// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
	// Default is seconds.
	// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Lifetimes []IkePolicyLifetime `pulumi:"lifetimes"`
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name *string `pulumi:"name"`
	// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
	// Changing this updates the existing policy.
	Pfs *string `pulumi:"pfs"`
	// The IKE mode. A valid value is main, which is the default.
	// Changing this updates the existing policy.
	Phase1NegotiationMode *string `pulumi:"phase1NegotiationMode"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region *string `pulumi:"region"`
	// The owner of the policy. Required if admin wants to
	// create a service for another policy. Changing this creates a new policy.
	TenantId *string `pulumi:"tenantId"`
	// Map of additional options.
	ValueSpecs map[string]interface{} `pulumi:"valueSpecs"`
}

// The set of arguments for constructing a IkePolicy resource.
type IkePolicyArgs struct {
	// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
	// Default is sha1. Changing this updates the algorithm of the existing policy.
	AuthAlgorithm pulumi.StringPtrInput
	// The human-readable description for the policy.
	// Changing this updates the description of the existing policy.
	Description pulumi.StringPtrInput
	// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
	// The default value is aes-128. Changing this updates the existing policy.
	EncryptionAlgorithm pulumi.StringPtrInput
	// The IKE mode. A valid value is v1 or v2. Default is v1.
	// Changing this updates the existing policy.
	IkeVersion pulumi.StringPtrInput
	// The lifetime of the security association. Consists of Unit and Value.
	// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
	// Default is seconds.
	// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
	// Default is 3600.
	Lifetimes IkePolicyLifetimeArrayInput
	// The name of the policy. Changing this updates the name of
	// the existing policy.
	Name pulumi.StringPtrInput
	// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
	// Changing this updates the existing policy.
	Pfs pulumi.StringPtrInput
	// The IKE mode. A valid value is main, which is the default.
	// Changing this updates the existing policy.
	Phase1NegotiationMode pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a VPN service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// service.
	Region pulumi.StringPtrInput
	// The owner of the policy. Required if admin wants to
	// create a service for another policy. Changing this creates a new policy.
	TenantId pulumi.StringPtrInput
	// Map of additional options.
	ValueSpecs pulumi.MapInput
}

func (IkePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ikePolicyArgs)(nil)).Elem()
}
