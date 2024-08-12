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
from ._inputs import *

__all__ = ['VolumeV2Args', 'VolumeV2']

@pulumi.input_type
class VolumeV2Args:
    def __init__(__self__, *,
                 size: pulumi.Input[int],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 consistency_group_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_replica: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VolumeV2 resource.
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] consistency_group_id: The consistency group to place the volume
               in.
        :param pulumi.Input[str] description: A description of the volume. Changing this updates
               the volume's description.
        :param pulumi.Input[str] image_id: The image ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to associate with the volume.
               Changing this updates the existing volume metadata.
        :param pulumi.Input[str] name: A unique name for the volume. Changing this updates the
               volume's name.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new volume.
        :param pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]] scheduler_hints: Provide the Cinder scheduler with hints on where
               to instantiate a volume in the OpenStack cloud. The available hints are described below.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_replica: The volume ID to replicate with.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        pulumi.set(__self__, "size", size)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if consistency_group_id is not None:
            pulumi.set(__self__, "consistency_group_id", consistency_group_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if scheduler_hints is not None:
            pulumi.set(__self__, "scheduler_hints", scheduler_hints)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if source_replica is not None:
            pulumi.set(__self__, "source_replica", source_replica)
        if source_vol_id is not None:
            pulumi.set(__self__, "source_vol_id", source_vol_id)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        The size of the volume to create (in gigabytes). Changing
        this creates a new volume.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The availability zone for the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="consistencyGroupId")
    def consistency_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The consistency group to place the volume
        in.
        """
        return pulumi.get(self, "consistency_group_id")

    @consistency_group_id.setter
    def consistency_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consistency_group_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the volume. Changing this updates
        the volume's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The image ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Metadata key/value pairs to associate with the volume.
        Changing this updates the existing volume metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the volume. Changing this updates the
        volume's name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new volume.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="schedulerHints")
    def scheduler_hints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]]]:
        """
        Provide the Cinder scheduler with hints on where
        to instantiate a volume in the OpenStack cloud. The available hints are described below.
        """
        return pulumi.get(self, "scheduler_hints")

    @scheduler_hints.setter
    def scheduler_hints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]]]):
        pulumi.set(self, "scheduler_hints", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="sourceReplica")
    def source_replica(self) -> Optional[pulumi.Input[str]]:
        """
        The volume ID to replicate with.
        """
        return pulumi.get(self, "source_replica")

    @source_replica.setter
    def source_replica(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_replica", value)

    @property
    @pulumi.getter(name="sourceVolId")
    def source_vol_id(self) -> Optional[pulumi.Input[str]]:
        """
        The volume ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "source_vol_id")

    @source_vol_id.setter
    def source_vol_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_vol_id", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of volume to create.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


@pulumi.input_type
class _VolumeV2State:
    def __init__(__self__, *,
                 attachments: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2AttachmentArgs']]]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 consistency_group_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_replica: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VolumeV2 resources.
        :param pulumi.Input[Sequence[pulumi.Input['VolumeV2AttachmentArgs']]] attachments: If a volume is attached to an instance, this attribute will
               display the Attachment ID, Instance ID, and the Device as the Instance
               sees it.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] consistency_group_id: The consistency group to place the volume
               in.
        :param pulumi.Input[str] description: A description of the volume. Changing this updates
               the volume's description.
        :param pulumi.Input[str] image_id: The image ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to associate with the volume.
               Changing this updates the existing volume metadata.
        :param pulumi.Input[str] name: A unique name for the volume. Changing this updates the
               volume's name.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new volume.
        :param pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]] scheduler_hints: Provide the Cinder scheduler with hints on where
               to instantiate a volume in the OpenStack cloud. The available hints are described below.
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_replica: The volume ID to replicate with.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        if attachments is not None:
            pulumi.set(__self__, "attachments", attachments)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if consistency_group_id is not None:
            pulumi.set(__self__, "consistency_group_id", consistency_group_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if scheduler_hints is not None:
            pulumi.set(__self__, "scheduler_hints", scheduler_hints)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if source_replica is not None:
            pulumi.set(__self__, "source_replica", source_replica)
        if source_vol_id is not None:
            pulumi.set(__self__, "source_vol_id", source_vol_id)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def attachments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2AttachmentArgs']]]]:
        """
        If a volume is attached to an instance, this attribute will
        display the Attachment ID, Instance ID, and the Device as the Instance
        sees it.
        """
        return pulumi.get(self, "attachments")

    @attachments.setter
    def attachments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2AttachmentArgs']]]]):
        pulumi.set(self, "attachments", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The availability zone for the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="consistencyGroupId")
    def consistency_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The consistency group to place the volume
        in.
        """
        return pulumi.get(self, "consistency_group_id")

    @consistency_group_id.setter
    def consistency_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consistency_group_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the volume. Changing this updates
        the volume's description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The image ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Metadata key/value pairs to associate with the volume.
        Changing this updates the existing volume metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the volume. Changing this updates the
        volume's name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new volume.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="schedulerHints")
    def scheduler_hints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]]]:
        """
        Provide the Cinder scheduler with hints on where
        to instantiate a volume in the OpenStack cloud. The available hints are described below.
        """
        return pulumi.get(self, "scheduler_hints")

    @scheduler_hints.setter
    def scheduler_hints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV2SchedulerHintArgs']]]]):
        pulumi.set(self, "scheduler_hints", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the volume to create (in gigabytes). Changing
        this creates a new volume.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter(name="sourceReplica")
    def source_replica(self) -> Optional[pulumi.Input[str]]:
        """
        The volume ID to replicate with.
        """
        return pulumi.get(self, "source_replica")

    @source_replica.setter
    def source_replica(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_replica", value)

    @property
    @pulumi.getter(name="sourceVolId")
    def source_vol_id(self) -> Optional[pulumi.Input[str]]:
        """
        The volume ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "source_vol_id")

    @source_vol_id.setter
    def source_vol_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_vol_id", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of volume to create.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


class VolumeV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 consistency_group_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2SchedulerHintArgs', 'VolumeV2SchedulerHintArgsDict']]]]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_replica: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V2 volume resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        volume1 = openstack.blockstorage.VolumeV2("volume_1",
            region="RegionOne",
            name="volume_1",
            description="first test volume",
            size=3)
        ```

        ## Import

        Volumes can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:blockstorage/volumeV2:VolumeV2 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] consistency_group_id: The consistency group to place the volume
               in.
        :param pulumi.Input[str] description: A description of the volume. Changing this updates
               the volume's description.
        :param pulumi.Input[str] image_id: The image ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to associate with the volume.
               Changing this updates the existing volume metadata.
        :param pulumi.Input[str] name: A unique name for the volume. Changing this updates the
               volume's name.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new volume.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2SchedulerHintArgs', 'VolumeV2SchedulerHintArgsDict']]]] scheduler_hints: Provide the Cinder scheduler with hints on where
               to instantiate a volume in the OpenStack cloud. The available hints are described below.
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_replica: The volume ID to replicate with.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 volume resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        volume1 = openstack.blockstorage.VolumeV2("volume_1",
            region="RegionOne",
            name="volume_1",
            description="first test volume",
            size=3)
        ```

        ## Import

        Volumes can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:blockstorage/volumeV2:VolumeV2 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
        ```

        :param str resource_name: The name of the resource.
        :param VolumeV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 consistency_group_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2SchedulerHintArgs', 'VolumeV2SchedulerHintArgsDict']]]]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_replica: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeV2Args.__new__(VolumeV2Args)

            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["consistency_group_id"] = consistency_group_id
            __props__.__dict__["description"] = description
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["scheduler_hints"] = scheduler_hints
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["source_replica"] = source_replica
            __props__.__dict__["source_vol_id"] = source_vol_id
            __props__.__dict__["volume_type"] = volume_type
            __props__.__dict__["attachments"] = None
        super(VolumeV2, __self__).__init__(
            'openstack:blockstorage/volumeV2:VolumeV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachments: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2AttachmentArgs', 'VolumeV2AttachmentArgsDict']]]]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            consistency_group_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2SchedulerHintArgs', 'VolumeV2SchedulerHintArgsDict']]]]] = None,
            size: Optional[pulumi.Input[int]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None,
            source_replica: Optional[pulumi.Input[str]] = None,
            source_vol_id: Optional[pulumi.Input[str]] = None,
            volume_type: Optional[pulumi.Input[str]] = None) -> 'VolumeV2':
        """
        Get an existing VolumeV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2AttachmentArgs', 'VolumeV2AttachmentArgsDict']]]] attachments: If a volume is attached to an instance, this attribute will
               display the Attachment ID, Instance ID, and the Device as the Instance
               sees it.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] consistency_group_id: The consistency group to place the volume
               in.
        :param pulumi.Input[str] description: A description of the volume. Changing this updates
               the volume's description.
        :param pulumi.Input[str] image_id: The image ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to associate with the volume.
               Changing this updates the existing volume metadata.
        :param pulumi.Input[str] name: A unique name for the volume. Changing this updates the
               volume's name.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new volume.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VolumeV2SchedulerHintArgs', 'VolumeV2SchedulerHintArgsDict']]]] scheduler_hints: Provide the Cinder scheduler with hints on where
               to instantiate a volume in the OpenStack cloud. The available hints are described below.
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_replica: The volume ID to replicate with.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeV2State.__new__(_VolumeV2State)

        __props__.__dict__["attachments"] = attachments
        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["consistency_group_id"] = consistency_group_id
        __props__.__dict__["description"] = description
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["scheduler_hints"] = scheduler_hints
        __props__.__dict__["size"] = size
        __props__.__dict__["snapshot_id"] = snapshot_id
        __props__.__dict__["source_replica"] = source_replica
        __props__.__dict__["source_vol_id"] = source_vol_id
        __props__.__dict__["volume_type"] = volume_type
        return VolumeV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Sequence['outputs.VolumeV2Attachment']]:
        """
        If a volume is attached to an instance, this attribute will
        display the Attachment ID, Instance ID, and the Device as the Instance
        sees it.
        """
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The availability zone for the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="consistencyGroupId")
    def consistency_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The consistency group to place the volume
        in.
        """
        return pulumi.get(self, "consistency_group_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the volume. Changing this updates
        the volume's description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[Optional[str]]:
        """
        The image ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Metadata key/value pairs to associate with the volume.
        Changing this updates the existing volume metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the volume. Changing this updates the
        volume's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new volume.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="schedulerHints")
    def scheduler_hints(self) -> pulumi.Output[Optional[Sequence['outputs.VolumeV2SchedulerHint']]]:
        """
        Provide the Cinder scheduler with hints on where
        to instantiate a volume in the OpenStack cloud. The available hints are described below.
        """
        return pulumi.get(self, "scheduler_hints")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The size of the volume to create (in gigabytes). Changing
        this creates a new volume.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        The snapshot ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="sourceReplica")
    def source_replica(self) -> pulumi.Output[Optional[str]]:
        """
        The volume ID to replicate with.
        """
        return pulumi.get(self, "source_replica")

    @property
    @pulumi.getter(name="sourceVolId")
    def source_vol_id(self) -> pulumi.Output[Optional[str]]:
        """
        The volume ID from which to create the volume.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "source_vol_id")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> pulumi.Output[str]:
        """
        The type of volume to create.
        Changing this creates a new volume.
        """
        return pulumi.get(self, "volume_type")

