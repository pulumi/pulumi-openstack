// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package images

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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
//
//	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/images"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := images.GetImageIds(ctx, &images.GetImageIdsArgs{
//				NameRegex: pulumi.StringRef("^Ubuntu 16\\.04.*-amd64"),
//				Properties: map[string]interface{}{
//					"key": "value",
//				},
//				Sort: pulumi.StringRef("updated_at"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetImageIds(ctx *pulumi.Context, args *GetImageIdsArgs, opts ...pulumi.InvokeOption) (*GetImageIdsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
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
	// A list of tags required to be set on the image
	// (all specified tags must be in the images tag list for it to be matched).
	Tags []string `pulumi:"tags"`
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
	SortKey    *string  `pulumi:"sortKey"`
	Tag        *string  `pulumi:"tag"`
	Tags       []string `pulumi:"tags"`
	Visibility *string  `pulumi:"visibility"`
}

func GetImageIdsOutput(ctx *pulumi.Context, args GetImageIdsOutputArgs, opts ...pulumi.InvokeOption) GetImageIdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageIdsResult, error) {
			args := v.(GetImageIdsArgs)
			r, err := GetImageIds(ctx, &args, opts...)
			var s GetImageIdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageIdsResultOutput)
}

// A collection of arguments for invoking getImageIds.
type GetImageIdsOutputArgs struct {
	// The status of the image. Must be one of
	// "accepted", "pending", "rejected", or "all".
	MemberStatus pulumi.StringPtrInput `pulumi:"memberStatus"`
	// The name of the image. Cannot be used simultaneously
	// with `nameRegex`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The regular expressian of the name of the image.
	// Cannot be used simultaneously with `name`. Unlike filtering by `name` the
	// `nameRegex` filtering does by client on the result of OpenStack search
	// query.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The owner (UUID) of the image.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
	// a map of key/value pairs to match an image with.
	// All specified properties must be matched. Unlike other options filtering
	// by `properties` does by client on the result of OpenStack search query.
	Properties pulumi.MapInput `pulumi:"properties"`
	// The region in which to obtain the V2 Glance client.
	// A Glance client is needed to create an Image that can be used with
	// a compute instance. If omitted, the `region` argument of the provider
	// is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The maximum size (in bytes) of the image to return.
	SizeMax pulumi.IntPtrInput `pulumi:"sizeMax"`
	// The minimum size (in bytes) of the image to return.
	SizeMin pulumi.IntPtrInput `pulumi:"sizeMin"`
	// Sorts the response by one or more attribute and sort
	// direction combinations. You can also set multiple sort keys and directions.
	// Default direction is `desc`. Use the comma (,) character to separate
	// multiple values. For example expression `sort = "name:asc,status"`
	// sorts ascending by name and descending by status. `sort` cannot be used
	// simultaneously with `sortKey`. If both are present in a configuration
	// then only `sort` will be used.
	Sort pulumi.StringPtrInput `pulumi:"sort"`
	// Order the results in either `asc` or `desc`.
	// Can be applied only with `sortKey`. Defaults to `asc`
	//
	// Deprecated: Use option 'sort' instead.
	SortDirection pulumi.StringPtrInput `pulumi:"sortDirection"`
	// Sort images based on a certain key. Defaults to
	// `name`. `sortKey` cannot be used simultaneously with `sort`. If both
	// are present in a configuration then only `sort` will be used.
	//
	// Deprecated: Use option 'sort' instead.
	SortKey pulumi.StringPtrInput `pulumi:"sortKey"`
	// Search for images with a specific tag.
	Tag pulumi.StringPtrInput `pulumi:"tag"`
	// A list of tags required to be set on the image
	// (all specified tags must be in the images tag list for it to be matched).
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The visibility of the image. Must be one of
	// "public", "private", "community", or "shared". Defaults to "private".
	Visibility pulumi.StringPtrInput `pulumi:"visibility"`
}

func (GetImageIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageIdsArgs)(nil)).Elem()
}

// A collection of values returned by getImageIds.
type GetImageIdsResultOutput struct{ *pulumi.OutputState }

func (GetImageIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageIdsResult)(nil)).Elem()
}

func (o GetImageIdsResultOutput) ToGetImageIdsResultOutput() GetImageIdsResultOutput {
	return o
}

func (o GetImageIdsResultOutput) ToGetImageIdsResultOutputWithContext(ctx context.Context) GetImageIdsResultOutput {
	return o
}

func (o GetImageIdsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetImageIdsResult] {
	return pulumix.Output[GetImageIdsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetImageIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetImageIdsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetImageIdsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetImageIdsResultOutput) MemberStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.MemberStatus }).(pulumi.StringPtrOutput)
}

func (o GetImageIdsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetImageIdsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetImageIdsResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o GetImageIdsResultOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v GetImageIdsResult) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

func (o GetImageIdsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageIdsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetImageIdsResultOutput) SizeMax() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *int { return v.SizeMax }).(pulumi.IntPtrOutput)
}

func (o GetImageIdsResultOutput) SizeMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *int { return v.SizeMin }).(pulumi.IntPtrOutput)
}

func (o GetImageIdsResultOutput) Sort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.Sort }).(pulumi.StringPtrOutput)
}

// Deprecated: Use option 'sort' instead.
func (o GetImageIdsResultOutput) SortDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.SortDirection }).(pulumi.StringPtrOutput)
}

// Deprecated: Use option 'sort' instead.
func (o GetImageIdsResultOutput) SortKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.SortKey }).(pulumi.StringPtrOutput)
}

func (o GetImageIdsResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o GetImageIdsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetImageIdsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GetImageIdsResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageIdsResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageIdsResultOutput{})
}
