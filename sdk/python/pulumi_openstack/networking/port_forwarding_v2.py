# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PortForwardingV2Args', 'PortForwardingV2']

@pulumi.input_type
class PortForwardingV2Args:
    def __init__(__self__, *,
                 external_port: pulumi.Input[int],
                 floatingip_id: pulumi.Input[str],
                 internal_ip_address: pulumi.Input[str],
                 internal_port: pulumi.Input[int],
                 internal_port_id: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PortForwardingV2 resource.
        :param pulumi.Input[int] external_port: The TCP/UDP/other protocol port number of the port forwarding. Changing this
               updates the `external_port` of an existing port forwarding.
        :param pulumi.Input[str] floatingip_id: The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        :param pulumi.Input[str] internal_ip_address: The fixed IPv4 address of the Neutron port associated with the port forwarding.
               Changing this updates the `internal_ip_address` of an existing port forwarding.
        :param pulumi.Input[int] internal_port: The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
               port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        :param pulumi.Input[str] internal_port_id: The ID of the Neutron port associated with the port forwarding. Changing
               this updates the `internal_port_id` of an existing port forwarding.
        :param pulumi.Input[str] protocol: The IP protocol used in the port forwarding. Changing this updates the `protocol`
               of an existing port forwarding.
        :param pulumi.Input[str] description: A text describing the port forwarding. Changing this
               updates the `description` of an existing port forwarding.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port forwarding. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               port forwarding.
        """
        PortForwardingV2Args._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            external_port=external_port,
            floatingip_id=floatingip_id,
            internal_ip_address=internal_ip_address,
            internal_port=internal_port,
            internal_port_id=internal_port_id,
            protocol=protocol,
            description=description,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             external_port: Optional[pulumi.Input[int]] = None,
             floatingip_id: Optional[pulumi.Input[str]] = None,
             internal_ip_address: Optional[pulumi.Input[str]] = None,
             internal_port: Optional[pulumi.Input[int]] = None,
             internal_port_id: Optional[pulumi.Input[str]] = None,
             protocol: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if external_port is None and 'externalPort' in kwargs:
            external_port = kwargs['externalPort']
        if external_port is None:
            raise TypeError("Missing 'external_port' argument")
        if floatingip_id is None and 'floatingipId' in kwargs:
            floatingip_id = kwargs['floatingipId']
        if floatingip_id is None:
            raise TypeError("Missing 'floatingip_id' argument")
        if internal_ip_address is None and 'internalIpAddress' in kwargs:
            internal_ip_address = kwargs['internalIpAddress']
        if internal_ip_address is None:
            raise TypeError("Missing 'internal_ip_address' argument")
        if internal_port is None and 'internalPort' in kwargs:
            internal_port = kwargs['internalPort']
        if internal_port is None:
            raise TypeError("Missing 'internal_port' argument")
        if internal_port_id is None and 'internalPortId' in kwargs:
            internal_port_id = kwargs['internalPortId']
        if internal_port_id is None:
            raise TypeError("Missing 'internal_port_id' argument")
        if protocol is None:
            raise TypeError("Missing 'protocol' argument")

        _setter("external_port", external_port)
        _setter("floatingip_id", floatingip_id)
        _setter("internal_ip_address", internal_ip_address)
        _setter("internal_port", internal_port)
        _setter("internal_port_id", internal_port_id)
        _setter("protocol", protocol)
        if description is not None:
            _setter("description", description)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> pulumi.Input[int]:
        """
        The TCP/UDP/other protocol port number of the port forwarding. Changing this
        updates the `external_port` of an existing port forwarding.
        """
        return pulumi.get(self, "external_port")

    @external_port.setter
    def external_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "external_port", value)

    @property
    @pulumi.getter(name="floatingipId")
    def floatingip_id(self) -> pulumi.Input[str]:
        """
        The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        """
        return pulumi.get(self, "floatingip_id")

    @floatingip_id.setter
    def floatingip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "floatingip_id", value)

    @property
    @pulumi.getter(name="internalIpAddress")
    def internal_ip_address(self) -> pulumi.Input[str]:
        """
        The fixed IPv4 address of the Neutron port associated with the port forwarding.
        Changing this updates the `internal_ip_address` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_ip_address")

    @internal_ip_address.setter
    def internal_ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_ip_address", value)

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> pulumi.Input[int]:
        """
        The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
        port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_port")

    @internal_port.setter
    def internal_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "internal_port", value)

    @property
    @pulumi.getter(name="internalPortId")
    def internal_port_id(self) -> pulumi.Input[str]:
        """
        The ID of the Neutron port associated with the port forwarding. Changing
        this updates the `internal_port_id` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_port_id")

    @internal_port_id.setter
    def internal_port_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "internal_port_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The IP protocol used in the port forwarding. Changing this updates the `protocol`
        of an existing port forwarding.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A text describing the port forwarding. Changing this
        updates the `description` of an existing port forwarding.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a port forwarding. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        port forwarding.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _PortForwardingV2State:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[int]] = None,
                 floatingip_id: Optional[pulumi.Input[str]] = None,
                 internal_ip_address: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[int]] = None,
                 internal_port_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PortForwardingV2 resources.
        :param pulumi.Input[str] description: A text describing the port forwarding. Changing this
               updates the `description` of an existing port forwarding.
        :param pulumi.Input[int] external_port: The TCP/UDP/other protocol port number of the port forwarding. Changing this
               updates the `external_port` of an existing port forwarding.
        :param pulumi.Input[str] floatingip_id: The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        :param pulumi.Input[str] internal_ip_address: The fixed IPv4 address of the Neutron port associated with the port forwarding.
               Changing this updates the `internal_ip_address` of an existing port forwarding.
        :param pulumi.Input[int] internal_port: The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
               port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        :param pulumi.Input[str] internal_port_id: The ID of the Neutron port associated with the port forwarding. Changing
               this updates the `internal_port_id` of an existing port forwarding.
        :param pulumi.Input[str] protocol: The IP protocol used in the port forwarding. Changing this updates the `protocol`
               of an existing port forwarding.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port forwarding. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               port forwarding.
        """
        _PortForwardingV2State._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            external_port=external_port,
            floatingip_id=floatingip_id,
            internal_ip_address=internal_ip_address,
            internal_port=internal_port,
            internal_port_id=internal_port_id,
            protocol=protocol,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             external_port: Optional[pulumi.Input[int]] = None,
             floatingip_id: Optional[pulumi.Input[str]] = None,
             internal_ip_address: Optional[pulumi.Input[str]] = None,
             internal_port: Optional[pulumi.Input[int]] = None,
             internal_port_id: Optional[pulumi.Input[str]] = None,
             protocol: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if external_port is None and 'externalPort' in kwargs:
            external_port = kwargs['externalPort']
        if floatingip_id is None and 'floatingipId' in kwargs:
            floatingip_id = kwargs['floatingipId']
        if internal_ip_address is None and 'internalIpAddress' in kwargs:
            internal_ip_address = kwargs['internalIpAddress']
        if internal_port is None and 'internalPort' in kwargs:
            internal_port = kwargs['internalPort']
        if internal_port_id is None and 'internalPortId' in kwargs:
            internal_port_id = kwargs['internalPortId']

        if description is not None:
            _setter("description", description)
        if external_port is not None:
            _setter("external_port", external_port)
        if floatingip_id is not None:
            _setter("floatingip_id", floatingip_id)
        if internal_ip_address is not None:
            _setter("internal_ip_address", internal_ip_address)
        if internal_port is not None:
            _setter("internal_port", internal_port)
        if internal_port_id is not None:
            _setter("internal_port_id", internal_port_id)
        if protocol is not None:
            _setter("protocol", protocol)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A text describing the port forwarding. Changing this
        updates the `description` of an existing port forwarding.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> Optional[pulumi.Input[int]]:
        """
        The TCP/UDP/other protocol port number of the port forwarding. Changing this
        updates the `external_port` of an existing port forwarding.
        """
        return pulumi.get(self, "external_port")

    @external_port.setter
    def external_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "external_port", value)

    @property
    @pulumi.getter(name="floatingipId")
    def floatingip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        """
        return pulumi.get(self, "floatingip_id")

    @floatingip_id.setter
    def floatingip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floatingip_id", value)

    @property
    @pulumi.getter(name="internalIpAddress")
    def internal_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The fixed IPv4 address of the Neutron port associated with the port forwarding.
        Changing this updates the `internal_ip_address` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_ip_address")

    @internal_ip_address.setter
    def internal_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_ip_address", value)

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> Optional[pulumi.Input[int]]:
        """
        The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
        port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_port")

    @internal_port.setter
    def internal_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "internal_port", value)

    @property
    @pulumi.getter(name="internalPortId")
    def internal_port_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Neutron port associated with the port forwarding. Changing
        this updates the `internal_port_id` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_port_id")

    @internal_port_id.setter
    def internal_port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_port_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The IP protocol used in the port forwarding. Changing this updates the `protocol`
        of an existing port forwarding.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a port forwarding. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        port forwarding.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class PortForwardingV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[int]] = None,
                 floatingip_id: Optional[pulumi.Input[str]] = None,
                 internal_ip_address: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[int]] = None,
                 internal_port_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V2 portforwarding resource within OpenStack.

        ## Example Usage
        ### Simple portforwarding

        ```python
        import pulumi
        import pulumi_openstack as openstack

        pf1 = openstack.networking.PortForwardingV2("pf1",
            external_port=7233,
            floatingip_id="7a52eb59-7d47-415d-a884-046666a6fbae",
            internal_port=25,
            internal_port_id="b930d7f6-ceb7-40a0-8b81-a425dd994ccf",
            protocol="tcp")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A text describing the port forwarding. Changing this
               updates the `description` of an existing port forwarding.
        :param pulumi.Input[int] external_port: The TCP/UDP/other protocol port number of the port forwarding. Changing this
               updates the `external_port` of an existing port forwarding.
        :param pulumi.Input[str] floatingip_id: The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        :param pulumi.Input[str] internal_ip_address: The fixed IPv4 address of the Neutron port associated with the port forwarding.
               Changing this updates the `internal_ip_address` of an existing port forwarding.
        :param pulumi.Input[int] internal_port: The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
               port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        :param pulumi.Input[str] internal_port_id: The ID of the Neutron port associated with the port forwarding. Changing
               this updates the `internal_port_id` of an existing port forwarding.
        :param pulumi.Input[str] protocol: The IP protocol used in the port forwarding. Changing this updates the `protocol`
               of an existing port forwarding.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port forwarding. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               port forwarding.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PortForwardingV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 portforwarding resource within OpenStack.

        ## Example Usage
        ### Simple portforwarding

        ```python
        import pulumi
        import pulumi_openstack as openstack

        pf1 = openstack.networking.PortForwardingV2("pf1",
            external_port=7233,
            floatingip_id="7a52eb59-7d47-415d-a884-046666a6fbae",
            internal_port=25,
            internal_port_id="b930d7f6-ceb7-40a0-8b81-a425dd994ccf",
            protocol="tcp")
        ```

        :param str resource_name: The name of the resource.
        :param PortForwardingV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PortForwardingV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            PortForwardingV2Args._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_port: Optional[pulumi.Input[int]] = None,
                 floatingip_id: Optional[pulumi.Input[str]] = None,
                 internal_ip_address: Optional[pulumi.Input[str]] = None,
                 internal_port: Optional[pulumi.Input[int]] = None,
                 internal_port_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PortForwardingV2Args.__new__(PortForwardingV2Args)

            __props__.__dict__["description"] = description
            if external_port is None and not opts.urn:
                raise TypeError("Missing required property 'external_port'")
            __props__.__dict__["external_port"] = external_port
            if floatingip_id is None and not opts.urn:
                raise TypeError("Missing required property 'floatingip_id'")
            __props__.__dict__["floatingip_id"] = floatingip_id
            if internal_ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'internal_ip_address'")
            __props__.__dict__["internal_ip_address"] = internal_ip_address
            if internal_port is None and not opts.urn:
                raise TypeError("Missing required property 'internal_port'")
            __props__.__dict__["internal_port"] = internal_port
            if internal_port_id is None and not opts.urn:
                raise TypeError("Missing required property 'internal_port_id'")
            __props__.__dict__["internal_port_id"] = internal_port_id
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
        super(PortForwardingV2, __self__).__init__(
            'openstack:networking/portForwardingV2:PortForwardingV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_port: Optional[pulumi.Input[int]] = None,
            floatingip_id: Optional[pulumi.Input[str]] = None,
            internal_ip_address: Optional[pulumi.Input[str]] = None,
            internal_port: Optional[pulumi.Input[int]] = None,
            internal_port_id: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'PortForwardingV2':
        """
        Get an existing PortForwardingV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A text describing the port forwarding. Changing this
               updates the `description` of an existing port forwarding.
        :param pulumi.Input[int] external_port: The TCP/UDP/other protocol port number of the port forwarding. Changing this
               updates the `external_port` of an existing port forwarding.
        :param pulumi.Input[str] floatingip_id: The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        :param pulumi.Input[str] internal_ip_address: The fixed IPv4 address of the Neutron port associated with the port forwarding.
               Changing this updates the `internal_ip_address` of an existing port forwarding.
        :param pulumi.Input[int] internal_port: The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
               port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        :param pulumi.Input[str] internal_port_id: The ID of the Neutron port associated with the port forwarding. Changing
               this updates the `internal_port_id` of an existing port forwarding.
        :param pulumi.Input[str] protocol: The IP protocol used in the port forwarding. Changing this updates the `protocol`
               of an existing port forwarding.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port forwarding. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               port forwarding.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PortForwardingV2State.__new__(_PortForwardingV2State)

        __props__.__dict__["description"] = description
        __props__.__dict__["external_port"] = external_port
        __props__.__dict__["floatingip_id"] = floatingip_id
        __props__.__dict__["internal_ip_address"] = internal_ip_address
        __props__.__dict__["internal_port"] = internal_port
        __props__.__dict__["internal_port_id"] = internal_port_id
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        return PortForwardingV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A text describing the port forwarding. Changing this
        updates the `description` of an existing port forwarding.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> pulumi.Output[int]:
        """
        The TCP/UDP/other protocol port number of the port forwarding. Changing this
        updates the `external_port` of an existing port forwarding.
        """
        return pulumi.get(self, "external_port")

    @property
    @pulumi.getter(name="floatingipId")
    def floatingip_id(self) -> pulumi.Output[str]:
        """
        The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
        """
        return pulumi.get(self, "floatingip_id")

    @property
    @pulumi.getter(name="internalIpAddress")
    def internal_ip_address(self) -> pulumi.Output[str]:
        """
        The fixed IPv4 address of the Neutron port associated with the port forwarding.
        Changing this updates the `internal_ip_address` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_ip_address")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> pulumi.Output[int]:
        """
        The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
        port forwarding. Changing this updates the `internal_port` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="internalPortId")
    def internal_port_id(self) -> pulumi.Output[str]:
        """
        The ID of the Neutron port associated with the port forwarding. Changing
        this updates the `internal_port_id` of an existing port forwarding.
        """
        return pulumi.get(self, "internal_port_id")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The IP protocol used in the port forwarding. Changing this updates the `protocol`
        of an existing port forwarding.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a port forwarding. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        port forwarding.
        """
        return pulumi.get(self, "region")

