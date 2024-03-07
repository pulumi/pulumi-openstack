# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSubnetIdsV2Result',
    'AwaitableGetSubnetIdsV2Result',
    'get_subnet_ids_v2',
    'get_subnet_ids_v2_output',
]

@pulumi.output_type
class GetSubnetIdsV2Result:
    """
    A collection of values returned by getSubnetIdsV2.
    """
    def __init__(__self__, cidr=None, description=None, dhcp_enabled=None, gateway_ip=None, id=None, ids=None, ip_version=None, ipv6_address_mode=None, ipv6_ra_mode=None, name=None, name_regex=None, network_id=None, region=None, sort_direction=None, sort_key=None, subnetpool_id=None, tags=None, tenant_id=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dhcp_enabled and not isinstance(dhcp_enabled, bool):
            raise TypeError("Expected argument 'dhcp_enabled' to be a bool")
        pulumi.set(__self__, "dhcp_enabled", dhcp_enabled)
        if gateway_ip and not isinstance(gateway_ip, str):
            raise TypeError("Expected argument 'gateway_ip' to be a str")
        pulumi.set(__self__, "gateway_ip", gateway_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if ip_version and not isinstance(ip_version, int):
            raise TypeError("Expected argument 'ip_version' to be a int")
        pulumi.set(__self__, "ip_version", ip_version)
        if ipv6_address_mode and not isinstance(ipv6_address_mode, str):
            raise TypeError("Expected argument 'ipv6_address_mode' to be a str")
        pulumi.set(__self__, "ipv6_address_mode", ipv6_address_mode)
        if ipv6_ra_mode and not isinstance(ipv6_ra_mode, str):
            raise TypeError("Expected argument 'ipv6_ra_mode' to be a str")
        pulumi.set(__self__, "ipv6_ra_mode", ipv6_ra_mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        pulumi.set(__self__, "sort_direction", sort_direction)
        if sort_key and not isinstance(sort_key, str):
            raise TypeError("Expected argument 'sort_key' to be a str")
        pulumi.set(__self__, "sort_key", sort_key)
        if subnetpool_id and not isinstance(subnetpool_id, str):
            raise TypeError("Expected argument 'subnetpool_id' to be a str")
        pulumi.set(__self__, "subnetpool_id", subnetpool_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[str]:
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dhcpEnabled")
    def dhcp_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "dhcp_enabled")

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> Optional[str]:
        return pulumi.get(self, "gateway_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[int]:
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="ipv6AddressMode")
    def ipv6_address_mode(self) -> Optional[str]:
        return pulumi.get(self, "ipv6_address_mode")

    @property
    @pulumi.getter(name="ipv6RaMode")
    def ipv6_ra_mode(self) -> str:
        return pulumi.get(self, "ipv6_ra_mode")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[str]:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sortDirection")
    def sort_direction(self) -> Optional[str]:
        return pulumi.get(self, "sort_direction")

    @property
    @pulumi.getter(name="sortKey")
    def sort_key(self) -> Optional[str]:
        return pulumi.get(self, "sort_key")

    @property
    @pulumi.getter(name="subnetpoolId")
    def subnetpool_id(self) -> Optional[str]:
        return pulumi.get(self, "subnetpool_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


class AwaitableGetSubnetIdsV2Result(GetSubnetIdsV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetIdsV2Result(
            cidr=self.cidr,
            description=self.description,
            dhcp_enabled=self.dhcp_enabled,
            gateway_ip=self.gateway_ip,
            id=self.id,
            ids=self.ids,
            ip_version=self.ip_version,
            ipv6_address_mode=self.ipv6_address_mode,
            ipv6_ra_mode=self.ipv6_ra_mode,
            name=self.name,
            name_regex=self.name_regex,
            network_id=self.network_id,
            region=self.region,
            sort_direction=self.sort_direction,
            sort_key=self.sort_key,
            subnetpool_id=self.subnetpool_id,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_subnet_ids_v2(cidr: Optional[str] = None,
                      description: Optional[str] = None,
                      dhcp_enabled: Optional[bool] = None,
                      gateway_ip: Optional[str] = None,
                      ip_version: Optional[int] = None,
                      ipv6_address_mode: Optional[str] = None,
                      ipv6_ra_mode: Optional[str] = None,
                      name: Optional[str] = None,
                      name_regex: Optional[str] = None,
                      network_id: Optional[str] = None,
                      region: Optional[str] = None,
                      sort_direction: Optional[str] = None,
                      sort_key: Optional[str] = None,
                      subnetpool_id: Optional[str] = None,
                      tags: Optional[Sequence[str]] = None,
                      tenant_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetIdsV2Result:
    """
    Use this data source to get a list of Openstack Subnet IDs matching the
    specified criteria.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_openstack as openstack

    subnets = openstack.networking.get_subnet_ids_v2(name_regex="public",
        tags=["public"])
    ```
    <!--End PulumiCodeChooser -->


    :param str cidr: The CIDR of the subnet.
    :param str description: Human-readable description of the subnet.
    :param bool dhcp_enabled: If the subnet has DHCP enabled.
    :param str gateway_ip: The IP of the subnet's gateway.
    :param int ip_version: The IP version of the subnet (either 4 or 6).
    :param str ipv6_address_mode: The IPv6 address mode. Valid values are
           `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
    :param str ipv6_ra_mode: The IPv6 Router Advertisement mode. Valid values
           are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
    :param str name: The name of the subnet.
    :param str network_id: The ID of the network the subnet belongs to.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve subnet ids. If omitted, the
           `region` argument of the provider is used.
    :param str sort_direction: Order the results in either `asc` or `desc`.
           Defaults to none.
    :param str sort_key: Sort subnets based on a certain key. Defaults to none.
    :param str subnetpool_id: The ID of the subnetpool associated with the subnet.
    :param Sequence[str] tags: The list of subnet tags to filter.
    :param str tenant_id: The owner of the subnet.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['description'] = description
    __args__['dhcpEnabled'] = dhcp_enabled
    __args__['gatewayIp'] = gateway_ip
    __args__['ipVersion'] = ip_version
    __args__['ipv6AddressMode'] = ipv6_address_mode
    __args__['ipv6RaMode'] = ipv6_ra_mode
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['networkId'] = network_id
    __args__['region'] = region
    __args__['sortDirection'] = sort_direction
    __args__['sortKey'] = sort_key
    __args__['subnetpoolId'] = subnetpool_id
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getSubnetIdsV2:getSubnetIdsV2', __args__, opts=opts, typ=GetSubnetIdsV2Result).value

    return AwaitableGetSubnetIdsV2Result(
        cidr=pulumi.get(__ret__, 'cidr'),
        description=pulumi.get(__ret__, 'description'),
        dhcp_enabled=pulumi.get(__ret__, 'dhcp_enabled'),
        gateway_ip=pulumi.get(__ret__, 'gateway_ip'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        ipv6_address_mode=pulumi.get(__ret__, 'ipv6_address_mode'),
        ipv6_ra_mode=pulumi.get(__ret__, 'ipv6_ra_mode'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        network_id=pulumi.get(__ret__, 'network_id'),
        region=pulumi.get(__ret__, 'region'),
        sort_direction=pulumi.get(__ret__, 'sort_direction'),
        sort_key=pulumi.get(__ret__, 'sort_key'),
        subnetpool_id=pulumi.get(__ret__, 'subnetpool_id'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_subnet_ids_v2)
def get_subnet_ids_v2_output(cidr: Optional[pulumi.Input[Optional[str]]] = None,
                             description: Optional[pulumi.Input[Optional[str]]] = None,
                             dhcp_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                             gateway_ip: Optional[pulumi.Input[Optional[str]]] = None,
                             ip_version: Optional[pulumi.Input[Optional[int]]] = None,
                             ipv6_address_mode: Optional[pulumi.Input[Optional[str]]] = None,
                             ipv6_ra_mode: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             network_id: Optional[pulumi.Input[Optional[str]]] = None,
                             region: Optional[pulumi.Input[Optional[str]]] = None,
                             sort_direction: Optional[pulumi.Input[Optional[str]]] = None,
                             sort_key: Optional[pulumi.Input[Optional[str]]] = None,
                             subnetpool_id: Optional[pulumi.Input[Optional[str]]] = None,
                             tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetIdsV2Result]:
    """
    Use this data source to get a list of Openstack Subnet IDs matching the
    specified criteria.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_openstack as openstack

    subnets = openstack.networking.get_subnet_ids_v2(name_regex="public",
        tags=["public"])
    ```
    <!--End PulumiCodeChooser -->


    :param str cidr: The CIDR of the subnet.
    :param str description: Human-readable description of the subnet.
    :param bool dhcp_enabled: If the subnet has DHCP enabled.
    :param str gateway_ip: The IP of the subnet's gateway.
    :param int ip_version: The IP version of the subnet (either 4 or 6).
    :param str ipv6_address_mode: The IPv6 address mode. Valid values are
           `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
    :param str ipv6_ra_mode: The IPv6 Router Advertisement mode. Valid values
           are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
    :param str name: The name of the subnet.
    :param str network_id: The ID of the network the subnet belongs to.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve subnet ids. If omitted, the
           `region` argument of the provider is used.
    :param str sort_direction: Order the results in either `asc` or `desc`.
           Defaults to none.
    :param str sort_key: Sort subnets based on a certain key. Defaults to none.
    :param str subnetpool_id: The ID of the subnetpool associated with the subnet.
    :param Sequence[str] tags: The list of subnet tags to filter.
    :param str tenant_id: The owner of the subnet.
    """
    ...
