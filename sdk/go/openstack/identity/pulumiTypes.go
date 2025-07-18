// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ApplicationCredentialAccessRule struct {
	// The ID of the existing access rule. The access rule ID of
	// another application credential can be provided.
	Id *string `pulumi:"id"`
	// The request method that the application credential is
	// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
	// `HEAD`, `PATCH`, `PUT` and `DELETE`.
	Method string `pulumi:"method"`
	// The API path that the application credential is permitted
	// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
	// **\*** to match against any string in the path up to a **/**, or the recursive
	// wildcard **\*\*** to include **/** in the matched path.
	Path string `pulumi:"path"`
	// The service type identifier for the service that the
	// application credential is granted to access. Must be a service type that is
	// listed in the service catalog and not a code name for a service. E.g.
	// **identity**, **compute**, **volumev3**, **image**, **network**,
	// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
	Service string `pulumi:"service"`
}

// ApplicationCredentialAccessRuleInput is an input type that accepts ApplicationCredentialAccessRuleArgs and ApplicationCredentialAccessRuleOutput values.
// You can construct a concrete instance of `ApplicationCredentialAccessRuleInput` via:
//
//	ApplicationCredentialAccessRuleArgs{...}
type ApplicationCredentialAccessRuleInput interface {
	pulumi.Input

	ToApplicationCredentialAccessRuleOutput() ApplicationCredentialAccessRuleOutput
	ToApplicationCredentialAccessRuleOutputWithContext(context.Context) ApplicationCredentialAccessRuleOutput
}

type ApplicationCredentialAccessRuleArgs struct {
	// The ID of the existing access rule. The access rule ID of
	// another application credential can be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The request method that the application credential is
	// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
	// `HEAD`, `PATCH`, `PUT` and `DELETE`.
	Method pulumi.StringInput `pulumi:"method"`
	// The API path that the application credential is permitted
	// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
	// **\*** to match against any string in the path up to a **/**, or the recursive
	// wildcard **\*\*** to include **/** in the matched path.
	Path pulumi.StringInput `pulumi:"path"`
	// The service type identifier for the service that the
	// application credential is granted to access. Must be a service type that is
	// listed in the service catalog and not a code name for a service. E.g.
	// **identity**, **compute**, **volumev3**, **image**, **network**,
	// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
	Service pulumi.StringInput `pulumi:"service"`
}

func (ApplicationCredentialAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredentialAccessRule)(nil)).Elem()
}

func (i ApplicationCredentialAccessRuleArgs) ToApplicationCredentialAccessRuleOutput() ApplicationCredentialAccessRuleOutput {
	return i.ToApplicationCredentialAccessRuleOutputWithContext(context.Background())
}

func (i ApplicationCredentialAccessRuleArgs) ToApplicationCredentialAccessRuleOutputWithContext(ctx context.Context) ApplicationCredentialAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialAccessRuleOutput)
}

// ApplicationCredentialAccessRuleArrayInput is an input type that accepts ApplicationCredentialAccessRuleArray and ApplicationCredentialAccessRuleArrayOutput values.
// You can construct a concrete instance of `ApplicationCredentialAccessRuleArrayInput` via:
//
//	ApplicationCredentialAccessRuleArray{ ApplicationCredentialAccessRuleArgs{...} }
type ApplicationCredentialAccessRuleArrayInput interface {
	pulumi.Input

	ToApplicationCredentialAccessRuleArrayOutput() ApplicationCredentialAccessRuleArrayOutput
	ToApplicationCredentialAccessRuleArrayOutputWithContext(context.Context) ApplicationCredentialAccessRuleArrayOutput
}

type ApplicationCredentialAccessRuleArray []ApplicationCredentialAccessRuleInput

func (ApplicationCredentialAccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredentialAccessRule)(nil)).Elem()
}

func (i ApplicationCredentialAccessRuleArray) ToApplicationCredentialAccessRuleArrayOutput() ApplicationCredentialAccessRuleArrayOutput {
	return i.ToApplicationCredentialAccessRuleArrayOutputWithContext(context.Background())
}

func (i ApplicationCredentialAccessRuleArray) ToApplicationCredentialAccessRuleArrayOutputWithContext(ctx context.Context) ApplicationCredentialAccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationCredentialAccessRuleArrayOutput)
}

type ApplicationCredentialAccessRuleOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationCredentialAccessRule)(nil)).Elem()
}

func (o ApplicationCredentialAccessRuleOutput) ToApplicationCredentialAccessRuleOutput() ApplicationCredentialAccessRuleOutput {
	return o
}

func (o ApplicationCredentialAccessRuleOutput) ToApplicationCredentialAccessRuleOutputWithContext(ctx context.Context) ApplicationCredentialAccessRuleOutput {
	return o
}

// The ID of the existing access rule. The access rule ID of
// another application credential can be provided.
func (o ApplicationCredentialAccessRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationCredentialAccessRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The request method that the application credential is
// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
// `HEAD`, `PATCH`, `PUT` and `DELETE`.
func (o ApplicationCredentialAccessRuleOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationCredentialAccessRule) string { return v.Method }).(pulumi.StringOutput)
}

// The API path that the application credential is permitted
// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
// **\*** to match against any string in the path up to a **/**, or the recursive
// wildcard **\*\*** to include **/** in the matched path.
func (o ApplicationCredentialAccessRuleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationCredentialAccessRule) string { return v.Path }).(pulumi.StringOutput)
}

// The service type identifier for the service that the
// application credential is granted to access. Must be a service type that is
// listed in the service catalog and not a code name for a service. E.g.
// **identity**, **compute**, **volumev3**, **image**, **network**,
// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
func (o ApplicationCredentialAccessRuleOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationCredentialAccessRule) string { return v.Service }).(pulumi.StringOutput)
}

type ApplicationCredentialAccessRuleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationCredentialAccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationCredentialAccessRule)(nil)).Elem()
}

func (o ApplicationCredentialAccessRuleArrayOutput) ToApplicationCredentialAccessRuleArrayOutput() ApplicationCredentialAccessRuleArrayOutput {
	return o
}

func (o ApplicationCredentialAccessRuleArrayOutput) ToApplicationCredentialAccessRuleArrayOutputWithContext(ctx context.Context) ApplicationCredentialAccessRuleArrayOutput {
	return o
}

func (o ApplicationCredentialAccessRuleArrayOutput) Index(i pulumi.IntInput) ApplicationCredentialAccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationCredentialAccessRule {
		return vs[0].([]ApplicationCredentialAccessRule)[vs[1].(int)]
	}).(ApplicationCredentialAccessRuleOutput)
}

type UserMultiFactorAuthRule struct {
	// A list of authentication plugins that the user must
	// authenticate with.
	Rules []string `pulumi:"rules"`
}

// UserMultiFactorAuthRuleInput is an input type that accepts UserMultiFactorAuthRuleArgs and UserMultiFactorAuthRuleOutput values.
// You can construct a concrete instance of `UserMultiFactorAuthRuleInput` via:
//
//	UserMultiFactorAuthRuleArgs{...}
type UserMultiFactorAuthRuleInput interface {
	pulumi.Input

	ToUserMultiFactorAuthRuleOutput() UserMultiFactorAuthRuleOutput
	ToUserMultiFactorAuthRuleOutputWithContext(context.Context) UserMultiFactorAuthRuleOutput
}

type UserMultiFactorAuthRuleArgs struct {
	// A list of authentication plugins that the user must
	// authenticate with.
	Rules pulumi.StringArrayInput `pulumi:"rules"`
}

func (UserMultiFactorAuthRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserMultiFactorAuthRule)(nil)).Elem()
}

func (i UserMultiFactorAuthRuleArgs) ToUserMultiFactorAuthRuleOutput() UserMultiFactorAuthRuleOutput {
	return i.ToUserMultiFactorAuthRuleOutputWithContext(context.Background())
}

func (i UserMultiFactorAuthRuleArgs) ToUserMultiFactorAuthRuleOutputWithContext(ctx context.Context) UserMultiFactorAuthRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMultiFactorAuthRuleOutput)
}

// UserMultiFactorAuthRuleArrayInput is an input type that accepts UserMultiFactorAuthRuleArray and UserMultiFactorAuthRuleArrayOutput values.
// You can construct a concrete instance of `UserMultiFactorAuthRuleArrayInput` via:
//
//	UserMultiFactorAuthRuleArray{ UserMultiFactorAuthRuleArgs{...} }
type UserMultiFactorAuthRuleArrayInput interface {
	pulumi.Input

	ToUserMultiFactorAuthRuleArrayOutput() UserMultiFactorAuthRuleArrayOutput
	ToUserMultiFactorAuthRuleArrayOutputWithContext(context.Context) UserMultiFactorAuthRuleArrayOutput
}

type UserMultiFactorAuthRuleArray []UserMultiFactorAuthRuleInput

func (UserMultiFactorAuthRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserMultiFactorAuthRule)(nil)).Elem()
}

func (i UserMultiFactorAuthRuleArray) ToUserMultiFactorAuthRuleArrayOutput() UserMultiFactorAuthRuleArrayOutput {
	return i.ToUserMultiFactorAuthRuleArrayOutputWithContext(context.Background())
}

func (i UserMultiFactorAuthRuleArray) ToUserMultiFactorAuthRuleArrayOutputWithContext(ctx context.Context) UserMultiFactorAuthRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMultiFactorAuthRuleArrayOutput)
}

type UserMultiFactorAuthRuleOutput struct{ *pulumi.OutputState }

func (UserMultiFactorAuthRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserMultiFactorAuthRule)(nil)).Elem()
}

func (o UserMultiFactorAuthRuleOutput) ToUserMultiFactorAuthRuleOutput() UserMultiFactorAuthRuleOutput {
	return o
}

func (o UserMultiFactorAuthRuleOutput) ToUserMultiFactorAuthRuleOutputWithContext(ctx context.Context) UserMultiFactorAuthRuleOutput {
	return o
}

// A list of authentication plugins that the user must
// authenticate with.
func (o UserMultiFactorAuthRuleOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UserMultiFactorAuthRule) []string { return v.Rules }).(pulumi.StringArrayOutput)
}

type UserMultiFactorAuthRuleArrayOutput struct{ *pulumi.OutputState }

func (UserMultiFactorAuthRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserMultiFactorAuthRule)(nil)).Elem()
}

func (o UserMultiFactorAuthRuleArrayOutput) ToUserMultiFactorAuthRuleArrayOutput() UserMultiFactorAuthRuleArrayOutput {
	return o
}

func (o UserMultiFactorAuthRuleArrayOutput) ToUserMultiFactorAuthRuleArrayOutputWithContext(ctx context.Context) UserMultiFactorAuthRuleArrayOutput {
	return o
}

func (o UserMultiFactorAuthRuleArrayOutput) Index(i pulumi.IntInput) UserMultiFactorAuthRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserMultiFactorAuthRule {
		return vs[0].([]UserMultiFactorAuthRule)[vs[1].(int)]
	}).(UserMultiFactorAuthRuleOutput)
}

type GetAuthScopeRole struct {
	// The ID of the role.
	RoleId string `pulumi:"roleId"`
	// The name of the role.
	RoleName string `pulumi:"roleName"`
}

// GetAuthScopeRoleInput is an input type that accepts GetAuthScopeRoleArgs and GetAuthScopeRoleOutput values.
// You can construct a concrete instance of `GetAuthScopeRoleInput` via:
//
//	GetAuthScopeRoleArgs{...}
type GetAuthScopeRoleInput interface {
	pulumi.Input

	ToGetAuthScopeRoleOutput() GetAuthScopeRoleOutput
	ToGetAuthScopeRoleOutputWithContext(context.Context) GetAuthScopeRoleOutput
}

type GetAuthScopeRoleArgs struct {
	// The ID of the role.
	RoleId pulumi.StringInput `pulumi:"roleId"`
	// The name of the role.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (GetAuthScopeRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeRole)(nil)).Elem()
}

func (i GetAuthScopeRoleArgs) ToGetAuthScopeRoleOutput() GetAuthScopeRoleOutput {
	return i.ToGetAuthScopeRoleOutputWithContext(context.Background())
}

func (i GetAuthScopeRoleArgs) ToGetAuthScopeRoleOutputWithContext(ctx context.Context) GetAuthScopeRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAuthScopeRoleOutput)
}

// GetAuthScopeRoleArrayInput is an input type that accepts GetAuthScopeRoleArray and GetAuthScopeRoleArrayOutput values.
// You can construct a concrete instance of `GetAuthScopeRoleArrayInput` via:
//
//	GetAuthScopeRoleArray{ GetAuthScopeRoleArgs{...} }
type GetAuthScopeRoleArrayInput interface {
	pulumi.Input

	ToGetAuthScopeRoleArrayOutput() GetAuthScopeRoleArrayOutput
	ToGetAuthScopeRoleArrayOutputWithContext(context.Context) GetAuthScopeRoleArrayOutput
}

type GetAuthScopeRoleArray []GetAuthScopeRoleInput

func (GetAuthScopeRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAuthScopeRole)(nil)).Elem()
}

func (i GetAuthScopeRoleArray) ToGetAuthScopeRoleArrayOutput() GetAuthScopeRoleArrayOutput {
	return i.ToGetAuthScopeRoleArrayOutputWithContext(context.Background())
}

func (i GetAuthScopeRoleArray) ToGetAuthScopeRoleArrayOutputWithContext(ctx context.Context) GetAuthScopeRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAuthScopeRoleArrayOutput)
}

type GetAuthScopeRoleOutput struct{ *pulumi.OutputState }

func (GetAuthScopeRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeRole)(nil)).Elem()
}

func (o GetAuthScopeRoleOutput) ToGetAuthScopeRoleOutput() GetAuthScopeRoleOutput {
	return o
}

func (o GetAuthScopeRoleOutput) ToGetAuthScopeRoleOutputWithContext(ctx context.Context) GetAuthScopeRoleOutput {
	return o
}

// The ID of the role.
func (o GetAuthScopeRoleOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeRole) string { return v.RoleId }).(pulumi.StringOutput)
}

// The name of the role.
func (o GetAuthScopeRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeRole) string { return v.RoleName }).(pulumi.StringOutput)
}

type GetAuthScopeRoleArrayOutput struct{ *pulumi.OutputState }

func (GetAuthScopeRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAuthScopeRole)(nil)).Elem()
}

func (o GetAuthScopeRoleArrayOutput) ToGetAuthScopeRoleArrayOutput() GetAuthScopeRoleArrayOutput {
	return o
}

func (o GetAuthScopeRoleArrayOutput) ToGetAuthScopeRoleArrayOutputWithContext(ctx context.Context) GetAuthScopeRoleArrayOutput {
	return o
}

func (o GetAuthScopeRoleArrayOutput) Index(i pulumi.IntInput) GetAuthScopeRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAuthScopeRole {
		return vs[0].([]GetAuthScopeRole)[vs[1].(int)]
	}).(GetAuthScopeRoleOutput)
}

type GetAuthScopeServiceCatalog struct {
	// A list of endpoints for the service.
	Endpoints []GetAuthScopeServiceCatalogEndpoint `pulumi:"endpoints"`
	// The ID of the endpoint.
	Id string `pulumi:"id"`
	// The name of the scope. This is an arbitrary name which is
	// only used as a unique identifier so an actual token isn't used as the ID.
	Name string `pulumi:"name"`
	// The type of the service.
	Type string `pulumi:"type"`
}

// GetAuthScopeServiceCatalogInput is an input type that accepts GetAuthScopeServiceCatalogArgs and GetAuthScopeServiceCatalogOutput values.
// You can construct a concrete instance of `GetAuthScopeServiceCatalogInput` via:
//
//	GetAuthScopeServiceCatalogArgs{...}
type GetAuthScopeServiceCatalogInput interface {
	pulumi.Input

	ToGetAuthScopeServiceCatalogOutput() GetAuthScopeServiceCatalogOutput
	ToGetAuthScopeServiceCatalogOutputWithContext(context.Context) GetAuthScopeServiceCatalogOutput
}

type GetAuthScopeServiceCatalogArgs struct {
	// A list of endpoints for the service.
	Endpoints GetAuthScopeServiceCatalogEndpointArrayInput `pulumi:"endpoints"`
	// The ID of the endpoint.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the scope. This is an arbitrary name which is
	// only used as a unique identifier so an actual token isn't used as the ID.
	Name pulumi.StringInput `pulumi:"name"`
	// The type of the service.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetAuthScopeServiceCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeServiceCatalog)(nil)).Elem()
}

func (i GetAuthScopeServiceCatalogArgs) ToGetAuthScopeServiceCatalogOutput() GetAuthScopeServiceCatalogOutput {
	return i.ToGetAuthScopeServiceCatalogOutputWithContext(context.Background())
}

func (i GetAuthScopeServiceCatalogArgs) ToGetAuthScopeServiceCatalogOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAuthScopeServiceCatalogOutput)
}

// GetAuthScopeServiceCatalogArrayInput is an input type that accepts GetAuthScopeServiceCatalogArray and GetAuthScopeServiceCatalogArrayOutput values.
// You can construct a concrete instance of `GetAuthScopeServiceCatalogArrayInput` via:
//
//	GetAuthScopeServiceCatalogArray{ GetAuthScopeServiceCatalogArgs{...} }
type GetAuthScopeServiceCatalogArrayInput interface {
	pulumi.Input

	ToGetAuthScopeServiceCatalogArrayOutput() GetAuthScopeServiceCatalogArrayOutput
	ToGetAuthScopeServiceCatalogArrayOutputWithContext(context.Context) GetAuthScopeServiceCatalogArrayOutput
}

type GetAuthScopeServiceCatalogArray []GetAuthScopeServiceCatalogInput

func (GetAuthScopeServiceCatalogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAuthScopeServiceCatalog)(nil)).Elem()
}

func (i GetAuthScopeServiceCatalogArray) ToGetAuthScopeServiceCatalogArrayOutput() GetAuthScopeServiceCatalogArrayOutput {
	return i.ToGetAuthScopeServiceCatalogArrayOutputWithContext(context.Background())
}

func (i GetAuthScopeServiceCatalogArray) ToGetAuthScopeServiceCatalogArrayOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAuthScopeServiceCatalogArrayOutput)
}

type GetAuthScopeServiceCatalogOutput struct{ *pulumi.OutputState }

func (GetAuthScopeServiceCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeServiceCatalog)(nil)).Elem()
}

func (o GetAuthScopeServiceCatalogOutput) ToGetAuthScopeServiceCatalogOutput() GetAuthScopeServiceCatalogOutput {
	return o
}

func (o GetAuthScopeServiceCatalogOutput) ToGetAuthScopeServiceCatalogOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogOutput {
	return o
}

// A list of endpoints for the service.
func (o GetAuthScopeServiceCatalogOutput) Endpoints() GetAuthScopeServiceCatalogEndpointArrayOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalog) []GetAuthScopeServiceCatalogEndpoint { return v.Endpoints }).(GetAuthScopeServiceCatalogEndpointArrayOutput)
}

// The ID of the endpoint.
func (o GetAuthScopeServiceCatalogOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalog) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the scope. This is an arbitrary name which is
// only used as a unique identifier so an actual token isn't used as the ID.
func (o GetAuthScopeServiceCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalog) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the service.
func (o GetAuthScopeServiceCatalogOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalog) string { return v.Type }).(pulumi.StringOutput)
}

type GetAuthScopeServiceCatalogArrayOutput struct{ *pulumi.OutputState }

func (GetAuthScopeServiceCatalogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAuthScopeServiceCatalog)(nil)).Elem()
}

func (o GetAuthScopeServiceCatalogArrayOutput) ToGetAuthScopeServiceCatalogArrayOutput() GetAuthScopeServiceCatalogArrayOutput {
	return o
}

func (o GetAuthScopeServiceCatalogArrayOutput) ToGetAuthScopeServiceCatalogArrayOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogArrayOutput {
	return o
}

func (o GetAuthScopeServiceCatalogArrayOutput) Index(i pulumi.IntInput) GetAuthScopeServiceCatalogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAuthScopeServiceCatalog {
		return vs[0].([]GetAuthScopeServiceCatalog)[vs[1].(int)]
	}).(GetAuthScopeServiceCatalogOutput)
}

type GetAuthScopeServiceCatalogEndpoint struct {
	// The ID of the endpoint.
	Id string `pulumi:"id"`
	// The interface of the endpoint.
	Interface string `pulumi:"interface"`
	// The region in which to obtain the V3 Identity client.
	// A Identity client is needed to retrieve tokens IDs. If omitted, the
	// `region` argument of the provider is used.
	Region string `pulumi:"region"`
	// The region ID of the endpoint.
	RegionId string `pulumi:"regionId"`
	// The URL of the endpoint.
	Url string `pulumi:"url"`
}

// GetAuthScopeServiceCatalogEndpointInput is an input type that accepts GetAuthScopeServiceCatalogEndpointArgs and GetAuthScopeServiceCatalogEndpointOutput values.
// You can construct a concrete instance of `GetAuthScopeServiceCatalogEndpointInput` via:
//
//	GetAuthScopeServiceCatalogEndpointArgs{...}
type GetAuthScopeServiceCatalogEndpointInput interface {
	pulumi.Input

	ToGetAuthScopeServiceCatalogEndpointOutput() GetAuthScopeServiceCatalogEndpointOutput
	ToGetAuthScopeServiceCatalogEndpointOutputWithContext(context.Context) GetAuthScopeServiceCatalogEndpointOutput
}

type GetAuthScopeServiceCatalogEndpointArgs struct {
	// The ID of the endpoint.
	Id pulumi.StringInput `pulumi:"id"`
	// The interface of the endpoint.
	Interface pulumi.StringInput `pulumi:"interface"`
	// The region in which to obtain the V3 Identity client.
	// A Identity client is needed to retrieve tokens IDs. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringInput `pulumi:"region"`
	// The region ID of the endpoint.
	RegionId pulumi.StringInput `pulumi:"regionId"`
	// The URL of the endpoint.
	Url pulumi.StringInput `pulumi:"url"`
}

func (GetAuthScopeServiceCatalogEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeServiceCatalogEndpoint)(nil)).Elem()
}

func (i GetAuthScopeServiceCatalogEndpointArgs) ToGetAuthScopeServiceCatalogEndpointOutput() GetAuthScopeServiceCatalogEndpointOutput {
	return i.ToGetAuthScopeServiceCatalogEndpointOutputWithContext(context.Background())
}

func (i GetAuthScopeServiceCatalogEndpointArgs) ToGetAuthScopeServiceCatalogEndpointOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAuthScopeServiceCatalogEndpointOutput)
}

// GetAuthScopeServiceCatalogEndpointArrayInput is an input type that accepts GetAuthScopeServiceCatalogEndpointArray and GetAuthScopeServiceCatalogEndpointArrayOutput values.
// You can construct a concrete instance of `GetAuthScopeServiceCatalogEndpointArrayInput` via:
//
//	GetAuthScopeServiceCatalogEndpointArray{ GetAuthScopeServiceCatalogEndpointArgs{...} }
type GetAuthScopeServiceCatalogEndpointArrayInput interface {
	pulumi.Input

	ToGetAuthScopeServiceCatalogEndpointArrayOutput() GetAuthScopeServiceCatalogEndpointArrayOutput
	ToGetAuthScopeServiceCatalogEndpointArrayOutputWithContext(context.Context) GetAuthScopeServiceCatalogEndpointArrayOutput
}

type GetAuthScopeServiceCatalogEndpointArray []GetAuthScopeServiceCatalogEndpointInput

func (GetAuthScopeServiceCatalogEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAuthScopeServiceCatalogEndpoint)(nil)).Elem()
}

func (i GetAuthScopeServiceCatalogEndpointArray) ToGetAuthScopeServiceCatalogEndpointArrayOutput() GetAuthScopeServiceCatalogEndpointArrayOutput {
	return i.ToGetAuthScopeServiceCatalogEndpointArrayOutputWithContext(context.Background())
}

func (i GetAuthScopeServiceCatalogEndpointArray) ToGetAuthScopeServiceCatalogEndpointArrayOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetAuthScopeServiceCatalogEndpointArrayOutput)
}

type GetAuthScopeServiceCatalogEndpointOutput struct{ *pulumi.OutputState }

func (GetAuthScopeServiceCatalogEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeServiceCatalogEndpoint)(nil)).Elem()
}

func (o GetAuthScopeServiceCatalogEndpointOutput) ToGetAuthScopeServiceCatalogEndpointOutput() GetAuthScopeServiceCatalogEndpointOutput {
	return o
}

func (o GetAuthScopeServiceCatalogEndpointOutput) ToGetAuthScopeServiceCatalogEndpointOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogEndpointOutput {
	return o
}

// The ID of the endpoint.
func (o GetAuthScopeServiceCatalogEndpointOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalogEndpoint) string { return v.Id }).(pulumi.StringOutput)
}

// The interface of the endpoint.
func (o GetAuthScopeServiceCatalogEndpointOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalogEndpoint) string { return v.Interface }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Identity client.
// A Identity client is needed to retrieve tokens IDs. If omitted, the
// `region` argument of the provider is used.
func (o GetAuthScopeServiceCatalogEndpointOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalogEndpoint) string { return v.Region }).(pulumi.StringOutput)
}

// The region ID of the endpoint.
func (o GetAuthScopeServiceCatalogEndpointOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalogEndpoint) string { return v.RegionId }).(pulumi.StringOutput)
}

// The URL of the endpoint.
func (o GetAuthScopeServiceCatalogEndpointOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeServiceCatalogEndpoint) string { return v.Url }).(pulumi.StringOutput)
}

type GetAuthScopeServiceCatalogEndpointArrayOutput struct{ *pulumi.OutputState }

func (GetAuthScopeServiceCatalogEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetAuthScopeServiceCatalogEndpoint)(nil)).Elem()
}

func (o GetAuthScopeServiceCatalogEndpointArrayOutput) ToGetAuthScopeServiceCatalogEndpointArrayOutput() GetAuthScopeServiceCatalogEndpointArrayOutput {
	return o
}

func (o GetAuthScopeServiceCatalogEndpointArrayOutput) ToGetAuthScopeServiceCatalogEndpointArrayOutputWithContext(ctx context.Context) GetAuthScopeServiceCatalogEndpointArrayOutput {
	return o
}

func (o GetAuthScopeServiceCatalogEndpointArrayOutput) Index(i pulumi.IntInput) GetAuthScopeServiceCatalogEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetAuthScopeServiceCatalogEndpoint {
		return vs[0].([]GetAuthScopeServiceCatalogEndpoint)[vs[1].(int)]
	}).(GetAuthScopeServiceCatalogEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialAccessRuleInput)(nil)).Elem(), ApplicationCredentialAccessRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationCredentialAccessRuleArrayInput)(nil)).Elem(), ApplicationCredentialAccessRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMultiFactorAuthRuleInput)(nil)).Elem(), UserMultiFactorAuthRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMultiFactorAuthRuleArrayInput)(nil)).Elem(), UserMultiFactorAuthRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAuthScopeRoleInput)(nil)).Elem(), GetAuthScopeRoleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAuthScopeRoleArrayInput)(nil)).Elem(), GetAuthScopeRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAuthScopeServiceCatalogInput)(nil)).Elem(), GetAuthScopeServiceCatalogArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAuthScopeServiceCatalogArrayInput)(nil)).Elem(), GetAuthScopeServiceCatalogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAuthScopeServiceCatalogEndpointInput)(nil)).Elem(), GetAuthScopeServiceCatalogEndpointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetAuthScopeServiceCatalogEndpointArrayInput)(nil)).Elem(), GetAuthScopeServiceCatalogEndpointArray{})
	pulumi.RegisterOutputType(ApplicationCredentialAccessRuleOutput{})
	pulumi.RegisterOutputType(ApplicationCredentialAccessRuleArrayOutput{})
	pulumi.RegisterOutputType(UserMultiFactorAuthRuleOutput{})
	pulumi.RegisterOutputType(UserMultiFactorAuthRuleArrayOutput{})
	pulumi.RegisterOutputType(GetAuthScopeRoleOutput{})
	pulumi.RegisterOutputType(GetAuthScopeRoleArrayOutput{})
	pulumi.RegisterOutputType(GetAuthScopeServiceCatalogOutput{})
	pulumi.RegisterOutputType(GetAuthScopeServiceCatalogArrayOutput{})
	pulumi.RegisterOutputType(GetAuthScopeServiceCatalogEndpointOutput{})
	pulumi.RegisterOutputType(GetAuthScopeServiceCatalogEndpointArrayOutput{})
}
