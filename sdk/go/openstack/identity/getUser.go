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

// Use this data source to get the ID of an OpenStack user.
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("openstack:identity/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The domain this user belongs to.
	DomainId *string `pulumi:"domainId"`
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// The identity provider ID of the user.
	IdpId *string `pulumi:"idpId"`
	// The name of the user.
	Name *string `pulumi:"name"`
	// Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
	PasswordExpiresAt *string `pulumi:"passwordExpiresAt"`
	// The protocol ID of the user.
	ProtocolId *string `pulumi:"protocolId"`
	// The region the user is located in.
	Region *string `pulumi:"region"`
	// The unique ID of the user.
	UniqueId *string `pulumi:"uniqueId"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// See Argument Reference above.
	DefaultProjectId string `pulumi:"defaultProjectId"`
	// A description of the user.
	Description string `pulumi:"description"`
	// See Argument Reference above.
	DomainId string `pulumi:"domainId"`
	// See Argument Reference above.
	Enabled *bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IdpId *string `pulumi:"idpId"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	PasswordExpiresAt *string `pulumi:"passwordExpiresAt"`
	// See Argument Reference above.
	ProtocolId *string `pulumi:"protocolId"`
	// The region the user is located in.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	UniqueId *string `pulumi:"uniqueId"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// The domain this user belongs to.
	DomainId pulumi.StringPtrInput `pulumi:"domainId"`
	// Whether the user is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
	// The identity provider ID of the user.
	IdpId pulumi.StringPtrInput `pulumi:"idpId"`
	// The name of the user.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
	PasswordExpiresAt pulumi.StringPtrInput `pulumi:"passwordExpiresAt"`
	// The protocol ID of the user.
	ProtocolId pulumi.StringPtrInput `pulumi:"protocolId"`
	// The region the user is located in.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The unique ID of the user.
	UniqueId pulumi.StringPtrInput `pulumi:"uniqueId"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserResult] {
	return pulumix.Output[LookupUserResult]{
		OutputState: o.OutputState,
	}
}

// See Argument Reference above.
func (o LookupUserResultOutput) DefaultProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.DefaultProjectId }).(pulumi.StringOutput)
}

// A description of the user.
func (o LookupUserResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Description }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.DomainId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) IdpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.IdpId }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) PasswordExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.PasswordExpiresAt }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) ProtocolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.ProtocolId }).(pulumi.StringPtrOutput)
}

// The region the user is located in.
func (o LookupUserResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) UniqueId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.UniqueId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
