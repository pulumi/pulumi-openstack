// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V3 Project resource within OpenStack Keystone.
//
// Note: You _must_ have admin privileges in your OpenStack cloud to use
// this resource.
type Project struct {
	pulumi.CustomResourceState

	// A description of the project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain this project belongs to.
	DomainId pulumi.StringOutput `pulumi:"domainId"`
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain pulumi.BoolPtrOutput `pulumi:"isDomain"`
	// The name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of this project.
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region pulumi.StringOutput `pulumi:"region"`
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
	// values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain *bool `pulumi:"isDomain"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The parent of this project.
	ParentId *string `pulumi:"parentId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region *string `pulumi:"region"`
}

type ProjectState struct {
	// A description of the project.
	Description pulumi.StringPtrInput
	// The domain this project belongs to.
	DomainId pulumi.StringPtrInput
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrInput
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The parent of this project.
	ParentId pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region pulumi.StringPtrInput
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
	// values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain *bool `pulumi:"isDomain"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The parent of this project.
	ParentId *string `pulumi:"parentId"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// A description of the project.
	Description pulumi.StringPtrInput
	// The domain this project belongs to.
	DomainId pulumi.StringPtrInput
	// Whether the project is enabled or disabled. Valid
	// values are `true` and `false`.
	Enabled pulumi.BoolPtrInput
	// Whether this project is a domain. Valid values
	// are `true` and `false`.
	IsDomain pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The parent of this project.
	ParentId pulumi.StringPtrInput
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used. Changing this
	// creates a new User.
	Region pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}
