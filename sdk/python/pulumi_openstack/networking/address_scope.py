# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AddressScope(pulumi.CustomResource):
    ip_version: pulumi.Output[float]
    """
    IP version, either 4 (default) or 6. Changing this
    creates a new address-scope.
    """
    name: pulumi.Output[str]
    """
    The name of the address-scope. Changing this updates the
    name of the existing address-scope.
    """
    project_id: pulumi.Output[str]
    """
    The owner of the address-scope. Required if admin
    wants to create a address-scope for another project. Changing this creates a
    new address-scope.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron address-scope. If omitted,
    the `region` argument of the provider is used. Changing this creates a new
    address-scope.
    """
    shared: pulumi.Output[bool]
    """
    Indicates whether this address-scope is shared across
    all projects. Changing this updates the shared status of the existing
    address-scope.
    """
    def __init__(__self__, resource_name, opts=None, ip_version=None, name=None, project_id=None, region=None, shared=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron addressscope resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] ip_version: IP version, either 4 (default) or 6. Changing this
               creates a new address-scope.
        :param pulumi.Input[str] name: The name of the address-scope. Changing this updates the
               name of the existing address-scope.
        :param pulumi.Input[str] project_id: The owner of the address-scope. Required if admin
               wants to create a address-scope for another project. Changing this creates a
               new address-scope.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron address-scope. If omitted,
               the `region` argument of the provider is used. Changing this creates a new
               address-scope.
        :param pulumi.Input[bool] shared: Indicates whether this address-scope is shared across
               all projects. Changing this updates the shared status of the existing
               address-scope.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_addressscope_v2.html.markdown.
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

        __props__['ip_version'] = ip_version

        __props__['name'] = name

        __props__['project_id'] = project_id

        __props__['region'] = region

        __props__['shared'] = shared

        super(AddressScope, __self__).__init__(
            'openstack:networking/addressScope:AddressScope',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

