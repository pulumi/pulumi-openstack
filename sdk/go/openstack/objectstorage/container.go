// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 container resource within OpenStack.
//
// ## Example Usage
//
// ### Basic Container
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region: pulumi.String("RegionOne"),
//				Name:   pulumi.String("tf-test-container-1"),
//				Metadata: pulumi.Map{
//					"test": pulumi.Any("true"),
//				},
//				ContentType: pulumi.String("application/json"),
//				Versioning:  pulumi.Bool(true),
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
// ### Basic Container with legacy versioning
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region: pulumi.String("RegionOne"),
//				Name:   pulumi.String("tf-test-container-1"),
//				Metadata: pulumi.Map{
//					"test": pulumi.Any("true"),
//				},
//				ContentType: pulumi.String("application/json"),
//				VersioningLegacy: &objectstorage.ContainerVersioningLegacyArgs{
//					Type:     pulumi.String("versions"),
//					Location: pulumi.String("tf-test-container-versions"),
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
// ### Global Read Access
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Requires that a user know the object name they are attempting to download
//			_, err := objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region:        pulumi.String("RegionOne"),
//				Name:          pulumi.String("tf-test-container-1"),
//				ContainerRead: pulumi.String(".r:*"),
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
// ### Global Read and List Access
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Any user can read any object, and list all objects in the container
//			_, err := objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region:        pulumi.String("RegionOne"),
//				Name:          pulumi.String("tf-test-container-1"),
//				ContainerRead: pulumi.String(".r:*,.rlistings"),
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
// ### Write-Only Access for a User
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/identity"
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := identity.GetAuthScope(ctx, &identity.GetAuthScopeArgs{
//				Name: "current",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			// The named user can only upload objects, not read objects or list the container
//			_, err = objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region:         pulumi.String("RegionOne"),
//				Name:           pulumi.String("tf-test-container-1"),
//				ContainerRead:  pulumi.Sprintf(".r:-%v", username),
//				ContainerWrite: pulumi.Sprintf("%v:%v", current.ProjectId, username),
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
// This resource can be imported by specifying the name of the container:
//
// Some attributes can't be imported :
// * `force_destroy`
// * `content_type`
// * `metadata`
// * `container_sync_to`
// * `container_sync_key`
//
// So you'll have to `pulumi preview` and `pulumi up` after the import to fix those missing attributes.
//
// ```sh
// $ pulumi import openstack:objectstorage/container:Container container_1 container_name
// ```
type Container struct {
	pulumi.CustomResourceState

	// Sets an access control list (ACL) that grants
	// read access. This header can contain a comma-delimited list of users that
	// can read the container (allows the GET method for all objects in the
	// container). Changing this updates the access control list read access.
	ContainerRead pulumi.StringPtrOutput `pulumi:"containerRead"`
	// The secret key for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncKey pulumi.StringPtrOutput `pulumi:"containerSyncKey"`
	// The destination for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncTo pulumi.StringPtrOutput `pulumi:"containerSyncTo"`
	// Sets an ACL that grants write access.
	// Changing this updates the access control list write access.
	ContainerWrite pulumi.StringPtrOutput `pulumi:"containerWrite"`
	// The MIME type for the container. Changing this
	// updates the MIME type.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Custom key/value pairs to associate with the container.
	// Changing this updates the existing container metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// A unique name for the container. Changing this creates a
	// new container.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region pulumi.StringOutput `pulumi:"region"`
	// The storage policy to be used for the container.
	// Changing this creates a new container.
	StoragePolicy pulumi.StringOutput `pulumi:"storagePolicy"`
	// A boolean that can enable or disable object
	// versioning. The default value is `false`. To use this feature, your Swift
	// version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
	// and a cloud administrator must have set the `allowObjectVersioning = true`
	// configuration option in Swift. If you cannot set this versioning type, you may
	// want to consider using `versioningLegacy` instead.
	Versioning pulumi.BoolPtrOutput `pulumi:"versioning"`
	// Enable legacy object versioning. The structure is described below.
	//
	// Deprecated: Use newer "versioning" implementation
	VersioningLegacy ContainerVersioningLegacyPtrOutput `pulumi:"versioningLegacy"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		args = &ContainerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Container
	err := ctx.RegisterResource("openstack:objectstorage/container:Container", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerState, opts ...pulumi.ResourceOption) (*Container, error) {
	var resource Container
	err := ctx.ReadResource("openstack:objectstorage/container:Container", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Container resources.
type containerState struct {
	// Sets an access control list (ACL) that grants
	// read access. This header can contain a comma-delimited list of users that
	// can read the container (allows the GET method for all objects in the
	// container). Changing this updates the access control list read access.
	ContainerRead *string `pulumi:"containerRead"`
	// The secret key for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncKey *string `pulumi:"containerSyncKey"`
	// The destination for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncTo *string `pulumi:"containerSyncTo"`
	// Sets an ACL that grants write access.
	// Changing this updates the access control list write access.
	ContainerWrite *string `pulumi:"containerWrite"`
	// The MIME type for the container. Changing this
	// updates the MIME type.
	ContentType *string `pulumi:"contentType"`
	// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Custom key/value pairs to associate with the container.
	// Changing this updates the existing container metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// A unique name for the container. Changing this creates a
	// new container.
	Name *string `pulumi:"name"`
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region *string `pulumi:"region"`
	// The storage policy to be used for the container.
	// Changing this creates a new container.
	StoragePolicy *string `pulumi:"storagePolicy"`
	// A boolean that can enable or disable object
	// versioning. The default value is `false`. To use this feature, your Swift
	// version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
	// and a cloud administrator must have set the `allowObjectVersioning = true`
	// configuration option in Swift. If you cannot set this versioning type, you may
	// want to consider using `versioningLegacy` instead.
	Versioning *bool `pulumi:"versioning"`
	// Enable legacy object versioning. The structure is described below.
	//
	// Deprecated: Use newer "versioning" implementation
	VersioningLegacy *ContainerVersioningLegacy `pulumi:"versioningLegacy"`
}

type ContainerState struct {
	// Sets an access control list (ACL) that grants
	// read access. This header can contain a comma-delimited list of users that
	// can read the container (allows the GET method for all objects in the
	// container). Changing this updates the access control list read access.
	ContainerRead pulumi.StringPtrInput
	// The secret key for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncKey pulumi.StringPtrInput
	// The destination for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncTo pulumi.StringPtrInput
	// Sets an ACL that grants write access.
	// Changing this updates the access control list write access.
	ContainerWrite pulumi.StringPtrInput
	// The MIME type for the container. Changing this
	// updates the MIME type.
	ContentType pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// Custom key/value pairs to associate with the container.
	// Changing this updates the existing container metadata.
	Metadata pulumi.MapInput
	// A unique name for the container. Changing this creates a
	// new container.
	Name pulumi.StringPtrInput
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region pulumi.StringPtrInput
	// The storage policy to be used for the container.
	// Changing this creates a new container.
	StoragePolicy pulumi.StringPtrInput
	// A boolean that can enable or disable object
	// versioning. The default value is `false`. To use this feature, your Swift
	// version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
	// and a cloud administrator must have set the `allowObjectVersioning = true`
	// configuration option in Swift. If you cannot set this versioning type, you may
	// want to consider using `versioningLegacy` instead.
	Versioning pulumi.BoolPtrInput
	// Enable legacy object versioning. The structure is described below.
	//
	// Deprecated: Use newer "versioning" implementation
	VersioningLegacy ContainerVersioningLegacyPtrInput
}

func (ContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerState)(nil)).Elem()
}

type containerArgs struct {
	// Sets an access control list (ACL) that grants
	// read access. This header can contain a comma-delimited list of users that
	// can read the container (allows the GET method for all objects in the
	// container). Changing this updates the access control list read access.
	ContainerRead *string `pulumi:"containerRead"`
	// The secret key for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncKey *string `pulumi:"containerSyncKey"`
	// The destination for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncTo *string `pulumi:"containerSyncTo"`
	// Sets an ACL that grants write access.
	// Changing this updates the access control list write access.
	ContainerWrite *string `pulumi:"containerWrite"`
	// The MIME type for the container. Changing this
	// updates the MIME type.
	ContentType *string `pulumi:"contentType"`
	// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Custom key/value pairs to associate with the container.
	// Changing this updates the existing container metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// A unique name for the container. Changing this creates a
	// new container.
	Name *string `pulumi:"name"`
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region *string `pulumi:"region"`
	// The storage policy to be used for the container.
	// Changing this creates a new container.
	StoragePolicy *string `pulumi:"storagePolicy"`
	// A boolean that can enable or disable object
	// versioning. The default value is `false`. To use this feature, your Swift
	// version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
	// and a cloud administrator must have set the `allowObjectVersioning = true`
	// configuration option in Swift. If you cannot set this versioning type, you may
	// want to consider using `versioningLegacy` instead.
	Versioning *bool `pulumi:"versioning"`
	// Enable legacy object versioning. The structure is described below.
	//
	// Deprecated: Use newer "versioning" implementation
	VersioningLegacy *ContainerVersioningLegacy `pulumi:"versioningLegacy"`
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// Sets an access control list (ACL) that grants
	// read access. This header can contain a comma-delimited list of users that
	// can read the container (allows the GET method for all objects in the
	// container). Changing this updates the access control list read access.
	ContainerRead pulumi.StringPtrInput
	// The secret key for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncKey pulumi.StringPtrInput
	// The destination for container synchronization.
	// Changing this updates container synchronization.
	ContainerSyncTo pulumi.StringPtrInput
	// Sets an ACL that grants write access.
	// Changing this updates the access control list write access.
	ContainerWrite pulumi.StringPtrInput
	// The MIME type for the container. Changing this
	// updates the MIME type.
	ContentType pulumi.StringPtrInput
	// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// Custom key/value pairs to associate with the container.
	// Changing this updates the existing container metadata.
	Metadata pulumi.MapInput
	// A unique name for the container. Changing this creates a
	// new container.
	Name pulumi.StringPtrInput
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region pulumi.StringPtrInput
	// The storage policy to be used for the container.
	// Changing this creates a new container.
	StoragePolicy pulumi.StringPtrInput
	// A boolean that can enable or disable object
	// versioning. The default value is `false`. To use this feature, your Swift
	// version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
	// and a cloud administrator must have set the `allowObjectVersioning = true`
	// configuration option in Swift. If you cannot set this versioning type, you may
	// want to consider using `versioningLegacy` instead.
	Versioning pulumi.BoolPtrInput
	// Enable legacy object versioning. The structure is described below.
	//
	// Deprecated: Use newer "versioning" implementation
	VersioningLegacy ContainerVersioningLegacyPtrInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (*Container) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i *Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

// ContainerArrayInput is an input type that accepts ContainerArray and ContainerArrayOutput values.
// You can construct a concrete instance of `ContainerArrayInput` via:
//
//	ContainerArray{ ContainerArgs{...} }
type ContainerArrayInput interface {
	pulumi.Input

	ToContainerArrayOutput() ContainerArrayOutput
	ToContainerArrayOutputWithContext(context.Context) ContainerArrayOutput
}

type ContainerArray []ContainerInput

func (ContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Container)(nil)).Elem()
}

func (i ContainerArray) ToContainerArrayOutput() ContainerArrayOutput {
	return i.ToContainerArrayOutputWithContext(context.Background())
}

func (i ContainerArray) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerArrayOutput)
}

// ContainerMapInput is an input type that accepts ContainerMap and ContainerMapOutput values.
// You can construct a concrete instance of `ContainerMapInput` via:
//
//	ContainerMap{ "key": ContainerArgs{...} }
type ContainerMapInput interface {
	pulumi.Input

	ToContainerMapOutput() ContainerMapOutput
	ToContainerMapOutputWithContext(context.Context) ContainerMapOutput
}

type ContainerMap map[string]ContainerInput

func (ContainerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Container)(nil)).Elem()
}

func (i ContainerMap) ToContainerMapOutput() ContainerMapOutput {
	return i.ToContainerMapOutputWithContext(context.Background())
}

func (i ContainerMap) ToContainerMapOutputWithContext(ctx context.Context) ContainerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerMapOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

// Sets an access control list (ACL) that grants
// read access. This header can contain a comma-delimited list of users that
// can read the container (allows the GET method for all objects in the
// container). Changing this updates the access control list read access.
func (o ContainerOutput) ContainerRead() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.ContainerRead }).(pulumi.StringPtrOutput)
}

// The secret key for container synchronization.
// Changing this updates container synchronization.
func (o ContainerOutput) ContainerSyncKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.ContainerSyncKey }).(pulumi.StringPtrOutput)
}

// The destination for container synchronization.
// Changing this updates container synchronization.
func (o ContainerOutput) ContainerSyncTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.ContainerSyncTo }).(pulumi.StringPtrOutput)
}

// Sets an ACL that grants write access.
// Changing this updates the access control list write access.
func (o ContainerOutput) ContainerWrite() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.ContainerWrite }).(pulumi.StringPtrOutput)
}

// The MIME type for the container. Changing this
// updates the MIME type.
func (o ContainerOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
func (o ContainerOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// Custom key/value pairs to associate with the container.
// Changing this updates the existing container metadata.
func (o ContainerOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *Container) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// A unique name for the container. Changing this creates a
// new container.
func (o ContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the container. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new container.
func (o ContainerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The storage policy to be used for the container.
// Changing this creates a new container.
func (o ContainerOutput) StoragePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Container) pulumi.StringOutput { return v.StoragePolicy }).(pulumi.StringOutput)
}

// A boolean that can enable or disable object
// versioning. The default value is `false`. To use this feature, your Swift
// version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
// and a cloud administrator must have set the `allowObjectVersioning = true`
// configuration option in Swift. If you cannot set this versioning type, you may
// want to consider using `versioningLegacy` instead.
func (o ContainerOutput) Versioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Container) pulumi.BoolPtrOutput { return v.Versioning }).(pulumi.BoolPtrOutput)
}

// Enable legacy object versioning. The structure is described below.
//
// Deprecated: Use newer "versioning" implementation
func (o ContainerOutput) VersioningLegacy() ContainerVersioningLegacyPtrOutput {
	return o.ApplyT(func(v *Container) ContainerVersioningLegacyPtrOutput { return v.VersioningLegacy }).(ContainerVersioningLegacyPtrOutput)
}

type ContainerArrayOutput struct{ *pulumi.OutputState }

func (ContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Container)(nil)).Elem()
}

func (o ContainerArrayOutput) ToContainerArrayOutput() ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) Index(i pulumi.IntInput) ContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Container {
		return vs[0].([]*Container)[vs[1].(int)]
	}).(ContainerOutput)
}

type ContainerMapOutput struct{ *pulumi.OutputState }

func (ContainerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Container)(nil)).Elem()
}

func (o ContainerMapOutput) ToContainerMapOutput() ContainerMapOutput {
	return o
}

func (o ContainerMapOutput) ToContainerMapOutputWithContext(ctx context.Context) ContainerMapOutput {
	return o
}

func (o ContainerMapOutput) MapIndex(k pulumi.StringInput) ContainerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Container {
		return vs[0].(map[string]*Container)[vs[1].(string)]
	}).(ContainerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInput)(nil)).Elem(), &Container{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerArrayInput)(nil)).Elem(), ContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerMapInput)(nil)).Elem(), ContainerMap{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerMapOutput{})
}
