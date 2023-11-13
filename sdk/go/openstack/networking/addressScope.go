// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Neutron addressscope resource within OpenStack.
//
// ## Example Usage
// ### Create an Address-scope
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networking.NewAddressScope(ctx, "addressscope1", &networking.AddressScopeArgs{
//				IpVersion: pulumi.Int(6),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create a Subnet Pool from an Address-scope
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/networking"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			addressscope1, err := networking.NewAddressScope(ctx, "addressscope1", &networking.AddressScopeArgs{
//				IpVersion: pulumi.Int(6),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewSubnetPool(ctx, "subnetpool1", &networking.SubnetPoolArgs{
//				Prefixes: pulumi.StringArray{
//					pulumi.String("fdf7:b13d:dead:beef::/64"),
//					pulumi.String("fd65:86cc:a334:39b7::/64"),
//				},
//				AddressScopeId: addressscope1.ID(),
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
// Address-scopes can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import openstack:networking/addressScope:AddressScope addressscope_1 9cc35860-522a-4d35-974d-51d4b011801e
//
// ```
type AddressScope struct {
	pulumi.CustomResourceState

	// IP version, either 4 (default) or 6. Changing this
	// creates a new address-scope.
	IpVersion pulumi.IntPtrOutput `pulumi:"ipVersion"`
	// The name of the address-scope. Changing this updates the
	// name of the existing address-scope.
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner of the address-scope. Required if admin
	// wants to create a address-scope for another project. Changing this creates a
	// new address-scope.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron address-scope. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// address-scope.
	Region pulumi.StringOutput `pulumi:"region"`
	// Indicates whether this address-scope is shared across
	// all projects. Changing this updates the shared status of the existing
	// address-scope.
	Shared pulumi.BoolOutput `pulumi:"shared"`
}

// NewAddressScope registers a new resource with the given unique name, arguments, and options.
func NewAddressScope(ctx *pulumi.Context,
	name string, args *AddressScopeArgs, opts ...pulumi.ResourceOption) (*AddressScope, error) {
	if args == nil {
		args = &AddressScopeArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AddressScope
	err := ctx.RegisterResource("openstack:networking/addressScope:AddressScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddressScope gets an existing AddressScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddressScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressScopeState, opts ...pulumi.ResourceOption) (*AddressScope, error) {
	var resource AddressScope
	err := ctx.ReadResource("openstack:networking/addressScope:AddressScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AddressScope resources.
type addressScopeState struct {
	// IP version, either 4 (default) or 6. Changing this
	// creates a new address-scope.
	IpVersion *int `pulumi:"ipVersion"`
	// The name of the address-scope. Changing this updates the
	// name of the existing address-scope.
	Name *string `pulumi:"name"`
	// The owner of the address-scope. Required if admin
	// wants to create a address-scope for another project. Changing this creates a
	// new address-scope.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron address-scope. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// address-scope.
	Region *string `pulumi:"region"`
	// Indicates whether this address-scope is shared across
	// all projects. Changing this updates the shared status of the existing
	// address-scope.
	Shared *bool `pulumi:"shared"`
}

type AddressScopeState struct {
	// IP version, either 4 (default) or 6. Changing this
	// creates a new address-scope.
	IpVersion pulumi.IntPtrInput
	// The name of the address-scope. Changing this updates the
	// name of the existing address-scope.
	Name pulumi.StringPtrInput
	// The owner of the address-scope. Required if admin
	// wants to create a address-scope for another project. Changing this creates a
	// new address-scope.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron address-scope. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// address-scope.
	Region pulumi.StringPtrInput
	// Indicates whether this address-scope is shared across
	// all projects. Changing this updates the shared status of the existing
	// address-scope.
	Shared pulumi.BoolPtrInput
}

func (AddressScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressScopeState)(nil)).Elem()
}

type addressScopeArgs struct {
	// IP version, either 4 (default) or 6. Changing this
	// creates a new address-scope.
	IpVersion *int `pulumi:"ipVersion"`
	// The name of the address-scope. Changing this updates the
	// name of the existing address-scope.
	Name *string `pulumi:"name"`
	// The owner of the address-scope. Required if admin
	// wants to create a address-scope for another project. Changing this creates a
	// new address-scope.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron address-scope. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// address-scope.
	Region *string `pulumi:"region"`
	// Indicates whether this address-scope is shared across
	// all projects. Changing this updates the shared status of the existing
	// address-scope.
	Shared *bool `pulumi:"shared"`
}

// The set of arguments for constructing a AddressScope resource.
type AddressScopeArgs struct {
	// IP version, either 4 (default) or 6. Changing this
	// creates a new address-scope.
	IpVersion pulumi.IntPtrInput
	// The name of the address-scope. Changing this updates the
	// name of the existing address-scope.
	Name pulumi.StringPtrInput
	// The owner of the address-scope. Required if admin
	// wants to create a address-scope for another project. Changing this creates a
	// new address-scope.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron address-scope. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// address-scope.
	Region pulumi.StringPtrInput
	// Indicates whether this address-scope is shared across
	// all projects. Changing this updates the shared status of the existing
	// address-scope.
	Shared pulumi.BoolPtrInput
}

func (AddressScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressScopeArgs)(nil)).Elem()
}

type AddressScopeInput interface {
	pulumi.Input

	ToAddressScopeOutput() AddressScopeOutput
	ToAddressScopeOutputWithContext(ctx context.Context) AddressScopeOutput
}

func (*AddressScope) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressScope)(nil)).Elem()
}

func (i *AddressScope) ToAddressScopeOutput() AddressScopeOutput {
	return i.ToAddressScopeOutputWithContext(context.Background())
}

func (i *AddressScope) ToAddressScopeOutputWithContext(ctx context.Context) AddressScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressScopeOutput)
}

// AddressScopeArrayInput is an input type that accepts AddressScopeArray and AddressScopeArrayOutput values.
// You can construct a concrete instance of `AddressScopeArrayInput` via:
//
//	AddressScopeArray{ AddressScopeArgs{...} }
type AddressScopeArrayInput interface {
	pulumi.Input

	ToAddressScopeArrayOutput() AddressScopeArrayOutput
	ToAddressScopeArrayOutputWithContext(context.Context) AddressScopeArrayOutput
}

type AddressScopeArray []AddressScopeInput

func (AddressScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressScope)(nil)).Elem()
}

func (i AddressScopeArray) ToAddressScopeArrayOutput() AddressScopeArrayOutput {
	return i.ToAddressScopeArrayOutputWithContext(context.Background())
}

func (i AddressScopeArray) ToAddressScopeArrayOutputWithContext(ctx context.Context) AddressScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressScopeArrayOutput)
}

// AddressScopeMapInput is an input type that accepts AddressScopeMap and AddressScopeMapOutput values.
// You can construct a concrete instance of `AddressScopeMapInput` via:
//
//	AddressScopeMap{ "key": AddressScopeArgs{...} }
type AddressScopeMapInput interface {
	pulumi.Input

	ToAddressScopeMapOutput() AddressScopeMapOutput
	ToAddressScopeMapOutputWithContext(context.Context) AddressScopeMapOutput
}

type AddressScopeMap map[string]AddressScopeInput

func (AddressScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressScope)(nil)).Elem()
}

func (i AddressScopeMap) ToAddressScopeMapOutput() AddressScopeMapOutput {
	return i.ToAddressScopeMapOutputWithContext(context.Background())
}

func (i AddressScopeMap) ToAddressScopeMapOutputWithContext(ctx context.Context) AddressScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressScopeMapOutput)
}

type AddressScopeOutput struct{ *pulumi.OutputState }

func (AddressScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressScope)(nil)).Elem()
}

func (o AddressScopeOutput) ToAddressScopeOutput() AddressScopeOutput {
	return o
}

func (o AddressScopeOutput) ToAddressScopeOutputWithContext(ctx context.Context) AddressScopeOutput {
	return o
}

// IP version, either 4 (default) or 6. Changing this
// creates a new address-scope.
func (o AddressScopeOutput) IpVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AddressScope) pulumi.IntPtrOutput { return v.IpVersion }).(pulumi.IntPtrOutput)
}

// The name of the address-scope. Changing this updates the
// name of the existing address-scope.
func (o AddressScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The owner of the address-scope. Required if admin
// wants to create a address-scope for another project. Changing this creates a
// new address-scope.
func (o AddressScopeOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressScope) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Networking client.
// A Networking client is needed to create a Neutron address-scope. If omitted,
// the `region` argument of the provider is used. Changing this creates a new
// address-scope.
func (o AddressScopeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AddressScope) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Indicates whether this address-scope is shared across
// all projects. Changing this updates the shared status of the existing
// address-scope.
func (o AddressScopeOutput) Shared() pulumi.BoolOutput {
	return o.ApplyT(func(v *AddressScope) pulumi.BoolOutput { return v.Shared }).(pulumi.BoolOutput)
}

type AddressScopeArrayOutput struct{ *pulumi.OutputState }

func (AddressScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AddressScope)(nil)).Elem()
}

func (o AddressScopeArrayOutput) ToAddressScopeArrayOutput() AddressScopeArrayOutput {
	return o
}

func (o AddressScopeArrayOutput) ToAddressScopeArrayOutputWithContext(ctx context.Context) AddressScopeArrayOutput {
	return o
}

func (o AddressScopeArrayOutput) Index(i pulumi.IntInput) AddressScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AddressScope {
		return vs[0].([]*AddressScope)[vs[1].(int)]
	}).(AddressScopeOutput)
}

type AddressScopeMapOutput struct{ *pulumi.OutputState }

func (AddressScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AddressScope)(nil)).Elem()
}

func (o AddressScopeMapOutput) ToAddressScopeMapOutput() AddressScopeMapOutput {
	return o
}

func (o AddressScopeMapOutput) ToAddressScopeMapOutputWithContext(ctx context.Context) AddressScopeMapOutput {
	return o
}

func (o AddressScopeMapOutput) MapIndex(k pulumi.StringInput) AddressScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AddressScope {
		return vs[0].(map[string]*AddressScope)[vs[1].(string)]
	}).(AddressScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressScopeInput)(nil)).Elem(), &AddressScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressScopeArrayInput)(nil)).Elem(), AddressScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressScopeMapInput)(nil)).Elem(), AddressScopeMap{})
	pulumi.RegisterOutputType(AddressScopeOutput{})
	pulumi.RegisterOutputType(AddressScopeArrayOutput{})
	pulumi.RegisterOutputType(AddressScopeMapOutput{})
}
