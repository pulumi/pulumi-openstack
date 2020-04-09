// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package blockstorage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about an existing volume.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_volume_v3.html.markdown.
func GetVolumeV3(ctx *pulumi.Context, args *GetVolumeV3Args, opts ...pulumi.InvokeOption) (*GetVolumeV3Result, error) {
	var rv GetVolumeV3Result
	err := ctx.Invoke("openstack:blockstorage/getVolumeV3:getVolumeV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolumeV3.
type GetVolumeV3Args struct {
	// Indicates if the volume is bootable.
	Bootable *string `pulumi:"bootable"`
	// Metadata key/value pairs associated with the volume.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the volume.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the volume.
	Status *string `pulumi:"status"`
	// The type of the volume.
	VolumeType *string `pulumi:"volumeType"`
}

// A collection of values returned by getVolumeV3.
type GetVolumeV3Result struct {
	// Indicates if the volume is bootable.
	Bootable string `pulumi:"bootable"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Indicates if the volume can be attached to more then one server.
	Multiattach bool `pulumi:"multiattach"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The size of the volume in GBs.
	Size int `pulumi:"size"`
	// The ID of the volume from which the current volume was created.
	SourceVolumeId string `pulumi:"sourceVolumeId"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
	// The type of the volume.
	VolumeType string `pulumi:"volumeType"`
}
