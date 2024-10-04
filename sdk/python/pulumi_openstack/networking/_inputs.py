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
    'NetworkSegmentArgs',
    'PortAllowedAddressPairArgs',
    'PortBindingArgs',
    'PortExtraDhcpOptionArgs',
    'PortFixedIpArgs',
    'RouterExternalFixedIpArgs',
    'RouterVendorOptionsArgs',
    'SubnetAllocationPoolArgs',
    'TrunkSubPortArgs',
]

@pulumi.input_type
class NetworkSegmentArgs:
    def __init__(__self__, *,
                 network_type: Optional[pulumi.Input[str]] = None,
                 physical_network: Optional[pulumi.Input[str]] = None,
                 segmentation_id: Optional[pulumi.Input[int]] = None):
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if physical_network is not None:
            pulumi.set(__self__, "physical_network", physical_network)
        if segmentation_id is not None:
            pulumi.set(__self__, "segmentation_id", segmentation_id)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="physicalNetwork")
    def physical_network(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "physical_network")

    @physical_network.setter
    def physical_network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "physical_network", value)

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "segmentation_id")

    @segmentation_id.setter
    def segmentation_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "segmentation_id", value)


@pulumi.input_type
class PortAllowedAddressPairArgs:
    def __init__(__self__, *,
                 ip_address: pulumi.Input[str],
                 mac_address: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "ip_address", ip_address)
        if mac_address is not None:
            pulumi.set(__self__, "mac_address", mac_address)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_address", value)


@pulumi.input_type
class PortBindingArgs:
    def __init__(__self__, *,
                 host_id: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 vif_details: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vif_type: Optional[pulumi.Input[str]] = None,
                 vnic_type: Optional[pulumi.Input[str]] = None):
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
    def host_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "host_id")

    @host_id.setter
    def host_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_id", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="vifDetails")
    def vif_details(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "vif_details")

    @vif_details.setter
    def vif_details(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "vif_details", value)

    @property
    @pulumi.getter(name="vifType")
    def vif_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vif_type")

    @vif_type.setter
    def vif_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vif_type", value)

    @property
    @pulumi.getter(name="vnicType")
    def vnic_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vnic_type")

    @vnic_type.setter
    def vnic_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vnic_type", value)


@pulumi.input_type
class PortExtraDhcpOptionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str],
                 ip_version: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_version", value)


@pulumi.input_type
class PortFixedIpArgs:
    def __init__(__self__, *,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class RouterExternalFixedIpArgs:
    def __init__(__self__, *,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class RouterVendorOptionsArgs:
    def __init__(__self__, *,
                 set_router_gateway_after_create: Optional[pulumi.Input[bool]] = None):
        if set_router_gateway_after_create is not None:
            pulumi.set(__self__, "set_router_gateway_after_create", set_router_gateway_after_create)

    @property
    @pulumi.getter(name="setRouterGatewayAfterCreate")
    def set_router_gateway_after_create(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "set_router_gateway_after_create")

    @set_router_gateway_after_create.setter
    def set_router_gateway_after_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_router_gateway_after_create", value)


@pulumi.input_type
class SubnetAllocationPoolArgs:
    def __init__(__self__, *,
                 end: pulumi.Input[str],
                 start: pulumi.Input[str]):
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> pulumi.Input[str]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: pulumi.Input[str]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[str]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[str]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class TrunkSubPortArgs:
    def __init__(__self__, *,
                 port_id: pulumi.Input[str],
                 segmentation_id: pulumi.Input[int],
                 segmentation_type: pulumi.Input[str]):
        pulumi.set(__self__, "port_id", port_id)
        pulumi.set(__self__, "segmentation_id", segmentation_id)
        pulumi.set(__self__, "segmentation_type", segmentation_type)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> pulumi.Input[int]:
        return pulumi.get(self, "segmentation_id")

    @segmentation_id.setter
    def segmentation_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "segmentation_id", value)

    @property
    @pulumi.getter(name="segmentationType")
    def segmentation_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "segmentation_type")

    @segmentation_type.setter
    def segmentation_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "segmentation_type", value)


