// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available Shared File System share network.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/sharedfilesystem"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "sharenetwork_1"
// 		_, err := sharedfilesystem.LookupShareNetwork(ctx, &sharedfilesystem.LookupShareNetworkArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupShareNetwork(ctx *pulumi.Context, args *LookupShareNetworkArgs, opts ...pulumi.InvokeOption) (*LookupShareNetworkResult, error) {
	var rv LookupShareNetworkResult
	err := ctx.Invoke("openstack:sharedfilesystem/getShareNetwork:getShareNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getShareNetwork.
type LookupShareNetworkArgs struct {
	// The human-readable description of the share network.
	Description *string `pulumi:"description"`
	// The IP version of the share network. Can either be 4 or 6.
	IpVersion *int `pulumi:"ipVersion"`
	// The name of the share network.
	Name *string `pulumi:"name"`
	// The share network type. Can either be VLAN, VXLAN,
	// GRE, or flat.
	NetworkType *string `pulumi:"networkType"`
	// The neutron network UUID of the share network.
	NeutronNetId *string `pulumi:"neutronNetId"`
	// The neutron subnet UUID of the share network.
	NeutronSubnetId *string `pulumi:"neutronSubnetId"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to read a share network. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The security service IDs associated with
	// the share network.
	SecurityServiceId *string `pulumi:"securityServiceId"`
	// The share network segmentation ID.
	SegmentationId *int `pulumi:"segmentationId"`
}

// A collection of values returned by getShareNetwork.
type LookupShareNetworkResult struct {
	// See Argument Reference above.
	Cidr string `pulumi:"cidr"`
	// See Argument Reference above.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	IpVersion int `pulumi:"ipVersion"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	NetworkType string `pulumi:"networkType"`
	// See Argument Reference above.
	NeutronNetId string `pulumi:"neutronNetId"`
	// See Argument Reference above.
	NeutronSubnetId string `pulumi:"neutronSubnetId"`
	// The owner of the Share Network.
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	SecurityServiceId *string `pulumi:"securityServiceId"`
	// The list of security service IDs associated with
	// the share network.
	SecurityServiceIds []string `pulumi:"securityServiceIds"`
	// See Argument Reference above.
	SegmentationId int `pulumi:"segmentationId"`
}
