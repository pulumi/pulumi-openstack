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

__all__ = ['RouterAssociateV2Args', 'RouterAssociateV2']

@pulumi.input_type
class RouterAssociateV2Args:
    def __init__(__self__, *,
                 bgpvpn_id: pulumi.Input[builtins.str],
                 router_id: pulumi.Input[builtins.str],
                 advertise_extra_routes: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a RouterAssociateV2 resource.
        :param pulumi.Input[builtins.str] bgpvpn_id: The ID of the BGP VPN to which the router will be
               associated. Changing this creates a new BGP VPN router association.
        :param pulumi.Input[builtins.str] router_id: The ID of the router to be associated with the BGP
               VPN. Changing this creates a new BGP VPN router association.
        :param pulumi.Input[builtins.bool] advertise_extra_routes: A boolean flag indicating whether extra
               routes should be advertised. Defaults to true.
        :param pulumi.Input[builtins.str] project_id: The ID of the project that owns the BGP VPN router
               association. Only administrative and users with `advsvc` role can specify a
               project ID other than their own. Changing this creates a new BGP VPN router
               association.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN router association. If
               omitted, the `region` argument of the provider is used. Changing this creates
               a new BGP VPN router association.
        """
        pulumi.set(__self__, "bgpvpn_id", bgpvpn_id)
        pulumi.set(__self__, "router_id", router_id)
        if advertise_extra_routes is not None:
            pulumi.set(__self__, "advertise_extra_routes", advertise_extra_routes)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="bgpvpnId")
    def bgpvpn_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the BGP VPN to which the router will be
        associated. Changing this creates a new BGP VPN router association.
        """
        return pulumi.get(self, "bgpvpn_id")

    @bgpvpn_id.setter
    def bgpvpn_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "bgpvpn_id", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the router to be associated with the BGP
        VPN. Changing this creates a new BGP VPN router association.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter(name="advertiseExtraRoutes")
    def advertise_extra_routes(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A boolean flag indicating whether extra
        routes should be advertised. Defaults to true.
        """
        return pulumi.get(self, "advertise_extra_routes")

    @advertise_extra_routes.setter
    def advertise_extra_routes(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "advertise_extra_routes", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the project that owns the BGP VPN router
        association. Only administrative and users with `advsvc` role can specify a
        project ID other than their own. Changing this creates a new BGP VPN router
        association.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a BGP VPN router association. If
        omitted, the `region` argument of the provider is used. Changing this creates
        a new BGP VPN router association.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _RouterAssociateV2State:
    def __init__(__self__, *,
                 advertise_extra_routes: Optional[pulumi.Input[builtins.bool]] = None,
                 bgpvpn_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering RouterAssociateV2 resources.
        :param pulumi.Input[builtins.bool] advertise_extra_routes: A boolean flag indicating whether extra
               routes should be advertised. Defaults to true.
        :param pulumi.Input[builtins.str] bgpvpn_id: The ID of the BGP VPN to which the router will be
               associated. Changing this creates a new BGP VPN router association.
        :param pulumi.Input[builtins.str] project_id: The ID of the project that owns the BGP VPN router
               association. Only administrative and users with `advsvc` role can specify a
               project ID other than their own. Changing this creates a new BGP VPN router
               association.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN router association. If
               omitted, the `region` argument of the provider is used. Changing this creates
               a new BGP VPN router association.
        :param pulumi.Input[builtins.str] router_id: The ID of the router to be associated with the BGP
               VPN. Changing this creates a new BGP VPN router association.
        """
        if advertise_extra_routes is not None:
            pulumi.set(__self__, "advertise_extra_routes", advertise_extra_routes)
        if bgpvpn_id is not None:
            pulumi.set(__self__, "bgpvpn_id", bgpvpn_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if router_id is not None:
            pulumi.set(__self__, "router_id", router_id)

    @property
    @pulumi.getter(name="advertiseExtraRoutes")
    def advertise_extra_routes(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        A boolean flag indicating whether extra
        routes should be advertised. Defaults to true.
        """
        return pulumi.get(self, "advertise_extra_routes")

    @advertise_extra_routes.setter
    def advertise_extra_routes(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "advertise_extra_routes", value)

    @property
    @pulumi.getter(name="bgpvpnId")
    def bgpvpn_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the BGP VPN to which the router will be
        associated. Changing this creates a new BGP VPN router association.
        """
        return pulumi.get(self, "bgpvpn_id")

    @bgpvpn_id.setter
    def bgpvpn_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "bgpvpn_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the project that owns the BGP VPN router
        association. Only administrative and users with `advsvc` role can specify a
        project ID other than their own. Changing this creates a new BGP VPN router
        association.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a BGP VPN router association. If
        omitted, the `region` argument of the provider is used. Changing this creates
        a new BGP VPN router association.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the router to be associated with the BGP
        VPN. Changing this creates a new BGP VPN router association.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "router_id", value)


class RouterAssociateV2(pulumi.CustomResource):

    pulumi_type = "openstack:bgpvpn/routerAssociateV2:RouterAssociateV2"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advertise_extra_routes: Optional[pulumi.Input[builtins.bool]] = None,
                 bgpvpn_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a V2 BGP VPN router association resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        association1 = openstack.bgpvpn.RouterAssociateV2("association_1",
            bgpvpn_id="d57d39e1-dc63-44fd-8cbd-a4e1488100c5",
            router_id="423fa80f-e0d7-4d02-a9a5-8b8c05812bf6")
        ```

        ## Import

        BGP VPN router associations can be imported using the BGP VPN ID and BGP VPN

        router association ID separated by a slash, e.g.:

        hcl

        ```sh
        $ pulumi import openstack:bgpvpn/routerAssociateV2:RouterAssociateV2 association_1 e26d509e-fc2d-4fb5-8562-619911a9a6bc/3cc9df2d-80db-4536-8ba6-295d1d0f723f
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] advertise_extra_routes: A boolean flag indicating whether extra
               routes should be advertised. Defaults to true.
        :param pulumi.Input[builtins.str] bgpvpn_id: The ID of the BGP VPN to which the router will be
               associated. Changing this creates a new BGP VPN router association.
        :param pulumi.Input[builtins.str] project_id: The ID of the project that owns the BGP VPN router
               association. Only administrative and users with `advsvc` role can specify a
               project ID other than their own. Changing this creates a new BGP VPN router
               association.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN router association. If
               omitted, the `region` argument of the provider is used. Changing this creates
               a new BGP VPN router association.
        :param pulumi.Input[builtins.str] router_id: The ID of the router to be associated with the BGP
               VPN. Changing this creates a new BGP VPN router association.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouterAssociateV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 BGP VPN router association resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        association1 = openstack.bgpvpn.RouterAssociateV2("association_1",
            bgpvpn_id="d57d39e1-dc63-44fd-8cbd-a4e1488100c5",
            router_id="423fa80f-e0d7-4d02-a9a5-8b8c05812bf6")
        ```

        ## Import

        BGP VPN router associations can be imported using the BGP VPN ID and BGP VPN

        router association ID separated by a slash, e.g.:

        hcl

        ```sh
        $ pulumi import openstack:bgpvpn/routerAssociateV2:RouterAssociateV2 association_1 e26d509e-fc2d-4fb5-8562-619911a9a6bc/3cc9df2d-80db-4536-8ba6-295d1d0f723f
        ```

        :param str resource_name: The name of the resource.
        :param RouterAssociateV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterAssociateV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advertise_extra_routes: Optional[pulumi.Input[builtins.bool]] = None,
                 bgpvpn_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouterAssociateV2Args.__new__(RouterAssociateV2Args)

            __props__.__dict__["advertise_extra_routes"] = advertise_extra_routes
            if bgpvpn_id is None and not opts.urn:
                raise TypeError("Missing required property 'bgpvpn_id'")
            __props__.__dict__["bgpvpn_id"] = bgpvpn_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if router_id is None and not opts.urn:
                raise TypeError("Missing required property 'router_id'")
            __props__.__dict__["router_id"] = router_id
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RouterAssociateV2, __self__).__init__(
            'openstack:bgpvpn/routerAssociateV2:RouterAssociateV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            advertise_extra_routes: Optional[pulumi.Input[builtins.bool]] = None,
            bgpvpn_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            router_id: Optional[pulumi.Input[builtins.str]] = None) -> 'RouterAssociateV2':
        """
        Get an existing RouterAssociateV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] advertise_extra_routes: A boolean flag indicating whether extra
               routes should be advertised. Defaults to true.
        :param pulumi.Input[builtins.str] bgpvpn_id: The ID of the BGP VPN to which the router will be
               associated. Changing this creates a new BGP VPN router association.
        :param pulumi.Input[builtins.str] project_id: The ID of the project that owns the BGP VPN router
               association. Only administrative and users with `advsvc` role can specify a
               project ID other than their own. Changing this creates a new BGP VPN router
               association.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN router association. If
               omitted, the `region` argument of the provider is used. Changing this creates
               a new BGP VPN router association.
        :param pulumi.Input[builtins.str] router_id: The ID of the router to be associated with the BGP
               VPN. Changing this creates a new BGP VPN router association.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterAssociateV2State.__new__(_RouterAssociateV2State)

        __props__.__dict__["advertise_extra_routes"] = advertise_extra_routes
        __props__.__dict__["bgpvpn_id"] = bgpvpn_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["router_id"] = router_id
        return RouterAssociateV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="advertiseExtraRoutes")
    def advertise_extra_routes(self) -> pulumi.Output[builtins.bool]:
        """
        A boolean flag indicating whether extra
        routes should be advertised. Defaults to true.
        """
        return pulumi.get(self, "advertise_extra_routes")

    @property
    @pulumi.getter(name="bgpvpnId")
    def bgpvpn_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the BGP VPN to which the router will be
        associated. Changing this creates a new BGP VPN router association.
        """
        return pulumi.get(self, "bgpvpn_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the project that owns the BGP VPN router
        association. Only administrative and users with `advsvc` role can specify a
        project ID other than their own. Changing this creates a new BGP VPN router
        association.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a BGP VPN router association. If
        omitted, the `region` argument of the provider is used. Changing this creates
        a new BGP VPN router association.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the router to be associated with the BGP
        VPN. Changing this creates a new BGP VPN router association.
        """
        return pulumi.get(self, "router_id")

