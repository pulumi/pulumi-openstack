// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### Append a security group to an existing port
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
//			systemPort, err := networking.LookupPort(ctx, &networking.LookupPortArgs{
//				FixedIp: pulumi.StringRef("10.0.0.10"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			secgroup, err := networking.LookupSecGroup(ctx, &networking.LookupSecGroupArgs{
//				Name: pulumi.StringRef("secgroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewPortSecGroupAssociate(ctx, "port_1", &networking.PortSecGroupAssociateArgs{
//				PortId: pulumi.String(systemPort.Id),
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.String(secgroup.Id),
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
// ### Enforce a security group to an existing port
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
//			systemPort, err := networking.LookupPort(ctx, &networking.LookupPortArgs{
//				FixedIp: pulumi.StringRef("10.0.0.10"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			secgroup, err := networking.LookupSecGroup(ctx, &networking.LookupSecGroupArgs{
//				Name: pulumi.StringRef("secgroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewPortSecGroupAssociate(ctx, "port_1", &networking.PortSecGroupAssociateArgs{
//				PortId:  pulumi.String(systemPort.Id),
//				Enforce: pulumi.Bool(true),
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.String(secgroup.Id),
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
// ### Remove all security groups from an existing port
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
//			systemPort, err := networking.LookupPort(ctx, &networking.LookupPortArgs{
//				FixedIp: pulumi.StringRef("10.0.0.10"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = networking.NewPortSecGroupAssociate(ctx, "port_1", &networking.PortSecGroupAssociateArgs{
//				PortId:           pulumi.String(systemPort.Id),
//				Enforce:          pulumi.Bool(true),
//				SecurityGroupIds: pulumi.StringArray{},
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
// Port security group association can be imported using the `id` of the port, e.g.
//
// ```sh
// $ pulumi import openstack:networking/portSecGroupAssociate:PortSecGroupAssociate port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
// ```
type PortSecGroupAssociate struct {
	pulumi.CustomResourceState

	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds pulumi.StringArrayOutput `pulumi:"allSecurityGroupIds"`
	// Whether to replace or append the list of security
	// groups, specified in the `securityGroupIds`. Defaults to `false`.
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// An UUID of the port to apply security groups to.
	PortId pulumi.StringOutput `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
}

// NewPortSecGroupAssociate registers a new resource with the given unique name, arguments, and options.
func NewPortSecGroupAssociate(ctx *pulumi.Context,
	name string, args *PortSecGroupAssociateArgs, opts ...pulumi.ResourceOption) (*PortSecGroupAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PortId == nil {
		return nil, errors.New("invalid value for required argument 'PortId'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PortSecGroupAssociate
	err := ctx.RegisterResource("openstack:networking/portSecGroupAssociate:PortSecGroupAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortSecGroupAssociate gets an existing PortSecGroupAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortSecGroupAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortSecGroupAssociateState, opts ...pulumi.ResourceOption) (*PortSecGroupAssociate, error) {
	var resource PortSecGroupAssociate
	err := ctx.ReadResource("openstack:networking/portSecGroupAssociate:PortSecGroupAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortSecGroupAssociate resources.
type portSecGroupAssociateState struct {
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds []string `pulumi:"allSecurityGroupIds"`
	// Whether to replace or append the list of security
	// groups, specified in the `securityGroupIds`. Defaults to `false`.
	Enforce *bool `pulumi:"enforce"`
	// An UUID of the port to apply security groups to.
	PortId *string `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// resource.
	Region *string `pulumi:"region"`
	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
}

type PortSecGroupAssociateState struct {
	// The collection of Security Group IDs on the port
	// which have been explicitly and implicitly added.
	AllSecurityGroupIds pulumi.StringArrayInput
	// Whether to replace or append the list of security
	// groups, specified in the `securityGroupIds`. Defaults to `false`.
	Enforce pulumi.BoolPtrInput
	// An UUID of the port to apply security groups to.
	PortId pulumi.StringPtrInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// resource.
	Region pulumi.StringPtrInput
	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
}

func (PortSecGroupAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*portSecGroupAssociateState)(nil)).Elem()
}

type portSecGroupAssociateArgs struct {
	// Whether to replace or append the list of security
	// groups, specified in the `securityGroupIds`. Defaults to `false`.
	Enforce *bool `pulumi:"enforce"`
	// An UUID of the port to apply security groups to.
	PortId string `pulumi:"portId"`
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// resource.
	Region *string `pulumi:"region"`
	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
}

// The set of arguments for constructing a PortSecGroupAssociate resource.
type PortSecGroupAssociateArgs struct {
	// Whether to replace or append the list of security
	// groups, specified in the `securityGroupIds`. Defaults to `false`.
	Enforce pulumi.BoolPtrInput
	// An UUID of the port to apply security groups to.
	PortId pulumi.StringInput
	// The region in which to obtain the V2 networking client.
	// A networking client is needed to manage a port. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// resource.
	Region pulumi.StringPtrInput
	// A list of security group IDs to apply to
	// the port. The security groups must be specified by ID and not name (as
	// opposed to how they are configured with the Compute Instance).
	SecurityGroupIds pulumi.StringArrayInput
}

func (PortSecGroupAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portSecGroupAssociateArgs)(nil)).Elem()
}

type PortSecGroupAssociateInput interface {
	pulumi.Input

	ToPortSecGroupAssociateOutput() PortSecGroupAssociateOutput
	ToPortSecGroupAssociateOutputWithContext(ctx context.Context) PortSecGroupAssociateOutput
}

func (*PortSecGroupAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**PortSecGroupAssociate)(nil)).Elem()
}

func (i *PortSecGroupAssociate) ToPortSecGroupAssociateOutput() PortSecGroupAssociateOutput {
	return i.ToPortSecGroupAssociateOutputWithContext(context.Background())
}

func (i *PortSecGroupAssociate) ToPortSecGroupAssociateOutputWithContext(ctx context.Context) PortSecGroupAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortSecGroupAssociateOutput)
}

// PortSecGroupAssociateArrayInput is an input type that accepts PortSecGroupAssociateArray and PortSecGroupAssociateArrayOutput values.
// You can construct a concrete instance of `PortSecGroupAssociateArrayInput` via:
//
//	PortSecGroupAssociateArray{ PortSecGroupAssociateArgs{...} }
type PortSecGroupAssociateArrayInput interface {
	pulumi.Input

	ToPortSecGroupAssociateArrayOutput() PortSecGroupAssociateArrayOutput
	ToPortSecGroupAssociateArrayOutputWithContext(context.Context) PortSecGroupAssociateArrayOutput
}

type PortSecGroupAssociateArray []PortSecGroupAssociateInput

func (PortSecGroupAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortSecGroupAssociate)(nil)).Elem()
}

func (i PortSecGroupAssociateArray) ToPortSecGroupAssociateArrayOutput() PortSecGroupAssociateArrayOutput {
	return i.ToPortSecGroupAssociateArrayOutputWithContext(context.Background())
}

func (i PortSecGroupAssociateArray) ToPortSecGroupAssociateArrayOutputWithContext(ctx context.Context) PortSecGroupAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortSecGroupAssociateArrayOutput)
}

// PortSecGroupAssociateMapInput is an input type that accepts PortSecGroupAssociateMap and PortSecGroupAssociateMapOutput values.
// You can construct a concrete instance of `PortSecGroupAssociateMapInput` via:
//
//	PortSecGroupAssociateMap{ "key": PortSecGroupAssociateArgs{...} }
type PortSecGroupAssociateMapInput interface {
	pulumi.Input

	ToPortSecGroupAssociateMapOutput() PortSecGroupAssociateMapOutput
	ToPortSecGroupAssociateMapOutputWithContext(context.Context) PortSecGroupAssociateMapOutput
}

type PortSecGroupAssociateMap map[string]PortSecGroupAssociateInput

func (PortSecGroupAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortSecGroupAssociate)(nil)).Elem()
}

func (i PortSecGroupAssociateMap) ToPortSecGroupAssociateMapOutput() PortSecGroupAssociateMapOutput {
	return i.ToPortSecGroupAssociateMapOutputWithContext(context.Background())
}

func (i PortSecGroupAssociateMap) ToPortSecGroupAssociateMapOutputWithContext(ctx context.Context) PortSecGroupAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortSecGroupAssociateMapOutput)
}

type PortSecGroupAssociateOutput struct{ *pulumi.OutputState }

func (PortSecGroupAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PortSecGroupAssociate)(nil)).Elem()
}

func (o PortSecGroupAssociateOutput) ToPortSecGroupAssociateOutput() PortSecGroupAssociateOutput {
	return o
}

func (o PortSecGroupAssociateOutput) ToPortSecGroupAssociateOutputWithContext(ctx context.Context) PortSecGroupAssociateOutput {
	return o
}

// The collection of Security Group IDs on the port
// which have been explicitly and implicitly added.
func (o PortSecGroupAssociateOutput) AllSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PortSecGroupAssociate) pulumi.StringArrayOutput { return v.AllSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Whether to replace or append the list of security
// groups, specified in the `securityGroupIds`. Defaults to `false`.
func (o PortSecGroupAssociateOutput) Enforce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PortSecGroupAssociate) pulumi.BoolPtrOutput { return v.Enforce }).(pulumi.BoolPtrOutput)
}

// An UUID of the port to apply security groups to.
func (o PortSecGroupAssociateOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *PortSecGroupAssociate) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 networking client.
// A networking client is needed to manage a port. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// resource.
func (o PortSecGroupAssociateOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PortSecGroupAssociate) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of security group IDs to apply to
// the port. The security groups must be specified by ID and not name (as
// opposed to how they are configured with the Compute Instance).
func (o PortSecGroupAssociateOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PortSecGroupAssociate) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

type PortSecGroupAssociateArrayOutput struct{ *pulumi.OutputState }

func (PortSecGroupAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PortSecGroupAssociate)(nil)).Elem()
}

func (o PortSecGroupAssociateArrayOutput) ToPortSecGroupAssociateArrayOutput() PortSecGroupAssociateArrayOutput {
	return o
}

func (o PortSecGroupAssociateArrayOutput) ToPortSecGroupAssociateArrayOutputWithContext(ctx context.Context) PortSecGroupAssociateArrayOutput {
	return o
}

func (o PortSecGroupAssociateArrayOutput) Index(i pulumi.IntInput) PortSecGroupAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PortSecGroupAssociate {
		return vs[0].([]*PortSecGroupAssociate)[vs[1].(int)]
	}).(PortSecGroupAssociateOutput)
}

type PortSecGroupAssociateMapOutput struct{ *pulumi.OutputState }

func (PortSecGroupAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PortSecGroupAssociate)(nil)).Elem()
}

func (o PortSecGroupAssociateMapOutput) ToPortSecGroupAssociateMapOutput() PortSecGroupAssociateMapOutput {
	return o
}

func (o PortSecGroupAssociateMapOutput) ToPortSecGroupAssociateMapOutputWithContext(ctx context.Context) PortSecGroupAssociateMapOutput {
	return o
}

func (o PortSecGroupAssociateMapOutput) MapIndex(k pulumi.StringInput) PortSecGroupAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PortSecGroupAssociate {
		return vs[0].(map[string]*PortSecGroupAssociate)[vs[1].(string)]
	}).(PortSecGroupAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PortSecGroupAssociateInput)(nil)).Elem(), &PortSecGroupAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortSecGroupAssociateArrayInput)(nil)).Elem(), PortSecGroupAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PortSecGroupAssociateMapInput)(nil)).Elem(), PortSecGroupAssociateMap{})
	pulumi.RegisterOutputType(PortSecGroupAssociateOutput{})
	pulumi.RegisterOutputType(PortSecGroupAssociateArrayOutput{})
	pulumi.RegisterOutputType(PortSecGroupAssociateMapOutput{})
}
