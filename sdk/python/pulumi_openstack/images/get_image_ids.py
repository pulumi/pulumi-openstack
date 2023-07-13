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
    'GetImageIdsResult',
    'AwaitableGetImageIdsResult',
    'get_image_ids',
    'get_image_ids_output',
]

@pulumi.output_type
class GetImageIdsResult:
    """
    A collection of values returned by getImageIds.
    """
    def __init__(__self__, id=None, ids=None, member_status=None, name=None, name_regex=None, owner=None, properties=None, region=None, size_max=None, size_min=None, sort=None, sort_direction=None, sort_key=None, tag=None, tags=None, visibility=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if member_status and not isinstance(member_status, str):
            raise TypeError("Expected argument 'member_status' to be a str")
        pulumi.set(__self__, "member_status", member_status)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if size_max and not isinstance(size_max, int):
            raise TypeError("Expected argument 'size_max' to be a int")
        pulumi.set(__self__, "size_max", size_max)
        if size_min and not isinstance(size_min, int):
            raise TypeError("Expected argument 'size_min' to be a int")
        pulumi.set(__self__, "size_min", size_min)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        pulumi.set(__self__, "sort_direction", sort_direction)
        if sort_key and not isinstance(sort_key, str):
            raise TypeError("Expected argument 'sort_key' to be a str")
        pulumi.set(__self__, "sort_key", sort_key)
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        pulumi.set(__self__, "tag", tag)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

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
    @pulumi.getter(name="memberStatus")
    def member_status(self) -> Optional[str]:
        return pulumi.get(self, "member_status")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def properties(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sizeMax")
    def size_max(self) -> Optional[int]:
        return pulumi.get(self, "size_max")

    @property
    @pulumi.getter(name="sizeMin")
    def size_min(self) -> Optional[int]:
        return pulumi.get(self, "size_min")

    @property
    @pulumi.getter
    def sort(self) -> Optional[str]:
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter(name="sortDirection")
    def sort_direction(self) -> Optional[str]:
        warnings.warn("""Use option 'sort' instead.""", DeprecationWarning)
        pulumi.log.warn("""sort_direction is deprecated: Use option 'sort' instead.""")

        return pulumi.get(self, "sort_direction")

    @property
    @pulumi.getter(name="sortKey")
    def sort_key(self) -> Optional[str]:
        warnings.warn("""Use option 'sort' instead.""", DeprecationWarning)
        pulumi.log.warn("""sort_key is deprecated: Use option 'sort' instead.""")

        return pulumi.get(self, "sort_key")

    @property
    @pulumi.getter
    def tag(self) -> Optional[str]:
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        return pulumi.get(self, "visibility")


class AwaitableGetImageIdsResult(GetImageIdsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageIdsResult(
            id=self.id,
            ids=self.ids,
            member_status=self.member_status,
            name=self.name,
            name_regex=self.name_regex,
            owner=self.owner,
            properties=self.properties,
            region=self.region,
            size_max=self.size_max,
            size_min=self.size_min,
            sort=self.sort,
            sort_direction=self.sort_direction,
            sort_key=self.sort_key,
            tag=self.tag,
            tags=self.tags,
            visibility=self.visibility)


def get_image_ids(member_status: Optional[str] = None,
                  name: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  owner: Optional[str] = None,
                  properties: Optional[Mapping[str, Any]] = None,
                  region: Optional[str] = None,
                  size_max: Optional[int] = None,
                  size_min: Optional[int] = None,
                  sort: Optional[str] = None,
                  sort_direction: Optional[str] = None,
                  sort_key: Optional[str] = None,
                  tag: Optional[str] = None,
                  tags: Optional[Sequence[str]] = None,
                  visibility: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageIdsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['memberStatus'] = member_status
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['owner'] = owner
    __args__['properties'] = properties
    __args__['region'] = region
    __args__['sizeMax'] = size_max
    __args__['sizeMin'] = size_min
    __args__['sort'] = sort
    __args__['sortDirection'] = sort_direction
    __args__['sortKey'] = sort_key
    __args__['tag'] = tag
    __args__['tags'] = tags
    __args__['visibility'] = visibility
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:images/getImageIds:getImageIds', __args__, opts=opts, typ=GetImageIdsResult).value

    return AwaitableGetImageIdsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        member_status=pulumi.get(__ret__, 'member_status'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        owner=pulumi.get(__ret__, 'owner'),
        properties=pulumi.get(__ret__, 'properties'),
        region=pulumi.get(__ret__, 'region'),
        size_max=pulumi.get(__ret__, 'size_max'),
        size_min=pulumi.get(__ret__, 'size_min'),
        sort=pulumi.get(__ret__, 'sort'),
        sort_direction=pulumi.get(__ret__, 'sort_direction'),
        sort_key=pulumi.get(__ret__, 'sort_key'),
        tag=pulumi.get(__ret__, 'tag'),
        tags=pulumi.get(__ret__, 'tags'),
        visibility=pulumi.get(__ret__, 'visibility'))


@_utilities.lift_output_func(get_image_ids)
def get_image_ids_output(member_status: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         owner: Optional[pulumi.Input[Optional[str]]] = None,
                         properties: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         size_max: Optional[pulumi.Input[Optional[int]]] = None,
                         size_min: Optional[pulumi.Input[Optional[int]]] = None,
                         sort: Optional[pulumi.Input[Optional[str]]] = None,
                         sort_direction: Optional[pulumi.Input[Optional[str]]] = None,
                         sort_key: Optional[pulumi.Input[Optional[str]]] = None,
                         tag: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         visibility: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageIdsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
