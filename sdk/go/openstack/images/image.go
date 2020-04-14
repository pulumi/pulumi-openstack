// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a V2 Image resource within OpenStack Glance.
//
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
// In addition, the `directUrl` property is also automatically reconciled if the
// Image Service set it.
type Image struct {
	pulumi.CustomResourceState

	// The checksum of the data associated with the image.
	Checksum pulumi.StringOutput `pulumi:"checksum"`
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat pulumi.StringOutput `pulumi:"containerFormat"`
	// The date the image was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat pulumi.StringOutput `pulumi:"diskFormat"`
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File           pulumi.StringOutput    `pulumi:"file"`
	ImageCachePath pulumi.StringPtrOutput `pulumi:"imageCachePath"`
	// This is the url of the raw image that will
	// be downloaded in the `imageCachePath` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `localFilePath`.
	ImageSourceUrl pulumi.StringPtrOutput `pulumi:"imageSourceUrl"`
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl`.
	LocalFilePath pulumi.StringPtrOutput `pulumi:"localFilePath"`
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb pulumi.IntPtrOutput `pulumi:"minDiskGb"`
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb pulumi.IntPtrOutput `pulumi:"minRamMb"`
	// The name of the image.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the openstack user who owns the image.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further
	// information about properties.
	Properties pulumi.MapOutput `pulumi:"properties"`
	// If true, image will not be deletable.
	// Defaults to false.
	Protected pulumi.BoolPtrOutput `pulumi:"protected"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region pulumi.StringOutput `pulumi:"region"`
	// The path to the JSON-schema that represent
	// the image or image
	Schema pulumi.StringOutput `pulumi:"schema"`
	// The size in bytes of the data associated with the image.
	SizeBytes pulumi.IntOutput `pulumi:"sizeBytes"`
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// (**Deprecated** - use `updatedAt` instead)
	UpdateAt pulumi.StringOutput `pulumi:"updateAt"`
	// The date the image was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum pulumi.BoolPtrOutput `pulumi:"verifyChecksum"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility pulumi.StringPtrOutput `pulumi:"visibility"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil || args.ContainerFormat == nil {
		return nil, errors.New("missing required argument 'ContainerFormat'")
	}
	if args == nil || args.DiskFormat == nil {
		return nil, errors.New("missing required argument 'DiskFormat'")
	}
	if args == nil {
		args = &ImageArgs{}
	}
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
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat *string `pulumi:"containerFormat"`
	// The date the image was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat *string `pulumi:"diskFormat"`
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File           *string `pulumi:"file"`
	ImageCachePath *string `pulumi:"imageCachePath"`
	// This is the url of the raw image that will
	// be downloaded in the `imageCachePath` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `localFilePath`.
	ImageSourceUrl *string `pulumi:"imageSourceUrl"`
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl`.
	LocalFilePath *string `pulumi:"localFilePath"`
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb *int `pulumi:"minDiskGb"`
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb *int `pulumi:"minRamMb"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// The id of the openstack user who owns the image.
	Owner *string `pulumi:"owner"`
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further
	// information about properties.
	Properties map[string]interface{} `pulumi:"properties"`
	// If true, image will not be deletable.
	// Defaults to false.
	Protected *bool `pulumi:"protected"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region *string `pulumi:"region"`
	// The path to the JSON-schema that represent
	// the image or image
	Schema *string `pulumi:"schema"`
	// The size in bytes of the data associated with the image.
	SizeBytes *int `pulumi:"sizeBytes"`
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status *string `pulumi:"status"`
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags []string `pulumi:"tags"`
	// (**Deprecated** - use `updatedAt` instead)
	UpdateAt *string `pulumi:"updateAt"`
	// The date the image was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum *bool `pulumi:"verifyChecksum"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility *string `pulumi:"visibility"`
}

type ImageState struct {
	// The checksum of the data associated with the image.
	Checksum pulumi.StringPtrInput
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat pulumi.StringPtrInput
	// The date the image was created.
	CreatedAt pulumi.StringPtrInput
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat pulumi.StringPtrInput
	// the trailing path after the glance
	// endpoint that represent the location of the image
	// or the path to retrieve it.
	File           pulumi.StringPtrInput
	ImageCachePath pulumi.StringPtrInput
	// This is the url of the raw image that will
	// be downloaded in the `imageCachePath` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `localFilePath`.
	ImageSourceUrl pulumi.StringPtrInput
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl`.
	LocalFilePath pulumi.StringPtrInput
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata pulumi.MapInput
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb pulumi.IntPtrInput
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// The id of the openstack user who owns the image.
	Owner pulumi.StringPtrInput
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further
	// information about properties.
	Properties pulumi.MapInput
	// If true, image will not be deletable.
	// Defaults to false.
	Protected pulumi.BoolPtrInput
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region pulumi.StringPtrInput
	// The path to the JSON-schema that represent
	// the image or image
	Schema pulumi.StringPtrInput
	// The size in bytes of the data associated with the image.
	SizeBytes pulumi.IntPtrInput
	// The status of the image. It can be "queued", "active"
	// or "saving".
	Status pulumi.StringPtrInput
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags pulumi.StringArrayInput
	// (**Deprecated** - use `updatedAt` instead)
	UpdateAt pulumi.StringPtrInput
	// The date the image was last updated.
	UpdatedAt pulumi.StringPtrInput
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum pulumi.BoolPtrInput
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility pulumi.StringPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat string `pulumi:"containerFormat"`
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat     string  `pulumi:"diskFormat"`
	ImageCachePath *string `pulumi:"imageCachePath"`
	// This is the url of the raw image that will
	// be downloaded in the `imageCachePath` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `localFilePath`.
	ImageSourceUrl *string `pulumi:"imageSourceUrl"`
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl`.
	LocalFilePath *string `pulumi:"localFilePath"`
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb *int `pulumi:"minDiskGb"`
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb *int `pulumi:"minRamMb"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further
	// information about properties.
	Properties map[string]interface{} `pulumi:"properties"`
	// If true, image will not be deletable.
	// Defaults to false.
	Protected *bool `pulumi:"protected"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region *string `pulumi:"region"`
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags []string `pulumi:"tags"`
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum *bool `pulumi:"verifyChecksum"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The container format. Must be one of
	// "ami", "ari", "aki", "bare", "ovf".
	ContainerFormat pulumi.StringInput
	// The disk format. Must be one of
	// "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".
	DiskFormat     pulumi.StringInput
	ImageCachePath pulumi.StringPtrInput
	// This is the url of the raw image that will
	// be downloaded in the `imageCachePath` before being uploaded to Glance.
	// Glance is able to download image from internet but the `gophercloud` library
	// does not yet provide a way to do so.
	// Conflicts with `localFilePath`.
	ImageSourceUrl pulumi.StringPtrInput
	// This is the filepath of the raw image file
	// that will be uploaded to Glance. Conflicts with `imageSourceUrl`.
	LocalFilePath pulumi.StringPtrInput
	// Amount of disk space (in GB) required to boot image.
	// Defaults to 0.
	MinDiskGb pulumi.IntPtrInput
	// Amount of ram (in MB) required to boot image.
	// Defauts to 0.
	MinRamMb pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// A map of key/value pairs to set freeform
	// information about an image. See the "Notes" section for further
	// information about properties.
	Properties pulumi.MapInput
	// If true, image will not be deletable.
	// Defaults to false.
	Protected pulumi.BoolPtrInput
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used. Changing this creates a new Image.
	Region pulumi.StringPtrInput
	// The tags of the image. It must be a list of strings.
	// At this time, it is not possible to delete all tags of an image.
	Tags pulumi.StringArrayInput
	// If false, the checksum will not be verified
	// once the image is finished uploading. Defaults to true.
	VerifyChecksum pulumi.BoolPtrInput
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". The ability to set the
	// visibility depends upon the configuration of the OpenStack cloud.
	Visibility pulumi.StringPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}
