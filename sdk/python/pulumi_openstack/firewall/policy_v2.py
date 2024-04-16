# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PolicyV2Args', 'PolicyV2']

@pulumi.input_type
class PolicyV2Args:
    def __init__(__self__, *,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PolicyV2 resource.
        :param pulumi.Input[bool] audited: Audit status of the firewall policy
               (must be "true" or "false" if provided - defaults to "false").
               This status is set to "false" whenever the firewall policy or any of its
               rules are changed. Changing this updates the `audited` status of an existing
               firewall policy.
        :param pulumi.Input[str] description: A description for the firewall policy. Changing
               this updates the `description` of an existing firewall policy.
        :param pulumi.Input[str] name: A name for the firewall policy. Changing this
               updates the `name` of an existing firewall policy.
        :param pulumi.Input[str] project_id: This argument conflicts and is interchangeable
               with `tenant_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another project. Changing this creates a new
               firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v2 networking client.
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
        :param pulumi.Input[str] tenant_id: This argument conflicts and is interchangeable
               with `project_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another tenant. Changing this creates a new
               firewall policy.
        """
        if audited is not None:
            pulumi.set(__self__, "audited", audited)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        This argument conflicts and is interchangeable
        with `tenant_id`. The owner of the firewall policy. Required if admin wants
        to create a firewall policy for another project. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v2 networking client.
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
        """
        This argument conflicts and is interchangeable
        with `project_id`. The owner of the firewall policy. Required if admin wants
        to create a firewall policy for another tenant. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _PolicyV2State:
    def __init__(__self__, *,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyV2 resources.
        :param pulumi.Input[bool] audited: Audit status of the firewall policy
               (must be "true" or "false" if provided - defaults to "false").
               This status is set to "false" whenever the firewall policy or any of its
               rules are changed. Changing this updates the `audited` status of an existing
               firewall policy.
        :param pulumi.Input[str] description: A description for the firewall policy. Changing
               this updates the `description` of an existing firewall policy.
        :param pulumi.Input[str] name: A name for the firewall policy. Changing this
               updates the `name` of an existing firewall policy.
        :param pulumi.Input[str] project_id: This argument conflicts and is interchangeable
               with `tenant_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another project. Changing this creates a new
               firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v2 networking client.
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
        :param pulumi.Input[str] tenant_id: This argument conflicts and is interchangeable
               with `project_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another tenant. Changing this creates a new
               firewall policy.
        """
        if audited is not None:
            pulumi.set(__self__, "audited", audited)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        This argument conflicts and is interchangeable
        with `tenant_id`. The owner of the firewall policy. Required if admin wants
        to create a firewall policy for another project. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v2 networking client.
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
        """
        This argument conflicts and is interchangeable
        with `project_id`. The owner of the firewall policy. Required if admin wants
        to create a firewall policy for another tenant. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class PolicyV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a v2 firewall policy resource within OpenStack.

        > **Note:** Firewall v2 has no support for OVN currently.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.RuleV2("rule_1",
            name="firewall_rule_1",
            description="drop TELNET traffic",
            action="deny",
            protocol="tcp",
            destination_port="23",
            enabled=True)
        rule2 = openstack.firewall.RuleV2("rule_2",
            name="firewall_rule_2",
            description="drop NTP traffic",
            action="deny",
            protocol="udp",
            destination_port="123",
            enabled=False)
        policy1 = openstack.firewall.PolicyV2("policy_1",
            name="firewall_policy",
            rules=[
                rule1.id,
                rule2.id,
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Firewall Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:firewall/policyV2:PolicyV2 policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
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
        :param pulumi.Input[str] project_id: This argument conflicts and is interchangeable
               with `tenant_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another project. Changing this creates a new
               firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v2 networking client.
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
        :param pulumi.Input[str] tenant_id: This argument conflicts and is interchangeable
               with `project_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another tenant. Changing this creates a new
               firewall policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PolicyV2Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a v2 firewall policy resource within OpenStack.

        > **Note:** Firewall v2 has no support for OVN currently.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.RuleV2("rule_1",
            name="firewall_rule_1",
            description="drop TELNET traffic",
            action="deny",
            protocol="tcp",
            destination_port="23",
            enabled=True)
        rule2 = openstack.firewall.RuleV2("rule_2",
            name="firewall_rule_2",
            description="drop NTP traffic",
            action="deny",
            protocol="udp",
            destination_port="123",
            enabled=False)
        policy1 = openstack.firewall.PolicyV2("policy_1",
            name="firewall_policy",
            rules=[
                rule1.id,
                rule2.id,
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Firewall Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:firewall/policyV2:PolicyV2 policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0
        ```

        :param str resource_name: The name of the resource.
        :param PolicyV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audited: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyV2Args.__new__(PolicyV2Args)

            __props__.__dict__["audited"] = audited
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["rules"] = rules
            __props__.__dict__["shared"] = shared
            __props__.__dict__["tenant_id"] = tenant_id
        super(PolicyV2, __self__).__init__(
            'openstack:firewall/policyV2:PolicyV2',
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
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'PolicyV2':
        """
        Get an existing PolicyV2 resource's state with the given name, id, and optional extra
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
        :param pulumi.Input[str] project_id: This argument conflicts and is interchangeable
               with `tenant_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another project. Changing this creates a new
               firewall policy.
        :param pulumi.Input[str] region: The region in which to obtain the v2 networking client.
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
        :param pulumi.Input[str] tenant_id: This argument conflicts and is interchangeable
               with `project_id`. The owner of the firewall policy. Required if admin wants
               to create a firewall policy for another tenant. Changing this creates a new
               firewall policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyV2State.__new__(_PolicyV2State)

        __props__.__dict__["audited"] = audited
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["rules"] = rules
        __props__.__dict__["shared"] = shared
        __props__.__dict__["tenant_id"] = tenant_id
        return PolicyV2(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        This argument conflicts and is interchangeable
        with `tenant_id`. The owner of the firewall policy. Required if admin wants
        to create a firewall policy for another project. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the v2 networking client.
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
        """
        This argument conflicts and is interchangeable
        with `project_id`. The owner of the firewall policy. Required if admin wants
        to create a firewall policy for another tenant. Changing this creates a new
        firewall policy.
        """
        return pulumi.get(self, "tenant_id")

