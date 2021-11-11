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
// Users can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
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
	return reflect.TypeOf((*User)(nil))
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

func (i *User) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *User) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

type UserPtrInput interface {
	pulumi.Input

	ToUserPtrOutput() UserPtrOutput
	ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput
}

type userPtrType UserArgs

func (*userPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (i *userPtrType) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *userPtrType) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
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

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
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

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) ToUserPtrOutput() UserPtrOutput {
	return o.ToUserPtrOutputWithContext(context.Background())
}

func (o UserOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v User) *User {
		return &v
	}).(UserPtrOutput)
}

type UserPtrOutput struct{ *pulumi.OutputState }

func (UserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (o UserPtrOutput) ToUserPtrOutput() UserPtrOutput {
	return o
}

func (o UserPtrOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o
}

func (o UserPtrOutput) Elem() UserOutput {
	return o.ApplyT(func(v *User) User {
		if v != nil {
			return *v
		}
		var ret User
		return ret
	}).(UserOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]User)(nil))
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) User {
		return vs[0].([]User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]User)(nil))
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) User {
		return vs[0].(map[string]User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPtrInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserPtrOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
