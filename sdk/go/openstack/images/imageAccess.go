// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages members for the shared OpenStack Glance V2 Image within the source
// project, which owns the Image.
type ImageAccess struct {
	pulumi.CustomResourceState

	// The date the image access was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The image ID.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The member ID, e.g. the target project ID.
	MemberId pulumi.StringOutput `pulumi:"memberId"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The member schema.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be `accepted`, `rejected` or
	// `pending`. Defaults to `pending`. Foridden for non-admin users.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date the image access was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewImageAccess registers a new resource with the given unique name, arguments, and options.
func NewImageAccess(ctx *pulumi.Context,
	name string, args *ImageAccessArgs, opts ...pulumi.ResourceOption) (*ImageAccess, error) {
	if args == nil || args.ImageId == nil {
		return nil, errors.New("missing required argument 'ImageId'")
	}
	if args == nil || args.MemberId == nil {
		return nil, errors.New("missing required argument 'MemberId'")
	}
	if args == nil {
		args = &ImageAccessArgs{}
	}
	var resource ImageAccess
	err := ctx.RegisterResource("openstack:images/imageAccess:ImageAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageAccess gets an existing ImageAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageAccessState, opts ...pulumi.ResourceOption) (*ImageAccess, error) {
	var resource ImageAccess
	err := ctx.ReadResource("openstack:images/imageAccess:ImageAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageAccess resources.
type imageAccessState struct {
	// The date the image access was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The image ID.
	ImageId *string `pulumi:"imageId"`
	// The member ID, e.g. the target project ID.
	MemberId *string `pulumi:"memberId"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// The member schema.
	Schema *string `pulumi:"schema"`
	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be `accepted`, `rejected` or
	// `pending`. Defaults to `pending`. Foridden for non-admin users.
	Status *string `pulumi:"status"`
	// The date the image access was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ImageAccessState struct {
	// The date the image access was created.
	CreatedAt pulumi.StringPtrInput
	// The image ID.
	ImageId pulumi.StringPtrInput
	// The member ID, e.g. the target project ID.
	MemberId pulumi.StringPtrInput
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// The member schema.
	Schema pulumi.StringPtrInput
	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be `accepted`, `rejected` or
	// `pending`. Defaults to `pending`. Foridden for non-admin users.
	Status pulumi.StringPtrInput
	// The date the image access was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (ImageAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageAccessState)(nil)).Elem()
}

type imageAccessArgs struct {
	// The image ID.
	ImageId string `pulumi:"imageId"`
	// The member ID, e.g. the target project ID.
	MemberId string `pulumi:"memberId"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be `accepted`, `rejected` or
	// `pending`. Defaults to `pending`. Foridden for non-admin users.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a ImageAccess resource.
type ImageAccessArgs struct {
	// The image ID.
	ImageId pulumi.StringInput
	// The member ID, e.g. the target project ID.
	MemberId pulumi.StringInput
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to manage Image members. If omitted, the `region`
	// argument of the provider is used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
	// The member proposal status. Optional if admin wants to
	// force the member proposal acceptance. Can either be `accepted`, `rejected` or
	// `pending`. Defaults to `pending`. Foridden for non-admin users.
	Status pulumi.StringPtrInput
}

func (ImageAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageAccessArgs)(nil)).Elem()
}
