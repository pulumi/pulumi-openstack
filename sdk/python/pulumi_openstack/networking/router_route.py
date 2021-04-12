# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['RouterRouteArgs', 'RouterRoute']

@pulumi.input_type
class RouterRouteArgs:
    def __init__(__self__, *,
                 destination_cidr: pulumi.Input[str],
                 next_hop: pulumi.Input[str],
                 router_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouterRoute resource.
        :param pulumi.Input[str] destination_cidr: CIDR block to match on the packet’s destination IP. Changing
               this creates a new routing entry.
        :param pulumi.Input[str] next_hop: IP address of the next hop gateway.  Changing
               this creates a new routing entry.
        :param pulumi.Input[str] router_id: ID of the router this routing entry belongs to. Changing
               this creates a new routing entry.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        """
        pulumi.set(__self__, "destination_cidr", destination_cidr)
        pulumi.set(__self__, "next_hop", next_hop)
        pulumi.set(__self__, "router_id", router_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> pulumi.Input[str]:
        """
        CIDR block to match on the packet’s destination IP. Changing
        this creates a new routing entry.
        """
        return pulumi.get(self, "destination_cidr")

    @destination_cidr.setter
    def destination_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr", value)

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> pulumi.Input[str]:
        """
        IP address of the next hop gateway.  Changing
        this creates a new routing entry.
        """
        return pulumi.get(self, "next_hop")

    @next_hop.setter
    def next_hop(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_hop", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Input[str]:
        """
        ID of the router this routing entry belongs to. Changing
        this creates a new routing entry.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to configure a routing entry on a router. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        routing entry.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class RouterRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a routing entry on a OpenStack V2 router.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        router1 = openstack.networking.Router("router1", admin_state_up=True)
        network1 = openstack.networking.Network("network1", admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        int1 = openstack.networking.RouterInterface("int1",
            router_id=router1.id,
            subnet_id=subnet1.id)
        router_route1 = openstack.networking.RouterRoute("routerRoute1",
            destination_cidr="10.0.1.0/24",
            next_hop="192.168.199.254",
            router_id=router1.id,
            opts=pulumi.ResourceOptions(depends_on=["openstack_networking_router_interface_v2.int_1"]))
        ```
        ## Notes

        The `next_hop` IP address must be directly reachable from the router at the ``networking.RouterRoute``
        resource creation time.  You can ensure that by explicitly specifying a dependency on the ``networking.RouterInterface``
        resource that connects the next hop to the router, as in the example above.

        ## Import

        Routing entries can be imported using a combined ID using the following format``<router_id>-route-<destination_cidr>-<next_hop>``

        ```sh
         $ pulumi import openstack:networking/routerRoute:RouterRoute router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr: CIDR block to match on the packet’s destination IP. Changing
               this creates a new routing entry.
        :param pulumi.Input[str] next_hop: IP address of the next hop gateway.  Changing
               this creates a new routing entry.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        :param pulumi.Input[str] router_id: ID of the router this routing entry belongs to. Changing
               this creates a new routing entry.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouterRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a routing entry on a OpenStack V2 router.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        router1 = openstack.networking.Router("router1", admin_state_up=True)
        network1 = openstack.networking.Network("network1", admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        int1 = openstack.networking.RouterInterface("int1",
            router_id=router1.id,
            subnet_id=subnet1.id)
        router_route1 = openstack.networking.RouterRoute("routerRoute1",
            destination_cidr="10.0.1.0/24",
            next_hop="192.168.199.254",
            router_id=router1.id,
            opts=pulumi.ResourceOptions(depends_on=["openstack_networking_router_interface_v2.int_1"]))
        ```
        ## Notes

        The `next_hop` IP address must be directly reachable from the router at the ``networking.RouterRoute``
        resource creation time.  You can ensure that by explicitly specifying a dependency on the ``networking.RouterInterface``
        resource that connects the next hop to the router, as in the example above.

        ## Import

        Routing entries can be imported using a combined ID using the following format``<router_id>-route-<destination_cidr>-<next_hop>``

        ```sh
         $ pulumi import openstack:networking/routerRoute:RouterRoute router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
        ```

        :param str resource_name: The name of the resource.
        :param RouterRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
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

            if destination_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr'")
            __props__['destination_cidr'] = destination_cidr
            if next_hop is None and not opts.urn:
                raise TypeError("Missing required property 'next_hop'")
            __props__['next_hop'] = next_hop
            __props__['region'] = region
            if router_id is None and not opts.urn:
                raise TypeError("Missing required property 'router_id'")
            __props__['router_id'] = router_id
        super(RouterRoute, __self__).__init__(
            'openstack:networking/routerRoute:RouterRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidr: Optional[pulumi.Input[str]] = None,
            next_hop: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            router_id: Optional[pulumi.Input[str]] = None) -> 'RouterRoute':
        """
        Get an existing RouterRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr: CIDR block to match on the packet’s destination IP. Changing
               this creates a new routing entry.
        :param pulumi.Input[str] next_hop: IP address of the next hop gateway.  Changing
               this creates a new routing entry.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        :param pulumi.Input[str] router_id: ID of the router this routing entry belongs to. Changing
               this creates a new routing entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_cidr"] = destination_cidr
        __props__["next_hop"] = next_hop
        __props__["region"] = region
        __props__["router_id"] = router_id
        return RouterRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> pulumi.Output[str]:
        """
        CIDR block to match on the packet’s destination IP. Changing
        this creates a new routing entry.
        """
        return pulumi.get(self, "destination_cidr")

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> pulumi.Output[str]:
        """
        IP address of the next hop gateway.  Changing
        this creates a new routing entry.
        """
        return pulumi.get(self, "next_hop")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to configure a routing entry on a router. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        routing entry.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[str]:
        """
        ID of the router this routing entry belongs to. Changing
        this creates a new routing entry.
        """
        return pulumi.get(self, "router_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

