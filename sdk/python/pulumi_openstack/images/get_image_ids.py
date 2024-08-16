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
    def __init__(__self__, container_format=None, disk_format=None, hidden=None, id=None, ids=None, member_status=None, name=None, name_regex=None, owner=None, properties=None, region=None, size_max=None, size_min=None, sort=None, tag=None, tags=None, visibility=None):
        if container_format and not isinstance(container_format, str):
            raise TypeError("Expected argument 'container_format' to be a str")
        pulumi.set(__self__, "container_format", container_format)
        if disk_format and not isinstance(disk_format, str):
            raise TypeError("Expected argument 'disk_format' to be a str")
        pulumi.set(__self__, "disk_format", disk_format)
        if hidden and not isinstance(hidden, bool):
            raise TypeError("Expected argument 'hidden' to be a bool")
        pulumi.set(__self__, "hidden", hidden)
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
    @pulumi.getter(name="containerFormat")
    def container_format(self) -> Optional[str]:
        return pulumi.get(self, "container_format")

    @property
    @pulumi.getter(name="diskFormat")
    def disk_format(self) -> Optional[str]:
        return pulumi.get(self, "disk_format")

    @property
    @pulumi.getter
    def hidden(self) -> Optional[bool]:
        return pulumi.get(self, "hidden")

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
    def properties(self) -> Optional[Mapping[str, str]]:
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
            container_format=self.container_format,
            disk_format=self.disk_format,
            hidden=self.hidden,
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
            tag=self.tag,
            tags=self.tags,
            visibility=self.visibility)


def get_image_ids(container_format: Optional[str] = None,
                  disk_format: Optional[str] = None,
                  hidden: Optional[bool] = None,
                  member_status: Optional[str] = None,
                  name: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  owner: Optional[str] = None,
                  properties: Optional[Mapping[str, str]] = None,
                  region: Optional[str] = None,
                  size_max: Optional[int] = None,
                  size_min: Optional[int] = None,
                  sort: Optional[str] = None,
                  tag: Optional[str] = None,
                  tags: Optional[Sequence[str]] = None,
                  visibility: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageIdsResult:
    """
    Use this data source to get a list of Openstack Image IDs matching the
    specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    images = openstack.images.get_image_ids(name_regex="^Ubuntu 16\\\\.04.*-amd64",
        sort="updated_at",
        properties={
            "key": "value",
        })
    ```


    :param str container_format: The container format of the image.
    :param str disk_format: The disk format of the image.
    :param bool hidden: Whether or not the image is hidden from public list.
    :param str member_status: The status of the image. Must be one of
           "accepted", "pending", "rejected", or "all".
    :param str name: The name of the image. Cannot be used simultaneously with
           `name_regex`.
    :param str name_regex: The regular expressian of the name of the image.
           Cannot be used simultaneously with `name`. Unlike filtering by `name` the
           `name_regex` filtering does by client on the result of OpenStack search
           query.
    :param str owner: The owner (UUID) of the image.
    :param Mapping[str, str] properties: a map of key/value pairs to match an image with.
           All specified properties must be matched. Unlike other options filtering by
           `properties` does by client on the result of OpenStack search query.
    :param str region: The region in which to obtain the V2 Glance client. A
           Glance client is needed to create an Image that can be used with a compute
           instance. If omitted, the `region` argument of the provider is used.
    :param int size_max: The maximum size (in bytes) of the image to return.
    :param int size_min: The minimum size (in bytes) of the image to return.
    :param str sort: Sorts the response by one or more attribute and sort
           direction combinations. You can also set multiple sort keys and directions.
           Default direction is `desc`. Use the comma (,) character to separate multiple
           values. For example expression `sort = "name:asc,status"` sorts ascending by
           name and descending by status.
    :param str tag: Search for images with a specific tag.
    :param Sequence[str] tags: A list of tags required to be set on the image (all
           specified tags must be in the images tag list for it to be matched).
    :param str visibility: The visibility of the image. Must be one of
           "public", "private", "community", or "shared". Defaults to "private".
    """
    __args__ = dict()
    __args__['containerFormat'] = container_format
    __args__['diskFormat'] = disk_format
    __args__['hidden'] = hidden
    __args__['memberStatus'] = member_status
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['owner'] = owner
    __args__['properties'] = properties
    __args__['region'] = region
    __args__['sizeMax'] = size_max
    __args__['sizeMin'] = size_min
    __args__['sort'] = sort
    __args__['tag'] = tag
    __args__['tags'] = tags
    __args__['visibility'] = visibility
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:images/getImageIds:getImageIds', __args__, opts=opts, typ=GetImageIdsResult).value

    return AwaitableGetImageIdsResult(
        container_format=pulumi.get(__ret__, 'container_format'),
        disk_format=pulumi.get(__ret__, 'disk_format'),
        hidden=pulumi.get(__ret__, 'hidden'),
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
        tag=pulumi.get(__ret__, 'tag'),
        tags=pulumi.get(__ret__, 'tags'),
        visibility=pulumi.get(__ret__, 'visibility'))


@_utilities.lift_output_func(get_image_ids)
def get_image_ids_output(container_format: Optional[pulumi.Input[Optional[str]]] = None,
                         disk_format: Optional[pulumi.Input[Optional[str]]] = None,
                         hidden: Optional[pulumi.Input[Optional[bool]]] = None,
                         member_status: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         owner: Optional[pulumi.Input[Optional[str]]] = None,
                         properties: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         size_max: Optional[pulumi.Input[Optional[int]]] = None,
                         size_min: Optional[pulumi.Input[Optional[int]]] = None,
                         sort: Optional[pulumi.Input[Optional[str]]] = None,
                         tag: Optional[pulumi.Input[Optional[str]]] = None,
                         tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         visibility: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageIdsResult]:
    """
    Use this data source to get a list of Openstack Image IDs matching the
    specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    images = openstack.images.get_image_ids(name_regex="^Ubuntu 16\\\\.04.*-amd64",
        sort="updated_at",
        properties={
            "key": "value",
        })
    ```


    :param str container_format: The container format of the image.
    :param str disk_format: The disk format of the image.
    :param bool hidden: Whether or not the image is hidden from public list.
    :param str member_status: The status of the image. Must be one of
           "accepted", "pending", "rejected", or "all".
    :param str name: The name of the image. Cannot be used simultaneously with
           `name_regex`.
    :param str name_regex: The regular expressian of the name of the image.
           Cannot be used simultaneously with `name`. Unlike filtering by `name` the
           `name_regex` filtering does by client on the result of OpenStack search
           query.
    :param str owner: The owner (UUID) of the image.
    :param Mapping[str, str] properties: a map of key/value pairs to match an image with.
           All specified properties must be matched. Unlike other options filtering by
           `properties` does by client on the result of OpenStack search query.
    :param str region: The region in which to obtain the V2 Glance client. A
           Glance client is needed to create an Image that can be used with a compute
           instance. If omitted, the `region` argument of the provider is used.
    :param int size_max: The maximum size (in bytes) of the image to return.
    :param int size_min: The minimum size (in bytes) of the image to return.
    :param str sort: Sorts the response by one or more attribute and sort
           direction combinations. You can also set multiple sort keys and directions.
           Default direction is `desc`. Use the comma (,) character to separate multiple
           values. For example expression `sort = "name:asc,status"` sorts ascending by
           name and descending by status.
    :param str tag: Search for images with a specific tag.
    :param Sequence[str] tags: A list of tags required to be set on the image (all
           specified tags must be in the images tag list for it to be matched).
    :param str visibility: The visibility of the image. Must be one of
           "public", "private", "community", or "shared". Defaults to "private".
    """
    ...
