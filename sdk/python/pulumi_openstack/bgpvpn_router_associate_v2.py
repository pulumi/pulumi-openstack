# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BgpvpnRouterAssociateV2Args', 'BgpvpnRouterAssociateV2']

@pulumi.input_type
class BgpvpnRouterAssociateV2Args:
    def __init__(__self__, *,
                 bgpvpn_id: pulumi.Input[str],
                 router_id: pulumi.Input[str],
                 advertise_extra_routes: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BgpvpnRouterAssociateV2 resource.
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
    def bgpvpn_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bgpvpn_id")

    @bgpvpn_id.setter
    def bgpvpn_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bgpvpn_id", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter(name="advertiseExtraRoutes")
    def advertise_extra_routes(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "advertise_extra_routes")

    @advertise_extra_routes.setter
    def advertise_extra_routes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "advertise_extra_routes", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _BgpvpnRouterAssociateV2State:
    def __init__(__self__, *,
                 advertise_extra_routes: Optional[pulumi.Input[bool]] = None,
                 bgpvpn_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpvpnRouterAssociateV2 resources.
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
    def advertise_extra_routes(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "advertise_extra_routes")

    @advertise_extra_routes.setter
    def advertise_extra_routes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "advertise_extra_routes", value)

    @property
    @pulumi.getter(name="bgpvpnId")
    def bgpvpn_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bgpvpn_id")

    @bgpvpn_id.setter
    def bgpvpn_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgpvpn_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router_id", value)


warnings.warn("""openstack.index/bgpvpnrouterassociatev2.BgpvpnRouterAssociateV2 has been deprecated in favor of openstack.bgpvpn/routerassociatev2.RouterAssociateV2""", DeprecationWarning)


class BgpvpnRouterAssociateV2(pulumi.CustomResource):
    warnings.warn("""openstack.index/bgpvpnrouterassociatev2.BgpvpnRouterAssociateV2 has been deprecated in favor of openstack.bgpvpn/routerassociatev2.RouterAssociateV2""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advertise_extra_routes: Optional[pulumi.Input[bool]] = None,
                 bgpvpn_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BgpvpnRouterAssociateV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpvpnRouterAssociateV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BgpvpnRouterAssociateV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BgpvpnRouterAssociateV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpvpnRouterAssociateV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advertise_extra_routes: Optional[pulumi.Input[bool]] = None,
                 bgpvpn_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""BgpvpnRouterAssociateV2 is deprecated: openstack.index/bgpvpnrouterassociatev2.BgpvpnRouterAssociateV2 has been deprecated in favor of openstack.bgpvpn/routerassociatev2.RouterAssociateV2""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpvpnRouterAssociateV2Args.__new__(BgpvpnRouterAssociateV2Args)

            __props__.__dict__["advertise_extra_routes"] = advertise_extra_routes
            if bgpvpn_id is None and not opts.urn:
                raise TypeError("Missing required property 'bgpvpn_id'")
            __props__.__dict__["bgpvpn_id"] = bgpvpn_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if router_id is None and not opts.urn:
                raise TypeError("Missing required property 'router_id'")
            __props__.__dict__["router_id"] = router_id
        super(BgpvpnRouterAssociateV2, __self__).__init__(
            'openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            advertise_extra_routes: Optional[pulumi.Input[bool]] = None,
            bgpvpn_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            router_id: Optional[pulumi.Input[str]] = None) -> 'BgpvpnRouterAssociateV2':
        """
        Get an existing BgpvpnRouterAssociateV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpvpnRouterAssociateV2State.__new__(_BgpvpnRouterAssociateV2State)

        __props__.__dict__["advertise_extra_routes"] = advertise_extra_routes
        __props__.__dict__["bgpvpn_id"] = bgpvpn_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["router_id"] = router_id
        return BgpvpnRouterAssociateV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="advertiseExtraRoutes")
    def advertise_extra_routes(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "advertise_extra_routes")

    @property
    @pulumi.getter(name="bgpvpnId")
    def bgpvpn_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "bgpvpn_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "router_id")

