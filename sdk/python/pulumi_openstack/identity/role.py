# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Role(pulumi.CustomResource):
    domain_id: pulumi.Output[str]
    """
    The domain the role belongs to.
    """
    name: pulumi.Output[str]
    """
    The name of the role.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V3 Keystone client.
    If omitted, the `region` argument of the provider is used. Changing this
    creates a new Role.
    """
    def __init__(__self__, resource_name, opts=None, domain_id=None, name=None, region=None, __name__=None, __opts__=None):
        """
        Manages a V3 Role resource within OpenStack Keystone.
        
        Note: You _must_ have admin privileges in your OpenStack cloud to use
        this resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: The domain the role belongs to.
        :param pulumi.Input[str] name: The name of the role.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new Role.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_role_v3.html.markdown.
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

        __props__['domain_id'] = domain_id

        __props__['name'] = name

        __props__['region'] = region

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Role, __self__).__init__(
            'openstack:identity/role:Role',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

