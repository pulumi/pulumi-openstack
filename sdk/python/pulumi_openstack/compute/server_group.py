# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServerGroupArgs', 'ServerGroup']

@pulumi.input_type
class ServerGroupArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input['ServerGroupRulesArgs']] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a ServerGroup resource.
        :param pulumi.Input[str] name: A unique name for the server group. Changing this creates
               a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of exactly one policy name to associate with
               the server group. See the Policies section for more information. Changing this
               creates a new server group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used. Changing
               this creates a new server group.
        :param pulumi.Input['ServerGroupRulesArgs'] rules: The rules which are applied to specified `policy`. Currently,
               only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the server group. Changing this creates
        a new server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of exactly one policy name to associate with
        the server group. See the Policies section for more information. Changing this
        creates a new server group.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        If omitted, the `region` argument of the provider is used. Changing
        this creates a new server group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input['ServerGroupRulesArgs']]:
        """
        The rules which are applied to specified `policy`. Currently,
        only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input['ServerGroupRulesArgs']]):
        pulumi.set(self, "rules", value)

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
class _ServerGroupState:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input['ServerGroupRulesArgs']] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering ServerGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: The instances that are part of this server group.
        :param pulumi.Input[str] name: A unique name for the server group. Changing this creates
               a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of exactly one policy name to associate with
               the server group. See the Policies section for more information. Changing this
               creates a new server group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used. Changing
               this creates a new server group.
        :param pulumi.Input['ServerGroupRulesArgs'] rules: The rules which are applied to specified `policy`. Currently,
               only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The instances that are part of this server group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the server group. Changing this creates
        a new server group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of exactly one policy name to associate with
        the server group. See the Policies section for more information. Changing this
        creates a new server group.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        If omitted, the `region` argument of the provider is used. Changing
        this creates a new server group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input['ServerGroupRulesArgs']]:
        """
        The rules which are applied to specified `policy`. Currently,
        only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input['ServerGroupRulesArgs']]):
        pulumi.set(self, "rules", value)

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


class ServerGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[pulumi.InputType['ServerGroupRulesArgs']]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a V2 Server Group resource within OpenStack.

        ## Example Usage

        ### Compute service API version 2.63 or below:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_sg = openstack.compute.ServerGroup("test-sg",
            name="my-sg",
            policies=["anti-affinity"])
        ```

        ### Compute service API version 2.64 or above:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_sg = openstack.compute.ServerGroup("test-sg",
            name="my-sg",
            policies=["anti-affinity"],
            rules=openstack.compute.ServerGroupRulesArgs(
                max_server_per_host=3,
            ))
        ```

        ## Policies

        * `affinity` - All instances/servers launched in this group will be hosted on
          the same compute node.

        * `anti-affinity` - All instances/servers launched in this group will be
          hosted on different compute nodes.

        * `soft-affinity` - All instances/servers launched in this group will be hosted
          on the same compute node if possible, but if not possible they
          still will be scheduled instead of failure. To use this policy your
          OpenStack environment should support Compute service API 2.15 or above.

        * `soft-anti-affinity` - All instances/servers launched in this group will be
          hosted on different compute nodes if possible, but if not possible they
          still will be scheduled instead of failure. To use this policy your
          OpenStack environment should support Compute service API 2.15 or above.

        ## Import

        Server Groups can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:compute/serverGroup:ServerGroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A unique name for the server group. Changing this creates
               a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of exactly one policy name to associate with
               the server group. See the Policies section for more information. Changing this
               creates a new server group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used. Changing
               this creates a new server group.
        :param pulumi.Input[pulumi.InputType['ServerGroupRulesArgs']] rules: The rules which are applied to specified `policy`. Currently,
               only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServerGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 Server Group resource within OpenStack.

        ## Example Usage

        ### Compute service API version 2.63 or below:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_sg = openstack.compute.ServerGroup("test-sg",
            name="my-sg",
            policies=["anti-affinity"])
        ```

        ### Compute service API version 2.64 or above:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_sg = openstack.compute.ServerGroup("test-sg",
            name="my-sg",
            policies=["anti-affinity"],
            rules=openstack.compute.ServerGroupRulesArgs(
                max_server_per_host=3,
            ))
        ```

        ## Policies

        * `affinity` - All instances/servers launched in this group will be hosted on
          the same compute node.

        * `anti-affinity` - All instances/servers launched in this group will be
          hosted on different compute nodes.

        * `soft-affinity` - All instances/servers launched in this group will be hosted
          on the same compute node if possible, but if not possible they
          still will be scheduled instead of failure. To use this policy your
          OpenStack environment should support Compute service API 2.15 or above.

        * `soft-anti-affinity` - All instances/servers launched in this group will be
          hosted on different compute nodes if possible, but if not possible they
          still will be scheduled instead of failure. To use this policy your
          OpenStack environment should support Compute service API 2.15 or above.

        ## Import

        Server Groups can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:compute/serverGroup:ServerGroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
        ```

        :param str resource_name: The name of the resource.
        :param ServerGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[pulumi.InputType['ServerGroupRulesArgs']]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerGroupArgs.__new__(ServerGroupArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["policies"] = policies
            __props__.__dict__["region"] = region
            __props__.__dict__["rules"] = rules
            __props__.__dict__["value_specs"] = value_specs
            __props__.__dict__["members"] = None
        super(ServerGroup, __self__).__init__(
            'openstack:compute/serverGroup:ServerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[pulumi.InputType['ServerGroupRulesArgs']]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'ServerGroup':
        """
        Get an existing ServerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: The instances that are part of this server group.
        :param pulumi.Input[str] name: A unique name for the server group. Changing this creates
               a new server group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of exactly one policy name to associate with
               the server group. See the Policies section for more information. Changing this
               creates a new server group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used. Changing
               this creates a new server group.
        :param pulumi.Input[pulumi.InputType['ServerGroupRulesArgs']] rules: The rules which are applied to specified `policy`. Currently,
               only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerGroupState.__new__(_ServerGroupState)

        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["policies"] = policies
        __props__.__dict__["region"] = region
        __props__.__dict__["rules"] = rules
        __props__.__dict__["value_specs"] = value_specs
        return ServerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        The instances that are part of this server group.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the server group. Changing this creates
        a new server group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of exactly one policy name to associate with
        the server group. See the Policies section for more information. Changing this
        creates a new server group.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        If omitted, the `region` argument of the provider is used. Changing
        this creates a new server group.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output['outputs.ServerGroupRules']:
        """
        The rules which are applied to specified `policy`. Currently,
        only the `max_server_per_host` rule is supported for the `anti-affinity` policy.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

