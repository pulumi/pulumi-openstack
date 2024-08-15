// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V2 Image resource within OpenStack Glance.
//
// > **Note:** All arguments including the source image URL password will be
// stored in the raw state as plain-text. Read more about sensitive data in
// state.
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
//			_, err := images.NewImage(ctx, "rancheros", &images.ImageArgs{
//				Name:            pulumi.String("RancherOS"),
//				ImageSourceUrl:  pulumi.String("https://releases.rancher.com/os/latest/rancheros-openstack.img"),
//				ContainerFormat: pulumi.String("bare"),
//				DiskFormat:      pulumi.String("qcow2"),
//				Properties: pulumi.StringMap{
//					"key": pulumi.String("value"),
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
// ## Notes
//
// ### Properties
//
// This resource supports the ability to add properties to a resource during
// creation as well as add, update, and delete properties during an update of this
// resource.
//
// Newer versions of OpenStack are adding some read-only properties to each image.
// These properties start with the prefix `os_`. If these properties are detected,
// this resource will automatically reconcile these with the user-provided
// properties.
//
// In addition, the `directUrl` and `stores` properties are also automatically reconciled if the
// Image Service set it.
//
// ## Import
//
// Images can be imported using the `id`, e.g.
//
// ```sh
// $ pulumi import openstack:images/image:Image rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2
// ```
type Image struct {
	pulumi.CustomResourceState

	// The checksum of the data associated with the image.
	Checksum pulumi.StringOutput `pulumi:"checksum"`
	// The container format. Must be one of "bare",
	// "ovf", "aki", "ari", "ami", "ova", "docker", "compressed".
	ContainerFormat pulumi.StringOutput `pulumi:"containerFormat"`
	// The date the image was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// If true, this provider will decompress downloaded
	// image before uploading it to OpenStack. Decompression algorithm is chosen by
	// checking "Content-Type" or `Content-Disposition` header to detect the
	// filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
	// Defaults to false. Changing this creates a new Image.
	Decompress pulumi.BoolPtrOutput `pulumi:"decompress"`
	// The disk format. Must be one of "raw", "vhd",
	// "vhdx", "vmdk", "vdi", "iso", "ploop", "qcow2", "aki", "ari", "ami"
	DiskFormat pulumi.StringOutput `pulumi:"diskFormat"`
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File pulumi.StringOutput `pulumi:"file"`
	// If true, image will be hidden from public list.
	// Defaults to false.
	Hidden         pulumi.BoolPtrOutput   `pulumi:"hidden"`
	ImageCachePath pulumi.StringPtrOutput `pulumi:"imageCachePath"`
	// Unique ID (valid UUID) of image to create. Changing
	// this creates a new image.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The password of basic auth to download
	// `imageSourceUrl`.
	ImageSourcePassword pulumi.StringPtrOutput `pulumi:"imageSourcePassword"`
	// This is the url of the raw image. If
	// `webDownload` is not used, then the image will be downloaded in the
	// `imageCachePath` before being uploaded to Glance. Conflicts with
	// `localFilePath`.
	ImageSourceUrl pulumi.StringPtrOutput `pulumi:"imageSourceUrl"`
	// The username of basic auth to download
	// `imageSourceUrl`.
	ImageSourceUsername pulumi.StringPtrOutput `pulumi:"imageSourceUsername"`
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl` and
	// `webDownload`.
	LocalFilePath pulumi.StringPtrOutput `pulumi:"localFilePath"`
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Amount of disk space (in GB) required to boot
	// image. Defaults to 0.
	MinDiskGb pulumi.IntPtrOutput `pulumi:"minDiskGb"`
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb pulumi.IntPtrOutput `pulumi:"minRamMb"`
	// The name of the image.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the openstack user who owns the image.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further information
	// about properties.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// If true, image will not be deletable. Defaults to
	// false.
	Protected pulumi.BoolPtrOutput `pulumi:"protected"`
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used. Changing
	// this creates a new Image.
	Region pulumi.StringOutput `pulumi:"region"`
	// The path to the JSON-schema that represent
	// the image or image
	Schema pulumi.StringOutput `pulumi:"schema"`
	// The size in bytes of the data associated with the image.
	SizeBytes pulumi.IntOutput `pulumi:"sizeBytes"`
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of the image. It must be a list of strings. At
	// this time, it is not possible to delete all tags of an image.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date the image was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// If false, the checksum will not be verified
	// once the image is finished uploading. Conflicts with `webDownload`. Defaults
	// to true when not using `webDownload`.
	VerifyChecksum pulumi.BoolPtrOutput `pulumi:"verifyChecksum"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
	// If true, the "web-download" import method will be
	// used to let Openstack download the image directly from the remote source.
	// Conflicts with `localFilePath`. Defaults to false.
	WebDownload pulumi.BoolPtrOutput `pulumi:"webDownload"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerFormat == nil {
		return nil, errors.New("invalid value for required argument 'ContainerFormat'")
	}
	if args.DiskFormat == nil {
		return nil, errors.New("invalid value for required argument 'DiskFormat'")
	}
	if args.ImageSourcePassword != nil {
		args.ImageSourcePassword = pulumi.ToSecret(args.ImageSourcePassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"imageSourcePassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("openstack:images/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("openstack:images/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// The checksum of the data associated with the image.
	Checksum *string `pulumi:"checksum"`
	// The container format. Must be one of "bare",
	// "ovf", "aki", "ari", "ami", "ova", "docker", "compressed".
	ContainerFormat *string `pulumi:"containerFormat"`
	// The date the image was created.
	CreatedAt *string `pulumi:"createdAt"`
	// If true, this provider will decompress downloaded
	// image before uploading it to OpenStack. Decompression algorithm is chosen by
	// checking "Content-Type" or `Content-Disposition` header to detect the
	// filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
	// Defaults to false. Changing this creates a new Image.
	Decompress *bool `pulumi:"decompress"`
	// The disk format. Must be one of "raw", "vhd",
	// "vhdx", "vmdk", "vdi", "iso", "ploop", "qcow2", "aki", "ari", "ami"
	DiskFormat *string `pulumi:"diskFormat"`
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File *string `pulumi:"file"`
	// If true, image will be hidden from public list.
	// Defaults to false.
	Hidden         *bool   `pulumi:"hidden"`
	ImageCachePath *string `pulumi:"imageCachePath"`
	// Unique ID (valid UUID) of image to create. Changing
	// this creates a new image.
	ImageId *string `pulumi:"imageId"`
	// The password of basic auth to download
	// `imageSourceUrl`.
	ImageSourcePassword *string `pulumi:"imageSourcePassword"`
	// This is the url of the raw image. If
	// `webDownload` is not used, then the image will be downloaded in the
	// `imageCachePath` before being uploaded to Glance. Conflicts with
	// `localFilePath`.
	ImageSourceUrl *string `pulumi:"imageSourceUrl"`
	// The username of basic auth to download
	// `imageSourceUrl`.
	ImageSourceUsername *string `pulumi:"imageSourceUsername"`
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl` and
	// `webDownload`.
	LocalFilePath *string `pulumi:"localFilePath"`
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata map[string]string `pulumi:"metadata"`
	// Amount of disk space (in GB) required to boot
	// image. Defaults to 0.
	MinDiskGb *int `pulumi:"minDiskGb"`
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb *int `pulumi:"minRamMb"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// The id of the openstack user who owns the image.
	Owner *string `pulumi:"owner"`
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further information
	// about properties.
	Properties map[string]string `pulumi:"properties"`
	// If true, image will not be deletable. Defaults to
	// false.
	Protected *bool `pulumi:"protected"`
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used. Changing
	// this creates a new Image.
	Region *string `pulumi:"region"`
	// The path to the JSON-schema that represent
	// the image or image
	Schema *string `pulumi:"schema"`
	// The size in bytes of the data associated with the image.
	SizeBytes *int `pulumi:"sizeBytes"`
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status *string `pulumi:"status"`
	// The tags of the image. It must be a list of strings. At
	// this time, it is not possible to delete all tags of an image.
	Tags []string `pulumi:"tags"`
	// The date the image was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// If false, the checksum will not be verified
	// once the image is finished uploading. Conflicts with `webDownload`. Defaults
	// to true when not using `webDownload`.
	VerifyChecksum *bool `pulumi:"verifyChecksum"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility *string `pulumi:"visibility"`
	// If true, the "web-download" import method will be
	// used to let Openstack download the image directly from the remote source.
	// Conflicts with `localFilePath`. Defaults to false.
	WebDownload *bool `pulumi:"webDownload"`
}

type ImageState struct {
	// The checksum of the data associated with the image.
	Checksum pulumi.StringPtrInput
	// The container format. Must be one of "bare",
	// "ovf", "aki", "ari", "ami", "ova", "docker", "compressed".
	ContainerFormat pulumi.StringPtrInput
	// The date the image was created.
	CreatedAt pulumi.StringPtrInput
	// If true, this provider will decompress downloaded
	// image before uploading it to OpenStack. Decompression algorithm is chosen by
	// checking "Content-Type" or `Content-Disposition` header to detect the
	// filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
	// Defaults to false. Changing this creates a new Image.
	Decompress pulumi.BoolPtrInput
	// The disk format. Must be one of "raw", "vhd",
	// "vhdx", "vmdk", "vdi", "iso", "ploop", "qcow2", "aki", "ari", "ami"
	DiskFormat pulumi.StringPtrInput
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File pulumi.StringPtrInput
	// If true, image will be hidden from public list.
	// Defaults to false.
	Hidden         pulumi.BoolPtrInput
	ImageCachePath pulumi.StringPtrInput
	// Unique ID (valid UUID) of image to create. Changing
	// this creates a new image.
	ImageId pulumi.StringPtrInput
	// The password of basic auth to download
	// `imageSourceUrl`.
	ImageSourcePassword pulumi.StringPtrInput
	// This is the url of the raw image. If
	// `webDownload` is not used, then the image will be downloaded in the
	// `imageCachePath` before being uploaded to Glance. Conflicts with
	// `localFilePath`.
	ImageSourceUrl pulumi.StringPtrInput
	// The username of basic auth to download
	// `imageSourceUrl`.
	ImageSourceUsername pulumi.StringPtrInput
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl` and
	// `webDownload`.
	LocalFilePath pulumi.StringPtrInput
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata pulumi.StringMapInput
	// Amount of disk space (in GB) required to boot
	// image. Defaults to 0.
	MinDiskGb pulumi.IntPtrInput
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// The id of the openstack user who owns the image.
	Owner pulumi.StringPtrInput
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further information
	// about properties.
	Properties pulumi.StringMapInput
	// If true, image will not be deletable. Defaults to
	// false.
	Protected pulumi.BoolPtrInput
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used. Changing
	// this creates a new Image.
	Region pulumi.StringPtrInput
	// The path to the JSON-schema that represent
	// the image or image
	Schema pulumi.StringPtrInput
	// The size in bytes of the data associated with the image.
	SizeBytes pulumi.IntPtrInput
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status pulumi.StringPtrInput
	// The tags of the image. It must be a list of strings. At
	// this time, it is not possible to delete all tags of an image.
	Tags pulumi.StringArrayInput
	// The date the image was last updated.
	UpdatedAt pulumi.StringPtrInput
	// If false, the checksum will not be verified
	// once the image is finished uploading. Conflicts with `webDownload`. Defaults
	// to true when not using `webDownload`.
	VerifyChecksum pulumi.BoolPtrInput
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility pulumi.StringPtrInput
	// If true, the "web-download" import method will be
	// used to let Openstack download the image directly from the remote source.
	// Conflicts with `localFilePath`. Defaults to false.
	WebDownload pulumi.BoolPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The container format. Must be one of "bare",
	// "ovf", "aki", "ari", "ami", "ova", "docker", "compressed".
	ContainerFormat string `pulumi:"containerFormat"`
	// If true, this provider will decompress downloaded
	// image before uploading it to OpenStack. Decompression algorithm is chosen by
	// checking "Content-Type" or `Content-Disposition` header to detect the
	// filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
	// Defaults to false. Changing this creates a new Image.
	Decompress *bool `pulumi:"decompress"`
	// The disk format. Must be one of "raw", "vhd",
	// "vhdx", "vmdk", "vdi", "iso", "ploop", "qcow2", "aki", "ari", "ami"
	DiskFormat string `pulumi:"diskFormat"`
	// If true, image will be hidden from public list.
	// Defaults to false.
	Hidden         *bool   `pulumi:"hidden"`
	ImageCachePath *string `pulumi:"imageCachePath"`
	// Unique ID (valid UUID) of image to create. Changing
	// this creates a new image.
	ImageId *string `pulumi:"imageId"`
	// The password of basic auth to download
	// `imageSourceUrl`.
	ImageSourcePassword *string `pulumi:"imageSourcePassword"`
	// This is the url of the raw image. If
	// `webDownload` is not used, then the image will be downloaded in the
	// `imageCachePath` before being uploaded to Glance. Conflicts with
	// `localFilePath`.
	ImageSourceUrl *string `pulumi:"imageSourceUrl"`
	// The username of basic auth to download
	// `imageSourceUrl`.
	ImageSourceUsername *string `pulumi:"imageSourceUsername"`
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl` and
	// `webDownload`.
	LocalFilePath *string `pulumi:"localFilePath"`
	// Amount of disk space (in GB) required to boot
	// image. Defaults to 0.
	MinDiskGb *int `pulumi:"minDiskGb"`
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb *int `pulumi:"minRamMb"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further information
	// about properties.
	Properties map[string]string `pulumi:"properties"`
	// If true, image will not be deletable. Defaults to
	// false.
	Protected *bool `pulumi:"protected"`
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used. Changing
	// this creates a new Image.
	Region *string `pulumi:"region"`
	// The tags of the image. It must be a list of strings. At
	// this time, it is not possible to delete all tags of an image.
	Tags []string `pulumi:"tags"`
	// If false, the checksum will not be verified
	// once the image is finished uploading. Conflicts with `webDownload`. Defaults
	// to true when not using `webDownload`.
	VerifyChecksum *bool `pulumi:"verifyChecksum"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility *string `pulumi:"visibility"`
	// If true, the "web-download" import method will be
	// used to let Openstack download the image directly from the remote source.
	// Conflicts with `localFilePath`. Defaults to false.
	WebDownload *bool `pulumi:"webDownload"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The container format. Must be one of "bare",
	// "ovf", "aki", "ari", "ami", "ova", "docker", "compressed".
	ContainerFormat pulumi.StringInput
	// If true, this provider will decompress downloaded
	// image before uploading it to OpenStack. Decompression algorithm is chosen by
	// checking "Content-Type" or `Content-Disposition` header to detect the
	// filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
	// Defaults to false. Changing this creates a new Image.
	Decompress pulumi.BoolPtrInput
	// The disk format. Must be one of "raw", "vhd",
	// "vhdx", "vmdk", "vdi", "iso", "ploop", "qcow2", "aki", "ari", "ami"
	DiskFormat pulumi.StringInput
	// If true, image will be hidden from public list.
	// Defaults to false.
	Hidden         pulumi.BoolPtrInput
	ImageCachePath pulumi.StringPtrInput
	// Unique ID (valid UUID) of image to create. Changing
	// this creates a new image.
	ImageId pulumi.StringPtrInput
	// The password of basic auth to download
	// `imageSourceUrl`.
	ImageSourcePassword pulumi.StringPtrInput
	// This is the url of the raw image. If
	// `webDownload` is not used, then the image will be downloaded in the
	// `imageCachePath` before being uploaded to Glance. Conflicts with
	// `localFilePath`.
	ImageSourceUrl pulumi.StringPtrInput
	// The username of basic auth to download
	// `imageSourceUrl`.
	ImageSourceUsername pulumi.StringPtrInput
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl` and
	// `webDownload`.
	LocalFilePath pulumi.StringPtrInput
	// Amount of disk space (in GB) required to boot
	// image. Defaults to 0.
	MinDiskGb pulumi.IntPtrInput
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further information
	// about properties.
	Properties pulumi.StringMapInput
	// If true, image will not be deletable. Defaults to
	// false.
	Protected pulumi.BoolPtrInput
	// The region in which to obtain the V2 Glance client. A
	// Glance client is needed to create an Image that can be used with a compute
	// instance. If omitted, the `region` argument of the provider is used. Changing
	// this creates a new Image.
	Region pulumi.StringPtrInput
	// The tags of the image. It must be a list of strings. At
	// this time, it is not possible to delete all tags of an image.
	Tags pulumi.StringArrayInput
	// If false, the checksum will not be verified
	// once the image is finished uploading. Conflicts with `webDownload`. Defaults
	// to true when not using `webDownload`.
	VerifyChecksum pulumi.BoolPtrInput
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility pulumi.StringPtrInput
	// If true, the "web-download" import method will be
	// used to let Openstack download the image directly from the remote source.
	// Conflicts with `localFilePath`. Defaults to false.
	WebDownload pulumi.BoolPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//	ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//	ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// The checksum of the data associated with the image.
func (o ImageOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Checksum }).(pulumi.StringOutput)
}

// The container format. Must be one of "bare",
// "ovf", "aki", "ari", "ami", "ova", "docker", "compressed".
func (o ImageOutput) ContainerFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ContainerFormat }).(pulumi.StringOutput)
}

// The date the image was created.
func (o ImageOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// If true, this provider will decompress downloaded
// image before uploading it to OpenStack. Decompression algorithm is chosen by
// checking "Content-Type" or `Content-Disposition` header to detect the
// filename extension. Supported algorithms are: gzip, bzip2, xz and zst.
// Defaults to false. Changing this creates a new Image.
func (o ImageOutput) Decompress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.Decompress }).(pulumi.BoolPtrOutput)
}

// The disk format. Must be one of "raw", "vhd",
// "vhdx", "vmdk", "vdi", "iso", "ploop", "qcow2", "aki", "ari", "ami"
func (o ImageOutput) DiskFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.DiskFormat }).(pulumi.StringOutput)
}

// the trailing path after the glance
// endpoint that represent the location of the image
// or the path to retrieve it.
func (o ImageOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.File }).(pulumi.StringOutput)
}

// If true, image will be hidden from public list.
// Defaults to false.
func (o ImageOutput) Hidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.Hidden }).(pulumi.BoolPtrOutput)
}

func (o ImageOutput) ImageCachePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageCachePath }).(pulumi.StringPtrOutput)
}

// Unique ID (valid UUID) of image to create. Changing
// this creates a new image.
func (o ImageOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The password of basic auth to download
// `imageSourceUrl`.
func (o ImageOutput) ImageSourcePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageSourcePassword }).(pulumi.StringPtrOutput)
}

// This is the url of the raw image. If
// `webDownload` is not used, then the image will be downloaded in the
// `imageCachePath` before being uploaded to Glance. Conflicts with
// `localFilePath`.
func (o ImageOutput) ImageSourceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageSourceUrl }).(pulumi.StringPtrOutput)
}

// The username of basic auth to download
// `imageSourceUrl`.
func (o ImageOutput) ImageSourceUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageSourceUsername }).(pulumi.StringPtrOutput)
}

// This is the filepath of the raw image file
// that will be uploaded to Glance. Conflicts with `imageSourceUrl` and
// `webDownload`.
func (o ImageOutput) LocalFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.LocalFilePath }).(pulumi.StringPtrOutput)
}

// The metadata associated with the image.
// Image metadata allow for meaningfully define the image properties
// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
func (o ImageOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Amount of disk space (in GB) required to boot
// image. Defaults to 0.
func (o ImageOutput) MinDiskGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.IntPtrOutput { return v.MinDiskGb }).(pulumi.IntPtrOutput)
}

// Amount of ram (in MB) required to boot image.
// Defauts to 0.
func (o ImageOutput) MinRamMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.IntPtrOutput { return v.MinRamMb }).(pulumi.IntPtrOutput)
}

// The name of the image.
func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the openstack user who owns the image.
func (o ImageOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// A map of key/value pairs to set freeform
// information about an image. See the "Notes" section for further information
// about properties.
func (o ImageOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// If true, image will not be deletable. Defaults to
// false.
func (o ImageOutput) Protected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.Protected }).(pulumi.BoolPtrOutput)
}

// The region in which to obtain the V2 Glance client. A
// Glance client is needed to create an Image that can be used with a compute
// instance. If omitted, the `region` argument of the provider is used. Changing
// this creates a new Image.
func (o ImageOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The path to the JSON-schema that represent
// the image or image
func (o ImageOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// The size in bytes of the data associated with the image.
func (o ImageOutput) SizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *Image) pulumi.IntOutput { return v.SizeBytes }).(pulumi.IntOutput)
}

// The status of the image. It can be "queued", "active"
// or "saving".
func (o ImageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of the image. It must be a list of strings. At
// this time, it is not possible to delete all tags of an image.
func (o ImageOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date the image was last updated.
func (o ImageOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// If false, the checksum will not be verified
// once the image is finished uploading. Conflicts with `webDownload`. Defaults
// to true when not using `webDownload`.
func (o ImageOutput) VerifyChecksum() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.VerifyChecksum }).(pulumi.BoolPtrOutput)
}

// The visibility of the image. Must be one of
// "public", "private", "community", or "shared". The ability to set the
// visibility depends upon the configuration of the OpenStack cloud.
func (o ImageOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Visibility }).(pulumi.StringPtrOutput)
}

// If true, the "web-download" import method will be
// used to let Openstack download the image directly from the remote source.
// Conflicts with `localFilePath`. Defaults to false.
func (o ImageOutput) WebDownload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.WebDownload }).(pulumi.BoolPtrOutput)
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
