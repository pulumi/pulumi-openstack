// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V1 Barbican container resource within OpenStack.
//
// ## Example Usage
// ### Container with the ACL
//
// > **Note** Only read ACLs are supported
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/keymanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := keymanager.NewContainerV1(ctx, "tls1", &keymanager.ContainerV1Args{
// 			Acl: &keymanager.ContainerV1AclArgs{
// 				Read: &keymanager.ContainerV1AclReadArgs{
// 					ProjectAccess: pulumi.Bool(false),
// 					Users: pulumi.StringArray{
// 						pulumi.String("userid1"),
// 						pulumi.String("userid2"),
// 					},
// 				},
// 			},
// 			SecretRefs: keymanager.ContainerV1SecretRefArray{
// 				&keymanager.ContainerV1SecretRefArgs{
// 					Name:      pulumi.String("certificate"),
// 					SecretRef: pulumi.Any(openstack_keymanager_secret_v1.Certificate_1.Secret_ref),
// 				},
// 				&keymanager.ContainerV1SecretRefArgs{
// 					Name:      pulumi.String("private_key"),
// 					SecretRef: pulumi.Any(openstack_keymanager_secret_v1.Private_key_1.Secret_ref),
// 				},
// 				&keymanager.ContainerV1SecretRefArgs{
// 					Name:      pulumi.String("intermediates"),
// 					SecretRef: pulumi.Any(openstack_keymanager_secret_v1.Intermediate_1.Secret_ref),
// 				},
// 			},
// 			Type: pulumi.String("certificate"),
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
// Containers can be imported using the container id (the last part of the container reference), e.g.
//
// ```sh
//  $ pulumi import openstack:keymanager/containerV1:ContainerV1 container_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
// ```
type ContainerV1 struct {
	pulumi.CustomResourceState

	// Allows to control an access to a container. Currently only
	// the `read` operation is supported. If not specified, the container is
	// accessible project wide. The `read` structure is described below.
	Acl ContainerV1AclOutput `pulumi:"acl"`
	// The list of the container consumers. The structure is described below.
	Consumers ContainerV1ConsumerArrayOutput `pulumi:"consumers"`
	// The container reference / where to find the container.
	ContainerRef pulumi.StringOutput `pulumi:"containerRef"`
	// The date the container ACL was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The creator of the container.
	CreatorId pulumi.StringOutput `pulumi:"creatorId"`
	// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 container.
	Region pulumi.StringOutput `pulumi:"region"`
	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs ContainerV1SecretRefArrayOutput `pulumi:"secretRefs"`
	// The status of the container.
	Status pulumi.StringOutput `pulumi:"status"`
	// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The date the container ACL was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewContainerV1 registers a new resource with the given unique name, arguments, and options.
func NewContainerV1(ctx *pulumi.Context,
	name string, args *ContainerV1Args, opts ...pulumi.ResourceOption) (*ContainerV1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ContainerV1
	err := ctx.RegisterResource("openstack:keymanager/containerV1:ContainerV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerV1 gets an existing ContainerV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerV1State, opts ...pulumi.ResourceOption) (*ContainerV1, error) {
	var resource ContainerV1
	err := ctx.ReadResource("openstack:keymanager/containerV1:ContainerV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerV1 resources.
type containerV1State struct {
	// Allows to control an access to a container. Currently only
	// the `read` operation is supported. If not specified, the container is
	// accessible project wide. The `read` structure is described below.
	Acl *ContainerV1Acl `pulumi:"acl"`
	// The list of the container consumers. The structure is described below.
	Consumers []ContainerV1Consumer `pulumi:"consumers"`
	// The container reference / where to find the container.
	ContainerRef *string `pulumi:"containerRef"`
	// The date the container ACL was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The creator of the container.
	CreatorId *string `pulumi:"creatorId"`
	// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
	Name *string `pulumi:"name"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 container.
	Region *string `pulumi:"region"`
	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs []ContainerV1SecretRef `pulumi:"secretRefs"`
	// The status of the container.
	Status *string `pulumi:"status"`
	// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
	Type *string `pulumi:"type"`
	// The date the container ACL was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ContainerV1State struct {
	// Allows to control an access to a container. Currently only
	// the `read` operation is supported. If not specified, the container is
	// accessible project wide. The `read` structure is described below.
	Acl ContainerV1AclPtrInput
	// The list of the container consumers. The structure is described below.
	Consumers ContainerV1ConsumerArrayInput
	// The container reference / where to find the container.
	ContainerRef pulumi.StringPtrInput
	// The date the container ACL was created.
	CreatedAt pulumi.StringPtrInput
	// The creator of the container.
	CreatorId pulumi.StringPtrInput
	// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
	Name pulumi.StringPtrInput
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 container.
	Region pulumi.StringPtrInput
	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs ContainerV1SecretRefArrayInput
	// The status of the container.
	Status pulumi.StringPtrInput
	// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
	Type pulumi.StringPtrInput
	// The date the container ACL was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (ContainerV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*containerV1State)(nil)).Elem()
}

type containerV1Args struct {
	// Allows to control an access to a container. Currently only
	// the `read` operation is supported. If not specified, the container is
	// accessible project wide. The `read` structure is described below.
	Acl *ContainerV1Acl `pulumi:"acl"`
	// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
	Name *string `pulumi:"name"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 container.
	Region *string `pulumi:"region"`
	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs []ContainerV1SecretRef `pulumi:"secretRefs"`
	// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ContainerV1 resource.
type ContainerV1Args struct {
	// Allows to control an access to a container. Currently only
	// the `read` operation is supported. If not specified, the container is
	// accessible project wide. The `read` structure is described below.
	Acl ContainerV1AclPtrInput
	// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
	Name pulumi.StringPtrInput
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to create a container. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// V1 container.
	Region pulumi.StringPtrInput
	// A set of dictionaries containing references to secrets. The structure is described
	// below.
	SecretRefs ContainerV1SecretRefArrayInput
	// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
	Type pulumi.StringInput
}

func (ContainerV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*containerV1Args)(nil)).Elem()
}

type ContainerV1Input interface {
	pulumi.Input

	ToContainerV1Output() ContainerV1Output
	ToContainerV1OutputWithContext(ctx context.Context) ContainerV1Output
}

func (*ContainerV1) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerV1)(nil))
}

func (i *ContainerV1) ToContainerV1Output() ContainerV1Output {
	return i.ToContainerV1OutputWithContext(context.Background())
}

func (i *ContainerV1) ToContainerV1OutputWithContext(ctx context.Context) ContainerV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerV1Output)
}

type ContainerV1Output struct {
	*pulumi.OutputState
}

func (ContainerV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerV1)(nil))
}

func (o ContainerV1Output) ToContainerV1Output() ContainerV1Output {
	return o
}

func (o ContainerV1Output) ToContainerV1OutputWithContext(ctx context.Context) ContainerV1Output {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContainerV1Output{})
}
