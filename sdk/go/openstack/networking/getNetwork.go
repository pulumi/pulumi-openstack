// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack network.
func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	var rv LookupNetworkResult
	err := ctx.Invoke("openstack:networking/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type LookupNetworkArgs struct {
	// Human-readable description of the network.
	Description *string `pulumi:"description"`
	// The external routing facility of the network.
	External *bool `pulumi:"external"`
	// The CIDR of a subnet within the network.
	MatchingSubnetCidr *string `pulumi:"matchingSubnetCidr"`
	// The network MTU to filter. Available, when Neutron `net-mtu`
	// extension is enabled.
	Mtu *int `pulumi:"mtu"`
	// The name of the network.
	Name *string `pulumi:"name"`
	// The ID of the network.
	NetworkId *string `pulumi:"networkId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve networks ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The status of the network.
	Status *string `pulumi:"status"`
	// The list of network tags to filter.
	Tags []string `pulumi:"tags"`
	// The owner of the network.
	TenantId *string `pulumi:"tenantId"`
	// The VLAN transparent attribute for the
	// network.
	TransparentVlan *bool `pulumi:"transparentVlan"`
}

// A collection of values returned by getNetwork.
type LookupNetworkResult struct {
	// The administrative state of the network.
	AdminStateUp string `pulumi:"adminStateUp"`
	// The set of string tags applied on the network.
	AllTags []string `pulumi:"allTags"`
	// The availability zone candidates for the network.
	AvailabilityZoneHints []string `pulumi:"availabilityZoneHints"`
	// See Argument Reference above.
	Description *string `pulumi:"description"`
	// The network DNS domain. Available, when Neutron DNS extension
	// is enabled
	DnsDomain string `pulumi:"dnsDomain"`
	// See Argument Reference above.
	External *bool `pulumi:"external"`
	// id is the provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	MatchingSubnetCidr *string `pulumi:"matchingSubnetCidr"`
	// See Argument Reference above.
	Mtu *int `pulumi:"mtu"`
	// See Argument Reference above.
	Name      *string `pulumi:"name"`
	NetworkId *string `pulumi:"networkId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// Specifies whether the network resource can be accessed by any
	// tenant or not.
	Shared   string   `pulumi:"shared"`
	Status   *string  `pulumi:"status"`
	Tags     []string `pulumi:"tags"`
	TenantId *string  `pulumi:"tenantId"`
	// See Argument Reference above.
	TransparentVlan *bool `pulumi:"transparentVlan"`
}
