# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input[bool] audited: Audit status of the firewall policy
               (must be "true" or "false" if provided - defaults to "false").
               This status is set to "false" whenever the firewall policy or any of its
               rules are changed. Changing this updates the `audited` status of an existing
               firewall policy.
        :param pulumi.Input[str] description: A description for the firewall policy. Changing
               this updates the `description` of an existing firewall policy.
        :param pulumi.Input[str] name: A name for the firewall policy. Changing this
               updates the `name` of an existing firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rules: An array of one or more firewall rules that comprise
               the policy. Changing this results in adding/removing rules from the
               existing firewall policy.
        :param pulumi.Input[bool] shared: Sharing status of the firewall policy (must be "true"
               or "false" if provided). If this is "true" the policy is visible to, and
               can be used in, firewalls in other tenants. Changing this updates the
               `shared` status of an existing firewall policy. Only administrative users
               can specify if the policy should be shared.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        PolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            audited=audited,
            description=description,
            name=name,
            region=region,
            rules=rules,
            shared=shared,
            tenant_id=tenant_id,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             audited: Optional[pulumi.Input[bool]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             shared: Optional[pulumi.Input[bool]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if audited is not None:
            _setter("audited", audited)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)
        if rules is not None:
            _setter("rules", rules)
        if shared is not None:
            _setter("shared", shared)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter
    def audited(self) -> Optional[pulumi.Input[bool]]:
        """
        Audit status of the firewall policy
        (must be "true" or "false" if provided - defaults to "false").
        This status is set to "false" whenever the firewall policy or any of its
        rules are changed. Changing this updates the `audited` status of an existing
        firewall policy.
        """
        return pulumi.get(self, "audited")

    @audited.setter
    def audited(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "audited", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the firewall policy. Changing
        this updates the `description` of an existing firewall policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the firewall policy. Changing this
        updates the `name` of an existing firewall policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of one or more firewall rules that comprise
        the policy. Changing this results in adding/removing rules from the
        existing firewall policy.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Sharing status of the firewall policy (must be "true"
        or "false" if provided). If this is "true" the policy is visible to, and
        can be used in, firewalls in other tenants. Changing this updates the
        `shared` status of an existing firewall policy. Only administrative users
        can specify if the policy should be shared.
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
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
class _PolicyState:
    def __init__(__self__, *,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Policy resources.
        :param pulumi.Input[bool] audited: Audit status of the firewall policy
               (must be "true" or "false" if provided - defaults to "false").
               This status is set to "false" whenever the firewall policy or any of its
               rules are changed. Changing this updates the `audited` status of an existing
               firewall policy.
        :param pulumi.Input[str] description: A description for the firewall policy. Changing
               this updates the `description` of an existing firewall policy.
        :param pulumi.Input[str] name: A name for the firewall policy. Changing this
               updates the `name` of an existing firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rules: An array of one or more firewall rules that comprise
               the policy. Changing this results in adding/removing rules from the
               existing firewall policy.
        :param pulumi.Input[bool] shared: Sharing status of the firewall policy (must be "true"
               or "false" if provided). If this is "true" the policy is visible to, and
               can be used in, firewalls in other tenants. Changing this updates the
               `shared` status of an existing firewall policy. Only administrative users
               can specify if the policy should be shared.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        _PolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            audited=audited,
            description=description,
            name=name,
            region=region,
            rules=rules,
            shared=shared,
            tenant_id=tenant_id,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             audited: Optional[pulumi.Input[bool]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             shared: Optional[pulumi.Input[bool]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if audited is not None:
            _setter("audited", audited)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)
        if rules is not None:
            _setter("rules", rules)
        if shared is not None:
            _setter("shared", shared)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter
    def audited(self) -> Optional[pulumi.Input[bool]]:
        """
        Audit status of the firewall policy
        (must be "true" or "false" if provided - defaults to "false").
        This status is set to "false" whenever the firewall policy or any of its
        rules are changed. Changing this updates the `audited` status of an existing
        firewall policy.
        """
        return pulumi.get(self, "audited")

    @audited.setter
    def audited(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "audited", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the firewall policy. Changing
        this updates the `description` of an existing firewall policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the firewall policy. Changing this
        updates the `name` of an existing firewall policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of one or more firewall rules that comprise
        the policy. Changing this results in adding/removing rules from the
        existing firewall policy.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Sharing status of the firewall policy (must be "true"
        or "false" if provided). If this is "true" the policy is visible to, and
        can be used in, firewalls in other tenants. Changing this updates the
        `shared` status of an existing firewall policy. Only administrative users
        can specify if the policy should be shared.
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
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


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a v1 firewall policy resource within OpenStack.

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
        ```

        ## Import

        Firewall Policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:firewall/policy:Policy policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] audited: Audit status of the firewall policy
               (must be "true" or "false" if provided - defaults to "false").
               This status is set to "false" whenever the firewall policy or any of its
               rules are changed. Changing this updates the `audited` status of an existing
               firewall policy.
        :param pulumi.Input[str] description: A description for the firewall policy. Changing
               this updates the `description` of an existing firewall policy.
        :param pulumi.Input[str] name: A name for the firewall policy. Changing this
               updates the `name` of an existing firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rules: An array of one or more firewall rules that comprise
               the policy. Changing this results in adding/removing rules from the
               existing firewall policy.
        :param pulumi.Input[bool] shared: Sharing status of the firewall policy (must be "true"
               or "false" if provided). If this is "true" the policy is visible to, and
               can be used in, firewalls in other tenants. Changing this updates the
               `shared` status of an existing firewall policy. Only administrative users
               can specify if the policy should be shared.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a v1 firewall policy resource within OpenStack.

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
        ```

        ## Import

        Firewall Policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:firewall/policy:Policy policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            PolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyArgs.__new__(PolicyArgs)

            __props__.__dict__["audited"] = audited
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["rules"] = rules
            __props__.__dict__["shared"] = shared
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["value_specs"] = value_specs
        super(Policy, __self__).__init__(
            'openstack:firewall/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            audited: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] audited: Audit status of the firewall policy
               (must be "true" or "false" if provided - defaults to "false").
               This status is set to "false" whenever the firewall policy or any of its
               rules are changed. Changing this updates the `audited` status of an existing
               firewall policy.
        :param pulumi.Input[str] description: A description for the firewall policy. Changing
               this updates the `description` of an existing firewall policy.
        :param pulumi.Input[str] name: A name for the firewall policy. Changing this
               updates the `name` of an existing firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v1 networking client.
               A networking client is needed to create a firewall policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rules: An array of one or more firewall rules that comprise
               the policy. Changing this results in adding/removing rules from the
               existing firewall policy.
        :param pulumi.Input[bool] shared: Sharing status of the firewall policy (must be "true"
               or "false" if provided). If this is "true" the policy is visible to, and
               can be used in, firewalls in other tenants. Changing this updates the
               `shared` status of an existing firewall policy. Only administrative users
               can specify if the policy should be shared.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyState.__new__(_PolicyState)

        __props__.__dict__["audited"] = audited
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["rules"] = rules
        __props__.__dict__["shared"] = shared
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["value_specs"] = value_specs
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def audited(self) -> pulumi.Output[Optional[bool]]:
        """
        Audit status of the firewall policy
        (must be "true" or "false" if provided - defaults to "false").
        This status is set to "false" whenever the firewall policy or any of its
        rules are changed. Changing this updates the `audited` status of an existing
        firewall policy.
        """
        return pulumi.get(self, "audited")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the firewall policy. Changing
        this updates the `description` of an existing firewall policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A name for the firewall policy. Changing this
        updates the `name` of an existing firewall policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the v1 networking client.
        A networking client is needed to create a firewall policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of one or more firewall rules that comprise
        the policy. Changing this results in adding/removing rules from the
        existing firewall policy.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[Optional[bool]]:
        """
        Sharing status of the firewall policy (must be "true"
        or "false" if provided). If this is "true" the policy is visible to, and
        can be used in, firewalls in other tenants. Changing this updates the
        `shared` status of an existing firewall policy. Only administrative users
        can specify if the policy should be shared.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

