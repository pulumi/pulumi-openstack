# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EndpointGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The human-readable description for the group.
    Changing this updates the description of the existing group.
    """
    endpoints: pulumi.Output[list]
    """
    List of endpoints of the same type, for the endpoint group. The values will depend on the type.
    Changing this creates a new group.
    """
    name: pulumi.Output[str]
    """
    The name of the group. Changing this updates the name of
    the existing group.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create an endpoint group. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    group.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the group. Required if admin wants to
    create an endpoint group for another project. Changing this creates a new group.
    """
    type: pulumi.Output[str]
    """
    The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
    Changing this creates a new group.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, description=None, endpoints=None, name=None, region=None, tenant_id=None, type=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron Endpoint Group resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The human-readable description for the group.
               Changing this updates the description of the existing group.
        :param pulumi.Input[list] endpoints: List of endpoints of the same type, for the endpoint group. The values will depend on the type.
               Changing this creates a new group.
        :param pulumi.Input[str] name: The name of the group. Changing this updates the name of
               the existing group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an endpoint group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               group.
        :param pulumi.Input[str] tenant_id: The owner of the group. Required if admin wants to
               create an endpoint group for another project. Changing this creates a new group.
        :param pulumi.Input[str] type: The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
               Changing this creates a new group.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/vpnaas_endpoint_group_v2.html.markdown.
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

        __props__['description'] = description

        __props__['endpoints'] = endpoints

        __props__['name'] = name

        __props__['region'] = region

        __props__['tenant_id'] = tenant_id

        __props__['type'] = type

        __props__['value_specs'] = value_specs

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(EndpointGroup, __self__).__init__(
            'openstack:vpnaas/endpointGroup:EndpointGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

