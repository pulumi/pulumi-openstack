# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetSubnetResult',
    'AwaitableGetSubnetResult',
    'get_subnet',
    'get_subnet_output',
]

@pulumi.output_type
class GetSubnetResult:
    """
    A collection of values returned by getSubnet.
    """
    def __init__(__self__, all_tags=None, allocation_pools=None, cidr=None, description=None, dhcp_enabled=None, dns_nameservers=None, dns_publish_fixed_ip=None, enable_dhcp=None, gateway_ip=None, host_routes=None, id=None, ip_version=None, ipv6_address_mode=None, ipv6_ra_mode=None, name=None, network_id=None, region=None, service_types=None, subnet_id=None, subnetpool_id=None, tags=None, tenant_id=None):
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if allocation_pools and not isinstance(allocation_pools, list):
            raise TypeError("Expected argument 'allocation_pools' to be a list")
        pulumi.set(__self__, "allocation_pools", allocation_pools)
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dhcp_enabled and not isinstance(dhcp_enabled, bool):
            raise TypeError("Expected argument 'dhcp_enabled' to be a bool")
        pulumi.set(__self__, "dhcp_enabled", dhcp_enabled)
        if dns_nameservers and not isinstance(dns_nameservers, list):
            raise TypeError("Expected argument 'dns_nameservers' to be a list")
        pulumi.set(__self__, "dns_nameservers", dns_nameservers)
        if dns_publish_fixed_ip and not isinstance(dns_publish_fixed_ip, bool):
            raise TypeError("Expected argument 'dns_publish_fixed_ip' to be a bool")
        pulumi.set(__self__, "dns_publish_fixed_ip", dns_publish_fixed_ip)
        if enable_dhcp and not isinstance(enable_dhcp, bool):
            raise TypeError("Expected argument 'enable_dhcp' to be a bool")
        pulumi.set(__self__, "enable_dhcp", enable_dhcp)
        if gateway_ip and not isinstance(gateway_ip, str):
            raise TypeError("Expected argument 'gateway_ip' to be a str")
        pulumi.set(__self__, "gateway_ip", gateway_ip)
        if host_routes and not isinstance(host_routes, list):
            raise TypeError("Expected argument 'host_routes' to be a list")
        pulumi.set(__self__, "host_routes", host_routes)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_types and not isinstance(service_types, list):
            raise TypeError("Expected argument 'service_types' to be a list")
        pulumi.set(__self__, "service_types", service_types)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
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
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Sequence[str]:
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="allocationPools")
    def allocation_pools(self) -> Sequence['outputs.GetSubnetAllocationPoolResult']:
        return pulumi.get(self, "allocation_pools")

    @property
    @pulumi.getter
    def cidr(self) -> str:
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dhcpEnabled")
    def dhcp_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "dhcp_enabled")

    @property
    @pulumi.getter(name="dnsNameservers")
    def dns_nameservers(self) -> Sequence[str]:
        return pulumi.get(self, "dns_nameservers")

    @property
    @pulumi.getter(name="dnsPublishFixedIp")
    def dns_publish_fixed_ip(self) -> Optional[bool]:
        return pulumi.get(self, "dns_publish_fixed_ip")

    @property
    @pulumi.getter(name="enableDhcp")
    def enable_dhcp(self) -> bool:
        return pulumi.get(self, "enable_dhcp")

    @property
    @pulumi.getter(name="gatewayIp")
    def gateway_ip(self) -> str:
        return pulumi.get(self, "gateway_ip")

    @property
    @pulumi.getter(name="hostRoutes")
    def host_routes(self) -> Sequence['outputs.GetSubnetHostRouteResult']:
        return pulumi.get(self, "host_routes")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> int:
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="ipv6AddressMode")
    def ipv6_address_mode(self) -> str:
        return pulumi.get(self, "ipv6_address_mode")

    @property
    @pulumi.getter(name="ipv6RaMode")
    def ipv6_ra_mode(self) -> str:
        return pulumi.get(self, "ipv6_ra_mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceTypes")
    def service_types(self) -> Sequence[str]:
        return pulumi.get(self, "service_types")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="subnetpoolId")
    def subnetpool_id(self) -> str:
        return pulumi.get(self, "subnetpool_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        return pulumi.get(self, "tenant_id")


class AwaitableGetSubnetResult(GetSubnetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetResult(
            all_tags=self.all_tags,
            allocation_pools=self.allocation_pools,
            cidr=self.cidr,
            description=self.description,
            dhcp_enabled=self.dhcp_enabled,
            dns_nameservers=self.dns_nameservers,
            dns_publish_fixed_ip=self.dns_publish_fixed_ip,
            enable_dhcp=self.enable_dhcp,
            gateway_ip=self.gateway_ip,
            host_routes=self.host_routes,
            id=self.id,
            ip_version=self.ip_version,
            ipv6_address_mode=self.ipv6_address_mode,
            ipv6_ra_mode=self.ipv6_ra_mode,
            name=self.name,
            network_id=self.network_id,
            region=self.region,
            service_types=self.service_types,
            subnet_id=self.subnet_id,
            subnetpool_id=self.subnetpool_id,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_subnet(cidr: Optional[str] = None,
               description: Optional[str] = None,
               dhcp_enabled: Optional[bool] = None,
               dns_publish_fixed_ip: Optional[bool] = None,
               gateway_ip: Optional[str] = None,
               ip_version: Optional[int] = None,
               ipv6_address_mode: Optional[str] = None,
               ipv6_ra_mode: Optional[str] = None,
               name: Optional[str] = None,
               network_id: Optional[str] = None,
               region: Optional[str] = None,
               subnet_id: Optional[str] = None,
               subnetpool_id: Optional[str] = None,
               tags: Optional[Sequence[str]] = None,
               tenant_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['description'] = description
    __args__['dhcpEnabled'] = dhcp_enabled
    __args__['dnsPublishFixedIp'] = dns_publish_fixed_ip
    __args__['gatewayIp'] = gateway_ip
    __args__['ipVersion'] = ip_version
    __args__['ipv6AddressMode'] = ipv6_address_mode
    __args__['ipv6RaMode'] = ipv6_ra_mode
    __args__['name'] = name
    __args__['networkId'] = network_id
    __args__['region'] = region
    __args__['subnetId'] = subnet_id
    __args__['subnetpoolId'] = subnetpool_id
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getSubnet:getSubnet', __args__, opts=opts, typ=GetSubnetResult).value

    return AwaitableGetSubnetResult(
        all_tags=pulumi.get(__ret__, 'all_tags'),
        allocation_pools=pulumi.get(__ret__, 'allocation_pools'),
        cidr=pulumi.get(__ret__, 'cidr'),
        description=pulumi.get(__ret__, 'description'),
        dhcp_enabled=pulumi.get(__ret__, 'dhcp_enabled'),
        dns_nameservers=pulumi.get(__ret__, 'dns_nameservers'),
        dns_publish_fixed_ip=pulumi.get(__ret__, 'dns_publish_fixed_ip'),
        enable_dhcp=pulumi.get(__ret__, 'enable_dhcp'),
        gateway_ip=pulumi.get(__ret__, 'gateway_ip'),
        host_routes=pulumi.get(__ret__, 'host_routes'),
        id=pulumi.get(__ret__, 'id'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        ipv6_address_mode=pulumi.get(__ret__, 'ipv6_address_mode'),
        ipv6_ra_mode=pulumi.get(__ret__, 'ipv6_ra_mode'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        region=pulumi.get(__ret__, 'region'),
        service_types=pulumi.get(__ret__, 'service_types'),
        subnet_id=pulumi.get(__ret__, 'subnet_id'),
        subnetpool_id=pulumi.get(__ret__, 'subnetpool_id'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_subnet)
def get_subnet_output(cidr: Optional[pulumi.Input[Optional[str]]] = None,
                      description: Optional[pulumi.Input[Optional[str]]] = None,
                      dhcp_enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                      dns_publish_fixed_ip: Optional[pulumi.Input[Optional[bool]]] = None,
                      gateway_ip: Optional[pulumi.Input[Optional[str]]] = None,
                      ip_version: Optional[pulumi.Input[Optional[int]]] = None,
                      ipv6_address_mode: Optional[pulumi.Input[Optional[str]]] = None,
                      ipv6_ra_mode: Optional[pulumi.Input[Optional[str]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      network_id: Optional[pulumi.Input[Optional[str]]] = None,
                      region: Optional[pulumi.Input[Optional[str]]] = None,
                      subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                      subnetpool_id: Optional[pulumi.Input[Optional[str]]] = None,
                      tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
