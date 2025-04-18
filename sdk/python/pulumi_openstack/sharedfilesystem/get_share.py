# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'GetShareResult',
    'AwaitableGetShareResult',
    'get_share',
    'get_share_output',
]

@pulumi.output_type
class GetShareResult:
    """
    A collection of values returned by getShare.
    """
    def __init__(__self__, availability_zone=None, description=None, export_location_path=None, export_locations=None, id=None, is_public=None, metadata=None, name=None, project_id=None, region=None, share_network_id=None, share_proto=None, size=None, snapshot_id=None, status=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if export_location_path and not isinstance(export_location_path, str):
            raise TypeError("Expected argument 'export_location_path' to be a str")
        pulumi.set(__self__, "export_location_path", export_location_path)
        if export_locations and not isinstance(export_locations, list):
            raise TypeError("Expected argument 'export_locations' to be a list")
        pulumi.set(__self__, "export_locations", export_locations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_public and not isinstance(is_public, bool):
            raise TypeError("Expected argument 'is_public' to be a bool")
        pulumi.set(__self__, "is_public", is_public)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if share_network_id and not isinstance(share_network_id, str):
            raise TypeError("Expected argument 'share_network_id' to be a str")
        pulumi.set(__self__, "share_network_id", share_network_id)
        if share_proto and not isinstance(share_proto, str):
            raise TypeError("Expected argument 'share_proto' to be a str")
        pulumi.set(__self__, "share_proto", share_proto)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> builtins.str:
        """
        The share availability zone.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="exportLocationPath")
    def export_location_path(self) -> Optional[builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "export_location_path")

    @property
    @pulumi.getter(name="exportLocations")
    def export_locations(self) -> Sequence['outputs.GetShareExportLocationResult']:
        """
        A list of export locations. For example, when a share
        server has more than one network interface, it can have multiple export
        locations.
        """
        return pulumi.get(self, "export_locations")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> builtins.bool:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        The region in which to obtain the V2 Shared File System client.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="shareNetworkId")
    def share_network_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "share_network_id")

    @property
    @pulumi.getter(name="shareProto")
    def share_proto(self) -> builtins.str:
        """
        The share protocol.
        """
        return pulumi.get(self, "share_proto")

    @property
    @pulumi.getter
    def size(self) -> builtins.int:
        """
        The share size, in GBs.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "status")


class AwaitableGetShareResult(GetShareResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetShareResult(
            availability_zone=self.availability_zone,
            description=self.description,
            export_location_path=self.export_location_path,
            export_locations=self.export_locations,
            id=self.id,
            is_public=self.is_public,
            metadata=self.metadata,
            name=self.name,
            project_id=self.project_id,
            region=self.region,
            share_network_id=self.share_network_id,
            share_proto=self.share_proto,
            size=self.size,
            snapshot_id=self.snapshot_id,
            status=self.status)


def get_share(description: Optional[builtins.str] = None,
              export_location_path: Optional[builtins.str] = None,
              is_public: Optional[builtins.bool] = None,
              metadata: Optional[Mapping[str, builtins.str]] = None,
              name: Optional[builtins.str] = None,
              region: Optional[builtins.str] = None,
              share_network_id: Optional[builtins.str] = None,
              snapshot_id: Optional[builtins.str] = None,
              status: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetShareResult:
    """
    Use this data source to get the ID of an available Shared File System share.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    share1 = openstack.sharedfilesystem.get_share(name="share_1")
    ```


    :param builtins.str description: The human-readable description for the share.
    :param builtins.str export_location_path: The export location path of the share. Available
           since Manila API version 2.35.
    :param builtins.bool is_public: The level of visibility for the share.
           length.
    :param Mapping[str, builtins.str] metadata: One or more metadata key and value pairs as a dictionary of
           strings.
    :param builtins.str name: The name of the share.
    :param builtins.str region: The region in which to obtain the V2 Shared File System client.
    :param builtins.str share_network_id: The UUID of the share's share network.
    :param builtins.str snapshot_id: The UUID of the share's base snapshot.
    :param builtins.str status: A share status filter. A valid value is `creating`,
           `error`, `available`, `deleting`, `error_deleting`, `manage_starting`,
           `manage_error`, `unmanage_starting`, `unmanage_error`, `unmanaged`,
           `extending`, `extending_error`, `shrinking`, `shrinking_error`, or
           `shrinking_possible_data_loss_error`.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['exportLocationPath'] = export_location_path
    __args__['isPublic'] = is_public
    __args__['metadata'] = metadata
    __args__['name'] = name
    __args__['region'] = region
    __args__['shareNetworkId'] = share_network_id
    __args__['snapshotId'] = snapshot_id
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getShare:getShare', __args__, opts=opts, typ=GetShareResult).value

    return AwaitableGetShareResult(
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        description=pulumi.get(__ret__, 'description'),
        export_location_path=pulumi.get(__ret__, 'export_location_path'),
        export_locations=pulumi.get(__ret__, 'export_locations'),
        id=pulumi.get(__ret__, 'id'),
        is_public=pulumi.get(__ret__, 'is_public'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        share_network_id=pulumi.get(__ret__, 'share_network_id'),
        share_proto=pulumi.get(__ret__, 'share_proto'),
        size=pulumi.get(__ret__, 'size'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        status=pulumi.get(__ret__, 'status'))
def get_share_output(description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     export_location_path: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     is_public: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                     metadata: Optional[pulumi.Input[Optional[Mapping[str, builtins.str]]]] = None,
                     name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     share_network_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     snapshot_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     status: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetShareResult]:
    """
    Use this data source to get the ID of an available Shared File System share.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    share1 = openstack.sharedfilesystem.get_share(name="share_1")
    ```


    :param builtins.str description: The human-readable description for the share.
    :param builtins.str export_location_path: The export location path of the share. Available
           since Manila API version 2.35.
    :param builtins.bool is_public: The level of visibility for the share.
           length.
    :param Mapping[str, builtins.str] metadata: One or more metadata key and value pairs as a dictionary of
           strings.
    :param builtins.str name: The name of the share.
    :param builtins.str region: The region in which to obtain the V2 Shared File System client.
    :param builtins.str share_network_id: The UUID of the share's share network.
    :param builtins.str snapshot_id: The UUID of the share's base snapshot.
    :param builtins.str status: A share status filter. A valid value is `creating`,
           `error`, `available`, `deleting`, `error_deleting`, `manage_starting`,
           `manage_error`, `unmanage_starting`, `unmanage_error`, `unmanaged`,
           `extending`, `extending_error`, `shrinking`, `shrinking_error`, or
           `shrinking_possible_data_loss_error`.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['exportLocationPath'] = export_location_path
    __args__['isPublic'] = is_public
    __args__['metadata'] = metadata
    __args__['name'] = name
    __args__['region'] = region
    __args__['shareNetworkId'] = share_network_id
    __args__['snapshotId'] = snapshot_id
    __args__['status'] = status
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:sharedfilesystem/getShare:getShare', __args__, opts=opts, typ=GetShareResult)
    return __ret__.apply(lambda __response__: GetShareResult(
        availability_zone=pulumi.get(__response__, 'availability_zone'),
        description=pulumi.get(__response__, 'description'),
        export_location_path=pulumi.get(__response__, 'export_location_path'),
        export_locations=pulumi.get(__response__, 'export_locations'),
        id=pulumi.get(__response__, 'id'),
        is_public=pulumi.get(__response__, 'is_public'),
        metadata=pulumi.get(__response__, 'metadata'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        share_network_id=pulumi.get(__response__, 'share_network_id'),
        share_proto=pulumi.get(__response__, 'share_proto'),
        size=pulumi.get(__response__, 'size'),
        snapshot_id=pulumi.get(__response__, 'snapshot_id'),
        status=pulumi.get(__response__, 'status')))
