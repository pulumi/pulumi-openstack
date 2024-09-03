// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack image.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/images"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := images.LookupImage(ctx, &images.LookupImageArgs{
//				Name:       pulumi.StringRef("Ubuntu 16.04"),
//				MostRecent: pulumi.BoolRef(true),
//				Properties: map[string]interface{}{
//					"key": "value",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupImageResult
	err := ctx.Invoke("openstack:images/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// The container format of the image.
	ContainerFormat *string `pulumi:"containerFormat"`
	// The disk format of the image.
	DiskFormat *string `pulumi:"diskFormat"`
	// Whether or not the image is hidden from public list.
	Hidden *bool `pulumi:"hidden"`
	// The status of the image. Must be one of
	// "accepted", "pending", "rejected", or "all".
	MemberStatus *string `pulumi:"memberStatus"`
	// If more than one result is returned, use the most
	// recent image.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the image. Cannot be used simultaneously with
	// `nameRegex`.
	Name *string `pulumi:"name"`
	// The regular expressian of the name of the image.
	// Cannot be used simultaneously with `name`. Unlike filtering by `name` the
	// `nameRegex` filtering does by client on the result of OpenStack search
	// query.
	NameRegex *string `pulumi:"nameRegex"`
	// The owner (UUID) of the image.
	Owner *string `pulumi:"owner"`
	// a map of key/value pairs to match an image with.
	// All specified properties must be matched. Unlike other options filtering by
	// `properties` does by client on the result of OpenStack search query.
	// Filtering is applied if server responce contains at least 2 images. In case
	// there is only one image the `properties` ignores.
	Properties map[string]string `pulumi:"properties"`
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The maximum size (in bytes) of the image to return.
	SizeMax *int `pulumi:"sizeMax"`
	// The minimum size (in bytes) of the image to return.
	SizeMin *int `pulumi:"sizeMin"`
	// Sorts the response by one or more attribute and sort
	// direction combinations. You can also set multiple sort keys and directions.
	// Default direction is `desc`. Use the comma (,) character to separate multiple
	// values. For example expression `sort = "name:asc,status"` sorts ascending by
	// name and descending by status.
	Sort *string `pulumi:"sort"`
	// Search for images with a specific tag.
	Tag *string `pulumi:"tag"`
	// A list of tags required to be set on the image (all
	// specified tags must be in the images tag list for it to be matched).
	Tags []string `pulumi:"tags"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". Defaults to "private".
	Visibility *string `pulumi:"visibility"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// The checksum of the data associated with the image.
	Checksum string `pulumi:"checksum"`
	// The format of the image's container.
	ContainerFormat *string `pulumi:"containerFormat"`
	// The date the image was created.
	CreatedAt string `pulumi:"createdAt"`
	// The format of the image's disk.
	DiskFormat *string `pulumi:"diskFormat"`
	// the trailing path after the glance endpoint that represent the
	// location of the image or the path to retrieve it.
	File   string `pulumi:"file"`
	Hidden *bool  `pulumi:"hidden"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	MemberStatus *string `pulumi:"memberStatus"`
	// The metadata associated with the image. Image metadata allow for
	// meaningfully define the image properties and tags. See
	// <https://docs.openstack.org/glance/latest/user/metadefs-concepts.html>.
	Metadata map[string]string `pulumi:"metadata"`
	// The minimum amount of disk space required to use the image.
	MinDiskGb int `pulumi:"minDiskGb"`
	// The minimum amount of ram required to use the image.
	MinRamMb   int     `pulumi:"minRamMb"`
	MostRecent *bool   `pulumi:"mostRecent"`
	Name       *string `pulumi:"name"`
	NameRegex  *string `pulumi:"nameRegex"`
	Owner      *string `pulumi:"owner"`
	// Freeform information about the image.
	Properties map[string]string `pulumi:"properties"`
	// Whether or not the image is protected.
	Protected bool   `pulumi:"protected"`
	Region    string `pulumi:"region"`
	// The path to the JSON-schema that represent the image
	Schema string `pulumi:"schema"`
	// The size of the image (in bytes).
	SizeBytes int     `pulumi:"sizeBytes"`
	SizeMax   *int    `pulumi:"sizeMax"`
	SizeMin   *int    `pulumi:"sizeMin"`
	Sort      *string `pulumi:"sort"`
	Tag       *string `pulumi:"tag"`
	// The tags list of the image.
	Tags []string `pulumi:"tags"`
	// The date the image was last updated.
	UpdatedAt  string  `pulumi:"updatedAt"`
	Visibility *string `pulumi:"visibility"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			var s LookupImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImageResultOutput)
}

// A collection of arguments for invoking getImage.
type LookupImageOutputArgs struct {
	// The container format of the image.
	ContainerFormat pulumi.StringPtrInput `pulumi:"containerFormat"`
	// The disk format of the image.
	DiskFormat pulumi.StringPtrInput `pulumi:"diskFormat"`
	// Whether or not the image is hidden from public list.
	Hidden pulumi.BoolPtrInput `pulumi:"hidden"`
	// The status of the image. Must be one of
	// "accepted", "pending", "rejected", or "all".
	MemberStatus pulumi.StringPtrInput `pulumi:"memberStatus"`
	// If more than one result is returned, use the most
	// recent image.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// The name of the image. Cannot be used simultaneously with
	// `nameRegex`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The regular expressian of the name of the image.
	// Cannot be used simultaneously with `name`. Unlike filtering by `name` the
	// `nameRegex` filtering does by client on the result of OpenStack search
	// query.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The owner (UUID) of the image.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
	// a map of key/value pairs to match an image with.
	// All specified properties must be matched. Unlike other options filtering by
	// `properties` does by client on the result of OpenStack search query.
	// Filtering is applied if server responce contains at least 2 images. In case
	// there is only one image the `properties` ignores.
	Properties pulumi.StringMapInput `pulumi:"properties"`
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The maximum size (in bytes) of the image to return.
	SizeMax pulumi.IntPtrInput `pulumi:"sizeMax"`
	// The minimum size (in bytes) of the image to return.
	SizeMin pulumi.IntPtrInput `pulumi:"sizeMin"`
	// Sorts the response by one or more attribute and sort
	// direction combinations. You can also set multiple sort keys and directions.
	// Default direction is `desc`. Use the comma (,) character to separate multiple
	// values. For example expression `sort = "name:asc,status"` sorts ascending by
	// name and descending by status.
	Sort pulumi.StringPtrInput `pulumi:"sort"`
	// Search for images with a specific tag.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// A list of tags required to be set on the image (all
	// specified tags must be in the images tag list for it to be matched).
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". Defaults to "private".
	Visibility pulumi.StringPtrInput `pulumi:"visibility"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

// A collection of values returned by getImage.
type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

// The checksum of the data associated with the image.
func (o LookupImageResultOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Checksum }).(pulumi.StringOutput)
}

// The format of the image's container.
func (o LookupImageResultOutput) ContainerFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.ContainerFormat }).(pulumi.StringPtrOutput)
}

// The date the image was created.
func (o LookupImageResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The format of the image's disk.
func (o LookupImageResultOutput) DiskFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.DiskFormat }).(pulumi.StringPtrOutput)
}

// the trailing path after the glance endpoint that represent the
// location of the image or the path to retrieve it.
func (o LookupImageResultOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.File }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Hidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *bool { return v.Hidden }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) MemberStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.MemberStatus }).(pulumi.StringPtrOutput)
}

// The metadata associated with the image. Image metadata allow for
// meaningfully define the image properties and tags. See
// <https://docs.openstack.org/glance/latest/user/metadefs-concepts.html>.
func (o LookupImageResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The minimum amount of disk space required to use the image.
func (o LookupImageResultOutput) MinDiskGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.MinDiskGb }).(pulumi.IntOutput)
}

// The minimum amount of ram required to use the image.
func (o LookupImageResultOutput) MinRamMb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.MinRamMb }).(pulumi.IntOutput)
}

func (o LookupImageResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

func (o LookupImageResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

// Freeform information about the image.
func (o LookupImageResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Whether or not the image is protected.
func (o LookupImageResultOutput) Protected() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.Protected }).(pulumi.BoolOutput)
}

func (o LookupImageResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Region }).(pulumi.StringOutput)
}

// The path to the JSON-schema that represent the image
func (o LookupImageResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Schema }).(pulumi.StringOutput)
}

// The size of the image (in bytes).
func (o LookupImageResultOutput) SizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.SizeBytes }).(pulumi.IntOutput)
}

func (o LookupImageResultOutput) SizeMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *int { return v.SizeMax }).(pulumi.IntPtrOutput)
}

func (o LookupImageResultOutput) SizeMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *int { return v.SizeMin }).(pulumi.IntPtrOutput)
}

func (o LookupImageResultOutput) Sort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Sort }).(pulumi.StringPtrOutput)
}

func (o LookupImageResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

// The tags list of the image.
func (o LookupImageResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date the image was last updated.
func (o LookupImageResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupImageResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImageResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
