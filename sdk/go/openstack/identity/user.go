// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a V3 User resource within OpenStack Keystone.
//
// > **Note:** All arguments including the user password will be stored in the
// raw state as plain-text. Read more about sensitive data in
// state.
//
// > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
//
// ## Example Usage
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
//			project1, err := identity.NewProject(ctx, "project1", nil)
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewUser(ctx, "user1", &identity.UserArgs{
//				DefaultProjectId:                 project1.ID(),
//				Description:                      pulumi.String("A user"),
//				Password:                         pulumi.String("password123"),
//				IgnoreChangePasswordUponFirstUse: pulumi.Bool(true),
//				MultiFactorAuthEnabled:           pulumi.Bool(true),
//				MultiFactorAuthRules: identity.UserMultiFactorAuthRuleArray{
//					&identity.UserMultiFactorAuthRuleArgs{
//						Rules: pulumi.StringArray{
//							pulumi.String("password"),
//							pulumi.String("totp"),
//						},
//					},
//					&identity.UserMultiFactorAuthRuleArgs{
//						Rules: pulumi.StringArray{
//							pulumi.String("password"),
//						},
//					},
//				},
//				Extra: pulumi.Map{
//					"email": pulumi.Any("user_1@foobar.com"),
//				},
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
// Users can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
//
// ```
type User struct {
	pulumi.CustomResourceState

	// The default project this user belongs to.
	DefaultProjectId pulumi.StringOutput `pulumi:"defaultProjectId"`
	// A description of the user.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain this user belongs to.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Free-form key/value pairs of extra information.
	Extra pulumi.MapOutput `pulumi:"extra"`
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse pulumi.BoolPtrOutput `pulumi:"ignoreChangePasswordUponFirstUse"`
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts pulumi.BoolPtrOutput `pulumi:"ignoreLockoutFailureAttempts"`
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry pulumi.BoolPtrOutput `pulumi:"ignorePasswordExpiry"`
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled pulumi.BoolPtrOutput `pulumi:"multiFactorAuthEnabled"`
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules UserMultiFactorAuthRuleArrayOutput `pulumi:"multiFactorAuthRules"`
	// The name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for the user.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("openstack:identity/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("openstack:identity/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The default project this user belongs to.
	DefaultProjectId *string `pulumi:"defaultProjectId"`
	// A description of the user.
	Description *string `pulumi:"description"`
	// The domain this user belongs to.
	DomainId *string `pulumi:"domainId"`
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// Free-form key/value pairs of extra information.
	Extra map[string]interface{} `pulumi:"extra"`
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse *bool `pulumi:"ignoreChangePasswordUponFirstUse"`
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts *bool `pulumi:"ignoreLockoutFailureAttempts"`
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry *bool `pulumi:"ignorePasswordExpiry"`
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled *bool `pulumi:"multiFactorAuthEnabled"`
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules []UserMultiFactorAuthRule `pulumi:"multiFactorAuthRules"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password for the user.
	Password *string `pulumi:"password"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region *string `pulumi:"region"`
}

type UserState struct {
	// The default project this user belongs to.
	DefaultProjectId pulumi.StringPtrInput
	// A description of the user.
	Description pulumi.StringPtrInput
	// The domain this user belongs to.
	DomainId pulumi.StringPtrInput
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrInput
	// Free-form key/value pairs of extra information.
	Extra pulumi.MapInput
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse pulumi.BoolPtrInput
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts pulumi.BoolPtrInput
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry pulumi.BoolPtrInput
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled pulumi.BoolPtrInput
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules UserMultiFactorAuthRuleArrayInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password for the user.
	Password pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The default project this user belongs to.
	DefaultProjectId *string `pulumi:"defaultProjectId"`
	// A description of the user.
	Description *string `pulumi:"description"`
	// The domain this user belongs to.
	DomainId *string `pulumi:"domainId"`
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// Free-form key/value pairs of extra information.
	Extra map[string]interface{} `pulumi:"extra"`
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse *bool `pulumi:"ignoreChangePasswordUponFirstUse"`
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts *bool `pulumi:"ignoreLockoutFailureAttempts"`
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry *bool `pulumi:"ignorePasswordExpiry"`
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled *bool `pulumi:"multiFactorAuthEnabled"`
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules []UserMultiFactorAuthRule `pulumi:"multiFactorAuthRules"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// The password for the user.
	Password *string `pulumi:"password"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The default project this user belongs to.
	DefaultProjectId pulumi.StringPtrInput
	// A description of the user.
	Description pulumi.StringPtrInput
	// The domain this user belongs to.
	DomainId pulumi.StringPtrInput
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrInput
	// Free-form key/value pairs of extra information.
	Extra pulumi.MapInput
	// User will not have to
	// change their password upon first use. Valid values are `true` and `false`.
	IgnoreChangePasswordUponFirstUse pulumi.BoolPtrInput
	// User will not have a failure
	// lockout placed on their account. Valid values are `true` and `false`.
	IgnoreLockoutFailureAttempts pulumi.BoolPtrInput
	// User's password will not expire.
	// Valid values are `true` and `false`.
	IgnorePasswordExpiry pulumi.BoolPtrInput
	// Whether to enable multi-factor
	// authentication. Valid values are `true` and `false`.
	MultiFactorAuthEnabled pulumi.BoolPtrInput
	// A multi-factor authentication rule.
	// The structure is documented below. Please see the
	// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
	// for more information on how to use mulit-factor rules.
	MultiFactorAuthRules UserMultiFactorAuthRuleArrayInput
	// The name of the user.
	Name pulumi.StringPtrInput
	// The password for the user.
	Password pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

func (i *User) ToOutput(ctx context.Context) pulumix.Output[*User] {
	return pulumix.Output[*User]{
		OutputState: i.ToUserOutputWithContext(ctx).OutputState,
	}
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

func (i UserArray) ToOutput(ctx context.Context) pulumix.Output[[]*User] {
	return pulumix.Output[[]*User]{
		OutputState: i.ToUserArrayOutputWithContext(ctx).OutputState,
	}
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

func (i UserMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*User] {
	return pulumix.Output[map[string]*User]{
		OutputState: i.ToUserMapOutputWithContext(ctx).OutputState,
	}
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) ToOutput(ctx context.Context) pulumix.Output[*User] {
	return pulumix.Output[*User]{
		OutputState: o.OutputState,
	}
}

// The default project this user belongs to.
func (o UserOutput) DefaultProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.DefaultProjectId }).(pulumi.StringOutput)
}

// A description of the user.
func (o UserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain this user belongs to.
func (o UserOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// Whether the user is enabled or disabled. Valid
// values are `true` and `false`.
func (o UserOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Free-form key/value pairs of extra information.
func (o UserOutput) Extra() pulumi.MapOutput {
	return o.ApplyT(func(v *User) pulumi.MapOutput { return v.Extra }).(pulumi.MapOutput)
}

// User will not have to
// change their password upon first use. Valid values are `true` and `false`.
func (o UserOutput) IgnoreChangePasswordUponFirstUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IgnoreChangePasswordUponFirstUse }).(pulumi.BoolPtrOutput)
}

// User will not have a failure
// lockout placed on their account. Valid values are `true` and `false`.
func (o UserOutput) IgnoreLockoutFailureAttempts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IgnoreLockoutFailureAttempts }).(pulumi.BoolPtrOutput)
}

// User's password will not expire.
// Valid values are `true` and `false`.
func (o UserOutput) IgnorePasswordExpiry() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IgnorePasswordExpiry }).(pulumi.BoolPtrOutput)
}

// Whether to enable multi-factor
// authentication. Valid values are `true` and `false`.
func (o UserOutput) MultiFactorAuthEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.MultiFactorAuthEnabled }).(pulumi.BoolPtrOutput)
}

// A multi-factor authentication rule.
// The structure is documented below. Please see the
// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
// for more information on how to use mulit-factor rules.
func (o UserOutput) MultiFactorAuthRules() UserMultiFactorAuthRuleArrayOutput {
	return o.ApplyT(func(v *User) UserMultiFactorAuthRuleArrayOutput { return v.MultiFactorAuthRules }).(UserMultiFactorAuthRuleArrayOutput)
}

// The name of the user.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for the user.
func (o UserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new User.
func (o UserOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*User] {
	return pulumix.Output[[]*User]{
		OutputState: o.OutputState,
	}
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*User] {
	return pulumix.Output[map[string]*User]{
		OutputState: o.OutputState,
	}
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
