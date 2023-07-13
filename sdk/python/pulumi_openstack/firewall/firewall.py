# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
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
        """
        pulumi.set(__self__, "policy_id", policy_id)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if associated_routers is not None:
            pulumi.set(__self__, "associated_routers", associated_routers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_routers is not None:
            pulumi.set(__self__, "no_routers", no_routers)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="associatedRouters")
    def associated_routers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "associated_routers")

    @associated_routers.setter
    def associated_routers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associated_routers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noRouters")
    def no_routers(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_routers")

    @no_routers.setter
    def no_routers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_routers", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

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
        """
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if associated_routers is not None:
            pulumi.set(__self__, "associated_routers", associated_routers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_routers is not None:
            pulumi.set(__self__, "no_routers", no_routers)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="associatedRouters")
    def associated_routers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "associated_routers")

    @associated_routers.setter
    def associated_routers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associated_routers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noRouters")
    def no_routers(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_routers")

    @no_routers.setter
    def no_routers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_routers", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

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
        Create a Firewall resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Firewall resource with the given unique name, props, and options.
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
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="associatedRouters")
    def associated_routers(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "associated_routers")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noRouters")
    def no_routers(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "no_routers")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "value_specs")

