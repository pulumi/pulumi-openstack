# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class User(pulumi.CustomResource):
    databases: pulumi.Output[list]
    """
    A list of database user should have access to.
    """
    host: pulumi.Output[str]
    instance_id: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    A unique name for the resource.
    """
    password: pulumi.Output[str]
    """
    User's password.
    """
    region: pulumi.Output[str]
    """
    Openstack region resource is created in.
    """
    def __init__(__self__, resource_name, opts=None, databases=None, host=None, instance_id=None, name=None, password=None, region=None, __name__=None, __opts__=None):
        """
        Manages a V1 DB user resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] databases: A list of database user should have access to.
        :param pulumi.Input[str] name: A unique name for the resource.
        :param pulumi.Input[str] password: User's password.
        :param pulumi.Input[str] region: Openstack region resource is created in.
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

        __props__['databases'] = databases

        __props__['host'] = host

        if instance_id is None:
            raise TypeError("Missing required property 'instance_id'")
        __props__['instance_id'] = instance_id

        __props__['name'] = name

        if password is None:
            raise TypeError("Missing required property 'password'")
        __props__['password'] = password

        __props__['region'] = region

        super(User, __self__).__init__(
            'openstack:database/user:User',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

