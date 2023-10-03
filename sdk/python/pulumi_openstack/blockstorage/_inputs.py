# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'VolumeAttachmentArgs',
    'VolumeSchedulerHintArgs',
    'VolumeV1AttachmentArgs',
    'VolumeV2AttachmentArgs',
    'VolumeV2SchedulerHintArgs',
]

@pulumi.input_type
class VolumeAttachmentArgs:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        VolumeAttachmentArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            device=device,
            id=id,
            instance_id=instance_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             device: Optional[pulumi.Input[str]] = None,
             id: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if device is not None:
            _setter("device", device)
        if id is not None:
            _setter("id", id)
        if instance_id is not None:
            _setter("instance_id", instance_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class VolumeSchedulerHintArgs:
    def __init__(__self__, *,
                 additional_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 different_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 local_to_instance: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 same_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, Any]] additional_properties: Arbitrary key/value pairs of additional
               properties to pass to the scheduler.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] different_hosts: The volume should be scheduled on a 
               different host from the set of volumes specified in the list provided.
        :param pulumi.Input[str] local_to_instance: An instance UUID. The volume should be 
               scheduled on the same host as the instance.
        :param pulumi.Input[str] query: A conditional query that a back-end must pass in
               order to host a volume. The query must use the `JsonFilter` syntax
               which is described
               [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
               At this time, only simple queries are supported. Compound queries using
               `and`, `or`, or `not` are not supported. An example of a simple query is:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] same_hosts: A list of volume UUIDs. The volume should be
               scheduled on the same host as another volume specified in the list provided.
        """
        VolumeSchedulerHintArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            additional_properties=additional_properties,
            different_hosts=different_hosts,
            local_to_instance=local_to_instance,
            query=query,
            same_hosts=same_hosts,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             additional_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             different_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             local_to_instance: Optional[pulumi.Input[str]] = None,
             query: Optional[pulumi.Input[str]] = None,
             same_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if additional_properties is not None:
            _setter("additional_properties", additional_properties)
        if different_hosts is not None:
            _setter("different_hosts", different_hosts)
        if local_to_instance is not None:
            _setter("local_to_instance", local_to_instance)
        if query is not None:
            _setter("query", query)
        if same_hosts is not None:
            _setter("same_hosts", same_hosts)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary key/value pairs of additional
        properties to pass to the scheduler.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter(name="differentHosts")
    def different_hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The volume should be scheduled on a 
        different host from the set of volumes specified in the list provided.
        """
        return pulumi.get(self, "different_hosts")

    @different_hosts.setter
    def different_hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "different_hosts", value)

    @property
    @pulumi.getter(name="localToInstance")
    def local_to_instance(self) -> Optional[pulumi.Input[str]]:
        """
        An instance UUID. The volume should be 
        scheduled on the same host as the instance.
        """
        return pulumi.get(self, "local_to_instance")

    @local_to_instance.setter
    def local_to_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_to_instance", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        """
        A conditional query that a back-end must pass in
        order to host a volume. The query must use the `JsonFilter` syntax
        which is described
        [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
        At this time, only simple queries are supported. Compound queries using
        `and`, `or`, or `not` are not supported. An example of a simple query is:
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter(name="sameHosts")
    def same_hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of volume UUIDs. The volume should be
        scheduled on the same host as another volume specified in the list provided.
        """
        return pulumi.get(self, "same_hosts")

    @same_hosts.setter
    def same_hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "same_hosts", value)


@pulumi.input_type
class VolumeV1AttachmentArgs:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        VolumeV1AttachmentArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            device=device,
            id=id,
            instance_id=instance_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             device: Optional[pulumi.Input[str]] = None,
             id: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if device is not None:
            _setter("device", device)
        if id is not None:
            _setter("id", id)
        if instance_id is not None:
            _setter("instance_id", instance_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class VolumeV2AttachmentArgs:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        VolumeV2AttachmentArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            device=device,
            id=id,
            instance_id=instance_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             device: Optional[pulumi.Input[str]] = None,
             id: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if device is not None:
            _setter("device", device)
        if id is not None:
            _setter("id", id)
        if instance_id is not None:
            _setter("instance_id", instance_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class VolumeV2SchedulerHintArgs:
    def __init__(__self__, *,
                 additional_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 different_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 local_to_instance: Optional[pulumi.Input[str]] = None,
                 query: Optional[pulumi.Input[str]] = None,
                 same_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, Any]] additional_properties: Arbitrary key/value pairs of additional
               properties to pass to the scheduler.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] different_hosts: The volume should be scheduled on a 
               different host from the set of volumes specified in the list provided.
        :param pulumi.Input[str] local_to_instance: An instance UUID. The volume should be 
               scheduled on the same host as the instance.
        :param pulumi.Input[str] query: A conditional query that a back-end must pass in
               order to host a volume. The query must use the `JsonFilter` syntax
               which is described
               [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
               At this time, only simple queries are supported. Compound queries using
               `and`, `or`, or `not` are not supported. An example of a simple query is:
        :param pulumi.Input[Sequence[pulumi.Input[str]]] same_hosts: A list of volume UUIDs. The volume should be
               scheduled on the same host as another volume specified in the list provided.
        """
        VolumeV2SchedulerHintArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            additional_properties=additional_properties,
            different_hosts=different_hosts,
            local_to_instance=local_to_instance,
            query=query,
            same_hosts=same_hosts,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             additional_properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             different_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             local_to_instance: Optional[pulumi.Input[str]] = None,
             query: Optional[pulumi.Input[str]] = None,
             same_hosts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if additional_properties is not None:
            _setter("additional_properties", additional_properties)
        if different_hosts is not None:
            _setter("different_hosts", different_hosts)
        if local_to_instance is not None:
            _setter("local_to_instance", local_to_instance)
        if query is not None:
            _setter("query", query)
        if same_hosts is not None:
            _setter("same_hosts", same_hosts)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary key/value pairs of additional
        properties to pass to the scheduler.
        """
        return pulumi.get(self, "additional_properties")

    @additional_properties.setter
    def additional_properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "additional_properties", value)

    @property
    @pulumi.getter(name="differentHosts")
    def different_hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The volume should be scheduled on a 
        different host from the set of volumes specified in the list provided.
        """
        return pulumi.get(self, "different_hosts")

    @different_hosts.setter
    def different_hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "different_hosts", value)

    @property
    @pulumi.getter(name="localToInstance")
    def local_to_instance(self) -> Optional[pulumi.Input[str]]:
        """
        An instance UUID. The volume should be 
        scheduled on the same host as the instance.
        """
        return pulumi.get(self, "local_to_instance")

    @local_to_instance.setter
    def local_to_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_to_instance", value)

    @property
    @pulumi.getter
    def query(self) -> Optional[pulumi.Input[str]]:
        """
        A conditional query that a back-end must pass in
        order to host a volume. The query must use the `JsonFilter` syntax
        which is described
        [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
        At this time, only simple queries are supported. Compound queries using
        `and`, `or`, or `not` are not supported. An example of a simple query is:
        """
        return pulumi.get(self, "query")

    @query.setter
    def query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query", value)

    @property
    @pulumi.getter(name="sameHosts")
    def same_hosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of volume UUIDs. The volume should be
        scheduled on the same host as another volume specified in the list provided.
        """
        return pulumi.get(self, "same_hosts")

    @same_hosts.setter
    def same_hosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "same_hosts", value)


