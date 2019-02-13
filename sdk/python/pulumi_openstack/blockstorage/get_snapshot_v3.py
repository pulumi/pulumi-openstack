# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSnapshotV3Result(object):
    """
    A collection of values returned by getSnapshotV3.
    """
    def __init__(__self__, description=None, metadata=None, name=None, region=None, size=None, status=None, volume_id=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError('Expected argument description to be a str')
        __self__.description = description
        """
        The snapshot's description.
        """
        if metadata and not isinstance(metadata, dict):
            raise TypeError('Expected argument metadata to be a dict')
        __self__.metadata = metadata
        """
        The snapshot's metadata.
        """
        if name and not isinstance(name, str):
            raise TypeError('Expected argument name to be a str')
        __self__.name = name
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        """
        See Argument Reference above.
        """
        if size and not isinstance(size, int):
            raise TypeError('Expected argument size to be a int')
        __self__.size = size
        """
        The size of the snapshot.
        """
        if status and not isinstance(status, str):
            raise TypeError('Expected argument status to be a str')
        __self__.status = status
        """
        See Argument Reference above.
        """
        if volume_id and not isinstance(volume_id, str):
            raise TypeError('Expected argument volume_id to be a str')
        __self__.volume_id = volume_id
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_snapshot_v3(most_recent=None, name=None, region=None, status=None, volume_id=None):
    """
    Use this data source to get information about an existing snapshot.
    """
    __args__ = dict()

    __args__['mostRecent'] = most_recent
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['volumeId'] = volume_id
    __ret__ = await pulumi.runtime.invoke('openstack:blockstorage/getSnapshotV3:getSnapshotV3', __args__)

    return GetSnapshotV3Result(
        description=__ret__.get('description'),
        metadata=__ret__.get('metadata'),
        name=__ret__.get('name'),
        region=__ret__.get('region'),
        size=__ret__.get('size'),
        status=__ret__.get('status'),
        volume_id=__ret__.get('volumeId'),
        id=__ret__.get('id'))
