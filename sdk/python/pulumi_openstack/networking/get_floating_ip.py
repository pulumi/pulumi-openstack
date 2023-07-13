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
    'GetFloatingIpResult',
    'AwaitableGetFloatingIpResult',
    'get_floating_ip',
    'get_floating_ip_output',
]

@pulumi.output_type
class GetFloatingIpResult:
    """
    A collection of values returned by getFloatingIp.
    """
    def __init__(__self__, address=None, all_tags=None, description=None, dns_domain=None, dns_name=None, fixed_ip=None, id=None, pool=None, port_id=None, region=None, status=None, tags=None, tenant_id=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dns_domain and not isinstance(dns_domain, str):
            raise TypeError("Expected argument 'dns_domain' to be a str")
        pulumi.set(__self__, "dns_domain", dns_domain)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if fixed_ip and not isinstance(fixed_ip, str):
            raise TypeError("Expected argument 'fixed_ip' to be a str")
        pulumi.set(__self__, "fixed_ip", fixed_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pool and not isinstance(pool, str):
            raise TypeError("Expected argument 'pool' to be a str")
        pulumi.set(__self__, "pool", pool)
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        pulumi.set(__self__, "port_id", port_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Sequence[str]:
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsDomain")
    def dns_domain(self) -> str:
        return pulumi.get(self, "dns_domain")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[str]:
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def pool(self) -> Optional[str]:
        return pulumi.get(self, "pool")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[str]:
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


class AwaitableGetFloatingIpResult(GetFloatingIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFloatingIpResult(
            address=self.address,
            all_tags=self.all_tags,
            description=self.description,
            dns_domain=self.dns_domain,
            dns_name=self.dns_name,
            fixed_ip=self.fixed_ip,
            id=self.id,
            pool=self.pool,
            port_id=self.port_id,
            region=self.region,
            status=self.status,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_floating_ip(address: Optional[str] = None,
                    description: Optional[str] = None,
                    fixed_ip: Optional[str] = None,
                    pool: Optional[str] = None,
                    port_id: Optional[str] = None,
                    region: Optional[str] = None,
                    status: Optional[str] = None,
                    tags: Optional[Sequence[str]] = None,
                    tenant_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFloatingIpResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['address'] = address
    __args__['description'] = description
    __args__['fixedIp'] = fixed_ip
    __args__['pool'] = pool
    __args__['portId'] = port_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getFloatingIp:getFloatingIp', __args__, opts=opts, typ=GetFloatingIpResult).value

    return AwaitableGetFloatingIpResult(
        address=pulumi.get(__ret__, 'address'),
        all_tags=pulumi.get(__ret__, 'all_tags'),
        description=pulumi.get(__ret__, 'description'),
        dns_domain=pulumi.get(__ret__, 'dns_domain'),
        dns_name=pulumi.get(__ret__, 'dns_name'),
        fixed_ip=pulumi.get(__ret__, 'fixed_ip'),
        id=pulumi.get(__ret__, 'id'),
        pool=pulumi.get(__ret__, 'pool'),
        port_id=pulumi.get(__ret__, 'port_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_floating_ip)
def get_floating_ip_output(address: Optional[pulumi.Input[Optional[str]]] = None,
                           description: Optional[pulumi.Input[Optional[str]]] = None,
                           fixed_ip: Optional[pulumi.Input[Optional[str]]] = None,
                           pool: Optional[pulumi.Input[Optional[str]]] = None,
                           port_id: Optional[pulumi.Input[Optional[str]]] = None,
                           region: Optional[pulumi.Input[Optional[str]]] = None,
                           status: Optional[pulumi.Input[Optional[str]]] = None,
                           tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFloatingIpResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
