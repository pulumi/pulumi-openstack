// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

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
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/identity"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := identity.NewProject(ctx, "project1", &identity.ProjectArgs{
// 			Description: pulumi.String("A project"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Projects can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:identity/project:Project project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
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
	return reflect.TypeOf((*Project)(nil))
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

func (i *Project) ToProjectPtrOutput() ProjectPtrOutput {
	return i.ToProjectPtrOutputWithContext(context.Background())
}

func (i *Project) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPtrOutput)
}

type ProjectPtrInput interface {
	pulumi.Input

	ToProjectPtrOutput() ProjectPtrOutput
	ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput
}

type projectPtrType ProjectArgs

func (*projectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil))
}

func (i *projectPtrType) ToProjectPtrOutput() ProjectPtrOutput {
	return i.ToProjectPtrOutputWithContext(context.Background())
}

func (i *projectPtrType) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPtrOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//          ProjectArray{ ProjectArgs{...} }
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
//          ProjectMap{ "key": ProjectArgs{...} }
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
	return reflect.TypeOf((*Project)(nil))
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o.ToProjectPtrOutputWithContext(context.Background())
}

func (o ProjectOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Project) *Project {
		return &v
	}).(ProjectPtrOutput)
}

type ProjectPtrOutput struct{ *pulumi.OutputState }

func (ProjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil))
}

func (o ProjectPtrOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) Elem() ProjectOutput {
	return o.ApplyT(func(v *Project) Project {
		if v != nil {
			return *v
		}
		var ret Project
		return ret
	}).(ProjectOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Project)(nil))
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Project {
		return vs[0].([]Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Project)(nil))
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Project {
		return vs[0].(map[string]Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPtrInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectPtrOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
