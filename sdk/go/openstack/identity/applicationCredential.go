// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Predefined secret
//
// Application credential below will have only one `swiftoperator` role.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewApplicationCredential(ctx, "swift", &identity.ApplicationCredentialArgs{
//				Description: pulumi.String("Swift technical application credential"),
//				ExpiresAt:   pulumi.String("2019-02-13T12:12:12Z"),
//				Roles: pulumi.StringArray{
//					pulumi.String("swiftoperator"),
//				},
//				Secret: pulumi.String("supersecret"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Unrestricted with autogenerated secret and unlimited TTL
//
// Application credential below will inherit all the current user's roles.
//
// !> **WARNING:** Restrictions on these Identity operations are deliberately
// imposed as a safeguard to prevent a compromised application credential from
// regenerating itself. Disabling this restriction poses an inherent added risk.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			unrestricted, err := identity.NewApplicationCredential(ctx, "unrestricted", &identity.ApplicationCredentialArgs{
//				Description:  pulumi.String("Unrestricted application credential"),
//				Unrestricted: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("applicationCredentialSecret", unrestricted.Secret)
//			return nil
//		})
//	}
//
// ```
// ### Application credential with access rules
//
// > **Note:** Application Credential access rules are supported only in Keystone
// starting from [Train](https://releases.openstack.org/train/highlights.html#keystone-identity-service) release.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewApplicationCredential(ctx, "monitoring", &identity.ApplicationCredentialArgs{
//				AccessRules: identity.ApplicationCredentialAccessRuleArray{
//					&identity.ApplicationCredentialAccessRuleArgs{
//						Method:  pulumi.String("GET"),
//						Path:    pulumi.String("/v2.0/metrics"),
//						Service: pulumi.String("monitoring"),
//					},
//					&identity.ApplicationCredentialAccessRuleArgs{
//						Method:  pulumi.String("PUT"),
//						Path:    pulumi.String("/v2.0/metrics"),
//						Service: pulumi.String("monitoring"),
//					},
//				},
//				ExpiresAt: pulumi.String("2019-02-13T12:12:12Z"),
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
// Application Credentials can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:identity/applicationCredential:ApplicationCredential application_credential_1 c17304b7-0953-4738-abb0-67005882b0a0
//
// ```
type ApplicationCredential struct {
	pulumi.CustomResourceState

	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	AccessRules ApplicationCredentialAccessRuleArrayOutput `pulumi:"accessRules"`
	// A description of the application credential.
	// Changing this creates a new application credential.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The expiration time of the application credential
	// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
	// an application credential will never expire. Changing this creates a new
	// application credential.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// A name of the application credential. Changing this
	// creates a new application credential.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project the application credential was created
	// for and that authentication requests using this application credential will
	// be scoped to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new application credential.
	Region pulumi.StringOutput `pulumi:"region"`
	// A collection of one or more role names, which this
	// application credential has to be associated with its project. If omitted,
	// all the current user's roles within the scoped project will be inherited by
	// a new application credential. Changing this creates a new application
	// credential.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The secret for the application credential. If omitted,
	// it will be generated by the server. Changing this creates a new application
	// credential.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// A flag indicating whether the application
	// credential may be used for creation or destruction of other application
	// credentials or trusts. Changing this creates a new application credential.
	Unrestricted pulumi.BoolPtrOutput `pulumi:"unrestricted"`
}

// NewApplicationCredential registers a new resource with the given unique name, arguments, and options.
func NewApplicationCredential(ctx *pulumi.Context,
	name string, args *ApplicationCredentialArgs, opts ...pulumi.ResourceOption) (*ApplicationCredential, error) {
	if args == nil {
		args = &ApplicationCredentialArgs{}
	}

	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	var resource ApplicationCredential
	err := ctx.RegisterResource("openstack:identity/applicationCredential:ApplicationCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationCredential gets an existing ApplicationCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationCredentialState, opts ...pulumi.ResourceOption) (*ApplicationCredential, error) {
	var resource ApplicationCredential
	err := ctx.ReadResource("openstack:identity/applicationCredential:ApplicationCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationCredential resources.
type applicationCredentialState struct {
	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	AccessRules []ApplicationCredentialAccessRule `pulumi:"accessRules"`
	// A description of the application credential.
	// Changing this creates a new application credential.
	Description *string `pulumi:"description"`
	// The expiration time of the application credential
	// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
	// an application credential will never expire. Changing this creates a new
	// application credential.
	ExpiresAt *string `pulumi:"expiresAt"`
	// A name of the application credential. Changing this
	// creates a new application credential.
	Name *string `pulumi:"name"`
	// The ID of the project the application credential was created
	// for and that authentication requests using this application credential will
	// be scoped to.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new application credential.
	Region *string `pulumi:"region"`
	// A collection of one or more role names, which this
	// application credential has to be associated with its project. If omitted,
	// all the current user's roles within the scoped project will be inherited by
	// a new application credential. Changing this creates a new application
	// credential.
	Roles []string `pulumi:"roles"`
	// The secret for the application credential. If omitted,
	// it will be generated by the server. Changing this creates a new application
	// credential.
	Secret *string `pulumi:"secret"`
	// A flag indicating whether the application
	// credential may be used for creation or destruction of other application
	// credentials or trusts. Changing this creates a new application credential.
	Unrestricted *bool `pulumi:"unrestricted"`
}

type ApplicationCredentialState struct {
	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	AccessRules ApplicationCredentialAccessRuleArrayInput
	// A description of the application credential.
	// Changing this creates a new application credential.
	Description pulumi.StringPtrInput
	// The expiration time of the application credential
	// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
	// an application credential will never expire. Changing this creates a new
	// application credential.
	ExpiresAt pulumi.StringPtrInput
	// A name of the application credential. Changing this
	// creates a new application credential.
	Name pulumi.StringPtrInput
	// The ID of the project the application credential was created
	// for and that authentication requests using this application credential will
	// be scoped to.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new application credential.
	Region pulumi.StringPtrInput
	// A collection of one or more role names, which this
	// application credential has to be associated with its project. If omitted,
	// all the current user's roles within the scoped project will be inherited by
	// a new application credential. Changing this creates a new application
	// credential.
	Roles pulumi.StringArrayInput
	// The secret for the application credential. If omitted,
	// it will be generated by the server. Changing this creates a new application
	// credential.
	Secret pulumi.StringPtrInput
	// A flag indicating whether the application
	// credential may be used for creation or destruction of other application
	// credentials or trusts. Changing this creates a new application credential.
	Unrestricted pulumi.BoolPtrInput
}

func (ApplicationCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCredentialState)(nil)).Elem()
}

type applicationCredentialArgs struct {
	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	AccessRules []ApplicationCredentialAccessRule `pulumi:"accessRules"`
	// A description of the application credential.
	// Changing this creates a new application credential.
	Description *string `pulumi:"description"`
	// The expiration time of the application credential
	// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
	// an application credential will never expire. Changing this creates a new
	// application credential.
	ExpiresAt *string `pulumi:"expiresAt"`
	// A name of the application credential. Changing this
	// creates a new application credential.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new application credential.
	Region *string `pulumi:"region"`
	// A collection of one or more role names, which this
	// application credential has to be associated with its project. If omitted,
	// all the current user's roles within the scoped project will be inherited by
	// a new application credential. Changing this creates a new application
	// credential.
	Roles []string `pulumi:"roles"`
	// The secret for the application credential. If omitted,
	// it will be generated by the server. Changing this creates a new application
	// credential.
	Secret *string `pulumi:"secret"`
	// A flag indicating whether the application
	// credential may be used for creation or destruction of other application
	// credentials or trusts. Changing this creates a new application credential.
	Unrestricted *bool `pulumi:"unrestricted"`
}

// The set of arguments for constructing a ApplicationCredential resource.
type ApplicationCredentialArgs struct {
	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	AccessRules ApplicationCredentialAccessRuleArrayInput
	// A description of the application credential.
	// Changing this creates a new application credential.
	Description pulumi.StringPtrInput
	// The expiration time of the application credential
	// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
	// an application credential will never expire. Changing this creates a new
	// application credential.
	ExpiresAt pulumi.StringPtrInput
	// A name of the application credential. Changing this
	// creates a new application credential.
	Name pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new application credential.
	Region pulumi.StringPtrInput
	// A collection of one or more role names, which this
	// application credential has to be associated with its project. If omitted,
	// all the current user's roles within the scoped project will be inherited by
	// a new application credential. Changing this creates a new application
	// credential.
	Roles pulumi.StringArrayInput
	// The secret for the application credential. If omitted,
	// it will be generated by the server. Changing this creates a new application
	// credential.
	Secret pulumi.StringPtrInput
	// A flag indicating whether the application
	// credential may be used for creation or destruction of other application
	// credentials or trusts. Changing this creates a new application credential.
	Unrestricted pulumi.BoolPtrInput
}

func (ApplicationCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationCredentialArgs)(nil)).Elem()
}

type ApplicationCredentialInput interface {
	pulumi.Input

	ToApplicationCredentialOutput() ApplicationCredentialOutput
	ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput
}

func (*ApplicationCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCredential)(nil)).Elem()
}

func (i *ApplicationCredential) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return i.ToApplicationCredentialOutputWithContext(context.Background())
}

func (i *ApplicationCredential) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialOutput)
}

// ApplicationCredentialArrayInput is an input type that accepts ApplicationCredentialArray and ApplicationCredentialArrayOutput values.
// You can construct a concrete instance of `ApplicationCredentialArrayInput` via:
//
//	ApplicationCredentialArray{ ApplicationCredentialArgs{...} }
type ApplicationCredentialArrayInput interface {
	pulumi.Input

	ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput
	ToApplicationCredentialArrayOutputWithContext(context.Context) ApplicationCredentialArrayOutput
}

type ApplicationCredentialArray []ApplicationCredentialInput

func (ApplicationCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationCredential)(nil)).Elem()
}

func (i ApplicationCredentialArray) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return i.ToApplicationCredentialArrayOutputWithContext(context.Background())
}

func (i ApplicationCredentialArray) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialArrayOutput)
}

// ApplicationCredentialMapInput is an input type that accepts ApplicationCredentialMap and ApplicationCredentialMapOutput values.
// You can construct a concrete instance of `ApplicationCredentialMapInput` via:
//
//	ApplicationCredentialMap{ "key": ApplicationCredentialArgs{...} }
type ApplicationCredentialMapInput interface {
	pulumi.Input

	ToApplicationCredentialMapOutput() ApplicationCredentialMapOutput
	ToApplicationCredentialMapOutputWithContext(context.Context) ApplicationCredentialMapOutput
}

type ApplicationCredentialMap map[string]ApplicationCredentialInput

func (ApplicationCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationCredential)(nil)).Elem()
}

func (i ApplicationCredentialMap) ToApplicationCredentialMapOutput() ApplicationCredentialMapOutput {
	return i.ToApplicationCredentialMapOutputWithContext(context.Background())
}

func (i ApplicationCredentialMap) ToApplicationCredentialMapOutputWithContext(ctx context.Context) ApplicationCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialMapOutput)
}

type ApplicationCredentialOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return o
}

// A collection of one or more access rules, which
// this application credential allows to follow. The structure is described
// below. Changing this creates a new application credential.
func (o ApplicationCredentialOutput) AccessRules() ApplicationCredentialAccessRuleArrayOutput {
	return o.ApplyT(func(v *ApplicationCredential) ApplicationCredentialAccessRuleArrayOutput { return v.AccessRules }).(ApplicationCredentialAccessRuleArrayOutput)
}

// A description of the application credential.
// Changing this creates a new application credential.
func (o ApplicationCredentialOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The expiration time of the application credential
// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
// an application credential will never expire. Changing this creates a new
// application credential.
func (o ApplicationCredentialOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// A name of the application credential. Changing this
// creates a new application credential.
func (o ApplicationCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project the application credential was created
// for and that authentication requests using this application credential will
// be scoped to.
func (o ApplicationCredentialOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new application credential.
func (o ApplicationCredentialOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A collection of one or more role names, which this
// application credential has to be associated with its project. If omitted,
// all the current user's roles within the scoped project will be inherited by
// a new application credential. Changing this creates a new application
// credential.
func (o ApplicationCredentialOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The secret for the application credential. If omitted,
// it will be generated by the server. Changing this creates a new application
// credential.
func (o ApplicationCredentialOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// A flag indicating whether the application
// credential may be used for creation or destruction of other application
// credentials or trusts. Changing this creates a new application credential.
func (o ApplicationCredentialOutput) Unrestricted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationCredential) pulumi.BoolPtrOutput { return v.Unrestricted }).(pulumi.BoolPtrOutput)
}

type ApplicationCredentialArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) Index(i pulumi.IntInput) ApplicationCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationCredential {
		return vs[0].([]*ApplicationCredential)[vs[1].(int)]
	}).(ApplicationCredentialOutput)
}

type ApplicationCredentialMapOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationCredential)(nil)).Elem()
}

func (o ApplicationCredentialMapOutput) ToApplicationCredentialMapOutput() ApplicationCredentialMapOutput {
	return o
}

func (o ApplicationCredentialMapOutput) ToApplicationCredentialMapOutputWithContext(ctx context.Context) ApplicationCredentialMapOutput {
	return o
}

func (o ApplicationCredentialMapOutput) MapIndex(k pulumi.StringInput) ApplicationCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationCredential {
		return vs[0].(map[string]*ApplicationCredential)[vs[1].(string)]
	}).(ApplicationCredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialInput)(nil)).Elem(), &ApplicationCredential{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialArrayInput)(nil)).Elem(), ApplicationCredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialMapInput)(nil)).Elem(), ApplicationCredentialMap{})
	pulumi.RegisterOutputType(ApplicationCredentialOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialArrayOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialMapOutput{})
}
