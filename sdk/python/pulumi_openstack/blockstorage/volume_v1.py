# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VolumeV1Args', 'VolumeV1']

@pulumi.input_type
class VolumeV1Args:
    def __init__(__self__, *,
                 size: pulumi.Input[int],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VolumeV1 resource.
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
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
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        VolumeV1Args._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            size=size,
            availability_zone=availability_zone,
            description=description,
            image_id=image_id,
            metadata=metadata,
            name=name,
            region=region,
            snapshot_id=snapshot_id,
            source_vol_id=source_vol_id,
            volume_type=volume_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             size: Optional[pulumi.Input[int]] = None,
             availability_zone: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             image_id: Optional[pulumi.Input[str]] = None,
             metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             snapshot_id: Optional[pulumi.Input[str]] = None,
             source_vol_id: Optional[pulumi.Input[str]] = None,
             volume_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if size is None:
            raise TypeError("Missing 'size' argument")
        if availability_zone is None and 'availabilityZone' in kwargs:
            availability_zone = kwargs['availabilityZone']
        if image_id is None and 'imageId' in kwargs:
            image_id = kwargs['imageId']
        if snapshot_id is None and 'snapshotId' in kwargs:
            snapshot_id = kwargs['snapshotId']
        if source_vol_id is None and 'sourceVolId' in kwargs:
            source_vol_id = kwargs['sourceVolId']
        if volume_type is None and 'volumeType' in kwargs:
            volume_type = kwargs['volumeType']

        _setter("size", size)
        if availability_zone is not None:
            _setter("availability_zone", availability_zone)
        if description is not None:
            _setter("description", description)
        if image_id is not None:
            _setter("image_id", image_id)
        if metadata is not None:
            _setter("metadata", metadata)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)
        if snapshot_id is not None:
            _setter("snapshot_id", snapshot_id)
        if source_vol_id is not None:
            _setter("source_vol_id", source_vol_id)
        if volume_type is not None:
            _setter("volume_type", volume_type)

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
class _VolumeV1State:
    def __init__(__self__, *,
                 attachments: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV1AttachmentArgs']]]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VolumeV1 resources.
        :param pulumi.Input[Sequence[pulumi.Input['VolumeV1AttachmentArgs']]] attachments: If a volume is attached to an instance, this attribute will
               display the Attachment ID, Instance ID, and the Device as the Instance
               sees it.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
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
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        _VolumeV1State._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            attachments=attachments,
            availability_zone=availability_zone,
            description=description,
            image_id=image_id,
            metadata=metadata,
            name=name,
            region=region,
            size=size,
            snapshot_id=snapshot_id,
            source_vol_id=source_vol_id,
            volume_type=volume_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             attachments: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV1AttachmentArgs']]]] = None,
             availability_zone: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             image_id: Optional[pulumi.Input[str]] = None,
             metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             size: Optional[pulumi.Input[int]] = None,
             snapshot_id: Optional[pulumi.Input[str]] = None,
             source_vol_id: Optional[pulumi.Input[str]] = None,
             volume_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if availability_zone is None and 'availabilityZone' in kwargs:
            availability_zone = kwargs['availabilityZone']
        if image_id is None and 'imageId' in kwargs:
            image_id = kwargs['imageId']
        if snapshot_id is None and 'snapshotId' in kwargs:
            snapshot_id = kwargs['snapshotId']
        if source_vol_id is None and 'sourceVolId' in kwargs:
            source_vol_id = kwargs['sourceVolId']
        if volume_type is None and 'volumeType' in kwargs:
            volume_type = kwargs['volumeType']

        if attachments is not None:
            _setter("attachments", attachments)
        if availability_zone is not None:
            _setter("availability_zone", availability_zone)
        if description is not None:
            _setter("description", description)
        if image_id is not None:
            _setter("image_id", image_id)
        if metadata is not None:
            _setter("metadata", metadata)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)
        if size is not None:
            _setter("size", size)
        if snapshot_id is not None:
            _setter("snapshot_id", snapshot_id)
        if source_vol_id is not None:
            _setter("source_vol_id", source_vol_id)
        if volume_type is not None:
            _setter("volume_type", volume_type)

    @property
    @pulumi.getter
    def attachments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV1AttachmentArgs']]]]:
        """
        If a volume is attached to an instance, this attribute will
        display the Attachment ID, Instance ID, and the Device as the Instance
        sees it.
        """
        return pulumi.get(self, "attachments")

    @attachments.setter
    def attachments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VolumeV1AttachmentArgs']]]]):
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


class VolumeV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V1 volume resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        volume1 = openstack.blockstorage.VolumeV1("volume1",
            description="first test volume",
            region="RegionOne",
            size=3)
        ```

        ## Import

        Volumes can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/volumeV1:VolumeV1 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
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
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V1 volume resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        volume1 = openstack.blockstorage.VolumeV1("volume1",
            description="first test volume",
            region="RegionOne",
            size=3)
        ```

        ## Import

        Volumes can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/volumeV1:VolumeV1 volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d
        ```

        :param str resource_name: The name of the resource.
        :param VolumeV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VolumeV1Args._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 source_vol_id: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeV1Args.__new__(VolumeV1Args)

            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["description"] = description
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["source_vol_id"] = source_vol_id
            __props__.__dict__["volume_type"] = volume_type
            __props__.__dict__["attachments"] = None
        super(VolumeV1, __self__).__init__(
            'openstack:blockstorage/volumeV1:VolumeV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeV1AttachmentArgs']]]]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None,
            source_vol_id: Optional[pulumi.Input[str]] = None,
            volume_type: Optional[pulumi.Input[str]] = None) -> 'VolumeV1':
        """
        Get an existing VolumeV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeV1AttachmentArgs']]]] attachments: If a volume is attached to an instance, this attribute will
               display the Attachment ID, Instance ID, and the Device as the Instance
               sees it.
        :param pulumi.Input[str] availability_zone: The availability zone for the volume.
               Changing this creates a new volume.
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
        :param pulumi.Input[int] size: The size of the volume to create (in gigabytes). Changing
               this creates a new volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] source_vol_id: The volume ID from which to create the volume.
               Changing this creates a new volume.
        :param pulumi.Input[str] volume_type: The type of volume to create.
               Changing this creates a new volume.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeV1State.__new__(_VolumeV1State)

        __props__.__dict__["attachments"] = attachments
        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["description"] = description
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["size"] = size
        __props__.__dict__["snapshot_id"] = snapshot_id
        __props__.__dict__["source_vol_id"] = source_vol_id
        __props__.__dict__["volume_type"] = volume_type
        return VolumeV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Sequence['outputs.VolumeV1Attachment']]:
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

