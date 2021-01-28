// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages memberships status for the shared OpenStack Glance V2 Image within the
// destination project, which has a member proposal.
//
// ## Example Usage
//
// Accept a shared image membershipship proposal within the current project.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v2/go/openstack/images"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "all"
// 		opt1 := "RancherOS"
// 		opt2 := "shared"
// 		rancheros, err := images.LookupImage(ctx, &images.LookupImageArgs{
// 			MemberStatus: &opt0,
// 			Name:         &opt1,
// 			Visibility:   &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = images.NewImageAccessAccept(ctx, "rancherosMember", &images.ImageAccessAcceptArgs{
// 			ImageId: pulumi.String(rancheros.Id),
// 			Status:  pulumi.String("accepted"),
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
// Image access acceptance status can be imported using the `image_id`, e.g.
//
// ```sh
//  $ pulumi import openstack:images/imageAccessAccept:ImageAccessAccept openstack_images_image_access_accept_v2 89c60255-9bd6-460c-822a-e2b959ede9d2
// ```
type ImageAccessAccept struct {
	pulumi.CustomResourceState

	// The date the image membership was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The proposed image ID.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The member ID, e.g. the target project ID. Optional
	// for admin accounts. Defaults to the current scope project ID.
	MemberId pulumi.StringOutput `pulumi:"memberId"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image memberships. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// membership.
	Region pulumi.StringOutput `pulumi:"region"`
	// The membership schema.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// The membership proposal status. Can either be
	// `accepted`, `rejected` or `pending`.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date the image membership was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewImageAccessAccept registers a new resource with the given unique name, arguments, and options.
func NewImageAccessAccept(ctx *pulumi.Context,
	name string, args *ImageAccessAcceptArgs, opts ...pulumi.ResourceOption) (*ImageAccessAccept, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	var resource ImageAccessAccept
	err := ctx.RegisterResource("openstack:images/imageAccessAccept:ImageAccessAccept", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageAccessAccept gets an existing ImageAccessAccept resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageAccessAccept(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageAccessAcceptState, opts ...pulumi.ResourceOption) (*ImageAccessAccept, error) {
	var resource ImageAccessAccept
	err := ctx.ReadResource("openstack:images/imageAccessAccept:ImageAccessAccept", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageAccessAccept resources.
type imageAccessAcceptState struct {
	// The date the image membership was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The proposed image ID.
	ImageId *string `pulumi:"imageId"`
	// The member ID, e.g. the target project ID. Optional
	// for admin accounts. Defaults to the current scope project ID.
	MemberId *string `pulumi:"memberId"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image memberships. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// membership.
	Region *string `pulumi:"region"`
	// The membership schema.
	Schema *string `pulumi:"schema"`
	// The membership proposal status. Can either be
	// `accepted`, `rejected` or `pending`.
	Status *string `pulumi:"status"`
	// The date the image membership was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ImageAccessAcceptState struct {
	// The date the image membership was created.
	CreatedAt pulumi.StringPtrInput
	// The proposed image ID.
	ImageId pulumi.StringPtrInput
	// The member ID, e.g. the target project ID. Optional
	// for admin accounts. Defaults to the current scope project ID.
	MemberId pulumi.StringPtrInput
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image memberships. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// membership.
	Region pulumi.StringPtrInput
	// The membership schema.
	Schema pulumi.StringPtrInput
	// The membership proposal status. Can either be
	// `accepted`, `rejected` or `pending`.
	Status pulumi.StringPtrInput
	// The date the image membership was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (ImageAccessAcceptState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageAccessAcceptState)(nil)).Elem()
}

type imageAccessAcceptArgs struct {
	// The proposed image ID.
	ImageId string `pulumi:"imageId"`
	// The member ID, e.g. the target project ID. Optional
	// for admin accounts. Defaults to the current scope project ID.
	MemberId *string `pulumi:"memberId"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image memberships. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// membership.
	Region *string `pulumi:"region"`
	// The membership proposal status. Can either be
	// `accepted`, `rejected` or `pending`.
	Status string `pulumi:"status"`
}

// The set of arguments for constructing a ImageAccessAccept resource.
type ImageAccessAcceptArgs struct {
	// The proposed image ID.
	ImageId pulumi.StringInput
	// The member ID, e.g. the target project ID. Optional
	// for admin accounts. Defaults to the current scope project ID.
	MemberId pulumi.StringPtrInput
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image memberships. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// membership.
	Region pulumi.StringPtrInput
	// The membership proposal status. Can either be
	// `accepted`, `rejected` or `pending`.
	Status pulumi.StringInput
}

func (ImageAccessAcceptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageAccessAcceptArgs)(nil)).Elem()
}

type ImageAccessAcceptInput interface {
	pulumi.Input

	ToImageAccessAcceptOutput() ImageAccessAcceptOutput
	ToImageAccessAcceptOutputWithContext(ctx context.Context) ImageAccessAcceptOutput
}

func (*ImageAccessAccept) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAccessAccept)(nil))
}

func (i *ImageAccessAccept) ToImageAccessAcceptOutput() ImageAccessAcceptOutput {
	return i.ToImageAccessAcceptOutputWithContext(context.Background())
}

func (i *ImageAccessAccept) ToImageAccessAcceptOutputWithContext(ctx context.Context) ImageAccessAcceptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAccessAcceptOutput)
}

type ImageAccessAcceptOutput struct {
	*pulumi.OutputState
}

func (ImageAccessAcceptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageAccessAccept)(nil))
}

func (o ImageAccessAcceptOutput) ToImageAccessAcceptOutput() ImageAccessAcceptOutput {
	return o
}

func (o ImageAccessAcceptOutput) ToImageAccessAcceptOutputWithContext(ctx context.Context) ImageAccessAcceptOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageAccessAcceptOutput{})
}
