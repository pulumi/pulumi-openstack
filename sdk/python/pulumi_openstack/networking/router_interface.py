# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RouterInterface(pulumi.CustomResource):
    port_id: pulumi.Output[str]
    """
    ID of the port this interface connects to. Changing
    this creates a new router interface.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 networking client.
    A networking client is needed to create a router. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    router interface.
    """
    router_id: pulumi.Output[str]
    """
    ID of the router this interface belongs to. Changing
    this creates a new router interface.
    """
    subnet_id: pulumi.Output[str]
    """
    ID of the subnet this interface connects to. Changing
    this creates a new router interface.
    """
    def __init__(__self__, resource_name, opts=None, port_id=None, region=None, router_id=None, subnet_id=None, __name__=None, __opts__=None):
        """
        Manages a V2 router interface resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] port_id: ID of the port this interface connects to. Changing
               this creates a new router interface.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router interface.
        :param pulumi.Input[str] router_id: ID of the router this interface belongs to. Changing
               this creates a new router interface.
        :param pulumi.Input[str] subnet_id: ID of the subnet this interface connects to. Changing
               this creates a new router interface.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_router_interface_v2.html.markdown.
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

        __props__['port_id'] = port_id

        __props__['region'] = region

        if router_id is None:
            raise TypeError("Missing required property 'router_id'")
        __props__['router_id'] = router_id

        __props__['subnet_id'] = subnet_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RouterInterface, __self__).__init__(
            'openstack:networking/routerInterface:RouterInterface',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

