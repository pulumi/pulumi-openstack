# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = [
    'NetworkSegment',
    'PortAllowedAddressPair',
    'PortBinding',
    'PortExtraDhcpOption',
    'PortFixedIp',
    'RouterExternalFixedIp',
    'RouterVendorOptions',
    'SubnetAllocationPool',
    'TrunkSubPort',
    'GetNetworkSegmentResult',
    'GetPortAllowedAddressPairResult',
    'GetPortBindingResult',
    'GetPortExtraDhcpOptionResult',
    'GetRouterExternalFixedIpResult',
    'GetSubnetAllocationPoolResult',
    'GetSubnetHostRouteResult',
    'GetTrunkSubPortResult',
]

@pulumi.output_type
class NetworkSegment(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkType":
            suggest = "network_type"
        elif key == "physicalNetwork":
            suggest = "physical_network"
        elif key == "segmentationId":
            suggest = "segmentation_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkSegment. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkSegment.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkSegment.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 network_type: Optional[str] = None,
                 physical_network: Optional[str] = None,
                 segmentation_id: Optional[int] = None):
        """
        :param str network_type: The type of physical network.
        :param str physical_network: The physical network where this network is implemented.
        :param int segmentation_id: An isolated segment on the physical network.
        """
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if physical_network is not None:
            pulumi.set(__self__, "physical_network", physical_network)
        if segmentation_id is not None:
            pulumi.set(__self__, "segmentation_id", segmentation_id)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[str]:
        """
        The type of physical network.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="physicalNetwork")
    def physical_network(self) -> Optional[str]:
        """
        The physical network where this network is implemented.
        """
        return pulumi.get(self, "physical_network")

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> Optional[int]:
        """
        An isolated segment on the physical network.
        """
        return pulumi.get(self, "segmentation_id")


@pulumi.output_type
class PortAllowedAddressPair(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipAddress":
            suggest = "ip_address"
        elif key == "macAddress":
            suggest = "mac_address"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PortAllowedAddressPair. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PortAllowedAddressPair.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PortAllowedAddressPair.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_address: str,
                 mac_address: Optional[str] = None):
        """
        :param str ip_address: The additional IP address.
        :param str mac_address: The additional MAC address.
        """
        pulumi.set(__self__, "ip_address", ip_address)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The additional IP address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        """
        The additional MAC address.
        """
        return pulumi.get(self, "mac_address")


@pulumi.output_type
class PortBinding(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hostId":
            suggest = "host_id"
        elif key == "vifDetails":
            suggest = "vif_details"
        elif key == "vifType":
            suggest = "vif_type"
        elif key == "vnicType":
            suggest = "vnic_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PortBinding. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PortBinding.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PortBinding.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 host_id: Optional[str] = None,
                 profile: Optional[str] = None,
                 vif_details: Optional[Mapping[str, str]] = None,
                 vif_type: Optional[str] = None,
                 vnic_type: Optional[str] = None):
        """
        :param str host_id: The ID of the host to allocate port on.
        :param str profile: Custom data to be passed as `binding:profile`. Data
               must be passed as JSON.
        :param Mapping[str, str] vif_details: A map of JSON strings containing additional
               details for this specific binding.
        :param str vif_type: The VNIC type of the port binding.
        :param str vnic_type: VNIC type for the port. Can either be `direct`,
               `direct-physical`, `macvtap`, `normal`, `baremetal` or `virtio-forwarder`.
               Default value is `normal`. It can be updated on unbound ports only.
        """
        if host_id is not None:
            pulumi.set(__self__, "host_id", host_id)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if vif_details is not None:
            pulumi.set(__self__, "vif_details", vif_details)
        if vif_type is not None:
            pulumi.set(__self__, "vif_type", vif_type)
        if vnic_type is not None:
            pulumi.set(__self__, "vnic_type", vnic_type)

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> Optional[str]:
        """
        The ID of the host to allocate port on.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter
    def profile(self) -> Optional[str]:
        """
        Custom data to be passed as `binding:profile`. Data
        must be passed as JSON.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter(name="vifDetails")
    def vif_details(self) -> Optional[Mapping[str, str]]:
        """
        A map of JSON strings containing additional
        details for this specific binding.
        """
        return pulumi.get(self, "vif_details")

    @property
    @pulumi.getter(name="vifType")
    def vif_type(self) -> Optional[str]:
        """
        The VNIC type of the port binding.
        """
        return pulumi.get(self, "vif_type")

    @property
    @pulumi.getter(name="vnicType")
    def vnic_type(self) -> Optional[str]:
        """
        VNIC type for the port. Can either be `direct`,
        `direct-physical`, `macvtap`, `normal`, `baremetal` or `virtio-forwarder`.
        Default value is `normal`. It can be updated on unbound ports only.
        """
        return pulumi.get(self, "vnic_type")


@pulumi.output_type
class PortExtraDhcpOption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipVersion":
            suggest = "ip_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PortExtraDhcpOption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PortExtraDhcpOption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PortExtraDhcpOption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 value: str,
                 ip_version: Optional[int] = None):
        """
        :param str name: Name of the DHCP option.
        :param str value: Value of the DHCP option.
        :param int ip_version: IP protocol version. Defaults to 4.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the DHCP option.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Value of the DHCP option.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[int]:
        """
        IP protocol version. Defaults to 4.
        """
        return pulumi.get(self, "ip_version")


@pulumi.output_type
class PortFixedIp(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipAddress":
            suggest = "ip_address"
        elif key == "subnetId":
            suggest = "subnet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PortFixedIp. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PortFixedIp.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PortFixedIp.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_address: Optional[str] = None,
                 subnet_id: Optional[str] = None):
        """
        :param str ip_address: IP address desired in the subnet for this port. If
               you don't specify `ip_address`, an available IP address from the specified
               subnet will be allocated to this port. This field will not be populated if it
               is left blank or omitted. To retrieve the assigned IP address, use the
               `all_fixed_ips` attribute.
        :param str subnet_id: Subnet in which to allocate IP address for
               this port.
        """
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        IP address desired in the subnet for this port. If
        you don't specify `ip_address`, an available IP address from the specified
        subnet will be allocated to this port. This field will not be populated if it
        is left blank or omitted. To retrieve the assigned IP address, use the
        `all_fixed_ips` attribute.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        Subnet in which to allocate IP address for
        this port.
        """
        return pulumi.get(self, "subnet_id")


@pulumi.output_type
class RouterExternalFixedIp(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipAddress":
            suggest = "ip_address"
        elif key == "subnetId":
            suggest = "subnet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RouterExternalFixedIp. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RouterExternalFixedIp.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RouterExternalFixedIp.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_address: Optional[str] = None,
                 subnet_id: Optional[str] = None):
        """
        :param str ip_address: The IP address to set on the router.
        :param str subnet_id: Subnet in which the fixed IP belongs to.
        """
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The IP address to set on the router.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        Subnet in which the fixed IP belongs to.
        """
        return pulumi.get(self, "subnet_id")


@pulumi.output_type
class RouterVendorOptions(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "setRouterGatewayAfterCreate":
            suggest = "set_router_gateway_after_create"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in RouterVendorOptions. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        RouterVendorOptions.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        RouterVendorOptions.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 set_router_gateway_after_create: Optional[bool] = None):
        """
        :param bool set_router_gateway_after_create: Boolean to control whether
               the Router gateway is assigned during creation or updated after creation.
        """
        if set_router_gateway_after_create is not None:
            pulumi.set(__self__, "set_router_gateway_after_create", set_router_gateway_after_create)

    @property
    @pulumi.getter(name="setRouterGatewayAfterCreate")
    def set_router_gateway_after_create(self) -> Optional[bool]:
        """
        Boolean to control whether
        the Router gateway is assigned during creation or updated after creation.
        """
        return pulumi.get(self, "set_router_gateway_after_create")


@pulumi.output_type
class SubnetAllocationPool(dict):
    def __init__(__self__, *,
                 end: str,
                 start: str):
        """
        :param str end: The ending address.
        :param str start: The starting address.
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> str:
        """
        The ending address.
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> str:
        """
        The starting address.
        """
        return pulumi.get(self, "start")


@pulumi.output_type
class TrunkSubPort(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "portId":
            suggest = "port_id"
        elif key == "segmentationId":
            suggest = "segmentation_id"
        elif key == "segmentationType":
            suggest = "segmentation_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TrunkSubPort. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TrunkSubPort.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TrunkSubPort.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 port_id: str,
                 segmentation_id: int,
                 segmentation_type: str):
        """
        :param str port_id: The ID of the port to be made a subport of the trunk.
        :param int segmentation_id: The numeric id of the subport segment.
        :param str segmentation_type: The segmentation technology to use, e.g., "vlan".
        """
        pulumi.set(__self__, "port_id", port_id)
        pulumi.set(__self__, "segmentation_id", segmentation_id)
        pulumi.set(__self__, "segmentation_type", segmentation_type)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> str:
        """
        The ID of the port to be made a subport of the trunk.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> int:
        """
        The numeric id of the subport segment.
        """
        return pulumi.get(self, "segmentation_id")

    @property
    @pulumi.getter(name="segmentationType")
    def segmentation_type(self) -> str:
        """
        The segmentation technology to use, e.g., "vlan".
        """
        return pulumi.get(self, "segmentation_type")


@pulumi.output_type
class GetNetworkSegmentResult(dict):
    def __init__(__self__, *,
                 network_type: str,
                 physical_network: str,
                 segmentation_id: int):
        pulumi.set(__self__, "network_type", network_type)
        pulumi.set(__self__, "physical_network", physical_network)
        pulumi.set(__self__, "segmentation_id", segmentation_id)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="physicalNetwork")
    def physical_network(self) -> str:
        return pulumi.get(self, "physical_network")

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> int:
        return pulumi.get(self, "segmentation_id")


@pulumi.output_type
class GetPortAllowedAddressPairResult(dict):
    def __init__(__self__, *,
                 ip_address: str,
                 mac_address: str):
        """
        :param str ip_address: The additional IP address.
        :param str mac_address: The MAC address of the port.
        """
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "mac_address", mac_address)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        """
        The additional IP address.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        """
        The MAC address of the port.
        """
        return pulumi.get(self, "mac_address")


@pulumi.output_type
class GetPortBindingResult(dict):
    def __init__(__self__, *,
                 host_id: str,
                 profile: str,
                 vif_details: Mapping[str, str],
                 vif_type: str,
                 vnic_type: str):
        """
        :param str host_id: The ID of the host, which has the allocatee port.
        :param str profile: A JSON string containing the binding profile information.
        :param Mapping[str, str] vif_details: A map of JSON strings containing additional details for this
               specific binding.
        :param str vif_type: The VNIC type of the port binding.
        :param str vnic_type: VNIC type for the port.
        """
        pulumi.set(__self__, "host_id", host_id)
        pulumi.set(__self__, "profile", profile)
        pulumi.set(__self__, "vif_details", vif_details)
        pulumi.set(__self__, "vif_type", vif_type)
        pulumi.set(__self__, "vnic_type", vnic_type)

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> str:
        """
        The ID of the host, which has the allocatee port.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter
    def profile(self) -> str:
        """
        A JSON string containing the binding profile information.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter(name="vifDetails")
    def vif_details(self) -> Mapping[str, str]:
        """
        A map of JSON strings containing additional details for this
        specific binding.
        """
        return pulumi.get(self, "vif_details")

    @property
    @pulumi.getter(name="vifType")
    def vif_type(self) -> str:
        """
        The VNIC type of the port binding.
        """
        return pulumi.get(self, "vif_type")

    @property
    @pulumi.getter(name="vnicType")
    def vnic_type(self) -> str:
        """
        VNIC type for the port.
        """
        return pulumi.get(self, "vnic_type")


@pulumi.output_type
class GetPortExtraDhcpOptionResult(dict):
    def __init__(__self__, *,
                 ip_version: int,
                 name: str,
                 value: str):
        """
        :param int ip_version: IP protocol version
        :param str name: The name of the port.
        :param str value: Value of the DHCP option.
        """
        pulumi.set(__self__, "ip_version", ip_version)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> int:
        """
        IP protocol version
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the port.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Value of the DHCP option.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetRouterExternalFixedIpResult(dict):
    def __init__(__self__, *,
                 ip_address: Optional[str] = None,
                 subnet_id: Optional[str] = None):
        """
        :param str ip_address: The IP address to set on the router.
        :param str subnet_id: Subnet in which the fixed IP belongs to.
        """
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The IP address to set on the router.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        Subnet in which the fixed IP belongs to.
        """
        return pulumi.get(self, "subnet_id")


@pulumi.output_type
class GetSubnetAllocationPoolResult(dict):
    def __init__(__self__, *,
                 end: str,
                 start: str):
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> str:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def start(self) -> str:
        return pulumi.get(self, "start")


@pulumi.output_type
class GetSubnetHostRouteResult(dict):
    def __init__(__self__, *,
                 destination_cidr: str,
                 next_hop: str):
        pulumi.set(__self__, "destination_cidr", destination_cidr)
        pulumi.set(__self__, "next_hop", next_hop)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> str:
        return pulumi.get(self, "destination_cidr")

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> str:
        return pulumi.get(self, "next_hop")


@pulumi.output_type
class GetTrunkSubPortResult(dict):
    def __init__(__self__, *,
                 port_id: str,
                 segmentation_id: int,
                 segmentation_type: str):
        """
        :param str port_id: The ID of the trunk parent port.
        :param int segmentation_id: The numeric id of the subport segment.
        :param str segmentation_type: The segmenation tecnology used, e.g., "vlan".
        """
        pulumi.set(__self__, "port_id", port_id)
        pulumi.set(__self__, "segmentation_id", segmentation_id)
        pulumi.set(__self__, "segmentation_type", segmentation_type)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> str:
        """
        The ID of the trunk parent port.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> int:
        """
        The numeric id of the subport segment.
        """
        return pulumi.get(self, "segmentation_id")

    @property
    @pulumi.getter(name="segmentationType")
    def segmentation_type(self) -> str:
        """
        The segmenation tecnology used, e.g., "vlan".
        """
        return pulumi.get(self, "segmentation_type")


