# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 router_id: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input[str] router_id: The ID of the router. Changing this creates a new service.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the resource. Can either be up(true) or down(false).
               Changing this updates the administrative state of the existing service.
        :param pulumi.Input[str] description: The human-readable description for the service.
               Changing this updates the description of the existing service.
        :param pulumi.Input[str] name: The name of the service. Changing this updates the name of
               the existing service.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               service.
        :param pulumi.Input[str] subnet_id: SubnetID is the ID of the subnet. Default is null.
        :param pulumi.Input[str] tenant_id: The owner of the service. Required if admin wants to
               create a service for another project. Changing this creates a new service.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        pulumi.set(__self__, "router_id", router_id)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Input[str]:
        """
        The ID of the router. Changing this creates a new service.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the resource. Can either be up(true) or down(false).
        Changing this updates the administrative state of the existing service.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description for the service.
        Changing this updates the description of the existing service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service. Changing this updates the name of
        the existing service.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        service.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        SubnetID is the ID of the subnet. Default is null.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the service. Required if admin wants to
        create a service for another project. Changing this creates a new service.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _ServiceState:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_v4_ip: Optional[pulumi.Input[str]] = None,
                 external_v6_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Service resources.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the resource. Can either be up(true) or down(false).
               Changing this updates the administrative state of the existing service.
        :param pulumi.Input[str] description: The human-readable description for the service.
               Changing this updates the description of the existing service.
        :param pulumi.Input[str] external_v4_ip: The read-only external (public) IPv4 address that is used for the VPN service.
        :param pulumi.Input[str] external_v6_ip: The read-only external (public) IPv6 address that is used for the VPN service.
        :param pulumi.Input[str] name: The name of the service. Changing this updates the name of
               the existing service.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               service.
        :param pulumi.Input[str] router_id: The ID of the router. Changing this creates a new service.
        :param pulumi.Input[str] status: Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
        :param pulumi.Input[str] subnet_id: SubnetID is the ID of the subnet. Default is null.
        :param pulumi.Input[str] tenant_id: The owner of the service. Required if admin wants to
               create a service for another project. Changing this creates a new service.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_v4_ip is not None:
            pulumi.set(__self__, "external_v4_ip", external_v4_ip)
        if external_v6_ip is not None:
            pulumi.set(__self__, "external_v6_ip", external_v6_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if router_id is not None:
            pulumi.set(__self__, "router_id", router_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the resource. Can either be up(true) or down(false).
        Changing this updates the administrative state of the existing service.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description for the service.
        Changing this updates the description of the existing service.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalV4Ip")
    def external_v4_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The read-only external (public) IPv4 address that is used for the VPN service.
        """
        return pulumi.get(self, "external_v4_ip")

    @external_v4_ip.setter
    def external_v4_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_v4_ip", value)

    @property
    @pulumi.getter(name="externalV6Ip")
    def external_v6_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The read-only external (public) IPv6 address that is used for the VPN service.
        """
        return pulumi.get(self, "external_v6_ip")

    @external_v6_ip.setter
    def external_v6_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_v6_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service. Changing this updates the name of
        the existing service.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        service.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the router. Changing this creates a new service.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        SubnetID is the ID of the subnet. Default is null.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the service. Required if admin wants to
        create a service for another project. Changing this creates a new service.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a V2 Neutron VPN service resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        service1 = openstack.vpnaas.Service("service_1",
            name="my_service",
            router_id="14a75700-fc03-4602-9294-26ee44f366b3",
            admin_state_up=True)
        ```

        ## Import

        Services can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:vpnaas/service:Service service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the resource. Can either be up(true) or down(false).
               Changing this updates the administrative state of the existing service.
        :param pulumi.Input[str] description: The human-readable description for the service.
               Changing this updates the description of the existing service.
        :param pulumi.Input[str] name: The name of the service. Changing this updates the name of
               the existing service.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               service.
        :param pulumi.Input[str] router_id: The ID of the router. Changing this creates a new service.
        :param pulumi.Input[str] subnet_id: SubnetID is the ID of the subnet. Default is null.
        :param pulumi.Input[str] tenant_id: The owner of the service. Required if admin wants to
               create a service for another project. Changing this creates a new service.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 Neutron VPN service resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        service1 = openstack.vpnaas.Service("service_1",
            name="my_service",
            router_id="14a75700-fc03-4602-9294-26ee44f366b3",
            admin_state_up=True)
        ```

        ## Import

        Services can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:vpnaas/service:Service service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
        ```

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if router_id is None and not opts.urn:
                raise TypeError("Missing required property 'router_id'")
            __props__.__dict__["router_id"] = router_id
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["value_specs"] = value_specs
            __props__.__dict__["external_v4_ip"] = None
            __props__.__dict__["external_v6_ip"] = None
            __props__.__dict__["status"] = None
        super(Service, __self__).__init__(
            'openstack:vpnaas/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_v4_ip: Optional[pulumi.Input[str]] = None,
            external_v6_ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            router_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the resource. Can either be up(true) or down(false).
               Changing this updates the administrative state of the existing service.
        :param pulumi.Input[str] description: The human-readable description for the service.
               Changing this updates the description of the existing service.
        :param pulumi.Input[str] external_v4_ip: The read-only external (public) IPv4 address that is used for the VPN service.
        :param pulumi.Input[str] external_v6_ip: The read-only external (public) IPv6 address that is used for the VPN service.
        :param pulumi.Input[str] name: The name of the service. Changing this updates the name of
               the existing service.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a VPN service. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               service.
        :param pulumi.Input[str] router_id: The ID of the router. Changing this creates a new service.
        :param pulumi.Input[str] status: Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
        :param pulumi.Input[str] subnet_id: SubnetID is the ID of the subnet. Default is null.
        :param pulumi.Input[str] tenant_id: The owner of the service. Required if admin wants to
               create a service for another project. Changing this creates a new service.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceState.__new__(_ServiceState)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["description"] = description
        __props__.__dict__["external_v4_ip"] = external_v4_ip
        __props__.__dict__["external_v6_ip"] = external_v6_ip
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["router_id"] = router_id
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["value_specs"] = value_specs
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        """
        The administrative state of the resource. Can either be up(true) or down(false).
        Changing this updates the administrative state of the existing service.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The human-readable description for the service.
        Changing this updates the description of the existing service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalV4Ip")
    def external_v4_ip(self) -> pulumi.Output[str]:
        """
        The read-only external (public) IPv4 address that is used for the VPN service.
        """
        return pulumi.get(self, "external_v4_ip")

    @property
    @pulumi.getter(name="externalV6Ip")
    def external_v6_ip(self) -> pulumi.Output[str]:
        """
        The read-only external (public) IPv6 address that is used for the VPN service.
        """
        return pulumi.get(self, "external_v6_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the service. Changing this updates the name of
        the existing service.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a VPN service. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        service.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[str]:
        """
        The ID of the router. Changing this creates a new service.
        """
        return pulumi.get(self, "router_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        SubnetID is the ID of the subnet. Default is null.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the service. Required if admin wants to
        create a service for another project. Changing this creates a new service.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

