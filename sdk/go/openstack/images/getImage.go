// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
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
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/images"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		opt1 := "Ubuntu 16.04"
// 		_, err := images.LookupImage(ctx, &images.LookupImageArgs{
// 			MostRecent: &opt0,
// 			Name:       &opt1,
// 			Properties: map[string]interface{}{
// 				"key": "value",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("openstack:images/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// Whether or not the image is hidden from public list.
	Hidden *bool `pulumi:"hidden"`
	// The status of the image. Must be one of
	// "accepted", "pending", "rejected", or "all".
	MemberStatus *string `pulumi:"memberStatus"`
	// If more than one result is returned, use the most
	// recent image.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// The owner (UUID) of the image.
	Owner *string `pulumi:"owner"`
	// a map of key/value pairs to match an image with.
	// All specified properties must be matched. Unlike other options filtering
	// by `properties` does by client on the result of OpenStack search query.
	// Filtering is applied if server responce contains at least 2 images. In
	// case there is only one image the `properties` ignores.
	Properties map[string]interface{} `pulumi:"properties"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used.
	Region *string `pulumi:"region"`
	// The maximum size (in bytes) of the image to return.
	SizeMax *int `pulumi:"sizeMax"`
	// The minimum size (in bytes) of the image to return.
	SizeMin *int `pulumi:"sizeMin"`
	// Order the results in either `asc` or `desc`.
	SortDirection *string `pulumi:"sortDirection"`
	// Sort images based on a certain key. Defaults to `name`.
	SortKey *string `pulumi:"sortKey"`
	// Search for images with a specific tag.
	Tag *string `pulumi:"tag"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". Defaults to "private".
	Visibility *string `pulumi:"visibility"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// The checksum of the data associated with the image.
	Checksum        string `pulumi:"checksum"`
	ContainerFormat string `pulumi:"containerFormat"`
	// The date the image was created.
	// * `containerFormat`: The format of the image's container.
	// * `diskFormat`: The format of the image's disk.
	CreatedAt  string `pulumi:"createdAt"`
	DiskFormat string `pulumi:"diskFormat"`
	// the trailing path after the glance endpoint that represent the
	// location of the image or the path to retrieve it.
	File   string `pulumi:"file"`
	Hidden *bool  `pulumi:"hidden"`
	// The provider-assigned unique ID for this managed resource.
	Id           string  `pulumi:"id"`
	MemberStatus *string `pulumi:"memberStatus"`
	// The metadata associated with the image.
	// Image metadata allow for meaningfully define the image properties
	// and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The minimum amount of disk space required to use the image.
	MinDiskGb int `pulumi:"minDiskGb"`
	// The minimum amount of ram required to use the image.
	MinRamMb   int     `pulumi:"minRamMb"`
	MostRecent *bool   `pulumi:"mostRecent"`
	Name       *string `pulumi:"name"`
	Owner      *string `pulumi:"owner"`
	// Freeform information about the image.
	Properties map[string]interface{} `pulumi:"properties"`
	// Whether or not the image is protected.
	Protected bool   `pulumi:"protected"`
	Region    string `pulumi:"region"`
	// The path to the JSON-schema that represent
	// the image or image
	Schema string `pulumi:"schema"`
	// The size of the image (in bytes).
	SizeBytes     int     `pulumi:"sizeBytes"`
	SizeMax       *int    `pulumi:"sizeMax"`
	SizeMin       *int    `pulumi:"sizeMin"`
	SortDirection *string `pulumi:"sortDirection"`
	SortKey       *string `pulumi:"sortKey"`
	Tag           *string `pulumi:"tag"`
	// The tags list of the image.
	Tags []string `pulumi:"tags"`
	// The date the image was last updated.
	UpdatedAt  string  `pulumi:"updatedAt"`
	Visibility *string `pulumi:"visibility"`
}
