// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V1 container resource within OpenStack.
//
// ## Example Usage
// ### Basic Container
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/objectstorage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := objectstorage.NewContainer(ctx, "container1", &objectstorage.ContainerArgs{
// 			ContentType: pulumi.String("application/json"),
// 			Metadata: pulumi.StringMap{
// 				"test": pulumi.String("true"),
// 			},
// 			Region: pulumi.String("RegionOne"),
// 			Versioning: &objectstorage.ContainerVersioningArgs{
// 				Location: pulumi.String("tf-test-container-versions"),
// 				Type:     pulumi.String("versions"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Global Read Access
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/objectstorage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := objectstorage.NewContainer(ctx, "container1", &objectstorage.ContainerArgs{
// 			ContainerRead: pulumi.String(".r:*"),
// 			Region:        pulumi.String("RegionOne"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Global Read and List Access
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/objectstorage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := objectstorage.NewContainer(ctx, "container1", &objectstorage.ContainerArgs{
// 			ContainerRead: pulumi.String(".r:*,.rlistings"),
// 			Region:        pulumi.String("RegionOne"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Write-Only Access for a User
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/identity"
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/objectstorage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := identity.GetAuthScope(ctx, &identity.GetAuthScopeArgs{
// 			Name: "current",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = objectstorage.NewContainer(ctx, "container1", &objectstorage.ContainerArgs{
// 			ContainerRead:  pulumi.String(fmt.Sprintf("%v%v", ".r:-", _var.Username)),
// 			ContainerWrite: pulumi.String(fmt.Sprintf("%v%v%v", current.ProjectId, ":", _var.Username)),
// 			Region:         pulumi.String("RegionOne"),
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
// This resource can be imported by specifying the name of the containerSome attributes can't be imported * `force_destroy` * `content_type` * `metadata` * `container_sync_to` * `container_sync_key` So you'll have to `terraform plan` and `terraform apply` after the import to fix those missing attributes.
//
// ```sh
//  $ pulumi import openstack:objectstorage/container:Container container_1 <name>
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
	// Enable object versioning. The structure is described below.
	Versioning ContainerVersioningPtrOutput `pulumi:"versioning"`
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOption) (*Container, error) {
	if args == nil {
		args = &ContainerArgs{}
	}

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
	// Enable object versioning. The structure is described below.
	Versioning *ContainerVersioning `pulumi:"versioning"`
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
	// Enable object versioning. The structure is described below.
	Versioning ContainerVersioningPtrInput
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
	// Enable object versioning. The structure is described below.
	Versioning *ContainerVersioning `pulumi:"versioning"`
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
	// Enable object versioning. The structure is described below.
	Versioning ContainerVersioningPtrInput
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerArgs)(nil)).Elem()
}

type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(ctx context.Context) ContainerOutput
}

func (Container) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i Container) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i Container) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

type ContainerOutput struct {
	*pulumi.OutputState
}

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerOutput)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContainerOutput{})
}
