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
    'GetSubnetPoolResult',
    'AwaitableGetSubnetPoolResult',
    'get_subnet_pool',
    'get_subnet_pool_output',
]

@pulumi.output_type
class GetSubnetPoolResult:
    """
    A collection of values returned by getSubnetPool.
    """
    def __init__(__self__, address_scope_id=None, all_tags=None, created_at=None, default_prefixlen=None, default_quota=None, description=None, id=None, ip_version=None, is_default=None, max_prefixlen=None, min_prefixlen=None, name=None, prefixes=None, project_id=None, region=None, revision_number=None, shared=None, tags=None, updated_at=None):
        if address_scope_id and not isinstance(address_scope_id, str):
            raise TypeError("Expected argument 'address_scope_id' to be a str")
        pulumi.set(__self__, "address_scope_id", address_scope_id)
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if default_prefixlen and not isinstance(default_prefixlen, int):
            raise TypeError("Expected argument 'default_prefixlen' to be a int")
        pulumi.set(__self__, "default_prefixlen", default_prefixlen)
        if default_quota and not isinstance(default_quota, int):
            raise TypeError("Expected argument 'default_quota' to be a int")
        pulumi.set(__self__, "default_quota", default_quota)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_version and not isinstance(ip_version, int):
            raise TypeError("Expected argument 'ip_version' to be a int")
        pulumi.set(__self__, "ip_version", ip_version)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if max_prefixlen and not isinstance(max_prefixlen, int):
            raise TypeError("Expected argument 'max_prefixlen' to be a int")
        pulumi.set(__self__, "max_prefixlen", max_prefixlen)
        if min_prefixlen and not isinstance(min_prefixlen, int):
            raise TypeError("Expected argument 'min_prefixlen' to be a int")
        pulumi.set(__self__, "min_prefixlen", min_prefixlen)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if prefixes and not isinstance(prefixes, list):
            raise TypeError("Expected argument 'prefixes' to be a list")
        pulumi.set(__self__, "prefixes", prefixes)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if revision_number and not isinstance(revision_number, int):
            raise TypeError("Expected argument 'revision_number' to be a int")
        pulumi.set(__self__, "revision_number", revision_number)
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        pulumi.set(__self__, "shared", shared)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="addressScopeId")
    def address_scope_id(self) -> str:
        return pulumi.get(self, "address_scope_id")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Sequence[str]:
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultPrefixlen")
    def default_prefixlen(self) -> int:
        return pulumi.get(self, "default_prefixlen")

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> int:
        return pulumi.get(self, "default_quota")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="maxPrefixlen")
    def max_prefixlen(self) -> int:
        return pulumi.get(self, "max_prefixlen")

    @property
    @pulumi.getter(name="minPrefixlen")
    def min_prefixlen(self) -> int:
        return pulumi.get(self, "min_prefixlen")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def prefixes(self) -> Sequence[str]:
        return pulumi.get(self, "prefixes")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="revisionNumber")
    def revision_number(self) -> int:
        return pulumi.get(self, "revision_number")

    @property
    @pulumi.getter
    def shared(self) -> bool:
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetSubnetPoolResult(GetSubnetPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetPoolResult(
            address_scope_id=self.address_scope_id,
            all_tags=self.all_tags,
            created_at=self.created_at,
            default_prefixlen=self.default_prefixlen,
            default_quota=self.default_quota,
            description=self.description,
            id=self.id,
            ip_version=self.ip_version,
            is_default=self.is_default,
            max_prefixlen=self.max_prefixlen,
            min_prefixlen=self.min_prefixlen,
            name=self.name,
            prefixes=self.prefixes,
            project_id=self.project_id,
            region=self.region,
            revision_number=self.revision_number,
            shared=self.shared,
            tags=self.tags,
            updated_at=self.updated_at)


def get_subnet_pool(address_scope_id: Optional[str] = None,
                    default_prefixlen: Optional[int] = None,
                    default_quota: Optional[int] = None,
                    description: Optional[str] = None,
                    ip_version: Optional[int] = None,
                    is_default: Optional[bool] = None,
                    max_prefixlen: Optional[int] = None,
                    min_prefixlen: Optional[int] = None,
                    name: Optional[str] = None,
                    project_id: Optional[str] = None,
                    region: Optional[str] = None,
                    shared: Optional[bool] = None,
                    tags: Optional[Sequence[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetPoolResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressScopeId'] = address_scope_id
    __args__['defaultPrefixlen'] = default_prefixlen
    __args__['defaultQuota'] = default_quota
    __args__['description'] = description
    __args__['ipVersion'] = ip_version
    __args__['isDefault'] = is_default
    __args__['maxPrefixlen'] = max_prefixlen
    __args__['minPrefixlen'] = min_prefixlen
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getSubnetPool:getSubnetPool', __args__, opts=opts, typ=GetSubnetPoolResult).value

    return AwaitableGetSubnetPoolResult(
        address_scope_id=pulumi.get(__ret__, 'address_scope_id'),
        all_tags=pulumi.get(__ret__, 'all_tags'),
        created_at=pulumi.get(__ret__, 'created_at'),
        default_prefixlen=pulumi.get(__ret__, 'default_prefixlen'),
        default_quota=pulumi.get(__ret__, 'default_quota'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        is_default=pulumi.get(__ret__, 'is_default'),
        max_prefixlen=pulumi.get(__ret__, 'max_prefixlen'),
        min_prefixlen=pulumi.get(__ret__, 'min_prefixlen'),
        name=pulumi.get(__ret__, 'name'),
        prefixes=pulumi.get(__ret__, 'prefixes'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        revision_number=pulumi.get(__ret__, 'revision_number'),
        shared=pulumi.get(__ret__, 'shared'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_subnet_pool)
def get_subnet_pool_output(address_scope_id: Optional[pulumi.Input[Optional[str]]] = None,
                           default_prefixlen: Optional[pulumi.Input[Optional[int]]] = None,
                           default_quota: Optional[pulumi.Input[Optional[int]]] = None,
                           description: Optional[pulumi.Input[Optional[str]]] = None,
                           ip_version: Optional[pulumi.Input[Optional[int]]] = None,
                           is_default: Optional[pulumi.Input[Optional[bool]]] = None,
                           max_prefixlen: Optional[pulumi.Input[Optional[int]]] = None,
                           min_prefixlen: Optional[pulumi.Input[Optional[int]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           project_id: Optional[pulumi.Input[Optional[str]]] = None,
                           region: Optional[pulumi.Input[Optional[str]]] = None,
                           shared: Optional[pulumi.Input[Optional[bool]]] = None,
                           tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetPoolResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
