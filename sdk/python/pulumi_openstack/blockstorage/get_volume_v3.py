# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetVolumeV3Result',
    'AwaitableGetVolumeV3Result',
    'get_volume_v3',
    'get_volume_v3_output',
]

@pulumi.output_type
class GetVolumeV3Result:
    """
    A collection of values returned by getVolumeV3.
    """
    def __init__(__self__, attachments=None, bootable=None, host=None, id=None, metadata=None, multiattach=None, name=None, region=None, size=None, source_volume_id=None, status=None, volume_type=None):
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        pulumi.set(__self__, "attachments", attachments)
        if bootable and not isinstance(bootable, str):
            raise TypeError("Expected argument 'bootable' to be a str")
        pulumi.set(__self__, "bootable", bootable)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if multiattach and not isinstance(multiattach, bool):
            raise TypeError("Expected argument 'multiattach' to be a bool")
        pulumi.set(__self__, "multiattach", multiattach)
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
    def attachments(self) -> Sequence['outputs.GetVolumeV3AttachmentResult']:
        """
        If a volume is attached to an instance, this attribute will
        display the Attachment ID, Instance ID, and the Device as the Instance
        sees it.
        """
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def bootable(self) -> str:
        """
        Indicates if the volume is bootable.
        """
        return pulumi.get(self, "bootable")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The OpenStack host on which the volume is located.
        """
        return pulumi.get(self, "host")

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
    def multiattach(self) -> bool:
        """
        Indicates if the volume can be attached to more then one server.
        """
        return pulumi.get(self, "multiattach")

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


class AwaitableGetVolumeV3Result(GetVolumeV3Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeV3Result(
            attachments=self.attachments,
            bootable=self.bootable,
            host=self.host,
            id=self.id,
            metadata=self.metadata,
            multiattach=self.multiattach,
            name=self.name,
            region=self.region,
            size=self.size,
            source_volume_id=self.source_volume_id,
            status=self.status,
            volume_type=self.volume_type)


def get_volume_v3(bootable: Optional[str] = None,
                  host: Optional[str] = None,
                  metadata: Optional[Mapping[str, Any]] = None,
                  name: Optional[str] = None,
                  region: Optional[str] = None,
                  status: Optional[str] = None,
                  volume_type: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeV3Result:
    """
    Use this data source to get information about an existing volume.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_openstack as openstack

    volume1 = openstack.blockstorage.get_volume_v3(name="volume_1")
    ```
    <!--End PulumiCodeChooser -->


    :param str bootable: Indicates if the volume is bootable.
    :param str host: The OpenStack host on which the volume is located.
    :param Mapping[str, Any] metadata: Metadata key/value pairs associated with the volume.
    :param str name: The name of the volume.
    :param str region: The region in which to obtain the V3 Block Storage
           client. If omitted, the `region` argument of the provider is used.
    :param str status: The status of the volume.
    :param str volume_type: The type of the volume.
    """
    __args__ = dict()
    __args__['bootable'] = bootable
    __args__['host'] = host
    __args__['metadata'] = metadata
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['volumeType'] = volume_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:blockstorage/getVolumeV3:getVolumeV3', __args__, opts=opts, typ=GetVolumeV3Result).value

    return AwaitableGetVolumeV3Result(
        attachments=pulumi.get(__ret__, 'attachments'),
        bootable=pulumi.get(__ret__, 'bootable'),
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        metadata=pulumi.get(__ret__, 'metadata'),
        multiattach=pulumi.get(__ret__, 'multiattach'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        size=pulumi.get(__ret__, 'size'),
        source_volume_id=pulumi.get(__ret__, 'source_volume_id'),
        status=pulumi.get(__ret__, 'status'),
        volume_type=pulumi.get(__ret__, 'volume_type'))


@_utilities.lift_output_func(get_volume_v3)
def get_volume_v3_output(bootable: Optional[pulumi.Input[Optional[str]]] = None,
                         host: Optional[pulumi.Input[Optional[str]]] = None,
                         metadata: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         status: Optional[pulumi.Input[Optional[str]]] = None,
                         volume_type: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVolumeV3Result]:
    """
    Use this data source to get information about an existing volume.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_openstack as openstack

    volume1 = openstack.blockstorage.get_volume_v3(name="volume_1")
    ```
    <!--End PulumiCodeChooser -->


    :param str bootable: Indicates if the volume is bootable.
    :param str host: The OpenStack host on which the volume is located.
    :param Mapping[str, Any] metadata: Metadata key/value pairs associated with the volume.
    :param str name: The name of the volume.
    :param str region: The region in which to obtain the V3 Block Storage
           client. If omitted, the `region` argument of the provider is used.
    :param str status: The status of the volume.
    :param str volume_type: The type of the volume.
    """
    ...
