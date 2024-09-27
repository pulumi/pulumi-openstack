// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V3 Project resource within OpenStack Keystone.
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
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewProject(ctx, "project_1", &identity.ProjectArgs{
//				Name:        pulumi.String("project_1"),
//				Description: pulumi.String("A project"),
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
// Projects can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:identity/project:Project project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
// ```
type Project struct {
	pulumi.CustomResourceState

	// A description of the project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain this project belongs to.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`. Default is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`. Default is `false`. Changing this creates a new
	// project/domain.
	IsDomain pulumi.BoolPtrOutput `pulumi:"isDomain"`
	// The name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of this project. Changing this creates
	// a new project.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new project.
	Region pulumi.StringOutput `pulumi:"region"`
	// Tags for the project. Changing this updates the existing
	// project.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("openstack:identity/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("openstack:identity/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// A description of the project.
	Description *string `pulumi:"description"`
	// The domain this project belongs to.
	DomainId *string `pulumi:"domainId"`
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`. Default is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`. Default is `false`. Changing this creates a new
	// project/domain.
	IsDomain *bool `pulumi:"isDomain"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The parent of this project. Changing this creates
	// a new project.
	ParentId *string `pulumi:"parentId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new project.
	Region *string `pulumi:"region"`
	// Tags for the project. Changing this updates the existing
	// project.
	Tags []string `pulumi:"tags"`
}

type ProjectState struct {
	// A description of the project.
	Description pulumi.StringPtrInput
	// The domain this project belongs to.
	DomainId pulumi.StringPtrInput
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`. Default is `true`.
	Enabled pulumi.BoolPtrInput
	// Whether this project is a domain. Valid values
	// are `true` and `false`. Default is `false`. Changing this creates a new
	// project/domain.
	IsDomain pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The parent of this project. Changing this creates
	// a new project.
	ParentId pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new project.
	Region pulumi.StringPtrInput
	// Tags for the project. Changing this updates the existing
	// project.
	Tags pulumi.StringArrayInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// A description of the project.
	Description *string `pulumi:"description"`
	// The domain this project belongs to.
	DomainId *string `pulumi:"domainId"`
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`. Default is `true`.
	Enabled *bool `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`. Default is `false`. Changing this creates a new
	// project/domain.
	IsDomain *bool `pulumi:"isDomain"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The parent of this project. Changing this creates
	// a new project.
	ParentId *string `pulumi:"parentId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new project.
	Region *string `pulumi:"region"`
	// Tags for the project. Changing this updates the existing
	// project.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// A description of the project.
	Description pulumi.StringPtrInput
	// The domain this project belongs to.
	DomainId pulumi.StringPtrInput
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`. Default is `true`.
	Enabled pulumi.BoolPtrInput
	// Whether this project is a domain. Valid values
	// are `true` and `false`. Default is `false`. Changing this creates a new
	// project/domain.
	IsDomain pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The parent of this project. Changing this creates
	// a new project.
	ParentId pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new project.
	Region pulumi.StringPtrInput
	// Tags for the project. Changing this updates the existing
	// project.
	Tags pulumi.StringArrayInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// A description of the project.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain this project belongs to.
func (o ProjectOutput) DomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.DomainId }).(pulumi.StringOutput)
}

// Whether the project is enabled or disabled. Valid
// values are `true` and `false`. Default is `true`.
func (o ProjectOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Whether this project is a domain. Valid values
// are `true` and `false`. Default is `false`. Changing this creates a new
// project/domain.
func (o ProjectOutput) IsDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.IsDomain }).(pulumi.BoolPtrOutput)
}

// The name of the project.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent of this project. Changing this creates
// a new project.
func (o ProjectOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

// The region in which to obtain the V3 Keystone client.
// If omitted, the `region` argument of the provider is used. Changing this
// creates a new project.
func (o ProjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Tags for the project. Changing this updates the existing
// project.
func (o ProjectOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Project) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
