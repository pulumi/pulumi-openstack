// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Application Credentials can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:identity/applicationCredential:ApplicationCredential application_credential_1 c17304b7-0953-4738-abb0-67005882b0a0
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
	return reflect.TypeOf((*ApplicationCredential)(nil))
}

func (i *ApplicationCredential) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return i.ToApplicationCredentialOutputWithContext(context.Background())
}

func (i *ApplicationCredential) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialOutput)
}

func (i *ApplicationCredential) ToApplicationCredentialPtrOutput() ApplicationCredentialPtrOutput {
	return i.ToApplicationCredentialPtrOutputWithContext(context.Background())
}

func (i *ApplicationCredential) ToApplicationCredentialPtrOutputWithContext(ctx context.Context) ApplicationCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialPtrOutput)
}

type ApplicationCredentialPtrInput interface {
	pulumi.Input

	ToApplicationCredentialPtrOutput() ApplicationCredentialPtrOutput
	ToApplicationCredentialPtrOutputWithContext(ctx context.Context) ApplicationCredentialPtrOutput
}

type applicationCredentialPtrType ApplicationCredentialArgs

func (*applicationCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCredential)(nil))
}

func (i *applicationCredentialPtrType) ToApplicationCredentialPtrOutput() ApplicationCredentialPtrOutput {
	return i.ToApplicationCredentialPtrOutputWithContext(context.Background())
}

func (i *applicationCredentialPtrType) ToApplicationCredentialPtrOutputWithContext(ctx context.Context) ApplicationCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialPtrOutput)
}

// ApplicationCredentialArrayInput is an input type that accepts ApplicationCredentialArray and ApplicationCredentialArrayOutput values.
// You can construct a concrete instance of `ApplicationCredentialArrayInput` via:
//
//          ApplicationCredentialArray{ ApplicationCredentialArgs{...} }
type ApplicationCredentialArrayInput interface {
	pulumi.Input

	ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput
	ToApplicationCredentialArrayOutputWithContext(context.Context) ApplicationCredentialArrayOutput
}

type ApplicationCredentialArray []ApplicationCredentialInput

func (ApplicationCredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApplicationCredential)(nil))
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
//          ApplicationCredentialMap{ "key": ApplicationCredentialArgs{...} }
type ApplicationCredentialMapInput interface {
	pulumi.Input

	ToApplicationCredentialMapOutput() ApplicationCredentialMapOutput
	ToApplicationCredentialMapOutputWithContext(context.Context) ApplicationCredentialMapOutput
}

type ApplicationCredentialMap map[string]ApplicationCredentialInput

func (ApplicationCredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApplicationCredential)(nil))
}

func (i ApplicationCredentialMap) ToApplicationCredentialMapOutput() ApplicationCredentialMapOutput {
	return i.ToApplicationCredentialMapOutputWithContext(context.Background())
}

func (i ApplicationCredentialMap) ToApplicationCredentialMapOutputWithContext(ctx context.Context) ApplicationCredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialMapOutput)
}

type ApplicationCredentialOutput struct {
	*pulumi.OutputState
}

func (ApplicationCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredential)(nil))
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutput() ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) ToApplicationCredentialOutputWithContext(ctx context.Context) ApplicationCredentialOutput {
	return o
}

func (o ApplicationCredentialOutput) ToApplicationCredentialPtrOutput() ApplicationCredentialPtrOutput {
	return o.ToApplicationCredentialPtrOutputWithContext(context.Background())
}

func (o ApplicationCredentialOutput) ToApplicationCredentialPtrOutputWithContext(ctx context.Context) ApplicationCredentialPtrOutput {
	return o.ApplyT(func(v ApplicationCredential) *ApplicationCredential {
		return &v
	}).(ApplicationCredentialPtrOutput)
}

type ApplicationCredentialPtrOutput struct {
	*pulumi.OutputState
}

func (ApplicationCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationCredential)(nil))
}

func (o ApplicationCredentialPtrOutput) ToApplicationCredentialPtrOutput() ApplicationCredentialPtrOutput {
	return o
}

func (o ApplicationCredentialPtrOutput) ToApplicationCredentialPtrOutputWithContext(ctx context.Context) ApplicationCredentialPtrOutput {
	return o
}

type ApplicationCredentialArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredential)(nil))
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutput() ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) ToApplicationCredentialArrayOutputWithContext(ctx context.Context) ApplicationCredentialArrayOutput {
	return o
}

func (o ApplicationCredentialArrayOutput) Index(i pulumi.IntInput) ApplicationCredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationCredential {
		return vs[0].([]ApplicationCredential)[vs[1].(int)]
	}).(ApplicationCredentialOutput)
}

type ApplicationCredentialMapOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApplicationCredential)(nil))
}

func (o ApplicationCredentialMapOutput) ToApplicationCredentialMapOutput() ApplicationCredentialMapOutput {
	return o
}

func (o ApplicationCredentialMapOutput) ToApplicationCredentialMapOutputWithContext(ctx context.Context) ApplicationCredentialMapOutput {
	return o
}

func (o ApplicationCredentialMapOutput) MapIndex(k pulumi.StringInput) ApplicationCredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApplicationCredential {
		return vs[0].(map[string]ApplicationCredential)[vs[1].(string)]
	}).(ApplicationCredentialOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationCredentialOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialPtrOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialArrayOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialMapOutput{})
}
