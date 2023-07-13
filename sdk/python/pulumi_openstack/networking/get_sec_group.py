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
    'GetSecGroupResult',
    'AwaitableGetSecGroupResult',
    'get_sec_group',
    'get_sec_group_output',
]

@pulumi.output_type
class GetSecGroupResult:
    """
    A collection of values returned by getSecGroup.
    """
    def __init__(__self__, all_tags=None, description=None, id=None, name=None, region=None, secgroup_id=None, tags=None, tenant_id=None):
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if secgroup_id and not isinstance(secgroup_id, str):
            raise TypeError("Expected argument 'secgroup_id' to be a str")
        pulumi.set(__self__, "secgroup_id", secgroup_id)
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
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

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
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secgroupId")
    def secgroup_id(self) -> Optional[str]:
        return pulumi.get(self, "secgroup_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        return pulumi.get(self, "tenant_id")


class AwaitableGetSecGroupResult(GetSecGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecGroupResult(
            all_tags=self.all_tags,
            description=self.description,
            id=self.id,
            name=self.name,
            region=self.region,
            secgroup_id=self.secgroup_id,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_sec_group(description: Optional[str] = None,
                  name: Optional[str] = None,
                  region: Optional[str] = None,
                  secgroup_id: Optional[str] = None,
                  tags: Optional[Sequence[str]] = None,
                  tenant_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecGroupResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['name'] = name
    __args__['region'] = region
    __args__['secgroupId'] = secgroup_id
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getSecGroup:getSecGroup', __args__, opts=opts, typ=GetSecGroupResult).value

    return AwaitableGetSecGroupResult(
        all_tags=pulumi.get(__ret__, 'all_tags'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        secgroup_id=pulumi.get(__ret__, 'secgroup_id'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_sec_group)
def get_sec_group_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         secgroup_id: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecGroupResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
