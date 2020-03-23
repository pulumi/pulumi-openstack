# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecGroup(pulumi.CustomResource):
    all_tags: pulumi.Output[list]
    """
    The collection of tags assigned on the security group, which have
    been explicitly and implicitly added.
    """
    delete_default_rules: pulumi.Output[bool]
    """
    Whether or not to delete the default
    egress security rules. This is `false` by default. See the below note
    for more information.
    """
    description: pulumi.Output[str]
    """
    A unique name for the security group.
    """
    name: pulumi.Output[str]
    """
    A unique name for the security group.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 networking client.
    A networking client is needed to create a port. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    security group.
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the security group.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the security group. Required if admin
    wants to create a port for another tenant. Changing this creates a new
    security group.
    """
    def __init__(__self__, resource_name, opts=None, delete_default_rules=None, description=None, name=None, region=None, tags=None, tenant_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 neutron security group resource within OpenStack.
        Unlike Nova security groups, neutron separates the group from the rules
        and also allows an admin to target a specific tenant_id.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_secgroup_v2.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_default_rules: Whether or not to delete the default
               egress security rules. This is `false` by default. See the below note
               for more information.
        :param pulumi.Input[str] description: A unique name for the security group.
        :param pulumi.Input[str] name: A unique name for the security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[list] tags: A set of string tags for the security group.
        :param pulumi.Input[str] tenant_id: The owner of the security group. Required if admin
               wants to create a port for another tenant. Changing this creates a new
               security group.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['delete_default_rules'] = delete_default_rules
            __props__['description'] = description
            __props__['name'] = name
            __props__['region'] = region
            __props__['tags'] = tags
            __props__['tenant_id'] = tenant_id
            __props__['all_tags'] = None
        super(SecGroup, __self__).__init__(
            'openstack:networking/secGroup:SecGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, all_tags=None, delete_default_rules=None, description=None, name=None, region=None, tags=None, tenant_id=None):
        """
        Get an existing SecGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] all_tags: The collection of tags assigned on the security group, which have
               been explicitly and implicitly added.
        :param pulumi.Input[bool] delete_default_rules: Whether or not to delete the default
               egress security rules. This is `false` by default. See the below note
               for more information.
        :param pulumi.Input[str] description: A unique name for the security group.
        :param pulumi.Input[str] name: A unique name for the security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[list] tags: A set of string tags for the security group.
        :param pulumi.Input[str] tenant_id: The owner of the security group. Required if admin
               wants to create a port for another tenant. Changing this creates a new
               security group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["all_tags"] = all_tags
        __props__["delete_default_rules"] = delete_default_rules
        __props__["description"] = description
        __props__["name"] = name
        __props__["region"] = region
        __props__["tags"] = tags
        __props__["tenant_id"] = tenant_id
        return SecGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

