// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAuthScope(ctx *pulumi.Context, args *GetAuthScopeArgs, opts ...pulumi.InvokeOption) (*GetAuthScopeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAuthScopeResult
	err := ctx.Invoke("openstack:identity/getAuthScope:getAuthScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthScope.
type GetAuthScopeArgs struct {
	// The name of the scope. This is an arbitrary name which is
	// only used as a unique identifier so an actual token isn't used as the ID.
	Name string `pulumi:"name"`
	// The region in which to obtain the V3 Identity client.
	// A Identity client is needed to retrieve tokens IDs. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// A boolean argument that determines whether to
	// export the current auth scope token ID. When set to `true`, the `tokenId`
	// attribute will contain an unencrypted token that can be used for further API
	// calls. **Warning**: please note that the leaked token may allow unauthorized
	// access to other OpenStack services within the current auth scope, so use this
	// option with caution.
	SetTokenId *bool `pulumi:"setTokenId"`
}

// A collection of values returned by getAuthScope.
type GetAuthScopeResult struct {
	// The domain ID of the scope.
	DomainId string `pulumi:"domainId"`
	// The domain name of the scope.
	DomainName string `pulumi:"domainName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the service.
	Name string `pulumi:"name"`
	// The domain ID of the project.
	ProjectDomainId string `pulumi:"projectDomainId"`
	// The domain name of the project.
	ProjectDomainName string `pulumi:"projectDomainName"`
	// The project ID of the scope.
	ProjectId string `pulumi:"projectId"`
	// The project name of the scope.
	ProjectName string `pulumi:"projectName"`
	// The region of the endpoint.
	Region string `pulumi:"region"`
	// A list of roles in the current scope. See reference below.
	Roles []GetAuthScopeRole `pulumi:"roles"`
	// A list of service catalog entries returned with the token.
	ServiceCatalogs []GetAuthScopeServiceCatalog `pulumi:"serviceCatalogs"`
	SetTokenId      *bool                        `pulumi:"setTokenId"`
	// The token ID of the scope.
	TokenId string `pulumi:"tokenId"`
	// The domain ID of the user.
	UserDomainId string `pulumi:"userDomainId"`
	// The domain name of the user.
	UserDomainName string `pulumi:"userDomainName"`
	// The user ID the of the scope.
	UserId string `pulumi:"userId"`
	// The username of the scope.
	UserName string `pulumi:"userName"`
}

func GetAuthScopeOutput(ctx *pulumi.Context, args GetAuthScopeOutputArgs, opts ...pulumi.InvokeOption) GetAuthScopeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetAuthScopeResultOutput, error) {
			args := v.(GetAuthScopeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:identity/getAuthScope:getAuthScope", args, GetAuthScopeResultOutput{}, options).(GetAuthScopeResultOutput), nil
		}).(GetAuthScopeResultOutput)
}

// A collection of arguments for invoking getAuthScope.
type GetAuthScopeOutputArgs struct {
	// The name of the scope. This is an arbitrary name which is
	// only used as a unique identifier so an actual token isn't used as the ID.
	Name pulumi.StringInput `pulumi:"name"`
	// The region in which to obtain the V3 Identity client.
	// A Identity client is needed to retrieve tokens IDs. If omitted, the
	// `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// A boolean argument that determines whether to
	// export the current auth scope token ID. When set to `true`, the `tokenId`
	// attribute will contain an unencrypted token that can be used for further API
	// calls. **Warning**: please note that the leaked token may allow unauthorized
	// access to other OpenStack services within the current auth scope, so use this
	// option with caution.
	SetTokenId pulumi.BoolPtrInput `pulumi:"setTokenId"`
}

func (GetAuthScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeArgs)(nil)).Elem()
}

// A collection of values returned by getAuthScope.
type GetAuthScopeResultOutput struct{ *pulumi.OutputState }

func (GetAuthScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthScopeResult)(nil)).Elem()
}

func (o GetAuthScopeResultOutput) ToGetAuthScopeResultOutput() GetAuthScopeResultOutput {
	return o
}

func (o GetAuthScopeResultOutput) ToGetAuthScopeResultOutputWithContext(ctx context.Context) GetAuthScopeResultOutput {
	return o
}

// The domain ID of the scope.
func (o GetAuthScopeResultOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.DomainId }).(pulumi.StringOutput)
}

// The domain name of the scope.
func (o GetAuthScopeResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAuthScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the service.
func (o GetAuthScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The domain ID of the project.
func (o GetAuthScopeResultOutput) ProjectDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.ProjectDomainId }).(pulumi.StringOutput)
}

// The domain name of the project.
func (o GetAuthScopeResultOutput) ProjectDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.ProjectDomainName }).(pulumi.StringOutput)
}

// The project ID of the scope.
func (o GetAuthScopeResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The project name of the scope.
func (o GetAuthScopeResultOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.ProjectName }).(pulumi.StringOutput)
}

// The region of the endpoint.
func (o GetAuthScopeResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.Region }).(pulumi.StringOutput)
}

// A list of roles in the current scope. See reference below.
func (o GetAuthScopeResultOutput) Roles() GetAuthScopeRoleArrayOutput {
	return o.ApplyT(func(v GetAuthScopeResult) []GetAuthScopeRole { return v.Roles }).(GetAuthScopeRoleArrayOutput)
}

// A list of service catalog entries returned with the token.
func (o GetAuthScopeResultOutput) ServiceCatalogs() GetAuthScopeServiceCatalogArrayOutput {
	return o.ApplyT(func(v GetAuthScopeResult) []GetAuthScopeServiceCatalog { return v.ServiceCatalogs }).(GetAuthScopeServiceCatalogArrayOutput)
}

func (o GetAuthScopeResultOutput) SetTokenId() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAuthScopeResult) *bool { return v.SetTokenId }).(pulumi.BoolPtrOutput)
}

// The token ID of the scope.
func (o GetAuthScopeResultOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.TokenId }).(pulumi.StringOutput)
}

// The domain ID of the user.
func (o GetAuthScopeResultOutput) UserDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.UserDomainId }).(pulumi.StringOutput)
}

// The domain name of the user.
func (o GetAuthScopeResultOutput) UserDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.UserDomainName }).(pulumi.StringOutput)
}

// The user ID the of the scope.
func (o GetAuthScopeResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.UserId }).(pulumi.StringOutput)
}

// The username of the scope.
func (o GetAuthScopeResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthScopeResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAuthScopeResultOutput{})
}
