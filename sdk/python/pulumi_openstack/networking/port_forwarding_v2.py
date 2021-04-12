# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

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
        pulumi.set(__self__, "external_port", external_port)
        pulumi.set(__self__, "floatingip_id", floatingip_id)
        pulumi.set(__self__, "internal_ip_address", internal_ip_address)
        pulumi.set(__self__, "internal_port", internal_port)
        pulumi.set(__self__, "internal_port_id", internal_port_id)
        pulumi.set(__self__, "protocol", protocol)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)

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
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            if external_port is None and not opts.urn:
                raise TypeError("Missing required property 'external_port'")
            __props__['external_port'] = external_port
            if floatingip_id is None and not opts.urn:
                raise TypeError("Missing required property 'floatingip_id'")
            __props__['floatingip_id'] = floatingip_id
            if internal_ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'internal_ip_address'")
            __props__['internal_ip_address'] = internal_ip_address
            if internal_port is None and not opts.urn:
                raise TypeError("Missing required property 'internal_port'")
            __props__['internal_port'] = internal_port
            if internal_port_id is None and not opts.urn:
                raise TypeError("Missing required property 'internal_port_id'")
            __props__['internal_port_id'] = internal_port_id
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            __props__['region'] = region
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

        __props__ = dict()

        __props__["description"] = description
        __props__["external_port"] = external_port
        __props__["floatingip_id"] = floatingip_id
        __props__["internal_ip_address"] = internal_ip_address
        __props__["internal_port"] = internal_port
        __props__["internal_port_id"] = internal_port_id
        __props__["protocol"] = protocol
        __props__["region"] = region
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

