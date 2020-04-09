# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Firewall(pulumi.CustomResource):
    admin_state_up: pulumi.Output[bool]
    """
    Administrative up/down status for the firewall
    (must be "true" or "false" if provided - defaults to "true").
    Changing this updates the `admin_state_up` of an existing firewall.
    """
    associated_routers: pulumi.Output[list]
    """
    Router(s) to associate this firewall instance
    with. Must be a list of strings. Changing this updates the associated routers
    of an existing firewall. Conflicts with `no_routers`.
    """
    description: pulumi.Output[str]
    """
    A description for the firewall. Changing this
    updates the `description` of an existing firewall.
    """
    name: pulumi.Output[str]
    """
    A name for the firewall. Changing this
    updates the `name` of an existing firewall.
    """
    no_routers: pulumi.Output[bool]
    """
    Should this firewall not be associated with any routers
    (must be "true" or "false" if provide - defaults to "false").
    Conflicts with `associated_routers`.
    """
    policy_id: pulumi.Output[str]
    """
    The policy resource id for the firewall. Changing
    this updates the `policy_id` of an existing firewall.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the v1 networking client.
    A networking client is needed to create a firewall. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    firewall.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the floating IP. Required if admin wants
    to create a firewall for another tenant. Changing this creates a new
    firewall.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, admin_state_up=None, associated_routers=None, description=None, name=None, no_routers=None, policy_id=None, region=None, tenant_id=None, value_specs=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a v1 firewall resource within OpenStack.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/fw_firewall_v1.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the firewall
               (must be "true" or "false" if provided - defaults to "true").
               Changing this updates the `admin_state_up` of an existing firewall.
        :param pulumi.Input[list] associated_routers: Router(s) to associate this firewall instance
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
        :param pulumi.Input[dict] value_specs: Map of additional options.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_state_up'] = admin_state_up
            __props__['associated_routers'] = associated_routers
            __props__['description'] = description
            __props__['name'] = name
            __props__['no_routers'] = no_routers
            if policy_id is None:
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
    def get(resource_name, id, opts=None, admin_state_up=None, associated_routers=None, description=None, name=None, no_routers=None, policy_id=None, region=None, tenant_id=None, value_specs=None):
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the firewall
               (must be "true" or "false" if provided - defaults to "true").
               Changing this updates the `admin_state_up` of an existing firewall.
        :param pulumi.Input[list] associated_routers: Router(s) to associate this firewall instance
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
        :param pulumi.Input[dict] value_specs: Map of additional options.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

