// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package blockstorage

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get a list of Block Storage availability zones from OpenStack
func LookupAvailabilityZonesV3(ctx *pulumi.Context, args *GetAvailabilityZonesV3Args) (*GetAvailabilityZonesV3Result, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["region"] = args.Region
		inputs["state"] = args.State
	}
	outputs, err := ctx.Invoke("openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAvailabilityZonesV3Result{
		Names: outputs["names"],
		Region: outputs["region"],
		State: outputs["state"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAvailabilityZonesV3.
type GetAvailabilityZonesV3Args struct {
	// The region in which to obtain the Block Storage client.
	// If omitted, the `region` argument of the provider is used.
	Region interface{}
	// The `state` of the availability zones to match. Can
	// either be `available` or `unavailable`. Default is `available`.
	State interface{}
}

// A collection of values returned by getAvailabilityZonesV3.
type GetAvailabilityZonesV3Result struct {
	// The names of the availability zones, ordered alphanumerically, that
	// match the queried `state`.
	Names interface{}
	// See Argument Reference above.
	Region interface{}
	// See Argument Reference above.
	State interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
