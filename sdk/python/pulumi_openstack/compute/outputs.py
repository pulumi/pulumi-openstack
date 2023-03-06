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
    'InstanceBlockDevice',
    'InstanceNetwork',
    'InstancePersonality',
    'InstanceSchedulerHint',
    'InstanceVendorOptions',
    'InstanceVolume',
    'SecGroupRule',
    'VolumeAttachVendorOptions',
    'GetInstanceV2NetworkResult',
]

@pulumi.output_type
class InstanceBlockDevice(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "sourceType":
            suggest = "source_type"
        elif key == "bootIndex":
            suggest = "boot_index"
        elif key == "deleteOnTermination":
            suggest = "delete_on_termination"
        elif key == "destinationType":
            suggest = "destination_type"
        elif key == "deviceType":
            suggest = "device_type"
        elif key == "diskBus":
            suggest = "disk_bus"
        elif key == "guestFormat":
            suggest = "guest_format"
        elif key == "volumeSize":
            suggest = "volume_size"
        elif key == "volumeType":
            suggest = "volume_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceBlockDevice. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceBlockDevice.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceBlockDevice.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 source_type: str,
                 boot_index: Optional[int] = None,
                 delete_on_termination: Optional[bool] = None,
                 destination_type: Optional[str] = None,
                 device_type: Optional[str] = None,
                 disk_bus: Optional[str] = None,
                 guest_format: Optional[str] = None,
                 uuid: Optional[str] = None,
                 volume_size: Optional[int] = None,
                 volume_type: Optional[str] = None):
        """
        :param str source_type: The source type of the device. Must be one of
               "blank", "image", "volume", or "snapshot". Changing this creates a new
               server.
        :param int boot_index: The boot index of the volume. It defaults to 0.
               Changing this creates a new server.
        :param bool delete_on_termination: Delete the volume / block device upon
               termination of the instance. Defaults to false. Changing this creates a
               new server.
        :param str destination_type: The type that gets created. Possible values
               are "volume" and "local". Changing this creates a new server.
        :param str device_type: The low-level device type that will be used. Most
               common thing is to leave this empty. Changing this creates a new server.
        :param str disk_bus: The low-level disk bus that will be used. Most common
               thing is to leave this empty. Changing this creates a new server.
        :param str guest_format: Specifies the guest server disk file system format,
               such as `ext2`, `ext3`, `ext4`, `xfs` or `swap`. Swap block device mappings
               have the following restrictions: source_type must be blank and destination_type
               must be local and only one swap disk per server and the size of the swap disk
               must be less than or equal to the swap size of the flavor. Changing this
               creates a new server.
        :param str uuid: The UUID of
               the image, volume, or snapshot. Changing this creates a new server.
        :param int volume_size: The size of the volume to create (in gigabytes). Required
               in the following combinations: source=image and destination=volume,
               source=blank and destination=local, and source=blank and destination=volume.
               Changing this creates a new server.
        :param str volume_type: The volume type that will be used, for example SSD
               or HDD storage. The available options depend on how your specific OpenStack
               cloud is configured and what classes of storage are provided. Changing this
               creates a new server.
        """
        pulumi.set(__self__, "source_type", source_type)
        if boot_index is not None:
            pulumi.set(__self__, "boot_index", boot_index)
        if delete_on_termination is not None:
            pulumi.set(__self__, "delete_on_termination", delete_on_termination)
        if destination_type is not None:
            pulumi.set(__self__, "destination_type", destination_type)
        if device_type is not None:
            pulumi.set(__self__, "device_type", device_type)
        if disk_bus is not None:
            pulumi.set(__self__, "disk_bus", disk_bus)
        if guest_format is not None:
            pulumi.set(__self__, "guest_format", guest_format)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="sourceType")
    def source_type(self) -> str:
        """
        The source type of the device. Must be one of
        "blank", "image", "volume", or "snapshot". Changing this creates a new
        server.
        """
        return pulumi.get(self, "source_type")

    @property
    @pulumi.getter(name="bootIndex")
    def boot_index(self) -> Optional[int]:
        """
        The boot index of the volume. It defaults to 0.
        Changing this creates a new server.
        """
        return pulumi.get(self, "boot_index")

    @property
    @pulumi.getter(name="deleteOnTermination")
    def delete_on_termination(self) -> Optional[bool]:
        """
        Delete the volume / block device upon
        termination of the instance. Defaults to false. Changing this creates a
        new server.
        """
        return pulumi.get(self, "delete_on_termination")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> Optional[str]:
        """
        The type that gets created. Possible values
        are "volume" and "local". Changing this creates a new server.
        """
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter(name="deviceType")
    def device_type(self) -> Optional[str]:
        """
        The low-level device type that will be used. Most
        common thing is to leave this empty. Changing this creates a new server.
        """
        return pulumi.get(self, "device_type")

    @property
    @pulumi.getter(name="diskBus")
    def disk_bus(self) -> Optional[str]:
        """
        The low-level disk bus that will be used. Most common
        thing is to leave this empty. Changing this creates a new server.
        """
        return pulumi.get(self, "disk_bus")

    @property
    @pulumi.getter(name="guestFormat")
    def guest_format(self) -> Optional[str]:
        """
        Specifies the guest server disk file system format,
        such as `ext2`, `ext3`, `ext4`, `xfs` or `swap`. Swap block device mappings
        have the following restrictions: source_type must be blank and destination_type
        must be local and only one swap disk per server and the size of the swap disk
        must be less than or equal to the swap size of the flavor. Changing this
        creates a new server.
        """
        return pulumi.get(self, "guest_format")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[str]:
        """
        The UUID of
        the image, volume, or snapshot. Changing this creates a new server.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[int]:
        """
        The size of the volume to create (in gigabytes). Required
        in the following combinations: source=image and destination=volume,
        source=blank and destination=local, and source=blank and destination=volume.
        Changing this creates a new server.
        """
        return pulumi.get(self, "volume_size")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[str]:
        """
        The volume type that will be used, for example SSD
        or HDD storage. The available options depend on how your specific OpenStack
        cloud is configured and what classes of storage are provided. Changing this
        creates a new server.
        """
        return pulumi.get(self, "volume_type")


@pulumi.output_type
class InstanceNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessNetwork":
            suggest = "access_network"
        elif key == "fixedIpV4":
            suggest = "fixed_ip_v4"
        elif key == "fixedIpV6":
            suggest = "fixed_ip_v6"
        elif key == "floatingIp":
            suggest = "floating_ip"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_network: Optional[bool] = None,
                 fixed_ip_v4: Optional[str] = None,
                 fixed_ip_v6: Optional[str] = None,
                 floating_ip: Optional[str] = None,
                 mac: Optional[str] = None,
                 name: Optional[str] = None,
                 port: Optional[str] = None,
                 uuid: Optional[str] = None):
        """
        :param bool access_network: Specifies if this network should be used for
               provisioning access. Accepts true or false. Defaults to false.
        :param str fixed_ip_v4: Specifies a fixed IPv4 address to be used on this
               network. Changing this creates a new server.
        :param str name: The human-readable
               name of the network. Changing this creates a new server.
        :param str port: The port UUID of a
               network to attach to the server. Changing this creates a new server.
        :param str uuid: The network UUID to
               attach to the server. Changing this creates a new server.
        """
        if access_network is not None:
            pulumi.set(__self__, "access_network", access_network)
        if fixed_ip_v4 is not None:
            pulumi.set(__self__, "fixed_ip_v4", fixed_ip_v4)
        if fixed_ip_v6 is not None:
            pulumi.set(__self__, "fixed_ip_v6", fixed_ip_v6)
        if floating_ip is not None:
            pulumi.set(__self__, "floating_ip", floating_ip)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="accessNetwork")
    def access_network(self) -> Optional[bool]:
        """
        Specifies if this network should be used for
        provisioning access. Accepts true or false. Defaults to false.
        """
        return pulumi.get(self, "access_network")

    @property
    @pulumi.getter(name="fixedIpV4")
    def fixed_ip_v4(self) -> Optional[str]:
        """
        Specifies a fixed IPv4 address to be used on this
        network. Changing this creates a new server.
        """
        return pulumi.get(self, "fixed_ip_v4")

    @property
    @pulumi.getter(name="fixedIpV6")
    def fixed_ip_v6(self) -> Optional[str]:
        return pulumi.get(self, "fixed_ip_v6")

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> Optional[str]:
        return pulumi.get(self, "floating_ip")

    @property
    @pulumi.getter
    def mac(self) -> Optional[str]:
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The human-readable
        name of the network. Changing this creates a new server.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> Optional[str]:
        """
        The port UUID of a
        network to attach to the server. Changing this creates a new server.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[str]:
        """
        The network UUID to
        attach to the server. Changing this creates a new server.
        """
        return pulumi.get(self, "uuid")


@pulumi.output_type
class InstancePersonality(dict):
    def __init__(__self__, *,
                 content: str,
                 file: str):
        """
        :param str content: The contents of the file. Limited to 255 bytes.
        :param str file: The absolute path of the destination file.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The contents of the file. Limited to 255 bytes.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def file(self) -> str:
        """
        The absolute path of the destination file.
        """
        return pulumi.get(self, "file")


@pulumi.output_type
class InstanceSchedulerHint(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "additionalProperties":
            suggest = "additional_properties"
        elif key == "buildNearHostIp":
            suggest = "build_near_host_ip"
        elif key == "differentCells":
            suggest = "different_cells"
        elif key == "differentHosts":
            suggest = "different_hosts"
        elif key == "sameHosts":
            suggest = "same_hosts"
        elif key == "targetCell":
            suggest = "target_cell"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceSchedulerHint. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceSchedulerHint.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceSchedulerHint.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 additional_properties: Optional[Mapping[str, Any]] = None,
                 build_near_host_ip: Optional[str] = None,
                 different_cells: Optional[Sequence[str]] = None,
                 different_hosts: Optional[Sequence[str]] = None,
                 group: Optional[str] = None,
                 queries: Optional[Sequence[str]] = None,
                 same_hosts: Optional[Sequence[str]] = None,
                 target_cell: Optional[str] = None):
        """
        :param Mapping[str, Any] additional_properties: Arbitrary key/value pairs of additional
               properties to pass to the scheduler.
        :param str build_near_host_ip: An IP Address in CIDR form. The instance
               will be placed on a compute node that is in the same subnet.
        :param Sequence[str] different_cells: The names of cells where not to build the instance.
        :param Sequence[str] different_hosts: A list of instance UUIDs. The instance will
               be scheduled on a different host than all other instances.
        :param str group: A UUID of a Server Group. The instance will be placed
               into that group.
        :param Sequence[str] queries: A conditional query that a compute node must pass in
               order to host an instance. The query must use the `JsonFilter` syntax
               which is described
               [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
               At this time, only simple queries are supported. Compound queries using
               `and`, `or`, or `not` are not supported. An example of a simple query is:
        :param Sequence[str] same_hosts: A list of instance UUIDs. The instance will be
               scheduled on the same host of those specified.
        :param str target_cell: The name of a cell to host the instance.
        """
        if additional_properties is not None:
            pulumi.set(__self__, "additional_properties", additional_properties)
        if build_near_host_ip is not None:
            pulumi.set(__self__, "build_near_host_ip", build_near_host_ip)
        if different_cells is not None:
            pulumi.set(__self__, "different_cells", different_cells)
        if different_hosts is not None:
            pulumi.set(__self__, "different_hosts", different_hosts)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if queries is not None:
            pulumi.set(__self__, "queries", queries)
        if same_hosts is not None:
            pulumi.set(__self__, "same_hosts", same_hosts)
        if target_cell is not None:
            pulumi.set(__self__, "target_cell", target_cell)

    @property
    @pulumi.getter(name="additionalProperties")
    def additional_properties(self) -> Optional[Mapping[str, Any]]:
        """
        Arbitrary key/value pairs of additional
        properties to pass to the scheduler.
        """
        return pulumi.get(self, "additional_properties")

    @property
    @pulumi.getter(name="buildNearHostIp")
    def build_near_host_ip(self) -> Optional[str]:
        """
        An IP Address in CIDR form. The instance
        will be placed on a compute node that is in the same subnet.
        """
        return pulumi.get(self, "build_near_host_ip")

    @property
    @pulumi.getter(name="differentCells")
    def different_cells(self) -> Optional[Sequence[str]]:
        """
        The names of cells where not to build the instance.
        """
        return pulumi.get(self, "different_cells")

    @property
    @pulumi.getter(name="differentHosts")
    def different_hosts(self) -> Optional[Sequence[str]]:
        """
        A list of instance UUIDs. The instance will
        be scheduled on a different host than all other instances.
        """
        return pulumi.get(self, "different_hosts")

    @property
    @pulumi.getter
    def group(self) -> Optional[str]:
        """
        A UUID of a Server Group. The instance will be placed
        into that group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def queries(self) -> Optional[Sequence[str]]:
        """
        A conditional query that a compute node must pass in
        order to host an instance. The query must use the `JsonFilter` syntax
        which is described
        [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
        At this time, only simple queries are supported. Compound queries using
        `and`, `or`, or `not` are not supported. An example of a simple query is:
        """
        return pulumi.get(self, "queries")

    @property
    @pulumi.getter(name="sameHosts")
    def same_hosts(self) -> Optional[Sequence[str]]:
        """
        A list of instance UUIDs. The instance will be
        scheduled on the same host of those specified.
        """
        return pulumi.get(self, "same_hosts")

    @property
    @pulumi.getter(name="targetCell")
    def target_cell(self) -> Optional[str]:
        """
        The name of a cell to host the instance.
        """
        return pulumi.get(self, "target_cell")


@pulumi.output_type
class InstanceVendorOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "detachPortsBeforeDestroy":
            suggest = "detach_ports_before_destroy"
        elif key == "ignoreResizeConfirmation":
            suggest = "ignore_resize_confirmation"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceVendorOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceVendorOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceVendorOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 detach_ports_before_destroy: Optional[bool] = None,
                 ignore_resize_confirmation: Optional[bool] = None):
        """
        :param bool detach_ports_before_destroy: Whether to try to detach all attached
               ports to the vm before destroying it to make sure the port state is correct
               after the vm destruction. This is helpful when the port is not deleted.
        :param bool ignore_resize_confirmation: Boolean to control whether
               to ignore manual confirmation of the instance resizing. This can be helpful
               to work with some OpenStack clouds which automatically confirm resizing of
               instances after some timeout.
        """
        if detach_ports_before_destroy is not None:
            pulumi.set(__self__, "detach_ports_before_destroy", detach_ports_before_destroy)
        if ignore_resize_confirmation is not None:
            pulumi.set(__self__, "ignore_resize_confirmation", ignore_resize_confirmation)

    @property
    @pulumi.getter(name="detachPortsBeforeDestroy")
    def detach_ports_before_destroy(self) -> Optional[bool]:
        """
        Whether to try to detach all attached
        ports to the vm before destroying it to make sure the port state is correct
        after the vm destruction. This is helpful when the port is not deleted.
        """
        return pulumi.get(self, "detach_ports_before_destroy")

    @property
    @pulumi.getter(name="ignoreResizeConfirmation")
    def ignore_resize_confirmation(self) -> Optional[bool]:
        """
        Boolean to control whether
        to ignore manual confirmation of the instance resizing. This can be helpful
        to work with some OpenStack clouds which automatically confirm resizing of
        instances after some timeout.
        """
        return pulumi.get(self, "ignore_resize_confirmation")


@pulumi.output_type
class InstanceVolume(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "volumeId":
            suggest = "volume_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceVolume. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceVolume.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceVolume.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 volume_id: str,
                 device: Optional[str] = None,
                 id: Optional[str] = None):
        pulumi.set(__self__, "volume_id", volume_id)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> str:
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


@pulumi.output_type
class SecGroupRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fromPort":
            suggest = "from_port"
        elif key == "ipProtocol":
            suggest = "ip_protocol"
        elif key == "toPort":
            suggest = "to_port"
        elif key == "fromGroupId":
            suggest = "from_group_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecGroupRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecGroupRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecGroupRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 from_port: int,
                 ip_protocol: str,
                 to_port: int,
                 cidr: Optional[str] = None,
                 from_group_id: Optional[str] = None,
                 id: Optional[str] = None,
                 self: Optional[bool] = None):
        """
        :param int from_port: An integer representing the lower bound of the port
               range to open. Changing this creates a new security group rule.
        :param str ip_protocol: The protocol type that will be allowed. Changing
               this creates a new security group rule.
        :param int to_port: An integer representing the upper bound of the port
               range to open. Changing this creates a new security group rule.
        :param str cidr: Required if `from_group_id` or `self` is empty. The IP range
               that will be the source of network traffic to the security group. Use 0.0.0.0/0
               to allow all IP addresses. Changing this creates a new security group rule. Cannot
               be combined with `from_group_id` or `self`.
        :param str from_group_id: Required if `cidr` or `self` is empty. The ID of a
               group from which to forward traffic to the parent group. Changing this creates a
               new security group rule. Cannot be combined with `cidr` or `self`.
        :param bool self: Required if `cidr` and `from_group_id` is empty. If true,
               the security group itself will be added as a source to this ingress rule. Cannot
               be combined with `cidr` or `from_group_id`.
        """
        pulumi.set(__self__, "from_port", from_port)
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        pulumi.set(__self__, "to_port", to_port)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if from_group_id is not None:
            pulumi.set(__self__, "from_group_id", from_group_id)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if self is not None:
            pulumi.set(__self__, "self", self)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> int:
        """
        An integer representing the lower bound of the port
        range to open. Changing this creates a new security group rule.
        """
        return pulumi.get(self, "from_port")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> str:
        """
        The protocol type that will be allowed. Changing
        this creates a new security group rule.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> int:
        """
        An integer representing the upper bound of the port
        range to open. Changing this creates a new security group rule.
        """
        return pulumi.get(self, "to_port")

    @property
    @pulumi.getter
    def cidr(self) -> Optional[str]:
        """
        Required if `from_group_id` or `self` is empty. The IP range
        that will be the source of network traffic to the security group. Use 0.0.0.0/0
        to allow all IP addresses. Changing this creates a new security group rule. Cannot
        be combined with `from_group_id` or `self`.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="fromGroupId")
    def from_group_id(self) -> Optional[str]:
        """
        Required if `cidr` or `self` is empty. The ID of a
        group from which to forward traffic to the parent group. Changing this creates a
        new security group rule. Cannot be combined with `cidr` or `self`.
        """
        return pulumi.get(self, "from_group_id")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def self(self) -> Optional[bool]:
        """
        Required if `cidr` and `from_group_id` is empty. If true,
        the security group itself will be added as a source to this ingress rule. Cannot
        be combined with `cidr` or `from_group_id`.
        """
        return pulumi.get(self, "self")


@pulumi.output_type
class VolumeAttachVendorOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ignoreVolumeConfirmation":
            suggest = "ignore_volume_confirmation"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VolumeAttachVendorOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VolumeAttachVendorOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VolumeAttachVendorOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ignore_volume_confirmation: Optional[bool] = None):
        """
        :param bool ignore_volume_confirmation: Boolean to control whether
               to ignore volume status confirmation of the attached volume. This can be helpful
               to work with some OpenStack clouds which don't have the Block Storage V3 API available.
        """
        if ignore_volume_confirmation is not None:
            pulumi.set(__self__, "ignore_volume_confirmation", ignore_volume_confirmation)

    @property
    @pulumi.getter(name="ignoreVolumeConfirmation")
    def ignore_volume_confirmation(self) -> Optional[bool]:
        """
        Boolean to control whether
        to ignore volume status confirmation of the attached volume. This can be helpful
        to work with some OpenStack clouds which don't have the Block Storage V3 API available.
        """
        return pulumi.get(self, "ignore_volume_confirmation")


@pulumi.output_type
class GetInstanceV2NetworkResult(dict):
    def __init__(__self__, *,
                 fixed_ip_v4: str,
                 fixed_ip_v6: str,
                 mac: str,
                 name: str,
                 port: str,
                 uuid: str):
        """
        :param str fixed_ip_v4: The IPv4 address assigned to this network port.
        :param str fixed_ip_v6: The IPv6 address assigned to this network port.
        :param str mac: The MAC address assigned to this network interface.
        :param str name: The name of the network
        :param str port: The port UUID for this network
        :param str uuid: The UUID of the network
        """
        pulumi.set(__self__, "fixed_ip_v4", fixed_ip_v4)
        pulumi.set(__self__, "fixed_ip_v6", fixed_ip_v6)
        pulumi.set(__self__, "mac", mac)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="fixedIpV4")
    def fixed_ip_v4(self) -> str:
        """
        The IPv4 address assigned to this network port.
        """
        return pulumi.get(self, "fixed_ip_v4")

    @property
    @pulumi.getter(name="fixedIpV6")
    def fixed_ip_v6(self) -> str:
        """
        The IPv6 address assigned to this network port.
        """
        return pulumi.get(self, "fixed_ip_v6")

    @property
    @pulumi.getter
    def mac(self) -> str:
        """
        The MAC address assigned to this network interface.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the network
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> str:
        """
        The port UUID for this network
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        The UUID of the network
        """
        return pulumi.get(self, "uuid")


