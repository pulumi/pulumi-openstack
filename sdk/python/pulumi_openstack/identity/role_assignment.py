# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RoleAssignment(pulumi.CustomResource):
    domain_id: pulumi.Output[str]
    """
    The domain to assign the role in.
    """
    group_id: pulumi.Output[str]
    """
    The group to assign the role to.
    """
    project_id: pulumi.Output[str]
    """
    The project to assign the role in.
    """
    region: pulumi.Output[str]
    role_id: pulumi.Output[str]
    """
    The role to assign.
    """
    user_id: pulumi.Output[str]
    """
    The user to assign the role to.
    """
    def __init__(__self__, resource_name, opts=None, domain_id=None, group_id=None, project_id=None, region=None, role_id=None, user_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V3 Role assignment within OpenStack Keystone.

        Note: You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_role_assignment_v3.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: The domain to assign the role in.
        :param pulumi.Input[str] group_id: The group to assign the role to.
        :param pulumi.Input[str] project_id: The project to assign the role in.
        :param pulumi.Input[str] role_id: The role to assign.
        :param pulumi.Input[str] user_id: The user to assign the role to.
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

            __props__['domain_id'] = domain_id
            __props__['group_id'] = group_id
            __props__['project_id'] = project_id
            __props__['region'] = region
            if role_id is None:
                raise TypeError("Missing required property 'role_id'")
            __props__['role_id'] = role_id
            __props__['user_id'] = user_id
        super(RoleAssignment, __self__).__init__(
            'openstack:identity/roleAssignment:RoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, domain_id=None, group_id=None, project_id=None, region=None, role_id=None, user_id=None):
        """
        Get an existing RoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_id: The domain to assign the role in.
        :param pulumi.Input[str] group_id: The group to assign the role to.
        :param pulumi.Input[str] project_id: The project to assign the role in.
        :param pulumi.Input[str] role_id: The role to assign.
        :param pulumi.Input[str] user_id: The user to assign the role to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["domain_id"] = domain_id
        __props__["group_id"] = group_id
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["role_id"] = role_id
        __props__["user_id"] = user_id
        return RoleAssignment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

