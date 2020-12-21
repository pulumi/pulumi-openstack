# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Firewall']


class Firewall(pulumi.CustomResource):
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
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a v1 firewall resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.Rule("rule1",
            action="deny",
            description="drop TELNET traffic",
            destination_port="23",
            enabled=True,
            protocol="tcp")
        rule2 = openstack.firewall.Rule("rule2",
            action="deny",
            description="drop NTP traffic",
            destination_port="123",
            enabled=False,
            protocol="udp")
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

            __props__['admin_state_up'] = admin_state_up
            __props__['associated_routers'] = associated_routers
            __props__['description'] = description
            __props__['name'] = name
            __props__['no_routers'] = no_routers
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            __props__['region'] = region
            __props__['tenant_id'] = tenant_id
            __props__['value_specs'] = value_specs
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

        __props__ = dict()

        __props__["admin_state_up"] = admin_state_up
        __props__["associated_routers"] = associated_routers
        __props__["description"] = description
        __props__["name"] = name
        __props__["no_routers"] = no_routers
        __props__["policy_id"] = policy_id
        __props__["region"] = region
        __props__["tenant_id"] = tenant_id
        __props__["value_specs"] = value_specs
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

