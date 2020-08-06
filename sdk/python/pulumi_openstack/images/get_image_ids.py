# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetImageIdsResult:
    """
    A collection of values returned by getImageIds.
    """
    def __init__(__self__, id=None, ids=None, member_status=None, name=None, name_regex=None, owner=None, properties=None, region=None, size_max=None, size_min=None, sort=None, sort_direction=None, sort_key=None, tag=None, visibility=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        __self__.ids = ids
        if member_status and not isinstance(member_status, str):
            raise TypeError("Expected argument 'member_status' to be a str")
        __self__.member_status = member_status
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        __self__.name_regex = name_regex
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        __self__.owner = owner
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if size_max and not isinstance(size_max, float):
            raise TypeError("Expected argument 'size_max' to be a float")
        __self__.size_max = size_max
        if size_min and not isinstance(size_min, float):
            raise TypeError("Expected argument 'size_min' to be a float")
        __self__.size_min = size_min
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        __self__.sort = sort
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        if sort_direction is not None:
            warnings.warn("Use option 'sort' instead.", DeprecationWarning)
            pulumi.log.warn("sort_direction is deprecated: Use option 'sort' instead.")

        __self__.sort_direction = sort_direction
        if sort_key and not isinstance(sort_key, str):
            raise TypeError("Expected argument 'sort_key' to be a str")
        if sort_key is not None:
            warnings.warn("Use option 'sort' instead.", DeprecationWarning)
            pulumi.log.warn("sort_key is deprecated: Use option 'sort' instead.")

        __self__.sort_key = sort_key
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        __self__.tag = tag
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        __self__.visibility = visibility
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
            visibility=self.visibility)

def get_image_ids(member_status=None,name=None,name_regex=None,owner=None,properties=None,region=None,size_max=None,size_min=None,sort=None,sort_direction=None,sort_key=None,tag=None,visibility=None,opts=None):
    """
    Use this data source to get a list of Openstack Image IDs matching the
    specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    images = openstack.images.get_image_ids(name_regex="^Ubuntu 16\\.04.*-amd64",
        properties={
            "key": "value",
        },
        sort="updated_at")
    ```


    :param str member_status: The status of the image. Must be one of
           "accepted", "pending", "rejected", or "all".
    :param str name: The name of the image. Cannot be used simultaneously
           with `name_regex`.
    :param str name_regex: The regular expressian of the name of the image.
           Cannot be used simultaneously with `name`. Unlike filtering by `name` the
           `name_regex` filtering does by client on the result of OpenStack search
           query.
    :param str owner: The owner (UUID) of the image.
    :param dict properties: a map of key/value pairs to match an image with.
           All specified properties must be matched. Unlike other options filtering
           by `properties` does by client on the result of OpenStack search query.
    :param str region: The region in which to obtain the V2 Glance client.
           A Glance client is needed to create an Image that can be used with
           a compute instance. If omitted, the `region` argument of the provider
           is used.
    :param float size_max: The maximum size (in bytes) of the image to return.
    :param float size_min: The minimum size (in bytes) of the image to return.
    :param str sort: Sorts the response by one or more attribute and sort
           direction combinations. You can also set multiple sort keys and directions.
           Default direction is `desc`. Use the comma (,) character to separate
           multiple values. For example expression `sort = "name:asc,status"`
           sorts ascending by name and descending by status. `sort` cannot be used
           simultaneously with `sort_key`. If both are present in a configuration
           then only `sort` will be used.
    :param str sort_direction: Order the results in either `asc` or `desc`.
           Can be applied only with `sort_key`. Defaults to `asc`
    :param str sort_key: Sort images based on a certain key. Defaults to
           `name`. `sort_key` cannot be used simultaneously with `sort`. If both
           are present in a configuration then only `sort` will be used.
    :param str tag: Search for images with a specific tag.
    :param str visibility: The visibility of the image. Must be one of
           "public", "private", "community", or "shared". Defaults to "private".
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
    __args__['visibility'] = visibility
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:images/getImageIds:getImageIds', __args__, opts=opts).value

    return AwaitableGetImageIdsResult(
        id=__ret__.get('id'),
        ids=__ret__.get('ids'),
        member_status=__ret__.get('memberStatus'),
        name=__ret__.get('name'),
        name_regex=__ret__.get('nameRegex'),
        owner=__ret__.get('owner'),
        properties=__ret__.get('properties'),
        region=__ret__.get('region'),
        size_max=__ret__.get('sizeMax'),
        size_min=__ret__.get('sizeMin'),
        sort=__ret__.get('sort'),
        sort_direction=__ret__.get('sortDirection'),
        sort_key=__ret__.get('sortKey'),
        tag=__ret__.get('tag'),
        visibility=__ret__.get('visibility'))
