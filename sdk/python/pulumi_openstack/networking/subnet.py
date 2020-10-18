# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Subnet']


class Subnet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolArgs']]]]] = None,
                 allocation_pools_collection: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolsCollectionArgs']]]]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_dhcp: Optional[pulumi.Input[bool]] = None,
                 gateway_ip: Optional[pulumi.Input[str]] = None,
                 host_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetHostRouteArgs']]]]] = None,
                 ip_version: Optional[pulumi.Input[int]] = None,
                 ipv6_address_mode: Optional[pulumi.Input[str]] = None,
                 ipv6_ra_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 no_gateway: Optional[pulumi.Input[bool]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnetpool_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V2 Neutron subnet resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            network_id=network1.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolArgs']]]] allocation_pools: A block declaring the start and end range of
               the IP addresses available for use with DHCP in this subnet. Multiple
               `allocation_pool` blocks can be declared, providing the subnet with more
               than one range of IP addresses to use with DHCP. However, each IP range
               must be from the same CIDR that the subnet is part of.
               The `allocation_pool` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolsCollectionArgs']]]] allocation_pools_collection: A block declaring the start and end range of the IP addresses available for
               use with DHCP in this subnet.
               The `allocation_pools` block is documented below.
        :param pulumi.Input[str] cidr: CIDR representing IP range for this subnet, based on IP
               version. You can omit this option if you are creating a subnet from a
               subnet pool.
        :param pulumi.Input[str] description: Human-readable description of the subnet. Changing this
               updates the name of the existing subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_nameservers: An array of DNS name server names used by hosts
               in this subnet. Changing this updates the DNS name servers for the existing
               subnet.
        :param pulumi.Input[bool] enable_dhcp: The administrative state of the network.
               Acceptable values are "true" and "false". Changing this value enables or
               disables the DHCP capabilities of the existing subnet. Defaults to true.
        :param pulumi.Input[str] gateway_ip: Default gateway used by devices in this subnet.
               Leaving this blank and not setting `no_gateway` will cause a default
               gateway of `.1` to be used. Changing this updates the gateway IP of the
               existing subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetHostRouteArgs']]]] host_routes: (**Deprecated** - use `networking.SubnetRoute`
               instead) An array of routes that should be used by devices
               with IPs from this subnet (not including local subnet route). The host_route
               object structure is documented below. Changing this updates the host routes
               for the existing subnet.
        :param pulumi.Input[int] ip_version: IP version, either 4 (default) or 6. Changing this creates a
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
        :param pulumi.Input[int] prefix_length: The prefix length to use when creating a subnet
               from a subnet pool. The default subnet pool prefix length that was defined
               when creating the subnet pool will be used if not provided. Changing this
               creates a new subnet.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               subnet.
        :param pulumi.Input[str] subnetpool_id: The ID of the subnetpool associated with the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the subnet.
        :param pulumi.Input[str] tenant_id: The owner of the subnet. Required if admin wants to
               create a subnet for another tenant. Changing this creates a new subnet.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['allocation_pools'] = allocation_pools
            if allocation_pools_collection is not None:
                warnings.warn("use allocation_pool instead", DeprecationWarning)
                pulumi.log.warn("allocation_pools_collection is deprecated: use allocation_pool instead")
            __props__['allocation_pools_collection'] = allocation_pools_collection
            __props__['cidr'] = cidr
            __props__['description'] = description
            __props__['dns_nameservers'] = dns_nameservers
            __props__['enable_dhcp'] = enable_dhcp
            __props__['gateway_ip'] = gateway_ip
            if host_routes is not None:
                warnings.warn("Use openstack_networking_subnet_route_v2 instead", DeprecationWarning)
                pulumi.log.warn("host_routes is deprecated: Use openstack_networking_subnet_route_v2 instead")
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

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            allocation_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolArgs']]]]] = None,
            allocation_pools_collection: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolsCollectionArgs']]]]] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dns_nameservers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enable_dhcp: Optional[pulumi.Input[bool]] = None,
            gateway_ip: Optional[pulumi.Input[str]] = None,
            host_routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetHostRouteArgs']]]]] = None,
            ip_version: Optional[pulumi.Input[int]] = None,
            ipv6_address_mode: Optional[pulumi.Input[str]] = None,
            ipv6_ra_mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            no_gateway: Optional[pulumi.Input[bool]] = None,
            prefix_length: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            subnetpool_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Subnet':
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] all_tags: The collection of ags assigned on the subnet, which have been
               explicitly and implicitly added.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolArgs']]]] allocation_pools: A block declaring the start and end range of
               the IP addresses available for use with DHCP in this subnet. Multiple
               `allocation_pool` blocks can be declared, providing the subnet with more
               than one range of IP addresses to use with DHCP. However, each IP range
               must be from the same CIDR that the subnet is part of.
               The `allocation_pool` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetAllocationPoolsCollectionArgs']]]] allocation_pools_collection: A block declaring the start and end range of the IP addresses available for
               use with DHCP in this subnet.
               The `allocation_pools` block is documented below.
        :param pulumi.Input[str] cidr: CIDR representing IP range for this subnet, based on IP
               version. You can omit this option if you are creating a subnet from a
               subnet pool.
        :param pulumi.Input[str] description: Human-readable description of the subnet. Changing this
               updates the name of the existing subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_nameservers: An array of DNS name server names used by hosts
               in this subnet. Changing this updates the DNS name servers for the existing
               subnet.
        :param pulumi.Input[bool] enable_dhcp: The administrative state of the network.
               Acceptable values are "true" and "false". Changing this value enables or
               disables the DHCP capabilities of the existing subnet. Defaults to true.
        :param pulumi.Input[str] gateway_ip: Default gateway used by devices in this subnet.
               Leaving this blank and not setting `no_gateway` will cause a default
               gateway of `.1` to be used. Changing this updates the gateway IP of the
               existing subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SubnetHostRouteArgs']]]] host_routes: (**Deprecated** - use `networking.SubnetRoute`
               instead) An array of routes that should be used by devices
               with IPs from this subnet (not including local subnet route). The host_route
               object structure is documented below. Changing this updates the host routes
               for the existing subnet.
        :param pulumi.Input[int] ip_version: IP version, either 4 (default) or 6. Changing this creates a
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
        :param pulumi.Input[int] prefix_length: The prefix length to use when creating a subnet
               from a subnet pool. The default subnet pool prefix length that was defined
               when creating the subnet pool will be used if not provided. Changing this
               creates a new subnet.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               subnet.
        :param pulumi.Input[str] subnetpool_id: The ID of the subnetpool associated with the subnet.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the subnet.
        :param pulumi.Input[str] tenant_id: The owner of the subnet. Required if admin wants to
               create a subnet for another tenant. Changing this creates a new subnet.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["all_tags"] = all_tags
        __props__["allocation_pools"] = allocation_pools
        __props__["allocation_pools_collection"] = allocation_pools_collection
        __props__["cidr"] = cidr
        __props__["description"] = description
        __props__["dns_nameservers"] = dns_nameservers
        __props__["enable_dhcp"] = enable_dhcp
        __props__["gateway_ip"] = gateway_ip
        __props__["host_routes"] = host_routes
        __props__["ip_version"] = ip_version
        __props__["ipv6_address_mode"] = ipv6_address_mode
        __props__["ipv6_ra_mode"] = ipv6_ra_mode
        __props__["name"] = name
        __props__["network_id"] = network_id
        __props__["no_gateway"] = no_gateway
        __props__["prefix_length"] = prefix_length
        __props__["region"] = region
        __props__["subnetpool_id"] = subnetpool_id
        __props__["tags"] = tags
        __props__["tenant_id"] = tenant_id
        __props__["value_specs"] = value_specs
        return Subnet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The collection of ags assigned on the subnet, which have been
        explicitly and implicitly added.
        """
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="allocationPools")
    def allocation_pools(self) -> pulumi.Output[Sequence['outputs.SubnetAllocationPool']]:
        """
        A block declaring the start and end range of
        the IP addresses available for use with DHCP in this subnet. Multiple
        `allocation_pool` blocks can be declared, providing the subnet with more
        than one range of IP addresses to use with DHCP. However, each IP range
        must be from the same CIDR that the subnet is part of.
        The `allocation_pool` block is documented below.
        """
        return pulumi.get(self, "allocation_pools")

    @property
    @pulumi.getter(name="allocationPoolsCollection")
    def allocation_pools_collection(self) -> pulumi.Output[Sequence['outputs.SubnetAllocationPoolsCollection']]:
        """
        A block declaring the start and end range of the IP addresses available for
        use with DHCP in this subnet.
        The `allocation_pools` block is documented below.
        """
        return pulumi.get(self, "allocation_pools_collection")

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        CIDR representing IP range for this subnet, based on IP
        version. You can omit this option if you are creating a subnet from a
        subnet pool.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-readable description of the subnet. Changing this
        updates the name of the existing subnet.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsNameservers")
    def dns_nameservers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of DNS name server names used by hosts
        in this subnet. Changing this updates the DNS name servers for the existing
        subnet.
        """
        return pulumi.get(self, "dns_nameservers")

    @property
    @pulumi.getter(name="enableDhcp")
    def enable_dhcp(self) -> pulumi.Output[Optional[bool]]:
        """
        The administrative state of the network.
        Acceptable values are "true" and "false". Changing this value enables or
        disables the DHCP capabilities of the existing subnet. Defaults to true.
        """
        return pulumi.get(self, "enable_dhcp")

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> pulumi.Output[str]:
        """
        Default gateway used by devices in this subnet.
        Leaving this blank and not setting `no_gateway` will cause a default
        gateway of `.1` to be used. Changing this updates the gateway IP of the
        existing subnet.
        """
        return pulumi.get(self, "gateway_ip")

    @property
    @pulumi.getter(name="hostRoutes")
    def host_routes(self) -> pulumi.Output[Optional[Sequence['outputs.SubnetHostRoute']]]:
        """
        (**Deprecated** - use `networking.SubnetRoute`
        instead) An array of routes that should be used by devices
        with IPs from this subnet (not including local subnet route). The host_route
        object structure is documented below. Changing this updates the host routes
        for the existing subnet.
        """
        return pulumi.get(self, "host_routes")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[Optional[int]]:
        """
        IP version, either 4 (default) or 6. Changing this creates a
        new subnet.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="ipv6AddressMode")
    def ipv6_address_mode(self) -> pulumi.Output[str]:
        """
        The IPv6 address mode. Valid values are
        `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        """
        return pulumi.get(self, "ipv6_address_mode")

    @property
    @pulumi.getter(name="ipv6RaMode")
    def ipv6_ra_mode(self) -> pulumi.Output[str]:
        """
        The IPv6 Router Advertisement mode. Valid values
        are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
        """
        return pulumi.get(self, "ipv6_ra_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the subnet. Changing this updates the name of
        the existing subnet.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The UUID of the parent network. Changing this
        creates a new subnet.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="noGateway")
    def no_gateway(self) -> pulumi.Output[Optional[bool]]:
        """
        Do not set a gateway IP on this subnet. Changing
        this removes or adds a default gateway IP of the existing subnet.
        """
        return pulumi.get(self, "no_gateway")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Output[Optional[int]]:
        """
        The prefix length to use when creating a subnet
        from a subnet pool. The default subnet pool prefix length that was defined
        when creating the subnet pool will be used if not provided. Changing this
        creates a new subnet.
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron subnet. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        subnet.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subnetpoolId")
    def subnetpool_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the subnetpool associated with the subnet.
        """
        return pulumi.get(self, "subnetpool_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of string tags for the subnet.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the subnet. Required if admin wants to
        create a subnet for another tenant. Changing this creates a new subnet.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

