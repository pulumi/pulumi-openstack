// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages members for the shared OpenStack Glance V2 Image within the source
// project, which owns the Image.
//
// ## Example Usage
//
// ### Unprivileged user
//
// Create a shared image and propose a membership to the
// `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/images"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rancheros, err := images.NewImage(ctx, "rancheros", &images.ImageArgs{
//				Name:            pulumi.String("RancherOS"),
//				ImageSourceUrl:  pulumi.String("https://releases.rancher.com/os/latest/rancheros-openstack.img"),
//				ContainerFormat: pulumi.String("bare"),
//				DiskFormat:      pulumi.String("qcow2"),
//				Visibility:      pulumi.String("shared"),
//				Properties: pulumi.StringMap{
//					"key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = images.NewImageAccess(ctx, "rancheros_member", &images.ImageAccessArgs{
//				ImageId:  rancheros.ID(),
//				MemberId: pulumi.String("bed6b6cbb86a4e2d8dc2735c2f1000e4"),
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
// ### Privileged user
//
// Create a shared image and set a membership to the
// `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/images"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			rancheros, err := images.NewImage(ctx, "rancheros", &images.ImageArgs{
//				Name:            pulumi.String("RancherOS"),
//				ImageSourceUrl:  pulumi.String("https://releases.rancher.com/os/latest/rancheros-openstack.img"),
//				ContainerFormat: pulumi.String("bare"),
//				DiskFormat:      pulumi.String("qcow2"),
//				Visibility:      pulumi.String("shared"),
//				Properties: pulumi.StringMap{
//					"key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = images.NewImageAccess(ctx, "rancheros_member", &images.ImageAccessArgs{
//				ImageId:  rancheros.ID(),
//				MemberId: pulumi.String("bed6b6cbb86a4e2d8dc2735c2f1000e4"),
//				Status:   pulumi.String("accepted"),
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
// Image access can be imported using the `image_id` and the `member_id`,
//
// separated by a slash, e.g.
//
// ```sh
// $ pulumi import openstack:images/imageAccess:ImageAccess openstack_images_image_access_v2 89c60255-9bd6-460c-822a-e2b959ede9d2/bed6b6cbb86a4e2d8dc2735c2f1000e4
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.MemberId == nil {
		return nil, errors.New("invalid value for required argument 'MemberId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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

type ImageAccessInput interface {
	pulumi.Input

	ToImageAccessOutput() ImageAccessOutput
	ToImageAccessOutputWithContext(ctx context.Context) ImageAccessOutput
}

func (*ImageAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageAccess)(nil)).Elem()
}

func (i *ImageAccess) ToImageAccessOutput() ImageAccessOutput {
	return i.ToImageAccessOutputWithContext(context.Background())
}

func (i *ImageAccess) ToImageAccessOutputWithContext(ctx context.Context) ImageAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAccessOutput)
}

// ImageAccessArrayInput is an input type that accepts ImageAccessArray and ImageAccessArrayOutput values.
// You can construct a concrete instance of `ImageAccessArrayInput` via:
//
//	ImageAccessArray{ ImageAccessArgs{...} }
type ImageAccessArrayInput interface {
	pulumi.Input

	ToImageAccessArrayOutput() ImageAccessArrayOutput
	ToImageAccessArrayOutputWithContext(context.Context) ImageAccessArrayOutput
}

type ImageAccessArray []ImageAccessInput

func (ImageAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageAccess)(nil)).Elem()
}

func (i ImageAccessArray) ToImageAccessArrayOutput() ImageAccessArrayOutput {
	return i.ToImageAccessArrayOutputWithContext(context.Background())
}

func (i ImageAccessArray) ToImageAccessArrayOutputWithContext(ctx context.Context) ImageAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAccessArrayOutput)
}

// ImageAccessMapInput is an input type that accepts ImageAccessMap and ImageAccessMapOutput values.
// You can construct a concrete instance of `ImageAccessMapInput` via:
//
//	ImageAccessMap{ "key": ImageAccessArgs{...} }
type ImageAccessMapInput interface {
	pulumi.Input

	ToImageAccessMapOutput() ImageAccessMapOutput
	ToImageAccessMapOutputWithContext(context.Context) ImageAccessMapOutput
}

type ImageAccessMap map[string]ImageAccessInput

func (ImageAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageAccess)(nil)).Elem()
}

func (i ImageAccessMap) ToImageAccessMapOutput() ImageAccessMapOutput {
	return i.ToImageAccessMapOutputWithContext(context.Background())
}

func (i ImageAccessMap) ToImageAccessMapOutputWithContext(ctx context.Context) ImageAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageAccessMapOutput)
}

type ImageAccessOutput struct{ *pulumi.OutputState }

func (ImageAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageAccess)(nil)).Elem()
}

func (o ImageAccessOutput) ToImageAccessOutput() ImageAccessOutput {
	return o
}

func (o ImageAccessOutput) ToImageAccessOutputWithContext(ctx context.Context) ImageAccessOutput {
	return o
}

// The date the image access was created.
func (o ImageAccessOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The image ID.
func (o ImageAccessOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The member ID, e.g. the target project ID.
func (o ImageAccessOutput) MemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.MemberId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Glance client.
// A Glance client is needed to manage Image members. If omitted, the `region`
// argument of the provider is used. Changing this creates a new resource.
func (o ImageAccessOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The member schema.
func (o ImageAccessOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// The member proposal status. Optional if admin wants to
// force the member proposal acceptance. Can either be `accepted`, `rejected` or
// `pending`. Defaults to `pending`. Foridden for non-admin users.
func (o ImageAccessOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date the image access was last updated.
func (o ImageAccessOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageAccess) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ImageAccessArrayOutput struct{ *pulumi.OutputState }

func (ImageAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageAccess)(nil)).Elem()
}

func (o ImageAccessArrayOutput) ToImageAccessArrayOutput() ImageAccessArrayOutput {
	return o
}

func (o ImageAccessArrayOutput) ToImageAccessArrayOutputWithContext(ctx context.Context) ImageAccessArrayOutput {
	return o
}

func (o ImageAccessArrayOutput) Index(i pulumi.IntInput) ImageAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageAccess {
		return vs[0].([]*ImageAccess)[vs[1].(int)]
	}).(ImageAccessOutput)
}

type ImageAccessMapOutput struct{ *pulumi.OutputState }

func (ImageAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageAccess)(nil)).Elem()
}

func (o ImageAccessMapOutput) ToImageAccessMapOutput() ImageAccessMapOutput {
	return o
}

func (o ImageAccessMapOutput) ToImageAccessMapOutputWithContext(ctx context.Context) ImageAccessMapOutput {
	return o
}

func (o ImageAccessMapOutput) MapIndex(k pulumi.StringInput) ImageAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageAccess {
		return vs[0].(map[string]*ImageAccess)[vs[1].(string)]
	}).(ImageAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAccessInput)(nil)).Elem(), &ImageAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAccessArrayInput)(nil)).Elem(), ImageAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageAccessMapInput)(nil)).Elem(), ImageAccessMap{})
	pulumi.RegisterOutputType(ImageAccessOutput{})
	pulumi.RegisterOutputType(ImageAccessArrayOutput{})
	pulumi.RegisterOutputType(ImageAccessMapOutput{})
}
