# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class SubnetPool(pulumi.CustomResource):
    """
    Manages a V2 Neutron subnetpool resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, address_scope_id=None, default_prefixlen=None, default_quota=None, description=None, ip_version=None, is_default=None, max_prefixlen=None, min_prefixlen=None, name=None, prefixes=None, project_id=None, region=None, shared=None, value_specs=None):
        """Create a SubnetPool resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['address_scope_id'] = address_scope_id

        __props__['default_prefixlen'] = default_prefixlen

        __props__['default_quota'] = default_quota

        __props__['description'] = description

        __props__['ip_version'] = ip_version

        __props__['is_default'] = is_default

        __props__['max_prefixlen'] = max_prefixlen

        __props__['min_prefixlen'] = min_prefixlen

        __props__['name'] = name

        if not prefixes:
            raise TypeError('Missing required property prefixes')
        __props__['prefixes'] = prefixes

        __props__['project_id'] = project_id

        __props__['region'] = region

        __props__['shared'] = shared

        __props__['value_specs'] = value_specs

        __props__['created_at'] = None
        __props__['revision_number'] = None
        __props__['updated_at'] = None

        super(SubnetPool, __self__).__init__(
            'openstack:networking/subnetPool:SubnetPool',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

