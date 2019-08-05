# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    audited: pulumi.Output[bool]
    """
    Audit status of the firewall policy
    (must be "true" or "false" if provided - defaults to "false").
    This status is set to "false" whenever the firewall policy or any of its
    rules are changed. Changing this updates the `audited` status of an existing
    firewall policy.
    """
    description: pulumi.Output[str]
    """
    A description for the firewall policy. Changing
    this updates the `description` of an existing firewall policy.
    """
    name: pulumi.Output[str]
    """
    A name for the firewall policy. Changing this
    updates the `name` of an existing firewall policy.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the v1 networking client.
    A networking client is needed to create a firewall policy. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    firewall policy.
    """
    rules: pulumi.Output[list]
    """
    An array of one or more firewall rules that comprise
    the policy. Changing this results in adding/removing rules from the
    existing firewall policy.
    """
    shared: pulumi.Output[bool]
    """
    Sharing status of the firewall policy (must be "true"
    or "false" if provided). If this is "true" the policy is visible to, and
    can be used in, firewalls in other tenants. Changing this updates the
    `shared` status of an existing firewall policy. Only administrative users
    can specify if the policy should be shared.
    """
    tenant_id: pulumi.Output[str]
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, audited=None, description=None, name=None, region=None, rules=None, shared=None, tenant_id=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a v1 firewall policy resource within OpenStack.
        
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
        :param pulumi.Input[list] rules: An array of one or more firewall rules that comprise
               the policy. Changing this results in adding/removing rules from the
               existing firewall policy.
        :param pulumi.Input[bool] shared: Sharing status of the firewall policy (must be "true"
               or "false" if provided). If this is "true" the policy is visible to, and
               can be used in, firewalls in other tenants. Changing this updates the
               `shared` status of an existing firewall policy. Only administrative users
               can specify if the policy should be shared.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/fw_policy_v1.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['audited'] = audited

        __props__['description'] = description

        __props__['name'] = name

        __props__['region'] = region

        __props__['rules'] = rules

        __props__['shared'] = shared

        __props__['tenant_id'] = tenant_id

        __props__['value_specs'] = value_specs

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Policy, __self__).__init__(
            'openstack:firewall/policy:Policy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

