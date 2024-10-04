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
    'GetRouterResult',
    'AwaitableGetRouterResult',
    'get_router',
    'get_router_output',
]

@pulumi.output_type
class GetRouterResult:
    """
    A collection of values returned by getRouter.
    """
    def __init__(__self__, admin_state_up=None, all_tags=None, availability_zone_hints=None, description=None, distributed=None, enable_snat=None, external_fixed_ips=None, external_network_id=None, id=None, name=None, region=None, router_id=None, status=None, tags=None, tenant_id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        pulumi.set(__self__, "admin_state_up", admin_state_up)
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if availability_zone_hints and not isinstance(availability_zone_hints, list):
            raise TypeError("Expected argument 'availability_zone_hints' to be a list")
        pulumi.set(__self__, "availability_zone_hints", availability_zone_hints)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if distributed and not isinstance(distributed, bool):
            raise TypeError("Expected argument 'distributed' to be a bool")
        pulumi.set(__self__, "distributed", distributed)
        if enable_snat and not isinstance(enable_snat, bool):
            raise TypeError("Expected argument 'enable_snat' to be a bool")
        pulumi.set(__self__, "enable_snat", enable_snat)
        if external_fixed_ips and not isinstance(external_fixed_ips, list):
            raise TypeError("Expected argument 'external_fixed_ips' to be a list")
        pulumi.set(__self__, "external_fixed_ips", external_fixed_ips)
        if external_network_id and not isinstance(external_network_id, str):
            raise TypeError("Expected argument 'external_network_id' to be a str")
        pulumi.set(__self__, "external_network_id", external_network_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        pulumi.set(__self__, "router_id", router_id)
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
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[bool]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Sequence[str]:
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="availabilityZoneHints")
    def availability_zone_hints(self) -> Sequence[str]:
        return pulumi.get(self, "availability_zone_hints")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def distributed(self) -> Optional[bool]:
        return pulumi.get(self, "distributed")

    @property
    @pulumi.getter(name="enableSnat")
    def enable_snat(self) -> bool:
        return pulumi.get(self, "enable_snat")

    @property
    @pulumi.getter(name="externalFixedIps")
    def external_fixed_ips(self) -> Sequence['outputs.GetRouterExternalFixedIpResult']:
        return pulumi.get(self, "external_fixed_ips")

    @property
    @pulumi.getter(name="externalNetworkId")
    def external_network_id(self) -> str:
        return pulumi.get(self, "external_network_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[str]:
        return pulumi.get(self, "router_id")

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


class AwaitableGetRouterResult(GetRouterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterResult(
            admin_state_up=self.admin_state_up,
            all_tags=self.all_tags,
            availability_zone_hints=self.availability_zone_hints,
            description=self.description,
            distributed=self.distributed,
            enable_snat=self.enable_snat,
            external_fixed_ips=self.external_fixed_ips,
            external_network_id=self.external_network_id,
            id=self.id,
            name=self.name,
            region=self.region,
            router_id=self.router_id,
            status=self.status,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_router(admin_state_up: Optional[bool] = None,
               description: Optional[str] = None,
               distributed: Optional[bool] = None,
               enable_snat: Optional[bool] = None,
               name: Optional[str] = None,
               region: Optional[str] = None,
               router_id: Optional[str] = None,
               status: Optional[str] = None,
               tags: Optional[Sequence[str]] = None,
               tenant_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['adminStateUp'] = admin_state_up
    __args__['description'] = description
    __args__['distributed'] = distributed
    __args__['enableSnat'] = enable_snat
    __args__['name'] = name
    __args__['region'] = region
    __args__['routerId'] = router_id
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getRouter:getRouter', __args__, opts=opts, typ=GetRouterResult).value

    return AwaitableGetRouterResult(
        admin_state_up=pulumi.get(__ret__, 'admin_state_up'),
        all_tags=pulumi.get(__ret__, 'all_tags'),
        availability_zone_hints=pulumi.get(__ret__, 'availability_zone_hints'),
        description=pulumi.get(__ret__, 'description'),
        distributed=pulumi.get(__ret__, 'distributed'),
        enable_snat=pulumi.get(__ret__, 'enable_snat'),
        external_fixed_ips=pulumi.get(__ret__, 'external_fixed_ips'),
        external_network_id=pulumi.get(__ret__, 'external_network_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        router_id=pulumi.get(__ret__, 'router_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_router)
def get_router_output(admin_state_up: Optional[pulumi.Input[Optional[bool]]] = None,
                      description: Optional[pulumi.Input[Optional[str]]] = None,
                      distributed: Optional[pulumi.Input[Optional[bool]]] = None,
                      enable_snat: Optional[pulumi.Input[Optional[bool]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      region: Optional[pulumi.Input[Optional[str]]] = None,
                      router_id: Optional[pulumi.Input[Optional[str]]] = None,
                      status: Optional[pulumi.Input[Optional[str]]] = None,
                      tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
