# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetVolumeV2Result:
    """
    A collection of values returned by getVolumeV2.
    """
    def __init__(__self__, bootable=None, id=None, metadata=None, name=None, region=None, size=None, source_volume_id=None, status=None, volume_type=None):
        if bootable and not isinstance(bootable, str):
            raise TypeError("Expected argument 'bootable' to be a str")
        __self__.bootable = bootable
        """
        Indicates if the volume is bootable.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        __self__.metadata = metadata
        """
        See Argument Reference above.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        __self__.size = size
        """
        The size of the volume in GBs.
        """
        if source_volume_id and not isinstance(source_volume_id, str):
            raise TypeError("Expected argument 'source_volume_id' to be a str")
        __self__.source_volume_id = source_volume_id
        """
        The ID of the volume from which the current volume was created.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        See Argument Reference above.
        """
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        __self__.volume_type = volume_type
        """
        The type of the volume.
        """
class AwaitableGetVolumeV2Result(GetVolumeV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeV2Result(
            bootable=self.bootable,
            id=self.id,
            metadata=self.metadata,
            name=self.name,
            region=self.region,
            size=self.size,
            source_volume_id=self.source_volume_id,
            status=self.status,
            volume_type=self.volume_type)

def get_volume_v2(bootable=None,metadata=None,name=None,region=None,status=None,volume_type=None,opts=None):
    """
    Use this data source to get information about an existing volume.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/blockstorage_volume_v2.html.markdown.


    :param str bootable: Indicates if the volume is bootable.
    :param dict metadata: Metadata key/value pairs associated with the volume.
    :param str name: The name of the volume.
    :param str region: The region in which to obtain the V2 Block Storage
           client. If omitted, the `region` argument of the provider is used.
    :param str status: The status of the volume.
    :param str volume_type: The type of the volume.
    """
    __args__ = dict()


    __args__['bootable'] = bootable
    __args__['metadata'] = metadata
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['volumeType'] = volume_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:blockstorage/getVolumeV2:getVolumeV2', __args__, opts=opts).value

    return AwaitableGetVolumeV2Result(
        bootable=__ret__.get('bootable'),
        id=__ret__.get('id'),
        metadata=__ret__.get('metadata'),
        name=__ret__.get('name'),
        region=__ret__.get('region'),
        size=__ret__.get('size'),
        source_volume_id=__ret__.get('sourceVolumeId'),
        status=__ret__.get('status'),
        volume_type=__ret__.get('volumeType'))
