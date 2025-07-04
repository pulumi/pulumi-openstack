// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 account resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := objectstorage.NewAccountV1(ctx, "account_1", &objectstorage.AccountV1Args{
//				Region: pulumi.String("RegionOne"),
//				Metadata: pulumi.StringMap{
//					"Temp-Url-Key": pulumi.String("testkey"),
//					"test":         pulumi.String("true"),
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
// This resource can be imported by specifying the project ID of the account:
//
// ```sh
// $ pulumi import openstack:objectstorage/accountV1:AccountV1 account_1 1202b3d0aaa44cfc8b79475c007b0711
// ```
type AccountV1 struct {
	pulumi.CustomResourceState

	// The number of bytes used by the account.
	BytesUsed pulumi.IntOutput `pulumi:"bytesUsed"`
	// The number of containers in the account.
	ContainerCount pulumi.IntOutput `pulumi:"containerCount"`
	// A map of headers returned for the account.
	Headers pulumi.StringMapOutput `pulumi:"headers"`
	// A map of custom key/value pairs to associate with the
	// account metadata. Changing the `Quota-Bytes` key value is allowed to be
	// updated only by the cloud administrator.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The number of objects in the account.
	ObjectCount pulumi.IntOutput `pulumi:"objectCount"`
	// The project ID of the corresponding account. If
	// omitted, the token's project ID is used. Changing this creates a new account.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The number of bytes allowed for the account.
	QuotaBytes pulumi.IntOutput `pulumi:"quotaBytes"`
	// The region in which to create the account. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// account.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewAccountV1 registers a new resource with the given unique name, arguments, and options.
func NewAccountV1(ctx *pulumi.Context,
	name string, args *AccountV1Args, opts ...pulumi.ResourceOption) (*AccountV1, error) {
	if args == nil {
		args = &AccountV1Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccountV1
	err := ctx.RegisterResource("openstack:objectstorage/accountV1:AccountV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountV1 gets an existing AccountV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountV1State, opts ...pulumi.ResourceOption) (*AccountV1, error) {
	var resource AccountV1
	err := ctx.ReadResource("openstack:objectstorage/accountV1:AccountV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountV1 resources.
type accountV1State struct {
	// The number of bytes used by the account.
	BytesUsed *int `pulumi:"bytesUsed"`
	// The number of containers in the account.
	ContainerCount *int `pulumi:"containerCount"`
	// A map of headers returned for the account.
	Headers map[string]string `pulumi:"headers"`
	// A map of custom key/value pairs to associate with the
	// account metadata. Changing the `Quota-Bytes` key value is allowed to be
	// updated only by the cloud administrator.
	Metadata map[string]string `pulumi:"metadata"`
	// The number of objects in the account.
	ObjectCount *int `pulumi:"objectCount"`
	// The project ID of the corresponding account. If
	// omitted, the token's project ID is used. Changing this creates a new account.
	ProjectId *string `pulumi:"projectId"`
	// The number of bytes allowed for the account.
	QuotaBytes *int `pulumi:"quotaBytes"`
	// The region in which to create the account. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// account.
	Region *string `pulumi:"region"`
}

type AccountV1State struct {
	// The number of bytes used by the account.
	BytesUsed pulumi.IntPtrInput
	// The number of containers in the account.
	ContainerCount pulumi.IntPtrInput
	// A map of headers returned for the account.
	Headers pulumi.StringMapInput
	// A map of custom key/value pairs to associate with the
	// account metadata. Changing the `Quota-Bytes` key value is allowed to be
	// updated only by the cloud administrator.
	Metadata pulumi.StringMapInput
	// The number of objects in the account.
	ObjectCount pulumi.IntPtrInput
	// The project ID of the corresponding account. If
	// omitted, the token's project ID is used. Changing this creates a new account.
	ProjectId pulumi.StringPtrInput
	// The number of bytes allowed for the account.
	QuotaBytes pulumi.IntPtrInput
	// The region in which to create the account. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// account.
	Region pulumi.StringPtrInput
}

func (AccountV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*accountV1State)(nil)).Elem()
}

type accountV1Args struct {
	// A map of custom key/value pairs to associate with the
	// account metadata. Changing the `Quota-Bytes` key value is allowed to be
	// updated only by the cloud administrator.
	Metadata map[string]string `pulumi:"metadata"`
	// The project ID of the corresponding account. If
	// omitted, the token's project ID is used. Changing this creates a new account.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to create the account. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// account.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a AccountV1 resource.
type AccountV1Args struct {
	// A map of custom key/value pairs to associate with the
	// account metadata. Changing the `Quota-Bytes` key value is allowed to be
	// updated only by the cloud administrator.
	Metadata pulumi.StringMapInput
	// The project ID of the corresponding account. If
	// omitted, the token's project ID is used. Changing this creates a new account.
	ProjectId pulumi.StringPtrInput
	// The region in which to create the account. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// account.
	Region pulumi.StringPtrInput
}

func (AccountV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*accountV1Args)(nil)).Elem()
}

type AccountV1Input interface {
	pulumi.Input

	ToAccountV1Output() AccountV1Output
	ToAccountV1OutputWithContext(ctx context.Context) AccountV1Output
}

func (*AccountV1) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountV1)(nil)).Elem()
}

func (i *AccountV1) ToAccountV1Output() AccountV1Output {
	return i.ToAccountV1OutputWithContext(context.Background())
}

func (i *AccountV1) ToAccountV1OutputWithContext(ctx context.Context) AccountV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(AccountV1Output)
}

// AccountV1ArrayInput is an input type that accepts AccountV1Array and AccountV1ArrayOutput values.
// You can construct a concrete instance of `AccountV1ArrayInput` via:
//
//	AccountV1Array{ AccountV1Args{...} }
type AccountV1ArrayInput interface {
	pulumi.Input

	ToAccountV1ArrayOutput() AccountV1ArrayOutput
	ToAccountV1ArrayOutputWithContext(context.Context) AccountV1ArrayOutput
}

type AccountV1Array []AccountV1Input

func (AccountV1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountV1)(nil)).Elem()
}

func (i AccountV1Array) ToAccountV1ArrayOutput() AccountV1ArrayOutput {
	return i.ToAccountV1ArrayOutputWithContext(context.Background())
}

func (i AccountV1Array) ToAccountV1ArrayOutputWithContext(ctx context.Context) AccountV1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountV1ArrayOutput)
}

// AccountV1MapInput is an input type that accepts AccountV1Map and AccountV1MapOutput values.
// You can construct a concrete instance of `AccountV1MapInput` via:
//
//	AccountV1Map{ "key": AccountV1Args{...} }
type AccountV1MapInput interface {
	pulumi.Input

	ToAccountV1MapOutput() AccountV1MapOutput
	ToAccountV1MapOutputWithContext(context.Context) AccountV1MapOutput
}

type AccountV1Map map[string]AccountV1Input

func (AccountV1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountV1)(nil)).Elem()
}

func (i AccountV1Map) ToAccountV1MapOutput() AccountV1MapOutput {
	return i.ToAccountV1MapOutputWithContext(context.Background())
}

func (i AccountV1Map) ToAccountV1MapOutputWithContext(ctx context.Context) AccountV1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountV1MapOutput)
}

type AccountV1Output struct{ *pulumi.OutputState }

func (AccountV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountV1)(nil)).Elem()
}

func (o AccountV1Output) ToAccountV1Output() AccountV1Output {
	return o
}

func (o AccountV1Output) ToAccountV1OutputWithContext(ctx context.Context) AccountV1Output {
	return o
}

// The number of bytes used by the account.
func (o AccountV1Output) BytesUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.IntOutput { return v.BytesUsed }).(pulumi.IntOutput)
}

// The number of containers in the account.
func (o AccountV1Output) ContainerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.IntOutput { return v.ContainerCount }).(pulumi.IntOutput)
}

// A map of headers returned for the account.
func (o AccountV1Output) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.StringMapOutput { return v.Headers }).(pulumi.StringMapOutput)
}

// A map of custom key/value pairs to associate with the
// account metadata. Changing the `Quota-Bytes` key value is allowed to be
// updated only by the cloud administrator.
func (o AccountV1Output) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The number of objects in the account.
func (o AccountV1Output) ObjectCount() pulumi.IntOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.IntOutput { return v.ObjectCount }).(pulumi.IntOutput)
}

// The project ID of the corresponding account. If
// omitted, the token's project ID is used. Changing this creates a new account.
func (o AccountV1Output) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The number of bytes allowed for the account.
func (o AccountV1Output) QuotaBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.IntOutput { return v.QuotaBytes }).(pulumi.IntOutput)
}

// The region in which to create the account. If omitted,
// the `region` argument of the provider is used. Changing this creates a new
// account.
func (o AccountV1Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AccountV1) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type AccountV1ArrayOutput struct{ *pulumi.OutputState }

func (AccountV1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccountV1)(nil)).Elem()
}

func (o AccountV1ArrayOutput) ToAccountV1ArrayOutput() AccountV1ArrayOutput {
	return o
}

func (o AccountV1ArrayOutput) ToAccountV1ArrayOutputWithContext(ctx context.Context) AccountV1ArrayOutput {
	return o
}

func (o AccountV1ArrayOutput) Index(i pulumi.IntInput) AccountV1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccountV1 {
		return vs[0].([]*AccountV1)[vs[1].(int)]
	}).(AccountV1Output)
}

type AccountV1MapOutput struct{ *pulumi.OutputState }

func (AccountV1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccountV1)(nil)).Elem()
}

func (o AccountV1MapOutput) ToAccountV1MapOutput() AccountV1MapOutput {
	return o
}

func (o AccountV1MapOutput) ToAccountV1MapOutputWithContext(ctx context.Context) AccountV1MapOutput {
	return o
}

func (o AccountV1MapOutput) MapIndex(k pulumi.StringInput) AccountV1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccountV1 {
		return vs[0].(map[string]*AccountV1)[vs[1].(string)]
	}).(AccountV1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccountV1Input)(nil)).Elem(), &AccountV1{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountV1ArrayInput)(nil)).Elem(), AccountV1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccountV1MapInput)(nil)).Elem(), AccountV1Map{})
	pulumi.RegisterOutputType(AccountV1Output{})
	pulumi.RegisterOutputType(AccountV1ArrayOutput{})
	pulumi.RegisterOutputType(AccountV1MapOutput{})
}
