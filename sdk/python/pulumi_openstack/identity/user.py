# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class User(pulumi.CustomResource):
    """
    Manages a V3 User resource within OpenStack Keystone.
    
    Note: You _must_ have admin privileges in your OpenStack cloud to use
    this resource.
    """
    def __init__(__self__, __name__, __opts__=None, default_project_id=None, description=None, domain_id=None, enabled=None, extra=None, ignore_change_password_upon_first_use=None, ignore_lockout_failure_attempts=None, ignore_password_expiry=None, multi_factor_auth_enabled=None, multi_factor_auth_rules=None, name=None, password=None, region=None):
        """Create a User resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['default_project_id'] = default_project_id

        __props__['description'] = description

        __props__['domain_id'] = domain_id

        __props__['enabled'] = enabled

        __props__['extra'] = extra

        __props__['ignore_change_password_upon_first_use'] = ignore_change_password_upon_first_use

        __props__['ignore_lockout_failure_attempts'] = ignore_lockout_failure_attempts

        __props__['ignore_password_expiry'] = ignore_password_expiry

        __props__['multi_factor_auth_enabled'] = multi_factor_auth_enabled

        __props__['multi_factor_auth_rules'] = multi_factor_auth_rules

        __props__['name'] = name

        __props__['password'] = password

        __props__['region'] = region

        super(User, __self__).__init__(
            'openstack:identity/user:User',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

