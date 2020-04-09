# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetImageResult:
    """
    A collection of values returned by getImage.
    """
    def __init__(__self__, checksum=None, container_format=None, created_at=None, disk_format=None, file=None, id=None, member_status=None, metadata=None, min_disk_gb=None, min_ram_mb=None, most_recent=None, name=None, owner=None, properties=None, protected=None, region=None, schema=None, size_bytes=None, size_max=None, size_min=None, sort_direction=None, sort_key=None, tag=None, tags=None, updated_at=None, visibility=None):
        if checksum and not isinstance(checksum, str):
            raise TypeError("Expected argument 'checksum' to be a str")
        __self__.checksum = checksum
        """
        The checksum of the data associated with the image.
        """
        if container_format and not isinstance(container_format, str):
            raise TypeError("Expected argument 'container_format' to be a str")
        __self__.container_format = container_format
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The date the image was created.
        * `container_format`: The format of the image's container.
        * `disk_format`: The format of the image's disk.
        """
        if disk_format and not isinstance(disk_format, str):
            raise TypeError("Expected argument 'disk_format' to be a str")
        __self__.disk_format = disk_format
        if file and not isinstance(file, str):
            raise TypeError("Expected argument 'file' to be a str")
        __self__.file = file
        """
        the trailing path after the glance endpoint that represent the
        location of the image or the path to retrieve it.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if member_status and not isinstance(member_status, str):
            raise TypeError("Expected argument 'member_status' to be a str")
        __self__.member_status = member_status
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        __self__.metadata = metadata
        """
        The metadata associated with the image.
        Image metadata allow for meaningfully define the image properties
        and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.
        """
        if min_disk_gb and not isinstance(min_disk_gb, float):
            raise TypeError("Expected argument 'min_disk_gb' to be a float")
        __self__.min_disk_gb = min_disk_gb
        """
        The minimum amount of disk space required to use the image.
        """
        if min_ram_mb and not isinstance(min_ram_mb, float):
            raise TypeError("Expected argument 'min_ram_mb' to be a float")
        __self__.min_ram_mb = min_ram_mb
        """
        The minimum amount of ram required to use the image.
        """
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        __self__.most_recent = most_recent
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        __self__.owner = owner
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Freeform information about the image.
        """
        if protected and not isinstance(protected, bool):
            raise TypeError("Expected argument 'protected' to be a bool")
        __self__.protected = protected
        """
        Whether or not the image is protected.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if schema and not isinstance(schema, str):
            raise TypeError("Expected argument 'schema' to be a str")
        __self__.schema = schema
        """
        The path to the JSON-schema that represent
        the image or image
        """
        if size_bytes and not isinstance(size_bytes, float):
            raise TypeError("Expected argument 'size_bytes' to be a float")
        __self__.size_bytes = size_bytes
        """
        The size of the image (in bytes).
        """
        if size_max and not isinstance(size_max, float):
            raise TypeError("Expected argument 'size_max' to be a float")
        __self__.size_max = size_max
        if size_min and not isinstance(size_min, float):
            raise TypeError("Expected argument 'size_min' to be a float")
        __self__.size_min = size_min
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        __self__.sort_direction = sort_direction
        if sort_key and not isinstance(sort_key, str):
            raise TypeError("Expected argument 'sort_key' to be a str")
        __self__.sort_key = sort_key
        if tag and not isinstance(tag, str):
            raise TypeError("Expected argument 'tag' to be a str")
        __self__.tag = tag
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        """
        The tags list of the image.
        """
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        __self__.updated_at = updated_at
        """
        The date the image was last updated.
        """
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        __self__.visibility = visibility
class AwaitableGetImageResult(GetImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageResult(
            checksum=self.checksum,
            container_format=self.container_format,
            created_at=self.created_at,
            disk_format=self.disk_format,
            file=self.file,
            id=self.id,
            member_status=self.member_status,
            metadata=self.metadata,
            min_disk_gb=self.min_disk_gb,
            min_ram_mb=self.min_ram_mb,
            most_recent=self.most_recent,
            name=self.name,
            owner=self.owner,
            properties=self.properties,
            protected=self.protected,
            region=self.region,
            schema=self.schema,
            size_bytes=self.size_bytes,
            size_max=self.size_max,
            size_min=self.size_min,
            sort_direction=self.sort_direction,
            sort_key=self.sort_key,
            tag=self.tag,
            tags=self.tags,
            updated_at=self.updated_at,
            visibility=self.visibility)

def get_image(member_status=None,most_recent=None,name=None,owner=None,properties=None,region=None,size_max=None,size_min=None,sort_direction=None,sort_key=None,tag=None,visibility=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack image.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/images_image_v2.html.markdown.


    :param str member_status: The status of the image. Must be one of
           "accepted", "pending", "rejected", or "all".
    :param bool most_recent: If more than one result is returned, use the most
           recent image.
    :param str name: The name of the image.
    :param str owner: The owner (UUID) of the image.
    :param dict properties: a map of key/value pairs to match an image with.
           All specified properties must be matched.
    :param str region: The region in which to obtain the V2 Glance client.
           A Glance client is needed to create an Image that can be used with
           a compute instance. If omitted, the `region` argument of the provider
           is used.
    :param float size_max: The maximum size (in bytes) of the image to return.
    :param float size_min: The minimum size (in bytes) of the image to return.
    :param str sort_direction: Order the results in either `asc` or `desc`.
    :param str sort_key: Sort images based on a certain key. Defaults to `name`.
    :param str tag: Search for images with a specific tag.
    :param str visibility: The visibility of the image. Must be one of
           "public", "private", "community", or "shared". Defaults to "private".
    """
    __args__ = dict()


    __args__['memberStatus'] = member_status
    __args__['mostRecent'] = most_recent
    __args__['name'] = name
    __args__['owner'] = owner
    __args__['properties'] = properties
    __args__['region'] = region
    __args__['sizeMax'] = size_max
    __args__['sizeMin'] = size_min
    __args__['sortDirection'] = sort_direction
    __args__['sortKey'] = sort_key
    __args__['tag'] = tag
    __args__['visibility'] = visibility
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:images/getImage:getImage', __args__, opts=opts).value

    return AwaitableGetImageResult(
        checksum=__ret__.get('checksum'),
        container_format=__ret__.get('containerFormat'),
        created_at=__ret__.get('createdAt'),
        disk_format=__ret__.get('diskFormat'),
        file=__ret__.get('file'),
        id=__ret__.get('id'),
        member_status=__ret__.get('memberStatus'),
        metadata=__ret__.get('metadata'),
        min_disk_gb=__ret__.get('minDiskGb'),
        min_ram_mb=__ret__.get('minRamMb'),
        most_recent=__ret__.get('mostRecent'),
        name=__ret__.get('name'),
        owner=__ret__.get('owner'),
        properties=__ret__.get('properties'),
        protected=__ret__.get('protected'),
        region=__ret__.get('region'),
        schema=__ret__.get('schema'),
        size_bytes=__ret__.get('sizeBytes'),
        size_max=__ret__.get('sizeMax'),
        size_min=__ret__.get('sizeMin'),
        sort_direction=__ret__.get('sortDirection'),
        sort_key=__ret__.get('sortKey'),
        tag=__ret__.get('tag'),
        tags=__ret__.get('tags'),
        updated_at=__ret__.get('updatedAt'),
        visibility=__ret__.get('visibility'))
