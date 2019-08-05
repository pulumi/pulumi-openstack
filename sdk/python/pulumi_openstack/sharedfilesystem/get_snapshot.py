# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, description=None, name=None, project_id=None, region=None, share_id=None, share_proto=None, share_size=None, size=None, status=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        See Argument Reference above.
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
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_snapshot(description=None,name=None,region=None,share_id=None,status=None,opts=None):
    """
    Use this data source to get the ID of an available Shared File System snapshot.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/sharedfilesystem_snapshot_v2.html.markdown.
    """
    __args__ = dict()

    __args__['description'] = description
    __args__['name'] = name
    __args__['region'] = region
    __args__['shareId'] = share_id
    __args__['status'] = status
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getSnapshot:getSnapshot', __args__, opts=opts).value

    return GetSnapshotResult(
        description=__ret__.get('description'),
        name=__ret__.get('name'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        share_id=__ret__.get('shareId'),
        share_proto=__ret__.get('shareProto'),
        share_size=__ret__.get('shareSize'),
        size=__ret__.get('size'),
        status=__ret__.get('status'),
        id=__ret__.get('id'))
