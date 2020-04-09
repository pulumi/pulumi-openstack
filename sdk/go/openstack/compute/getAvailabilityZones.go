// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get a list of availability zones from OpenStack
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_availability_zones_v2.html.markdown.
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("openstack:compute/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// The `region` to fetch availability zones from, defaults to the provider's `region`
	Region *string `pulumi:"region"`
	// The `state` of the availability zones to match, default ("available").
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The names of the availability zones, ordered alphanumerically, that match the queried `state`
	Names  []string `pulumi:"names"`
	Region string   `pulumi:"region"`
	State  *string  `pulumi:"state"`
}
