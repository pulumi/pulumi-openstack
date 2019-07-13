// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get a list of availability zones from OpenStack
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_availability_zones_v2.html.markdown.
func LookupAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs) (*GetAvailabilityZonesResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["region"] = args.Region
		inputs["state"] = args.State
	}
	outputs, err := ctx.Invoke("openstack:compute/getAvailabilityZones:getAvailabilityZones", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAvailabilityZonesResult{
		Names: outputs["names"],
		Region: outputs["region"],
		State: outputs["state"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// The `region` to fetch availability zones from, defaults to the provider's `region`
	Region interface{}
	// The `state` of the availability zones to match, default ("available").
	State interface{}
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	// The names of the availability zones, ordered alphanumerically, that match the queried `state`
	Names interface{}
	Region interface{}
	State interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}