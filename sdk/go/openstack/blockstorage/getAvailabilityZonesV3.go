// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get a list of Block Storage availability zones from OpenStack
func GetAvailabilityZonesV3(ctx *pulumi.Context, args *GetAvailabilityZonesV3Args, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesV3Result, error) {
	var rv GetAvailabilityZonesV3Result
	err := ctx.Invoke("openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZonesV3.
type GetAvailabilityZonesV3Args struct {
	// The region in which to obtain the Block Storage client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The `state` of the availability zones to match. Can
	// either be `available` or `unavailable`. Default is `available`.
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZonesV3.
type GetAvailabilityZonesV3Result struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The names of the availability zones, ordered alphanumerically, that
	// match the queried `state`.
	Names []string `pulumi:"names"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	State *string `pulumi:"state"`
}
