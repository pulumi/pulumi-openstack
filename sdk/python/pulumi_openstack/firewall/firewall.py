# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FirewallArgs', 'Firewall']

@pulumi.input_type
class FirewallArgs:
    def __init__(__self__, *,
                 policy_id: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_routers: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Firewall resource.
        :param pulumi.Input[str] policy_id: The policy resource id for the firewall. Changing
               this updates the `policy_id` of an existing firewall.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the firewall
               (must be "true" or "false" if provided - defaults to "true").
               Changing this updates the `admin_state_up` of an existing firewall.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_routers: Router(s) to associate this firewall instance
               with. Must be a list of strings. Changing this updates the associated routers
               of an existing firewall. Conflicts with `no_routers`.
        :param pulumi.Input[str] description: A description for the firewall. Changing this
               updates the `description` of an existing firewall.
        :param pulumi.Input[str] name: A name for the firewall. Changing this
               updates the `name` of an existing firewall.
        :param pulumi.Input[bool] no_routers: Should this firewall not be associated with any routers
               (must be "true" or "false" if provide - defaults to "false").
               Conflicts with `associated_routers`.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall.
        :param pulumi.Input[str] tenant_id: The owner of the floating IP. Required if admin wants
               to create a firewall for another tenant. Changing this creates a new
               firewall.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        FirewallArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            policy_id=policy_id,
            admin_state_up=admin_state_up,
            associated_routers=associated_routers,
            description=description,
            name=name,
            no_routers=no_routers,
            region=region,
            tenant_id=tenant_id,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             policy_id: Optional[pulumi.Input[str]] = None,
             admin_state_up: Optional[pulumi.Input[bool]] = None,
             associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             no_routers: Optional[pulumi.Input[bool]] = None,
             region: Optional[pulumi.Input[str]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if policy_id is None and 'policyId' in kwargs:
            policy_id = kwargs['policyId']
        if policy_id is None:
            raise TypeError("Missing 'policy_id' argument")
        if admin_state_up is None and 'adminStateUp' in kwargs:
            admin_state_up = kwargs['adminStateUp']
        if associated_routers is None and 'associatedRouters' in kwargs:
            associated_routers = kwargs['associatedRouters']
        if no_routers is None and 'noRouters' in kwargs:
            no_routers = kwargs['noRouters']
        if tenant_id is None and 'tenantId' in kwargs:
            tenant_id = kwargs['tenantId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        _setter("policy_id", policy_id)
        if admin_state_up is not None:
            _setter("admin_state_up", admin_state_up)
        if associated_routers is not None:
            _setter("associated_routers", associated_routers)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if no_routers is not None:
            _setter("no_routers", no_routers)
        if region is not None:
            _setter("region", region)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The policy resource id for the firewall. Changing
        this updates the `policy_id` of an existing firewall.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        Administrative up/down status for the firewall
        (must be "true" or "false" if provided - defaults to "true").
        Changing this updates the `admin_state_up` of an existing firewall.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="associatedRouters")
    def associated_routers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Router(s) to associate this firewall instance
        with. Must be a list of strings. Changing this updates the associated routers
        of an existing firewall. Conflicts with `no_routers`.
        """
        return pulumi.get(self, "associated_routers")

    @associated_routers.setter
    def associated_routers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associated_routers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the firewall. Changing this
        updates the `description` of an existing firewall.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the firewall. Changing this
        updates the `name` of an existing firewall.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noRouters")
    def no_routers(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this firewall not be associated with any routers
        (must be "true" or "false" if provide - defaults to "false").
        Conflicts with `associated_routers`.
        """
        return pulumi.get(self, "no_routers")

    @no_routers.setter
    def no_routers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_routers", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the floating IP. Required if admin wants
        to create a firewall for another tenant. Changing this creates a new
        firewall.
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
class _FirewallState:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_routers: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Firewall resources.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the firewall
               (must be "true" or "false" if provided - defaults to "true").
               Changing this updates the `admin_state_up` of an existing firewall.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_routers: Router(s) to associate this firewall instance
               with. Must be a list of strings. Changing this updates the associated routers
               of an existing firewall. Conflicts with `no_routers`.
        :param pulumi.Input[str] description: A description for the firewall. Changing this
               updates the `description` of an existing firewall.
        :param pulumi.Input[str] name: A name for the firewall. Changing this
               updates the `name` of an existing firewall.
        :param pulumi.Input[bool] no_routers: Should this firewall not be associated with any routers
               (must be "true" or "false" if provide - defaults to "false").
               Conflicts with `associated_routers`.
        :param pulumi.Input[str] policy_id: The policy resource id for the firewall. Changing
               this updates the `policy_id` of an existing firewall.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall.
        :param pulumi.Input[str] tenant_id: The owner of the floating IP. Required if admin wants
               to create a firewall for another tenant. Changing this creates a new
               firewall.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        _FirewallState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            admin_state_up=admin_state_up,
            associated_routers=associated_routers,
            description=description,
            name=name,
            no_routers=no_routers,
            policy_id=policy_id,
            region=region,
            tenant_id=tenant_id,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             admin_state_up: Optional[pulumi.Input[bool]] = None,
             associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             no_routers: Optional[pulumi.Input[bool]] = None,
             policy_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if admin_state_up is None and 'adminStateUp' in kwargs:
            admin_state_up = kwargs['adminStateUp']
        if associated_routers is None and 'associatedRouters' in kwargs:
            associated_routers = kwargs['associatedRouters']
        if no_routers is None and 'noRouters' in kwargs:
            no_routers = kwargs['noRouters']
        if policy_id is None and 'policyId' in kwargs:
            policy_id = kwargs['policyId']
        if tenant_id is None and 'tenantId' in kwargs:
            tenant_id = kwargs['tenantId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        if admin_state_up is not None:
            _setter("admin_state_up", admin_state_up)
        if associated_routers is not None:
            _setter("associated_routers", associated_routers)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if no_routers is not None:
            _setter("no_routers", no_routers)
        if policy_id is not None:
            _setter("policy_id", policy_id)
        if region is not None:
            _setter("region", region)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        Administrative up/down status for the firewall
        (must be "true" or "false" if provided - defaults to "true").
        Changing this updates the `admin_state_up` of an existing firewall.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="associatedRouters")
    def associated_routers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Router(s) to associate this firewall instance
        with. Must be a list of strings. Changing this updates the associated routers
        of an existing firewall. Conflicts with `no_routers`.
        """
        return pulumi.get(self, "associated_routers")

    @associated_routers.setter
    def associated_routers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associated_routers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the firewall. Changing this
        updates the `description` of an existing firewall.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the firewall. Changing this
        updates the `name` of an existing firewall.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noRouters")
    def no_routers(self) -> Optional[pulumi.Input[bool]]:
        """
        Should this firewall not be associated with any routers
        (must be "true" or "false" if provide - defaults to "false").
        Conflicts with `associated_routers`.
        """
        return pulumi.get(self, "no_routers")

    @no_routers.setter
    def no_routers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_routers", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The policy resource id for the firewall. Changing
        this updates the `policy_id` of an existing firewall.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the floating IP. Required if admin wants
        to create a firewall for another tenant. Changing this creates a new
        firewall.
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


class Firewall(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_routers: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a v1 firewall resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.Rule("rule1",
            description="drop TELNET traffic",
            action="deny",
            protocol="tcp",
            destination_port="23",
            enabled=True)
        rule2 = openstack.firewall.Rule("rule2",
            description="drop NTP traffic",
            action="deny",
            protocol="udp",
            destination_port="123",
            enabled=False)
        policy1 = openstack.firewall.Policy("policy1", rules=[
            rule1.id,
            rule2.id,
        ])
        firewall1 = openstack.firewall.Firewall("firewall1", policy_id=policy1.id)
        ```

        ## Import

        Firewalls can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:firewall/firewall:Firewall firewall_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the firewall
               (must be "true" or "false" if provided - defaults to "true").
               Changing this updates the `admin_state_up` of an existing firewall.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_routers: Router(s) to associate this firewall instance
               with. Must be a list of strings. Changing this updates the associated routers
               of an existing firewall. Conflicts with `no_routers`.
        :param pulumi.Input[str] description: A description for the firewall. Changing this
               updates the `description` of an existing firewall.
        :param pulumi.Input[str] name: A name for the firewall. Changing this
               updates the `name` of an existing firewall.
        :param pulumi.Input[bool] no_routers: Should this firewall not be associated with any routers
               (must be "true" or "false" if provide - defaults to "false").
               Conflicts with `associated_routers`.
        :param pulumi.Input[str] policy_id: The policy resource id for the firewall. Changing
               this updates the `policy_id` of an existing firewall.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall.
        :param pulumi.Input[str] tenant_id: The owner of the floating IP. Required if admin wants
               to create a firewall for another tenant. Changing this creates a new
               firewall.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a v1 firewall resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.Rule("rule1",
            description="drop TELNET traffic",
            action="deny",
            protocol="tcp",
            destination_port="23",
            enabled=True)
        rule2 = openstack.firewall.Rule("rule2",
            description="drop NTP traffic",
            action="deny",
            protocol="udp",
            destination_port="123",
            enabled=False)
        policy1 = openstack.firewall.Policy("policy1", rules=[
            rule1.id,
            rule2.id,
        ])
        firewall1 = openstack.firewall.Firewall("firewall1", policy_id=policy1.id)
        ```

        ## Import

        Firewalls can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:firewall/firewall:Firewall firewall_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
        ```

        :param str resource_name: The name of the resource.
        :param FirewallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            FirewallArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_routers: Optional[pulumi.Input[bool]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallArgs.__new__(FirewallArgs)

            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["associated_routers"] = associated_routers
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["no_routers"] = no_routers
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["value_specs"] = value_specs
        super(Firewall, __self__).__init__(
            'openstack:firewall/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            associated_routers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            no_routers: Optional[pulumi.Input[bool]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Firewall':
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the firewall
               (must be "true" or "false" if provided - defaults to "true").
               Changing this updates the `admin_state_up` of an existing firewall.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_routers: Router(s) to associate this firewall instance
               with. Must be a list of strings. Changing this updates the associated routers
               of an existing firewall. Conflicts with `no_routers`.
        :param pulumi.Input[str] description: A description for the firewall. Changing this
               updates the `description` of an existing firewall.
        :param pulumi.Input[str] name: A name for the firewall. Changing this
               updates the `name` of an existing firewall.
        :param pulumi.Input[bool] no_routers: Should this firewall not be associated with any routers
               (must be "true" or "false" if provide - defaults to "false").
               Conflicts with `associated_routers`.
        :param pulumi.Input[str] policy_id: The policy resource id for the firewall. Changing
               this updates the `policy_id` of an existing firewall.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall.
        :param pulumi.Input[str] tenant_id: The owner of the floating IP. Required if admin wants
               to create a firewall for another tenant. Changing this creates a new
               firewall.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallState.__new__(_FirewallState)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["associated_routers"] = associated_routers
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["no_routers"] = no_routers
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["value_specs"] = value_specs
        return Firewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        """
        Administrative up/down status for the firewall
        (must be "true" or "false" if provided - defaults to "true").
        Changing this updates the `admin_state_up` of an existing firewall.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="associatedRouters")
    def associated_routers(self) -> pulumi.Output[Sequence[str]]:
        """
        Router(s) to associate this firewall instance
        with. Must be a list of strings. Changing this updates the associated routers
        of an existing firewall. Conflicts with `no_routers`.
        """
        return pulumi.get(self, "associated_routers")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the firewall. Changing this
        updates the `description` of an existing firewall.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A name for the firewall. Changing this
        updates the `name` of an existing firewall.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noRouters")
    def no_routers(self) -> pulumi.Output[Optional[bool]]:
        """
        Should this firewall not be associated with any routers
        (must be "true" or "false" if provide - defaults to "false").
        Conflicts with `associated_routers`.
        """
        return pulumi.get(self, "no_routers")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The policy resource id for the firewall. Changing
        this updates the `policy_id` of an existing firewall.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the floating IP. Required if admin wants
        to create a firewall for another tenant. Changing this creates a new
        firewall.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

