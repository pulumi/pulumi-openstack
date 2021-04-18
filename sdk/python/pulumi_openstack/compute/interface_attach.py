# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InterfaceAttachArgs', 'InterfaceAttach']

@pulumi.input_type
class InterfaceAttachArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InterfaceAttach resource.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] fixed_ip: An IP address to assosciate with the port.
               _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created automatically.
               _NOTE_: This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               _NOTE_: This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the interface attachment.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new attachment.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the Instance to attach the Port or Network to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address to assosciate with the port.
        _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        """
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Network to attach to an Instance. A port will be created automatically.
        _NOTE_: This option and `port_id` are mutually exclusive.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Port to attach to an Instance.
        _NOTE_: This option and `network_id` are mutually exclusive.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the interface attachment.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new attachment.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _InterfaceAttachState:
    def __init__(__self__, *,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InterfaceAttach resources.
        :param pulumi.Input[str] fixed_ip: An IP address to assosciate with the port.
               _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created automatically.
               _NOTE_: This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               _NOTE_: This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the interface attachment.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new attachment.
        """
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        """
        An IP address to assosciate with the port.
        _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        """
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Instance to attach the Port or Network to.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Network to attach to an Instance. A port will be created automatically.
        _NOTE_: This option and `port_id` are mutually exclusive.
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Port to attach to an Instance.
        _NOTE_: This option and `network_id` are mutually exclusive.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the interface attachment.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new attachment.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class InterfaceAttach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Attaches a Network Interface (a Port) to an Instance using the OpenStack
        Compute (Nova) v2 API.

        ## Example Usage
        ### Basic Attachment

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        instance1 = openstack.compute.Instance("instance1", security_groups=["default"])
        ai1 = openstack.compute.InterfaceAttach("ai1",
            instance_id=instance1.id,
            network_id=openstack_networking_port_v2["network_1"]["id"])
        ```
        ### Attachment Specifying a Fixed IP

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        instance1 = openstack.compute.Instance("instance1", security_groups=["default"])
        ai1 = openstack.compute.InterfaceAttach("ai1",
            fixed_ip="10.0.10.10",
            instance_id=instance1.id,
            network_id=openstack_networking_port_v2["network_1"]["id"])
        ```
        ### Attachment Using an Existing Port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        port1 = openstack.networking.Port("port1",
            admin_state_up=True,
            network_id=network1.id)
        instance1 = openstack.compute.Instance("instance1", security_groups=["default"])
        ai1 = openstack.compute.InterfaceAttach("ai1",
            instance_id=instance1.id,
            port_id=port1.id)
        ```

        ## Import

        Interface Attachments can be imported using the Instance ID and Port ID separated by a slash, e.g.

        ```sh
         $ pulumi import openstack:compute/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: An IP address to assosciate with the port.
               _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created automatically.
               _NOTE_: This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               _NOTE_: This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the interface attachment.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new attachment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InterfaceAttachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attaches a Network Interface (a Port) to an Instance using the OpenStack
        Compute (Nova) v2 API.

        ## Example Usage
        ### Basic Attachment

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        instance1 = openstack.compute.Instance("instance1", security_groups=["default"])
        ai1 = openstack.compute.InterfaceAttach("ai1",
            instance_id=instance1.id,
            network_id=openstack_networking_port_v2["network_1"]["id"])
        ```
        ### Attachment Specifying a Fixed IP

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        instance1 = openstack.compute.Instance("instance1", security_groups=["default"])
        ai1 = openstack.compute.InterfaceAttach("ai1",
            fixed_ip="10.0.10.10",
            instance_id=instance1.id,
            network_id=openstack_networking_port_v2["network_1"]["id"])
        ```
        ### Attachment Using an Existing Port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        port1 = openstack.networking.Port("port1",
            admin_state_up=True,
            network_id=network1.id)
        instance1 = openstack.compute.Instance("instance1", security_groups=["default"])
        ai1 = openstack.compute.InterfaceAttach("ai1",
            instance_id=instance1.id,
            port_id=port1.id)
        ```

        ## Import

        Interface Attachments can be imported using the Instance ID and Port ID separated by a slash, e.g.

        ```sh
         $ pulumi import openstack:compute/interfaceAttach:InterfaceAttach ai_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666
        ```

        :param str resource_name: The name of the resource.
        :param InterfaceAttachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InterfaceAttachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = InterfaceAttachArgs.__new__(InterfaceAttachArgs)

            __props__.__dict__["fixed_ip"] = fixed_ip
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["region"] = region
        super(InterfaceAttach, __self__).__init__(
            'openstack:compute/interfaceAttach:InterfaceAttach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fixed_ip: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'InterfaceAttach':
        """
        Get an existing InterfaceAttach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: An IP address to assosciate with the port.
               _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created automatically.
               _NOTE_: This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               _NOTE_: This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the interface attachment.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new attachment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InterfaceAttachState.__new__(_InterfaceAttachState)

        __props__.__dict__["fixed_ip"] = fixed_ip
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["region"] = region
        return InterfaceAttach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> pulumi.Output[str]:
        """
        An IP address to assosciate with the port.
        _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        """
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the Instance to attach the Port or Network to.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The ID of the Network to attach to an Instance. A port will be created automatically.
        _NOTE_: This option and `port_id` are mutually exclusive.
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[str]:
        """
        The ID of the Port to attach to an Instance.
        _NOTE_: This option and `network_id` are mutually exclusive.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the interface attachment.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new attachment.
        """
        return pulumi.get(self, "region")

