# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ServerGroup(pulumi.CustomResource):
    members: pulumi.Output[list]
    """
    The instances that are part of this server group.
    """
    name: pulumi.Output[str]
    """
    A unique name for the server group. Changing this creates
    a new server group.
    """
    policies: pulumi.Output[list]
    """
    The set of policies for the server group. All policies
    are mutually exclusive. See the Policies section for more information.
    Changing this creates a new server group.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    If omitted, the `region` argument of the provider is used. Changing
    this creates a new server group.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, name=None, policies=None, region=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a V2 Server Group resource within OpenStack.
        
        ## Policies
        
        * `affinity` - All instances/servers launched in this group will be hosted on
            the same compute node.
        
        * `anti-affinity` - All instances/servers launched in this group will be
            hosted on different compute nodes.
        
        * `soft-affinity` - All instances/servers launched in this group will be hosted
            on the same compute node if possible, but if not possible they
            still will be scheduled instead of failure. To use this policy your
            OpenStack environment should support Compute service API 2.15 or above.
        
        * `soft-anti-affinity` - All instances/servers launched in this group will be
            hosted on different compute nodes if possible, but if not possible they
            still will be scheduled instead of failure. To use this policy your
            OpenStack environment should support Compute service API 2.15 or above.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A unique name for the server group. Changing this creates
               a new server group.
        :param pulumi.Input[list] policies: The set of policies for the server group. All policies
               are mutually exclusive. See the Policies section for more information.
               Changing this creates a new server group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used. Changing
               this creates a new server group.
        :param pulumi.Input[dict] value_specs: Map of additional options.
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

        __props__['name'] = name

        __props__['policies'] = policies

        __props__['region'] = region

        __props__['value_specs'] = value_specs

        __props__['members'] = None

        super(ServerGroup, __self__).__init__(
            'openstack:compute/serverGroup:ServerGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

