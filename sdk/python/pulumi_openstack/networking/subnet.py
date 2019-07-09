# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Subnet(pulumi.CustomResource):
    all_tags: pulumi.Output[list]
    """
    The collection of ags assigned on the subnet, which have been
    explicitly and implicitly added.
    """
    allocation_pools: pulumi.Output[list]
    """
    An array of sub-ranges of CIDR available for
    dynamic allocation to ports. The allocation_pool object structure is
    documented below.
    """
    allocation_pools_collection: pulumi.Output[list]
    """
    
    An array of sub-ranges of CIDR available for dynamic allocation to ports.
    The allocation_pools object structure is documented below.
    """
    cidr: pulumi.Output[str]
    """
    CIDR representing IP range for this subnet, based on IP
    version. You can omit this option if you are creating a subnet from a
    subnet pool.
    """
    description: pulumi.Output[str]
    """
    Human-readable description of the subnet. Changing this
    updates the name of the existing subnet.
    """
    dns_nameservers: pulumi.Output[list]
    """
    An array of DNS name server names used by hosts
    in this subnet. Changing this updates the DNS name servers for the existing
    subnet.
    """
    enable_dhcp: pulumi.Output[bool]
    """
    The administrative state of the network.
    Acceptable values are "true" and "false". Changing this value enables or
    disables the DHCP capabilities of the existing subnet. Defaults to true.
    """
    gateway_ip: pulumi.Output[str]
    """
    Default gateway used by devices in this subnet.
    Leaving this blank and not setting `no_gateway` will cause a default
    gateway of `.1` to be used. Changing this updates the gateway IP of the
    existing subnet.
    """
    host_routes: pulumi.Output[list]
    """
    (**Deprecated** - use `openstack_networking_subnet_route_v2`
    instead) An array of routes that should be used by devices
    with IPs from this subnet (not including local subnet route). The host_route
    object structure is documented below. Changing this updates the host routes
    for the existing subnet.
    """
    ip_version: pulumi.Output[float]
    """
    IP version, either 4 (default) or 6. Changing this creates a
    new subnet.
    """
    ipv6_address_mode: pulumi.Output[str]
    """
    The IPv6 address mode. Valid values are
    `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
    """
    ipv6_ra_mode: pulumi.Output[str]
    """
    The IPv6 Router Advertisement mode. Valid values
    are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
    """
    name: pulumi.Output[str]
    """
    The name of the subnet. Changing this updates the name of
    the existing subnet.
    """
    network_id: pulumi.Output[str]
    """
    The UUID of the parent network. Changing this
    creates a new subnet.
    """
    no_gateway: pulumi.Output[bool]
    """
    Do not set a gateway IP on this subnet. Changing
    this removes or adds a default gateway IP of the existing subnet.
    """
    prefix_length: pulumi.Output[float]
    """
    The prefix length to use when creating a subnet
    from a subnet pool. The default subnet pool prefix length that was defined
    when creating the subnet pool will be used if not provided. Changing this
    creates a new subnet.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron subnet. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    subnet.
    """
    subnetpool_id: pulumi.Output[str]
    """
    The ID of the subnetpool associated with the subnet.
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the subnet.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the subnet. Required if admin wants to
    create a subnet for another tenant. Changing this creates a new subnet.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, allocation_pools=None, allocation_pools_collection=None, cidr=None, description=None, dns_nameservers=None, enable_dhcp=None, gateway_ip=None, host_routes=None, ip_version=None, ipv6_address_mode=None, ipv6_ra_mode=None, name=None, network_id=None, no_gateway=None, prefix_length=None, region=None, subnetpool_id=None, tags=None, tenant_id=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron subnet resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allocation_pools: An array of sub-ranges of CIDR available for
               dynamic allocation to ports. The allocation_pool object structure is
               documented below.
        :param pulumi.Input[list] allocation_pools_collection: 
               An array of sub-ranges of CIDR available for dynamic allocation to ports.
               The allocation_pools object structure is documented below.
        :param pulumi.Input[str] cidr: CIDR representing IP range for this subnet, based on IP
               version. You can omit this option if you are creating a subnet from a
               subnet pool.
        :param pulumi.Input[str] description: Human-readable description of the subnet. Changing this
               updates the name of the existing subnet.
        :param pulumi.Input[list] dns_nameservers: An array of DNS name server names used by hosts
               in this subnet. Changing this updates the DNS name servers for the existing
               subnet.
        :param pulumi.Input[bool] enable_dhcp: The administrative state of the network.
               Acceptable values are "true" and "false". Changing this value enables or
               disables the DHCP capabilities of the existing subnet. Defaults to true.
        :param pulumi.Input[str] gateway_ip: Default gateway used by devices in this subnet.
               Leaving this blank and not setting `no_gateway` will cause a default
               gateway of `.1` to be used. Changing this updates the gateway IP of the
               existing subnet.
        :param pulumi.Input[list] host_routes: (**Deprecated** - use `openstack_networking_subnet_route_v2`
               instead) An array of routes that should be used by devices
               with IPs from this subnet (not including local subnet route). The host_route
               object structure is documented below. Changing this updates the host routes
               for the existing subnet.
        :param pulumi.Input[float] ip_version: IP version, either 4 (default) or 6. Changing this creates a
               new subnet.
        :param pulumi.Input[str] ipv6_address_mode: The IPv6 address mode. Valid values are
               `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        :param pulumi.Input[str] ipv6_ra_mode: The IPv6 Router Advertisement mode. Valid values
               are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        :param pulumi.Input[str] name: The name of the subnet. Changing this updates the name of
               the existing subnet.
        :param pulumi.Input[str] network_id: The UUID of the parent network. Changing this
               creates a new subnet.
        :param pulumi.Input[bool] no_gateway: Do not set a gateway IP on this subnet. Changing
               this removes or adds a default gateway IP of the existing subnet.
        :param pulumi.Input[float] prefix_length: The prefix length to use when creating a subnet
               from a subnet pool. The default subnet pool prefix length that was defined
               when creating the subnet pool will be used if not provided. Changing this
               creates a new subnet.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               subnet.
        :param pulumi.Input[str] subnetpool_id: The ID of the subnetpool associated with the subnet.
        :param pulumi.Input[list] tags: A set of string tags for the subnet.
        :param pulumi.Input[str] tenant_id: The owner of the subnet. Required if admin wants to
               create a subnet for another tenant. Changing this creates a new subnet.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_subnet_v2.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allocation_pools'] = allocation_pools

        __props__['allocation_pools_collection'] = allocation_pools_collection

        __props__['cidr'] = cidr

        __props__['description'] = description

        __props__['dns_nameservers'] = dns_nameservers

        __props__['enable_dhcp'] = enable_dhcp

        __props__['gateway_ip'] = gateway_ip

        __props__['host_routes'] = host_routes

        __props__['ip_version'] = ip_version

        __props__['ipv6_address_mode'] = ipv6_address_mode

        __props__['ipv6_ra_mode'] = ipv6_ra_mode

        __props__['name'] = name

        if network_id is None:
            raise TypeError("Missing required property 'network_id'")
        __props__['network_id'] = network_id

        __props__['no_gateway'] = no_gateway

        __props__['prefix_length'] = prefix_length

        __props__['region'] = region

        __props__['subnetpool_id'] = subnetpool_id

        __props__['tags'] = tags

        __props__['tenant_id'] = tenant_id

        __props__['value_specs'] = value_specs

        __props__['all_tags'] = None

        super(Subnet, __self__).__init__(
            'openstack:networking/subnet:Subnet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

