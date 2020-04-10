# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Project(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description of the project.
    """
    domain_id: pulumi.Output[str]
    """
    The domain this project belongs to.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the project is enabled or disabled. Valid
    values are `true` and `false`.
    """
    is_domain: pulumi.Output[bool]
    """
    Whether this project is a domain. Valid values
    are `true` and `false`.
    """
    name: pulumi.Output[str]
    """
    The name of the project.
    """
    parent_id: pulumi.Output[str]
    """
    The parent of this project.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V3 Keystone client.
    If omitted, the `region` argument of the provider is used. Changing this
    creates a new User.
    """
    def __init__(__self__, resource_name, opts=None, description=None, domain_id=None, enabled=None, is_domain=None, name=None, parent_id=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V3 Project resource within OpenStack Keystone.

        Note: You _must_ have admin privileges in your OpenStack cloud to use
        this resource.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the project.
        :param pulumi.Input[str] domain_id: The domain this project belongs to.
        :param pulumi.Input[bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`.
        :param pulumi.Input[bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`.
        :param pulumi.Input[str] name: The name of the project.
        :param pulumi.Input[str] parent_id: The parent of this project.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new User.
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

            __props__['description'] = description
            __props__['domain_id'] = domain_id
            __props__['enabled'] = enabled
            __props__['is_domain'] = is_domain
            __props__['name'] = name
            __props__['parent_id'] = parent_id
            __props__['region'] = region
        super(Project, __self__).__init__(
            'openstack:identity/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, domain_id=None, enabled=None, is_domain=None, name=None, parent_id=None, region=None):
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the project.
        :param pulumi.Input[str] domain_id: The domain this project belongs to.
        :param pulumi.Input[bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`.
        :param pulumi.Input[bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`.
        :param pulumi.Input[str] name: The name of the project.
        :param pulumi.Input[str] parent_id: The parent of this project.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new User.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["domain_id"] = domain_id
        __props__["enabled"] = enabled
        __props__["is_domain"] = is_domain
        __props__["name"] = name
        __props__["parent_id"] = parent_id
        __props__["region"] = region
        return Project(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

