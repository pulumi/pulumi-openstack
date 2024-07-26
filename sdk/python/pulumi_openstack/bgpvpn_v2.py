# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BgpvpnV2Args', 'BgpvpnV2']

@pulumi.input_type
class BgpvpnV2Args:
    def __init__(__self__, *,
                 export_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 import_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 local_pref: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 route_distinguishers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a BgpvpnV2 resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] export_targets: A list of additional Route Targets that will be
               used for export.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] import_targets: A list of additional Route Targets that will be
               imported.
        :param pulumi.Input[int] local_pref: The default BGP LOCAL\\_PREF of routes that will be
               advertised to the BGP VPN, unless overridden per-route.
        :param pulumi.Input[str] name: The name of the BGP VPN. Changing this updates the name of
               the existing BGP VPN.
        :param pulumi.Input[str] project_id: The ID of the project that owns the BGPVPN. Only
               administrative and users with `advsvc` role can specify a project ID other
               than their own. Changing this creates a new BGP VPN.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_distinguishers: A list of route distinguisher strings. If
               specified, one of these RDs will be used to advertise VPN routes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_targets: A list of Route Targets that will be both
               imported and used for export.
        :param pulumi.Input[str] type: The type of the BGP VPN (either `l2` or `l3`). Changing this
               creates a new BGP VPN. Defaults to `l3`.
        :param pulumi.Input[int] vni: The globally-assigned VXLAN VNI for the BGP VPN. Changing
               this creates a new BGP VPN.
        """
        if export_targets is not None:
            pulumi.set(__self__, "export_targets", export_targets)
        if import_targets is not None:
            pulumi.set(__self__, "import_targets", import_targets)
        if local_pref is not None:
            pulumi.set(__self__, "local_pref", local_pref)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if route_distinguishers is not None:
            pulumi.set(__self__, "route_distinguishers", route_distinguishers)
        if route_targets is not None:
            pulumi.set(__self__, "route_targets", route_targets)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vni is not None:
            pulumi.set(__self__, "vni", vni)

    @property
    @pulumi.getter(name="exportTargets")
    def export_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of additional Route Targets that will be
        used for export.
        """
        return pulumi.get(self, "export_targets")

    @export_targets.setter
    def export_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "export_targets", value)

    @property
    @pulumi.getter(name="importTargets")
    def import_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of additional Route Targets that will be
        imported.
        """
        return pulumi.get(self, "import_targets")

    @import_targets.setter
    def import_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "import_targets", value)

    @property
    @pulumi.getter(name="localPref")
    def local_pref(self) -> Optional[pulumi.Input[int]]:
        """
        The default BGP LOCAL\\_PREF of routes that will be
        advertised to the BGP VPN, unless overridden per-route.
        """
        return pulumi.get(self, "local_pref")

    @local_pref.setter
    def local_pref(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "local_pref", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the BGP VPN. Changing this updates the name of
        the existing BGP VPN.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project that owns the BGPVPN. Only
        administrative and users with `advsvc` role can specify a project ID other
        than their own. Changing this creates a new BGP VPN.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a BGP VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        BGP VPN.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="routeDistinguishers")
    def route_distinguishers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of route distinguisher strings. If
        specified, one of these RDs will be used to advertise VPN routes.
        """
        return pulumi.get(self, "route_distinguishers")

    @route_distinguishers.setter
    def route_distinguishers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "route_distinguishers", value)

    @property
    @pulumi.getter(name="routeTargets")
    def route_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of Route Targets that will be both
        imported and used for export.
        """
        return pulumi.get(self, "route_targets")

    @route_targets.setter
    def route_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "route_targets", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the BGP VPN (either `l2` or `l3`). Changing this
        creates a new BGP VPN. Defaults to `l3`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vni(self) -> Optional[pulumi.Input[int]]:
        """
        The globally-assigned VXLAN VNI for the BGP VPN. Changing
        this creates a new BGP VPN.
        """
        return pulumi.get(self, "vni")

    @vni.setter
    def vni(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vni", value)


@pulumi.input_type
class _BgpvpnV2State:
    def __init__(__self__, *,
                 export_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 import_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 local_pref: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 route_distinguishers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering BgpvpnV2 resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] export_targets: A list of additional Route Targets that will be
               used for export.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] import_targets: A list of additional Route Targets that will be
               imported.
        :param pulumi.Input[int] local_pref: The default BGP LOCAL\\_PREF of routes that will be
               advertised to the BGP VPN, unless overridden per-route.
        :param pulumi.Input[str] name: The name of the BGP VPN. Changing this updates the name of
               the existing BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] networks: A list of network IDs that are associated with the BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ports: A list of port IDs that are associated with the BGP VPN.
        :param pulumi.Input[str] project_id: The ID of the project that owns the BGPVPN. Only
               administrative and users with `advsvc` role can specify a project ID other
               than their own. Changing this creates a new BGP VPN.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_distinguishers: A list of route distinguisher strings. If
               specified, one of these RDs will be used to advertise VPN routes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_targets: A list of Route Targets that will be both
               imported and used for export.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] routers: A list of router IDs that are associated with the BGP VPN.
        :param pulumi.Input[bool] shared: Indicates whether the BGP VPN is shared across projects.
        :param pulumi.Input[str] type: The type of the BGP VPN (either `l2` or `l3`). Changing this
               creates a new BGP VPN. Defaults to `l3`.
        :param pulumi.Input[int] vni: The globally-assigned VXLAN VNI for the BGP VPN. Changing
               this creates a new BGP VPN.
        """
        if export_targets is not None:
            pulumi.set(__self__, "export_targets", export_targets)
        if import_targets is not None:
            pulumi.set(__self__, "import_targets", import_targets)
        if local_pref is not None:
            pulumi.set(__self__, "local_pref", local_pref)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if route_distinguishers is not None:
            pulumi.set(__self__, "route_distinguishers", route_distinguishers)
        if route_targets is not None:
            pulumi.set(__self__, "route_targets", route_targets)
        if routers is not None:
            pulumi.set(__self__, "routers", routers)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vni is not None:
            pulumi.set(__self__, "vni", vni)

    @property
    @pulumi.getter(name="exportTargets")
    def export_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of additional Route Targets that will be
        used for export.
        """
        return pulumi.get(self, "export_targets")

    @export_targets.setter
    def export_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "export_targets", value)

    @property
    @pulumi.getter(name="importTargets")
    def import_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of additional Route Targets that will be
        imported.
        """
        return pulumi.get(self, "import_targets")

    @import_targets.setter
    def import_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "import_targets", value)

    @property
    @pulumi.getter(name="localPref")
    def local_pref(self) -> Optional[pulumi.Input[int]]:
        """
        The default BGP LOCAL\\_PREF of routes that will be
        advertised to the BGP VPN, unless overridden per-route.
        """
        return pulumi.get(self, "local_pref")

    @local_pref.setter
    def local_pref(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "local_pref", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the BGP VPN. Changing this updates the name of
        the existing BGP VPN.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of network IDs that are associated with the BGP VPN.
        """
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of port IDs that are associated with the BGP VPN.
        """
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project that owns the BGPVPN. Only
        administrative and users with `advsvc` role can specify a project ID other
        than their own. Changing this creates a new BGP VPN.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a BGP VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        BGP VPN.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="routeDistinguishers")
    def route_distinguishers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of route distinguisher strings. If
        specified, one of these RDs will be used to advertise VPN routes.
        """
        return pulumi.get(self, "route_distinguishers")

    @route_distinguishers.setter
    def route_distinguishers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "route_distinguishers", value)

    @property
    @pulumi.getter(name="routeTargets")
    def route_targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of Route Targets that will be both
        imported and used for export.
        """
        return pulumi.get(self, "route_targets")

    @route_targets.setter
    def route_targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "route_targets", value)

    @property
    @pulumi.getter
    def routers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of router IDs that are associated with the BGP VPN.
        """
        return pulumi.get(self, "routers")

    @routers.setter
    def routers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "routers", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the BGP VPN is shared across projects.
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the BGP VPN (either `l2` or `l3`). Changing this
        creates a new BGP VPN. Defaults to `l3`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vni(self) -> Optional[pulumi.Input[int]]:
        """
        The globally-assigned VXLAN VNI for the BGP VPN. Changing
        this creates a new BGP VPN.
        """
        return pulumi.get(self, "vni")

    @vni.setter
    def vni(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vni", value)


class BgpvpnV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 export_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 import_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 local_pref: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 route_distinguishers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a V2 BGP VPN service resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        bgpvpn1 = openstack.BgpvpnV2("bgpvpn_1",
            name="bgpvpn1",
            route_distinguishers=["64512:1"],
            route_targets=["64512:1"],
            import_targets=["64512:2"],
            export_targets=["64512:3"])
        ```

        ## Import

        BGP VPNs can be imported using the `id`, e.g.

        hcl

        ```sh
        $ pulumi import openstack:index/bgpvpnV2:BgpvpnV2 bgpvpn_1 1eec2c66-6be2-4305-af3f-354c9b81f18c
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] export_targets: A list of additional Route Targets that will be
               used for export.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] import_targets: A list of additional Route Targets that will be
               imported.
        :param pulumi.Input[int] local_pref: The default BGP LOCAL\\_PREF of routes that will be
               advertised to the BGP VPN, unless overridden per-route.
        :param pulumi.Input[str] name: The name of the BGP VPN. Changing this updates the name of
               the existing BGP VPN.
        :param pulumi.Input[str] project_id: The ID of the project that owns the BGPVPN. Only
               administrative and users with `advsvc` role can specify a project ID other
               than their own. Changing this creates a new BGP VPN.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_distinguishers: A list of route distinguisher strings. If
               specified, one of these RDs will be used to advertise VPN routes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_targets: A list of Route Targets that will be both
               imported and used for export.
        :param pulumi.Input[str] type: The type of the BGP VPN (either `l2` or `l3`). Changing this
               creates a new BGP VPN. Defaults to `l3`.
        :param pulumi.Input[int] vni: The globally-assigned VXLAN VNI for the BGP VPN. Changing
               this creates a new BGP VPN.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BgpvpnV2Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 BGP VPN service resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        bgpvpn1 = openstack.BgpvpnV2("bgpvpn_1",
            name="bgpvpn1",
            route_distinguishers=["64512:1"],
            route_targets=["64512:1"],
            import_targets=["64512:2"],
            export_targets=["64512:3"])
        ```

        ## Import

        BGP VPNs can be imported using the `id`, e.g.

        hcl

        ```sh
        $ pulumi import openstack:index/bgpvpnV2:BgpvpnV2 bgpvpn_1 1eec2c66-6be2-4305-af3f-354c9b81f18c
        ```

        :param str resource_name: The name of the resource.
        :param BgpvpnV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpvpnV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 export_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 import_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 local_pref: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 route_distinguishers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 route_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpvpnV2Args.__new__(BgpvpnV2Args)

            __props__.__dict__["export_targets"] = export_targets
            __props__.__dict__["import_targets"] = import_targets
            __props__.__dict__["local_pref"] = local_pref
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["route_distinguishers"] = route_distinguishers
            __props__.__dict__["route_targets"] = route_targets
            __props__.__dict__["type"] = type
            __props__.__dict__["vni"] = vni
            __props__.__dict__["networks"] = None
            __props__.__dict__["ports"] = None
            __props__.__dict__["routers"] = None
            __props__.__dict__["shared"] = None
        super(BgpvpnV2, __self__).__init__(
            'openstack:index/bgpvpnV2:BgpvpnV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            export_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            import_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            local_pref: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            route_distinguishers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            route_targets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vni: Optional[pulumi.Input[int]] = None) -> 'BgpvpnV2':
        """
        Get an existing BgpvpnV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] export_targets: A list of additional Route Targets that will be
               used for export.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] import_targets: A list of additional Route Targets that will be
               imported.
        :param pulumi.Input[int] local_pref: The default BGP LOCAL\\_PREF of routes that will be
               advertised to the BGP VPN, unless overridden per-route.
        :param pulumi.Input[str] name: The name of the BGP VPN. Changing this updates the name of
               the existing BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] networks: A list of network IDs that are associated with the BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ports: A list of port IDs that are associated with the BGP VPN.
        :param pulumi.Input[str] project_id: The ID of the project that owns the BGPVPN. Only
               administrative and users with `advsvc` role can specify a project ID other
               than their own. Changing this creates a new BGP VPN.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a BGP VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               BGP VPN.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_distinguishers: A list of route distinguisher strings. If
               specified, one of these RDs will be used to advertise VPN routes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] route_targets: A list of Route Targets that will be both
               imported and used for export.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] routers: A list of router IDs that are associated with the BGP VPN.
        :param pulumi.Input[bool] shared: Indicates whether the BGP VPN is shared across projects.
        :param pulumi.Input[str] type: The type of the BGP VPN (either `l2` or `l3`). Changing this
               creates a new BGP VPN. Defaults to `l3`.
        :param pulumi.Input[int] vni: The globally-assigned VXLAN VNI for the BGP VPN. Changing
               this creates a new BGP VPN.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpvpnV2State.__new__(_BgpvpnV2State)

        __props__.__dict__["export_targets"] = export_targets
        __props__.__dict__["import_targets"] = import_targets
        __props__.__dict__["local_pref"] = local_pref
        __props__.__dict__["name"] = name
        __props__.__dict__["networks"] = networks
        __props__.__dict__["ports"] = ports
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["route_distinguishers"] = route_distinguishers
        __props__.__dict__["route_targets"] = route_targets
        __props__.__dict__["routers"] = routers
        __props__.__dict__["shared"] = shared
        __props__.__dict__["type"] = type
        __props__.__dict__["vni"] = vni
        return BgpvpnV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="exportTargets")
    def export_targets(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of additional Route Targets that will be
        used for export.
        """
        return pulumi.get(self, "export_targets")

    @property
    @pulumi.getter(name="importTargets")
    def import_targets(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of additional Route Targets that will be
        imported.
        """
        return pulumi.get(self, "import_targets")

    @property
    @pulumi.getter(name="localPref")
    def local_pref(self) -> pulumi.Output[Optional[int]]:
        """
        The default BGP LOCAL\\_PREF of routes that will be
        advertised to the BGP VPN, unless overridden per-route.
        """
        return pulumi.get(self, "local_pref")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the BGP VPN. Changing this updates the name of
        the existing BGP VPN.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of network IDs that are associated with the BGP VPN.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def ports(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of port IDs that are associated with the BGP VPN.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project that owns the BGPVPN. Only
        administrative and users with `advsvc` role can specify a project ID other
        than their own. Changing this creates a new BGP VPN.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a BGP VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        BGP VPN.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routeDistinguishers")
    def route_distinguishers(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of route distinguisher strings. If
        specified, one of these RDs will be used to advertise VPN routes.
        """
        return pulumi.get(self, "route_distinguishers")

    @property
    @pulumi.getter(name="routeTargets")
    def route_targets(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of Route Targets that will be both
        imported and used for export.
        """
        return pulumi.get(self, "route_targets")

    @property
    @pulumi.getter
    def routers(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of router IDs that are associated with the BGP VPN.
        """
        return pulumi.get(self, "routers")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[bool]:
        """
        Indicates whether the BGP VPN is shared across projects.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the BGP VPN (either `l2` or `l3`). Changing this
        creates a new BGP VPN. Defaults to `l3`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vni(self) -> pulumi.Output[Optional[int]]:
        """
        The globally-assigned VXLAN VNI for the BGP VPN. Changing
        this creates a new BGP VPN.
        """
        return pulumi.get(self, "vni")
