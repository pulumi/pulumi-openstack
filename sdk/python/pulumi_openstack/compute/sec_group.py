# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SecGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description for the security group. Changing this
    updates the `description` of an existing security group.
    """
    name: pulumi.Output[str]
    """
    A unique name for the security group. Changing this
    updates the `name` of an existing security group.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    A Compute client is needed to create a security group. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    security group.
    """
    rules: pulumi.Output[list]
    """
    A rule describing how the security group operates. The
    rule object structure is documented below. Changing this updates the
    security group rules. As shown in the example above, multiple rule blocks
    may be used.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, region=None, rules=None, __name__=None, __opts__=None):
        """
        Manages a V2 security group resource within OpenStack.
        
        Please note that managing security groups through the OpenStack Compute API
        has been deprecated. Unless you are using an older OpenStack environment, it is
        recommended to use the `openstack_networking_secgroup_v2`
        and `openstack_networking_secgroup_rule_v2`
        resources instead, which uses the OpenStack Networking API.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the security group. Changing this
               updates the `description` of an existing security group.
        :param pulumi.Input[str] name: A unique name for the security group. Changing this
               updates the `name` of an existing security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a security group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[list] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_secgroup_v2.html.markdown.
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

        if description is None:
            raise TypeError("Missing required property 'description'")
        __props__['description'] = description

        __props__['name'] = name

        __props__['region'] = region

        __props__['rules'] = rules

        super(SecGroup, __self__).__init__(
            'openstack:compute/secGroup:SecGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

