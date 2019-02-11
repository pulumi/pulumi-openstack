# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class L7RuleV2(pulumi.CustomResource):
    admin_state_up: pulumi.Output[bool]
    """
    The administrative state of the L7 Rule.
    A valid value is true (UP) or false (DOWN).
    """
    compare_type: pulumi.Output[str]
    """
    The comparison type for the L7 rule - can either be
    CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
    """
    invert: pulumi.Output[bool]
    """
    When true the logic of the rule is inverted. For example, with invert
    true, equal to would become not equal to. Default is false.
    """
    key: pulumi.Output[str]
    """
    The key to use for the comparison. For example, the name of the cookie to
    evaluate. Valid when `type` is set to COOKIE or HEADER.
    """
    l7policy_id: pulumi.Output[str]
    """
    The ID of the L7 Policy to query. Changing this creates a new
    L7 Rule.
    """
    listener_id: pulumi.Output[str]
    """
    The ID of the Listener owning this resource.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create an . If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    L7 Rule.
    """
    tenant_id: pulumi.Output[str]
    """
    Required for admins. The UUID of the tenant who owns
    the L7 Rule.  Only administrative users can specify a tenant UUID
    other than their own. Changing this creates a new L7 Rule.
    """
    type: pulumi.Output[str]
    """
    The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
    HOST\_NAME or PATH.
    """
    value: pulumi.Output[str]
    """
    The value to use for the comparison. For example, the file type to
    compare.
    """
    def __init__(__self__, __name__, __opts__=None, admin_state_up=None, compare_type=None, invert=None, key=None, l7policy_id=None, region=None, tenant_id=None, type=None, value=None):
        """
        Manages a V2 L7 Rule resource within OpenStack.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the L7 Rule.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[str] compare_type: The comparison type for the L7 rule - can either be
               CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        :param pulumi.Input[bool] invert: When true the logic of the rule is inverted. For example, with invert
               true, equal to would become not equal to. Default is false.
        :param pulumi.Input[str] key: The key to use for the comparison. For example, the name of the cookie to
               evaluate. Valid when `type` is set to COOKIE or HEADER.
        :param pulumi.Input[str] l7policy_id: The ID of the L7 Policy to query. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the L7 Rule.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new L7 Rule.
        :param pulumi.Input[str] type: The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
               HOST\_NAME or PATH.
        :param pulumi.Input[str] value: The value to use for the comparison. For example, the file type to
               compare.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['admin_state_up'] = admin_state_up

        if not compare_type:
            raise TypeError('Missing required property compare_type')
        __props__['compare_type'] = compare_type

        __props__['invert'] = invert

        __props__['key'] = key

        if not l7policy_id:
            raise TypeError('Missing required property l7policy_id')
        __props__['l7policy_id'] = l7policy_id

        __props__['region'] = region

        __props__['tenant_id'] = tenant_id

        if not type:
            raise TypeError('Missing required property type')
        __props__['type'] = type

        if not value:
            raise TypeError('Missing required property value')
        __props__['value'] = value

        __props__['listener_id'] = None

        super(L7RuleV2, __self__).__init__(
            'openstack:loadbalancer/l7RuleV2:L7RuleV2',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
