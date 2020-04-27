# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, description=None, id=None, name=None, project_id=None, region=None, share_id=None, share_proto=None, share_size=None, size=None, status=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if share_id and not isinstance(share_id, str):
            raise TypeError("Expected argument 'share_id' to be a str")
        __self__.share_id = share_id
        """
        The UUID of the source share that was used to create the snapshot.
        """
        if share_proto and not isinstance(share_proto, str):
            raise TypeError("Expected argument 'share_proto' to be a str")
        __self__.share_proto = share_proto
        """
        The file system protocol of a share snapshot.
        """
        if share_size and not isinstance(share_size, float):
            raise TypeError("Expected argument 'share_size' to be a float")
        __self__.share_size = share_size
        """
        The share snapshot size, in GBs.
        """
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        __self__.size = size
        """
        The snapshot size, in GBs.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        See Argument Reference above.
        """
class AwaitableGetSnapshotResult(GetSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotResult(
            description=self.description,
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            region=self.region,
            share_id=self.share_id,
            share_proto=self.share_proto,
            share_size=self.share_size,
            size=self.size,
            status=self.status)

def get_snapshot(description=None,name=None,region=None,share_id=None,status=None,opts=None):
    """
    Use this data source to get the ID of an available Shared File System snapshot.




    :param str description: The human-readable description of the snapshot.
    :param str name: The name of the snapshot.
    :param str region: The region in which to obtain the V2 Shared File System client.
    :param str share_id: The UUID of the source share that was used to create the snapshot.
    :param str status: A snapshot status filter. A valid value is `available`, `error`,
           `creating`, `deleting`, `manage_starting`, `manage_error`, `unmanage_starting`,
           `unmanage_error` or `error_deleting`.
    """
    __args__ = dict()


    __args__['description'] = description
    __args__['name'] = name
    __args__['region'] = region
    __args__['shareId'] = share_id
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getSnapshot:getSnapshot', __args__, opts=opts).value

    return AwaitableGetSnapshotResult(
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        share_id=__ret__.get('shareId'),
        share_proto=__ret__.get('shareProto'),
        share_size=__ret__.get('shareSize'),
        size=__ret__.get('size'),
        status=__ret__.get('status'))
