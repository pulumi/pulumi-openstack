// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of Openstack Image IDs matching the
// specified criteria.
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
// 		opt0 := "^Ubuntu 16\\.04.*-amd64"
// 		opt1 := "updated_at"
// 		_, err := images.GetImageIds(ctx, &images.GetImageIdsArgs{
// 			NameRegex: &opt0,
// 			Properties: map[string]interface{}{
// 				"key": "value",
// 			},
// 			Sort: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetImageIds(ctx *pulumi.Context, args *GetImageIdsArgs, opts ...pulumi.InvokeOption) (*GetImageIdsResult, error) {
	var rv GetImageIdsResult
	err := ctx.Invoke("openstack:images/getImageIds:getImageIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageIds.
type GetImageIdsArgs struct {
	// The status of the image. Must be one of
	// "accepted", "pending", "rejected", or "all".
	MemberStatus *string `pulumi:"memberStatus"`
	// The name of the image. Cannot be used simultaneously
	// with `nameRegex`.
	Name *string `pulumi:"name"`
	// The regular expressian of the name of the image.
	// Cannot be used simultaneously with `name`. Unlike filtering by `name` the
	// `nameRegex` filtering does by client on the result of OpenStack search
	// query.
	NameRegex *string `pulumi:"nameRegex"`
	// The owner (UUID) of the image.
	Owner *string `pulumi:"owner"`
	// a map of key/value pairs to match an image with.
	// All specified properties must be matched. Unlike other options filtering
	// by `properties` does by client on the result of OpenStack search query.
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
	// Sorts the response by one or more attribute and sort
	// direction combinations. You can also set multiple sort keys and directions.
	// Default direction is `desc`. Use the comma (,) character to separate
	// multiple values. For example expression `sort = "name:asc,status"`
	// sorts ascending by name and descending by status. `sort` cannot be used
	// simultaneously with `sortKey`. If both are present in a configuration
	// then only `sort` will be used.
	Sort *string `pulumi:"sort"`
	// Order the results in either `asc` or `desc`.
	// Can be applied only with `sortKey`. Defaults to `asc`
	//
	// Deprecated: Use option 'sort' instead.
	SortDirection *string `pulumi:"sortDirection"`
	// Sort images based on a certain key. Defaults to
	// `name`. `sortKey` cannot be used simultaneously with `sort`. If both
	// are present in a configuration then only `sort` will be used.
	//
	// Deprecated: Use option 'sort' instead.
	SortKey *string `pulumi:"sortKey"`
	// Search for images with a specific tag.
	Tag *string `pulumi:"tag"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". Defaults to "private".
	Visibility *string `pulumi:"visibility"`
}

// A collection of values returned by getImageIds.
type GetImageIdsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string                 `pulumi:"id"`
	Ids          []string               `pulumi:"ids"`
	MemberStatus *string                `pulumi:"memberStatus"`
	Name         *string                `pulumi:"name"`
	NameRegex    *string                `pulumi:"nameRegex"`
	Owner        *string                `pulumi:"owner"`
	Properties   map[string]interface{} `pulumi:"properties"`
	Region       string                 `pulumi:"region"`
	SizeMax      *int                   `pulumi:"sizeMax"`
	SizeMin      *int                   `pulumi:"sizeMin"`
	Sort         *string                `pulumi:"sort"`
	// Deprecated: Use option 'sort' instead.
	SortDirection *string `pulumi:"sortDirection"`
	// Deprecated: Use option 'sort' instead.
	SortKey    *string `pulumi:"sortKey"`
	Tag        *string `pulumi:"tag"`
	Visibility *string `pulumi:"visibility"`
}
