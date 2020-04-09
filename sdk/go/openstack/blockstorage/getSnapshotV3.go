// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package blockstorage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about an existing snapshot.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_snapshot_v3.html.markdown.
func GetSnapshotV3(ctx *pulumi.Context, args *GetSnapshotV3Args, opts ...pulumi.InvokeOption) (*GetSnapshotV3Result, error) {
	var rv GetSnapshotV3Result
	err := ctx.Invoke("openstack:blockstorage/getSnapshotV3:getSnapshotV3", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshotV3.
type GetSnapshotV3Args struct {
	// Pick the most recently created snapshot if there
	// are multiple results.
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the snapshot.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Block Storage
	// client. If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the snapshot.
	Status *string `pulumi:"status"`
	// The ID of the snapshot's volume.
	VolumeId *string `pulumi:"volumeId"`
}

// A collection of values returned by getSnapshotV3.
type GetSnapshotV3Result struct {
	// The snapshot's description.
	Description string `pulumi:"description"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The snapshot's metadata.
	Metadata   map[string]interface{} `pulumi:"metadata"`
	MostRecent *bool                  `pulumi:"mostRecent"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// The size of the snapshot.
	Size int `pulumi:"size"`
	// See Argument Reference above.
	Status string `pulumi:"status"`
	// See Argument Reference above.
	VolumeId string `pulumi:"volumeId"`
}
