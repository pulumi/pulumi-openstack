// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack subnet.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_subnet_v2.html.markdown.
func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("openstack:networking/getSubnet:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnet.
type LookupSubnetArgs struct {
	// The CIDR of the subnet.
	Cidr *string `pulumi:"cidr"`
	// Human-readable description for the subnet.
	Description *string `pulumi:"description"`
	// If the subnet has DHCP disabled.
	DhcpDisabled *bool `pulumi:"dhcpDisabled"`
	// If the subnet has DHCP enabled.
	DhcpEnabled *bool `pulumi:"dhcpEnabled"`
	// The IP of the subnet's gateway.
	GatewayIp *string `pulumi:"gatewayIp"`
	// The IP version of the subnet (either 4 or 6).
	IpVersion *int `pulumi:"ipVersion"`
	// The IPv6 address mode. Valid values are
	// `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6AddressMode *string `pulumi:"ipv6AddressMode"`
	// The IPv6 Router Advertisement mode. Valid values
	// are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
	Ipv6RaMode *string `pulumi:"ipv6RaMode"`
	// The name of the subnet.
	Name *string `pulumi:"name"`
	// The ID of the network the subnet belongs to.
	NetworkId *string `pulumi:"networkId"`
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve subnet ids. If omitted, the
	// `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The ID of the subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The ID of the subnetpool associated with the subnet.
	SubnetpoolId *string `pulumi:"subnetpoolId"`
	// The list of subnet tags to filter.
	Tags []string `pulumi:"tags"`
	// The owner of the subnet.
	TenantId *string `pulumi:"tenantId"`
}


// A collection of values returned by getSubnet.
type LookupSubnetResult struct {
	// A set of string tags applied on the subnet.
	AllTags []string `pulumi:"allTags"`
	// Allocation pools of the subnet.
	AllocationPools []GetSubnetAllocationPool `pulumi:"allocationPools"`
	Cidr string `pulumi:"cidr"`
	Description string `pulumi:"description"`
	DhcpDisabled *bool `pulumi:"dhcpDisabled"`
	DhcpEnabled *bool `pulumi:"dhcpEnabled"`
	// DNS Nameservers of the subnet.
	DnsNameservers []string `pulumi:"dnsNameservers"`
	// Whether the subnet has DHCP enabled or not.
	EnableDhcp bool `pulumi:"enableDhcp"`
	GatewayIp string `pulumi:"gatewayIp"`
	// Host Routes of the subnet.
	HostRoutes []GetSubnetHostRoute `pulumi:"hostRoutes"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	IpVersion int `pulumi:"ipVersion"`
	Ipv6AddressMode string `pulumi:"ipv6AddressMode"`
	Ipv6RaMode string `pulumi:"ipv6RaMode"`
	Name string `pulumi:"name"`
	NetworkId string `pulumi:"networkId"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	SubnetId string `pulumi:"subnetId"`
	SubnetpoolId string `pulumi:"subnetpoolId"`
	Tags []string `pulumi:"tags"`
	TenantId string `pulumi:"tenantId"`
}

