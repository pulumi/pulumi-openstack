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
    'NetworkSegmentArgs',
    'PortAllowedAddressPairArgs',
    'PortBindingArgs',
    'PortExtraDhcpOptionArgs',
    'PortFixedIpArgs',
    'RouterExternalFixedIpArgs',
    'RouterVendorOptionsArgs',
    'SubnetAllocationPoolArgs',
    'SubnetAllocationPoolsCollectionArgs',
    'SubnetHostRouteArgs',
    'TrunkSubPortArgs',
]

@pulumi.input_type
class NetworkSegmentArgs:
    def __init__(__self__, *,
                 network_type: Optional[pulumi.Input[str]] = None,
                 physical_network: Optional[pulumi.Input[str]] = None,
                 segmentation_id: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] network_type: The type of physical network.
        :param pulumi.Input[str] physical_network: The physical network where this network is implemented.
        :param pulumi.Input[int] segmentation_id: An isolated segment on the physical network.
        """
        NetworkSegmentArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            network_type=network_type,
            physical_network=physical_network,
            segmentation_id=segmentation_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             network_type: Optional[pulumi.Input[str]] = None,
             physical_network: Optional[pulumi.Input[str]] = None,
             segmentation_id: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if network_type is None and 'networkType' in kwargs:
            network_type = kwargs['networkType']
        if physical_network is None and 'physicalNetwork' in kwargs:
            physical_network = kwargs['physicalNetwork']
        if segmentation_id is None and 'segmentationId' in kwargs:
            segmentation_id = kwargs['segmentationId']

        if network_type is not None:
            _setter("network_type", network_type)
        if physical_network is not None:
            _setter("physical_network", physical_network)
        if segmentation_id is not None:
            _setter("segmentation_id", segmentation_id)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of physical network.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="physicalNetwork")
    def physical_network(self) -> Optional[pulumi.Input[str]]:
        """
        The physical network where this network is implemented.
        """
        return pulumi.get(self, "physical_network")

    @physical_network.setter
    def physical_network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "physical_network", value)

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> Optional[pulumi.Input[int]]:
        """
        An isolated segment on the physical network.
        """
        return pulumi.get(self, "segmentation_id")

    @segmentation_id.setter
    def segmentation_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "segmentation_id", value)


@pulumi.input_type
class PortAllowedAddressPairArgs:
    def __init__(__self__, *,
                 ip_address: pulumi.Input[str],
                 mac_address: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] ip_address: The additional IP address.
        :param pulumi.Input[str] mac_address: The additional MAC address.
        """
        PortAllowedAddressPairArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ip_address=ip_address,
            mac_address=mac_address,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ip_address: Optional[pulumi.Input[str]] = None,
             mac_address: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if ip_address is None and 'ipAddress' in kwargs:
            ip_address = kwargs['ipAddress']
        if ip_address is None:
            raise TypeError("Missing 'ip_address' argument")
        if mac_address is None and 'macAddress' in kwargs:
            mac_address = kwargs['macAddress']

        _setter("ip_address", ip_address)
        if mac_address is not None:
            _setter("mac_address", mac_address)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Input[str]:
        """
        The additional IP address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[pulumi.Input[str]]:
        """
        The additional MAC address.
        """
        return pulumi.get(self, "mac_address")

    @mac_address.setter
    def mac_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_address", value)


@pulumi.input_type
class PortBindingArgs:
    def __init__(__self__, *,
                 host_id: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 vif_details: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vif_type: Optional[pulumi.Input[str]] = None,
                 vnic_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] host_id: The ID of the host to allocate port on.
        :param pulumi.Input[str] profile: Custom data to be passed as `binding:profile`. Data
               must be passed as JSON.
        :param pulumi.Input[Mapping[str, Any]] vif_details: A map of JSON strings containing additional
               details for this specific binding.
        :param pulumi.Input[str] vif_type: The VNIC type of the port binding.
        :param pulumi.Input[str] vnic_type: VNIC type for the port. Can either be `direct`,
               `direct-physical`, `macvtap`, `normal`, `baremetal` or `virtio-forwarder`.
               Default value is `normal`.
        """
        PortBindingArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            host_id=host_id,
            profile=profile,
            vif_details=vif_details,
            vif_type=vif_type,
            vnic_type=vnic_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             host_id: Optional[pulumi.Input[str]] = None,
             profile: Optional[pulumi.Input[str]] = None,
             vif_details: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             vif_type: Optional[pulumi.Input[str]] = None,
             vnic_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if host_id is None and 'hostId' in kwargs:
            host_id = kwargs['hostId']
        if vif_details is None and 'vifDetails' in kwargs:
            vif_details = kwargs['vifDetails']
        if vif_type is None and 'vifType' in kwargs:
            vif_type = kwargs['vifType']
        if vnic_type is None and 'vnicType' in kwargs:
            vnic_type = kwargs['vnicType']

        if host_id is not None:
            _setter("host_id", host_id)
        if profile is not None:
            _setter("profile", profile)
        if vif_details is not None:
            _setter("vif_details", vif_details)
        if vif_type is not None:
            _setter("vif_type", vif_type)
        if vnic_type is not None:
            _setter("vnic_type", vnic_type)

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the host to allocate port on.
        """
        return pulumi.get(self, "host_id")

    @host_id.setter
    def host_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_id", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Custom data to be passed as `binding:profile`. Data
        must be passed as JSON.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="vifDetails")
    def vif_details(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map of JSON strings containing additional
        details for this specific binding.
        """
        return pulumi.get(self, "vif_details")

    @vif_details.setter
    def vif_details(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "vif_details", value)

    @property
    @pulumi.getter(name="vifType")
    def vif_type(self) -> Optional[pulumi.Input[str]]:
        """
        The VNIC type of the port binding.
        """
        return pulumi.get(self, "vif_type")

    @vif_type.setter
    def vif_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vif_type", value)

    @property
    @pulumi.getter(name="vnicType")
    def vnic_type(self) -> Optional[pulumi.Input[str]]:
        """
        VNIC type for the port. Can either be `direct`,
        `direct-physical`, `macvtap`, `normal`, `baremetal` or `virtio-forwarder`.
        Default value is `normal`.
        """
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
        """
        :param pulumi.Input[str] name: Name of the DHCP option.
        :param pulumi.Input[str] value: Value of the DHCP option.
        :param pulumi.Input[int] ip_version: IP protocol version. Defaults to 4.
        """
        PortExtraDhcpOptionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            value=value,
            ip_version=ip_version,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             ip_version: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if name is None:
            raise TypeError("Missing 'name' argument")
        if value is None:
            raise TypeError("Missing 'value' argument")
        if ip_version is None and 'ipVersion' in kwargs:
            ip_version = kwargs['ipVersion']

        _setter("name", name)
        _setter("value", value)
        if ip_version is not None:
            _setter("ip_version", ip_version)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the DHCP option.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Value of the DHCP option.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[int]]:
        """
        IP protocol version. Defaults to 4.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_version", value)


@pulumi.input_type
class PortFixedIpArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 ip_address: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] subnet_id: Subnet in which to allocate IP address for
               this port.
        :param pulumi.Input[str] ip_address: IP address desired in the subnet for this port. If
               you don't specify `ip_address`, an available IP address from the specified
               subnet will be allocated to this port. This field will not be populated if it
               is left blank or omitted. To retrieve the assigned IP address, use the
               `all_fixed_ips` attribute.
        """
        PortFixedIpArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            subnet_id=subnet_id,
            ip_address=ip_address,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             subnet_id: Optional[pulumi.Input[str]] = None,
             ip_address: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if subnet_id is None and 'subnetId' in kwargs:
            subnet_id = kwargs['subnetId']
        if subnet_id is None:
            raise TypeError("Missing 'subnet_id' argument")
        if ip_address is None and 'ipAddress' in kwargs:
            ip_address = kwargs['ipAddress']

        _setter("subnet_id", subnet_id)
        if ip_address is not None:
            _setter("ip_address", ip_address)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Subnet in which to allocate IP address for
        this port.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address desired in the subnet for this port. If
        you don't specify `ip_address`, an available IP address from the specified
        subnet will be allocated to this port. This field will not be populated if it
        is left blank or omitted. To retrieve the assigned IP address, use the
        `all_fixed_ips` attribute.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)


@pulumi.input_type
class RouterExternalFixedIpArgs:
    def __init__(__self__, *,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] ip_address: The IP address to set on the router.
        :param pulumi.Input[str] subnet_id: Subnet in which the fixed IP belongs to.
        """
        RouterExternalFixedIpArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ip_address=ip_address,
            subnet_id=subnet_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ip_address: Optional[pulumi.Input[str]] = None,
             subnet_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if ip_address is None and 'ipAddress' in kwargs:
            ip_address = kwargs['ipAddress']
        if subnet_id is None and 'subnetId' in kwargs:
            subnet_id = kwargs['subnetId']

        if ip_address is not None:
            _setter("ip_address", ip_address)
        if subnet_id is not None:
            _setter("subnet_id", subnet_id)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address to set on the router.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet in which the fixed IP belongs to.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class RouterVendorOptionsArgs:
    def __init__(__self__, *,
                 set_router_gateway_after_create: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] set_router_gateway_after_create: Boolean to control whether
               the Router gateway is assigned during creation or updated after creation.
        """
        RouterVendorOptionsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            set_router_gateway_after_create=set_router_gateway_after_create,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             set_router_gateway_after_create: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if set_router_gateway_after_create is None and 'setRouterGatewayAfterCreate' in kwargs:
            set_router_gateway_after_create = kwargs['setRouterGatewayAfterCreate']

        if set_router_gateway_after_create is not None:
            _setter("set_router_gateway_after_create", set_router_gateway_after_create)

    @property
    @pulumi.getter(name="setRouterGatewayAfterCreate")
    def set_router_gateway_after_create(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean to control whether
        the Router gateway is assigned during creation or updated after creation.
        """
        return pulumi.get(self, "set_router_gateway_after_create")

    @set_router_gateway_after_create.setter
    def set_router_gateway_after_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_router_gateway_after_create", value)


@pulumi.input_type
class SubnetAllocationPoolArgs:
    def __init__(__self__, *,
                 end: pulumi.Input[str],
                 start: pulumi.Input[str]):
        """
        :param pulumi.Input[str] end: The ending address.
        :param pulumi.Input[str] start: The starting address.
        """
        SubnetAllocationPoolArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            end=end,
            start=start,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             end: Optional[pulumi.Input[str]] = None,
             start: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if end is None:
            raise TypeError("Missing 'end' argument")
        if start is None:
            raise TypeError("Missing 'start' argument")

        _setter("end", end)
        _setter("start", start)

    @property
    @pulumi.getter
    def end(self) -> pulumi.Input[str]:
        """
        The ending address.
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: pulumi.Input[str]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[str]:
        """
        The starting address.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[str]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class SubnetAllocationPoolsCollectionArgs:
    def __init__(__self__, *,
                 end: pulumi.Input[str],
                 start: pulumi.Input[str]):
        """
        :param pulumi.Input[str] end: The ending address.
        :param pulumi.Input[str] start: The starting address.
        """
        SubnetAllocationPoolsCollectionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            end=end,
            start=start,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             end: Optional[pulumi.Input[str]] = None,
             start: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if end is None:
            raise TypeError("Missing 'end' argument")
        if start is None:
            raise TypeError("Missing 'start' argument")

        _setter("end", end)
        _setter("start", start)

    @property
    @pulumi.getter
    def end(self) -> pulumi.Input[str]:
        """
        The ending address.
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: pulumi.Input[str]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[str]:
        """
        The starting address.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[str]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class SubnetHostRouteArgs:
    def __init__(__self__, *,
                 destination_cidr: pulumi.Input[str],
                 next_hop: pulumi.Input[str]):
        """
        :param pulumi.Input[str] destination_cidr: The destination CIDR.
        :param pulumi.Input[str] next_hop: The next hop in the route.
        """
        SubnetHostRouteArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            destination_cidr=destination_cidr,
            next_hop=next_hop,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             destination_cidr: Optional[pulumi.Input[str]] = None,
             next_hop: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if destination_cidr is None and 'destinationCidr' in kwargs:
            destination_cidr = kwargs['destinationCidr']
        if destination_cidr is None:
            raise TypeError("Missing 'destination_cidr' argument")
        if next_hop is None and 'nextHop' in kwargs:
            next_hop = kwargs['nextHop']
        if next_hop is None:
            raise TypeError("Missing 'next_hop' argument")

        _setter("destination_cidr", destination_cidr)
        _setter("next_hop", next_hop)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> pulumi.Input[str]:
        """
        The destination CIDR.
        """
        return pulumi.get(self, "destination_cidr")

    @destination_cidr.setter
    def destination_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr", value)

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> pulumi.Input[str]:
        """
        The next hop in the route.
        """
        return pulumi.get(self, "next_hop")

    @next_hop.setter
    def next_hop(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_hop", value)


@pulumi.input_type
class TrunkSubPortArgs:
    def __init__(__self__, *,
                 port_id: pulumi.Input[str],
                 segmentation_id: pulumi.Input[int],
                 segmentation_type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] port_id: The ID of the port to be made a subport of the trunk.
        :param pulumi.Input[int] segmentation_id: The numeric id of the subport segment.
        :param pulumi.Input[str] segmentation_type: The segmentation technology to use, e.g., "vlan".
        """
        TrunkSubPortArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            port_id=port_id,
            segmentation_id=segmentation_id,
            segmentation_type=segmentation_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             port_id: Optional[pulumi.Input[str]] = None,
             segmentation_id: Optional[pulumi.Input[int]] = None,
             segmentation_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if port_id is None and 'portId' in kwargs:
            port_id = kwargs['portId']
        if port_id is None:
            raise TypeError("Missing 'port_id' argument")
        if segmentation_id is None and 'segmentationId' in kwargs:
            segmentation_id = kwargs['segmentationId']
        if segmentation_id is None:
            raise TypeError("Missing 'segmentation_id' argument")
        if segmentation_type is None and 'segmentationType' in kwargs:
            segmentation_type = kwargs['segmentationType']
        if segmentation_type is None:
            raise TypeError("Missing 'segmentation_type' argument")

        _setter("port_id", port_id)
        _setter("segmentation_id", segmentation_id)
        _setter("segmentation_type", segmentation_type)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Input[str]:
        """
        The ID of the port to be made a subport of the trunk.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> pulumi.Input[int]:
        """
        The numeric id of the subport segment.
        """
        return pulumi.get(self, "segmentation_id")

    @segmentation_id.setter
    def segmentation_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "segmentation_id", value)

    @property
    @pulumi.getter(name="segmentationType")
    def segmentation_type(self) -> pulumi.Input[str]:
        """
        The segmentation technology to use, e.g., "vlan".
        """
        return pulumi.get(self, "segmentation_type")

    @segmentation_type.setter
    def segmentation_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "segmentation_type", value)


