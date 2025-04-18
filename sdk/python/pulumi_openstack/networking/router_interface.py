# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['RouterInterfaceArgs', 'RouterInterface']

@pulumi.input_type
class RouterInterfaceArgs:
    def __init__(__self__, *,
                 router_id: pulumi.Input[builtins.str],
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a RouterInterface resource.
        :param pulumi.Input[builtins.str] router_id: ID of the router this interface belongs to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.bool] force_destroy: A boolean indicating whether the routes from the
               corresponding router ID should be deleted so that the router interface can
               be destroyed without any errors. The default value is `false`.
        :param pulumi.Input[builtins.str] port_id: ID of the port this interface connects to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router interface.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet this interface connects to. Changing
               this creates a new router interface.
        """
        pulumi.set(__self__, "router_id", router_id)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the router this interface belongs to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A boolean indicating whether the routes from the
        corresponding router ID should be deleted so that the router interface can
        be destroyed without any errors. The default value is `false`.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the port this interface connects to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a router. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        router interface.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the subnet this interface connects to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class _RouterInterfaceState:
    def __init__(__self__, *,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router_id: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering RouterInterface resources.
        :param pulumi.Input[builtins.bool] force_destroy: A boolean indicating whether the routes from the
               corresponding router ID should be deleted so that the router interface can
               be destroyed without any errors. The default value is `false`.
        :param pulumi.Input[builtins.str] port_id: ID of the port this interface connects to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router interface.
        :param pulumi.Input[builtins.str] router_id: ID of the router this interface belongs to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet this interface connects to. Changing
               this creates a new router interface.
        """
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if router_id is not None:
            pulumi.set(__self__, "router_id", router_id)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A boolean indicating whether the routes from the
        corresponding router ID should be deleted so that the router interface can
        be destroyed without any errors. The default value is `false`.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the port this interface connects to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "port_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a router. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        router interface.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the router this interface belongs to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the subnet this interface connects to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "subnet_id", value)


class RouterInterface(pulumi.CustomResource):

    pulumi_type = "openstack:networking/routerInterface:RouterInterface"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router_id: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a V2 router interface resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network_1",
            name="tf_test_network",
            admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet_1",
            network_id=network1.id,
            cidr="192.168.199.0/24",
            ip_version=4)
        router1 = openstack.networking.Router("router_1",
            name="my_router",
            external_network_id="f67f0d72-0ddf-11e4-9d95-e1f29f417e2f")
        router_interface1 = openstack.networking.RouterInterface("router_interface_1",
            router_id=router1.id,
            subnet_id=subnet1.id)
        ```

        ## Import

        Router Interfaces can be imported using the port `id`, e.g.

        $ openstack port list --router <router name or id>

        ```sh
        $ pulumi import openstack:networking/routerInterface:RouterInterface int_1 port_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] force_destroy: A boolean indicating whether the routes from the
               corresponding router ID should be deleted so that the router interface can
               be destroyed without any errors. The default value is `false`.
        :param pulumi.Input[builtins.str] port_id: ID of the port this interface connects to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router interface.
        :param pulumi.Input[builtins.str] router_id: ID of the router this interface belongs to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet this interface connects to. Changing
               this creates a new router interface.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouterInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 router interface resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network_1",
            name="tf_test_network",
            admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet_1",
            network_id=network1.id,
            cidr="192.168.199.0/24",
            ip_version=4)
        router1 = openstack.networking.Router("router_1",
            name="my_router",
            external_network_id="f67f0d72-0ddf-11e4-9d95-e1f29f417e2f")
        router_interface1 = openstack.networking.RouterInterface("router_interface_1",
            router_id=router1.id,
            subnet_id=subnet1.id)
        ```

        ## Import

        Router Interfaces can be imported using the port `id`, e.g.

        $ openstack port list --router <router name or id>

        ```sh
        $ pulumi import openstack:networking/routerInterface:RouterInterface int_1 port_id
        ```

        :param str resource_name: The name of the resource.
        :param RouterInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router_id: Optional[pulumi.Input[builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouterInterfaceArgs.__new__(RouterInterfaceArgs)

            __props__.__dict__["force_destroy"] = force_destroy
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["region"] = region
            if router_id is None and not opts.urn:
                raise TypeError("Missing required property 'router_id'")
            __props__.__dict__["router_id"] = router_id
            __props__.__dict__["subnet_id"] = subnet_id
        super(RouterInterface, __self__).__init__(
            'openstack:networking/routerInterface:RouterInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
            port_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            router_id: Optional[pulumi.Input[builtins.str]] = None,
            subnet_id: Optional[pulumi.Input[builtins.str]] = None) -> 'RouterInterface':
        """
        Get an existing RouterInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] force_destroy: A boolean indicating whether the routes from the
               corresponding router ID should be deleted so that the router interface can
               be destroyed without any errors. The default value is `false`.
        :param pulumi.Input[builtins.str] port_id: ID of the port this interface connects to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router interface.
        :param pulumi.Input[builtins.str] router_id: ID of the router this interface belongs to. Changing
               this creates a new router interface.
        :param pulumi.Input[builtins.str] subnet_id: ID of the subnet this interface connects to. Changing
               this creates a new router interface.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterInterfaceState.__new__(_RouterInterfaceState)

        __props__.__dict__["force_destroy"] = force_destroy
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["region"] = region
        __props__.__dict__["router_id"] = router_id
        __props__.__dict__["subnet_id"] = subnet_id
        return RouterInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        A boolean indicating whether the routes from the
        corresponding router ID should be deleted so that the router interface can
        be destroyed without any errors. The default value is `false`.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the port this interface connects to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a router. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        router interface.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the router this interface belongs to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the subnet this interface connects to. Changing
        this creates a new router interface.
        """
        return pulumi.get(self, "subnet_id")

