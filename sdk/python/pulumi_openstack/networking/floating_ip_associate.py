# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FloatingIpAssociateArgs', 'FloatingIpAssociate']

@pulumi.input_type
class FloatingIpAssociateArgs:
    def __init__(__self__, *,
                 floating_ip: pulumi.Input[str],
                 port_id: pulumi.Input[str],
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FloatingIpAssociate resource.
        :param pulumi.Input[str] floating_ip: IP Address of an existing floating IP.
        :param pulumi.Input[str] port_id: ID of an existing port with at least one IP address to
               associate with this floating IP.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a floating IP that can be used with
               another networking resource, such as a load balancer. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               floating IP (which may or may not have a different address).
        """
        pulumi.set(__self__, "floating_ip", floating_ip)
        pulumi.set(__self__, "port_id", port_id)
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> pulumi.Input[str]:
        """
        IP Address of an existing floating IP.
        """
        return pulumi.get(self, "floating_ip")

    @floating_ip.setter
    def floating_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "floating_ip", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Input[str]:
        """
        ID of an existing port with at least one IP address to
        associate with this floating IP.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a floating IP that can be used with
        another networking resource, such as a load balancer. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        floating IP (which may or may not have a different address).
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _FloatingIpAssociateState:
    def __init__(__self__, *,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FloatingIpAssociate resources.
        :param pulumi.Input[str] floating_ip: IP Address of an existing floating IP.
        :param pulumi.Input[str] port_id: ID of an existing port with at least one IP address to
               associate with this floating IP.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a floating IP that can be used with
               another networking resource, such as a load balancer. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               floating IP (which may or may not have a different address).
        """
        if fixed_ip is not None:
            pulumi.set(__self__, "fixed_ip", fixed_ip)
        if floating_ip is not None:
            pulumi.set(__self__, "floating_ip", floating_ip)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP Address of an existing floating IP.
        """
        return pulumi.get(self, "floating_ip")

    @floating_ip.setter
    def floating_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floating_ip", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of an existing port with at least one IP address to
        associate with this floating IP.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a floating IP that can be used with
        another networking resource, such as a load balancer. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        floating IP (which may or may not have a different address).
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class FloatingIpAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Associates a floating IP to a port. This is useful for situations
        where you have a pre-allocated floating IP or are unable to use the
        `networking.FloatingIp` resource to create a floating IP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        port1 = openstack.networking.Port("port1", network_id="a5bbd213-e1d3-49b6-aed1-9df60ea94b9a")
        fip1 = openstack.networking.FloatingIpAssociate("fip1",
            floating_ip="1.2.3.4",
            port_id=port1.id)
        ```

        ## Import

        Floating IP associations can be imported using the `id` of the floating IP, e.g.

        ```sh
         $ pulumi import openstack:networking/floatingIpAssociate:FloatingIpAssociate fip 2c7f39f3-702b-48d1-940c-b50384177ee1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] floating_ip: IP Address of an existing floating IP.
        :param pulumi.Input[str] port_id: ID of an existing port with at least one IP address to
               associate with this floating IP.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a floating IP that can be used with
               another networking resource, such as a load balancer. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               floating IP (which may or may not have a different address).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FloatingIpAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Associates a floating IP to a port. This is useful for situations
        where you have a pre-allocated floating IP or are unable to use the
        `networking.FloatingIp` resource to create a floating IP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        port1 = openstack.networking.Port("port1", network_id="a5bbd213-e1d3-49b6-aed1-9df60ea94b9a")
        fip1 = openstack.networking.FloatingIpAssociate("fip1",
            floating_ip="1.2.3.4",
            port_id=port1.id)
        ```

        ## Import

        Floating IP associations can be imported using the `id` of the floating IP, e.g.

        ```sh
         $ pulumi import openstack:networking/floatingIpAssociate:FloatingIpAssociate fip 2c7f39f3-702b-48d1-940c-b50384177ee1
        ```

        :param str resource_name: The name of the resource.
        :param FloatingIpAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FloatingIpAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input[str]] = None,
                 port_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FloatingIpAssociateArgs.__new__(FloatingIpAssociateArgs)

            __props__.__dict__["fixed_ip"] = fixed_ip
            if floating_ip is None and not opts.urn:
                raise TypeError("Missing required property 'floating_ip'")
            __props__.__dict__["floating_ip"] = floating_ip
            if port_id is None and not opts.urn:
                raise TypeError("Missing required property 'port_id'")
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["region"] = region
        super(FloatingIpAssociate, __self__).__init__(
            'openstack:networking/floatingIpAssociate:FloatingIpAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fixed_ip: Optional[pulumi.Input[str]] = None,
            floating_ip: Optional[pulumi.Input[str]] = None,
            port_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'FloatingIpAssociate':
        """
        Get an existing FloatingIpAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] floating_ip: IP Address of an existing floating IP.
        :param pulumi.Input[str] port_id: ID of an existing port with at least one IP address to
               associate with this floating IP.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a floating IP that can be used with
               another networking resource, such as a load balancer. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               floating IP (which may or may not have a different address).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FloatingIpAssociateState.__new__(_FloatingIpAssociateState)

        __props__.__dict__["fixed_ip"] = fixed_ip
        __props__.__dict__["floating_ip"] = floating_ip
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["region"] = region
        return FloatingIpAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> pulumi.Output[str]:
        """
        IP Address of an existing floating IP.
        """
        return pulumi.get(self, "floating_ip")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[str]:
        """
        ID of an existing port with at least one IP address to
        associate with this floating IP.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a floating IP that can be used with
        another networking resource, such as a load balancer. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        floating IP (which may or may not have a different address).
        """
        return pulumi.get(self, "region")

