// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networking

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available OpenStack port.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_port_v2.html.markdown.
func LookupPort(ctx *pulumi.Context, args *GetPortArgs) (*GetPortResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["adminStateUp"] = args.AdminStateUp
		inputs["description"] = args.Description
		inputs["deviceId"] = args.DeviceId
		inputs["deviceOwner"] = args.DeviceOwner
		inputs["dnsName"] = args.DnsName
		inputs["fixedIp"] = args.FixedIp
		inputs["macAddress"] = args.MacAddress
		inputs["name"] = args.Name
		inputs["networkId"] = args.NetworkId
		inputs["portId"] = args.PortId
		inputs["projectId"] = args.ProjectId
		inputs["region"] = args.Region
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["status"] = args.Status
		inputs["tags"] = args.Tags
		inputs["tenantId"] = args.TenantId
	}
	outputs, err := ctx.Invoke("openstack:networking/getPort:getPort", inputs)
	if err != nil {
		return nil, err
	}
	return &GetPortResult{
		AdminStateUp: outputs["adminStateUp"],
		AllFixedIps: outputs["allFixedIps"],
		AllSecurityGroupIds: outputs["allSecurityGroupIds"],
		AllTags: outputs["allTags"],
		AllowedAddressPairs: outputs["allowedAddressPairs"],
		Bindings: outputs["bindings"],
		Description: outputs["description"],
		DeviceId: outputs["deviceId"],
		DeviceOwner: outputs["deviceOwner"],
		DnsAssignments: outputs["dnsAssignments"],
		DnsName: outputs["dnsName"],
		ExtraDhcpOptions: outputs["extraDhcpOptions"],
		FixedIp: outputs["fixedIp"],
		MacAddress: outputs["macAddress"],
		Name: outputs["name"],
		NetworkId: outputs["networkId"],
		PortId: outputs["portId"],
		ProjectId: outputs["projectId"],
		Region: outputs["region"],
		SecurityGroupIds: outputs["securityGroupIds"],
		Status: outputs["status"],
		Tags: outputs["tags"],
		TenantId: outputs["tenantId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getPort.
type GetPortArgs struct {
	// The administrative state of the port.
	AdminStateUp interface{}
	// Human-readable description of the port.
	Description interface{}
	// The ID of the device the port belongs to.
	DeviceId interface{}
	// The device owner of the port.
	DeviceOwner interface{}
	// The port DNS name to filter. Available, when Neutron
	// DNS extension is enabled.
	DnsName interface{}
	// The port IP address filter.
	FixedIp interface{}
	// The MAC address of the port.
	MacAddress interface{}
	// The name of the port.
	Name interface{}
	// The ID of the network the port belongs to.
	NetworkId interface{}
	// The ID of the port.
	PortId interface{}
	// The owner of the port.
	ProjectId interface{}
	// The region in which to obtain the V2 Neutron client.
	// A Neutron client is needed to retrieve port ids. If omitted, the
	// `region` argument of the provider is used.
	Region interface{}
	// The list of port security group IDs to filter.
	SecurityGroupIds interface{}
	// The status of the port.
	Status interface{}
	// The list of port tags to filter.
	Tags interface{}
	TenantId interface{}
}

// A collection of values returned by getPort.
type GetPortResult struct {
	// See Argument Reference above.
	AdminStateUp interface{}
	// The collection of Fixed IP addresses on the port in the
	// order returned by the Network v2 API.
	AllFixedIps interface{}
	// The set of security group IDs applied on the port.
	AllSecurityGroupIds interface{}
	// The set of string tags applied on the port.
	AllTags interface{}
	// An IP/MAC Address pair of additional IP
	// addresses that can be active on this port. The structure is described
	// below.
	AllowedAddressPairs interface{}
	// The port binding information. The structure is described below.
	Bindings interface{}
	// See Argument Reference above.
	Description interface{}
	// See Argument Reference above.
	DeviceId interface{}
	// See Argument Reference above.
	DeviceOwner interface{}
	// The list of maps representing port DNS assignments.
	DnsAssignments interface{}
	// See Argument Reference above.
	DnsName interface{}
	// An extra DHCP option configured on the port.
	// The structure is described below.
	ExtraDhcpOptions interface{}
	FixedIp interface{}
	// The additional MAC address.
	MacAddress interface{}
	// Name of the DHCP option.
	Name interface{}
	// See Argument Reference above.
	NetworkId interface{}
	// See Argument Reference above.
	PortId interface{}
	// See Argument Reference above.
	ProjectId interface{}
	// See Argument Reference above.
	Region interface{}
	SecurityGroupIds interface{}
	Status interface{}
	Tags interface{}
	TenantId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
