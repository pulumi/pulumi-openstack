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
    'GetVolumeV2Result',
    'AwaitableGetVolumeV2Result',
    'get_volume_v2',
    'get_volume_v2_output',
]

@pulumi.output_type
class GetVolumeV2Result:
    """
    A collection of values returned by getVolumeV2.
    """
    def __init__(__self__, bootable=None, id=None, metadata=None, name=None, region=None, size=None, source_volume_id=None, status=None, volume_type=None):
        if bootable and not isinstance(bootable, str):
            raise TypeError("Expected argument 'bootable' to be a str")
        pulumi.set(__self__, "bootable", bootable)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if source_volume_id and not isinstance(source_volume_id, str):
            raise TypeError("Expected argument 'source_volume_id' to be a str")
        pulumi.set(__self__, "source_volume_id", source_volume_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def bootable(self) -> str:
        """
        Indicates if the volume is bootable.
        """
        return pulumi.get(self, "bootable")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, Any]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of the volume in GBs.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="sourceVolumeId")
    def source_volume_id(self) -> str:
        """
        The ID of the volume from which the current volume was created.
        """
        return pulumi.get(self, "source_volume_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        """
        The type of the volume.
        """
        return pulumi.get(self, "volume_type")


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


def get_volume_v2(bootable: Optional[str] = None,
                  metadata: Optional[Mapping[str, Any]] = None,
                  name: Optional[str] = None,
                  region: Optional[str] = None,
                  status: Optional[str] = None,
                  volume_type: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeV2Result:
    """
    Use this data source to get information about an existing volume.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_openstack as openstack

    volume1 = openstack.blockstorage.get_volume_v2(name="volume_1")
    ```
    <!--End PulumiCodeChooser -->


    :param str bootable: Indicates if the volume is bootable.
    :param Mapping[str, Any] metadata: Metadata key/value pairs associated with the volume.
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
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:blockstorage/getVolumeV2:getVolumeV2', __args__, opts=opts, typ=GetVolumeV2Result).value

    return AwaitableGetVolumeV2Result(
        bootable=pulumi.get(__ret__, 'bootable'),
        id=pulumi.get(__ret__, 'id'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        size=pulumi.get(__ret__, 'size'),
        source_volume_id=pulumi.get(__ret__, 'source_volume_id'),
        status=pulumi.get(__ret__, 'status'),
        volume_type=pulumi.get(__ret__, 'volume_type'))


@_utilities.lift_output_func(get_volume_v2)
def get_volume_v2_output(bootable: Optional[pulumi.Input[Optional[str]]] = None,
                         metadata: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         volume_type: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVolumeV2Result]:
    """
    Use this data source to get information about an existing volume.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_openstack as openstack

    volume1 = openstack.blockstorage.get_volume_v2(name="volume_1")
    ```
    <!--End PulumiCodeChooser -->


    :param str bootable: Indicates if the volume is bootable.
    :param Mapping[str, Any] metadata: Metadata key/value pairs associated with the volume.
    :param str name: The name of the volume.
    :param str region: The region in which to obtain the V2 Block Storage
           client. If omitted, the `region` argument of the provider is used.
    :param str status: The status of the volume.
    :param str volume_type: The type of the volume.
    """
    ...
