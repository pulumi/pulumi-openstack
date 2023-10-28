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
    'GetSnapshotResult',
    'AwaitableGetSnapshotResult',
    'get_snapshot',
    'get_snapshot_output',
]

@pulumi.output_type
class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, description=None, id=None, name=None, project_id=None, region=None, share_id=None, share_proto=None, share_size=None, size=None, status=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if share_id and not isinstance(share_id, str):
            raise TypeError("Expected argument 'share_id' to be a str")
        pulumi.set(__self__, "share_id", share_id)
        if share_proto and not isinstance(share_proto, str):
            raise TypeError("Expected argument 'share_proto' to be a str")
        pulumi.set(__self__, "share_proto", share_proto)
        if share_size and not isinstance(share_size, int):
            raise TypeError("Expected argument 'share_size' to be a int")
        pulumi.set(__self__, "share_size", share_size)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        See Argument Reference above.
        """
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
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> str:
        """
        The UUID of the source share that was used to create the snapshot.
        """
        return pulumi.get(self, "share_id")

    @property
    @pulumi.getter(name="shareProto")
    def share_proto(self) -> str:
        """
        The file system protocol of a share snapshot.
        """
        return pulumi.get(self, "share_proto")

    @property
    @pulumi.getter(name="shareSize")
    def share_size(self) -> int:
        """
        The share snapshot size, in GBs.
        """
        return pulumi.get(self, "share_size")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The snapshot size, in GBs.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "status")


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


def get_snapshot(description: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 share_id: Optional[str] = None,
                 status: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotResult:
    """
    Use this data source to get the ID of an available Shared File System snapshot.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    snapshot1 = openstack.sharedfilesystem.get_snapshot(name="snapshot_1")
    ```


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
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getSnapshot:getSnapshot', __args__, opts=opts, typ=GetSnapshotResult).value

    return AwaitableGetSnapshotResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        share_id=pulumi.get(__ret__, 'share_id'),
        share_proto=pulumi.get(__ret__, 'share_proto'),
        share_size=pulumi.get(__ret__, 'share_size'),
        size=pulumi.get(__ret__, 'size'),
        status=pulumi.get(__ret__, 'status'))


@_utilities.lift_output_func(get_snapshot)
def get_snapshot_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        share_id: Optional[pulumi.Input[Optional[str]]] = None,
                        status: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSnapshotResult]:
    """
    Use this data source to get the ID of an available Shared File System snapshot.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    snapshot1 = openstack.sharedfilesystem.get_snapshot(name="snapshot_1")
    ```


    :param str description: The human-readable description of the snapshot.
    :param str name: The name of the snapshot.
    :param str region: The region in which to obtain the V2 Shared File System client.
    :param str share_id: The UUID of the source share that was used to create the snapshot.
    :param str status: A snapshot status filter. A valid value is `available`, `error`,
           `creating`, `deleting`, `manage_starting`, `manage_error`, `unmanage_starting`,
           `unmanage_error` or `error_deleting`.
    """
    ...
